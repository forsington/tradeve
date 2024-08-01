package main

import (
	"fmt"
	"os"
	"time"

	"github.com/boltdb/bolt"
	"github.com/forsington/tradeve/analysis"
	"github.com/forsington/tradeve/config"
	"github.com/forsington/tradeve/esi"
	"github.com/forsington/tradeve/history"
	"github.com/forsington/tradeve/item"
	"github.com/forsington/tradeve/order"
	"github.com/hashicorp/go-hclog"
)

func main() {
	conf := config.DefaultConfig()
	conf, err := config.ParseConfigFile("config.json", conf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conf = config.ParseFlags(conf)

	logger := hclog.New(&hclog.LoggerOptions{
		Name:  "tradeve",
		Level: hclog.LevelFromString(conf.LogLevel),
	})
	logger.Debug("config read", "file", "config.json")
	fmt.Printf("config read: \n%s\n", conf.String())

	logger.Info("starting TradEVE")
	esiClient := esi.NewClient("tranquility")

	err = ensureSDEFileExists(conf.ForceSDEDownload, logger)
	if err != nil {
		logger.Error("error ensuring SDE file exists", "error", err)
		os.Exit(0)
	}

	historyRepo := history.NewESIRepository(esiClient)
	orderRepo := order.NewESIRepository(esiClient)

	if conf.CacheEnabled {
		// BoltDB
		// Open the data file in the current directory.
		// It will be created if it doesn't exist.
		db, err := bolt.Open("bolt.db", 0600, nil)
		if err != nil {
			logger.Error("error setting upp boltdb", "error", err)
			os.Exit(1)
		}
		defer db.Close()

		// A history repository that will first check the cache and then the ESI API
		boltHistoryRepo, err := history.NewBoltRepository(db, logger)
		if err != nil {
			logger.Error("error creating boltdb repo", "error", err)
			return
		}
		historyRepo = history.NewCacheFirstRepository(boltHistoryRepo, historyRepo)

		// An order repository that will first check the cache and then the ESI API
		boltOrderRepo, err := order.NewBoltRepository(db, logger)
		if err != nil {
			logger.Error("error creating boltdb repo", "error", err)
			return
		}
		orderRepo = order.NewCacheFirstRepository(boltOrderRepo, orderRepo)
	}

	historyService := history.NewService(logger, historyRepo)
	orderService := order.NewService(orderRepo)

	// Item service
	itemService := item.NewService(logger, orderService, historyService)

	run(logger, esiClient, itemService, conf)
}

func run(logger hclog.Logger, esiClient *esi.Client, itemService *item.Service, conf *config.Configuration) {
	types, err := loadCSV(SDEFilename)
	if err != nil {
		logger.Error("error loading types from CSV", "error", err)
		os.Exit(0)
	}
	logger.Info("found types", "count", len(types))

	hasActiveOrders, err := esiClient.GetMarketTypes(conf.RegionId())
	if err != nil {
		logger.Error("error fetching market types", "error", err)
		os.Exit(0)
	}

	types = types.Active(hasActiveOrders)
	logger.Info("types with active orders in region", "count", len(types), "region", conf.RegionId())

	types = types.ExcludeGroups(conf.ExcludeGroups)
	logger.Info("types after removing exclude groups", "count", len(types))

	start := time.Now()

	depth := len(types)
	estimatedRunTime := (int64(esi.MarketHistoryCallTime) * int64(depth)) + int64(esi.MarketOrdersCallTime)*int64(depth)
	logger.Info("fetching market history and active orders for items", "count", depth)
	logger.Info("estimated run time (if cache empty or first run of the day)", "time", time.Duration(estimatedRunTime))
	items := itemService.GetItems(types, conf.RegionId(), conf.HistoryDays)

	elapsed := time.Since(start)

	a := analysis.NewAnalyzer(conf.HistoryDays, conf.BrokerFeeBuy, conf.BrokerFeeSell, conf.SalesTax, conf.Filters, logger)
	a.Run(items)

	logger.Info("finished", "elapsed", elapsed.String())
	logger.Info("total ESI requests", "count", esiClient.Calls())
	logger.Info("average ESI request time", "seconds", fmt.Sprintf("%.2f", elapsed.Seconds()/float64(esiClient.Calls())))

	a.Print(analysis.SortFromString(conf.SortBy))
}

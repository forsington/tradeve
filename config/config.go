package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/dustin/go-humanize"
	"github.com/pkg/errors"
)

type Configuration struct {
	LogLevel string `json:"log_level"`

	// CacheEnabled ESI responses
	CacheEnabled bool `json:"cache_enabled"`

	// Download the SDE file even if it exists
	ForceSDEDownload bool `json:"force_sde_download"`

	RegionId int     `json:"region_id"`
	Region   *Region `json:"-"`

	BrokerFeeBuy  float64 `json:"broker_fee_buy"`
	BrokerFeeSell float64 `json:"broker_fee_sell"`
	SalesTax      float64 `json:"sales_tax"`

	// How many days of history to check
	HistoryDays int `json:"history_days"`

	// Valid options "margin", "lpdp", "volume", "flippers", "score"
	SortBy string `json:"sort_by"`

	// Filtering options
	Filters *Filters `json:"filters"`

	// Exclude groups of items
	ExcludeGroups *ExcludeGroups `json:"exclude_groups"`
}

type ExcludeGroups struct {
	Skins      bool `json:"skins"`
	Wearables  bool `json:"wearables"`
	Skillbooks bool `json:"skillbooks"`
	Blueprints bool `json:"blueprints"`
	Skinr      bool `json:"skinr"`
	Crates     bool `json:"crates"`
	Deprecated bool `json:"deprecated"`
}

type Filters struct {
	LowIskBoundary  float64 `json:"low_isk_boundary"`
	HighIskBoundary float64 `json:"high_isk_boundary"`
	MinLPDP         float64 `json:"min_lpdp"` // Volume * Margin / 2
	MinVolume       int     `json:"min_volume"`
	MinProfitMargin float64 `json:"min_profit_margin"` // in percent
	MaxCenterOffset float64 `json:"max_center_offset"` // How close to the centre of the buy/sell spread the item is, 1 = 100% buy, -1 = 100% sell
	MaxFlippers     int     `json:"max_flippers"`      // Orders that have been updated in the last 24h
	MinScore        int     `json:"min_score"`
}

func (c *Configuration) LowIskString() string {
	return fmt.Sprintf("%s ISK", humanize.Comma(int64(c.Filters.LowIskBoundary)))
}

func (c *Configuration) HighIskString() string {
	return fmt.Sprintf("%s ISK", humanize.Comma(int64(c.Filters.HighIskBoundary)))
}

func DefaultConfig() *Configuration {
	config := &Configuration{
		LogLevel:         "info",
		CacheEnabled:     true,
		ForceSDEDownload: false,
		RegionId:         10000002,
		BrokerFeeBuy:     0.005,
		BrokerFeeSell:    0.015,
		SalesTax:         0.0202,
		HistoryDays:      7,
		SortBy:           "lpdp",
		Filters: &Filters{
			LowIskBoundary:  10,
			HighIskBoundary: 100000000000,
			MinLPDP:         100000000,
			MinVolume:       5,
			MinProfitMargin: 0.05,
			MaxCenterOffset: 0.5,
			MaxFlippers:     100,
			MinScore:        0,
		},
		ExcludeGroups: &ExcludeGroups{
			Skins:      true,
			Wearables:  true,
			Skillbooks: true,
			Blueprints: true,
			Skinr:      true,
			Crates:     true,
			Deprecated: true,
		},
	}

	config.Region = regions.GetRegionById(config.RegionId)

	return config
}

// PartialConfiguration is used for unmarshaling and selectively updating
type PartialConfiguration struct {
	LogLevel         *string        `json:"log_level"`
	CacheEnabled     *bool          `json:"cache_enabled"`
	ForceSdeDownload *bool          `json:"force_sde_download"`
	RegionId         *int           `json:"region_id"`
	BrokerFeeBuy     *float64       `json:"broker_fee_buy"`
	BrokerFeeSell    *float64       `json:"broker_fee_sell"`
	SalesTax         *float64       `json:"sales_tax"`
	HistoryDays      *int           `json:"history_days"`
	SortBy           *string        `json:"sort_by"`
	Filters          *Filters       `json:"filters"`
	ExcludeGroups    *ExcludeGroups `json:"exclude_groups"`
}

func ParseConfigFile(configFileName string, config *Configuration) (*Configuration, error) {
	file, _ := os.Open(configFileName)
	decoder := json.NewDecoder(file)
	partial := &PartialConfiguration{}
	err := decoder.Decode(&partial)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing configuration file")
	}

	// Update fields only if they are not nil
	if partial.LogLevel != nil {
		config.LogLevel = *partial.LogLevel
	}
	if partial.CacheEnabled != nil {
		config.CacheEnabled = *partial.CacheEnabled
	}
	if partial.ForceSdeDownload != nil {
		config.ForceSDEDownload = *partial.ForceSdeDownload
	}
	if partial.RegionId != nil {
		config.RegionId = *partial.RegionId
	}
	if partial.BrokerFeeBuy != nil {
		config.BrokerFeeBuy = *partial.BrokerFeeBuy
	}
	if partial.BrokerFeeSell != nil {
		config.BrokerFeeSell = *partial.BrokerFeeSell
	}
	if partial.SalesTax != nil {
		config.SalesTax = *partial.SalesTax
	}
	if partial.HistoryDays != nil {
		config.HistoryDays = *partial.HistoryDays
	}
	if partial.SortBy != nil {
		config.SortBy = *partial.SortBy
	}
	if partial.Filters != nil {
		config.Filters = partial.Filters
	}
	if partial.ExcludeGroups != nil {
		config.ExcludeGroups = partial.ExcludeGroups
	}

	config.Region = regions.GetRegionById(config.RegionId)

	return config, nil
}

func ParseFlags(config *Configuration) *Configuration {
	fDebug := flag.Bool("d", false, "Debug mode")
	fForceSDEDownload := flag.Bool("f", false, "Force download of SDE file")
	fNoCache := flag.Bool("no-cache", false, "Disable cache")

	flag.Parse()

	if *fDebug {
		config.LogLevel = "DEBUG"
	}

	if *fForceSDEDownload {
		config.ForceSDEDownload = true
	}

	if *fNoCache {
		config.CacheEnabled = false
	}

	return config
}

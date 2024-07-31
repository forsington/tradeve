package analysis

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/forsington/tradeve/config"

	"github.com/dustin/go-humanize"
	"github.com/forsington/tradeve/item"
	"github.com/hashicorp/go-hclog"
	"github.com/olekukonko/tablewriter"
)

type Analysis struct {
	Item *item.Item

	Days                int
	AveragePrice        float64
	AverageVolume       int64
	AverageSpread       float64
	AverageHigh         float64
	AverageLow          float64
	AverageProfitMargin float64
	AverageOffset       float64
	ISKVol              float64

	// Hypothetical profit if you were to execute every order of the day
	LPDP float64
	// Orders that have been updated in the last 24h
	Flippers int
	// Aggregation of factors based on weights
	Score int
}

func (a *Analysis) score() int {
	score := 0
	if a.Flippers == 0 {
		score = int(a.LPDP)
	} else {
		score = int(a.LPDP) / a.Flippers
	}

	score = score / 1000000

	return score
}

// How important each factor is in the final score, 0-10
type Weights struct {
	LPDP     float64
	Flippers float64
	Margin   float64
	Offset   float64
}

type Analyses []*Analysis

type Analyzer struct {
	days          int
	buyBrokerFee  float64
	sellBrokerFee float64
	salesTax      float64

	filters  *config.Filters
	analyses Analyses

	logger hclog.Logger
}

func NewAnalyzer(days int, buyBrokerFee, sellBrokerFee, salesTax float64, filters *config.Filters, logger hclog.Logger) *Analyzer {
	return &Analyzer{
		days:          days,
		buyBrokerFee:  buyBrokerFee,
		sellBrokerFee: sellBrokerFee,
		salesTax:      salesTax,
		filters:       filters,
		logger:        logger.Named("analyzer"),
	}
}

func (a *Analyzer) Run(items item.Items) {
	analyses := Analyses{}

	for _, item := range items {
		analysis := a.SummarizeHistory(item)

		analysis.Item = item
		analysis.Days = a.days
		analysis.LPDP = analysis.AverageSpread * float64(analysis.AverageVolume) / 2
		analysis.Flippers = item.Orders.Flippers()

		analysis.Score = analysis.score()

		if a.IsOutlier(analysis) {
			a.logger.Debug("outlier detected, skipping", "item", item.Type.TypeName, "margin", analysis.AverageProfitMargin)
			continue
		}
		if a.PassesFilters(analysis) {
			analyses = append(analyses, analysis)
		}

	}

	a.analyses = analyses
}

func (a *Analyzer) IsOutlier(analysis *Analysis) bool {
	// if any day in the history varies more than 50% from the average price, it's an outlier
	for _, history := range analysis.Item.History {
		if history.Average > analysis.AveragePrice*1.5 || history.Average < analysis.AveragePrice*0.5 {
			return true
		}
	}
	return false
}

func (a *Analyzer) SummarizeHistory(item *item.Item) *Analysis {
	var sumPrice, sumSpread, sumAverageOffset, sumHigh, sumLow float64
	var sumVol int64

	// sum all the days of history
	for _, history := range item.History {
		sumPrice = sumPrice + history.Average
		sumVol = sumVol + history.Volume

		// Broker fee from Buy order, sales tax + broker fee from Sell order
		brokerFeeBuyOrder := history.Low * a.buyBrokerFee
		brokerFeeSellOrder := history.High * a.sellBrokerFee
		salesTax := history.High * a.salesTax
		spread := (history.High - history.Low) - (brokerFeeBuyOrder + brokerFeeSellOrder + salesTax)

		sumSpread = sumSpread + spread
		sumHigh = sumHigh + history.High
		sumLow = sumLow + history.Low

		// 1 = average price is only buy orders filled
		// -1 = average price is only sell orders filled
		offset := ((history.High - history.Average) / ((history.High - history.Low) / 2)) - 1
		// offset is negative if it's closer to the high, but we only care about the absolute value
		if offset < 0 {
			offset = -offset
		}
		sumAverageOffset = sumAverageOffset + offset

	}

	fDays := float64(a.days)
	analysis := &Analysis{
		AveragePrice:        sumPrice / fDays,
		AverageVolume:       sumVol / int64(a.days),
		AverageSpread:       sumSpread / fDays,
		AverageHigh:         sumHigh / fDays,
		AverageLow:          sumLow / fDays,
		AverageProfitMargin: sumSpread / sumLow,
		AverageOffset:       sumAverageOffset / fDays,
		ISKVol:              (sumPrice / fDays) * float64(sumVol/int64(a.days)),
	}

	return analysis
}

func (a *Analyzer) PassesFilters(analysis *Analysis) bool {
	if analysis.AveragePrice > a.filters.HighIskBoundary {
		a.logger.Debug("HighIskBoundary exceeded", "item", analysis.Item.Type.TypeName, "high", analysis.AveragePrice, "boundary", a.filters.HighIskBoundary)
		return false
	}
	if analysis.AveragePrice < a.filters.LowIskBoundary {
		a.logger.Debug("LowIskBoundary exceeded", "item", analysis.Item.Type.TypeName, "low", analysis.AveragePrice, "boundary", a.filters.LowIskBoundary)
		return false
	}
	if analysis.AverageOffset > a.filters.MaxCenterOffset || analysis.AverageOffset < -a.filters.MaxCenterOffset {
		a.logger.Debug("MaxCenterOffset exceeded", "item", analysis.Item.Type.TypeName, "offset", analysis.AverageOffset, "max", a.filters.MaxCenterOffset)
		return false
	}
	if analysis.LPDP < a.filters.MinLPDP {
		a.logger.Debug("MinLPDP exceeded", "item", analysis.Item.Type.TypeName, "lpdp", analysis.LPDP, "min", a.filters.MinLPDP)
		return false
	}
	if analysis.AverageProfitMargin < a.filters.MinProfitMargin {
		a.logger.Debug("MinMargin exceeded", "item", analysis.Item.Type.TypeName, "margin", analysis.AverageProfitMargin, "min", a.filters.MinProfitMargin)
		return false
	}
	if analysis.AverageVolume < int64(a.filters.MinVolume) {
		a.logger.Debug("MinVolume exceeded", "item", analysis.Item.Type.TypeName, "volume", analysis.AverageVolume, "min", a.filters.MinVolume)
		return false
	}

	if analysis.Flippers > a.filters.MaxFlippers {
		a.logger.Info("MaxTraders exceeded", "item", analysis.Item.Type.TypeName, "traders", analysis.Flippers, "max", a.filters.MaxFlippers)
		return false
	}

	if analysis.Score < a.filters.MinScore {
		a.logger.Debug("MinScore exceeded", "item", analysis.Item.Type.TypeName, "score", analysis.Score, "min", a.filters.MinScore)
		return false
	}

	return true
}

type Sort int

var (
	ByMargin   Sort = 0
	ByLPDP     Sort = 1
	ByVolume   Sort = 2
	ByFlippers Sort = 3
	ByScore    Sort = 4
)

func SortFromString(sort string) Sort {
	switch sort {
	case "margin":
		return ByMargin
	case "lpdp":
		return ByLPDP
	case "volume":
		return ByVolume
	case "flippers":
		return ByFlippers
	case "score":
		return ByScore
	default:
		return ByLPDP
	}
}

func (a *Analyzer) Print(sort Sort) {
	analyses := a.analyses.Sort(sort)

	fmt.Printf("\nAnalysis complete. %d items match criteria over the last %d days (daily average):\n\n", len(a.analyses), a.days)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetColWidth(50)
	table.SetHeader([]string{"#", "Score", "LPDP", "Margin", "Take Home", "Item", "Price", "Volume", "Flippers", "ISK Vol"})
	for i, a := range analyses {
		table.Append([]string{
			strconv.Itoa(i + 1),
			fmt.Sprintf("%d", a.Score),
			iskFmt(a.LPDP),
			fmt.Sprintf("%.1f%%", a.AverageProfitMargin*100),
			iskFmt(a.AverageSpread),
			a.Item.Type.TypeName,
			iskFmt(a.AveragePrice),
			fmt.Sprintf("%d", a.AverageVolume),
			fmt.Sprintf("%d", a.Flippers),
			iskFmt(a.ISKVol),
		})
	}
	table.Render()
	fmt.Println()
}

func iskFmt(amount float64) string {
	return fmt.Sprintf("%s ISK", humanize.Comma(int64(amount)))
}
func (analyses Analyses) Sort(sortBy Sort) Analyses {
	if len(analyses) == 0 {
		return analyses
	}

	switch sortBy {
	case ByMargin:
		sort.Slice(analyses, func(i, j int) bool {
			return analyses[i].AverageProfitMargin > analyses[j].AverageProfitMargin
		})
	case ByLPDP:
		sort.Slice(analyses, func(i, j int) bool {
			return analyses[i].LPDP > analyses[j].LPDP
		})
	case ByVolume:
		sort.Slice(analyses, func(i, j int) bool {
			return analyses[i].AverageVolume > analyses[j].AverageVolume
		})
	case ByFlippers:
		sort.Slice(analyses, func(i, j int) bool {
			return analyses[i].Flippers > analyses[j].Flippers
		})
	case ByScore:
		sort.Slice(analyses, func(i, j int) bool {
			return analyses[i].Score > analyses[j].Score
		})
	default:
		// No sorting
	}
	return analyses
}

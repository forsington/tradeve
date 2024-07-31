package analysis

import (
	"testing"

	"github.com/forsington/tradeve/config"
	"github.com/forsington/tradeve/history"
	"github.com/forsington/tradeve/item"
	"github.com/hashicorp/go-hclog"
)

// SummarizeHistory
func TestSummarizeHistory(t *testing.T) {
	item := &item.Item{
		Type: &item.Type{TypeId: 34, TypeName: "Tritanium"},
	}

	cases := []struct {
		name          string
		history       []*history.MarketHistory
		expected      *Analysis
		brokerFeeBuy  float64
		brokerFeeSell float64
		salesTax      float64
	}{
		{
			name: "no taxes / fees",
			history: []*history.MarketHistory{
				{Average: 3, Volume: 2, High: 4, Low: 2},
				{Average: 3, Volume: 2, High: 4, Low: 2},
			},
			expected: &Analysis{
				Item:                item,
				AveragePrice:        3,
				AverageVolume:       2,
				AverageSpread:       2,
				AverageHigh:         4,
				AverageLow:          2,
				AverageProfitMargin: 1, // 100%
				AverageOffset:       0,
			},
			brokerFeeBuy:  0,
			brokerFeeSell: 0,
			salesTax:      0,
		},
		{
			name: "no taxes / fees, different values",
			history: []*history.MarketHistory{
				{Average: 300, Volume: 200, High: 600, Low: 200},
				{Average: 300, Volume: 300, High: 500, Low: 300},
			},
			expected: &Analysis{
				Item:                item,
				AveragePrice:        300,
				AverageVolume:       250,
				AverageSpread:       300,
				AverageHigh:         550,
				AverageLow:          250,
				AverageProfitMargin: 1.2, // 120%
				AverageOffset:       0.75,
			},
			brokerFeeBuy:  0,
			brokerFeeSell: 0,
			salesTax:      0,
		},
		{
			name: "taxes / fees applied",
			history: []*history.MarketHistory{
				{Average: 300, Volume: 200, High: 600, Low: 200},
				{Average: 300, Volume: 300, High: 500, Low: 300},
			},
			expected: &Analysis{
				Item:                item,
				AveragePrice:        300,
				AverageVolume:       250,
				AverageSpread:       270,
				AverageHigh:         550,
				AverageLow:          250,
				AverageProfitMargin: 1.08, // 108%
				AverageOffset:       0.75,
			},
			brokerFeeBuy:  0.01,
			brokerFeeSell: 0.02,
			salesTax:      0.03,
		},
	}

	for _, c := range cases {
		a := NewAnalyzer(len(c.history), c.brokerFeeBuy, c.brokerFeeSell, c.salesTax, &config.Filters{}, hclog.NewNullLogger())

		item.History = c.history
		analysis := a.SummarizeHistory(item)

		// check fields in analysis
		if analysis.AveragePrice != c.expected.AveragePrice {
			t.Errorf("case %s , expected %f average price, got %f", c.name, c.expected.AveragePrice, analysis.AveragePrice)
		}
		if analysis.AverageVolume != c.expected.AverageVolume {
			t.Errorf("case %s, expected %d average volume, got %d", c.name, c.expected.AverageVolume, analysis.AverageVolume)
		}
		if analysis.AverageSpread != c.expected.AverageSpread {
			t.Errorf("case %s, expected %f average spread, got %f", c.name, c.expected.AverageSpread, analysis.AverageSpread)
		}
		if analysis.AverageHigh != c.expected.AverageHigh {
			t.Errorf("case %s, expected %f average high, got %f", c.name, c.expected.AverageHigh, analysis.AverageHigh)
		}
		if analysis.AverageLow != c.expected.AverageLow {
			t.Errorf("case %s, expected %f average low, got %f", c.name, c.expected.AverageLow, analysis.AverageLow)
		}
		if analysis.AverageProfitMargin != c.expected.AverageProfitMargin {
			t.Errorf("case %s, expected %f average margin, got %f", c.name, c.expected.AverageProfitMargin, analysis.AverageProfitMargin)
		}
		if analysis.AverageOffset != c.expected.AverageOffset {
			t.Errorf("case %s, expected %f, got %f", c.name, c.expected.AverageOffset, analysis.AverageOffset)
		}
	}
}

func TestPassesFilters(t *testing.T) {
	item := &item.Item{
		Type: &item.Type{
			TypeId:   34,
			TypeName: "Tritanium",
		},
	}

	filter := &config.Filters{
		LowIskBoundary:  10,
		HighIskBoundary: 10000,
		MinLPDP:         100000,
		MinVolume:       10,
		MinProfitMargin: 0.1,
		MaxCenterOffset: 0.1,
		MaxFlippers:     5,
	}

	okAnalysis := &Analysis{
		Item:                item,
		AveragePrice:        100,
		LPDP:                100000,
		AverageVolume:       10,
		AverageProfitMargin: 0.1,
		AverageOffset:       0.1,
		Flippers:            5,
	}

	nokCases := []struct {
		name     string
		analysis *Analysis
		expected bool
	}{
		{
			name:     "ok",
			analysis: okAnalysis,
			expected: true,
		},
		{
			name: "AveragePrice < LowIskBoundary",
			analysis: &Analysis{
				Item:                item,
				AveragePrice:        5,
				LPDP:                okAnalysis.LPDP,
				AverageVolume:       okAnalysis.AverageVolume,
				AverageProfitMargin: okAnalysis.AverageProfitMargin,
				AverageOffset:       okAnalysis.AverageOffset,
				Flippers:            okAnalysis.Flippers,
			},
			expected: false,
		},
		{
			name: "AveragePrice > HighIskBoundary",
			analysis: &Analysis{
				Item:                item,
				AveragePrice:        10001,
				LPDP:                okAnalysis.LPDP,
				AverageVolume:       okAnalysis.AverageVolume,
				AverageProfitMargin: okAnalysis.AverageProfitMargin,
				AverageOffset:       okAnalysis.AverageOffset,
				Flippers:            okAnalysis.Flippers,
			},
			expected: false,
		},
		{
			name: "LPDP < MinLPDP",
			analysis: &Analysis{
				Item:                item,
				AveragePrice:        okAnalysis.AveragePrice,
				LPDP:                99999,
				AverageVolume:       okAnalysis.AverageVolume,
				AverageProfitMargin: okAnalysis.AverageProfitMargin,
				AverageOffset:       okAnalysis.AverageOffset,
				Flippers:            okAnalysis.Flippers,
			},
			expected: false,
		},
		{
			name: "AverageVolume < MinVolume",
			analysis: &Analysis{
				Item:                item,
				AveragePrice:        okAnalysis.AveragePrice,
				LPDP:                okAnalysis.LPDP,
				AverageVolume:       9,
				AverageProfitMargin: okAnalysis.AverageProfitMargin,
				AverageOffset:       okAnalysis.AverageOffset,
				Flippers:            okAnalysis.Flippers,
			},
			expected: false,
		},
		{
			name: "AverageProfitMargin < MinProfitMargin",
			analysis: &Analysis{
				Item:                item,
				AveragePrice:        okAnalysis.AveragePrice,
				LPDP:                okAnalysis.LPDP,
				AverageVolume:       okAnalysis.AverageVolume,
				AverageProfitMargin: 0.09,
				AverageOffset:       okAnalysis.AverageOffset,
				Flippers:            okAnalysis.Flippers,
			},
			expected: false,
		},
		{
			name: "AverageOffset > MaxCenterOffset",
			analysis: &Analysis{
				Item:                item,
				AveragePrice:        okAnalysis.AveragePrice,
				LPDP:                okAnalysis.LPDP,
				AverageVolume:       okAnalysis.AverageVolume,
				AverageProfitMargin: okAnalysis.AverageProfitMargin,
				AverageOffset:       0.11,
				Flippers:            okAnalysis.Flippers,
			},
			expected: false,
		},
		{
			name: "AverageOffset < -MaxCenterOffset",
			analysis: &Analysis{
				Item:                item,
				AveragePrice:        okAnalysis.AveragePrice,
				LPDP:                okAnalysis.LPDP,
				AverageVolume:       okAnalysis.AverageVolume,
				AverageProfitMargin: okAnalysis.AverageProfitMargin,
				AverageOffset:       -0.11,
				Flippers:            okAnalysis.Flippers,
			},
			expected: false,
		},
		{
			name: "AverageActiveTraders > MaxTraders",
			analysis: &Analysis{
				Item:                item,
				AveragePrice:        okAnalysis.AveragePrice,
				LPDP:                okAnalysis.LPDP,
				AverageVolume:       okAnalysis.AverageVolume,
				AverageProfitMargin: okAnalysis.AverageProfitMargin,
				AverageOffset:       okAnalysis.AverageOffset,
				Flippers:            6,
			},
			expected: false,
		},
	}

	for _, c := range nokCases {
		a := NewAnalyzer(5, 0.1, 0.1, 0.1, filter, hclog.NewNullLogger())

		if a.PassesFilters(c.analysis) != c.expected {
			t.Errorf("case %s, expected %v, got %v", c.name, c.expected, !c.expected)
		}
	}
}

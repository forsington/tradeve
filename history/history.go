package history

import (
	"time"
)

// MarketHistory for an item
type MarketHistory struct {
	Average    float64
	Date       string
	High       float64
	Low        float64
	OrderCount int64
	Volume     int64
}

type MarketHistories []*MarketHistory

func (m MarketHistories) TrimDays(days int) MarketHistories {
	// ESI doesn't include today's or yesterday's data, so we need to go back 2 days
	twoDaysAgo := time.Now().AddDate(0, 0, -2)

	var newHistories MarketHistories

	for _, h := range m {
		date, _ := time.Parse("2006-01-02", h.Date)

		if date.After(twoDaysAgo.AddDate(0, 0, -int(days))) && date.Before(twoDaysAgo) {
			newHistories = append(newHistories, h)
		}
	}

	return newHistories
}

func (m MarketHistories) ExpectedDates(days int) []string {
	// ESI doesn't include today's or yesterday's data, so we need to go back 2 days
	twoDaysAgo := time.Now().AddDate(0, 0, -2)

	var dates []string

	for i := 0; i < days; i++ {
		dates = append(dates, twoDaysAgo.AddDate(0, 0, -i).Format("2006-01-02"))
	}

	return dates
}

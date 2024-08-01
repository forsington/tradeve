package history

import (
	"fmt"

	"github.com/forsington/tradeve/esi"
)

type ESIRepository struct {
	ESI *esi.Client
}

func NewESIRepository(esi *esi.Client) Repository {
	return &ESIRepository{
		ESI: esi,
	}
}

func (e *ESIRepository) Get(typeId, days, region int) (MarketHistories, error) {
	histories := MarketHistories{}
	esiHistories, err := e.ESI.GetMarketHistory(typeId, region)
	if err != nil {
		return nil, err
	}

	for _, history := range esiHistories {
		histories = append(histories, &MarketHistory{
			Average:    history.Average,
			Date:       history.Date,
			High:       history.Highest,
			Low:        history.Lowest,
			OrderCount: history.OrderCount,
			Volume:     history.Volume,
		})
	}

	trimmed := histories.TrimDays(days)

	// ESI doesn't return an entry for days where no items were traded, so we need to fill in the gaps
	if len(trimmed) != days {
		for _, date := range trimmed.ExpectedDates(days) {
			dayExists := false
			for _, history := range trimmed {
				if history.Date == date {
					dayExists = true
					break
				}
			}
			if !dayExists {
				trimmed = append(trimmed, &MarketHistory{
					Date: date,
				})
			}
		}
	}

	return trimmed, nil
}

func (e *ESIRepository) Upsert(typeId int, marketHistory MarketHistories) error {
	return fmt.Errorf("not implemented")
}

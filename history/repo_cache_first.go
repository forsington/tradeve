package history

import "fmt"

type CacheFirstRepository struct {
	cacheRepository Repository
	esiRepository   Repository
}

func NewCacheFirstRepository(cacheRepository, esiRepository Repository) *CacheFirstRepository {
	return &CacheFirstRepository{
		cacheRepository: cacheRepository,
		esiRepository:   esiRepository,
	}
}

// Check the cache first
// If the cache doesn't have the data, check ESI
// If ESI has the data, update the cache
func (c *CacheFirstRepository) Get(typeId, days, region int) (MarketHistories, error) {
	marketHistories, err := c.cacheRepository.Get(typeId, days, region)
	if err != nil {
		marketHistories, err = c.fetchAndUpdateCache(typeId, region)
		if err != nil {
			return nil, err
		}

		marketHistories = marketHistories.TrimDays(days)
	}

	if len(marketHistories) != days {
		marketHistories, err = c.fetchAndUpdateCache(typeId, region)
		if err != nil {
			return nil, err
		}
	}

	trimmed := marketHistories.TrimDays(days)

	return trimmed, nil
}

func (c *CacheFirstRepository) Upsert(typeId int, marketHistories MarketHistories) error {
	return fmt.Errorf("not implemented")
}

func (c *CacheFirstRepository) fetchAndUpdateCache(typeId, region int) (MarketHistories, error) {
	// Updating the cache, do all days
	days := 365
	marketHistories, err := c.esiRepository.Get(typeId, days, region)
	if err != nil {
		return nil, err
	}
	err = c.cacheRepository.Upsert(typeId, marketHistories)
	if err != nil {
		return nil, err
	}

	return marketHistories, nil
}

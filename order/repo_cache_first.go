package order

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
func (c *CacheFirstRepository) Get(typeId, region int) (Orders, error) {
	staleCache := false
	orders, err := c.cacheRepository.Get(typeId, region)

	if err != nil {
		staleCache = true
	}

	today := orders.Today()
	if len(today) == 0 {
		staleCache = true
	}

	if staleCache {
		orders, err = c.fetchAndUpdateCache(typeId, region)
		if err != nil {
			return nil, err
		}
	}

	return orders, nil
}

func (c *CacheFirstRepository) fetchAndUpdateCache(typeId, region int) (Orders, error) {
	// Updating the cache, do all days
	orders, err := c.esiRepository.Get(typeId, region)
	if err != nil {
		return nil, err
	}
	err = c.cacheRepository.Upsert(typeId, orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (c *CacheFirstRepository) Upsert(typeId int, orders Orders) error {
	return fmt.Errorf("not implemented")
}

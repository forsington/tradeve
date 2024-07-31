package history

type Repository interface {
	Get(typeId, days, region int) (MarketHistories, error)
	Upsert(typeId int, marketHistory MarketHistories) error
}

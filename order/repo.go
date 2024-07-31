package order

type Repository interface {
	Get(typeId, regionId int) (Orders, error)
	Upsert(typeId int, orders Orders) error
}

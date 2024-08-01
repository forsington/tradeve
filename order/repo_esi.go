package order

import (
	"fmt"
	"time"

	"github.com/forsington/tradeve/esi"
)

type ESIRepository struct {
	client *esi.Client
}

func NewESIRepository(client *esi.Client) Repository {
	return &ESIRepository{
		client: client,
	}
}

func (r *ESIRepository) Get(typeId, regionId int) (Orders, error) {
	esiOrders, err := r.client.GetMarketOrders(regionId, typeId)
	if err != nil {
		return nil, err
	}

	var orders Orders
	for _, order := range esiOrders {
		orders = append(orders, &Order{
			FetchDate:    time.Now(),
			Duration:     order.Duration,
			IsBuyOrder:   order.IsBuyOrder,
			Issued:       order.Issued,
			LocationId:   order.LocationId,
			MinVolume:    order.MinVolume,
			OrderId:      order.OrderId,
			Price:        order.Price,
			Range:        order.Range_,
			SystemId:     order.SystemId,
			TypeId:       order.TypeId,
			VolumeRemain: order.VolumeRemain,
			VolumeTotal:  order.VolumeTotal,
		})
	}

	return orders, nil
}

func (r *ESIRepository) Upsert(typeId int, orders Orders) error {
	return fmt.Errorf("not implemented")
}

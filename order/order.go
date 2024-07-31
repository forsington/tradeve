package order

import (
	"time"
)

type Order struct {
	FetchDate time.Time

	Duration     int32     `json:"duration"`
	IsBuyOrder   bool      `json:"is_buy_order"`
	Issued       time.Time `json:"issued"`
	LocationId   int64     `json:"location_id"`
	MinVolume    int32     `json:"min_volume"`
	OrderId      int64     `json:"order_id"`
	Price        float64   `json:"price"`
	Range        string    `json:"range"`
	SystemId     int32     `json:"system_id"`
	TypeId       int32     `json:"type_id"`
	VolumeRemain int32     `json:"volume_remain"`
	VolumeTotal  int32     `json:"volume_total"`
}

type Orders []*Order

// Flippers returns the number of orders that have been updated in the last 24 hours
func (o Orders) Flippers() int {
	flippers := 0
	for _, order := range o {
		if time.Since(order.Issued) < 24*time.Hour {
			flippers++
		}
	}
	return flippers

}

// Orders from today
func (o Orders) Today() Orders {
	today := Orders{}
	for _, order := range o {
		if order.FetchDate.After(time.Now().UTC().Truncate(24 * time.Hour)) {
			today = append(today, order)
		}
	}
	return today
}

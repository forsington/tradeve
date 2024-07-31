package item

import (
	"fmt"

	"github.com/forsington/tradeve/history"
	"github.com/forsington/tradeve/order"
)

// Item struct
type Item struct {
	Type    *Type
	History []*history.MarketHistory
	Orders  order.Orders

	// AveragePrice float64
}

type Items []*Item

func (items Items) Len() int {
	return len(items)
}

func (items Items) Append(item *Item) (Items, error) {
	// don't add duplicates
	for _, existing := range items {
		if existing.Type == nil {
			return nil, fmt.Errorf("existing item has no type")
		}
		if item.Type == nil {
			return nil, fmt.Errorf("new item has no type")
		}

		if existing.Type.TypeId == item.Type.TypeId {
			return items, nil
		}
	}

	return append(items, item), nil
}

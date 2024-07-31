package item

import (
	"github.com/forsington/tradeve/history"
	"github.com/forsington/tradeve/order"
	"github.com/hashicorp/go-hclog"
)

type Service struct {
	logger         hclog.Logger
	orderService   *order.Service
	historyService *history.Service
}

func NewService(logger hclog.Logger, orderService *order.Service, historyService *history.Service) *Service {
	return &Service{
		logger:         logger,
		orderService:   orderService,
		historyService: historyService,
	}
}

func (s *Service) GetItems(types []*Type, regionId, days int) Items {
	depth := len(types)
	items := make(Items, 0)

	for i, t := range types {
		history, err := s.historyService.Get(t.TypeId, regionId, days)
		if err != nil {
			s.logger.Error("error", "item", t.TypeName, "error", err)
			continue
		}

		orders, err := s.orderService.Get(t.TypeId, regionId)
		if err != nil {
			s.logger.Error("error fetching orders", "item", t.TypeName, "error", err)
			continue
		}

		if i%100 == 0 {
			s.logger.Info("fetched item", "current", i, "total", depth)
		}
		s.logger.Debug("fetched item", "index", i+1, "total", depth, "item", t.TypeName)

		s := &Item{
			Type:    t,
			History: history,
			Orders:  orders,
		}
		items = append(items, s)
	}
	s.logger.Info("fetched history", "current", depth, "total", depth)

	return items
}

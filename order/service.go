package order

import "fmt"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Get(typeId, regionId int) (Orders, error) {
	orders, err := s.repo.Get(typeId, regionId)
	if err != nil {
		return nil, err
	}

	orders = orders.Today()
	if len(orders) == 0 {
		return nil, fmt.Errorf("no orders found for %d today", typeId)
	}

	return orders, nil
}

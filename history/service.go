package history

import "github.com/hashicorp/go-hclog"

type Service struct {
	logger hclog.Logger
	repo   Repository
}

func NewService(logger hclog.Logger, repo Repository) *Service {
	return &Service{
		logger: logger,
		repo:   repo,
	}
}

func (s *Service) Get(typeId, regionId, days int) (MarketHistories, error) {
	history, err := s.repo.Get(typeId, days, regionId)
	if err != nil {
		return nil, err
	}

	return history, nil
}

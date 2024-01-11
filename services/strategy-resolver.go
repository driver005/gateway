package services

import (
	"context"
)

type StrategyResolverService struct {
	ctx context.Context
}

func NewStrategyResolverService() *StrategyResolverService {
	return &StrategyResolverService{
		context.Background(),
	}
}

func (s *StrategyResolverService) SetContext(context context.Context) *StrategyResolverService {
	s.ctx = context
	return s
}

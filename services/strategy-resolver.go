package services

import (
	"context"
)

type StrategyResolverService struct {
	ctx context.Context
}

func NewStrategyResolverService(ctx context.Context) *StrategyResolverService {
	return &StrategyResolverService{
		ctx,
	}
}

func (s StrategyResolverService) ResolveBatchJobByType(batchtype string)

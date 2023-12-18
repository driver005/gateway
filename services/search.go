package services

import (
	"context"
)

type DefaultSearchService struct {
	ctx context.Context
}

func NewDefaultSearchService(
	ctx context.Context,
) *DefaultSearchService {
	return &DefaultSearchService{
		ctx,
	}
}

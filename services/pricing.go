package services

import (
	"context"
)

type PricingService struct {
	ctx context.Context
}

func NewPricingService(
	ctx context.Context,
) *PricingService {
	return &PricingService{
		ctx,
	}
}

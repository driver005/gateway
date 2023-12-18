package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type ShippingTaxRateService struct {
	ctx  context.Context
	repo *repository.ShippingTaxRateRepo
}

func NewShippingTaxRateService(
	ctx context.Context,
	repo *repository.ShippingTaxRateRepo,
) *ShippingTaxRateService {
	return &ShippingTaxRateService{
		ctx,
		repo,
	}
}

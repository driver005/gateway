package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type ProductTaxRateService struct {
	ctx  context.Context
	repo *repository.ProductTaxRateRepo
}

func NewProductTaxRateService(
	ctx context.Context,
	repo *repository.ProductTaxRateRepo,
) *ProductTaxRateService {
	return &ProductTaxRateService{
		ctx,
		repo,
	}
}

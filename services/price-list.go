package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type PriceListService struct {
	ctx                      context.Context
	repo                     *repository.PriceListRepo
	moneyAmountRepository    *repository.MoneyAmountRepo
	productVariantRepository *repository.ProductVariantRepo
}

func NewPriceListService(
	ctx context.Context,
	repo *repository.PriceListRepo,
	moneyAmountRepository *repository.MoneyAmountRepo,
	productVariantRepository *repository.ProductVariantRepo,
) *PriceListService {
	return &PriceListService{
		ctx,
		repo,
		moneyAmountRepository,
		productVariantRepository,
	}
}

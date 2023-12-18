package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type ProductVariantService struct {
	ctx                          context.Context
	repo                         *repository.ProductVariantRepo
	productRepository            *repository.ProductRepo
	moneyAmountRepository        *repository.MoneyAmountRepo
	productOptionValueRepository *repository.ProductOptionValueRepo
	cartRepository               *repository.CartRepo
}

func NewProductVariantService(
	ctx context.Context,
	repo *repository.ProductVariantRepo,
	productRepository *repository.ProductRepo,
	moneyAmountRepository *repository.MoneyAmountRepo,
	productOptionValueRepository *repository.ProductOptionValueRepo,
	cartRepository *repository.CartRepo,
) *ProductVariantService {
	return &ProductVariantService{
		ctx,
		repo,
		productRepository,
		moneyAmountRepository,
		productOptionValueRepository,
		cartRepository,
	}
}

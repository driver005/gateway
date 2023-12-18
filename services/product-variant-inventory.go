package services

import (
	"context"
)

type ProductVariantInventoryService struct {
	ctx                   context.Context
	productVariantService *ProductVariantService
}

func NewProductVariantInventoryService(
	ctx context.Context,
	productVariantService *ProductVariantService,
) *ProductVariantInventoryService {
	return &ProductVariantInventoryService{
		ctx,
		productVariantService,
	}
}

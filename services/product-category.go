package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type ProductCategoryService struct {
	ctx  context.Context
	repo *repository.ProductCategoryRepo
}

func NewProductCategoryService(
	ctx context.Context,
	repo *repository.ProductCategoryRepo,
) *ProductCategoryService {
	return &ProductCategoryService{
		ctx,
		repo,
	}
}

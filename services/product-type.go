package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type ProductTypeService struct {
	ctx  context.Context
	repo *repository.ProductTypeRepo
}

func NewProductTypeService(
	ctx context.Context,
	repo *repository.ProductTypeRepo,
) *ProductTypeService {
	return &ProductTypeService{
		ctx,
		repo,
	}
}

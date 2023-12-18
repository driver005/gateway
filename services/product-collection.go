package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type ProductCollectionService struct {
	ctx  context.Context
	repo *repository.ProductCollectionRepo
}

func NewProductCollectionService(
	ctx context.Context,
	repo *repository.ProductCollectionRepo,
) *ProductCollectionService {
	return &ProductCollectionService{
		ctx,
		repo,
	}
}

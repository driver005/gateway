package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type ProductTagService struct {
	ctx  context.Context
	repo *repository.ProductTagRepo
}

func NewProductTagService(
	ctx context.Context,
	repo *repository.ProductTagRepo,
) *ProductTagService {
	return &ProductTagService{
		ctx,
		repo,
	}
}

package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type CartService struct {
	ctx  context.Context
	repo *repository.CartRepo
}

func NewCartService(ctx context.Context, repo *repository.CartRepo) *CartService {
	return &CartService{
		ctx,
		repo,
	}
}

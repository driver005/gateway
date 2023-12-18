package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type SwapService struct {
	ctx  context.Context
	repo *repository.SwapRepo
}

func NewSwapService(ctx context.Context, repo *repository.SwapRepo) *SwapService {
	return &SwapService{
		ctx,
		repo,
	}
}

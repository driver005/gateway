package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type IdempotencyKeyService struct {
	ctx  context.Context
	repo *repository.IdempotencyKeyRepo
}

func NewIdempotencyKeyService(
	ctx context.Context,
	repo *repository.IdempotencyKeyRepo,
) *IdempotencyKeyService {
	return &IdempotencyKeyService{
		ctx,
		repo,
	}
}

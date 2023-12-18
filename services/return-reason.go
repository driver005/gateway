package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type ReturnReasonService struct {
	ctx  context.Context
	repo *repository.ReturnReasonRepo
}

func NewReturnReasonService(
	ctx context.Context,
	repo *repository.ReturnReasonRepo,
) *ReturnReasonService {
	return &ReturnReasonService{
		ctx,
		repo,
	}
}

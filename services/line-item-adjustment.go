package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type LineItemAdjustmentService struct {
	ctx  context.Context
	repo *repository.LineItemAdjustmentRepo
}

func NewLineItemAdjustmentService(
	ctx context.Context,
	repo *repository.LineItemAdjustmentRepo,
) *LineItemAdjustmentService {
	return &LineItemAdjustmentService{
		ctx,
		repo,
	}
}

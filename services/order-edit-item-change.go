package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type OrderItemChangeService struct {
	ctx  context.Context
	repo *repository.OrderItemChangeRepo
}

func NewOrderItemChangeService(
	ctx context.Context,
	repo *repository.OrderItemChangeRepo,
) *OrderItemChangeService {
	return &OrderItemChangeService{
		ctx,
		repo,
	}
}

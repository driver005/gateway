package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type OrderService struct {
	ctx               context.Context
	repo              *repository.OrderRepo
	addressRepository *repository.AddressRepo
}

func NewOrderService(
	ctx context.Context,
	repo *repository.OrderRepo,
	addressRepository *repository.AddressRepo,
) *OrderService {
	return &OrderService{
		ctx,
		repo,
		addressRepository,
	}
}

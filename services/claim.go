package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type ClaimService struct {
	ctx                      context.Context
	repo                     *repository.ClaimRepo
	addressRepository        *repository.AddressRepo
	shippingMethodRepository *repository.ShippingMethodRepo
	lineItemRepository       *repository.LineItemRepo
}

func NewClaimService(
	ctx context.Context,
	repo *repository.ClaimRepo,
	addressRepository *repository.AddressRepo,
	shippingMethodRepository *repository.ShippingMethodRepo,
	lineItemRepository *repository.LineItemRepo,
) *ClaimService {
	return &ClaimService{
		ctx,
		repo,
		addressRepository,
		shippingMethodRepository,
		lineItemRepository,
	}
}

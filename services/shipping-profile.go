package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type ShippingProfileService struct {
	ctx               context.Context
	repo              *repository.ShippingProfileRepo
	productRepository *repository.ProductRepo
}

func NewShippingProfileService(
	ctx context.Context,
	repo *repository.ShippingProfileRepo,
	productRepository *repository.ProductRepo,
) *ShippingProfileService {
	return &ShippingProfileService{
		ctx,
		repo,
		productRepository,
	}
}

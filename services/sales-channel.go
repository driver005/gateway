package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type SalesChannelService struct {
	ctx  context.Context
	repo *repository.SalesChannelRepo
}

func NewSalesChannelService(
	ctx context.Context,
	repo *repository.SalesChannelRepo,
) *SalesChannelService {
	return &SalesChannelService{
		ctx,
		repo,
	}
}

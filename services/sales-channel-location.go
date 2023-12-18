package services

import (
	"context"
)

type SalesChannelLocationService struct {
	ctx                 context.Context
	salesChannelService *SalesChannelService
}

func NewSalesChannelLocationService(
	ctx context.Context,
	salesChannelService *SalesChannelService,
) *SalesChannelLocationService {
	return &SalesChannelLocationService{
		ctx,
		salesChannelService,
	}
}

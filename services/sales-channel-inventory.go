package services

import (
	"context"
)

type SalesChannelInventoryService struct {
	ctx context.Context
}

func NewSalesChannelInventoryService(
	ctx context.Context,
) *SalesChannelInventoryService {
	return &SalesChannelInventoryService{
		ctx,
	}
}

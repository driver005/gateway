package services

import (
	"context"

	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type SalesChannelInventoryService struct {
	ctx context.Context
	r   Registry
}

func NewSalesChannelInventoryService(
	r Registry,
) *SalesChannelInventoryService {
	return &SalesChannelInventoryService{
		context.Background(),
		r,
	}
}

func (s *SalesChannelInventoryService) SetContext(context context.Context) *SalesChannelInventoryService {
	s.ctx = context
	return s
}

func (s *SalesChannelInventoryService) RetrieveAvailableItemQuantity(salesChannelId uuid.UUID, inventoryItemId uuid.UUID) (int, *utils.ApplictaionError) {
	locationIds, err := s.r.SalesChannelLocationService().SetContext(s.ctx).ListLocationIds(salesChannelId)
	if err != nil {
		return 0, err
	}
	availableQuantity, err := s.r.InventoryService().RetrieveAvailableQuantity(s.ctx, inventoryItemId, locationIds)
	if err != nil {
		return 0, err
	}
	return availableQuantity, nil
}

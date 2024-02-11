package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FilterableSalesChannel struct {
	core.FilterModel
	Name        string `json:"name,omitempty" validate:"omitempty"`
	Description string `json:"description,omitempty" validate:"omitempty"`
}

type CreateSalesChannelInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty" validate:"omitempty"`
	IsDisabled  bool   `json:"is_disabled,omitempty" validate:"omitempty"`
}

type UpdateSalesChannelInput CreateSalesChannelInput

type ProductBatchSalesChannel struct {
	Id uuid.UUID `json:"id"`
}

type SalesChannelStockLocations struct {
	LocationId uuid.UUID `json:"location_id"`
}

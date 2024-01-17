package types

import "github.com/google/uuid"

type CreateSalesChannelInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty" validate:"omitempty"`
	IsDisabled  bool   `json:"is_disabled,omitempty" validate:"omitempty"`
}

type UpdateSalesChannelInput CreateSalesChannelInput

type ProductBatchSalesChannel struct {
	Id uuid.UUID `json:"id"`
}

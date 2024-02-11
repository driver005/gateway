package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FilterableReturn struct {
	core.FilterModel

	IdempotencyKey string `json:"idempotency_key,omitempty" validate:"omitempty"`
}

type OrderReturnItem struct {
	ItemId   uuid.UUID `json:"item_id"`
	Quantity int       `json:"quantity"`
	ReasonId uuid.UUID `json:"reason_id,omitempty" validate:"omitempty"`
	Note     string    `json:"note,omitempty" validate:"omitempty"`
}

type CreateReturnInput struct {
	OrderId        uuid.UUID                       `json:"order_id"`
	SwapId         uuid.UUID                       `json:"swap_id,omitempty" validate:"omitempty"`
	ClaimOrderId   uuid.UUID                       `json:"claim_order_id,omitempty" validate:"omitempty"`
	Items          []OrderReturnItem               `json:"items,omitempty" validate:"omitempty"`
	ShippingMethod *CreateClaimReturnShippingInput `json:"shipping_method,omitempty" validate:"omitempty"`
	NoNotification bool                            `json:"no_notification,omitempty" validate:"omitempty"`
	Metadata       core.JSONB                      `json:"metadata,omitempty" validate:"omitempty"`
	RefundAmount   float64                         `json:"refund_amount,omitempty" validate:"omitempty"`
	LocationId     uuid.UUID                       `json:"location_id,omitempty" validate:"omitempty"`
}

type UpdateReturnInput struct {
	Items          []OrderReturnItem               `json:"items,omitempty" validate:"omitempty"`
	ShippingMethod *CreateClaimReturnShippingInput `json:"shipping_method,omitempty" validate:"omitempty"`
	NoNotification bool                            `json:"no_notification,omitempty" validate:"omitempty"`
	Metadata       core.JSONB                      `json:"metadata,omitempty" validate:"omitempty"`
}

type ReturnReceive struct {
	Items      []OrderReturnItem `json:"items"`
	Refund     float64           `json:"refund,omitempty" validate:"omitempty"`
	LocationId uuid.UUID         `json:"location_id,omitempty" validate:"omitempty"`
}

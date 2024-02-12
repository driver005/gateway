package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FilterableSwap struct {
	core.FilterModel

	IdempotencyKey string `json:"idempotency_key,omitempty" validate:"omitempty"`
}

type CreateSwap struct {
	OrderId              uuid.UUID                            `json:"order_id" validate:"required"`
	ReturnItems          []OrderReturnItem                    `json:"return_items" validate:"dive"`
	AdditionalItems      []CreateClaimItemAdditionalItemInput `json:"additional_items" validate:"dive"`
	ReturnShippingOption uuid.UUID                            `json:"return_shipping_option,omitempty" validate:"omitempty"`
}

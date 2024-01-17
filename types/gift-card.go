package types

import (
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type CreateGiftCardInput struct {
	OrderId    uuid.UUID  `json:"order_id,omitempty" validate:"omitempty"`
	Value      float64    `json:"value,omitempty" validate:"omitempty"`
	Balance    float64    `json:"balance,omitempty" validate:"omitempty"`
	EndsAt     *time.Time `json:"ends_at,omitempty" validate:"omitempty"`
	IsDisabled bool       `json:"is_disabled,omitempty" validate:"omitempty"`
	RegionId   uuid.UUID  `json:"region_id"`
	Metadata   core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
	TaxRate    float64    `json:"tax_rate,omitempty" validate:"omitempty"`
}

type UpdateGiftCardInput struct {
	Balance    float64    `json:"balance,omitempty" validate:"omitempty"`
	EndsAt     *time.Time `json:"ends_at,omitempty" validate:"omitempty"`
	IsDisabled bool       `json:"is_disabled,omitempty" validate:"omitempty"`
	RegionId   uuid.UUID  `json:"region_id,omitempty" validate:"omitempty"`
	Metadata   core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

type CreateGiftCardTransactionInput struct {
	GiftCardId uuid.UUID `json:"gift_card_id"`
	OrderId    uuid.UUID `json:"order_id"`
	Amount     float64   `json:"amount"`
	CreatedAt  time.Time `json:"created_at,omitempty" validate:"omitempty"`
	IsTaxable  bool      `json:"is_taxable,omitempty" validate:"omitempty"`
	TaxRate    float64   `json:"tax_rate,omitempty" validate:"omitempty"`
}

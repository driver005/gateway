package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type LineItemUpdate struct {
	Title                 string     `json:"title,omitempty" validate:"omitempty"`
	UnitPrice             float64    `json:"unit_price,omitempty" validate:"omitempty"`
	Quantity              int        `json:"quantity,omitempty" validate:"omitempty"`
	Metadata              core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
	RegionId              uuid.UUID  `json:"region_id,omitempty" validate:"omitempty"`
	VariantId             uuid.UUID  `json:"variant_id,omitempty" validate:"omitempty"`
	ShouldCalculatePrices bool       `json:"should_calculate_prices,omitempty" validate:"omitempty"`
	HasShipping           bool       `json:"has_shipping"`
}

type LineItemValidateData struct {
	Variant   *struct{ ProductId uuid.UUID }
	VariantId uuid.UUID
}

type GenerateInputData struct {
	VariantId uuid.UUID  `json:"variantId uuid.UUID"`
	Quantity  int        `json:"quantity"`
	Metadata  core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
	UnitPrice float64    `json:"unit_price,omitempty" validate:"omitempty"`
}

type GenerateLineItemContext struct {
	RegionId       uuid.UUID              `json:"region_id,omitempty" validate:"omitempty"`
	UnitPrice      float64                `json:"unit_price,omitempty" validate:"omitempty"`
	IncludesTax    bool                   `json:"includes_tax,omitempty" validate:"omitempty"`
	Metadata       core.JSONB             `json:"metadata,omitempty" validate:"omitempty"`
	CustomerId     uuid.UUID              `json:"customer_id,omitempty" validate:"omitempty"`
	OrderEditId    uuid.UUID              `json:"order_edit_id,omitempty" validate:"omitempty"`
	Cart           *models.Cart           `json:"cart,omitempty" validate:"omitempty"`
	VariantPricing *ProductVariantPricing `json:"variant_pricing,omitempty" validate:"omitempty"`
}

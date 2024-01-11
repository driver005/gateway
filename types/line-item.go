package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type GenerateInputData struct {
	VariantId uuid.UUID  `json:"variantId"`
	Quantity  int        `json:"quantity"`
	Metadata  core.JSONB `json:"metadata,omitempty"`
	UnitPrice float64    `json:"unit_price,omitempty"`
}

type GenerateLineItemContext struct {
	RegionId       uuid.UUID              `json:"region_id,omitempty"`
	UnitPrice      float64                `json:"unit_price,omitempty"`
	IncludesTax    bool                   `json:"includes_tax,omitempty"`
	Metadata       core.JSONB             `json:"metadata,omitempty"`
	CustomerId     uuid.UUID              `json:"customer_id,omitempty"`
	OrderEditId    uuid.UUID              `json:"order_edit_id,omitempty"`
	Cart           *models.Cart           `json:"cart,omitempty"`
	VariantPricing *ProductVariantPricing `json:"variant_pricing,omitempty"`
}

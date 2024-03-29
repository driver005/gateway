package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type TaxLinesMaps struct {
	LineItemsTaxLines       map[uuid.UUID][]models.LineItemTaxLine
	ShippingMethodsTaxLines map[uuid.UUID][]models.ShippingMethodTaxLine
}

type TaxServiceRate struct {
	Rate float64
	Name string
	Code string
}

type ProviderLineItemTaxLine struct {
	Rate     float64    `json:"rate"`
	Name     string     `json:"name"`
	Code     string     `json:"code"`
	ItemId   uuid.UUID  `json:"item_id"`
	Metadata core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

type ProviderShippingMethodTaxLine struct {
	Rate             float64    `json:"rate"`
	Name             string     `json:"name"`
	Code             string     `json:"code"`
	ShippingMethodId uuid.UUID  `json:"shipping_method_id"`
	Metadata         core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

type ProviderTaxLine struct {
	Rate             float64    `json:"rate"`
	Name             string     `json:"name"`
	Code             string     `json:"code"`
	ItemId           uuid.UUID  `json:"item_id"`
	ShippingMethodId uuid.UUID  `json:"shipping_method_id"`
	Metadata         core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

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
	Rate             float64
	Name             string
	Code             *string
	ItemID           uuid.UUID
	ShippingMethodID uuid.UUID
	Metadata         core.JSONB
}

type ProviderTaxLine struct {
	Rate             float64
	Name             string
	Code             string
	ItemID           uuid.UUID
	ShippingMethodID uuid.UUID
	Metadata         core.JSONB
}

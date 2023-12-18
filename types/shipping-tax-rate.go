package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FilterableShippingTaxRate struct {
	core.FilterModel

	ShippingOptionID uuid.UUID `json:"shipping_option_id,omitempty"`
	RateID           string    `json:"rate_id,omitempty"`
}

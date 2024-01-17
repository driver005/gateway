package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FilterableShippingTaxRate struct {
	core.FilterModel

	ShippingOptionId uuid.UUID `json:"shipping_option_id,omitempty" validate:"omitempty"`
	RateId           uuid.UUID `json:"rate_id,omitempty" validate:"omitempty"`
}

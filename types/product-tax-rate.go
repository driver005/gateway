package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// FilterableProductTaxRate represents the filters that can be applied when querying for shipping tax rates.
type FilterableProductTaxRate struct {
	core.FilterModel
	ProductId uuid.UUIDs `json:"product_id,omitempty" validate:"omitempty"`
	RateId    uuid.UUIDs `json:"rate_id,omitempty" validate:"omitempty"`
}

package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// FilterableProductTaxRate represents the filters that can be applied when querying for shipping tax rates.
type FilterableProductTaxRate struct {
	ProductId uuid.UUIDs
	RateId    uuid.UUIDs
	CreatedAt *core.TimeModel
	UpdatedAt *core.TimeModel
}

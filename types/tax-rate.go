package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FilterableTaxRate struct {
	core.FilterModel

	RegionId uuid.UUIDs       `json:"region_id,omitempty"`
	Code     []string         `json:"code,omitempty"`
	Name     []string         `json:"name,omitempty"`
	Rate     core.NumberModel `json:"rate,omitempty"`
}

type UpdateTaxRateInput struct {
	RegionId uuid.UUID `json:"region_id,omitempty"`
	Code     string    `json:"code,omitempty"`
	Name     string    `json:"name,omitempty"`
	Rate     float64   `json:"rate,omitempty"`
}

type CreateTaxRateInput struct {
	RegionId uuid.UUID `json:"region_id"`
	Code     string    `json:"code"`
	Name     string    `json:"name"`
	Rate     *int      `json:"rate,omitempty"`
}

type TaxRateListByConfig struct {
	RegionId uuid.UUID `json:"region_id,omitempty"`
}

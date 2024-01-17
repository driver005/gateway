package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FilterableTaxRate struct {
	core.FilterModel

	RegionId uuid.UUIDs       `json:"region_id,omitempty" validate:"omitempty"`
	Code     []string         `json:"code,omitempty" validate:"omitempty"`
	Name     []string         `json:"name,omitempty" validate:"omitempty"`
	Rate     core.NumberModel `json:"rate,omitempty" validate:"omitempty"`
}

type UpdateTaxRateInput struct {
	RegionId uuid.UUID `json:"region_id,omitempty" validate:"omitempty"`
	Code     string    `json:"code,omitempty" validate:"omitempty"`
	Name     string    `json:"name,omitempty" validate:"omitempty"`
	Rate     float64   `json:"rate,omitempty" validate:"omitempty"`
}

type CreateTaxRateInput struct {
	RegionId uuid.UUID `json:"region_id"`
	Code     string    `json:"code"`
	Name     string    `json:"name"`
	Rate     float64   `json:"rate,omitempty" validate:"omitempty"`
}

type TaxRateListByConfig struct {
	RegionId uuid.UUID `json:"region_id,omitempty" validate:"omitempty"`
}

package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FilterableProductType struct {
	core.FilterModel
	Value               string    `json:"value,omitempty" validate:"omitempty"`
	DiscountConditionId uuid.UUID `json:"discount_condition_id,omitempty" validate:"omitempty"`
}

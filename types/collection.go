package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FilterableCollectionProps struct {
	core.FilterModel
	Title               string    `json:"title,omitempty" validate:"omitempty"`
	Handle              string    `json:"handle,omitempty" validate:"omitempty"`
	DiscountConditionId uuid.UUID `json:"discount_condition_id,omitempty" validate:"omitempty"`
}
package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type CustomerGroupsBatchCustomer struct {
	Id uuid.UUID `json:"id"`
}

type CreateCustomerGroup struct {
	Name     string     `json:"name,omitempty" validate:"omitempty"`
	Metadata core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

type UpdateCustomerGroup struct {
	Name     string     `json:"name,omitempty" validate:"omitempty"`
	Metadata core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

type FilterableCustomerGroup struct {
	core.FilterModel

	Name                []string  `json:"name,omitempty" validate:"omitempty"`
	DiscountConditionId uuid.UUID `json:"discount_condition_id,omitempty" validate:"omitempty"`
}

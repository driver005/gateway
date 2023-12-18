package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FilterableCustomerGroup struct {
	core.FilterModel

	Q                   string    `json:"q,omitempty"`
	Name                []string  `json:"name,omitempty"`
	DiscountConditionID uuid.UUID `json:"discount_condition_id,omitempty"`
}

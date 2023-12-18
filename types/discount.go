package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
)

type FilterableDiscount struct {
	core.FilterModel

	Q          string                         `json:"q,omitempty"`
	IsDynamic  *bool                          `json:"is_dynamic,omitempty"`
	IsDisabled *bool                          `json:"is_disabled,omitempty"`
	Rule       *AdminGetDiscountsDiscountRule `json:"rule,omitempty"`
}

type AdminGetDiscountsDiscountRule struct {
	Type       *models.DiscountRule   `json:"type,omitempty"`
	Allocation *models.AllocationType `json:"allocation,omitempty"`
}

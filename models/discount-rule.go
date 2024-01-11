package models

import (
	"database/sql/driver"

	"github.com/driver005/gateway/core"
)

type DiscountRuleType string

const (
	DiscountRuleFixed        DiscountRuleType = "fixed"
	DiscountRulePersentage   DiscountRuleType = "percentage"
	DiscountRuleFreeShipping DiscountRuleType = "free_shipping"
)

func (pl *DiscountRuleType) Scan(value interface{}) error {
	*pl = DiscountRuleType(value.([]byte))
	return nil
}

func (pl DiscountRuleType) Value() (driver.Value, error) {
	return string(pl), nil
}

type AllocationType string

const (
	AllocationTotal AllocationType = "total"
	AllocationItem  AllocationType = "item"
)

// DiscountRule - Holds the rules that governs how a Discount is calculated when applied to a Cart.
type DiscountRule struct {
	core.Model

	// The type of the Discount, can be `fixed` for discounts that reduce the price by a fixed amount, `percentage` for percentage reductions or `free_shipping` for shipping vouchers.
	Type DiscountRuleType `json:"type"`

	// A short description of the discount
	Description string `json:"description" gorm:"default:null"`

	// The value that the discount represents; this will depend on the type of the discount
	Value float64 `json:"value"`

	// The scope that the discount should apply to.
	Allocation AllocationType `json:"allocation" gorm:"default:null"`

	// A set of conditions that can be used to limit when  the discount can be used. Available if the relation `conditions` is expanded.
	Conditions []DiscountCondition `json:"conditions" gorm:"foreignKey:id"`
}

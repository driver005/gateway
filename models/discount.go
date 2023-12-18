package models

import (
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// Discount - Represents a discount that can be applied to a cart for promotional purposes.
type Discount struct {
	core.Model

	// A unique code for the discount - this will be used by the customer to apply the discount
	Code string `json:"code"`

	// A flag to indicate if multiple instances of the discount can be generated. I.e. for newsletter discounts
	IsDynamic bool `json:"is_dynamic"`

	// The Discount Rule that governs the behaviour of the Discount
	RuleId uuid.NullUUID `json:"rule_id" gorm:"default:null"`

	Rule *DiscountRule `json:"rule" gorm:"foreignKey:id;references:rule_id"`

	// Whether the Discount has been disabled. Disabled discounts cannot be applied to carts
	IsDisabled bool `json:"is_disabled" gorm:"default:null"`

	// The Discount that the discount was created from. This will always be a dynamic discount
	ParentDiscountId uuid.NullUUID `json:"parent_discount_id" gorm:"default:null"`

	ParentDiscount *Discount `json:"parent_discount" gorm:"foreignKey:id;references:parent_discount_id"`

	// The time at which the discount can be used.
	StartsAt time.Time `json:"starts_at" gorm:"default:null"`

	// The time at which the discount can no longer be used.
	EndsAt time.Time `json:"ends_at" gorm:"default:null"`

	// Duration the discount runs between
	ValidDuration string `json:"valid_duration" gorm:"default:null"`

	// The Regions in which the Discount can be used. Available if the relation `regions` is expanded.
	Regions []Region `json:"regions" gorm:"foreignKey:id"`

	// The maximum number of times that a discount can be used.
	UsageLimit int32 `json:"usage_limit" gorm:"default:null"`

	// The number of times a discount has been used.
	UsageCount int32 `json:"usage_count" gorm:"default:null"`
}

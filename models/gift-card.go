package models

import (
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// GiftCard - Gift Cards are redeemable and represent a value that can be used towards the payment of an Order.
type GiftCard struct {
	core.Model

	// The unique code that identifies the Gift Card. This is used by the Customer to redeem the value of the Gift Card.
	Code string `json:"code"`

	// The value that the Gift Card represents.
	Value float64 `json:"value"`

	// The remaining value on the Gift Card.
	Balance float64 `json:"balance"`

	// The id of the Region in which the Gift Card is available.
	RegionId uuid.NullUUID `json:"region_id"`

	// A region object. Available if the relation `region` is expanded.
	Region *Region `json:"region" gorm:"foreignKey:id;references:region_id"`

	// The id of the Order that the Gift Card was purchased in.
	OrderId uuid.NullUUID `json:"order_id" gorm:"default:null"`

	// An order object. Available if the relation `order` is expanded.
	Order *Order `json:"order" gorm:"foreignKey:id;references:order_id"`

	// Whether the Gift Card has been disabled. Disabled Gift Cards cannot be applied to carts.
	IsDisabled bool `json:"is_disabled" gorm:"default:null"`

	TaxRate float64 `json:"tax_rate"`

	// The time at which the Gift Card can no longer be used.
	EndsAt *time.Time `json:"ends_at" gorm:"default:null"`
}

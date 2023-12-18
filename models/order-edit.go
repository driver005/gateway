package models

import (
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// OrderEdit - Order edit keeps track of order items changes.
type OrderEdit struct {
	core.Model

	// The ID of the order that is edited
	OrderId uuid.NullUUID `json:"order_id"`

	Order *Order `json:"order" gorm:"foreignKey:id;references:order_id"`

	// Line item changes array.
	Changes []OrderItemChange `json:"changes" gorm:"foreignKey:id"`

	// An optional note with additional details about the order edit.
	InternalNote string `json:"internal_note" gorm:"default:null"`

	// The unique identifier of the user or customer who created the order edit.
	CreatedBy string `json:"created_by"`

	// The unique identifier of the user or customer who requested the order edit.
	RequestedBy string `json:"requested_by" gorm:"default:null"`

	// The date with timezone at which the edit was requested.
	RequestedAt time.Time `json:"requested_at" gorm:"default:null"`

	// The unique identifier of the user or customer who confirmed the order edit.
	ConfirmedBy string `json:"confirmed_by" gorm:"default:null"`

	// The date with timezone at which the edit was confirmed.
	ConfirmedAt time.Time `json:"confirmed_at" gorm:"default:null"`

	// The unique identifier of the user or customer who declined the order edit.
	DeclinedBy string `json:"declined_by" gorm:"default:null"`

	// The date with timezone at which the edit was declined.
	DeclinedAt time.Time `json:"declined_at" gorm:"default:null"`

	// An optional note why  the order edit is declined.
	DeclinedReason string `json:"declined_reason" gorm:"default:null"`

	// The subtotal for line items computed from changes.
	Subtotal int32 `json:"subtotal" gorm:"default:null"`

	// The total of discount
	DiscountTotal int32 `json:"discount_total" gorm:"default:null"`

	// The total of tax
	TaxTotal int32 `json:"tax_total" gorm:"default:null"`

	// The total amount of the edited order.
	Total int32 `json:"total" gorm:"default:null"`

	// The difference between the total amount of the order and total amount of edited order.
	DifferenceDue int32 `json:"difference_due" gorm:"default:null"`

	// Computed line items from the changes.
	Items []LineItem `json:"items" gorm:"foreignKey:id"`
}

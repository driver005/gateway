package models

import (
	"database/sql/driver"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// DraftOrder - Represents a draft order
type DraftOrder struct {
	core.Model

	// The status of the draft order
	Status DraftOrderStatus `json:"status" gorm:"default:open"`

	// The draft order's display ID
	DisplayId uuid.NullUUID `json:"display_id" gorm:"default:null"`

	// The ID of the cart associated with the draft order.
	CartId uuid.NullUUID `json:"cart_id" gorm:"default:null"`

	// A cart object. Available if the relation `cart` is expanded.
	Cart Cart `json:"cart" gorm:"default:null"`

	// The ID of the order associated with the draft order.
	OrderId uuid.NullUUID `json:"order_id" gorm:"default:null"`

	// An order object. Available if the relation `order` is expanded.
	Order *Order `json:"order" gorm:"foreignKey:id;references:order_id"`

	// The date the draft order was canceled at.
	CanceledAt time.Time `json:"canceled_at" gorm:"default:null"`

	// The date the draft order was completed at.
	CompletedAt time.Time `json:"completed_at" gorm:"default:null"`

	// Whether to send the customer notifications regarding order updates.
	NoNotificationOrder bool `json:"no_notification_order" gorm:"default:null"`

	// Randomly generated key used to continue the completion of the cart associated with the draft order in case of failure.
	IdempotencyKey string `json:"idempotency_key" gorm:"default:null"`
}

// The status of the Price List
type DraftOrderStatus string

// Defines values for DraftOrderStatus.
const (
	DraftOrderStatusCompleted DraftOrderStatus = "completed"
	DraftOrderStatusOpen      DraftOrderStatus = "open"
)

func (pl *DraftOrderStatus) Scan(value interface{}) error {
	*pl = DraftOrderStatus(value.([]byte))
	return nil
}

func (pl DraftOrderStatus) Value() (driver.Value, error) {
	return string(pl), nil
}

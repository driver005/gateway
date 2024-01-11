package models

import (
	"database/sql/driver"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// Return - Return orders hold information about Line Items that a Customer wishes to send back, along with how the items will be returned. Returns can be used as part of a Swap.
type Return struct {
	core.Model

	// Status of the Return.
	Status ReturnStatus `json:"status" gorm:"default:null"`

	// The Return Items that will be shipped back to the warehouse. Available if the relation `items` is expanded.
	Items []ReturnItem `json:"items" gorm:"foreignKey:return_id"`

	// The ID of the Swap that the Return is a part of.
	SwapId uuid.NullUUID `json:"swap_id" gorm:"default:null"`

	// A swap object. Available if the relation `swap` is expanded.
	Swap *Swap `json:"swap" gorm:"foreignKey:id;references:swap_id"`

	// The ID of the Order that the Return is made from.
	OrderId uuid.NullUUID `json:"order_id" gorm:"default:null"`

	// An order object. Available if the relation `order` is expanded.
	Order *Order `json:"order" gorm:"foreignKey:id;references:order_id"`

	// The ID of the Claim that the Return is a part of.
	ClaimOrderId uuid.NullUUID `json:"claim_order_id" gorm:"default:null"`

	// A claim order object. Available if the relation `claim_order` is expanded.
	ClaimOrder *ClaimOrder `json:"claim_order" gorm:"foreignKey:id;references:claim_order_id"`

	// The Shipping Method that will be used to send the Return back. Can be null if the Customer facilitates the return shipment themselves. Available if the relation `shipping_method` is expanded.
	ShippingMethod *ShippingMethod `json:"shipping_method" gorm:"foreignKey:id"`

	// Data about the return shipment as provided by the Fulfilment Provider that handles the return shipment.
	ShippingData core.JSONB `json:"shipping_data" gorm:"default:null"`

	// The amount that should be refunded as a result of the return.
	RefundAmount float64 `json:"refund_amount"`

	// When set to true, no notification will be sent related to this return.
	NoNotification bool `json:"no_notification" gorm:"default:null"`

	LocationId string `json:"location_id"`

	// Randomly generated key used to continue the completion of the return in case of failure.
	IdempotencyKey string `json:"idempotency_key" gorm:"default:null"`

	// The date with timezone at which the return was received.
	ReceivedAt *time.Time `json:"received_at" gorm:"default:null"`
}

// ReturnStatus represents the status of a return
type ReturnStatus string

// Enum values for ReturnStatus
const (
	ReturnRequested      ReturnStatus = "requested"
	ReturnReceived       ReturnStatus = "received"
	ReturnRequiresAction ReturnStatus = "requires_action"
	ReturnCanceled       ReturnStatus = "canceled"
)

func (pl *ReturnStatus) Scan(value interface{}) error {
	*pl = ReturnStatus(value.([]byte))
	return nil
}

func (pl ReturnStatus) Value() (driver.Value, error) {
	return string(pl), nil
}

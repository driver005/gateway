package models

import (
	"database/sql/driver"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:DraftOrder
// title: "DraftOrder"
// description: "A draft order is created by an admin without direct involvement of the customer. Once its payment is marked as captured, it is transformed into an order."
// type: object
// required:
//   - canceled_at
//   - cart_id
//   - completed_at
//   - created_at
//   - display_id
//   - id
//   - idempotency_key
//   - metadata
//   - no_notification_order
//   - order_id
//   - status
//   - updated_at
//
// properties:
//
//	id:
//	  description: The draft order's ID
//	  type: string
//	  example: dorder_01G8TJFKBG38YYFQ035MSVG03C
//	status:
//	  description: The status of the draft order. It's changed to `completed` when it's transformed to an order.
//	  type: string
//	  enum:
//	    - open
//	    - completed
//	  default: open
//	display_id:
//	  description: The draft order's display ID
//	  type: string
//	  example: 2
//	cart_id:
//	  description: The ID of the cart associated with the draft order.
//	  nullable: true
//	  type: string
//	  example: cart_01G8ZH853Y6TFXWPG5EYE81X63
//	cart:
//	  description: The details of the cart associated with the draft order.
//	  x-expandable: "cart"
//	  nullable: true
//	  $ref: "#/components/schemas/Cart"
//	order_id:
//	  description: The ID of the order created from the draft order when its payment is captured.
//	  nullable: true
//	  type: string
//	  example: order_01G8TJSYT9M6AVS5N4EMNFS1EK
//	order:
//	  description: The details of the order created from the draft order when its payment is captured.
//	  x-expandable: "order"
//	  nullable: true
//	  $ref: "#/components/schemas/Order"
//	canceled_at:
//	  description: The date the draft order was canceled at.
//	  nullable: true
//	  type: string
//	  format: date-time
//	completed_at:
//	  description: The date the draft order was completed at.
//	  nullable: true
//	  type: string
//	  format: date-time
//	no_notification_order:
//	  description: Whether to send the customer notifications regarding order updates.
//	  nullable: true
//	  type: boolean
//	  example: false
//	idempotency_key:
//	  description: Randomly generated key used to continue the completion of the cart associated with the draft order in case of failure.
//	  nullable: true
//	  type: string
//	  externalDocs:
//	    url: https://docs.medusajs.com/development/idempotency-key/overview.md
//	    description: Learn more how to use the idempotency key.
//	created_at:
//	  description: The date with timezone at which the resource was created.
//	  type: string
//	  format: date-time
//	updated_at:
//	  description: The date with timezone at which the resource was updated.
//	  type: string
//	  format: date-time
//	metadata:
//	  description: An optional key-value map with additional details
//	  nullable: true
//	  type: object
//	  example: {car: "white"}
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type DraftOrder struct {
	core.BaseModel

	Status              DraftOrderStatus `json:"status" gorm:"column:status;default:'open'"`
	DisplayId           uuid.NullUUID    `json:"display_id" gorm:"column:display_id"`
	CartId              uuid.NullUUID    `json:"cart_id" gorm:"column:cart_id"`
	Cart                *Cart            `json:"cart" gorm:"foreignKey:CartId"`
	OrderId             uuid.NullUUID    `json:"order_id" gorm:"column:order_id"`
	Order               *Order           `json:"order" gorm:"foreignKey:OrderId"`
	CanceledAt          *time.Time       `json:"canceled_at" gorm:"column:canceled_at"`
	CompletedAt         *time.Time       `json:"completed_at" gorm:"column:completed_at"`
	NoNotificationOrder bool             `json:"no_notification_order" gorm:"column:no_notification_order"`
	IdempotencyKey      string           `json:"idempotency_key" gorm:"column:idempotency_key"`
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

package models

import (
	"database/sql/driver"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:Return
// title: "Return"
// description: "A Return holds information about Line Items that a Customer wishes to send back, along with how the items will be returned. Returns can also be used as part of a Swap or a Claim."
// type: object
// required:
//   - claim_order_id
//   - created_at
//   - id
//   - idempotency_key
//   - location_id
//   - metadata
//   - no_notification
//   - order_id
//   - received_at
//   - refund_amount
//   - shipping_data
//   - status
//   - swap_id
//   - updated_at
//
// properties:
//
//	id:
//	  description: The return's ID
//	  type: string
//	  example: ret_01F0YET7XPCMF8RZ0Y151NZV2V
//	status:
//	  description: Status of the Return.
//	  type: string
//	  enum:
//	    - requested
//	    - received
//	    - requires_action
//	    - canceled
//	  default: requested
//	items:
//	  description: The details of the items that the customer is returning.
//	  type: array
//	  x-expandable: "items"
//	  items:
//	    $ref: "#/components/schemas/ReturnItem"
//	swap_id:
//	  description: The ID of the swap that the return may belong to.
//	  nullable: true
//	  type: string
//	  example: null
//	swap:
//	  description: The details of the swap that the return may belong to.
//	  x-expandable: "swap"
//	  nullable: true
//	  $ref: "#/components/schemas/Swap"
//	claim_order_id:
//	  description: The ID of the claim that the return may belong to.
//	  nullable: true
//	  type: string
//	  example: null
//	claim_order:
//	  description: The details of the claim that the return may belong to.
//	  x-expandable: "claim_order"
//	  nullable: true
//	  $ref: "#/components/schemas/ClaimOrder"
//	order_id:
//	  description: The ID of the order that the return was created for.
//	  nullable: true
//	  type: string
//	  example: order_01G8TJSYT9M6AVS5N4EMNFS1EK
//	order:
//	  description: The details of the order that the return was created for.
//	  x-expandable: "order"
//	  nullable: true
//	  $ref: "#/components/schemas/Order"
//	shipping_method:
//	  description: The details of the Shipping Method that will be used to send the Return back. Can be null if the Customer will handle the return shipment themselves.
//	  x-expandable: "shipping_method"
//	  nullable: true
//	  $ref: "#/components/schemas/ShippingMethod"
//	shipping_data:
//	  description: Data about the return shipment as provided by the Fulfilment Provider that handles the return shipment.
//	  nullable: true
//	  type: object
//	  example: {}
//	location_id:
//	  description: The ID of the stock location the return will be added back.
//	  nullable: true
//	  type: string
//	  example: sloc_01G8TJSYT9M6AVS5N4EMNFS1EK
//	refund_amount:
//	  description: The amount that should be refunded as a result of the return.
//	  type: integer
//	  example: 1000
//	no_notification:
//	  description: When set to true, no notification will be sent related to this return.
//	  nullable: true
//	  type: boolean
//	  example: false
//	idempotency_key:
//	  description: Randomly generated key used to continue the completion of the return in case of failure.
//	  nullable: true
//	  type: string
//	  externalDocs:
//	    url: https://docs.medusajs.com/development/idempotency-key/overview.md
//	    description: Learn more how to use the idempotency key.
//	received_at:
//	  description: The date with timezone at which the return was received.
//	  nullable: true
//	  type: string
//	  format: date-time
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
type Return struct {
	core.Model

	Status         ReturnStatus    `json:"status" gorm:"column:status;default:'requested'"`
	Items          []ReturnItem    `json:"items" gorm:"foreignKey:ReturnId"`
	SwapId         uuid.NullUUID   `json:"swap_id" gorm:"column:swap_id"`
	Swap           *Swap           `json:"swap" gorm:"foreignKey:SwapId"`
	OrderId        uuid.NullUUID   `json:"order_id" gorm:"column:order_id"`
	Order          *Order          `json:"order" gorm:"foreignKey:OrderId"`
	ClaimOrderId   uuid.NullUUID   `json:"claim_order_id" gorm:"column:claim_order_id"`
	ClaimOrder     *ClaimOrder     `json:"claim_order" gorm:"foreignKey:ClaimOrderId"`
	ShippingMethod *ShippingMethod `json:"shipping_method" gorm:"foreignKey:Id"`
	ShippingData   core.JSONB      `json:"shipping_data" gorm:"column:shipping_data"`
	RefundAmount   float64         `json:"refund_amount" gorm:"column:refund_amount"`
	NoNotification bool            `json:"no_notification" gorm:"column:no_notification"`
	LocationId     uuid.NullUUID   `json:"location_id" gorm:"column:location_id"`
	IdempotencyKey string          `json:"idempotency_key" gorm:"column:idempotency_key"`
	ReceivedAt     *time.Time      `json:"received_at" gorm:"column:received_at"`
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

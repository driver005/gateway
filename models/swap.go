package models

import (
	"database/sql/driver"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:Swap
// title: "Swap"
// description: "A swap can be created when a Customer wishes to exchange Products that they have purchased with different Products. It consists of a Return of previously purchased Products and a Fulfillment of new Products. It also includes information on any additional payment or refund required based on the difference between the exchanged products."
// type: object
// required:
//   - allow_backorder
//   - canceled_at
//   - cart_id
//   - confirmed_at
//   - created_at
//   - deleted_at
//   - difference_due
//   - fulfillment_status
//   - id
//   - idempotency_key
//   - metadata
//   - no_notification
//   - order_id
//   - payment_status
//   - shipping_address_id
//   - updated_at
//
// properties:
//
//	id:
//	  description: The swap's ID
//	  type: string
//	  example: swap_01F0YET86Y9G92D3YDR9Y6V676
//	fulfillment_status:
//	  description: The status of the Fulfillment of the Swap.
//	  type: string
//	  enum:
//	    - not_fulfilled
//	    - fulfilled
//	    - shipped
//	    - partially_shipped
//	    - canceled
//	    - requires_action
//	  example: not_fulfilled
//	payment_status:
//	  description: The status of the Payment of the Swap. The payment may either refer to the refund of an amount or the authorization of a new amount.
//	  type: string
//	  enum:
//	    - not_paid
//	    - awaiting
//	    - captured
//	    - confirmed
//	    - canceled
//	    - difference_refunded
//	    - partially_refunded
//	    - refunded
//	    - requires_action
//	  example: not_paid
//	order_id:
//	  description: The ID of the order that the swap belongs to.
//	  type: string
//	  example: order_01G8TJSYT9M6AVS5N4EMNFS1EK
//	order:
//	  description: The details of the order that the swap belongs to.
//	  x-expandable: "order"
//	  nullable: true
//	  $ref: "#/components/schemas/Order"
//	additional_items:
//	  description: The details of the new products to send to the customer, represented as line items.
//	  type: array
//	  x-expandable: "additional_items"
//	  items:
//	    $ref: "#/components/schemas/LineItem"
//	return_order:
//	  description: The details of the return that belongs to the swap, which holds the details on the items being returned.
//	  x-expandable: "return_order"
//	  nullable: true
//	  $ref: "#/components/schemas/Return"
//	fulfillments:
//	  description: The details of the fulfillments that are used to send the new items to the customer.
//	  x-expandable: "fulfillments"
//	  type: array
//	  items:
//	    $ref: "#/components/schemas/Fulfillment"
//	payment:
//	  description: The details of the additional payment authorized by the customer when `difference_due` is positive.
//	  x-expandable: "payment"
//	  nullable: true
//	  $ref: "#/components/schemas/Payment"
//	difference_due:
//	  description: The difference amount between the order’s original total and the new total imposed by the swap. If its value is negative, a refund must be issues to the customer. If it's positive, additional payment must be authorized by the customer. Otherwise, no payment processing is required.
//	  nullable: true
//	  type: integer
//	  example: 0
//	shipping_address_id:
//	  description: The Address to send the new Line Items to - in most cases this will be the same as the shipping address on the Order.
//	  nullable: true
//	  type: string
//	  example: addr_01G8ZH853YPY9B94857DY91YGW
//	shipping_address:
//	  description: The details of the shipping address that the new items should be sent to.
//	  x-expandable: "shipping_address"
//	  nullable: true
//	  $ref: "#/components/schemas/Address"
//	shipping_methods:
//	  description: The details of the shipping methods used to fulfill the additional items purchased.
//	  type: array
//	  x-expandable: "shipping_methods"
//	  items:
//	    $ref: "#/components/schemas/ShippingMethod"
//	cart_id:
//	  description: The ID of the cart that the customer uses to complete the swap.
//	  nullable: true
//	  type: string
//	  example: cart_01G8ZH853Y6TFXWPG5EYE81X63
//	cart:
//	  description: The details of the cart that the customer uses to complete the swap.
//	  x-expandable: "cart"
//	  nullable: true
//	  $ref: "#/components/schemas/Cart"
//	confirmed_at:
//	  description: The date with timezone at which the Swap was confirmed by the Customer.
//	  nullable: true
//	  type: string
//	  format: date-time
//	canceled_at:
//	  description: The date with timezone at which the Swap was canceled.
//	  nullable: true
//	  type: string
//	  format: date-time
//	no_notification:
//	  description: If set to true, no notification will be sent related to this swap
//	  nullable: true
//	  type: boolean
//	  example: false
//	allow_backorder:
//	  description: If true, swaps can be completed with items out of stock
//	  type: boolean
//	  default: false
//	idempotency_key:
//	  description: Randomly generated key used to continue the completion of the swap in case of failure.
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
//	deleted_at:
//	  description: The date with timezone at which the resource was deleted.
//	  nullable: true
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
type Swap struct {
	core.Model

	FulfillmentStatus SwapFulfillmentStatus `json:"fulfillment_status" gorm:"column:fulfillment_status"`
	PaymentStatus     SwapPaymentStatus     `json:"payment_status" gorm:"column:payment_status"`
	OrderId           uuid.NullUUID         `json:"order_id" gorm:"column:order_id"`
	Order             *Order                `json:"order" gorm:"foreignKey:OrderId"`
	AdditionalItems   []LineItem            `json:"additional_items" gorm:"foreignKey:Id"`
	ReturnOrder       *Return               `json:"return_order" gorm:"foreignKey:Id"`
	Fulfillments      []Fulfillment         `json:"fulfillments" gorm:"foreignKey:Id"`
	Payment           *Payment              `json:"payment" gorm:"foreignKey:Id"`
	DifferenceDue     float64               `json:"difference_due" gorm:"column:difference_due"`
	ShippingAddressId uuid.NullUUID         `json:"shipping_address_id" gorm:"column:shipping_address_id"`
	ShippingAddress   *Address              `json:"shipping_address" gorm:"foreignKey:ShippingAddressId"`
	ShippingMethods   []ShippingMethod      `json:"shipping_methods" gorm:"foreignKey:Id"`
	CartId            uuid.NullUUID         `json:"cart_id" gorm:"column:cart_id"`
	Cart              *Cart                 `json:"cart" gorm:"foreignKey:CartId"`
	AllowBackorder    bool                  `json:"allow_backorder" gorm:"column:allow_backorder;default:false"`
	IdempotencyKey    string                `json:"idempotency_key" gorm:"column:idempotency_key"`
	ConfirmedAt       *time.Time            `json:"confirmed_at" gorm:"column:confirmed_at"`
	CanceledAt        *time.Time            `json:"canceled_at" gorm:"column:canceled_at"`
	NoNotification    bool                  `json:"no_notification" gorm:"column:no_notification"`
}

type SwapFulfillmentStatus string

const (
	SwapFulfillmentNotFulfilled     SwapFulfillmentStatus = "not_fulfilled"
	SwapFulfillmentFulfilled        SwapFulfillmentStatus = "fulfilled"
	SwapFulfillmentShipped          SwapFulfillmentStatus = "shipped"
	SwapFulfillmentPartiallyShipped SwapFulfillmentStatus = "partially_shipped"
	SwapFulfillmentCanceled         SwapFulfillmentStatus = "canceled"
	SwapFulfillmentRequiresAction   SwapFulfillmentStatus = "requires_action"
)

func (pl *SwapFulfillmentStatus) Scan(value interface{}) error {
	*pl = SwapFulfillmentStatus(value.([]byte))
	return nil
}

func (pl SwapFulfillmentStatus) Value() (driver.Value, error) {
	return string(pl), nil
}

type SwapPaymentStatus string

const (
	SwapPaymentNotPaid            SwapPaymentStatus = "not_paid"
	SwapPaymentAwaiting           SwapPaymentStatus = "awaiting"
	SwapPaymentCaptured           SwapPaymentStatus = "captured"
	SwapPaymentConfirmed          SwapPaymentStatus = "confirmed"
	SwapPaymentCanceled           SwapPaymentStatus = "canceled"
	SwapPaymentDifferenceRefunded SwapPaymentStatus = "difference_refunded"
	SwapPaymentPartiallyRefunded  SwapPaymentStatus = "partially_refunded"
	SwapPaymentRefunded           SwapPaymentStatus = "refunded"
	SwapPaymentRequiresAction     SwapPaymentStatus = "requires_action"
)

func (pl *SwapPaymentStatus) Scan(value interface{}) error {
	*pl = SwapPaymentStatus(value.([]byte))
	return nil
}

func (pl SwapPaymentStatus) Value() (driver.Value, error) {
	return string(pl), nil
}

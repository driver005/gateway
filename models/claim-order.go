package models

import (
	"database/sql/driver"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:ClaimOrder
// title: "Claim"
// description: "A Claim represents a group of faulty or missing items. It consists of claim items that refer to items in the original order that should be replaced or refunded. It also includes details related to shipping and fulfillment."
// type: object
// required:
//   - canceled_at
//   - created_at
//   - deleted_at
//   - fulfillment_status
//   - id
//   - idempotency_key
//   - metadata
//   - no_notification
//   - order_id
//   - payment_status
//   - refund_amount
//   - shipping_address_id
//   - type
//   - updated_at
//
// properties:
//
//	id:
//	  description: The claim's ID
//	  type: string
//	  example: claim_01G8ZH853Y6TFXWPG5EYE81X63
//	type:
//	  description: The claim's type
//	  type: string
//	  enum:
//	    - refund
//	    - replace
//	payment_status:
//	  description: The status of the claim's payment
//	  type: string
//	  enum:
//	    - na
//	    - not_refunded
//	    - refunded
//	  default: na
//	fulfillment_status:
//	  description: The claim's fulfillment status
//	  type: string
//	  enum:
//	    - not_fulfilled
//	    - partially_fulfilled
//	    - fulfilled
//	    - partially_shipped
//	    - shipped
//	    - partially_returned
//	    - returned
//	    - canceled
//	    - requires_action
//	  default: not_fulfilled
//	claim_items:
//	  description: The details of the items that should be replaced or refunded.
//	  type: array
//	  x-expandable: "claim_items"
//	  items:
//	    $ref: "#/components/schemas/ClaimItem"
//	additional_items:
//	  description: The details of the new items to be shipped when the claim's type is `replace`
//	  type: array
//	  x-expandable: "additional_items"
//	  items:
//	    $ref: "#/components/schemas/LineItem"
//	order_id:
//	  description: The ID of the order that the claim comes from.
//	  type: string
//	  example: order_01G8TJSYT9M6AVS5N4EMNFS1EK
//	order:
//	  description: The details of the order that this claim was created for.
//	  x-expandable: "order"
//	  nullable: true
//	  $ref: "#/components/schemas/Order"
//	return_order:
//	  description: The details of the return associated with the claim if the claim's type is `replace`.
//	  x-expandable: "return_order"
//	  nullable: true
//	  $ref: "#/components/schemas/Return"
//	shipping_address_id:
//	  description: The ID of the address that the new items should be shipped to
//	  nullable: true
//	  type: string
//	  example: addr_01G8ZH853YPY9B94857DY91YGW
//	shipping_address:
//	  description: The details of the address that new items should be shipped to.
//	  x-expandable: "shipping_address"
//	  nullable: true
//	  $ref: "#/components/schemas/Address"
//	shipping_methods:
//	  description: The details of the shipping methods that the claim order will be shipped with.
//	  type: array
//	  x-expandable: "shipping_methods"
//	  items:
//	    $ref: "#/components/schemas/ShippingMethod"
//	fulfillments:
//	  description: The fulfillments of the new items to be shipped
//	  type: array
//	  x-expandable: "fulfillments"
//	  items:
//	    $ref: "#/components/schemas/Fulfillment"
//	refund_amount:
//	  description: The amount that will be refunded in conjunction with the claim
//	  nullable: true
//	  type: integer
//	  example: 1000
//	canceled_at:
//	  description: The date with timezone at which the claim was canceled.
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
//	no_notification:
//	  description: Flag for describing whether or not notifications related to this should be send.
//	  nullable: true
//	  type: boolean
//	  example: false
//	idempotency_key:
//	  description: Randomly generated key used to continue the completion of the cart associated with the claim in case of failure.
//	  nullable: true
//	  type: string
//	  externalDocs:
//	    url: https://docs.medusajs.com/development/idempotency-key/overview.md
//	    description: Learn more how to use the idempotency key.
type ClaimOrder struct {
	core.Model

	Type              ClaimStatus            `json:"type"  gorm:"column:type;default:'na'"`
	PaymentStatus     ClaimPaymentStatus     `json:"payment_status"  gorm:"column:payment_status;default:'not_fulfilled'"`
	FulfillmentStatus ClaimFulfillmentStatus `json:"fulfillment_status"  gorm:"column:fulfillment_status"`
	ClaimItems        []ClaimItem            `json:"claim_items"  gorm:"column:claim_items;foreignKey:Id"`
	AdditionalItems   []LineItem             `json:"additional_items"  gorm:"column:additional_items;foreignKey:Id"`
	OrderId           uuid.NullUUID          `json:"order_id"  gorm:"column:order_id"`
	Order             *Order                 `json:"order"  gorm:"column:order;foreignKey:OrderId"`
	ReturnOrder       *Return                `json:"return_order"  gorm:"column:return_order;foreignKey:Id"`
	ShippingAddressId uuid.NullUUID          `json:"shipping_address_id"  gorm:"column:shipping_address_id"`
	ShippingAddress   *Address               `json:"shipping_address"  gorm:"column:shipping_address;foreignKey:ShippingAddressId"`
	ShippingMethods   []ShippingMethod       `json:"shipping_methods"  gorm:"column:shipping_methods;foreignKey:Id"`
	Fulfillments      []Fulfillment          `json:"fulfillments"  gorm:"column:fulfillments;foreignKey:Id"`
	RefundAmount      float64                `json:"refund_amount"  gorm:"column:refund_amount"`
	CanceledAt        *time.Time             `json:"canceled_at"  gorm:"column:canceled_at"`
	NoNotification    bool                   `json:"no_notification"  gorm:"column:no_notification"`
	IdempotencyKey    string                 `json:"idempotency_key"  gorm:"column:idempotency_key"`
}

// The status of the Price List
type ClaimStatus string

// Defines values for ClaimStatus.
const (
	ClaimStatusReplace ClaimStatus = "replace"
	ClaimStatusRefund  ClaimStatus = "refund"
)

func (pl *ClaimStatus) Scan(value interface{}) error {
	*pl = ClaimStatus(value.([]byte))
	return nil
}

func (pl ClaimStatus) Value() (driver.Value, error) {
	return string(pl), nil
}

type ClaimPaymentStatus string

const (
	ClaimPaymentStatusNa          ClaimPaymentStatus = "na"
	ClaimPaymentStatusNotRefunded ClaimPaymentStatus = "not_refunded"
	ClaimPaymentStatusRefunded    ClaimPaymentStatus = "refunded"
)

func (pl *ClaimPaymentStatus) Scan(value interface{}) error {
	*pl = ClaimPaymentStatus(value.([]byte))
	return nil
}

func (pl ClaimPaymentStatus) Value() (driver.Value, error) {
	return string(pl), nil
}

type ClaimFulfillmentStatus string

const (
	ClaimFulfillmentStatusNotFulfilled       ClaimFulfillmentStatus = "not_fulfilled"
	ClaimFulfillmentStatusPartiallyFulfilled ClaimFulfillmentStatus = "partially_fulfilled"
	ClaimFulfillmentStatusFulfilled          ClaimFulfillmentStatus = "fulfilled"
	ClaimFulfillmentStatusPartiallyShipped   ClaimFulfillmentStatus = "partially_shipped"
	ClaimFulfillmentStatusShipped            ClaimFulfillmentStatus = "shipped"
	ClaimFulfillmentStatusPartiallyReturned  ClaimFulfillmentStatus = "partially_returned"
	ClaimFulfillmentStatusReturned           ClaimFulfillmentStatus = "returned"
	ClaimFulfillmentStatusCanceled           ClaimFulfillmentStatus = "canceled"
	ClaimFulfillmentStatusRequiresAction     ClaimFulfillmentStatus = "requires_action"
)

func (pl *ClaimFulfillmentStatus) Scan(value interface{}) error {
	*pl = ClaimFulfillmentStatus(value.([]byte))
	return nil
}

func (pl ClaimFulfillmentStatus) Value() (driver.Value, error) {
	return string(pl), nil
}

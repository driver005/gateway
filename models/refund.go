package models

import (
	"database/sql/driver"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:Refund
// title: "Refund"
// description: "A refund represents an amount of money transfered back to the customer for a given reason. Refunds may occur in relation to Returns, Swaps and Claims, but can also be initiated by an admin for an order."
// type: object
// required:
//   - amount
//   - created_at
//   - id
//   - idempotency_key
//   - metadata
//   - note
//   - order_id
//   - payment_id
//   - reason
//   - updated_at
//
// properties:
//
//	id:
//	  description: The refund's ID
//	  type: string
//	  example: ref_01G1G5V27GYX4QXNARRQCW1N8T
//	order_id:
//	  description: The ID of the order this refund was created for.
//	  nullable: true
//	  type: string
//	  example: order_01G8TJSYT9M6AVS5N4EMNFS1EK
//	order:
//	  description: The details of the order this refund was created for.
//	  x-expandable: "order"
//	  nullable: true
//	  $ref: "#/components/schemas/Order"
//	payment_id:
//	  description: The payment's ID, if available.
//	  nullable: true
//	  type: string
//	  example: pay_01G8ZCC5W42ZNY842124G7P5R9
//	payment:
//	  description: The details of the payment associated with the refund.
//	  x-expandable: "payment"
//	  nullable: true
//	  $ref: "#/components/schemas/Payment"
//	amount:
//	  description: The amount that has be refunded to the Customer.
//	  type: integer
//	  example: 1000
//	note:
//	  description: An optional note explaining why the amount was refunded.
//	  nullable: true
//	  type: string
//	  example: I didn't like it
//	reason:
//	  description: The reason given for the Refund, will automatically be set when processed as part of a Swap, Claim or Return.
//	  type: string
//	  enum:
//	    - discount
//	    - return
//	    - swap
//	    - claim
//	    - other
//	  example: return
//	idempotency_key:
//	  description: Randomly generated key used to continue the completion of the refund in case of failure.
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
type Refund struct {
	core.BaseModel

	OrderId        uuid.NullUUID `json:"order_id" gorm:"column:order_id"`
	PaymentId      uuid.NullUUID `json:"payment_id" gorm:"column:payment_id"`
	Order          *Order        `json:"order" gorm:"foreignKey:OrderId"`
	Payment        *Payment      `json:"payment" gorm:"foreignKey:PaymentId"`
	Amount         float64       `json:"amount" gorm:"column:amount"`
	Note           string        `json:"note" gorm:"column:note"`
	Reason         RefundReason  `json:"reason" gorm:"column:reason"`
	IdempotencyKey string        `json:"idempotency_key" gorm:"column:idempotency_key"`
}

// The status of the product
type RefundReason string

const (
	RefundReasonDiscount RefundReason = "discount"
	RefundReasonReturn   RefundReason = "return"
	RefundReasonSwap     RefundReason = "swap"
	RefundReasonClaim    RefundReason = "claim"
	RefundReasonOther    RefundReason = "other"
)

func (ps *RefundReason) Scan(value interface{}) error {
	*ps = RefundReason(value.([]byte))
	return nil
}

func (ps RefundReason) Value() (driver.Value, error) {
	return string(ps), nil
}

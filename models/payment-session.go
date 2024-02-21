package models

import (
	"database/sql/driver"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:PaymentSession
// title: "Payment Session"
// description: "A Payment Session is created when a Customer initilizes the checkout flow, and can be used to hold the state of a payment flow. Each Payment Session is controlled by a Payment Provider, which is responsible for the communication with external payment services. Authorized Payment Sessions will eventually get promoted to Payments to indicate that they are authorized for payment processing such as capture or refund. Payment sessions can also be used as part of payment collections."
// type: object
// required:
//   - amount
//   - cart_id
//   - created_at
//   - data
//   - id
//   - is_initiated
//   - is_selected
//   - idempotency_key
//   - payment_authorized_at
//   - provider_id
//   - status
//   - updated_at
//
// properties:
//
//	id:
//	  description: The payment session's ID
//	  type: string
//	  example: ps_01G901XNSRM2YS3ASN9H5KG3FZ
//	cart_id:
//	  description: The ID of the cart that the payment session was created for.
//	  nullable: true
//	  type: string
//	  example: cart_01G8ZH853Y6TFXWPG5EYE81X63
//	cart:
//	  description: The details of the cart that the payment session was created for.
//	  x-expandable: "cart"
//	  nullable: true
//	  $ref: "#/components/schemas/Cart"
//	provider_id:
//	  description: The ID of the Payment Provider that is responsible for the Payment Session
//	  type: string
//	  example: manual
//	is_selected:
//	  description: A flag to indicate if the Payment Session has been selected as the method that will be used to complete the purchase.
//	  nullable: true
//	  type: boolean
//	  example: true
//	is_initiated:
//	  description: A flag to indicate if a communication with the third party provider has been initiated.
//	  type: boolean
//	  default: false
//	  example: true
//	status:
//	  description: Indicates the status of the Payment Session. Will default to `pending`, and will eventually become `authorized`. Payment Sessions may have the status of `requires_more` to indicate that further actions are to be completed by the Customer.
//	  type: string
//	  enum:
//	    - authorized
//	    - pending
//	    - requires_more
//	    - error
//	    - canceled
//	  example: pending
//	data:
//	  description: The data required for the Payment Provider to identify, modify and process the Payment Session. Typically this will be an object that holds an id to the external payment session, but can be an empty object if the Payment Provider doesn't hold any state.
//	  type: object
//	  example: {}
//	idempotency_key:
//	  description: Randomly generated key used to continue the completion of a cart in case of failure.
//	  nullable: true
//	  type: string
//	  externalDocs:
//	    url: https://docs.medusajs.com/development/idempotency-key/overview.md
//	    description: Learn more how to use the idempotency key.
//	amount:
//	  description: The amount that the Payment Session has been authorized for.
//	  nullable: true
//	  type: integer
//	  example: 100
//	payment_authorized_at:
//	  description: The date with timezone at which the Payment Session was authorized.
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
type PaymentSession struct {
	core.Model

	CartId              uuid.NullUUID        `json:"cart_id" gorm:"column:cart_id"`
	Cart                *Cart                `json:"cart" gorm:"foreignKey:CartId"`
	ProviderId          uuid.NullUUID        `json:"provider_id" gorm:"column:provider_id"`
	IsSelected          bool                 `json:"is_selected" gorm:"column:is_selected"`
	IsInitiated         bool                 `json:"is_initiated" gorm:"column:is_initiated;default:false"`
	Status              PaymentSessionStatus `json:"status" gorm:"column:status"`
	Data                core.JSONB           `json:"data" gorm:"column:data"`
	Amount              float64              `json:"amount" gorm:"column:amount"`
	PaymentAuthorizedAt *time.Time           `json:"payment_authorized_at" gorm:"column:payment_authorized_at"`
	IdempotencyKey      string               `json:"idempotency_key" gorm:"column:idempotency_key"`
}

type PaymentSessionStatus string

const (
	PaymentSessionStatusAuthorized   PaymentSessionStatus = "authorized"
	PaymentSessionStatusPending      PaymentSessionStatus = "pending"
	PaymentSessionStatusRequiresMore PaymentSessionStatus = "requires_more"
	PaymentSessionStatusError        PaymentSessionStatus = "error"
	PaymentSessionStatusCanceled     PaymentSessionStatus = "canceled"
	PaymentSessionStatusSuccess      PaymentSessionStatus = "success"
)

func (pl *PaymentSessionStatus) Scan(value interface{}) error {
	*pl = PaymentSessionStatus(value.([]byte))
	return nil
}

func (pl PaymentSessionStatus) Value() (driver.Value, error) {
	return string(pl), nil
}

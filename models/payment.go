package models

import (
	"database/sql/driver"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:Payment
// title: "Payment"
// description: "A payment is originally created from a payment session. Once a payment session is authorized, the payment is created to represent the authorized amount with a given payment method. Payments can be captured, canceled or refunded. Payments can be made towards orders, swaps, order edits, or other resources."
// type: object
// required:
//   - amount
//   - amount_refunded
//   - canceled_at
//   - captured_at
//   - cart_id
//   - created_at
//   - currency_code
//   - data
//   - id
//   - idempotency_key
//   - metadata
//   - order_id
//   - provider_id
//   - swap_id
//   - updated_at
//
// properties:
//
//	id:
//	  description: The payment's ID
//	  type: string
//	  example: pay_01G2SJNT6DEEWDFNAJ4XWDTHKE
//	swap_id:
//	  description: The ID of the swap that this payment was potentially created for.
//	  nullable: true
//	  type: string
//	  example: null
//	swap:
//	  description: The details of the swap that this payment was potentially created for.
//	  x-expandable: "swap"
//	  nullable: true
//	  $ref: "#/components/schemas/Swap"
//	cart_id:
//	  description: The ID of the cart that the payment session was potentially created for.
//	  nullable: true
//	  type: string
//	cart:
//	  description: The details of the cart that the payment session was potentially created for.
//	  x-expandable: "cart"
//	  nullable: true
//	  $ref: "#/components/schemas/Cart"
//	order_id:
//	  description: The ID of the order that the payment session was potentially created for.
//	  nullable: true
//	  type: string
//	  example: order_01G8TJSYT9M6AVS5N4EMNFS1EK
//	order:
//	  description: The details of the order that the payment session was potentially created for.
//	  x-expandable: "order"
//	  nullable: true
//	  $ref: "#/components/schemas/Order"
//	amount:
//	  description: The amount that the Payment has been authorized for.
//	  type: integer
//	  example: 100
//	currency_code:
//	  description: The 3 character ISO currency code of the payment.
//	  type: string
//	  example: usd
//	  externalDocs:
//	    url: https://en.wikipedia.org/wiki/ISO_4217#Active_codes
//	    description: See a list of codes.
//	currency:
//	  description: The details of the currency of the payment.
//	  x-expandable: "currency"
//	  nullable: true
//	  $ref: "#/components/schemas/Currency"
//	amount_refunded:
//	  description: The amount of the original Payment amount that has been refunded back to the Customer.
//	  type: integer
//	  default: 0
//	  example: 0
//	provider_id:
//	  description: The id of the Payment Provider that is responsible for the Payment
//	  type: string
//	  example: manual
//	data:
//	  description: The data required for the Payment Provider to identify, modify and process the Payment. Typically this will be an object that holds an id to the external payment session, but can be an empty object if the Payment Provider doesn't hold any state.
//	  type: object
//	  example: {}
//	captured_at:
//	  description: The date with timezone at which the Payment was captured.
//	  nullable: true
//	  type: string
//	  format: date-time
//	canceled_at:
//	  description: The date with timezone at which the Payment was canceled.
//	  nullable: true
//	  type: string
//	  format: date-time
//	idempotency_key:
//	  description: Randomly generated key used to continue the completion of a payment in case of failure.
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
type Payment struct {
	core.BaseModel

	SwapId         uuid.NullUUID `json:"swap_id" gorm:"column:swap_id"`
	Swap           *Swap         `json:"swap" gorm:"foreignKey:SwapId"`
	CartId         uuid.NullUUID `json:"cart_id" gorm:"column:cart_id"`
	Cart           *Cart         `json:"cart" gorm:"foreignKey:CartId"`
	OrderId        uuid.NullUUID `json:"order_id" gorm:"column:order_id"`
	Order          *Order        `json:"order" gorm:"foreignKey:OrderId"`
	Amount         float64       `json:"amount" gorm:"column:amount"`
	CurrencyCode   string        `json:"currency_code" gorm:"column:currency_code"`
	Currency       *Currency     `json:"currency" gorm:"foreignKey:CurrencyCode;foreignKey:Code"`
	AmountRefunded float64       `json:"amount_refunded" gorm:"column:amount_refunded;default:0"`
	ProviderId     uuid.NullUUID `json:"provider_id" gorm:"column:provider_id"`
	Data           core.JSONB    `json:"data" gorm:"column:data"`
	CapturedAt     *time.Time    `json:"captured_at" gorm:"column:captured_at"`
	CanceledAt     *time.Time    `json:"canceled_at" gorm:"column:canceled_at"`
	IdempotencyKey string        `json:"idempotency_key" gorm:"column:idempotency_key"`
}

type PaymentStatus string

const (
	PaymentStatusNotPaid           PaymentStatus = "not_paid"
	PaymentStatusAwaiting          PaymentStatus = "awaiting"
	PaymentStatusCaptured          PaymentStatus = "captured"
	PaymentStatusPartiallyRefunded PaymentStatus = "partially_refunded"
	PaymentStatusRefunded          PaymentStatus = "refunded"
	PaymentStatusCanceled          PaymentStatus = "canceled"
	PaymentStatusRequiresAction    PaymentStatus = "requires_action"
)

func (pl *PaymentStatus) Scan(value interface{}) error {
	*pl = PaymentStatus(value.([]byte))
	return nil
}

func (pl PaymentStatus) Value() (driver.Value, error) {
	return string(pl), nil
}

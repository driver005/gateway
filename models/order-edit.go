package models

import (
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:OrderEdit
// title: "Order Edit"
// description: "Order edit allows modifying items in an order, such as adding, updating, or deleting items from the original order. Once the order edit is confirmed, the changes are reflected on the original order."
// type: object
// required:
//   - canceled_at
//   - canceled_by
//   - confirmed_by
//   - confirmed_at
//   - created_at
//   - created_by
//   - declined_at
//   - declined_by
//   - declined_reason
//   - id
//   - internal_note
//   - order_id
//   - payment_collection_id
//   - requested_at
//   - requested_by
//   - status
//   - updated_at
//
// properties:
//
//	id:
//	  description: The order edit's ID
//	  type: string
//	  example: oe_01G8TJSYT9M6AVS5N4EMNFS1EK
//	order_id:
//	  description: The ID of the order that is edited
//	  type: string
//	  example: order_01G2SG30J8C85S4A5CHM2S1NS2
//	order:
//	  description: The details of the order that this order edit was created for.
//	  x-expandable: "order"
//	  nullable: true
//	  $ref: "#/components/schemas/Order"
//	changes:
//	  description: The details of all the changes on the original order's line items.
//	  x-expandable: "changes"
//	  type: array
//	  items:
//	    $ref: "#/components/schemas/OrderItemChange"
//	internal_note:
//	  description: An optional note with additional details about the order edit.
//	  nullable: true
//	  type: string
//	  example: Included two more items B to the order.
//	created_by:
//	  description: The unique identifier of the user or customer who created the order edit.
//	  type: string
//	requested_by:
//	  description: The unique identifier of the user or customer who requested the order edit.
//	  nullable: true
//	  type: string
//	requested_at:
//	  description: The date with timezone at which the edit was requested.
//	  nullable: true
//	  type: string
//	  format: date-time
//	confirmed_by:
//	  description: The unique identifier of the user or customer who confirmed the order edit.
//	  nullable: true
//	  type: string
//	confirmed_at:
//	  description: The date with timezone at which the edit was confirmed.
//	  nullable: true
//	  type: string
//	  format: date-time
//	declined_by:
//	  description: The unique identifier of the user or customer who declined the order edit.
//	  nullable: true
//	  type: string
//	declined_at:
//	  description: The date with timezone at which the edit was declined.
//	  nullable: true
//	  type: string
//	  format: date-time
//	declined_reason:
//	  description: An optional note why  the order edit is declined.
//	  nullable: true
//	  type: string
//	canceled_by:
//	  description: The unique identifier of the user or customer who cancelled the order edit.
//	  nullable: true
//	  type: string
//	canceled_at:
//	  description: The date with timezone at which the edit was cancelled.
//	  nullable: true
//	  type: string
//	  format: date-time
//	subtotal:
//	  description: The total of subtotal
//	  type: integer
//	  example: 8000
//	discount_total:
//	  description: The total of discount
//	  type: integer
//	  example: 800
//	shipping_total:
//	  description: The total of the shipping amount
//	  type: integer
//	  example: 800
//	gift_card_total:
//	  description: The total of the gift card amount
//	  type: integer
//	  example: 800
//	gift_card_tax_total:
//	  description: The total of the gift card tax amount
//	  type: integer
//	  example: 800
//	tax_total:
//	  description: The total of tax
//	  type: integer
//	  example: 0
//	total:
//	  description: The total amount of the edited order.
//	  type: integer
//	  example: 8200
//	difference_due:
//	  description: The difference between the total amount of the order and total amount of edited order.
//	  type: integer
//	  example: 8200
//	status:
//	  description: The status of the order edit.
//	  type: string
//	  enum:
//	    - confirmed
//	    - declined
//	    - requested
//	    - created
//	    - canceled
//	items:
//	  description: The details of the cloned items from the original order with the new changes. Once the order edit is confirmed, these line items are associated with the original order.
//	  type: array
//	  x-expandable: "items"
//	  items:
//	    $ref: "#/components/schemas/LineItem"
//	payment_collection_id:
//	  description: The ID of the payment collection
//	  nullable: true
//	  type: string
//	  example: paycol_01G8TJSYT9M6AVS5N4EMNFS1EK
//	payment_collection:
//	  description: The details of the payment collection used to authorize additional payment if necessary.
//	  x-expandable: "payment_collection"
//	  nullable: true
//	  $ref: "#/components/schemas/PaymentCollection"
//	created_at:
//	  description: The date with timezone at which the resource was created.
//	  type: string
//	  format: date-time
//	updated_at:
//	  description: The date with timezone at which the resource was updated.
//	  type: string
//	  format: date-time
type OrderEdit struct {
	core.Model

	Status              OrderEditStatus    `json:"status" gorm:"column:status"`
	OrderId             uuid.NullUUID      `json:"order_id" gorm:"column:order_id"`
	Order               *Order             `json:"order" gorm:"foreignKey:OrderId"`
	Changes             []OrderItemChange  `json:"changes" gorm:"foreignKey:Id"`
	InternalNote        string             `json:"internal_note" gorm:"column:internal_note"`
	CreatedBy           uuid.NullUUID      `json:"created_by" gorm:"column:created_by"`
	RequestedBy         uuid.NullUUID      `json:"requested_by" gorm:"column:requested_by"`
	RequestedAt         *time.Time         `json:"requested_at" gorm:"column:requested_at"`
	ConfirmedBy         uuid.NullUUID      `json:"confirmed_by" gorm:"column:confirmed_by"`
	ConfirmedAt         *time.Time         `json:"confirmed_at" gorm:"column:confirmed_at"`
	DeclinedBy          uuid.NullUUID      `json:"declined_by" gorm:"column:declined_by"`
	CanceledBy          uuid.NullUUID      `json:"canceled_by" gorm:"column:canceled_by"`
	DeclinedAt          *time.Time         `json:"declined_at" gorm:"column:declined_at"`
	DeclinedReason      string             `json:"declined_reason" gorm:"column:declined_reason"`
	CanceledAt          *time.Time         `json:"canceled_at" gorm:"column:canceled_at"`
	ShippingTotal       float64            `json:"shipping_total" gorm:"column:shipping_total"`
	Subtotal            float64            `json:"subtotal" gorm:"column:subtotal"`
	DiscountTotal       float64            `json:"discount_total" gorm:"column:discount_total"`
	TaxTotal            float64            `json:"tax_total" gorm:"column:tax_total"`
	Total               float64            `json:"total" gorm:"column:total"`
	DifferenceDue       float64            `json:"difference_due" gorm:"column:difference_due"`
	Items               []LineItem         `json:"items" gorm:"foreignKey:Id"`
	PaymentCollectionId uuid.NullUUID      `json:"payment_collection_id" gorm:"column:payment_collection_id"`
	PaymentCollection   *PaymentCollection `json:"region" gorm:"foreignKey:PaymentCollectionId"`
	GiftCardTotal       float64            `json:"gift_card_total" gorm:"column:gift_card_total"`
	GiftCardTaxTotal    float64            `json:"gift_card_tax_total" gorm:"column:gift_card_tax_total"`
}

type OrderEditStatus string

const (
	OrderEditStatusConfirmed OrderEditStatus = "confirmed"
	OrderEditStatusDeclined  OrderEditStatus = "declined"
	OrderEditStatusRequested OrderEditStatus = "requested"
	OrderEditStatusCreated   OrderEditStatus = "created"
	OrderEditStatusCanceled  OrderEditStatus = "canceled"
)

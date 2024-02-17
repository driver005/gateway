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

	Status              OrderEditStatus    `json:"status"`
	OrderId             uuid.NullUUID      `json:"order_id"`
	Order               *Order             `json:"order" gorm:"foreignKey:id;references:order_id"`
	Changes             []OrderItemChange  `json:"changes" gorm:"foreignKey:id"`
	InternalNote        string             `json:"internal_note" gorm:"default:null"`
	CreatedBy           uuid.UUID          `json:"created_by"`
	RequestedBy         uuid.UUID          `json:"requested_by" gorm:"default:null"`
	RequestedAt         *time.Time         `json:"requested_at" gorm:"default:null"`
	ConfirmedBy         uuid.UUID          `json:"confirmed_by" gorm:"default:null"`
	ConfirmedAt         *time.Time         `json:"confirmed_at" gorm:"default:null"`
	DeclinedBy          uuid.UUID          `json:"declined_by" gorm:"default:null"`
	CanceledBy          uuid.UUID          `json:"canceled_by" gorm:"default:null"`
	DeclinedAt          *time.Time         `json:"declined_at" gorm:"default:null"`
	DeclinedReason      string             `json:"declined_reason" gorm:"default:null"`
	CanceledAt          *time.Time         `json:"canceled_at" gorm:"default:null"`
	ShippingTotal       float64            `json:"shipping_total" gorm:"default:null"`
	Subtotal            float64            `json:"subtotal" gorm:"default:null"`
	DiscountTotal       float64            `json:"discount_total" gorm:"default:null"`
	TaxTotal            float64            `json:"tax_total" gorm:"default:null"`
	Total               float64            `json:"total" gorm:"default:null"`
	DifferenceDue       float64            `json:"difference_due" gorm:"default:null"`
	Items               []LineItem         `json:"items" gorm:"foreignKey:id"`
	PaymentCollectionId uuid.NullUUID      `json:"payment_collection_id"`
	PaymentCollection   *PaymentCollection `json:"region,omitempty" gorm:"foreignKey:id;references:payment_collection_id"`
	GiftCardTotal       float64            `json:"gift_card_total" gorm:"default:null"`
	GiftCardTaxTotal    float64            `json:"gift_card_tax_total" gorm:"default:null"`
}

type OrderEditStatus string

const (
	OrderEditStatusConfirmed OrderEditStatus = "confirmed"
	OrderEditStatusDeclined  OrderEditStatus = "declined"
	OrderEditStatusRequested OrderEditStatus = "requested"
	OrderEditStatusCreated   OrderEditStatus = "created"
	OrderEditStatusCanceled  OrderEditStatus = "canceled"
)

package models

import (
	"database/sql/driver"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:Order
// title: "Order"
// description: "An order is a purchase made by a customer. It holds details about payment and fulfillment of the order. An order may also be created from a draft order, which is created by an admin user."
// type: object
// required:
//   - billing_address_id
//   - canceled_at
//   - cart_id
//   - created_at
//   - currency_code
//   - customer_id
//   - draft_order_id
//   - display_id
//   - email
//   - external_id
//   - fulfillment_status
//   - id
//   - idempotency_key
//   - metadata
//   - no_notification
//   - object
//   - payment_status
//   - region_id
//   - shipping_address_id
//   - status
//   - tax_rate
//   - updated_at
//
// properties:
//
//	id:
//	  description: The order's ID
//	  type: string
//	  example: order_01G8TJSYT9M6AVS5N4EMNFS1EK
//	status:
//	  description: The order's status
//	  type: string
//	  enum:
//	    - pending
//	    - completed
//	    - archived
//	    - canceled
//	    - requires_action
//	  default: pending
//	fulfillment_status:
//	  description: The order's fulfillment status
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
//	payment_status:
//	  description: The order's payment status
//	  type: string
//	  enum:
//	    - not_paid
//	    - awaiting
//	    - captured
//	    - partially_refunded
//	    - refunded
//	    - canceled
//	    - requires_action
//	  default: not_paid
//	display_id:
//	  description: The order's display ID
//	  type: integer
//	  example: 2
//	cart_id:
//	  description: The ID of the cart associated with the order
//	  nullable: true
//	  type: string
//	  example: cart_01G8ZH853Y6TFXWPG5EYE81X63
//	cart:
//	  description: The details of the cart associated with the order.
//	  x-expandable: "cart"
//	  nullable: true
//	  $ref: "#/components/schemas/Cart"
//	customer_id:
//	  description: The ID of the customer associated with the order
//	  type: string
//	  example: cus_01G2SG30J8C85S4A5CHM2S1NS2
//	customer:
//	  description: The details of the customer associated with the order.
//	  x-expandable: "customer"
//	  nullable: true
//	  $ref: "#/components/schemas/Customer"
//	email:
//	  description: The email associated with the order
//	  type: string
//	  format: email
//	billing_address_id:
//	  description: The ID of the billing address associated with the order
//	  nullable: true
//	  type: string
//	  example: addr_01G8ZH853YPY9B94857DY91YGW
//	billing_address:
//	  description: The details of the billing address associated with the order.
//	  x-expandable: "billing_address"
//	  nullable: true
//	  $ref: "#/components/schemas/Address"
//	shipping_address_id:
//	  description: The ID of the shipping address associated with the order
//	  nullable: true
//	  type: string
//	  example: addr_01G8ZH853YPY9B94857DY91YGW
//	shipping_address:
//	  description: The details of the shipping address associated with the order.
//	  x-expandable: "shipping_address"
//	  nullable: true
//	  $ref: "#/components/schemas/Address"
//	region_id:
//	  description: The ID of the region this order was created in.
//	  type: string
//	  example: reg_01G1G5V26T9H8Y0M4JNE3YGA4G
//	region:
//	  description: The details of the region this order was created in.
//	  x-expandable: "region"
//	  nullable: true
//	  $ref: "#/components/schemas/Region"
//	currency_code:
//	  description: The 3 character currency code that is used in the order
//	  type: string
//	  example: usd
//	  externalDocs:
//	    url: https://en.wikipedia.org/wiki/ISO_4217#Active_codes
//	    description: See a list of codes.
//	currency:
//	  description: The details of the currency used in the order.
//	  x-expandable: "currency"
//	  nullable: true
//	  $ref: "#/components/schemas/Currency"
//	tax_rate:
//	  description: The order's tax rate
//	  nullable: true
//	  type: number
//	  example: 0
//	discounts:
//	  description: The details of the discounts applied on the order.
//	  type: array
//	  x-expandable: "discounts"
//	  items:
//	    $ref: "#/components/schemas/Discount"
//	gift_cards:
//	  description: The details of the gift card used in the order.
//	  type: array
//	  x-expandable: "gift_cards"
//	  items:
//	    $ref: "#/components/schemas/GiftCard"
//	shipping_methods:
//	  description: The details of the shipping methods used in the order.
//	  type: array
//	  x-expandable: "shipping_methods"
//	  items:
//	    $ref: "#/components/schemas/ShippingMethod"
//	payments:
//	  description: The details of the payments used in the order.
//	  type: array
//	  x-expandable: "payments"
//	  items:
//	    $ref: "#/components/schemas/Payment"
//	fulfillments:
//	  description: The details of the fulfillments created for the order.
//	  type: array
//	  x-expandable: "fulfillments"
//	  items:
//	    $ref: "#/components/schemas/Fulfillment"
//	returns:
//	  description: The details of the returns created for the order.
//	  type: array
//	  x-expandable: "returns"
//	  items:
//	    $ref: "#/components/schemas/Return"
//	claims:
//	  description: The details of the claims created for the order.
//	  type: array
//	  x-expandable: "claims"
//	  items:
//	    $ref: "#/components/schemas/ClaimOrder"
//	refunds:
//	  description: The details of the refunds created for the order.
//	  type: array
//	  x-expandable: "refunds"
//	  items:
//	    $ref: "#/components/schemas/Refund"
//	swaps:
//	  description: The details of the swaps created for the order.
//	  type: array
//	  x-expandable: "swaps"
//	  items:
//	    $ref: "#/components/schemas/Swap"
//	draft_order_id:
//	  description: The ID of the draft order this order was created from.
//	  nullable: true
//	  type: string
//	  example: null
//	draft_order:
//	  description: The details of the draft order this order was created from.
//	  x-expandable: "draft_order"
//	  nullable: true
//	  $ref: "#/components/schemas/DraftOrder"
//	items:
//	  description: The details of the line items that belong to the order.
//	  x-expandable: "items"
//	  type: array
//	  items:
//	    $ref: "#/components/schemas/LineItem"
//	edits:
//	  description: The details of the order edits done on the order.
//	  type: array
//	  x-expandable: "edits"
//	  items:
//	    $ref: "#/components/schemas/OrderEdit"
//	gift_card_transactions:
//	  description: The gift card transactions made in the order.
//	  type: array
//	  x-expandable: "gift_card_transactions"
//	  items:
//	    $ref: "#/components/schemas/GiftCardTransaction"
//	canceled_at:
//	  description: The date the order was canceled on.
//	  nullable: true
//	  type: string
//	  format: date-time
//	no_notification:
//	  description: Flag for describing whether or not notifications related to this should be send.
//	  nullable: true
//	  type: boolean
//	  example: false
//	idempotency_key:
//	  description: Randomly generated key used to continue the processing of the order in case of failure.
//	  nullable: true
//	  type: string
//	  externalDocs:
//	    url: https://docs.medusajs.com/development/idempotency-key/overview.md
//	    description: Learn more how to use the idempotency key.
//	external_id:
//	  description: The ID of an external order.
//	  nullable: true
//	  type: string
//	  example: null
//	sales_channel_id:
//	  description: The ID of the sales channel this order belongs to.
//	  nullable: true
//	  type: string
//	  example: null
//	sales_channel:
//	  description: The details of the sales channel this order belongs to.
//	  x-expandable: "sales_channel"
//	  nullable: true
//	  $ref: "#/components/schemas/SalesChannel"
//	shipping_total:
//	  type: integer
//	  description: The total of shipping
//	  example: 1000
//	  nullable: true
//	shipping_tax_total:
//	  type: integer
//	  description: The tax total applied on shipping
//	  example: 1000
//	raw_discount_total:
//	  description: The total of discount
//	  type: integer
//	  example: 800
//	discount_total:
//	  description: The total of discount rounded
//	  type: integer
//	  example: 800
//	tax_total:
//	  description: The total of tax
//	  type: integer
//	  example: 0
//	item_tax_total:
//	  description: The tax total applied on items
//	  type: integer
//	  example: 0
//	  nullable: true
//	refunded_total:
//	  description: The total amount refunded if the order is returned.
//	  type: integer
//	  example: 0
//	total:
//	  description: The total amount of the order
//	  type: integer
//	  example: 8200
//	subtotal:
//	  description: The subtotal of the order
//	  type: integer
//	  example: 8000
//	paid_total:
//	  description: The total amount paid
//	  type: integer
//	  example: 8000
//	refundable_amount:
//	  description: The amount that can be refunded
//	  type: integer
//	  example: 8200
//	gift_card_total:
//	  description: The total of gift cards
//	  type: integer
//	  example: 0
//	gift_card_tax_total:
//	  description: The total of gift cards with taxes
//	  type: integer
//	  example: 0
//	returnable_items:
//	  description: The details of the line items that are returnable as part of the order, swaps, or claims
//	  type: array
//	  x-expandable: "returnable_items"
//	  items:
//	    $ref: "#/components/schemas/LineItem"
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
//	sales_channels:
//	  description: The associated sales channels.
//	  type: array
//	  nullable: true
//	  x-expandable: "sales_channels"
//	  x-featureFlag: "medusa_v2"
//	  items:
//	    $ref: "#/components/schemas/SalesChannel"
type Order struct {
	core.Model

	Status               OrderStatus           `json:"status" gorm:"column:status;default:'pending'"`
	FulfillmentStatus    FulfillmentStatus     `json:"fulfillment_status" gorm:"column:fulfillment_status;default:'not_fulfilled'"`
	PaymentStatus        PaymentStatus         `json:"payment_status" gorm:"column:payment_status;default:'not_paid'"`
	DisplayId            int                   `json:"display_id" gorm:"column:display_id"`
	CartId               uuid.NullUUID         `json:"cart_id" gorm:"column:cart_id"`
	Cart                 *Cart                 `json:"cart" gorm:"foreignKey:CartId"`
	CustomerId           uuid.NullUUID         `json:"customer_id" gorm:"column:customer_id"`
	Customer             *Customer             `json:"customer" gorm:"foreignKey:CustomerId"`
	Email                string                `json:"email" gorm:"column:email"`
	BillingAddressId     uuid.NullUUID         `json:"billing_address_id" gorm:"column:billing_address_id"`
	BillingAddress       *Address              `json:"billing_address" gorm:"foreignKey:BillingAddressId"`
	ShippingAddressId    uuid.NullUUID         `json:"shipping_address_id" gorm:"column:shipping_address_id"`
	ShippingAddress      *Address              `json:"shipping_address" gorm:"foreignKey:ShippingAddressId"`
	RegionId             uuid.NullUUID         `json:"region_id" gorm:"column:region_id"`
	Region               *Region               `json:"region" gorm:"foreignKey:RegionId"`
	CurrencyCode         string                `json:"currency_code" gorm:"column:currency_code"`
	Currency             *Currency             `json:"currency" gorm:"foreignKey:CurrencyCode;foreignKey:Code"`
	TaxRate              float64               `json:"tax_rate" gorm:"column:tax_rate"`
	Discounts            []Discount            `json:"discounts" gorm:"many2many:order_discounts"`
	GiftCards            []GiftCard            `json:"gift_cards" gorm:"many2many:order_gift_cards"`
	ShippingMethods      []ShippingMethod      `json:"shipping_methods" gorm:"foreignKey:Id"`
	Payments             []Payment             `json:"payments" gorm:"foreignKey:Id"`
	Fulfillments         []Fulfillment         `json:"fulfillments" gorm:"foreignKey:Id"`
	Returns              []Return              `json:"returns" gorm:"foreignKey:Id"`
	Claims               []ClaimOrder          `json:"claims" gorm:"foreignKey:Id"`
	Refunds              []Refund              `json:"refunds" gorm:"foreignKey:Id"`
	Swaps                []Swap                `json:"swaps" gorm:"foreignKey:Id"`
	DraftOrderId         uuid.NullUUID         `json:"draft_order_id" gorm:"column:draft_order_id"`
	DraftOrder           *DraftOrder           `json:"draft_order" gorm:"foreignKey:DraftOrderId"`
	Items                []LineItem            `json:"items" gorm:"foreignKey:Id"`
	ReturnableItems      []LineItem            `json:"returnable_items" gorm:"foreignKey:Id"`
	Edits                []OrderEdit           `json:"edits" gorm:"foreignKey:Id"`
	GiftCardTransactions []GiftCardTransaction `json:"gift_card_transactions" gorm:"foreignKey:Id"`
	CanceledAt           *time.Time            `json:"canceled_at" gorm:"column:canceled_at"`
	NoNotification       bool                  `json:"no_notification" gorm:"column:no_notification"`
	IdempotencyKey       string                `json:"idempotency_key" gorm:"column:idempotency_key"`
	ExternalId           uuid.NullUUID         `json:"external_id" gorm:"column:external_id"`
	SalesChannelId       uuid.NullUUID         `json:"sales_channel_id" gorm:"column:sales_channel_id"`
	SalesChannel         *SalesChannel         `json:"sales_channel" gorm:"foreignKey:SalesChannelId"`
	ShippingTotal        float64               `json:"shipping_total" gorm:"column:shipping_total"`
	ShippingTaxTotal     float64               `json:"shipping_tax_total" gorm:"column:shipping_tax_total"`
	DiscountTotal        float64               `json:"discount_total" gorm:"column:discount_total"`
	RawDiscountTotal     float64               `json:"raw_discount_total" gorm:"column:raw_discount_total"`
	ItemTaxTotal         float64               `json:"item_tax_total" gorm:"column:item_tax_total"`
	TaxTotal             float64               `json:"tax_total" gorm:"column:tax_total"`
	RefundedTotal        float64               `json:"refunded_total" gorm:"column:refunded_total"`
	Total                float64               `json:"total" gorm:"column:total"`
	Subtotal             float64               `json:"subtotal" gorm:"column:subtotal"`
	PaidTotal            float64               `json:"paid_total" gorm:"column:paid_total"`
	RefundableAmount     float64               `json:"refundable_amount" gorm:"column:refundable_amount"`
	GiftCardTotal        float64               `json:"gift_card_total" gorm:"column:gift_card_total"`
	GiftCardTaxTotal     float64               `json:"gift_card_tax_total" gorm:"column:gift_card_tax_total"`
}

type OrderStatus string

const (
	OrderStatusPending        OrderStatus = "pending"
	OrderStatusCompleted      OrderStatus = "completed"
	OrderStatusArchived       OrderStatus = "archived"
	OrderStatusCanceled       OrderStatus = "canceled"
	OrderStatusRefunded       OrderStatus = "refunded"
	OrderStatusRequiresAction OrderStatus = "requires_action"
)

func (pl *OrderStatus) Scan(value interface{}) error {
	*pl = OrderStatus(value.([]byte))
	return nil
}

func (pl OrderStatus) Value() (driver.Value, error) {
	return string(pl), nil
}

package models

import (
	"database/sql/driver"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:Cart
// title: "Cart"
// description: "A cart represents a virtual shopping bag. It can be used to complete an order, a swap, or a claim."
// type: object
// required:
//   - billing_address_id
//   - completed_at
//   - context
//   - created_at
//   - customer_id
//   - deleted_at
//   - email
//   - id
//   - idempotency_key
//   - metadata
//   - payment_authorized_at
//   - payment_id
//   - payment_session
//   - region_id
//   - shipping_address_id
//   - type
//   - updated_at
//
// properties:
//
//	id:
//	  description: The cart's ID
//	  type: string
//	  example: cart_01G8ZH853Y6TFXWPG5EYE81X63
//	email:
//	  description: The email associated with the cart
//	  nullable: true
//	  type: string
//	  format: email
//	billing_address_id:
//	  description: The billing address's ID
//	  nullable: true
//	  type: string
//	  example: addr_01G8ZH853YPY9B94857DY91YGW
//	billing_address:
//	  description: The details of the billing address associated with the cart.
//	  x-expandable: "billing_address"
//	  nullable: true
//	  $ref: "#/components/schemas/Address"
//	shipping_address_id:
//	  description: The shipping address's ID
//	  nullable: true
//	  type: string
//	  example: addr_01G8ZH853YPY9B94857DY91YGW
//	shipping_address:
//	  description: The details of the shipping address associated with the cart.
//	  x-expandable: "shipping_address"
//	  nullable: true
//	  $ref: "#/components/schemas/Address"
//	items:
//	  description: The line items added to the cart.
//	  type: array
//	  x-expandable: "items"
//	  items:
//	    $ref: "#/components/schemas/LineItem"
//	region_id:
//	  description: The region's ID
//	  type: string
//	  example: reg_01G1G5V26T9H8Y0M4JNE3YGA4G
//	region:
//	  description: The details of the region associated with the cart.
//	  x-expandable: "region"
//	  nullable: true
//	  $ref: "#/components/schemas/Region"
//	discounts:
//	  description: An array of details of all discounts applied to the cart.
//	  type: array
//	  x-expandable: "discounts"
//	  items:
//	    $ref: "#/components/schemas/Discount"
//	gift_cards:
//	  description: An array of details of all gift cards applied to the cart.
//	  type: array
//	  x-expandable: "gift_cards"
//	  items:
//	    $ref: "#/components/schemas/GiftCard"
//	customer_id:
//	  description: The customer's ID
//	  nullable: true
//	  type: string
//	  example: cus_01G2SG30J8C85S4A5CHM2S1NS2
//	customer:
//	  description: The details of the customer the cart belongs to.
//	  x-expandable: "customer"
//	  nullable: true
//	  $ref: "#/components/schemas/Customer"
//	payment_session:
//	  description: The details of the selected payment session in the cart.
//	  x-expandable: "payment_session"
//	  nullable: true
//	  $ref: "#/components/schemas/PaymentSession"
//	payment_sessions:
//	  description: The details of all payment sessions created on the cart.
//	  type: array
//	  x-expandable: "payment_sessions"
//	  items:
//	    $ref: "#/components/schemas/PaymentSession"
//	payment_id:
//	  description: The payment's ID if available
//	  nullable: true
//	  type: string
//	  example: pay_01G8ZCC5W42ZNY842124G7P5R9
//	payment:
//	  description: The details of the payment associated with the cart.
//	  nullable: true
//	  x-expandable: "payment"
//	  $ref: "#/components/schemas/Payment"
//	shipping_methods:
//	  description: The details of the shipping methods added to the cart.
//	  type: array
//	  x-expandable: "shipping_methods"
//	  items:
//	    $ref: "#/components/schemas/ShippingMethod"
//	type:
//	  description: The cart's type.
//	  type: string
//	  enum:
//	    - default
//	    - swap
//	    - draft_order
//	    - payment_link
//	    - claim
//	  default: default
//	completed_at:
//	  description: The date with timezone at which the cart was completed.
//	  nullable: true
//	  type: string
//	  format: date-time
//	payment_authorized_at:
//	  description: The date with timezone at which the payment was authorized.
//	  nullable: true
//	  type: string
//	  format: date-time
//	idempotency_key:
//	  description: Randomly generated key used to continue the completion of a cart in case of failure.
//	  nullable: true
//	  type: string
//	  externalDocs:
//	    url: https://docs.medusajs.com/development/idempotency-key/overview.md
//	    description: Learn more how to use the idempotency key.
//	context:
//	  description: "The context of the cart which can include info like IP or user agent."
//	  nullable: true
//	  type: object
//	  example:
//	    ip: "::1"
//	    user_agent: "PostmanRuntime/7.29.2"
//	sales_channel_id:
//	  description: The sales channel ID the cart is associated with.
//	  nullable: true
//	  type: string
//	  example: null
//	sales_channel:
//	  description: The details of the sales channel associated with the cart.
//	  nullable: true
//	  x-expandable: "sales_channel"
//	  $ref: "#/components/schemas/SalesChannel"
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
//	shipping_total:
//	  description: The total of shipping
//	  type: integer
//	  example: 1000
//	discount_total:
//	  description: The total of discount rounded
//	  type: integer
//	  example: 800
//	raw_discount_total:
//	  description: The total of discount
//	  type: integer
//	  example: 800
//	item_tax_total:
//	  description: The total of items with taxes
//	  type: integer
//	  example: 8000
//	shipping_tax_total:
//	  description: The total of shipping with taxes
//	  type: integer
//	  example: 1000
//	tax_total:
//	  description: The total of tax
//	  type: integer
//	  example: 0
//	refunded_total:
//	  description: The total amount refunded if the order associated with this cart is returned.
//	  type: integer
//	  example: 0
//	total:
//	  description: The total amount of the cart
//	  type: integer
//	  example: 8200
//	subtotal:
//	  description: The subtotal of the cart
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
//	sales_channels:
//	  description: The associated sales channels.
//	  type: array
//	  nullable: true
//	  x-expandable: "sales_channels"
//	  items:
//	    $ref: "#/components/schemas/SalesChannel"
type Cart struct {
	core.SoftDeletableModel

	Email               string           `json:"email" gorm:"column:email"`
	BillingAddressId    uuid.NullUUID    `json:"billing_address_id" gorm:"column:billing_address_id"`
	BillingAddress      *Address         `json:"billing_address" gorm:"foreignKey:BillingAddressId"`
	ShippingAddressId   uuid.NullUUID    `json:"shipping_address_id" gorm:"column:shipping_address_id"`
	ShippingAddress     *Address         `json:"shipping_address" gorm:"foreignKey:ShippingAddressId"`
	Items               []LineItem       `json:"items" gorm:"foreignKey:Id"`
	RegionId            uuid.NullUUID    `json:"region_id" gorm:"column:region_id"`
	Region              *Region          `json:"region" gorm:"foreignKey:RegionId"`
	Discounts           []Discount       `json:"discounts" gorm:"many2many:cart_discounts"`
	GiftCards           []GiftCard       `json:"gift_cards" gorm:"many2many:cart_gift_cards"`
	CustomerId          uuid.NullUUID    `json:"customer_id" gorm:"column:customer_id"`
	Customer            *Customer        `json:"customer" gorm:"foreignKey:CustomerId"`
	PaymentSession      *PaymentSession  `json:"payment_session" gorm:"foreignKey:Id"`
	PaymentSessions     []PaymentSession `json:"payment_sessions" gorm:"foreignKey:Id"`
	PaymentId           uuid.NullUUID    `json:"payment_id" gorm:"column:payment_id"`
	Payment             *Payment         `json:"payment" gorm:"foreignKey:PaymentId"`
	ShippingMethods     []ShippingMethod `json:"shipping_methods" gorm:"foreignKey:Id"`
	Type                CartType         `json:"type" gorm:"column:type;default:'default'"`
	CompletedAt         *time.Time       `json:"completed_at" gorm:"column:completed_at"`
	PaymentAuthorizedAt *time.Time       `json:"payment_authorized_at" gorm:"column:payment_authorized_at"`
	IdempotencyKey      string           `json:"idempotency_key" gorm:"column:idempotency_key"`
	Context             core.JSONB       `json:"context" gorm:"column:context"`
	SalesChannelId      uuid.NullUUID    `json:"sales_channel_id" gorm:"column:sales_channel_id"`
	SalesChannel        *SalesChannel    `json:"sales_channel" gorm:"foreignKey:SalesChannelId"`
	ShippingTotal       float64          `json:"shipping_total" gorm:"column:shipping_total"`
	ShippingTaxTotal    float64          `json:"shipping_tax_total" gorm:"column:shipping_tax_total"`
	DiscountTotal       float64          `json:"discount_total" gorm:"column:discount_total"`
	RawDiscountTotal    float64          `json:"raw_discount_total" gorm:"column:raw_discount_total"`
	ItemTaxTotal        float64          `json:"item_tax_total" gorm:"column:item_tax_total"`
	TaxTotal            float64          `json:"tax_total" gorm:"column:tax_total"`
	RefundedTotal       float64          `json:"refunded_total" gorm:"column:refunded_total"`
	Total               float64          `json:"total" gorm:"column:total"`
	Subtotal            float64          `json:"subtotal" gorm:"column:subtotal"`
	RefundableAmount    float64          `json:"refundable_amount" gorm:"column:refundable_amount"`
	GiftCardTotal       float64          `json:"gift_card_total" gorm:"column:gift_card_total"`
	GiftCardTaxTotal    float64          `json:"gift_card_tax_total" gorm:"column:gift_card_tax_total"`
}

// The status of the Price List
type CartType string

// Defines values for CartType.
const (
	CartDefault     = "default"
	CartSwap        = "swap"
	CartDraftOrder  = "draft_order"
	CartPaymentLink = "payment_link"
	CartClaim       = "claim"
)

func (pl *CartType) Scan(value interface{}) error {
	*pl = CartType(value.([]byte))
	return nil
}

func (pl CartType) Value() (driver.Value, error) {
	return string(pl), nil
}

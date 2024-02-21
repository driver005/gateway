package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:Store
// title: "Store"
// description: "A store holds the main settings of the commerce shop. By default, only one store is created and used within the Medusa backend. It holds settings related to the name of the store, available currencies, and more."
// type: object
// required:
//   - created_at
//   - default_currency_code
//   - default_location_id
//   - id
//   - invite_link_template
//   - metadata
//   - name
//   - payment_link_template
//   - swap_link_template
//   - updated_at
//
// properties:
//
//	id:
//	  description: The store's ID
//	  type: string
//	  example: store_01G1G5V21KADXNGH29BJMAJ4B4
//	name:
//	  description: The name of the Store - this may be displayed to the Customer.
//	  type: string
//	  example: Medusa Store
//	  default: Medusa Store
//	default_currency_code:
//	  description: The three character currency code that is the default of the store.
//	  type: string
//	  example: usd
//	  externalDocs:
//	    url: https://en.wikipedia.org/wiki/ISO_4217#Active_codes
//	    description: See a list of codes.
//	default_currency:
//	  description: The details of the store's default currency.
//	  x-expandable: "default_currency"
//	  default: "usd"
//	  nullable: true
//	  $ref: "#/components/schemas/Currency"
//	currencies:
//	  description: The details of the enabled currencies in the store.
//	  type: array
//	  x-expandable: "currencies"
//	  items:
//	    $ref: "#/components/schemas/Currency"
//	swap_link_template:
//	  description: A template to generate Swap links from. Use {{cart_id}} to include the Swap's `cart_id` in the link.
//	  nullable: true
//	  type: string
//	  example: null
//	payment_link_template:
//	  description: A template to generate Payment links from. Use {{cart_id}} to include the payment's `cart_id` in the link.
//	  nullable: true
//	  type: string
//	  example: null
//	invite_link_template:
//	  description: A template to generate Invite links from
//	  nullable: true
//	  type: string
//	  example: null
//	default_location_id:
//	  description: The location ID the store is associated with.
//	  nullable: true
//	  type: string
//	  example: null
//	default_sales_channel_id:
//	  description: The ID of the store's default sales channel.
//	  nullable: true
//	  type: string
//	  example: null
//	default_sales_channel:
//	  description: The details of the store's default sales channel.
//	  x-expandable: "default_sales_channel"
//	  nullable: true
//	  $ref: "#/components/schemas/SalesChannel"
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
type Store struct {
	core.Model

	Name                  string        `json:"name" gorm:"column:name;default:'Store'"`
	DefaultCurrencyCode   string        `json:"default_currency_code" gorm:"column:default_currency_code;default:'usd'"`
	DefaultCurrency       *Currency     `json:"default_currency" gorm:"foreignKey:DefaultCurrencyCode;foreignKey:Code"`
	Currencies            []Currency    `json:"currencies" gorm:"many2many:store_currencies"`
	SwapLinkTemplate      string        `json:"swap_link_template" gorm:"column:swap_link_template"`
	PaymentLinkTemplate   string        `json:"payment_link_template" gorm:"column:payment_link_template"`
	InviteLinkTemplate    string        `json:"invite_link_template" gorm:"column:invite_link_template"`
	DefaultSalesChannelId uuid.NullUUID `json:"default_sales_channel_id" gorm:"column:default_sales_channel_id"`
	DefaultSalesChannel   *SalesChannel `json:"default_sales_channel" gorm:"foreignKey:DefaultSalesChannelId"`
}

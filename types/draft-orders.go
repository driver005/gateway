package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FilterableDraftOrder struct {
	core.FilterModel
}

type DraftOrderListSelector struct {
	Q string `json:"q,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostDraftOrdersReq
// type: object
// description: "The details of the draft order to create."
// required:
//   - email
//   - region_id
//   - shipping_methods
//
// properties:
//
//	status:
//	  description: >-
//	    The status of the draft order. The draft order's default status is `open`. It's changed to `completed` when its payment is marked as paid.
//	  type: string
//	  enum: [open, completed]
//	email:
//	  description: "The email of the customer of the draft order"
//	  type: string
//	  format: email
//	billing_address:
//	  description: "The Address to be used for billing purposes."
//	  anyOf:
//	    - $ref: "#/components/schemas/AddressPayload"
//	    - type: string
//	shipping_address:
//	  description: "The Address to be used for shipping purposes."
//	  anyOf:
//	    - $ref: "#/components/schemas/AddressPayload"
//	    - type: string
//	items:
//	  description: The draft order's line items.
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - quantity
//	    properties:
//	      variant_id:
//	        description: The ID of the Product Variant associated with the line item. If the line item is custom, the `variant_id` should be omitted.
//	        type: string
//	      unit_price:
//	        description: The custom price of the line item. If a `variant_id` is supplied, the price provided here will override the variant's price.
//	        type: integer
//	      title:
//	        description: The title of the line item if `variant_id` is not provided.
//	        type: string
//	      quantity:
//	        description: The quantity of the line item.
//	        type: integer
//	      metadata:
//	        description: The optional key-value map with additional details about the line item.
//	        type: object
//	        externalDocs:
//	          description: "Learn about the metadata attribute, and how to delete and update it."
//	          url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
//	region_id:
//	  description: The ID of the region for the draft order
//	  type: string
//	discounts:
//	  description: The discounts to add to the draft order
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - code
//	    properties:
//	      code:
//	        description: The code of the discount to apply
//	        type: string
//	customer_id:
//	  description: The ID of the customer this draft order is associated with.
//	  type: string
//	no_notification_order:
//	  description: An optional flag passed to the resulting order that indicates whether the customer should receive notifications about order updates.
//	  type: boolean
//	shipping_methods:
//	  description: The shipping methods for the draft order
//	  type: array
//	  items:
//	    type: object
//	    required:
//	       - option_id
//	    properties:
//	      option_id:
//	        description: The ID of the shipping option in use
//	        type: string
//	      data:
//	        description: The optional additional data needed for the shipping method
//	        type: object
//	      price:
//	        description: The price of the shipping method.
//	        type: integer
//	metadata:
//	  description: The optional key-value map with additional details about the Draft Order.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type DraftOrderCreate struct {
	Status              string           `json:"status,omitempty" validate:"omitempty"`
	Email               string           `json:"email"`
	BillingAddressId    uuid.UUID        `json:"billing_address_id,omitempty" validate:"omitempty"`
	BillingAddress      *AddressPayload  `json:"billing_address,omitempty" validate:"omitempty"`
	ShippingAddressId   uuid.UUID        `json:"shipping_address_id,omitempty" validate:"omitempty"`
	ShippingAddress     *AddressPayload  `json:"shipping_address,omitempty" validate:"omitempty"`
	Items               []Item           `json:"items,omitempty" validate:"omitempty"`
	RegionId            uuid.UUID        `json:"region_id"`
	Discounts           []Discount       `json:"discounts,omitempty" validate:"omitempty"`
	CustomerId          uuid.UUID        `json:"customer_id,omitempty" validate:"omitempty"`
	NoNotificationOrder bool             `json:"no_notification_order,omitempty" validate:"omitempty"`
	ShippingMethods     []ShippingMethod `json:"shipping_methods"`
	Metadata            core.JSONB       `json:"metadata,omitempty" validate:"omitempty"`
	IdempotencyKey      string           `json:"idempotency_key,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostDraftOrdersDraftOrderReq
// type: object
// description: "The details of the draft order to update."
// properties:
//
//	region_id:
//	  type: string
//	  description: The ID of the Region to create the Draft Order in.
//	country_code:
//	  type: string
//	  description: "The 2 character ISO code for the Country."
//	  externalDocs:
//	     url: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements
//	     description: See a list of codes.
//	email:
//	  type: string
//	  description: "An email to be used in the Draft Order."
//	  format: email
//	billing_address:
//	  description: "The Address to be used for billing purposes."
//	  anyOf:
//	    - $ref: "#/components/schemas/AddressPayload"
//	    - type: string
//	shipping_address:
//	  description: "The Address to be used for shipping purposes."
//	  anyOf:
//	    - $ref: "#/components/schemas/AddressPayload"
//	    - type: string
//	discounts:
//	  description: "An array of Discount codes to add to the Draft Order."
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - code
//	    properties:
//	      code:
//	        description: "The code that a Discount is identifed by."
//	        type: string
//	no_notification_order:
//	  description: "An optional flag passed to the resulting order that indicates whether the customer should receive notifications about order updates."
//	  type: boolean
//	customer_id:
//	  description: "The ID of the customer this draft order is associated with."
//	  type: string
type DraftOrderUpdate struct {
	RegionId            uuid.UUID       `json:"region_id,omitempty" validate:"omitempty,string"`
	CountryCode         string          `json:"country_code,omitempty" validate:"omitempty,string"`
	Email               string          `json:"email,omitempty" validate:"omitempty,email"`
	BillingAddressId    uuid.UUID       `json:"billing_address_id,omitempty" validate:"omitempty"`
	BillingAddress      *AddressPayload `json:"billing_address,omitempty" validate:"omitempty"`
	ShippingAddressId   uuid.UUID       `json:"shipping_address_id,omitempty" validate:"omitempty"`
	ShippingAddress     *AddressPayload `json:"shipping_address,omitempty" validate:"omitempty"`
	Discounts           []Discount      `json:"discounts,omitempty" validate:"omitempty,dive"`
	CustomerId          uuid.UUID       `json:"customer_id,omitempty" validate:"omitempty,string"`
	NoNotificationOrder bool            `json:"no_notification_order,omitempty" validate:"omitempty,boolean"`
}

type ShippingMethod struct {
	OptionId uuid.UUID  `json:"option_id"`
	Data     core.JSONB `json:"data,omitempty" validate:"omitempty"`
	Price    float64    `json:"price,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostDraftOrdersDraftOrderLineItemsReq
// type: object
// description: "The details of the line item to create."
// required:
//   - quantity
//
// properties:
//
//	variant_id:
//	  description: The ID of the Product Variant associated with the line item. If the line item is custom, the `variant_id` should be omitted.
//	  type: string
//	unit_price:
//	  description: The custom price of the line item. If a `variant_id` is supplied, the price provided here will override the variant's price.
//	  type: integer
//	title:
//	  description: The title of the line item if `variant_id` is not provided.
//	  type: string
//	  default: "Custom item"
//	quantity:
//	  description: The quantity of the line item.
//	  type: integer
//	metadata:
//	  description: The optional key-value map with additional details about the Line Item.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type Item struct {
	Title     string     `json:"title,omitempty" validate:"omitempty"`
	UnitPrice float64    `json:"unit_price,omitempty" validate:"omitempty"`
	VariantId uuid.UUID  `json:"variant_id,omitempty" validate:"omitempty"`
	Quantity  int        `json:"quantity"`
	Metadata  core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

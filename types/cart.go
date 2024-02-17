package types

import (
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type GiftCard struct {
	Code string `json:"code"`
}

type Discount struct {
	Code string `json:"code"`
}

type CartCreateProps struct {
	RegionId          uuid.UUID        `json:"region_id,omitempty" validate:"omitempty"`
	Region            *models.Region   `json:"region,omitempty" validate:"omitempty"`
	Email             string           `json:"email,omitempty" validate:"omitempty"`
	BillingAddressId  uuid.UUID        `json:"billing_address_id,omitempty" validate:"omitempty"`
	BillingAddress    *AddressPayload  `json:"billing_address,omitempty" validate:"omitempty"`
	ShippingAddressId uuid.UUID        `json:"shipping_address_id,omitempty" validate:"omitempty"`
	ShippingAddress   *AddressPayload  `json:"shipping_address,omitempty" validate:"omitempty"`
	GiftCards         []GiftCard       `json:"gift_cards,omitempty" validate:"omitempty"`
	Discounts         []Discount       `json:"discounts,omitempty" validate:"omitempty"`
	Customer          *models.Customer `json:"customer,omitempty" validate:"omitempty"`
	CustomerId        uuid.UUID        `json:"customer_id,omitempty" validate:"omitempty"`
	Type              models.CartType  `json:"type,omitempty" validate:"omitempty"`
	Context           core.JSONB       `json:"context,omitempty" validate:"omitempty"`
	Metadata          core.JSONB       `json:"metadata,omitempty" validate:"omitempty"`
	SalesChannelId    uuid.UUID        `json:"sales_channel_id,omitempty" validate:"omitempty"`
	CountryCode       string           `json:"country_code,omitempty" validate:"omitempty"`
}

// @oas:schema:StorePostCartsCartReq
// type: object
// description: "The details to update of the cart."
// properties:
//
//	region_id:
//	  type: string
//	  description: "The ID of the Region to create the Cart in. Setting the cart's region can affect the pricing of the items in the cart as well as the used currency."
//	country_code:
//	  type: string
//	  description: "The 2 character ISO country code to create the Cart in. Setting this parameter will set the country code of the shipping address."
//	  externalDocs:
//	    url: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements
//	    description: See a list of codes.
//	email:
//	  type: string
//	  description: "An email to be used on the Cart."
//	  format: email
//	sales_channel_id:
//	  type: string
//	  description: "The ID of the Sales channel to create the Cart in. The cart's sales channel affects which products can be added to the cart. If a product does not
//	   exist in the cart's sales channel, it cannot be added to the cart. If you add a publishable API key in the header of this request and specify a sales channel ID,
//	   the specified sales channel must be within the scope of the publishable API key's resources."
//	billing_address:
//	  description: "The Address to be used for billing purposes."
//	  anyOf:
//	    - $ref: "#/components/schemas/AddressPayload"
//	      description: A full billing address object.
//	    - type: string
//	      description: The billing address ID
//	shipping_address:
//	  description: "The Address to be used for shipping purposes."
//	  anyOf:
//	    - $ref: "#/components/schemas/AddressPayload"
//	      description: A full shipping address object.
//	    - type: string
//	      description: The shipping address ID
//	gift_cards:
//	  description: "An array of Gift Card codes to add to the Cart."
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - code
//	    properties:
//	      code:
//	        description: "The code of a gift card."
//	        type: string
//	discounts:
//	  description: "An array of Discount codes to add to the Cart."
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - code
//	    properties:
//	      code:
//	        description: "The code of the discount."
//	        type: string
//	customer_id:
//	  description: "The ID of the Customer to associate the Cart with."
//	  type: string
//	context:
//	  description: >-
//	    An object to provide context to the Cart. The `context` field is automatically populated with `ip` and `user_agent`
//	  type: object
//	  example:
//	    ip: "::1"
//	    user_agent: "Chrome"
type CartUpdateProps struct {
	RegionId            uuid.UUID       `json:"region_id,omitempty" validate:"omitempty"`
	CountryCode         string          `json:"country_code,omitempty" validate:"omitempty"`
	Email               string          `json:"email,omitempty" validate:"omitempty"`
	ShippingAddressId   uuid.UUID       `json:"shipping_address_id,omitempty" validate:"omitempty"`
	BillingAddressId    uuid.UUID       `json:"billing_address_id,omitempty" validate:"omitempty"`
	BillingAddress      *AddressPayload `json:"billing_address,omitempty" validate:"omitempty"`
	ShippingAddress     *AddressPayload `json:"shipping_address,omitempty" validate:"omitempty"`
	CompletedAt         *time.Time      `json:"completed_at,omitempty" validate:"omitempty"`
	PaymentAuthorizedAt *time.Time      `json:"payment_authorized_at,omitempty" validate:"omitempty"`
	GiftCards           []GiftCard      `json:"gift_cards,omitempty" validate:"omitempty"`
	Discounts           []Discount      `json:"discounts,omitempty" validate:"omitempty"`
	CustomerId          uuid.UUID       `json:"customer_id,omitempty" validate:"omitempty"`
	Context             core.JSONB      `json:"context,omitempty" validate:"omitempty"`
	Metadata            core.JSONB      `json:"metadata,omitempty" validate:"omitempty"`
	SalesChannelId      uuid.UUID       `json:"sales_channel_id,omitempty" validate:"omitempty"`
}

type FilterableCartProps struct {
	core.FilterModel
}

// @oas:schema:StorePostCartReq
// type: object
// description: "The details of the cart to be created."
// properties:
//
//	region_id:
//	  type: string
//	  description: "The ID of the Region to create the Cart in. Setting the cart's region can affect the pricing of the items in the cart as well as the used currency.
//	   If this parameter is not provided, the first region in the store is used by default."
//	sales_channel_id:
//	  type: string
//	  description: "The ID of the Sales channel to create the Cart in. The cart's sales channel affects which products can be added to the cart. If a product does not
//	   exist in the cart's sales channel, it cannot be added to the cart. If you add a publishable API key in the header of this request and specify a sales channel ID,
//	   the specified sales channel must be within the scope of the publishable API key's resources. If you add a publishable API key in the header of this request,
//	   you don't specify a sales channel ID, and the publishable API key is associated with one sales channel, that sales channel will be attached to the cart.
//	   If no sales channel is passed and no publishable API key header is passed or the publishable API key isn't associated with any sales channel,
//	   the cart will not be associated with any sales channel."
//	country_code:
//	  type: string
//	  description: "The two character ISO country code to create the Cart in. Setting this parameter will set the country code of the shipping address."
//	  externalDocs:
//	   url: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements
//	   description: See a list of codes.
//	items:
//	  description: "An array of product variants to generate line items from."
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - variant_id
//	      - quantity
//	    properties:
//	      variant_id:
//	        description: The ID of the Product Variant.
//	        type: string
//	      quantity:
//	        description: The quantity to add into the cart.
//	        type: integer
//	context:
//	  description: >-
//	    An object to provide context to the Cart. The `context` field is automatically populated with `ip` and `user_agent`
//	  type: object
//	  example:
//	    ip: "::1"
//	    user_agent: "Chrome"
type CreateCart struct {
	RegionId       uuid.UUID                            `json:"region_id,omitempty" validate:"omitempty,string"`
	CountryCode    string                               `json:"country_code,omitempty" validate:"omitempty,string"`
	Items          []CreateClaimItemAdditionalItemInput `json:"items,omitempty" validate:"omitempty,dive"`
	Context        core.JSONB                           `json:"context,omitempty" validate:"omitempty"`
	SalesChannelId uuid.UUID                            `json:"sales_channel_id,omitempty" validate:"omitempty,string"`
}

// @oas:schema:StorePostCartsCartPaymentSessionUpdateReq
// type: object
// required:
//   - data
//
// properties:
//
//	data:
//	  type: object
//	  description: The data to update the payment session with.
type UpdatePaymentSession struct {
	Data core.JSONB `json:"data"`
}

// @oas:schema:StorePostCartsCartShippingMethodReq
// type: object
// description: "The details of the shipping method to add to the cart."
// required:
//   - option_id
//
// properties:
//
//	option_id:
//	  type: string
//	  description: ID of the shipping option to create the method from.
//	data:
//	  type: object
//	  description: Used to hold any data that the shipping method may need to process the fulfillment of the order. This depends on the fulfillment provider you're using.
type AddShippingMethod struct {
	OptionId uuid.UUID  `json:"option_id"`
	Data     core.JSONB `json:"data,omitempty" validate:"omitempty"`
}

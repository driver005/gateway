package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:AdminPostStoreReq
// type: object
// description: "The details to update of the store."
// properties:
//
//	name:
//	  description: "The name of the Store"
//	  type: string
//	swap_link_template:
//	  description: >-
//	    A template for Swap links - use `{{cart_id}}` to insert the Swap Cart ID
//	  type: string
//	  example: "http://example.com/swaps/{{cart_id}}"
//	payment_link_template:
//	  description: "A template for payment links - use `{{cart_id}}` to insert the Cart ID"
//	  example: "http://example.com/payments/{{cart_id}}"
//	  type: string
//	invite_link_template:
//	  description: "A template for invite links - use `{{invite_token}}` to insert the invite token"
//	  example: "http://example.com/invite?token={{invite_token}}"
//	  type: string
//	default_currency_code:
//	  description: "The default currency code of the Store."
//	  type: string
//	  externalDocs:
//	    url: https://en.wikipedia.org/wiki/ISO_4217#Active_codes
//	    description: See a list of codes.
//	currencies:
//	  description: "Array of available currencies in the store. Each currency is in 3 character ISO code format."
//	  type: array
//	  items:
//	    type: string
//	    externalDocs:
//	      url: https://en.wikipedia.org/wiki/ISO_4217#Active_codes
//	      description: See a list of codes.
//	metadata:
//	  description: "An optional set of key-value pairs with additional information."
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type UpdateStoreInput struct {
	Name                  string     `json:"name,omitempty" validate:"omitempty"`
	SwapLinkTemplate      string     `json:"swap_link_template,omitempty" validate:"omitempty"`
	PaymentLinkTemplate   string     `json:"payment_link_template,omitempty" validate:"omitempty"`
	InviteLinkTemplate    string     `json:"invite_link_template,omitempty" validate:"omitempty"`
	DefaultCurrencyCode   string     `json:"default_currency_code,omitempty" validate:"omitempty"`
	Currencies            []string   `json:"currencies,omitempty" validate:"omitempty"`
	Metadata              core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
	DefaultSalesChannelId uuid.UUID  `json:"default_sales_channel_id,omitempty" validate:"omitempty"`
}

// type ExtendedStoreDTO struct {
// 	*models.Store
// 	PaymentProviders     []PaymentProvider
// 	FulfillmentProviders []FulfillmentProvider
// 	FeatureFlags         FeatureFlagsResponse
// 	Modules              ModulesResponse
// }

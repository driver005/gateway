package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FilterableRegion struct {
	core.FilterModel
	Name string `json:"name,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostRegionsRegionReq
// type: object
// description: "The details to update of the regions."
// properties:
//
//	name:
//	  description: "The name of the Region"
//	  type: string
//	currency_code:
//	  description: "The 3 character ISO currency code to use in the Region."
//	  type: string
//	  externalDocs:
//	    url: https://en.wikipedia.org/wiki/ISO_4217#Active_codes
//	    description: See a list of codes.
//	automatic_taxes:
//	  description: >-
//	    If set to `true`, the Medusa backend will automatically calculate taxes for carts in this region. If set to `false`, the taxes must be calculated manually.
//	  externalDocs:
//	    url: https://docs.medusajs.com/modules/taxes/storefront/manual-calculation
//	    description: How to calculate taxes in a storefront.
//	  type: boolean
//	gift_cards_taxable:
//	  description: >-
//	    If set to `true`, taxes will be applied on gift cards.
//	  type: boolean
//	tax_provider_id:
//	  description: "The ID of the tax provider to use. If none provided, the system tax provider is used."
//	  type: string
//	tax_code:
//	  description: "The tax code of the Region."
//	  type: string
//	tax_rate:
//	  description: "The tax rate to use in the Region."
//	  type: number
//	includes_tax:
//	  x-featureFlag: "tax_inclusive_pricing"
//	  description: "Whether taxes are included in the prices of the region."
//	  type: boolean
//	payment_providers:
//	  description: "A list of Payment Provider IDs that can be used in the Region"
//	  type: array
//	  items:
//	    type: string
//	fulfillment_providers:
//	  description: "A list of Fulfillment Provider IDs that can be used in the Region"
//	  type: array
//	  items:
//	    type: string
//	countries:
//	  description: "A list of countries' 2 ISO characters that should be included in the Region."
//	  type: array
//	  items:
//	    type: string
type UpdateRegionInput struct {
	Name                 string     `json:"name,omitempty" validate:"omitempty"`
	CurrencyCode         string     `json:"currency_code,omitempty" validate:"omitempty"`
	TaxCode              string     `json:"tax_code,omitempty" validate:"omitempty"`
	TaxRate              float64    `json:"tax_rate,omitempty" validate:"omitempty"`
	GiftCardsTaxable     bool       `json:"gift_cards_taxable,omitempty" validate:"omitempty"`
	AutomaticTaxes       bool       `json:"automatic_taxes,omitempty" validate:"omitempty"`
	TaxProviderId        uuid.UUID  `json:"tax_provider_id,omitempty" validate:"omitempty"`
	PaymentProviders     uuid.UUIDs `json:"payment_providers,omitempty" validate:"omitempty"`
	FulfillmentProviders uuid.UUIDs `json:"fulfillment_providers,omitempty" validate:"omitempty"`
	Countries            []string   `json:"countries,omitempty" validate:"omitempty"`
	IncludesTax          bool       `json:"includes_tax,omitempty" validate:"omitempty"`
	Metadata             core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostRegionsReq
// type: object
// description: "The details of the region to create."
// required:
//   - name
//   - currency_code
//   - tax_rate
//   - payment_providers
//   - fulfillment_providers
//   - countries
//
// properties:
//
//	name:
//	  description: "The name of the Region"
//	  type: string
//	currency_code:
//	  description: "The 3 character ISO currency code to use in the Region."
//	  type: string
//	  externalDocs:
//	    url: https://en.wikipedia.org/wiki/ISO_4217#Active_codes
//	    description: See a list of codes.
//	tax_code:
//	  description: "The tax code of the Region."
//	  type: string
//	tax_rate:
//	  description: "The tax rate to use in the Region."
//	  type: number
//	payment_providers:
//	  description: "A list of Payment Provider IDs that can be used in the Region"
//	  type: array
//	  items:
//	    type: string
//	fulfillment_providers:
//	  description: "A list of Fulfillment Provider IDs that can be used in the Region"
//	  type: array
//	  items:
//	    type: string
//	countries:
//	  description: "A list of countries' 2 ISO characters that should be included in the Region."
//	  example: ["US"]
//	  type: array
//	  items:
//	    type: string
//	includes_tax:
//	  x-featureFlag: "tax_inclusive_pricing"
//	  description: "Whether taxes are included in the prices of the region."
//	  type: boolean
type CreateRegionInput struct {
	Name                 string     `json:"name"`
	CurrencyCode         string     `json:"currency_code"`
	TaxCode              string     `json:"tax_code,omitempty" validate:"omitempty"`
	TaxRate              float64    `json:"tax_rate"`
	PaymentProviders     uuid.UUIDs `json:"payment_providers"`
	FulfillmentProviders uuid.UUIDs `json:"fulfillment_providers"`
	Countries            []string   `json:"countries"`
	IncludesTax          bool       `json:"includes_tax,omitempty" validate:"omitempty"`
	Metadata             core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostRegionsRegionCountriesReq
// type: object
// description: "The details of the country to add to the region."
// required:
//   - country_code
//
// properties:
//
//	country_code:
//	  description: "The 2 character ISO code for the Country."
//	  type: string
//	  externalDocs:
//	    url: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements
//	    description: See a list of codes.
type RegionCountries struct {
	CountryCode string `json:"country_code"`
}

// @oas:schema:AdminPostRegionsRegionFulfillmentProvidersReq
// type: object
// description: "The details of the fulfillment provider to add to the region."
// required:
//   - provider_id
//
// properties:
//
//	provider_id:
//	  description: "The ID of the Fulfillment Provider."
//	  type: string
type RegionFulfillmentProvider struct {
	ProviderId uuid.UUID `json:"provider_id"`
}

// @oas:schema:AdminPostRegionsRegionPaymentProvidersReq
// type: object
// description: "The details of the payment provider to add to the region."
// required:
//   - provider_id
//
// properties:
//
//	provider_id:
//	  description: "The ID of the Payment Provider."
//	  type: string
type RegionPaymentProvider struct {
	ProviderId uuid.UUID `json:"provider_id"`
}

package types

import (
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type FilterablePriceList struct {
	core.FilterModel

	Q              string                   `json:"q,omitempty" validate:"omitempty"`
	Status         []models.PriceListStatus `json:"status,omitempty" validate:"omitempty"`
	Name           string                   `json:"name,omitempty" validate:"omitempty"`
	CustomerGroups []string                 `json:"customer_groups,omitempty" validate:"omitempty"`
	Description    string                   `json:"description,omitempty" validate:"omitempty"`
	Type           []models.PriceListType   `json:"type,omitempty" validate:"omitempty"`
}

type AdminPriceListPricesUpdateReq struct {
	Id           uuid.UUID `json:"id,omitempty" validate:"omitempty"`
	RegionId     uuid.UUID `json:"region_id,omitempty" validate:"omitempty"`
	CurrencyCode string    `json:"currency_code,omitempty" validate:"omitempty"`
	VariantId    uuid.UUID `json:"variant_id"`
	Amount       float64   `json:"amount"`
	MinQuantity  int       `json:"min_quantity,omitempty" validate:"omitempty"`
	MaxQuantity  int       `json:"max_quantity,omitempty" validate:"omitempty"`
}

type AdminPriceListPricesCreateReq struct {
	RegionId     uuid.UUID `json:"region_id,omitempty" validate:"omitempty"`
	CurrencyCode string    `json:"currency_code,omitempty" validate:"omitempty"`
	Amount       float64   `json:"amount"`
	VariantId    uuid.UUID `json:"variant_id"`
	MinQuantity  int       `json:"min_quantity,omitempty" validate:"omitempty"`
	MaxQuantity  int       `json:"max_quantity,omitempty" validate:"omitempty"`
}

type CustomerGroups struct {
	Id uuid.UUID `json:"id"`
}

// @oas:schema:AdminPostPriceListsPriceListReq
// type: object
// description: "The details of the price list to create."
// required:
//   - name
//   - description
//   - type
//   - prices
//
// properties:
//
//	name:
//	  description: "The name of the Price List."
//	  type: string
//	description:
//	  description: "The description of the Price List."
//	  type: string
//	starts_at:
//	  description: "The date with timezone that the Price List starts being valid."
//	  type: string
//	  format: date
//	ends_at:
//	  description: "The date with timezone that the Price List ends being valid."
//	  type: string
//	  format: date
//	type:
//	  description: The type of the Price List.
//	  type: string
//	  enum:
//	   - sale
//	   - override
//	status:
//	  description: >-
//	    The status of the Price List. If the status is set to `draft`, the prices created in the price list will not be available of the customer.
//	  type: string
//	  enum:
//	    - active
//	    - draft
//	prices:
//	   description: The prices of the Price List.
//	   type: array
//	   items:
//	     type: object
//	     required:
//	       - amount
//	       - variant_id
//	     properties:
//	       region_id:
//	         description: The ID of the Region for which the price is used. This is only required if `currecny_code` is not provided.
//	         type: string
//	       currency_code:
//	         description: The 3 character ISO currency code for which the price will be used. This is only required if `region_id` is not provided.
//	         type: string
//	         externalDocs:
//	           url: https://en.wikipedia.org/wiki/ISO_4217#Active_codes
//	           description: See a list of codes.
//	       amount:
//	         description: The amount to charge for the Product Variant.
//	         type: integer
//	       variant_id:
//	         description: The ID of the Variant for which the price is used.
//	         type: string
//	       min_quantity:
//	         description: The minimum quantity for which the price will be used.
//	         type: integer
//	       max_quantity:
//	         description: The maximum quantity for which the price will be used.
//	         type: integer
//	customer_groups:
//	  type: array
//	  description: An array of customer groups that the Price List applies to.
//	  items:
//	    type: object
//	    required:
//	      - id
//	    properties:
//	      id:
//	        description: The ID of a customer group
//	        type: string
//	includes_tax:
//	   description: "Tax included in prices of price list"
//	   x-featureFlag: "tax_inclusive_pricing"
//	   type: boolean
type CreatePriceListInput struct {
	Name           string                      `json:"name"`
	Description    string                      `json:"description"`
	Type           models.PriceListType        `json:"type"`
	Status         models.PriceListStatus      `json:"status,omitempty" validate:"omitempty"`
	Prices         []PriceListPriceCreateInput `json:"prices"`
	CustomerGroups []CustomerGroups            `json:"customer_groups,omitempty" validate:"omitempty"`
	StartsAt       *time.Time                  `json:"starts_at,omitempty" validate:"omitempty"`
	EndsAt         *time.Time                  `json:"ends_at,omitempty" validate:"omitempty"`
	IncludesTax    bool                        `json:"includes_tax,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostPriceListsPriceListPriceListReq
// type: object
// description: "The details to update of the payment collection."
// properties:
//
//	name:
//	  description: "The name of the Price List"
//	  type: string
//	description:
//	  description: "The description of the Price List."
//	  type: string
//	starts_at:
//	  description: "The date with timezone that the Price List starts being valid."
//	  type: string
//	  format: date
//	ends_at:
//	  description: "The date with timezone that the Price List ends being valid."
//	  type: string
//	  format: date
//	type:
//	  description: The type of the Price List.
//	  type: string
//	  enum:
//	   - sale
//	   - override
//	status:
//	  description: >-
//	    The status of the Price List. If the status is set to `draft`, the prices created in the price list will not be available of the customer.
//	  type: string
//	  enum:
//	   - active
//	   - draft
//	prices:
//	  description: The prices of the Price List.
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - amount
//	      - variant_id
//	    properties:
//	      id:
//	        description: The ID of the price.
//	        type: string
//	      region_id:
//	        description: The ID of the Region for which the price is used. This is only required if `currecny_code` is not provided.
//	        type: string
//	      currency_code:
//	        description: The 3 character ISO currency code for which the price will be used. This is only required if `region_id` is not provided.
//	        type: string
//	        externalDocs:
//	           url: https://en.wikipedia.org/wiki/ISO_4217#Active_codes
//	           description: See a list of codes.
//	      variant_id:
//	        description: The ID of the Variant for which the price is used.
//	        type: string
//	      amount:
//	        description: The amount to charge for the Product Variant.
//	        type: integer
//	      min_quantity:
//	        description: The minimum quantity for which the price will be used.
//	        type: integer
//	      max_quantity:
//	        description: The maximum quantity for which the price will be used.
//	        type: integer
//	customer_groups:
//	  type: array
//	  description: An array of customer groups that the Price List applies to.
//	  items:
//	    type: object
//	    required:
//	      - id
//	    properties:
//	      id:
//	        description: The ID of a customer group
//	        type: string
//	includes_tax:
//	  description: "Tax included in prices of price list"
//	  x-featureFlag: "tax_inclusive_pricing"
//	  type: boolean
type UpdatePriceListInput struct {
	Name           string                      `json:"name,omitempty" validate:"omitempty"`
	Description    string                      `json:"description,omitempty" validate:"omitempty"`
	StartsAt       *time.Time                  `json:"starts_at,omitempty" validate:"omitempty"`
	EndsAt         *time.Time                  `json:"ends_at,omitempty" validate:"omitempty"`
	Status         models.PriceListStatus      `json:"status,omitempty" validate:"omitempty"`
	Type           models.PriceListType        `json:"type,omitempty" validate:"omitempty"`
	IncludesTax    bool                        `json:"includes_tax,omitempty" validate:"omitempty"`
	Prices         []PriceListPriceCreateInput `json:"prices,omitempty" validate:"omitempty"`
	CustomerGroups []CustomerGroups            `json:"customer_groups,omitempty" validate:"omitempty"`
}

type PriceListPriceUpdateInput struct {
	Id           uuid.UUID `json:"id,omitempty" validate:"omitempty"`
	VariantId    uuid.UUID `json:"variant_id,omitempty" validate:"omitempty"`
	RegionId     uuid.UUID `json:"region_id,omitempty" validate:"omitempty"`
	CurrencyCode string    `json:"currency_code,omitempty" validate:"omitempty"`
	Amount       float64   `json:"amount,omitempty" validate:"omitempty"`
	MinQuantity  int       `json:"min_quantity,omitempty" validate:"omitempty"`
	MaxQuantity  int       `json:"max_quantity,omitempty" validate:"omitempty"`
}

type PriceListPriceCreateInput struct {
	RegionId     uuid.UUID `json:"region_id,omitempty" validate:"omitempty"`
	CurrencyCode string    `json:"currency_code,omitempty" validate:"omitempty"`
	VariantId    uuid.UUID `json:"variant_id"`
	Amount       float64   `json:"amount"`
	MinQuantity  int       `json:"min_quantity,omitempty" validate:"omitempty"`
	MaxQuantity  int       `json:"max_quantity,omitempty" validate:"omitempty"`
}

type PriceListLoadConfig struct {
	IncludeDiscountPrices bool      `json:"include_discount_prices,omitempty" validate:"omitempty"`
	CustomerId            uuid.UUID `json:"customer_id,omitempty" validate:"omitempty"`
	CartId                uuid.UUID `json:"cart_id,omitempty" validate:"omitempty"`
	RegionId              uuid.UUID `json:"region_id,omitempty" validate:"omitempty"`
	CurrencyCode          string    `json:"currency_code,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostPriceListPricesPricesReq
// type: object
// description: "The details of the prices to add."
// properties:
//
//	prices:
//	  description: The prices to update or add.
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - amount
//	      - variant_id
//	    properties:
//	      id:
//	        description: The ID of the price.
//	        type: string
//	      region_id:
//	        description: The ID of the Region for which the price is used. This is only required if `currecny_code` is not provided.
//	        type: string
//	      currency_code:
//	        description: The 3 character ISO currency code for which the price will be used. This is only required if `region_id` is not provided.
//	        type: string
//	        externalDocs:
//	          url: https://en.wikipedia.org/wiki/ISO_4217#Active_codes
//	          description: See a list of codes.
//	      variant_id:
//	        description: The ID of the Variant for which the price is used.
//	        type: string
//	      amount:
//	        description: The amount to charge for the Product Variant.
//	        type: integer
//	      min_quantity:
//	        description: The minimum quantity for which the price will be used.
//	        type: integer
//	      max_quantity:
//	        description: The maximum quantity for which the price will be used.
//	        type: integer
//	override:
//	  description: >-
//	    If set to `true`, the prices will replace all existing prices associated with the Price List.
//	  type: boolean
type AddPriceListPrices struct {
	Prices   []PriceListPriceCreateInput `json:"prices"`
	Override bool                        `json:"override,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminDeletePriceListPricesPricesReq
// type: object
// description: "The details of the prices to delete."
// properties:
//
//	price_ids:
//	  description: The IDs of the prices to delete.
//	  type: array
//	  items:
//	    type: string
type DeletePriceListPrices struct {
	PriceIds uuid.UUIDs `json:"price_ids"`
}

// @oas:schema:AdminDeletePriceListsPriceListProductsPricesBatchReq
// type: object
// description: "The details of the products' prices to delete."
// properties:
//
//	product_ids:
//	  description: The IDs of the products to delete their associated prices.
//	  type: array
//	  items:
//	    type: string
type DeletePriceListPricesBatch struct {
	ProductIds uuid.UUIDs `json:"product_ids"`
}

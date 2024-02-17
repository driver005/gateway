package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FilterableTaxRate struct {
	core.FilterModel

	RegionId uuid.UUIDs       `json:"region_id,omitempty" validate:"omitempty"`
	Code     []string         `json:"code,omitempty" validate:"omitempty"`
	Name     []string         `json:"name,omitempty" validate:"omitempty"`
	Rate     core.NumberModel `json:"rate,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostTaxRatesTaxRateReq
// type: object
// description: "The details to update of the tax rate."
// properties:
//
//	code:
//	  type: string
//	  description: "The code of the tax rate."
//	name:
//	  type: string
//	  description: "The name of the tax rate."
//	region_id:
//	  type: string
//	  description: "The ID of the Region that the tax rate belongs to."
//	rate:
//	  type: number
//	  description: "The numeric rate to charge."
//	products:
//	  type: array
//	  description: "The IDs of the products associated with this tax rate"
//	  items:
//	    type: string
//	shipping_options:
//	  type: array
//	  description: "The IDs of the shipping options associated with this tax rate"
//	  items:
//	    type: string
//	product_types:
//	  type: array
//	  description: "The IDs of the types of product types associated with this tax rate"
//	  items:
//	    type: string
type UpdateTaxRateInput struct {
	RegionId uuid.UUID `json:"region_id,omitempty" validate:"omitempty"`
	Code     string    `json:"code,omitempty" validate:"omitempty"`
	Name     string    `json:"name,omitempty" validate:"omitempty"`
	Rate     float64   `json:"rate,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostTaxRatesReq
// type: object
// description: "The details of the tax rate to create."
// required:
//   - code
//   - name
//   - region_id
//
// properties:
//
//	code:
//	  type: string
//	  description: "The code of the tax rate."
//	name:
//	  type: string
//	  description: "The name of the tax rate."
//	region_id:
//	  type: string
//	  description: "The ID of the Region that the tax rate belongs to."
//	rate:
//	  type: number
//	  description: "The numeric rate to charge."
//	products:
//	  type: array
//	  description: "The IDs of the products associated with this tax rate."
//	  items:
//	    type: string
//	shipping_options:
//	  type: array
//	  description: "The IDs of the shipping options associated with this tax rate"
//	  items:
//	    type: string
//	product_types:
//	  type: array
//	  description: "The IDs of the types of products associated with this tax rate"
//	  items:
//	    type: string
type CreateTaxRateInput struct {
	RegionId uuid.UUID `json:"region_id"`
	Code     string    `json:"code"`
	Name     string    `json:"name"`
	Rate     float64   `json:"rate,omitempty" validate:"omitempty"`
}

type TaxRateListByConfig struct {
	RegionId uuid.UUID `json:"region_id,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostTaxRatesTaxRateProductTypesReq
// type: object
// description: "The product types to add to the tax rate."
// required:
//   - product_types
//
// properties:
//
//	product_types:
//	  type: array
//	  description: "The IDs of the types of products to associate with this tax rate"
//	  items:
//	    type: string
type TaxRateProductTypes struct {
	ProductTypes uuid.UUIDs `json:"product_types"`
}

// @oas:schema:AdminPostTaxRatesTaxRateShippingOptionsReq
// type: object
// description: "The details of the shipping options to associate with the tax rate."
// required:
//   - shipping_options
//
// properties:
//
//	shipping_options:
//	  type: array
//	  description: "The IDs of the shipping options to associate with this tax rate"
//	  items:
//	    type: string
type TaxRateShippingOptions struct {
	ShippingOptions uuid.UUIDs `json:"shipping_options"`
}

// @oas:schema:AdminPostTaxRatesTaxRateProductsReq
// type: object
// description: "The details of the products to associat with the tax rate."
// required:
//   - products
//
// properties:
//
//	products:
//	  type: array
//	  description: "The IDs of the products to associate with this tax rate"
//	  items:
//	    type: string
type TaxRateProducts struct {
	Products uuid.UUIDs `json:"products"`
}

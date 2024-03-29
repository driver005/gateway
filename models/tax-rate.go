package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:TaxRate
// title: "Tax Rate"
// description: "A Tax Rate can be used to define a custom rate to charge on specified products, product types, and shipping options within a given region."
// type: object
// required:
//   - code
//   - created_at
//   - id
//   - metadata
//   - name
//   - rate
//   - region_id
//   - updated_at
//
// properties:
//
//	id:
//	  description: The tax rate's ID
//	  type: string
//	  example: txr_01G8XDBAWKBHHJRKH0AV02KXBR
//	rate:
//	  description: The numeric rate to charge
//	  nullable: true
//	  type: number
//	  example: 10
//	code:
//	  description: A code to identify the tax type by
//	  nullable: true
//	  type: string
//	  example: tax01
//	name:
//	  description: A human friendly name for the tax
//	  type: string
//	  example: Tax Example
//	region_id:
//	  description: The ID of the region that the rate belongs to.
//	  type: string
//	  example: reg_01G1G5V26T9H8Y0M4JNE3YGA4G
//	region:
//	  description: The details of the region that the rate belongs to.
//	  x-expandable: "region"
//	  nullable: true
//	  $ref: "#/components/schemas/Region"
//	products:
//	  description: The details of the products that belong to this tax rate.
//	  type: array
//	  x-expandable: "products"
//	  items:
//	    $ref: "#/components/schemas/Product"
//	product_types:
//	  description: The details of the product types that belong to this tax rate.
//	  type: array
//	  x-expandable: "product_types"
//	  items:
//	    $ref: "#/components/schemas/ProductType"
//	shipping_options:
//	  description: The details of the shipping options that belong to this tax rate.
//	  type: array
//	  x-expandable: "shipping_options"
//	  items:
//	    $ref: "#/components/schemas/ShippingOption"
//	product_count:
//	  description: The count of products
//	  type: integer
//	  example: 10
//	product_type_count:
//	  description: The count of product types
//	  type: integer
//	  example: 2
//	shipping_option_count:
//	  description: The count of shipping options
//	  type: integer
//	  example: 1
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
type TaxRate struct {
	core.BaseModel

	Rate                float64          `json:"rate" gorm:"column:rate"`
	Code                string           `json:"code" gorm:"column:code"`
	Name                string           `json:"name" gorm:"column:name"`
	RegionId            uuid.NullUUID    `json:"region_id" gorm:"column:region_id"`
	Region              *Region          `json:"region" gorm:"foreignKey:RegionId"`
	Products            []Product        `json:"products" gorm:"many2many:product_tax_rate"`
	ProductTypes        []ProductType    `json:"product_types" gorm:"many2many:product_type_tax_rate"`
	ShippingOptions     []ShippingOption `json:"shipping_options" gorm:"many2many:shipping_tax_rate"`
	ProductCount        int32            `json:"product_count" gorm:"column:product_count"`
	ProductTypeCount    int32            `json:"product_type_count" gorm:"column:product_type_count"`
	ShippingOptionCount int32            `json:"shipping_option_count" gorm:"column:shipping_option_count"`
}

package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:ProductTypeTaxRate
// title: "Product Type Tax Rate"
// description: "This represents the association between a tax rate and a product type to indicate that the product type is taxed in a different way than the default."
// type: object
// required:
//   - created_at
//   - metadata
//   - product_type_id
//   - rate_id
//   - updated_at
//
// properties:
//
//	product_type_id:
//	  description: The ID of the Product type
//	  type: string
//	  example: ptyp_01G8X9A7ESKAJXG2H0E6F1MW7A
//	product_type:
//	  description: The details of the product type.
//	  x-expandable: "product_type"
//	  nullable: true
//	  $ref: "#/components/schemas/ProductType"
//	rate_id:
//	  description: The id of the Tax Rate
//	  type: string
//	  example: txr_01G8XDBAWKBHHJRKH0AV02KXBR
//	tax_rate:
//	  description: The details of the tax rate.
//	  x-expandable: "tax_rate"
//	  nullable: true
//	  $ref: "#/components/schemas/TaxRate"
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
type ProductTypeTaxRate struct {
	ProductTypeId uuid.NullUUID `json:"product_type_id" gorm:"column:product_type_id"`
	ProductType   *ProductType  `json:"product_type" gorm:"foreignKey:ProductTypeId"`
	RateId        uuid.NullUUID `json:"rate_id" gorm:"column:rate_id"`
	TaxRate       *TaxRate      `json:"tax_rate" gorm:"foreignKey:RateId"`
	Metadata      core.JSONB    `json:"metadata" gorm:"column:metadata"`
}

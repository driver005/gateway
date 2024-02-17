package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:ProductTaxRate
// title: "Product Tax Rate"
// description: "This represents the association between a tax rate and a product to indicate that the product is taxed in a way different than the default."
// type: object
// required:
//   - created_at
//   - metadata
//   - product_id
//   - rate_id
//   - updated_at
//
// properties:
//
//	product_id:
//	  description: The ID of the Product
//	  type: string
//	  example: prod_01G1G5V2MBA328390B5AXJ610F
//	product:
//	  description: The details of the product.
//	  x-expandable: "product"
//	  nullable: true
//	  $ref: "#/components/schemas/Product"
//	rate_id:
//	  description: The ID of the Tax Rate
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
type ProductTaxRate struct {
	ProductId uuid.NullUUID `json:"product_id"`
	Product   *Product      `json:"product" gorm:"foreignKey:id;references:product_id"`
	RateId    uuid.NullUUID `json:"rate_id"`
	TaxRate   *TaxRate      `json:"tax_rate" gorm:"foreignKey:id;references:rate_id"`
	Metadata  core.JSONB    `json:"metadata" gorm:"default:null"`
}

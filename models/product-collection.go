package models

import "github.com/driver005/gateway/core"

// @oas:schema:ProductCollection
// title: "Product Collection"
// description: "A Product Collection allows grouping together products for promotional purposes. For example, an admin can create a Summer collection, add products to it, and showcase it on the storefront."
// type: object
// required:
//   - created_at
//   - deleted_at
//   - handle
//   - id
//   - metadata
//   - title
//   - updated_at
//
// properties:
//
//	id:
//	  description: The product collection's ID
//	  type: string
//	  example: pcol_01F0YESBFAZ0DV6V831JXWH0BG
//	title:
//	  description: The title that the Product Collection is identified by.
//	  type: string
//	  example: Summer Collection
//	handle:
//	  description: A unique string that identifies the Product Collection - can for example be used in slug structures.
//	  nullable: true
//	  type: string
//	  example: summer-collection
//	products:
//	  description: The details of the products that belong to this product collection.
//	  type: array
//	  x-expandable: "products"
//	  items:
//	    $ref: "#/components/schemas/Product"
//	created_at:
//	  description: The date with timezone at which the resource was created.
//	  type: string
//	  format: date-time
//	updated_at:
//	  description: The date with timezone at which the resource was updated.
//	  type: string
//	  format: date-time
//	deleted_at:
//	  description: The date with timezone at which the resource was deleted.
//	  nullable: true
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
type ProductCollection struct {
	core.SoftDeletableModel

	Title    string    `json:"title" gorm:"column:title"`
	Handle   string    `json:"handle" gorm:"column:handle"`
	Products []Product `json:"products" gorm:"foreignKey:Id"`
}

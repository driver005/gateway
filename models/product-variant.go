package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:ProductVariant
// title: "Product Variant"
// description: "A Product Variant represents a Product with a specific set of Product Option configurations. The maximum number of Product Variants that a Product can have is given by the number of available Product Option combinations. A product must at least have one product variant."
// type: object
// required:
//   - allow_backorder
//   - barcode
//   - created_at
//   - deleted_at
//   - ean
//   - height
//   - hs_code
//   - id
//   - inventory_quantity
//   - length
//   - manage_inventory
//   - material
//   - metadata
//   - mid_code
//   - origin_country
//   - product_id
//   - sku
//   - title
//   - upc
//   - updated_at
//   - weight
//   - width
//
// properties:
//
//	id:
//	  description: The product variant's ID
//	  type: string
//	  example: variant_01G1G5V2MRX2V3PVSR2WXYPFB6
//	title:
//	  description: A title that can be displayed for easy identification of the Product Variant.
//	  type: string
//	  example: Small
//	product_id:
//	  description: The ID of the product that the product variant belongs to.
//	  type: string
//	  example: prod_01G1G5V2MBA328390B5AXJ610F
//	product:
//	  description: The details of the product that the product variant belongs to.
//	  x-expandable: "product"
//	  nullable: true
//	  $ref: "#/components/schemas/Product"
//	prices:
//	  description: The details of the prices of the Product Variant, each represented as a Money Amount. Each Money Amount represents a price in a given currency or a specific Region.
//	  type: array
//	  x-expandable: "prices"
//	  items:
//	    $ref: "#/components/schemas/MoneyAmount"
//	sku:
//	  description: The unique stock keeping unit used to identify the Product Variant. This will usually be a unique identifer for the item that is to be shipped, and can be referenced across multiple systems.
//	  nullable: true
//	  type: string
//	  example: shirt-123
//	barcode:
//	  description: A generic field for a GTIN number that can be used to identify the Product Variant.
//	  nullable: true
//	  type: string
//	  example: null
//	ean:
//	  description: An EAN barcode number that can be used to identify the Product Variant.
//	  nullable: true
//	  type: string
//	  example: null
//	upc:
//	  description: A UPC barcode number that can be used to identify the Product Variant.
//	  nullable: true
//	  type: string
//	  example: null
//	variant_rank:
//	  description: The ranking of this variant
//	  nullable: true
//	  type: number
//	  default: 0
//	inventory_quantity:
//	  description: The current quantity of the item that is stocked.
//	  type: integer
//	  example: 100
//	allow_backorder:
//	  description: Whether the Product Variant should be purchasable when `inventory_quantity` is 0.
//	  type: boolean
//	  default: false
//	manage_inventory:
//	  description: Whether Medusa should manage inventory for the Product Variant.
//	  type: boolean
//	  default: true
//	hs_code:
//	  description: The Harmonized System code of the Product Variant. May be used by Fulfillment Providers to pass customs information to shipping carriers.
//	  nullable: true
//	  type: string
//	  example: null
//	origin_country:
//	  description: The country in which the Product Variant was produced. May be used by Fulfillment Providers to pass customs information to shipping carriers.
//	  nullable: true
//	  type: string
//	  example: null
//	mid_code:
//	  description: The Manufacturers Identification code that identifies the manufacturer of the Product Variant. May be used by Fulfillment Providers to pass customs information to shipping carriers.
//	  nullable: true
//	  type: string
//	  example: null
//	material:
//	  description: The material and composition that the Product Variant is made of, May be used by Fulfillment Providers to pass customs information to shipping carriers.
//	  nullable: true
//	  type: string
//	  example: null
//	weight:
//	  description: The weight of the Product Variant. May be used in shipping rate calculations.
//	  nullable: true
//	  type: number
//	  example: null
//	length:
//	  description: "The length of the Product Variant. May be used in shipping rate calculations."
//	  nullable: true
//	  type: number
//	  example: null
//	height:
//	  description: The height of the Product Variant. May be used in shipping rate calculations.
//	  nullable: true
//	  type: number
//	  example: null
//	width:
//	  description: The width of the Product Variant. May be used in shipping rate calculations.
//	  nullable: true
//	  type: number
//	  example: null
//	options:
//	  description: The details of the product options that this product variant defines values for.
//	  type: array
//	  x-expandable: "options"
//	  items:
//	    $ref: "#/components/schemas/ProductOptionValue"
//	inventory_items:
//	  description: The details inventory items of the product variant.
//	  type: array
//	  x-expandable: "inventory_items"
//	  items:
//	    $ref: "#/components/schemas/ProductVariantInventoryItem"
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
//	purchasable:
//	  description: |
//	     Only used with the inventory modules.
//	     A boolean value indicating whether the Product Variant is purchasable.
//	     A variant is purchasable if:
//	       - inventory is not managed
//	       - it has no inventory items
//	       - it is in stock
//	       - it is backorderable.
//	  type: boolean
type ProductVariant struct {
	core.SoftDeletableModel

	AllowBackorder    bool                 `json:"allow_backorder" gorm:"column:allow_backorder;default:false"`
	Purchasable       bool                 `json:"purchasable" gorm:"column:purchasable"`
	Barcode           string               `json:"barcode" gorm:"column:barcode"`
	Ean               string               `json:"ean" gorm:"column:ean"`
	Height            float64              `json:"height" gorm:"column:height"`
	HsCode            string               `json:"hs_code" gorm:"column:hs_code"`
	InventoryQuantity int                  `json:"inventory_quantity" gorm:"column:inventory_quantity"`
	Length            float64              `json:"length" gorm:"column:length"`
	ManageInventory   bool                 `json:"manage_inventory" gorm:"column:manage_inventory;default:true"`
	Material          string               `json:"material" gorm:"column:material"`
	MIdCode           string               `json:"mid_code" gorm:"column:mid_code"`
	Options           []ProductOptionValue `json:"options" gorm:"foreignKey:Id"`
	OriginCountry     string               `json:"origin_country" gorm:"column:origin_country"`
	Prices            []MoneyAmount        `json:"prices" gorm:"many2many:product_variant_money_amount"`
	Product           *Product             `json:"product" gorm:"foreignKey:ProductId"`
	ProductId         uuid.NullUUID        `json:"product_id" gorm:"column:product_id"`
	Sku               string               `json:"sku" gorm:"column:sku"`
	Title             string               `json:"title" gorm:"column:title"`
	Upc               string               `json:"upc" gorm:"column:upc"`
	VariantRank       int                  `json:"variant_rank" gorm:"column:variant_rank;default:0"`
	Weight            float64              `json:"weight" gorm:"column:weight"`
	Width             float64              `json:"width" gorm:"column:width"`
}

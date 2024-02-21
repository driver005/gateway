package models

import (
	"database/sql/driver"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

//
// @oas:schema:Product
// title: "Product"
// description: "A product is a saleable item that holds general information such as name or description. It must include at least one Product Variant, where each product variant defines different options to purchase the product with (for example, different sizes or colors). The prices and inventory of the product are defined on the variant level."
// type: object
// required:
//   - collection_id
//   - created_at
//   - deleted_at
//   - description
//   - discountable
//   - external_id
//   - handle
//   - height
//   - hs_code
//   - id
//   - is_giftcard
//   - length
//   - material
//   - metadata
//   - mid_code
//   - origin_country
//   - profile_id
//   - status
//   - subtitle
//   - type_id
//   - thumbnail
//   - title
//   - updated_at
//   - weight
//   - width
// properties:
//   id:
//     description: The product's ID
//     type: string
//     example: prod_01G1G5V2MBA328390B5AXJ610F
//   title:
//     description: A title that can be displayed for easy identification of the Product.
//     type: string
//     example: Medusa Coffee Mug
//   subtitle:
//     description: An optional subtitle that can be used to further specify the Product.
//     nullable: true
//     type: string
//   description:
//     description: A short description of the Product.
//     nullable: true
//     type: string
//     example: Every programmer's best friend.
//   handle:
//     description: A unique identifier for the Product (e.g. for slug structure).
//     nullable: true
//     type: string
//     example: coffee-mug
//   is_giftcard:
//     description: Whether the Product represents a Gift Card. Products that represent Gift Cards will automatically generate a redeemable Gift Card code once they are purchased.
//     type: boolean
//     default: false
//   status:
//     description: The status of the product
//     type: string
//     enum:
//       - draft
//       - proposed
//       - published
//       - rejected
//     default: draft
//   images:
//     description: The details of the product's images.
//     type: array
//     x-expandable: "images"
//     items:
//       $ref: "#/components/schemas/Image"
//   thumbnail:
//     description: A URL to an image file that can be used to identify the Product.
//     nullable: true
//     type: string
//     format: uri
//   options:
//     description: The details of the Product Options that are defined for the Product. The product's variants will have a unique combination of values of the product's options.
//     type: array
//     x-expandable: "options"
//     items:
//       $ref: "#/components/schemas/ProductOption"
//   variants:
//     description: The details of the Product Variants that belong to the Product. Each will have a unique combination of values of the product's options.
//     type: array
//     x-expandable: "variants"
//     items:
//       $ref: "#/components/schemas/ProductVariant"
//   categories:
//     description: The details of the product categories that this product belongs to.
//     type: array
//     x-expandable: "categories"
//     x-featureFlag: "product_categories"
//     items:
//       $ref: "#/components/schemas/ProductCategory"
//   profile_id:
//     description: The ID of the shipping profile that the product belongs to. The shipping profile has a set of defined shipping options that can be used to fulfill the product.
//     type: string
//     example: sp_01G1G5V239ENSZ5MV4JAR737BM
//   profile:
//     description: The details of the shipping profile that the product belongs to. The shipping profile has a set of defined shipping options that can be used to fulfill the product.
//     x-expandable: "profile"
//     nullable: true
//     $ref: "#/components/schemas/ShippingProfile"
//   profiles:
//     description: Available if the relation `profiles` is expanded.
//     nullable: true
//     type: array
//     items:
//       $ref: "#/components/schemas/ShippingProfile"
//   weight:
//     description: The weight of the Product Variant. May be used in shipping rate calculations.
//     nullable: true
//     type: number
//     example: null
//   length:
//     description: The length of the Product Variant. May be used in shipping rate calculations.
//     nullable: true
//     type: number
//     example: null
//   height:
//     description: The height of the Product Variant. May be used in shipping rate calculations.
//     nullable: true
//     type: number
//     example: null
//   width:
//     description: The width of the Product Variant. May be used in shipping rate calculations.
//     nullable: true
//     type: number
//     example: null
//   hs_code:
//     description: The Harmonized System code of the Product Variant. May be used by Fulfillment Providers to pass customs information to shipping carriers.
//     nullable: true
//     type: string
//     example: null
//   origin_country:
//     description: The country in which the Product Variant was produced. May be used by Fulfillment Providers to pass customs information to shipping carriers.
//     nullable: true
//     type: string
//     example: null
//   mid_code:
//     description: The Manufacturers Identification code that identifies the manufacturer of the Product Variant. May be used by Fulfillment Providers to pass customs information to shipping carriers.
//     nullable: true
//     type: string
//     example: null
//   material:
//     description: The material and composition that the Product Variant is made of, May be used by Fulfillment Providers to pass customs information to shipping carriers.
//     nullable: true
//     type: string
//     example: null
//   collection_id:
//     description: The ID of the product collection that the product belongs to.
//     nullable: true
//     type: string
//     example: pcol_01F0YESBFAZ0DV6V831JXWH0BG
//   collection:
//     description: The details of the product collection that the product belongs to.
//     x-expandable: "collection"
//     nullable: true
//     $ref: "#/components/schemas/ProductCollection"
//   type_id:
//     description: The ID of the product type that the product belongs to.
//     nullable: true
//     type: string
//     example: ptyp_01G8X9A7ESKAJXG2H0E6F1MW7A
//   type:
//     description: The details of the product type that the product belongs to.
//     x-expandable: "type"
//     nullable: true
//     $ref: "#/components/schemas/ProductType"
//   tags:
//     description: The details of the product tags used in this product.
//     type: array
//     x-expandable: "type"
//     items:
//       $ref: "#/components/schemas/ProductTag"
//   discountable:
//     description: Whether the Product can be discounted. Discounts will not apply to Line Items of this Product when this flag is set to `false`.
//     type: boolean
//     default: true
//   external_id:
//     description: The external ID of the product
//     nullable: true
//     type: string
//     example: null
//   sales_channels:
//     description: The details of the sales channels this product is available in.
//     type: array
//     x-expandable: "sales_channels"
//     items:
//       $ref: "#/components/schemas/SalesChannel"
//   created_at:
//     description: The date with timezone at which the resource was created.
//     type: string
//     format: date-time
//   updated_at:
//     description: The date with timezone at which the resource was updated.
//     type: string
//     format: date-time
//   deleted_at:
//     description: The date with timezone at which the resource was deleted.
//     nullable: true
//     type: string
//     format: date-time
//   metadata:
//     description: An optional key-value map with additional details
//     nullable: true
//     type: object
//     example: {car: "white"}
//     externalDocs:
//       description: "Learn about the metadata attribute, and how to delete and update it."
//       url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
//

type Product struct {
	core.SoftDeletableModel

	Collection    *ProductCollection `json:"collection" gorm:"foreignKey:CollectionId"`
	CollectionId  uuid.NullUUID      `json:"collection_id" gorm:"column:collection_id"`
	Description   string             `json:"description" gorm:"column:description"`
	Discountable  bool               `json:"discountable" gorm:"column:discountable;default:false"`
	ExternalId    string             `json:"external_id" gorm:"column:external_id"`
	Handle        string             `json:"handle" gorm:"column:handle"`
	Height        float64            `json:"height" gorm:"column:height"`
	HsCode        string             `json:"hs_code" gorm:"column:hs_code"`
	Images        []Image            `json:"images" gorm:"many2many:product_images"`
	IsGiftcard    bool               `json:"is_giftcard" gorm:"column:is_giftcard;default:false"`
	Length        float64            `json:"length" gorm:"column:length"`
	Material      string             `json:"material" gorm:"column:material"`
	MIdCode       uuid.NullUUID      `json:"mid_code" gorm:"column:mid_code"`
	Options       []ProductOption    `json:"options" gorm:"foreignKey:Id"`
	OriginCountry string             `json:"origin_country" gorm:"column:origin_country"`
	Categories    []ProductCategory  `json:"categories" gorm:"many2many:product_category_product"`
	Profile       *ShippingProfile   `json:"profile" gorm:"foreignKey:ProfileId"`
	ProfileId     uuid.NullUUID      `json:"profile_id" gorm:"column:profile_id"`
	Profiles      []ShippingProfile  `json:"profiles" gorm:"many2many:product_shipping_profile"`
	SalesChannels []SalesChannel     `json:"sales_channels" gorm:"many2many:product_sales_channel"`
	Status        ProductStatus      `json:"status" gorm:"column:status;default:'draft'"`
	Subtitle      string             `json:"subtitle" gorm:"column:subtitle"`
	Tags          []ProductTag       `json:"tags" gorm:"many2many:product_tags"`
	Thumbnail     string             `json:"thumbnail" gorm:"column:thumbnail"`
	Title         string             `json:"title" gorm:"column:title"`
	Type          *ProductType       `json:"type" gorm:"foreignKey:TypeId"`
	TypeId        uuid.NullUUID      `json:"type_id" gorm:"column:type_id"`
	Variants      []ProductVariant   `json:"variants" gorm:"foreignKey:Id"`
	Weight        float64            `json:"weight" gorm:"column:weight"`
	Width         float64            `json:"width" gorm:"column:width"`
}

// The status of the product
type ProductStatus string

const (
	ProductStatusDraft     ProductStatus = "draft"
	ProductStatusProposed  ProductStatus = "proposed"
	ProductStatusPublished ProductStatus = "published"
	ProductStatusRejected  ProductStatus = "rejected"
)

func (ps *ProductStatus) Scan(value interface{}) error {
	*ps = ProductStatus(value.([]byte))
	return nil
}

func (ps ProductStatus) Value() (driver.Value, error) {
	return string(ps), nil
}

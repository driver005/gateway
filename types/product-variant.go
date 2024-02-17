package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type GetRegionPriceContext struct {
	RegionId              uuid.UUID `json:"regionId uuid.UUID"`
	Quantity              int       `json:"quantity,omitempty" validate:"omitempty"`
	CustomerId            uuid.UUID `json:"customer_id,omitempty" validate:"omitempty"`
	IncludeDiscountPrices bool      `json:"include_discount_prices,omitempty" validate:"omitempty"`
}

type ProductVariantOption struct {
	OptionId uuid.UUID `json:"option_id"`
	Value    string    `json:"value"`
}
type UpdateProductVariantData struct {
	Variant    *models.ProductVariant     `json:"variant"`
	UpdateData *UpdateProductVariantInput `json:"updateData"`
}

type UpdateVariantPricesData struct {
	VariantId uuid.UUID             `json:"variantId uuid.UUID"`
	Prices    []ProductVariantPrice `json:"prices"`
}

type UpdateVariantRegionPriceData struct {
	VariantId uuid.UUID            `json:"variantId uuid.UUID"`
	Price     *ProductVariantPrice `json:"price"`
}

type UpdateVariantCurrencyPriceData struct {
	VariantId uuid.UUID            `json:"variantId uuid.UUID"`
	Price     *ProductVariantPrice `json:"price"`
}

type ProductVariantPricesUpdateReq struct {
	Id           uuid.UUID `json:"id,omitempty" validate:"omitempty"`
	RegionId     uuid.UUID `json:"region_id,omitempty" validate:"omitempty"`
	CurrencyCode string    `json:"currency_code,omitempty" validate:"omitempty"`
	Amount       int       `json:"amount"`
	MinQuantity  int       `json:"min_quantity,omitempty" validate:"omitempty"`
	MaxQuantity  int       `json:"max_quantity,omitempty" validate:"omitempty"`
}

type ProductVariantPricesCreateReq struct {
	RegionId     uuid.UUID `json:"region_id,omitempty" validate:"omitempty"`
	CurrencyCode string    `json:"currency_code,omitempty" validate:"omitempty"`
	Amount       int       `json:"amount"`
	MinQuantity  int       `json:"min_quantity,omitempty" validate:"omitempty"`
	MaxQuantity  int       `json:"max_quantity,omitempty" validate:"omitempty"`
}

type FilterableProductVariant struct {
	core.FilterModel
	Title             []string         `json:"title,omitempty" validate:"omitempty"`
	ProductId         uuid.UUIDs       `json:"product_id,omitempty" validate:"omitempty"`
	Product           *models.Product  `json:"product,omitempty" validate:"omitempty"`
	SKU               []string         `json:"sku,omitempty" validate:"omitempty"`
	Barcode           []string         `json:"barcode,omitempty" validate:"omitempty"`
	EAN               []string         `json:"ean,omitempty" validate:"omitempty"`
	UPC               string           `json:"upc,omitempty" validate:"omitempty"`
	InventoryQuantity core.NumberModel `json:"inventory_quantity,omitempty" validate:"omitempty"`
	AllowBackorder    bool             `json:"allow_backorder,omitempty" validate:"omitempty"`
	ManageInventory   bool             `json:"manage_inventory,omitempty" validate:"omitempty"`
	HSCode            []string         `json:"hs_code,omitempty" validate:"omitempty"`
	OriginCountry     []string         `json:"origin_country,omitempty" validate:"omitempty"`
	MidCode           []string         `json:"mid_code,omitempty" validate:"omitempty"`
	Material          string           `json:"material,omitempty" validate:"omitempty"`
	Weight            core.NumberModel `json:"weight,omitempty" validate:"omitempty"`
	Length            core.NumberModel `json:"length,omitempty" validate:"omitempty"`
	Height            core.NumberModel `json:"height,omitempty" validate:"omitempty"`
	Width             core.NumberModel `json:"width,omitempty" validate:"omitempty"`
	Q                 string           `json:"q,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostProductsProductVariantsReq
// type: object
// description: "The details of the product variant to create."
// required:
//   - title
//   - prices
//   - options
//
// properties:
//
//	title:
//	  description: The title of the product variant.
//	  type: string
//	sku:
//	  description: The unique SKU of the product variant.
//	  type: string
//	ean:
//	  description: The EAN number of the product variant.
//	  type: string
//	upc:
//	  description: The UPC number of the product variant.
//	  type: string
//	barcode:
//	  description: A generic GTIN field of the product variant.
//	  type: string
//	hs_code:
//	  description: The Harmonized System code of the product variant.
//	  type: string
//	inventory_quantity:
//	  description: The amount of stock kept of the product variant.
//	  type: integer
//	  default: 0
//	allow_backorder:
//	  description: Whether the product variant can be purchased when out of stock.
//	  type: boolean
//	manage_inventory:
//	  description: Whether Medusa should keep track of the inventory of this product variant.
//	  type: boolean
//	  default: true
//	weight:
//	  description: The wieght of the product variant.
//	  type: number
//	length:
//	  description: The length of the product variant.
//	  type: number
//	height:
//	  description: The height of the product variant.
//	  type: number
//	width:
//	  description: The width of the product variant.
//	  type: number
//	origin_country:
//	  description: The country of origin of the product variant.
//	  type: string
//	mid_code:
//	  description: The Manufacturer Identification code of the product variant.
//	  type: string
//	material:
//	  description: The material composition of the product variant.
//	  type: string
//	metadata:
//	  description: An optional set of key-value pairs with additional information.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
//	prices:
//	  type: array
//	  description: An array of product variant prices. A product variant can have different prices for each region or currency code.
//	  externalDocs:
//	    url: https://docs.medusajs.com/modules/products/admin/manage-products#product-variant-prices
//	    description: Product variant pricing.
//	  items:
//	    type: object
//	    required:
//	      - amount
//	    properties:
//	      region_id:
//	        description: The ID of the Region the price will be used in. This is only required if `currency_code` is not provided.
//	        type: string
//	      currency_code:
//	        description: The 3 character ISO currency code the price will be used in. This is only required if `region_id` is not provided.
//	        type: string
//	        externalDocs:
//	          url: https://en.wikipedia.org/wiki/ISO_4217#Active_codes
//	          description: See a list of codes.
//	      amount:
//	        description: The price amount.
//	        type: integer
//	      min_quantity:
//	       description: The minimum quantity required to be added to the cart for the price to be used.
//	       type: integer
//	      max_quantity:
//	        description: The maximum quantity required to be added to the cart for the price to be used.
//	        type: integer
//	options:
//	  type: array
//	  description: An array of Product Option values that the variant corresponds to.
//	  items:
//	    type: object
//	    required:
//	      - option_id
//	      - value
//	    properties:
//	      option_id:
//	        description: The ID of the Product Option.
//	        type: string
//	      value:
//	        description: A value to give to the Product Option.
//	        type: string
type CreateProductVariantInput struct {
	Title             string                 `json:"title,omitempty" validate:"omitempty"`
	ProductId         uuid.UUID              `json:"product_id,omitempty" validate:"omitempty"`
	SKU               string                 `json:"sku,omitempty" validate:"omitempty"`
	Barcode           string                 `json:"barcode,omitempty" validate:"omitempty"`
	EAN               string                 `json:"ean,omitempty" validate:"omitempty"`
	UPC               string                 `json:"upc,omitempty" validate:"omitempty"`
	VariantRank       int                    `json:"variant_rank,omitempty" validate:"omitempty"`
	InventoryQuantity int                    `json:"inventory_quantity,omitempty" validate:"omitempty"`
	AllowBackorder    bool                   `json:"allow_backorder,omitempty" validate:"omitempty"`
	ManageInventory   bool                   `json:"manage_inventory,omitempty" validate:"omitempty"`
	HSCode            string                 `json:"hs_code,omitempty" validate:"omitempty"`
	OriginCountry     string                 `json:"origin_country,omitempty" validate:"omitempty"`
	MIdCode           string                 `json:"mid_code,omitempty" validate:"omitempty"`
	Material          string                 `json:"material,omitempty" validate:"omitempty"`
	Weight            float64                `json:"weight,omitempty" validate:"omitempty"`
	Length            float64                `json:"length,omitempty" validate:"omitempty"`
	Height            float64                `json:"height,omitempty" validate:"omitempty"`
	Width             float64                `json:"width,omitempty" validate:"omitempty"`
	Options           []ProductVariantOption `json:"options"`
	Prices            []ProductVariantPrice  `json:"prices"`
	Metadata          core.JSONB             `json:"metadata,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostProductsProductVariantsVariantReq
// type: object
// properties:
//
//	title:
//	  description: The title of the product variant.
//	  type: string
//	sku:
//	  description: The unique SKU of the product variant.
//	  type: string
//	ean:
//	  description: The EAN number of the item.
//	  type: string
//	upc:
//	  description: The UPC number of the item.
//	  type: string
//	barcode:
//	  description: A generic GTIN field of the product variant.
//	  type: string
//	hs_code:
//	  description: The Harmonized System code of the product variant.
//	  type: string
//	inventory_quantity:
//	  description: The amount of stock kept of the product variant.
//	  type: integer
//	allow_backorder:
//	  description: Whether the product variant can be purchased when out of stock.
//	  type: boolean
//	manage_inventory:
//	  description: Whether Medusa should keep track of the inventory of this product variant.
//	  type: boolean
//	weight:
//	  description: The weight of the product variant.
//	  type: number
//	length:
//	  description: The length of the product variant.
//	  type: number
//	height:
//	  description: The height of the product variant.
//	  type: number
//	width:
//	  description: The width of the product variant.
//	  type: number
//	origin_country:
//	  description: The country of origin of the product variant.
//	  type: string
//	mid_code:
//	  description: The Manufacturer Identification code of the product variant.
//	  type: string
//	material:
//	  description: The material composition of the product variant.
//	  type: string
//	metadata:
//	  description: An optional set of key-value pairs with additional information.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
//	prices:
//	  type: array
//	  description: An array of product variant prices. A product variant can have different prices for each region or currency code.
//	  externalDocs:
//	    url: https://docs.medusajs.com/modules/products/admin/manage-products#product-variant-prices
//	    description: Product variant pricing.
//	  items:
//	    type: object
//	    required:
//	      - amount
//	    properties:
//	      id:
//	        description: The ID of the price. If provided, the existing price will be updated. Otherwise, a new price will be created.
//	        type: string
//	      region_id:
//	        description: The ID of the Region the price will be used in. This is only required if `currency_code` is not provided.
//	        type: string
//	      currency_code:
//	        description: The 3 character ISO currency code the price will be used in. This is only required if `region_id` is not provided.
//	        type: string
//	        externalDocs:
//	          url: https://en.wikipedia.org/wiki/ISO_4217#Active_codes
//	          description: See a list of codes.
//	      amount:
//	        description: The price amount.
//	        type: integer
//	      min_quantity:
//	       description: The minimum quantity required to be added to the cart for the price to be used.
//	       type: integer
//	      max_quantity:
//	        description: The maximum quantity required to be added to the cart for the price to be used.
//	        type: integer
//	options:
//	  type: array
//	  description: An array of Product Option values that the variant corresponds to.
//	  items:
//	    type: object
//	    required:
//	      - option_id
//	      - value
//	    properties:
//	      option_id:
//	        description: The ID of the Product Option.
//	        type: string
//	      value:
//	        description: The value of the Product Option.
//	        type: string
type UpdateProductVariantInput struct {
	Title             string                 `json:"title,omitempty" validate:"omitempty"`
	ProductId         uuid.UUID              `json:"product_id,omitempty" validate:"omitempty"`
	SKU               string                 `json:"sku,omitempty" validate:"omitempty"`
	Barcode           string                 `json:"barcode,omitempty" validate:"omitempty"`
	EAN               string                 `json:"ean,omitempty" validate:"omitempty"`
	UPC               string                 `json:"upc,omitempty" validate:"omitempty"`
	InventoryQuantity int                    `json:"inventory_quantity,omitempty" validate:"omitempty"`
	AllowBackorder    bool                   `json:"allow_backorder,omitempty" validate:"omitempty"`
	ManageInventory   bool                   `json:"manage_inventory,omitempty" validate:"omitempty"`
	HSCode            string                 `json:"hs_code,omitempty" validate:"omitempty"`
	OriginCountry     string                 `json:"origin_country,omitempty" validate:"omitempty"`
	VariantRank       int                    `json:"variant_rank,omitempty" validate:"omitempty"`
	MIdCode           string                 `json:"mid_code,omitempty" validate:"omitempty"`
	Material          string                 `json:"material,omitempty" validate:"omitempty"`
	Weight            float64                `json:"weight,omitempty" validate:"omitempty"`
	Length            float64                `json:"length,omitempty" validate:"omitempty"`
	Height            float64                `json:"height,omitempty" validate:"omitempty"`
	Width             float64                `json:"width,omitempty" validate:"omitempty"`
	Options           []ProductVariantOption `json:"options,omitempty" validate:"omitempty"`
	Prices            []ProductVariantPrice  `json:"prices,omitempty" validate:"omitempty"`
	Metadata          core.JSONB             `json:"metadata,omitempty" validate:"omitempty"`
}

type ProductVariantPrice struct {
	Id           uuid.UUID `json:"id,omitempty" validate:"omitempty"`
	CurrencyCode string    `json:"currency_code,omitempty" validate:"omitempty"`
	RegionId     uuid.UUID `json:"region_id,omitempty" validate:"omitempty"`
	Amount       float64   `json:"amount"`
	MinQuantity  int       `json:"min_quantity,omitempty" validate:"omitempty"`
	MaxQuantity  int       `json:"max_quantity,omitempty" validate:"omitempty"`
}

type ProductVariantVariant struct {
	PriceSelectionParams
	SalesChannelId uuid.UUID `json:"sales_channel_id,omitempty" validate:"omitempty"`
}

type ProductVariantParams struct {
	PriceSelectionParams
	Ids               uuid.UUIDs       `json:"ids,omitempty" validate:"omitempty,alphanum"`
	SalesChannelId    uuid.UUID        `json:"sales_channel_id,omitempty" validate:"omitempty,alphanum"`
	Id                uuid.UUID        `json:"id,omitempty" validate:"omitempty"`
	Title             []string         `json:"title,omitempty" validate:"omitempty"`
	InventoryQuantity core.NumberModel `json:"inventory_quantity,omitempty" validate:"omitempty"`
}

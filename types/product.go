package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type FilterableProduct struct {
	core.FilterModel
	PriceSelectionParams
	Q                       string                 `json:"q,omitempty" validate:"omitempty"`
	Status                  []models.ProductStatus `json:"status,omitempty" validate:"omitempty"`
	PriceListId             uuid.UUIDs             `json:"price_list_id,omitempty" validate:"omitempty"`
	CollectionId            uuid.UUIDs             `json:"collection_id,omitempty" validate:"omitempty"`
	Tags                    []string               `json:"tags,omitempty" validate:"omitempty"`
	Title                   string                 `json:"title,omitempty" validate:"omitempty"`
	Description             string                 `json:"description,omitempty" validate:"omitempty"`
	Handle                  string                 `json:"handle,omitempty" validate:"omitempty"`
	IsGiftcard              bool                   `json:"is_giftcard,omitempty" validate:"omitempty"`
	TypeId                  uuid.UUIDs             `json:"type_id,omitempty" validate:"omitempty"`
	SalesChannelId          uuid.UUIDs             `json:"sales_channel_id,omitempty" validate:"omitempty"`
	DiscountConditionId     uuid.UUID              `json:"discount_condition_id,omitempty" validate:"omitempty"`
	CategoryId              uuid.UUIDs             `json:"category_id,omitempty" validate:"omitempty"`
	IncludeCategoryChildren bool                   `json:"include_category_children,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostProductsReq
// type: object
// description: "The details of the product to create."
// required:
//   - title
//
// properties:
//
//	title:
//	  description: "The title of the Product"
//	  type: string
//	subtitle:
//	  description: "The subtitle of the Product"
//	  type: string
//	description:
//	  description: "The description of the Product."
//	  type: string
//	is_giftcard:
//	  description: A flag to indicate if the Product represents a Gift Card. Purchasing Products with this flag set to `true` will result in a Gift Card being created.
//	  type: boolean
//	  default: false
//	discountable:
//	  description: A flag to indicate if discounts can be applied to the Line Items generated from this Product
//	  type: boolean
//	  default: true
//	images:
//	  description: An array of images of the Product. Each value in the array is a URL to the image. You can use the upload API Routes to upload the image and obtain a URL.
//	  type: array
//	  items:
//	    type: string
//	thumbnail:
//	  description: The thumbnail to use for the Product. The value is a URL to the thumbnail. You can use the upload API Routes to upload the thumbnail and obtain a URL.
//	  type: string
//	handle:
//	  description: A unique handle to identify the Product by. If not provided, the kebab-case version of the product title will be used. This can be used as a slug in URLs.
//	  type: string
//	status:
//	  description: The status of the product. The product is shown to the customer only if its status is `published`.
//	  type: string
//	  enum: [draft, proposed, published, rejected]
//	  default: draft
//	type:
//	  description: The Product Type to associate the Product with.
//	  type: object
//	  required:
//	    - value
//	  properties:
//	    id:
//	      description: The ID of an existing Product Type. If not provided, a new product type will be created.
//	      type: string
//	    value:
//	      description: The value of the Product Type.
//	      type: string
//	collection_id:
//	  description: The ID of the Product Collection the Product belongs to.
//	  type: string
//	tags:
//	  description: Product Tags to associate the Product with.
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - value
//	    properties:
//	      id:
//	        description: The ID of an existing Product Tag. If not provided, a new product tag will be created.
//	        type: string
//	      value:
//	        description: The value of the Tag. If the `id` is provided, the value of the existing tag will be updated.
//	        type: string
//	sales_channels:
//	  description: "Sales channels to associate the Product with."
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - id
//	    properties:
//	      id:
//	        description: The ID of an existing Sales channel.
//	        type: string
//	categories:
//	  description: "Product categories to add the Product to."
//	  x-featureFlag: "product_categories"
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - id
//	    properties:
//	      id:
//	        description: The ID of a Product Category.
//	        type: string
//	options:
//	  description: The Options that the Product should have. A new product option will be created for every item in the array.
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - title
//	    properties:
//	      title:
//	        description: The title of the Product Option.
//	        type: string
//	variants:
//	  description: An array of Product Variants to create with the Product. Each product variant must have a unique combination of Product Option values.
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - title
//	    properties:
//	      title:
//	        description: The title of the Product Variant.
//	        type: string
//	      sku:
//	        description: The unique SKU of the Product Variant.
//	        type: string
//	      ean:
//	        description: The EAN number of the item.
//	        type: string
//	      upc:
//	        description: The UPC number of the item.
//	        type: string
//	      barcode:
//	        description: A generic GTIN field of the Product Variant.
//	        type: string
//	      hs_code:
//	        description: The Harmonized System code of the Product Variant.
//	        type: string
//	      inventory_quantity:
//	        description: The amount of stock kept of the Product Variant.
//	        type: integer
//	        default: 0
//	      allow_backorder:
//	        description: Whether the Product Variant can be purchased when out of stock.
//	        type: boolean
//	      manage_inventory:
//	        description: Whether Medusa should keep track of the inventory of this Product Variant.
//	        type: boolean
//	      weight:
//	        description: The wieght of the Product Variant.
//	        type: number
//	      length:
//	        description: The length of the Product Variant.
//	        type: number
//	      height:
//	        description: The height of the Product Variant.
//	        type: number
//	      width:
//	        description: The width of the Product Variant.
//	        type: number
//	      origin_country:
//	        description: The country of origin of the Product Variant.
//	        type: string
//	      mid_code:
//	        description: The Manufacturer Identification code of the Product Variant.
//	        type: string
//	      material:
//	        description: The material composition of the Product Variant.
//	        type: string
//	      metadata:
//	        description: An optional set of key-value pairs with additional information.
//	        type: object
//	        externalDocs:
//	          description: "Learn about the metadata attribute, and how to delete and update it."
//	          url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
//	      prices:
//	        type: array
//	        description: An array of product variant prices. A product variant can have different prices for each region or currency code.
//	        externalDocs:
//	          url: https://docs.medusajs.com/modules/products/admin/manage-products#product-variant-prices
//	          description: Product variant pricing.
//	        items:
//	          type: object
//	          required:
//	            - amount
//	          properties:
//	            region_id:
//	              description: The ID of the Region the price will be used in. This is only required if `currency_code` is not provided.
//	              type: string
//	            currency_code:
//	              description: The 3 character ISO currency code the price will be used in. This is only required if `region_id` is not provided.
//	              type: string
//	              externalDocs:
//	                url: https://en.wikipedia.org/wiki/ISO_4217#Active_codes
//	                description: See a list of codes.
//	            amount:
//	              description: The price amount.
//	              type: integer
//	            min_quantity:
//	              description: The minimum quantity required to be added to the cart for the price to be used.
//	              type: integer
//	            max_quantity:
//	              description: The maximum quantity required to be added to the cart for the price to be used.
//	              type: integer
//	      options:
//	        type: array
//	        description: An array of Product Option values that the variant corresponds to. The option values should be added into the array in the same index as in the `options` field of the product.
//	        externalDocs:
//	          url: https://docs.medusajs.com/modules/products/admin/manage-products#create-a-product
//	          description: Example of how to create a product with options and variants
//	        items:
//	          type: object
//	          required:
//	            - value
//	          properties:
//	            value:
//	              description: The value to give for the Product Option at the same index in the Product's `options` field.
//	              type: string
//	weight:
//	  description: The weight of the Product.
//	  type: number
//	length:
//	  description: The length of the Product.
//	  type: number
//	height:
//	  description: The height of the Product.
//	  type: number
//	width:
//	  description: The width of the Product.
//	  type: number
//	hs_code:
//	  description: The Harmonized System code of the Product.
//	  type: string
//	origin_country:
//	  description: The country of origin of the Product.
//	  type: string
//	mid_code:
//	  description: The Manufacturer Identification code of the Product.
//	  type: string
//	material:
//	  description: The material composition of the Product.
//	  type: string
//	metadata:
//	  description: An optional set of key-value pairs with additional information.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type CreateProductInput struct {
	Title         string                                  `json:"title"`
	Subtitle      string                                  `json:"subtitle,omitempty" validate:"omitempty"`
	ProfileId     uuid.UUID                               `json:"profile_id,omitempty" validate:"omitempty"`
	Description   string                                  `json:"description,omitempty" validate:"omitempty"`
	IsGiftcard    bool                                    `json:"is_giftcard,omitempty" validate:"omitempty"`
	Discountable  bool                                    `json:"discountable,omitempty" validate:"omitempty"`
	Images        []string                                `json:"images,omitempty" validate:"omitempty"`
	Thumbnail     string                                  `json:"thumbnail,omitempty" validate:"omitempty"`
	Handle        string                                  `json:"handle,omitempty" validate:"omitempty"`
	Status        models.ProductStatus                    `json:"status,omitempty" validate:"omitempty"`
	Type          *CreateProductProductTypeInput          `json:"type,omitempty" validate:"omitempty"`
	CollectionId  uuid.UUID                               `json:"collection_id,omitempty" validate:"omitempty"`
	Tags          []CreateProductProductTagInput          `json:"tags,omitempty" validate:"omitempty"`
	Options       []CreateProductProductOption            `json:"options,omitempty" validate:"omitempty"`
	Variants      []CreateProductVariantInput             `json:"variants,omitempty" validate:"omitempty"`
	SalesChannels []CreateProductProductSalesChannelInput `json:"sales_channels,omitempty" validate:"omitempty"`
	Categories    []CreateProductProductCategoryInput     `json:"categories,omitempty" validate:"omitempty"`
	Weight        int                                     `json:"weight,omitempty" validate:"omitempty"`
	Length        int                                     `json:"length,omitempty" validate:"omitempty"`
	Height        int                                     `json:"height,omitempty" validate:"omitempty"`
	Width         int                                     `json:"width,omitempty" validate:"omitempty"`
	HSCode        string                                  `json:"hs_code,omitempty" validate:"omitempty"`
	OriginCountry string                                  `json:"origin_country,omitempty" validate:"omitempty"`
	MIdCode       uuid.UUID                               `json:"mid_code,omitempty" validate:"omitempty"`
	Material      string                                  `json:"material,omitempty" validate:"omitempty"`
	Metadata      core.JSONB                              `json:"metadata,omitempty" validate:"omitempty"`
	ExternalId    uuid.UUID                               `json:"external_id,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostProductsProductReq
// type: object
// description: "The details to update of the product."
// properties:
//
//	title:
//	  description: "The title of the Product"
//	  type: string
//	subtitle:
//	  description: "The subtitle of the Product"
//	  type: string
//	description:
//	  description: "The description of the Product."
//	  type: string
//	discountable:
//	  description: A flag to indicate if discounts can be applied to the Line Items generated from this Product
//	  type: boolean
//	images:
//	  description: An array of images of the Product. Each value in the array is a URL to the image. You can use the upload API Routes to upload the image and obtain a URL.
//	  type: array
//	  items:
//	    type: string
//	thumbnail:
//	  description: The thumbnail to use for the Product. The value is a URL to the thumbnail. You can use the upload API Routes to upload the thumbnail and obtain a URL.
//	  type: string
//	handle:
//	  description: A unique handle to identify the Product by. If not provided, the kebab-case version of the product title will be used. This can be used as a slug in URLs.
//	  type: string
//	status:
//	  description: The status of the product. The product is shown to the customer only if its status is `published`.
//	  type: string
//	  enum: [draft, proposed, published, rejected]
//	type:
//	  description: The Product Type to associate the Product with.
//	  type: object
//	  required:
//	    - value
//	  properties:
//	    id:
//	      description: The ID of an existing Product Type. If not provided, a new product type will be created.
//	      type: string
//	    value:
//	      description: The value of the Product Type.
//	      type: string
//	collection_id:
//	  description: The ID of the Product Collection the Product belongs to.
//	  type: string
//	tags:
//	  description: Product Tags to associate the Product with.
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - value
//	    properties:
//	      id:
//	        description: The ID of an existing Product Tag. If not provided, a new product tag will be created.
//	        type: string
//	      value:
//	        description: The value of the Tag. If the `id` is provided, the value of the existing tag will be updated.
//	        type: string
//	sales_channels:
//	  description: "Sales channels to associate the Product with."
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - id
//	    properties:
//	      id:
//	        description: The ID of an existing Sales channel.
//	        type: string
//	categories:
//	  description: "Product categories to add the Product to."
//	  x-featureFlag: "product_categories"
//	  type: array
//	  items:
//	    required:
//	      - id
//	    properties:
//	      id:
//	        description: The ID of a Product Category.
//	        type: string
//	variants:
//	  description: An array of Product Variants to create with the Product. Each product variant must have a unique combination of Product Option values.
//	  type: array
//	  items:
//	    type: object
//	    properties:
//	      id:
//	        description: The id of an existing product variant. If provided, the details of the product variant will be updated. If not, a new product variant will be created.
//	        type: string
//	      title:
//	        description: The title of the product variant.
//	        type: string
//	      sku:
//	        description: The unique SKU of the product variant.
//	        type: string
//	      ean:
//	        description: The EAN number of the product variant.
//	        type: string
//	      upc:
//	        description: The UPC number of the product variant.
//	        type: string
//	      barcode:
//	        description: A generic GTIN field of the product variant.
//	        type: string
//	      hs_code:
//	        description: The Harmonized System code of the product variant.
//	        type: string
//	      inventory_quantity:
//	        description: The amount of stock kept of the product variant.
//	        type: integer
//	      allow_backorder:
//	        description: Whether the product variant can be purchased when out of stock.
//	        type: boolean
//	      manage_inventory:
//	        description: Whether Medusa should keep track of the inventory of this product variant.
//	        type: boolean
//	      weight:
//	        description: The weight of the product variant.
//	        type: number
//	      length:
//	        description: The length of the product variant.
//	        type: number
//	      height:
//	        description: The height of the product variant.
//	        type: number
//	      width:
//	        description: The width of the product variant.
//	        type: number
//	      origin_country:
//	        description: The country of origin of the product variant.
//	        type: string
//	      mid_code:
//	        description: The Manufacturer Identification code of the product variant.
//	        type: string
//	      material:
//	        description: The material composition of the product variant.
//	        type: string
//	      metadata:
//	        description: An optional set of key-value pairs with additional information.
//	        type: object
//	        externalDocs:
//	          description: "Learn about the metadata attribute, and how to delete and update it."
//	          url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
//	      prices:
//	        type: array
//	        description: An array of product variant prices. A product variant can have different prices for each region or currency code.
//	        externalDocs:
//	          url: https://docs.medusajs.com/modules/products/admin/manage-products#product-variant-prices
//	          description: Product variant pricing.
//	        items:
//	          type: object
//	          required:
//	            - amount
//	          properties:
//	            id:
//	              description: The ID of the Price. If provided, the existing price will be updated. Otherwise, a new price will be created.
//	              type: string
//	            region_id:
//	              description: The ID of the Region the price will be used in. This is only required if `currency_code` is not provided.
//	              type: string
//	            currency_code:
//	              description: The 3 character ISO currency code the price will be used in. This is only required if `region_id` is not provided.
//	              type: string
//	              externalDocs:
//	                url: https://en.wikipedia.org/wiki/ISO_4217#Active_codes
//	                description: See a list of codes.
//	            amount:
//	              description: The price amount.
//	              type: integer
//	            min_quantity:
//	              description: The minimum quantity required to be added to the cart for the price to be used.
//	              type: integer
//	            max_quantity:
//	              description: The maximum quantity required to be added to the cart for the price to be used.
//	              type: integer
//	      options:
//	        type: array
//	        description: An array of Product Option values that the variant corresponds to.
//	        items:
//	          type: object
//	          required:
//	            - option_id
//	            - value
//	          properties:
//	            option_id:
//	              description: The ID of the Option.
//	              type: string
//	            value:
//	              description: The value of the Product Option.
//	              type: string
//	weight:
//	  description: The weight of the Product.
//	  type: number
//	length:
//	  description: The length of the Product.
//	  type: number
//	height:
//	  description: The height of the Product.
//	  type: number
//	width:
//	  description: The width of the Product.
//	  type: number
//	hs_code:
//	  description: The Harmonized System code of the product variant.
//	  type: string
//	origin_country:
//	  description: The country of origin of the Product.
//	  type: string
//	mid_code:
//	  description: The Manufacturer Identification code of the Product.
//	  type: string
//	material:
//	  description: The material composition of the Product.
//	  type: string
//	metadata:
//	  description: An optional set of key-value pairs with additional information.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type UpdateProductInput struct {
	Title         string                                  `json:"title,omitempty" validate:"omitempty"`
	Subtitle      string                                  `json:"subtitle,omitempty" validate:"omitempty"`
	ProfileId     uuid.UUID                               `json:"profile_id,omitempty" validate:"omitempty"`
	Description   string                                  `json:"description,omitempty" validate:"omitempty"`
	IsGiftcard    bool                                    `json:"is_giftcard,omitempty" validate:"omitempty"`
	Discountable  bool                                    `json:"discountable,omitempty" validate:"omitempty"`
	Images        []string                                `json:"images,omitempty" validate:"omitempty"`
	Thumbnail     string                                  `json:"thumbnail,omitempty" validate:"omitempty"`
	Handle        string                                  `json:"handle,omitempty" validate:"omitempty"`
	Status        models.ProductStatus                    `json:"status,omitempty" validate:"omitempty"`
	Type          *CreateProductProductTypeInput          `json:"type,omitempty" validate:"omitempty"`
	CollectionId  uuid.UUID                               `json:"collection_id,omitempty" validate:"omitempty"`
	Tags          []CreateProductProductTagInput          `json:"tags,omitempty" validate:"omitempty"`
	Options       []CreateProductProductOption            `json:"options,omitempty" validate:"omitempty"`
	SalesChannels []CreateProductProductSalesChannelInput `json:"sales_channels,omitempty" validate:"omitempty"`
	Categories    []CreateProductProductCategoryInput     `json:"categories,omitempty" validate:"omitempty"`
	Weight        float64                                 `json:"weight,omitempty" validate:"omitempty"`
	Length        float64                                 `json:"length,omitempty" validate:"omitempty"`
	Height        float64                                 `json:"height,omitempty" validate:"omitempty"`
	Width         float64                                 `json:"width,omitempty" validate:"omitempty"`
	HSCode        string                                  `json:"hs_code,omitempty" validate:"omitempty"`
	OriginCountry string                                  `json:"origin_country,omitempty" validate:"omitempty"`
	MIdCode       uuid.UUID                               `json:"mid_code,omitempty" validate:"omitempty"`
	Material      string                                  `json:"material,omitempty" validate:"omitempty"`
	Metadata      core.JSONB                              `json:"metadata,omitempty" validate:"omitempty"`
	ExternalId    string                                  `json:"external_id,omitempty" validate:"omitempty"`
	Variants      []UpdateProductProductVariantDTO        `json:"variants,omitempty" validate:"omitempty"`
}

type CreateProductProductTagInput struct {
	Id    uuid.UUID `json:"id,omitempty" validate:"omitempty"`
	Value string    `json:"value"`
}

type CreateProductProductSalesChannelInput struct {
	Id uuid.UUID `json:"id"`
}

type CreateProductProductCategoryInput struct {
	Id uuid.UUID `json:"id"`
}

type CreateProductProductTypeInput struct {
	Id    uuid.UUID `json:"id,omitempty" validate:"omitempty"`
	Value string    `json:"value"`
}

type CreateProductProductVariantInput struct {
	Title             string                                  `json:"title"`
	SKU               string                                  `json:"sku,omitempty" validate:"omitempty"`
	EAN               string                                  `json:"ean,omitempty" validate:"omitempty"`
	UPC               string                                  `json:"upc,omitempty" validate:"omitempty"`
	Barcode           string                                  `json:"barcode,omitempty" validate:"omitempty"`
	HSCode            string                                  `json:"hs_code,omitempty" validate:"omitempty"`
	InventoryQuantity int                                     `json:"inventory_quantity,omitempty" validate:"omitempty"`
	AllowBackorder    bool                                    `json:"allow_backorder,omitempty" validate:"omitempty"`
	ManageInventory   bool                                    `json:"manage_inventory,omitempty" validate:"omitempty"`
	Weight            int                                     `json:"weight,omitempty" validate:"omitempty"`
	Length            int                                     `json:"length,omitempty" validate:"omitempty"`
	Height            int                                     `json:"height,omitempty" validate:"omitempty"`
	Width             int                                     `json:"width,omitempty" validate:"omitempty"`
	OriginCountry     string                                  `json:"origin_country,omitempty" validate:"omitempty"`
	MIdCode           uuid.UUID                               `json:"mid_code,omitempty" validate:"omitempty"`
	Material          string                                  `json:"material,omitempty" validate:"omitempty"`
	Metadata          core.JSONB                              `json:"metadata,omitempty" validate:"omitempty"`
	Prices            []CreateProductProductVariantPriceInput `json:"prices,omitempty" validate:"omitempty"`
	Options           []ProductVariantOption                  `json:"options,omitempty" validate:"omitempty"`
}

type UpdateProductProductVariantDTO struct {
	Id                uuid.UUID                               `json:"id,omitempty" validate:"omitempty"`
	Title             string                                  `json:"title,omitempty" validate:"omitempty"`
	SKU               string                                  `json:"sku,omitempty" validate:"omitempty"`
	EAN               string                                  `json:"ean,omitempty" validate:"omitempty"`
	UPC               string                                  `json:"upc,omitempty" validate:"omitempty"`
	Barcode           string                                  `json:"barcode,omitempty" validate:"omitempty"`
	HSCode            string                                  `json:"hs_code,omitempty" validate:"omitempty"`
	InventoryQuantity int                                     `json:"inventory_quantity,omitempty" validate:"omitempty"`
	AllowBackorder    bool                                    `json:"allow_backorder,omitempty" validate:"omitempty"`
	ManageInventory   bool                                    `json:"manage_inventory,omitempty" validate:"omitempty"`
	Weight            int                                     `json:"weight,omitempty" validate:"omitempty"`
	Length            int                                     `json:"length,omitempty" validate:"omitempty"`
	Height            int                                     `json:"height,omitempty" validate:"omitempty"`
	Width             int                                     `json:"width,omitempty" validate:"omitempty"`
	OriginCountry     string                                  `json:"origin_country,omitempty" validate:"omitempty"`
	MIdCode           uuid.UUID                               `json:"mid_code,omitempty" validate:"omitempty"`
	Material          string                                  `json:"material,omitempty" validate:"omitempty"`
	Metadata          core.JSONB                              `json:"metadata,omitempty" validate:"omitempty"`
	Prices            []CreateProductProductVariantPriceInput `json:"prices,omitempty" validate:"omitempty"`
	Options           []ProductVariantOption                  `json:"options,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostProductsProductOptionsReq
// type: object
// description: "The details of the product option to create."
// required:
//   - title
//
// properties:
//
//	title:
//	  description: "The title the Product Option."
//	  type: string
//	  example: "Size"
type CreateProductProductOption struct {
	Title string `json:"title"`
}

type CreateProductProductVariantPriceInput struct {
	RegionId     uuid.UUID `json:"region_id,omitempty" validate:"omitempty"`
	CurrencyCode string    `json:"currency_code,omitempty" validate:"omitempty"`
	Amount       float64   `json:"amount"`
	MinQuantity  int       `json:"min_quantity,omitempty" validate:"omitempty"`
	MaxQuantity  int       `json:"max_quantity,omitempty" validate:"omitempty"`
}

type ProductOptionInput struct {
	Title  string                      `json:"title"`
	Values []models.ProductOptionValue `json:"values,omitempty" validate:"omitempty"`
}

type ProductSalesChannelReq struct {
	Id uuid.UUID `json:"id"`
}

type ProductProductCategoryReq struct {
	Id uuid.UUID `json:"id"`
}

type ProductTagReq struct {
	Id    uuid.UUID `json:"id,omitempty" validate:"omitempty"`
	Value string    `json:"value"`
}

type ProductTypeReq struct {
	Id    uuid.UUID `json:"id,omitempty" validate:"omitempty"`
	Value string    `json:"value"`
}

// @oas:schema:StorePostSearchReq
// type: object
// properties:
//
//	q:
//	  type: string
//	  description: The search query.
//	offset:
//	  type: number
//	  description: The number of products to skip when retrieving the products.
//	limit:
//	  type: number
//	  description: Limit the number of products returned.
//	filter:
//	  description: Pass filters based on the search service.
type ProductSearch struct {
	Q      string      `json:"q,omitempty" validate:"omitempty"`
	Offset int         `json:"offset,omitempty" validate:"omitempty"`
	Limit  int         `json:"limit,omitempty" validate:"omitempty"`
	Filter interface{} `json:"filter,omitempty" validate:"omitempty"`
}

// type ProductFilterOptions struct {
// 	PriceListId uuid.UUID             FindOperatorPriceList       `json:"price_list_id,omitempty" validate:"omitempty"`
// 	SalesChannelId uuid.UUID          FindOperatorSalesChannel    `json:"sales_channel_id,omitempty" validate:"omitempty"`
// 	CategoryId uuid.UUID              FindOperatorProductCategory `json:"category_id,omitempty" validate:"omitempty"`
// 	IncludeCategoryChildren bool                        `json:"include_category_children,omitempty" validate:"omitempty"`
// 	DiscountConditionId uuid.UUID                      `json:"discount_condition_id,omitempty" validate:"omitempty"`
// }

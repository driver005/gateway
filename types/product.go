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

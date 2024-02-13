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

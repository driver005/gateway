package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type GetRegionPriceContext struct {
	RegionId              string `json:"regionId"`
	Quantity              int    `json:"quantity,omitempty"`
	CustomerId            string `json:"customer_id,omitempty"`
	IncludeDiscountPrices bool   `json:"include_discount_prices,omitempty"`
}

type ProductVariantOption struct {
	OptionId string `json:"option_id"`
	Value    string `json:"value"`
}
type UpdateProductVariantData struct {
	Variant    models.ProductVariant `json:"variant"`
	UpdateData models.ProductVariant `json:"updateData"`
}

type UpdateVariantPricesData struct {
	VariantId string               `json:"variantId"`
	Prices    []models.MoneyAmount `json:"prices"`
}

type UpdateVariantRegionPriceData struct {
	VariantId string             `json:"variantId"`
	Price     models.MoneyAmount `json:"price"`
}

type UpdateVariantCurrencyPriceData struct {
	VariantId string             `json:"variantId"`
	Price     models.MoneyAmount `json:"price"`
}

type ProductVariantPricesUpdateReq struct {
	ID           string `json:"id,omitempty"`
	RegionId     string `json:"region_id,omitempty"`
	CurrencyCode string `json:"currency_code,omitempty"`
	Amount       int    `json:"amount"`
	MinQuantity  int    `json:"min_quantity,omitempty"`
	MaxQuantity  int    `json:"max_quantity,omitempty"`
}

type ProductVariantPricesCreateReq struct {
	RegionId     string `json:"region_id,omitempty"`
	CurrencyCode string `json:"currency_code,omitempty"`
	Amount       int    `json:"amount"`
	MinQuantity  int    `json:"min_quantity,omitempty"`
	MaxQuantity  int    `json:"max_quantity,omitempty"`
}

type FilterableProductVariant struct {
	core.FilterModel
	Title             []string         `json:"title,omitempty"`
	ProductId         uuid.UUIDs       `json:"product_id,omitempty"`
	Product           models.Product   `json:"product,omitempty"`
	SKU               []string         `json:"sku,omitempty"`
	Barcode           []string         `json:"barcode,omitempty"`
	EAN               []string         `json:"ean,omitempty"`
	UPC               string           `json:"upc,omitempty"`
	InventoryQuantity core.NumberModel `json:"inventory_quantity,omitempty"`
	AllowBackorder    bool             `json:"allow_backorder,omitempty"`
	ManageInventory   bool             `json:"manage_inventory,omitempty"`
	HSCode            []string         `json:"hs_code,omitempty"`
	OriginCountry     []string         `json:"origin_country,omitempty"`
	MidCode           []string         `json:"mid_code,omitempty"`
	Material          string           `json:"material,omitempty"`
	Weight            core.NumberModel `json:"weight,omitempty"`
	Length            core.NumberModel `json:"length,omitempty"`
	Height            core.NumberModel `json:"height,omitempty"`
	Width             core.NumberModel `json:"width,omitempty"`
	Q                 string           `json:"q,omitempty"`
}

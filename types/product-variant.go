package types

import "github.com/driver005/gateway/core"

type FilterableProductVariant struct {
	core.FilterModel
	Title             *[]string        `json:"title,omitempty"`
	ProductID         *[]string        `json:"product_id,omitempty"`
	SKU               *[]string        `json:"sku,omitempty"`
	Barcode           *[]string        `json:"barcode,omitempty"`
	EAN               *[]string        `json:"ean,omitempty"`
	UPC               string           `json:"upc,omitempty"`
	InventoryQuantity core.NumberModel `json:"inventory_quantity,omitempty"`
	AllowBackorder    *bool            `json:"allow_backorder,omitempty"`
	ManageInventory   *bool            `json:"manage_inventory,omitempty"`
	HSCode            *[]string        `json:"hs_code,omitempty"`
	OriginCountry     *[]string        `json:"origin_country,omitempty"`
	MidCode           *[]string        `json:"mid_code,omitempty"`
	Material          string           `json:"material,omitempty"`
	Weight            core.NumberModel `json:"weight,omitempty"`
	Length            core.NumberModel `json:"length,omitempty"`
	Height            core.NumberModel `json:"height,omitempty"`
	Width             core.NumberModel `json:"width,omitempty"`
	Q                 string           `json:"q,omitempty"`
}

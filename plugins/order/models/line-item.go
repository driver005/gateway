package models

import (
	"github.com/driver005/gateway/core"
)

type LineItem struct {
	core.BaseModel

	TotalsId              string               `gorm:"column:totals_id;type:text" json:"totals_id"`
	Title                 string               `gorm:"column:title;type:text;not null" json:"title"`
	Subtitle              string               `gorm:"column:subtitle;type:text" json:"subtitle"`
	Thumbnail             string               `gorm:"column:thumbnail;type:text" json:"thumbnail"`
	VariantId             string               `gorm:"column:variant_id;type:text;index:IDX_order_line_item_variant_id,priority:1" json:"variant_id"`
	ProductId             string               `gorm:"column:product_id;type:text;index:IDX_order_line_item_product_id,priority:1" json:"product_id"`
	ProductTitle          string               `gorm:"column:product_title;type:text" json:"product_title"`
	ProductDescription    string               `gorm:"column:product_description;type:text" json:"product_description"`
	ProductSubtitle       string               `gorm:"column:product_subtitle;type:text" json:"product_subtitle"`
	ProductType           string               `gorm:"column:product_type;type:text" json:"product_type"`
	ProductCollection     string               `gorm:"column:product_collection;type:text" json:"product_collection"`
	ProductHandle         string               `gorm:"column:product_handle;type:text" json:"product_handle"`
	VariantSku            string               `gorm:"column:variant_sku;type:text" json:"variant_sku"`
	VariantBarcode        string               `gorm:"column:variant_barcode;type:text" json:"variant_barcode"`
	VariantTitle          string               `gorm:"column:variant_title;type:text" json:"variant_title"`
	VariantOptionValues   core.JSONB           `gorm:"column:variant_option_values;type:jsonb" json:"variant_option_values"`
	RequiresShipping      bool                 `gorm:"column:requires_shipping;type:boolean;not null;default:true" json:"requires_shipping"`
	IsDiscountable        bool                 `gorm:"column:is_discountable;type:boolean;not null;default:true" json:"is_discountable"`
	IsTaxInclusive        bool                 `gorm:"column:is_tax_inclusive;type:boolean;not null" json:"is_tax_inclusive"`
	CompareAtUnitPrice    float64              `gorm:"column:compare_at_unit_price;type:numeric" json:"compare_at_unit_price"`
	RawCompareAtUnitPrice core.JSONB           `gorm:"column:raw_compare_at_unit_price;type:jsonb" json:"raw_compare_at_unit_price"`
	UnitPrice             float64              `gorm:"column:unit_price;type:numeric;not null" json:"unit_price"`
	RawUnitPrice          core.JSONB           `gorm:"column:raw_unit_price;type:jsonb;not null" json:"raw_unit_price"`
	TaxLines              []LineItemTaxLine    `gorm:"foreignkey:ItemId" json:"tax_lines"`
	Adjustments           []LineItemAdjustment `gorm:"foreignkey:ItemId" json:"adjustments"`
	Totals                *OrderDetail         `gorm:"foreignkey:TotalsId" json:"totals"`
}

func (*LineItem) TableName() string {
	return "order_line_item"
}

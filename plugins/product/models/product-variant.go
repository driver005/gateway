package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type ProductVariant struct {
	core.BaseModel

	Title             string               `gorm:"column:title;type:text;not null" json:"title"`
	Sku               string               `gorm:"column:sku;type:text;uniqueIndex:IDX_product_variant_sku_unique,priority:1" json:"sku"`
	Barcode           string               `gorm:"column:barcode;type:text;uniqueIndex:IDX_product_variant_barcode_unique,priority:1" json:"barcode"`
	Ean               string               `gorm:"column:ean;type:text;uniqueIndex:IDX_product_variant_ean_unique,priority:1" json:"ean"`
	Upc               string               `gorm:"column:upc;type:text;uniqueIndex:IDX_product_variant_upc_unique,priority:1" json:"upc"`
	InventoryQuantity float64              `gorm:"column:inventory_quantity;type:numeric;not null;default:100" json:"inventory_quantity"`
	AllowBackorder    bool                 `gorm:"column:allow_backorder;type:boolean;not null" json:"allow_backorder"`
	ManageInventory   bool                 `gorm:"column:manage_inventory;type:boolean;not null;default:true" json:"manage_inventory"`
	HsCode            string               `gorm:"column:hs_code;type:text" json:"hs_code"`
	OriginCountry     string               `gorm:"column:origin_country;type:text" json:"origin_country"`
	MidCode           string               `gorm:"column:mid_code;type:text" json:"mid_code"`
	Material          string               `gorm:"column:material;type:text" json:"material"`
	Weight            float64              `gorm:"column:weight;type:numeric" json:"weight"`
	Length            float64              `gorm:"column:length;type:numeric" json:"length"`
	Height            float64              `gorm:"column:height;type:numeric" json:"height"`
	Width             float64              `gorm:"column:width;type:numeric" json:"width"`
	VariantRank       float64              `gorm:"column:variant_rank;type:numeric" json:"variant_rank"`
	ProductId         string               `gorm:"column:product_id;type:text;not null;index:IDX_product_variant_product_id,priority:1" json:"product_id"`
	Product           *Product             `gorm:"foreignKey:ProductId;constraint:OnDelete:CASCADE"`
	Options           []ProductOptionValue `gorm:"foreignKey:VariantId;constraint:OnDelete:CASCADE;"`
	DeletedAt         gorm.DeletedAt       `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_product_variant_deleted_at,priority:1" json:"deleted_at"`
}

func (*ProductVariant) TableName() string {
	return "product_variant"
}

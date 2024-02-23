package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type FulfillmentItem struct {
	core.BaseModel

	Title           string         `gorm:"column:title;type:text;not null" json:"title"`
	Sku             string         `gorm:"column:sku;type:text;not null" json:"sku"`
	Barcode         string         `gorm:"column:barcode;type:text;not null" json:"barcode"`
	Quantity        float64        `gorm:"column:quantity;type:numeric;not null" json:"quantity"`
	LineItemId      string         `gorm:"column:line_item_id;type:text;index:IDX_fulfillment_item_line_item_id,priority:1" json:"line_item_id"`
	InventoryItemId string         `gorm:"column:inventory_item_id;type:text;index:IDX_fulfillment_item_inventory_item_id,priority:1" json:"inventory_item_id"`
	FulfillmentId   string         `gorm:"column:fulfillment_id;type:text;not null;index:IDX_fulfillment_item_fulfillment_id,priority:1" json:"fulfillment_id"`
	Fulfillment     *Fulfillment   `gorm:"foreignKey:FulfillmentID;references:ID;constraint:OnDelete:CASCADE" json:"fulfillment"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_fulfillment_item_deleted_at,priority:1" json:"deleted_at"`
}

func (*FulfillmentItem) TableName() string {
	return "fulfillment_item"
}

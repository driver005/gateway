package models

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/plugins/order/types"
)

type OrderDetail struct {
	core.BaseModel

	OrderId                    string             `gorm:"column:order_id;type:text;not null;uniqueIndex:IDX_order_detail_order_id_item_id_version,priority:1" json:"order_id"`
	Order                      *Order             `gorm:"foreignkey:OrderId" json:"order"`
	Version                    int32              `gorm:"column:version;type:integer;not null;uniqueIndex:IDX_order_detail_order_id_item_id_version,priority:2" json:"version"`
	ItemId                     string             `gorm:"column:item_id;type:text;not null;uniqueIndex:IDX_order_detail_order_id_item_id_version,priority:3" json:"item_id"`
	Item                       *LineItem          `gorm:"foreignkey:ItemId" json:"item"`
	Quantity                   float64            `gorm:"column:quantity;type:numeric;not null" json:"quantity"`
	RawQuantity                string             `gorm:"column:raw_quantity;type:jsonb;not null" json:"raw_quantity"`
	FulfilledQuantity          float64            `gorm:"column:fulfilled_quantity;type:numeric;not null" json:"fulfilled_quantity"`
	RawFulfilledQuantity       string             `gorm:"column:raw_fulfilled_quantity;type:jsonb;not null" json:"raw_fulfilled_quantity"`
	ShippedQuantity            float64            `gorm:"column:shipped_quantity;type:numeric;not null" json:"shipped_quantity"`
	RawShippedQuantity         string             `gorm:"column:raw_shipped_quantity;type:jsonb;not null" json:"raw_shipped_quantity"`
	ReturnRequestedQuantity    float64            `gorm:"column:return_requested_quantity;type:numeric;not null" json:"return_requested_quantity"`
	RawReturnRequestedQuantity string             `gorm:"column:raw_return_requested_quantity;type:jsonb;not null" json:"raw_return_requested_quantity"`
	ReturnReceivedQuantity     float64            `gorm:"column:return_received_quantity;type:numeric;not null" json:"return_received_quantity"`
	RawReturnReceivedQuantity  string             `gorm:"column:raw_return_received_quantity;type:jsonb;not null" json:"raw_return_received_quantity"`
	ReturnDismissedQuantity    float64            `gorm:"column:return_dismissed_quantity;type:numeric;not null" json:"return_dismissed_quantity"`
	RawReturnDismissedQuantity string             `gorm:"column:raw_return_dismissed_quantity;type:jsonb;not null" json:"raw_return_dismissed_quantity"`
	WrittenOffQuantity         float64            `gorm:"column:written_off_quantity;type:numeric;not null" json:"written_off_quantity"`
	RawWrittenOffQuantity      string             `gorm:"column:raw_written_off_quantity;type:jsonb;not null" json:"raw_written_off_quantity"`
	Summary                    *types.ItemSummary `gorm:"column:summary;type:jsonb;not null" json:"summary"`
}

func (*OrderDetail) TableName() string {
	return "order_detail"
}

package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type OrderSalesChannel struct {
	core.SoftDeletableModel

	OrderId        uuid.NullUUID `gorm:"column:order_id;type:character varying;primaryKey;uniqueIndex:order_sales_channel_order_id_unique,priority:1" json:"order_id"`
	SalesChannelId uuid.NullUUID `gorm:"column:sales_channel_id;type:character varying;primaryKey" json:"sales_channel_id"`
}

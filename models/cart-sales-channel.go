package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type CartSalesChannel struct {
	core.SoftDeletableModel

	CartId         uuid.NullUUID `gorm:"column:cart_id;type:character varying;primaryKey;uniqueIndex:cart_sales_channel_cart_id_unique,priority:1" json:"cart_id"`
	SalesChannelId uuid.NullUUID `gorm:"column:sales_channel_id;type:character varying;primaryKey" json:"sales_channel_id"`
}

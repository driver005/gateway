package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type ProductSalesChannel struct {
	core.BaseModel

	ProductId      uuid.NullUUID `gorm:"column:product_id;type:character varying;not null;uniqueIndex:product_sales_channel_product_id_sales_channel_id_unique,priority:1;index:IDX_5a4d5e1e60f97633547821ec8d,priority:1" json:"product_id"`
	SalesChannelId uuid.NullUUID `gorm:"column:sales_channel_id;type:character varying;not null;uniqueIndex:product_sales_channel_product_id_sales_channel_id_unique,priority:2;index:IDX_37341bad297fe5cca91f921032,priority:1" json:"sales_channel_id"`
}

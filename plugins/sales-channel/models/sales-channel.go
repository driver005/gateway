package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type SalesChannel struct {
	core.BaseModel

	Name        string         `gorm:"column:name;type:text;not null" json:"name"`
	Description string         `gorm:"column:description;type:text" json:"description"`
	IsDisabled  bool           `gorm:"column:is_disabled;type:boolean;not null" json:"is_disabled"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_sales_channel_deleted_at,priority:1" json:"deleted_at"`
}

func (*SalesChannel) TableName() string {
	return "sales_channel"
}

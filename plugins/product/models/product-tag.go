package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type ProductTag struct {
	core.BaseModel

	Value     string         `gorm:"column:value;type:text;not null" json:"value"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_product_tag_deleted_at,priority:1" json:"deleted_at"`
	Products  []Product      `gorm:"many2many:product_tags;" json:"products"`
}

func (*ProductTag) TableName() string {
	return "product_tag"
}

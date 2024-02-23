package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type ProductType struct {
	core.BaseModel

	Value     string         `gorm:"column:value;type:text;not null" json:"value"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_product_type_deleted_at,priority:1" json:"deleted_at"`
}

func (*ProductType) TableName() string {
	return "product_type"
}

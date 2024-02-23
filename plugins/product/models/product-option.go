package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type ProductOption struct {
	core.BaseModel

	Title     string               `gorm:"column:title;type:text;not null" json:"title"`
	ProductId string               `gorm:"column:product_id;type:text;not null;index:IDX_product_option_product_id,priority:1" json:"product_id"`
	Product   *Product             `gorm:"foreignKey:ProductId" json:"product"`
	Values    []ProductOptionValue `gorm:"foreignKey:OptionId" json:"values"`
	DeletedAt gorm.DeletedAt       `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_product_option_deleted_at,priority:1" json:"deleted_at"`
}

func (*ProductOption) TableName() string {
	return "product_option"
}

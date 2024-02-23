package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type ProductOptionValue struct {
	core.BaseModel

	Value     string          `gorm:"column:value;type:text;not null" json:"value"`
	OptionId  string          `gorm:"column:option_id;type:text;not null;index:IDX_product_option_value_option_id,priority:1" json:"option_id"`
	Option    *ProductOption  `gorm:"foreignKey:OptionId" json:"option"`
	VariantId string          `gorm:"column:variant_id;type:text;not null;index:IDX_product_option_value_variant_id,priority:1" json:"variant_id"`
	Variant   *ProductVariant `gorm:"foreignKey:VariantId" json:"variant"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_product_option_value_deleted_at,priority:1" json:"deleted_at"`
}

func (*ProductOptionValue) TableName() string {
	return "product_option_value"
}

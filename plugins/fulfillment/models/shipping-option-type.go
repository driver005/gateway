package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type ShippingOptionType struct {
	core.BaseModel

	Label            string          `gorm:"column:label;type:text;not null" json:"label"`
	Description      string          `gorm:"column:description;type:text" json:"description"`
	Code             string          `gorm:"column:code;type:text;not null" json:"code"`
	ShippingOptionId string          `gorm:"column:shipping_option_id;type:text;not null;index:IDX_shipping_option_type_shipping_option_id,priority:1" json:"shipping_option_id"`
	ShippingOption   *ShippingOption `gorm:"foreignKey:ShippingOptionId" json:"shipping_option"`
	DeletedAt        gorm.DeletedAt  `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_shipping_option_type_deleted_at,priority:1" json:"deleted_at"`
}

func (*ShippingOptionType) TableName() string {
	return "shipping_option_type"
}

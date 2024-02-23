package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type ShippingOptionRule struct {
	core.BaseModel

	Attribute        string          `gorm:"column:attribute;type:text;not null" json:"attribute"`
	Operator         string          `gorm:"column:operator;type:text;not null" json:"operator"`
	Value            string          `gorm:"column:value;type:jsonb" json:"value"`
	ShippingOptionId string          `gorm:"column:shipping_option_id;type:text;not null;index:IDX_shipping_option_rule_shipping_option_id,priority:1" json:"shipping_option_id"`
	ShippingOption   *ShippingOption `gorm:"foreignKey:ShippingOptionId" json:"shipping_option"`
	DeletedAt        gorm.DeletedAt  `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_shipping_option_rule_deleted_at,priority:1" json:"deleted_at"`
}

func (*ShippingOptionRule) TableName() string {
	return "shipping_option_rule"
}

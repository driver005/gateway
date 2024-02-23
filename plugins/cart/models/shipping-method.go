package models

import "github.com/driver005/gateway/core"

type ShippingMethod struct {
	core.BaseModel

	CartId                   string                     `gorm:"column:cart_id;type:text;not null;index:IDX_shipping_method_cart_id,priority:1" json:"cart_id"`
	Cart                     Cart                       `gorm:"foreignKey:CartId" json:"cart"`
	Name                     string                     `gorm:"column:name;type:text;not null" json:"name"`
	Description              string                     `gorm:"column:description;type:jsonb" json:"description"`
	Amount                   float64                    `gorm:"column:amount;type:numeric;not null" json:"amount"`
	RawAmount                string                     `gorm:"column:raw_amount;type:jsonb;not null" json:"raw_amount"`
	IsTaxInclusive           bool                       `gorm:"column:is_tax_inclusive;type:boolean;not null" json:"is_tax_inclusive"`
	ShippingOptionId         string                     `gorm:"column:shipping_option_id;type:text;index:IDX_shipping_method_option_id,priority:1" json:"shipping_option_id"`
	Data                     core.JSONB                 `gorm:"column:data;type:jsonb" json:"data"`
	ShippingMethodTaxLine    []ShippingMethodTaxLine    `gorm:"foreignkey:ShippingMethodId" json:"shipping_method_tax_line"`
	ShippingMethodAdjustment []ShippingMethodAdjustment `gorm:"foreignkey:ShippingMethodId" json:"shipping_method_adjustment"`
}

func (*ShippingMethod) TableName() string {
	return "cart_shipping_method"
}

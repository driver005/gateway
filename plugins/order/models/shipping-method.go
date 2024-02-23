package models

import "github.com/driver005/gateway/core"

type ShippingMethod struct {
	core.BaseModel

	OrderId                  string                     `gorm:"column:order_id;type:text;not null;index:IDX_order_shipping_method_order_id,priority:1" json:"order_id"`
	Order                    *Order                     `gorm:"foreignkey:OrderId" json:"order"`
	Name                     string                     `gorm:"column:name;type:text;not null" json:"name"`
	Description              string                     `gorm:"column:description;type:jsonb" json:"description"`
	Amount                   float64                    `gorm:"column:amount;type:numeric;not null" json:"amount"`
	RawAmount                string                     `gorm:"column:raw_amount;type:jsonb;not null" json:"raw_amount"`
	IsTaxInclusive           bool                       `gorm:"column:is_tax_inclusive;type:boolean;not null" json:"is_tax_inclusive"`
	ShippingOptionId         string                     `gorm:"column:shipping_option_id;type:text;index:IDX_order_shipping_method_shipping_option_id,priority:1" json:"shipping_option_id"`
	Data                     core.JSONB                 `gorm:"column:data;type:jsonb" json:"data"`
	ShippingMethodTaxLine    []ShippingMethodTaxLine    `gorm:"foreignkey:ShippingMethodId" json:"shipping_method_tax_line"`
	ShippingMethodAdjustment []ShippingMethodAdjustment `gorm:"foreignkey:ShippingMethodId" json:"shipping_method_adjustment"`
}

func (*ShippingMethod) TableName() string {
	return "order_shipping_method"
}

package models

import "github.com/driver005/gateway/core"

type ShippingMethodTaxLine struct {
	core.BaseModel

	Description      string          `gorm:"column:description;type:text" json:"description"`
	TaxRateId        string          `gorm:"column:tax_rate_id;type:text" json:"tax_rate_id"`
	Code             string          `gorm:"column:code;type:text;not null" json:"code"`
	Rate             float64         `gorm:"column:rate;type:numeric;not null" json:"rate"`
	RawRate          string          `gorm:"column:raw_rate;type:jsonb;not null" json:"raw_rate"`
	ProviderId       string          `gorm:"column:provider_id;type:text" json:"provider_id"`
	ShippingMethodId string          `gorm:"column:shipping_method_id;type:text;index:IDX_order_shipping_method_tax_line_shipping_method_id,priority:1" json:"shipping_method_id"`
	ShippingMethod   *ShippingMethod `gorm:"foreignkey:ShippingMethodId;association_foreignkey:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"shipping_method"`
}

func (*ShippingMethodTaxLine) TableName() string {
	return "order_shipping_method_tax_line"
}

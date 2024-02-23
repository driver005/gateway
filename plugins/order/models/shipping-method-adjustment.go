package models

import "github.com/driver005/gateway/core"

type ShippingMethodAdjustment struct {
	core.BaseModel

	Description      string          `gorm:"column:description;type:text" json:"description"`
	PromotionId      string          `gorm:"column:promotion_id;type:text" json:"promotion_id"`
	Code             string          `gorm:"column:code;type:text" json:"code"`
	Amount           float64         `gorm:"column:amount;type:numeric;not null" json:"amount"`
	RawAmount        string          `gorm:"column:raw_amount;type:jsonb;not null" json:"raw_amount"`
	ProviderId       string          `gorm:"column:provider_id;type:text" json:"provider_id"`
	ShippingMethodId string          `gorm:"column:shipping_method_id;type:text;index:IDX_order_shipping_method_adjustment_shipping_method_id,priority:1" json:"shipping_method_id"`
	ShippingMethod   *ShippingMethod `gorm:"foreignkey:ShippingMethodId;association_foreignkey:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"shipping_method"`
}

func (*ShippingMethodAdjustment) TableName() string {
	return "order_shipping_method_adjustment"
}

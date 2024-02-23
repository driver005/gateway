package models

import "github.com/driver005/gateway/core"

type LineItemAdjustment struct {
	core.BaseModel

	Description string    `gorm:"column:description;type:text" json:"description"`
	PromotionId string    `gorm:"column:promotion_id;type:text" json:"promotion_id"`
	Code        string    `gorm:"column:code;type:text" json:"code"`
	Amount      float64   `gorm:"column:amount;type:numeric;not null" json:"amount"`
	RawAmount   string    `gorm:"column:raw_amount;type:jsonb;not null" json:"raw_amount"`
	ProviderId  string    `gorm:"column:provider_id;type:text" json:"provider_id"`
	ItemId      string    `gorm:"column:item_id;type:text;index:IDX_order_line_item_adjustment_item_id,priority:1" json:"item_id"`
	Item        *LineItem `gorm:"foreignkey:ItemId;association_foreignkey:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"item"`
}

func (*LineItemAdjustment) TableName() string {
	return "order_line_item_adjustment"
}

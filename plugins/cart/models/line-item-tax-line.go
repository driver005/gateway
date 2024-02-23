package models

import (
	"github.com/driver005/gateway/core"
)

type LineItemTaxLine struct {
	core.BaseModel

	Description string    `gorm:"column:description;type:text" json:"description"`
	TaxRateId   string    `gorm:"column:tax_rate_id;type:text;index:IDX_line_item_tax_line_tax_rate_id,priority:1" json:"tax_rate_id"`
	Code        string    `gorm:"column:code;type:text;not null" json:"code"`
	Rate        float64   `gorm:"column:rate;type:numeric;not null" json:"rate"`
	RawRate     string    `gorm:"column:raw_rate;type:jsonb;not null" json:"raw_rate"`
	ProviderId  string    `gorm:"column:provider_id;type:text" json:"provider_id"`
	ItemId      string    `gorm:"type:text;index:IDX_tax_line_item_id" json:"item_id"`
	Item        *LineItem `gorm:"foreignkey:ItemId;association_foreignkey:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"item"`
}

func (*LineItemTaxLine) TableName() string {
	return "cart_line_item_tax_line"
}

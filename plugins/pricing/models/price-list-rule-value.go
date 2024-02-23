package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type PriceListRuleValue struct {
	core.BaseModel

	Value           string         `gorm:"column:value;type:text;not null" json:"value"`
	PriceListRuleId string         `gorm:"column:price_list_rule_id;type:text;not null;index:IDX_price_list_rule_price_list_rule_value_id,priority:1" json:"price_list_rule_id"`
	PriceListRule   *PriceListRule `gorm:"foreignKey:PriceListRuleId" json:"price_list_rule"`
	DeletedAt       gorm.DeletedAt `gorm:"index:idx_price_list_rule_value_deleted_at;column:deleted_at;type:timestamptz;default:null"`
}

func (*PriceListRuleValue) TableName() string {
	return "price_list_rule_value"
}

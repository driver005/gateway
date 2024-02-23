package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type PriceSetMoneyAmountRules struct {
	core.BaseModel

	PriceSetMoneyAmountId string         `gorm:"column:price_set_money_amount_id;type:text;not null;index:IDX_price_set_money_amount_rules_price_set_money_amount_id,priority:1" json:"price_set_money_amount_id"`
	RuleTypeId            string         `gorm:"column:rule_type_id;type:text;not null;index:IDX_price_set_money_amount_rules_rule_type_id,priority:1" json:"rule_type_id"`
	Value                 string         `gorm:"column:value;type:text;not null" json:"value"`
	DeletedAt             gorm.DeletedAt `gorm:"index:idx_price_set_money_amount_rules_deleted_at;column:deleted_at;type:timestamptz;default:null"`
}

func (*PriceSetMoneyAmountRules) TableName() string {
	return "price_set_money_amount_rules"
}

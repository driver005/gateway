package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type PriceRule struct {
	core.BaseModel

	PriceSetId            string         `gorm:"column:price_set_id;type:text;not null;index:IDX_price_rule_price_set_id,priority:1" json:"price_set_id"`
	RuleTypeId            string         `gorm:"column:rule_type_id;type:text;not null;index:IDX_price_rule_rule_type_id,priority:1" json:"rule_type_id"`
	Value                 string         `gorm:"column:value;type:text;not null" json:"value"`
	Priority              int32          `gorm:"column:priority;type:integer;not null" json:"priority"`
	PriceSetMoneyAmountId string         `gorm:"column:price_set_money_amount_id;type:text;not null;index:IDX_price_rule_price_set_money_amount_id,priority:1" json:"price_set_money_amount_id"`
	DeletedAt             gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_price_rule_deleted_at,priority:1" json:"deleted_at"`
}

func (*PriceRule) TableName() string {
	return "price_rule"
}

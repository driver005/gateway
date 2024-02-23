package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type PriceSetMoneyAmount struct {
	core.BaseModel

	Title                    string                     `gorm:"column:title;type:text;not null" json:"title"`
	PriceSetId               string                     `gorm:"column:price_set_id;type:text;not null;index:IDX_price_set_money_amount_price_set_id,priority:1" json:"price_set_id"`
	MoneyAmountId            string                     `gorm:"column:money_amount_id;type:text;not null;uniqueIndex:price_set_money_amount_money_amount_id_unique,priority:1;index:IDX_price_set_money_amount_money_amount_id,priority:1" json:"money_amount_id"`
	RulesCount               int32                      `gorm:"column:rules_count;type:integer;not null" json:"rules_count"`
	PriceListId              string                     `gorm:"column:price_list_id;type:text;index:IDX_price_rule_price_list_id,priority:1" json:"price_list_id"`
	PriceSet                 *PriceSet                  `gorm:"foreignKey:PriceSetId;references:id;constraint:OnDelete:CASCADE;index:idx_price_set_money_amount_price_set_id" json:"price_set"`
	MoneyAmount              *MoneyAmount               `gorm:"foreignKey:MoneyAmountId;references:id;constraint:OnDelete:CASCADE;index:idx_price_set_money_amount_money_amount_id" json:"money_amount"`
	PriceList                *PriceList                 `gorm:"foreignKey:PriceListId;references:id;constraint:OnDelete:CASCADE;index:idx_price_set_money_amount_price_list_id" json:"price_list"`
	PriceRules               []PriceRule                `gorm:"foreignKey:PriceSetMoneyAmountId;constraint:OnDelete:CASCADE;index:idx_price_set_money_amount_price_rules" json:"price_rules"`
	PriceSetMoneyAmountRules []PriceSetMoneyAmountRules `gorm:"foreignKey:PriceSetMoneyAmountId;constraint:OnDelete:CASCADE;index:idx_price_set_money_amount_rules" json:"price_set_money_amount_rules"`
	DeletedAt                gorm.DeletedAt             `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_price_set_money_amount_deleted_at,priority:1" json:"deleted_at"`
}

func (*PriceSetMoneyAmount) TableName() string {
	return "price_set_money_amount"
}

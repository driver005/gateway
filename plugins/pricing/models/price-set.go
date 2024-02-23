package models

import "github.com/driver005/gateway/core"

type PriceSet struct {
	core.BaseModel
	//TODO: check uuid
	// ID string `gorm:"column:id;type:text;primaryKey" json:"id"`
	PriceSetMoneyAmounts []PriceSetMoneyAmount `gorm:"foreignKey:PriceSetId;references:id;constraint:OnDelete:CASCADE;index:idx_price_set_price_set_money_amounts" json:"price_set_money_amounts"`
	PriceRules           []PriceRule           `gorm:"foreignKey:PriceSetId;references:id;constraint:OnDelete:CASCADE;index:idx_price_set_price_rules" json:"price_rules"`
	MoneyAmounts         []MoneyAmount         `gorm:"many2many:price_set_money_amounts;constraint:OnDelete:CASCADE;index:idx_price_set_money_amounts" json:"money_amounts"`
	RuleTypes            []RuleType            `gorm:"many2many:price_set_rule_types;constraint:OnDelete:CASCADE;index:idx_price_set_rule_types" json:"rule_types"`
}

func (*PriceSet) TableName() string {
	return "price_set"
}

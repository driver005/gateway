package models

import "github.com/driver005/gateway/core"

type PriceSetRuleType struct {
	core.BaseModel
	//TODO: check uuid
	// ID         string    `gorm:"column:id;type:text;primaryKey" json:"id"`
	PriceSetId string    `gorm:"column:price_set_id;type:text;not null;index:IDX_price_set_rule_type_price_set_id,priority:1" json:"price_set_id"`
	RuleTypeId string    `gorm:"column:rule_type_id;type:text;not null;index:IDX_price_set_rule_type_rule_type_id,priority:1" json:"rule_type_id"`
	PriceSet   *PriceSet `gorm:"foreignKey:PriceSetID;references:id;constraint:OnDelete:CASCADE;index:idx_price_set_rule_type_price_set_id"`
	RuleType   *RuleType `gorm:"foreignKey:RuleTypeID;references:id;constraint:OnDelete:CASCADE;index:idx_price_set_rule_type_rule_type_id"`
}

func (*PriceSetRuleType) TableName() string {
	return "price_set_rule_type"
}

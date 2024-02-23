package models

type PriceListRule struct {
	ID                  string               `gorm:"column:id;type:text;primaryKey" json:"id"`
	RuleTypeId          string               `gorm:"column:rule_type_id;type:text;not null;uniqueIndex:IDX_price_list_rule_rule_type_id_price_list_id_unique,priority:1;index:IDX_price_list_rule_rule_type_id,priority:1" json:"rule_type_id"`
	PriceListId         string               `gorm:"column:price_list_id;type:text;not null;uniqueIndex:IDX_price_list_rule_rule_type_id_price_list_id_unique,priority:2;index:IDX_price_list_rule_price_list_id,priority:1" json:"price_list_id"`
	RuleType            *RuleType            `gorm:"foreignKey:RuleTypeId" json:"rule_type"`
	PriceList           *PriceList           `gorm:"foreignKey:PriceListId" json:"price_list"`
	PriceListRuleValues []PriceListRuleValue `gorm:"foreignKey:PriceListRuleId" json:"price_list_rule_values"`
}

func (*PriceListRule) TableName() string {
	return "price_list_rule"
}

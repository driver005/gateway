package models

import "github.com/driver005/gateway/core"

type RuleType struct {
	core.BaseModel
	//TODO: check uuid
	// ID              string     `gorm:"column:id;type:text;primaryKey" json:"id"`
	Name            string     `gorm:"column:name;type:text;not null" json:"name"`
	RuleAttribute   string     `gorm:"column:rule_attribute;type:text;not null;index:IDX_rule_type_rule_attribute,priority:1" json:"rule_attribute"`
	DefaultPriority int32      `gorm:"column:default_priority;type:integer;not null" json:"default_priority"`
	PriceSets       []PriceSet `gorm:"many2many:price_set_rule_types;constraint:OnDelete:CASCADE;index:idx_price_set_rule_types" json:"price_sets"`
}

func (*RuleType) TableName() string {
	return "rule_type"
}

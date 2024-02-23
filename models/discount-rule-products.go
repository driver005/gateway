package models

import "github.com/google/uuid"

type DiscountRuleProduct struct {
	DiscountRuleId uuid.NullUUID `gorm:"column:discount_rule_id;type:character varying;primaryKey;index:IDX_4e0739e5f0244c08d41174ca08,priority:1" json:"discount_rule_id"`
	ProductId      uuid.NullUUID `gorm:"column:product_id;type:character varying;primaryKey;index:IDX_be66106a673b88a81c603abe7e,priority:1" json:"product_id"`
}

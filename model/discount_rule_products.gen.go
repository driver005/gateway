// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameDiscountRuleProduct = "discount_rule_products"

// DiscountRuleProduct mapped from table <discount_rule_products>
type DiscountRuleProduct struct {
	DiscountRuleID string `gorm:"column:discount_rule_id;type:character varying;primaryKey;index:IDX_4e0739e5f0244c08d41174ca08,priority:1" json:"discount_rule_id"`
	ProductID      string `gorm:"column:product_id;type:character varying;primaryKey;index:IDX_be66106a673b88a81c603abe7e,priority:1" json:"product_id"`
}

// TableName DiscountRuleProduct's table name
func (*DiscountRuleProduct) TableName() string {
	return TableNameDiscountRuleProduct
}
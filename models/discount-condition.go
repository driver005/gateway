package models

import (
	"database/sql/driver"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// DiscountCondition - Holds rule conditions for when a discount is applicable
type DiscountCondition struct {
	core.Model

	// The type of the Condition
	Type DiscountConditionType `json:"type" gorm:"type:enum('products','product_types','product_collections','product_tags','customer_groups')"`

	// The operator of the Condition
	Operator string `json:"operator"`

	// The ID of the discount rule associated with the condition
	DiscountRuleId uuid.NullUUID `json:"discount_rule_id"`

	DiscountRule *DiscountRule `json:"discount_rule" gorm:"foreignKey:id;references:discount_rule_id"`

	// products associated with this condition if type = products. Available if the relation `products` is expanded.
	Products []Product `json:"products" gorm:"many2many:discount_condition_product"`

	// product types associated with this condition if type = product_types. Available if the relation `product_types` is expanded.
	ProductTypes []ProductType `json:"product_types" gorm:"many2many:discount_condition_product_types"`

	// product tags associated with this condition if type = product_tags. Available if the relation `product_tags` is expanded.
	ProductTags []ProductTag `json:"product_tags" gorm:"many2many:discount_condition_product_tag"`

	// product collections associated with this condition if type = product_collections. Available if the relation `product_collections` is expanded.
	ProductCollections []ProductCollection `json:"product_collections" gorm:"many2many:discount_condition_product_collection"`

	// customer groups associated with this condition if type = customer_groups. Available if the relation `customer_groups` is expanded.
	CustomerGroups []CustomerGroup `json:"customer_groups" gorm:"many2many:discount_condition_customer_group"`
}

// The status of the Price List
type DiscountConditionType string

// Defines values for DiscountConditionType.
const (
	DiscountConditionTypeRroducts           DiscountConditionType = "products"
	DiscountConditionTypeProductTypes       DiscountConditionType = "product_types"
	DiscountConditionTypeProductCollections DiscountConditionType = "product_collections"
	DiscountConditionTypeProductTags        DiscountConditionType = "product_tags"
	DiscountConditionTypeCustomerGroups     DiscountConditionType = "customer_groups"
)

func (pl *DiscountConditionType) Scan(value interface{}) error {
	*pl = DiscountConditionType(value.([]byte))
	return nil
}

func (pl DiscountConditionType) Value() (driver.Value, error) {
	return string(pl), nil
}

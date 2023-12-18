package models

import "github.com/google/uuid"

// DiscountConditionProductType - Associates a discount condition with a product type
type DiscountConditionProductType struct {
	// The ID of the Product Tag
	ProductTypeId uuid.NullUUID `json:"product_type_id"`

	ProductType *ProductType `json:"product_type" gorm:"foreignKey:id;references:product_type_id"`

	// The ID of the Discount Condition
	ConditionId uuid.NullUUID `json:"condition_id"`

	DiscountCondition *DiscountCondition `json:"discount_condition" gorm:"foreignKey:id;references:condition_id"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" gorm:"default:null"`
}

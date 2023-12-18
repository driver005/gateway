package models

import "github.com/google/uuid"

// DiscountConditionProduct - Associates a discount condition with a product
type DiscountConditionProduct struct {
	// The ID of the Product Tag
	ProductId uuid.NullUUID `json:"product_id"`

	Product *Product `json:"product" gorm:"foreignKey:id;references:product_id"`

	// The ID of the Discount Condition
	ConditionId uuid.NullUUID `json:"condition_id"`

	DiscountCondition *DiscountCondition `json:"discount_condition" gorm:"foreignKey:id;references:condition_id"`

	// An optional key-value map with additional details
}

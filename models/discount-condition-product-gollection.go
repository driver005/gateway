package models

import "github.com/google/uuid"

// DiscountConditionProductCollection - Associates a discount condition with a product collection
type DiscountConditionProductCollection struct {
	// The ID of the Product Collection
	ProductCollectionId uuid.NullUUID `json:"product_collection_id"`

	ProductCollection *ProductCollection `json:"product_collection" gorm:"foreignKey:id;references:product_collection_id"`

	// The ID of the Discount Condition
	ConditionId uuid.NullUUID `json:"condition_id"`

	DiscountCondition *DiscountCondition `json:"discount_condition" gorm:"foreignKey:id;references:condition_id"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" gorm:"default:null"`
}

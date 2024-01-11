package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// A Product Variant Inventory Item links variants with inventory items and denotes the required quantity of the variant.
type ProductVariantInventoryItem struct {
	core.Model

	// The id of the inventory item
	InventoryItemId uuid.NullUUID `json:"inventory_item_id"`

	// The id of the variant.
	VariantId uuid.NullUUID `json:"variant_id"`

	// The details of the product variant.
	Variant *ProductVariant `json:"variant,omitempty"  gorm:"foreignKey:id;references:variant_id"`

	// The quantity of an inventory item required for the variant.
	RequiredQuantity int `json:"required_quantity"`
}

package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// ClaimItem - Represents a claimed item along with information about the reasons for the claim.
type ClaimItem struct {
	core.Model

	// Available if the relation `images` is expanded.
	Images []ClaimImage `json:"images" gorm:"foreignKey:id"`

	// The ID of the claim this item is associated with.
	ClaimOrderId uuid.NullUUID `json:"claim_order_id"`

	// A claim order object. Available if the relation `claim_order` is expanded.
	ClaimOrder *ClaimOrder `json:"claim_order" gorm:"foreignKey:id;references:claim_order_id"`

	// The ID of the line item that the claim item refers to.
	ItemId uuid.NullUUID `json:"item_id"`

	Item *LineItem `json:"item" gorm:"foreignKey:id;references:item_id"`

	// The ID of the product variant that is claimed.
	VariantId uuid.NullUUID `json:"variant_id"`

	// A variant object. Available if the relation `variant` is expanded.
	Variant *ProductVariant `json:"variant" gorm:"foreignKey:id;references:variant_id"`

	// The reason for the claim
	Reason string `json:"reason"`

	// An optional note about the claim, for additional information
	Note string `json:"note" gorm:"default:null"`

	// The quantity of the item that is being claimed; must be less than or equal to the amount purchased in the original order.
	Quantity int `json:"quantity"`

	// User defined tags for easy filtering and grouping. Available if the relation 'tags' is expanded.
	Tags []ClaimTag `json:"tags" gorm:"foreignKey:id"`
}

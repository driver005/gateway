package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// ClaimImage - Represents photo documentation of a claim.
type ClaimImage struct {
	core.Model

	// The ID of the claim item associated with the image
	ClaimItemId uuid.NullUUID `json:"claim_item_id"`

	// A claim item object. Available if the relation `claim_item` is expanded.
	ClaimItem *ClaimItem `json:"claim_item" gorm:"foreignKey:id;references:claim_item_id"`

	// The URL of the image
	Url string `json:"url"`
}

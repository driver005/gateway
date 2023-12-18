package models

import "github.com/google/uuid"

// ClaimTag - Claim Tags are user defined tags that can be assigned to claim items for easy filtering and grouping.
type ClaimTag struct {
	Id uuid.UUID `json:"id gorm:primarykey"`

	// The value that the claim tag holds
	Value string `json:"value"`
}

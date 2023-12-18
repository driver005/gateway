package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// Note - Notes are elements which we can use in association with different resources to allow users to describe additional information in relation to these.
type Note struct {
	core.Model

	// The type of resource that the Note refers to.
	ResourceType string `json:"resource_type"`

	// The ID of the resource that the Note refers to.
	ResourceId uuid.NullUUID `json:"resource_id"`

	// The contents of the note.
	Value string `json:"value"`

	// The ID of the author (user)
	AuthorId uuid.NullUUID `json:"author_id" gorm:"default:null"`

	Author *User `json:"author" gorm:"foreignKey:id;references:author_id"`
}

package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

// @oas:schema:AdminPostNotesReq
// type: object
// description: "The details of the note to be created."
// required:
//   - resource_id
//   - resource_type
//   - value
//
// properties:
//
//	resource_id:
//	  type: string
//	  description: The ID of the resource which the Note relates to. For example, an order ID.
//	resource_type:
//	  type: string
//	  description: The type of resource which the Note relates to. For example, `order`.
//	value:
//	  type: string
//	  description: The content of the Note to create.
type CreateNoteInput struct {
	Value        string       `json:"value"`
	ResourceType string       `json:"resource_type"`
	ResourceId   uuid.UUID    `json:"resource_id"`
	AuthorId     uuid.UUID    `json:"author_id,omitempty" validate:"omitempty"`
	Author       *models.User `json:"author,omitempty" validate:"omitempty"`
	Metadata     core.JSONB   `json:"metadata,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostNotesNoteReq
// type: object
// description: "The details to update of the note."
// required:
//   - value
//
// properties:
//
//	value:
//	  type: string
//	  description: The description of the Note.
type UpdateNoteInput struct {
	Value string `json:"value"`
}

type Selector struct {
	ResourceId uuid.UUID `json:"resource_id,omitempty" validate:"omitempty"`
}

type FilterableNote struct {
	core.FilterModel

	ResourceType string    `json:"resource_type,omitempty" validate:"omitempty"`
	ResourceId   uuid.UUID `json:"resource_id,omitempty" validate:"omitempty"`
	Value        string    `json:"value,omitempty" validate:"omitempty"`
	AuthorId     uuid.UUID `json:"author_id,omitempty" validate:"omitempty"`
}

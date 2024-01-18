package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type CreateNoteInput struct {
	Value        string       `json:"value"`
	ResourceType string       `json:"resource_type"`
	ResourceId   uuid.UUID    `json:"resource_id"`
	AuthorId     uuid.UUID    `json:"author_id,omitempty" validate:"omitempty"`
	Author       *models.User `json:"author,omitempty" validate:"omitempty"`
	Metadata     core.JSONB   `json:"metadata,omitempty" validate:"omitempty"`
}

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

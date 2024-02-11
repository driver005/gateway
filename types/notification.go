package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FilterableNotification struct {
	core.FilterModel

	EventName      string    `json:"event_name,omitempty"`
	ResourceType   string    `json:"resource_type,omitempty"`
	ResourceId     uuid.UUID `json:"resource_id,omitempty"`
	To             string    `json:"to,omitempty"`
	IncludeResends bool      `json:"include_resends,omitempty"`
}

type ResendNotification struct {
	To string `json:"to,omitempty" validate:"omitempty"`
}

package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type CreateReturnReason struct {
	Value                string     `json:"value"`
	Label                string     `json:"label"`
	ParentReturnReasonId uuid.UUID  `json:"parent_return_reason_id,omitempty" validate:"omitempty"`
	Description          string     `json:"description,omitempty" validate:"omitempty"`
	Metadata             core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

type UpdateReturnReason struct {
	Description          string     `json:"description,omitempty" validate:"omitempty"`
	Label                string     `json:"label,omitempty" validate:"omitempty"`
	ParentReturnReasonId uuid.UUID  `json:"parent_return_reason_id,omitempty" validate:"omitempty"`
	Metadata             core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

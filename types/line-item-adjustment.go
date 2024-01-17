package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FilterableLineItemAdjustmentProps struct {
	core.FilterModel

	ItemId      uuid.UUID `json:"item_id,,omitempty" validate:"omitempty"`
	Description string    `json:"description,,omitempty" validate:"omitempty"`
	ResourceId  uuid.UUID `json:"resource_id,,omitempty" validate:"omitempty"`
}

package types

import (
	"github.com/google/uuid"
)

type OrdersReturnItem struct {
	ItemId   uuid.UUID `json:"item_id"`
	Quantity int       `json:"quantity"`
	ReasonId uuid.UUID `json:"reason_id,omitempty"`
	Note     string    `json:"note,omitempty"`
}

type TotalsContext struct {
	ForceTaxes      bool `json:"force_taxes,omitempty"`
	ReturnableItems bool `json:"returnable_items,omitempty"`
}

package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// ReturnItem - Correlates a Line Item with a Return, keeping track of the quantity of the Line Item that will be returned.
type ReturnItem struct {

	// The id of the Return that the Return Item belongs to.
	ReturnId uuid.NullUUID `json:"return_id" gorm:"primarykey"`

	ReturnOrder *Return `json:"return_order" gorm:"foreignKey:id;references:return_id"`

	// The id of the Line Item that the Return Item references.
	ItemId uuid.NullUUID `json:"item_id"`

	Item *LineItem `json:"item" gorm:"foreignKey:id;references:item_id"`

	// The quantity of the Line Item that is included in the Return.
	Quantity int `json:"quantity" gorm:"default:null"`

	// Whether the Return Item was requested initially or received unexpectedly in the warehouse.
	IsRequested bool `json:"is_requested" gorm:"default:null"`

	// The quantity that was originally requested to be returned.
	RequestedQuantity int `json:"requested_quantity" gorm:"default:null"`

	// The quantity that was received in the warehouse.
	RecievedQuantity int `json:"recieved_quantity" gorm:"default:null"`

	// The ID of the reason for returning the item.
	ReasonId uuid.NullUUID `json:"reason_id" gorm:"default:null"`

	Reason *ReturnReason `json:"reason" gorm:"foreignKey:id;references:reason_id"`

	// An optional note with additional details about the Return.
	Note string `json:"note" gorm:"default:null"`

	Metadata core.JSONB `json:"metadata" gorm:"default:null"`
}

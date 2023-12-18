package models

import "github.com/google/uuid"

// FulfillmentItem - Correlates a Line Item with a Fulfillment, keeping track of the quantity of the Line Item.
type FulfillmentItem struct {

	// The id of the Fulfillment that the Fulfillment Item belongs to.
	FulfillmentId uuid.NullUUID `json:"fulfillment_id" gorm:"primarykey"`

	// A fulfillment object. Available if the relation `fulfillment` is expanded.
	Fulfillment *Fulfillment `json:"fulfillment" gorm:"foreignKey:id;references:fulfillment_id"`

	// The id of the Line Item that the Fulfillment Item references.
	ItemId uuid.NullUUID `json:"item_id"`

	Item *LineItem `json:"item" gorm:"foreignKey:id;references:item_id"`

	// The quantity of the Line Item that is included in the Fulfillment.
	Quantity int32 `json:"quantity"`
}

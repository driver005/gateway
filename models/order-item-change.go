package models

import (
	"database/sql/driver"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// OrderItemChange - Represents an order edit item change
type OrderItemChange struct {
	core.Model

	// The order's status
	Type OrderEditItemChangeType `json:"type"`

	// The ID of the order edit
	OrderEditId uuid.NullUUID `json:"order_edit_id"`

	OrderEdit *OrderEdit `json:"order_edit" gorm:"foreignKey:id;references:order_edit_id"`

	// The ID of the original line item in the order
	OriginalLineItemId uuid.NullUUID `json:"original_line_item_id" gorm:"default:null"`

	OriginalLineItem *LineItem `json:"original_line_item" gorm:"foreignKey:id;references:original_line_item_id"`

	// The ID of the cloned line item.
	LineItemId uuid.NullUUID `json:"line_item_id" gorm:"default:null"`

	LineItem *LineItem `json:"line_item" gorm:"foreignKey:id;references:line_item_id"`
}

type OrderEditItemChangeType string

const (
	OrderEditStatusItemAdd    OrderEditItemChangeType = "item_add"
	OrderEditStatusItemRemove OrderEditItemChangeType = "item_remove"
	OrderEditStatusItemUpdate OrderEditItemChangeType = "item_update"
)

func (pl *OrderEditItemChangeType) Scan(value interface{}) error {
	*pl = OrderEditItemChangeType(value.([]byte))
	return nil
}

func (pl OrderEditItemChangeType) Value() (driver.Value, error) {
	return string(pl), nil
}

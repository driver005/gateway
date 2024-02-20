package models

import (
	"database/sql/driver"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:OrderItemChange
// title: "Order Item Change"
// description: "An order item change is a change made within an order edit to an order's items. These changes are not reflected on the original order until the order edit is confirmed."
// type: object
// required:
//   - created_at
//   - deleted_at
//   - id
//   - line_item_id
//   - order_edit_id
//   - original_line_item_id
//   - type
//   - updated_at
//
// properties:
//
//	id:
//	  description: The order item change's ID
//	  type: string
//	  example: oic_01G8TJSYT9M6AVS5N4EMNFS1EK
//	type:
//	  description: The order item change's status
//	  type: string
//	  enum:
//	    - item_add
//	    - item_remove
//	    - item_update
//	order_edit_id:
//	  description: The ID of the order edit
//	  type: string
//	  example: oe_01G2SG30J8C85S4A5CHM2S1NS2
//	order_edit:
//	  description: The details of the order edit the item change is associated with.
//	  x-expandable: "order_edit"
//	  nullable: true
//	  $ref: "#/components/schemas/OrderEdit"
//	original_line_item_id:
//	   description: The ID of the original line item in the order
//	   nullable: true
//	   type: string
//	   example: item_01G8ZC9GWT6B2GP5FSXRXNFNGN
//	original_line_item:
//	   description: The details of the original line item this item change references. This is used if the item change updates or deletes the original item.
//	   x-expandable: "original_line_item"
//	   nullable: true
//	   $ref: "#/components/schemas/LineItem"
//	line_item_id:
//	   description: The ID of the cloned line item.
//	   nullable: true
//	   type: string
//	   example: item_01G8ZC9GWT6B2GP5FSXRXNFNGN
//	line_item:
//	   description: The details of the resulting line item after the item change. This line item is then used in the original order once the order edit is confirmed.
//	   x-expandable: "line_item"
//	   nullable: true
//	   $ref: "#/components/schemas/LineItem"
//	created_at:
//	  description: The date with timezone at which the resource was created.
//	  type: string
//	  format: date-time
//	updated_at:
//	  description: The date with timezone at which the resource was updated.
//	  type: string
//	  format: date-time
//	deleted_at:
//	  description: The date with timezone at which the resource was deleted.
//	  nullable: true
//	  type: string
//	  format: date-time
type OrderItemChange struct {
	core.Model

	Type               OrderEditItemChangeType `json:"type"  gorm:"column:type"`
	OrderEditId        uuid.NullUUID           `json:"order_edit_id"  gorm:"column:order_edit_id"`
	OrderEdit          *OrderEdit              `json:"order_edit"  gorm:"column:order_edit;foreignKey:OrderEditId"`
	OriginalLineItemId uuid.NullUUID           `json:"original_line_item_id"  gorm:"column:original_line_item_id"`
	OriginalLineItem   *LineItem               `json:"original_line_item"  gorm:"column:original_line_item;foreignKey:OriginalLineItemId"`
	LineItemId         uuid.NullUUID           `json:"line_item_id"  gorm:"column:line_item_id"`
	LineItem           *LineItem               `json:"line_item"  gorm:"column:line_item;foreignKey:LineItemId"`
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

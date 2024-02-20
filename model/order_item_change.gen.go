// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameOrderItemChange = "order_item_change"

// OrderItemChange mapped from table <order_item_change>
type OrderItemChange struct {
	ID                 string         `gorm:"column:id;type:character varying;primaryKey" json:"id"`
	CreatedAt          time.Time      `gorm:"column:created_at;type:timestamp with time zone;not null;default:now()" json:"created_at"`
	UpdatedAt          time.Time      `gorm:"column:updated_at;type:timestamp with time zone;not null;default:now()" json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone" json:"deleted_at"`
	Type               string         `gorm:"column:type;type:order_item_change_type_enum;not null" json:"type"`
	OrderEditID        string         `gorm:"column:order_edit_id;type:character varying;not null;uniqueIndex:UQ_da93cee3ca0dd50a5246268c2e9,priority:1;uniqueIndex:UQ_5b7a99181e4db2ea821be0b6196,priority:1" json:"order_edit_id"`
	OriginalLineItemID string         `gorm:"column:original_line_item_id;type:character varying;uniqueIndex:UQ_5b7a99181e4db2ea821be0b6196,priority:2" json:"original_line_item_id"`
	LineItemID         string         `gorm:"column:line_item_id;type:character varying;uniqueIndex:UQ_da93cee3ca0dd50a5246268c2e9,priority:2;uniqueIndex:REL_5f9688929761f7df108b630e64,priority:1" json:"line_item_id"`
}

// TableName OrderItemChange's table name
func (*OrderItemChange) TableName() string {
	return TableNameOrderItemChange
}

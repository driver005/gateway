package models

import "github.com/driver005/gateway/core"

type OrderChangeAction struct {
	core.BaseModel

	OrderChangeId string       `gorm:"column:order_change_id;type:text;not null;index:IDX_order_change_action_order_change_id,priority:1" json:"order_change_id"`
	OrderChange   *OrderChange `gorm:"foreignkey:OrderChangeID" json:"order_change"`
	Reference     string       `gorm:"column:reference;type:text;not null" json:"reference"`
	ReferenceId   string       `gorm:"column:reference_id;type:text;not null;index:IDX_order_change_action_reference_id,priority:1" json:"reference_id"`
	Action        core.JSONB   `gorm:"column:action;type:jsonb;not null" json:"action"`
	InternalNote  string       `gorm:"column:internal_note;type:text" json:"internal_note"`
}

func (*OrderChangeAction) TableName() string {
	return "order_change_action"
}

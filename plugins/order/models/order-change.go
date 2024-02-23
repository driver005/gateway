package models

import (
	"time"

	"github.com/driver005/gateway/core"
)

type OrderChange struct {
	core.BaseModel

	OrderId        string    `gorm:"column:order_id;type:text;not null;index:IDX_order_change_order_id,priority:1" json:"order_id"`
	Order          *Order    `gorm:"foreignkey:OrderId" json:"order"`
	Description    string    `gorm:"column:description;type:text" json:"description"`
	Status         string    `gorm:"column:status;type:text;not null;index:IDX_order_change_status,priority:1;default:pending" json:"status"`
	InternalNote   string    `gorm:"column:internal_note;type:text" json:"internal_note"`
	CreatedBy      string    `gorm:"column:created_by;type:text;not null" json:"created_by"`
	RequestedBy    string    `gorm:"column:requested_by;type:text" json:"requested_by"`
	RequestedAt    time.Time `gorm:"column:requested_at;type:timestamp with time zone" json:"requested_at"`
	ConfirmedBy    string    `gorm:"column:confirmed_by;type:text" json:"confirmed_by"`
	ConfirmedAt    time.Time `gorm:"column:confirmed_at;type:timestamp with time zone" json:"confirmed_at"`
	DeclinedBy     string    `gorm:"column:declined_by;type:text" json:"declined_by"`
	DeclinedReason string    `gorm:"column:declined_reason;type:text" json:"declined_reason"`
	DeclinedAt     time.Time `gorm:"column:declined_at;type:timestamp with time zone" json:"declined_at"`
	CanceledBy     string    `gorm:"column:canceled_by;type:text" json:"canceled_by"`
	CanceledAt     time.Time `gorm:"column:canceled_at;type:timestamp with time zone" json:"canceled_at"`
}

func (*OrderChange) TableName() string {
	return "order_change"
}

// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameFulfillment = "fulfillment"

// Fulfillment mapped from table <fulfillment>
type Fulfillment struct {
	ID              string    `gorm:"column:id;type:character varying;primaryKey" json:"id"`
	SwapID          string    `gorm:"column:swap_id;type:character varying;index:IDX_a52e234f729db789cf473297a5,priority:1" json:"swap_id"`
	OrderID         string    `gorm:"column:order_id;type:character varying;index:IDX_f129acc85e346a10eed12b86fc,priority:1" json:"order_id"`
	TrackingNumbers string    `gorm:"column:tracking_numbers;type:jsonb;not null;default:[]" json:"tracking_numbers"`
	Data            string    `gorm:"column:data;type:jsonb;not null" json:"data"`
	ShippedAt       time.Time `gorm:"column:shipped_at;type:timestamp with time zone" json:"shipped_at"`
	CanceledAt      time.Time `gorm:"column:canceled_at;type:timestamp with time zone" json:"canceled_at"`
	CreatedAt       time.Time `gorm:"column:created_at;type:timestamp with time zone;not null;default:now()" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;type:timestamp with time zone;not null;default:now()" json:"updated_at"`
	Metadata        string    `gorm:"column:metadata;type:jsonb" json:"metadata"`
	IdempotencyKey  string    `gorm:"column:idempotency_key;type:character varying" json:"idempotency_key"`
	ProviderID      string    `gorm:"column:provider_id;type:character varying;index:IDX_beb35a6de60a6c4f91d5ae57e4,priority:1" json:"provider_id"`
	ClaimOrderID    string    `gorm:"column:claim_order_id;type:character varying;index:IDX_d73e55964e0ff2db8f03807d52,priority:1" json:"claim_order_id"`
	NoNotification  bool      `gorm:"column:no_notification;type:boolean" json:"no_notification"`
	LocationID      string    `gorm:"column:location_id;type:character varying" json:"location_id"`
}

// TableName Fulfillment's table name
func (*Fulfillment) TableName() string {
	return TableNameFulfillment
}
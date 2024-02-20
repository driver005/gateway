// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameShippingOption = "shipping_option"

// ShippingOption mapped from table <shipping_option>
type ShippingOption struct {
	ID         string         `gorm:"column:id;type:character varying;primaryKey" json:"id"`
	Name       string         `gorm:"column:name;type:character varying;not null" json:"name"`
	RegionID   string         `gorm:"column:region_id;type:character varying;not null;index:IDX_5c58105f1752fca0f4ce69f466,priority:1" json:"region_id"`
	ProfileID  string         `gorm:"column:profile_id;type:character varying;not null;index:IDX_c951439af4c98bf2bd7fb8726c,priority:1" json:"profile_id"`
	ProviderID string         `gorm:"column:provider_id;type:character varying;not null;index:IDX_a0e206bfaed3cb63c186091734,priority:1" json:"provider_id"`
	PriceType  string         `gorm:"column:price_type;type:shipping_option_price_type_enum;not null" json:"price_type"`
	Amount     int32          `gorm:"column:amount;type:integer" json:"amount"`
	IsReturn   bool           `gorm:"column:is_return;type:boolean;not null" json:"is_return"`
	Data       string         `gorm:"column:data;type:jsonb;not null" json:"data"`
	CreatedAt  time.Time      `gorm:"column:created_at;type:timestamp with time zone;not null;default:now()" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;type:timestamp with time zone;not null;default:now()" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone" json:"deleted_at"`
	Metadata   string         `gorm:"column:metadata;type:jsonb" json:"metadata"`
	AdminOnly  bool           `gorm:"column:admin_only;type:boolean;not null" json:"admin_only"`
}

// TableName ShippingOption's table name
func (*ShippingOption) TableName() string {
	return TableNameShippingOption
}

package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type ShippingProfile struct {
	core.BaseModel

	Name            string           `gorm:"column:name;type:text;not null;uniqueIndex:IDX_shipping_profile_name_unique,priority:1" json:"name"`
	Type            string           `gorm:"column:type;type:text;not null" json:"type"`
	ShippingOptions []ShippingOption `gorm:"foreignKey:ShippingProfileId" json:"shipping_options"`
	DeletedAt       gorm.DeletedAt   `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_shipping_profile_deleted_at,priority:1" json:"deleted_at"`
}

func (*ShippingProfile) TableName() string {
	return "shipping_profile"
}

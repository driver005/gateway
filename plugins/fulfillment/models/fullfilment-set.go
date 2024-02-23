package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type FulfillmentSet struct {
	core.BaseModel

	Name         string         `gorm:"column:name;type:text;not null;uniqueIndex:IDX_fulfillment_set_name_unique,priority:1" json:"name"`
	Type         string         `gorm:"column:type;type:text;not null" json:"type"`
	ServiceZones []ServiceZone  `gorm:"foreignKey:FulfillmentSetId;references:ID;constraint:OnDelete:CASCADE" json:"service_zones"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_fulfillment_set_deleted_at,priority:1" json:"deleted_at"`
}

func (*FulfillmentSet) TableName() string {
	return "fulfillment_set"
}

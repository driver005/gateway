package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type ServiceZone struct {
	core.BaseModel

	Name             string           `gorm:"column:name;type:text;not null;uniqueIndex:IDX_service_zone_name_unique,priority:1" json:"name"`
	FulfillmentSetId string           `gorm:"column:fulfillment_set_id;type:text;not null;index:IDX_service_zone_fulfillment_set_id,priority:1" json:"fulfillment_set_id"`
	FulfillmentSet   *FulfillmentSet  `gorm:"foreignKey:FulfillmentSetId" json:"fulfillment_set"`
	GeoZones         []GeoZone        `gorm:"foreignKey:ServiceZoneId" json:"geo_zones"`
	ShippingOptions  []ShippingOption `gorm:"foreignKey:ServiceZoneId" json:"shipping_options"`
	DeletedAt        gorm.DeletedAt   `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_service_zone_deleted_at,priority:1" json:"deleted_at"`
}

func (*ServiceZone) TableName() string {
	return "service_zone"
}

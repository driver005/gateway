package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type GeoZone struct {
	core.BaseModel

	Type             string         `gorm:"column:type;type:text;not null;default:country" json:"type"`
	CountryCode      string         `gorm:"column:country_code;type:text;not null;index:IDX_geo_zone_country_code,priority:1" json:"country_code"`
	ProvinceCode     string         `gorm:"column:province_code;type:text;index:IDX_geo_zone_province_code,priority:1" json:"province_code"`
	City             string         `gorm:"column:city;type:text;index:IDX_geo_zone_city,priority:1" json:"city"`
	ServiceZoneId    string         `gorm:"column:service_zone_id;type:text;not null;index:IDX_geo_zone_service_zone_id,priority:1" json:"service_zone_id"`
	PostalExpression string         `gorm:"column:postal_expression;type:jsonb" json:"postal_expression"`
	ServiceZone      *ServiceZone   `gorm:"foreignKey:ServiceZoneId;references:ID;constraint:OnDelete:CASCADE" json:"service_zone"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_geo_zone_deleted_at,priority:1" json:"deleted_at"`
}

func (*GeoZone) TableName() string {
	return "geo_zone"
}

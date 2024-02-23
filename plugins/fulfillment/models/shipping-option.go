package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type ShippingOption struct {
	core.BaseModel

	Name                 string               `gorm:"column:name;type:text;not null" json:"name"`
	PriceType            string               `gorm:"column:price_type;type:text;not null;default:calculated" json:"price_type"`
	ServiceZoneId        string               `gorm:"column:service_zone_id;type:text;not null;index:IDX_shipping_option_service_zone_id,priority:1" json:"service_zone_id"`
	ShippingProfileId    string               `gorm:"column:shipping_profile_id;type:text;index:IDX_shipping_option_shipping_profile_id,priority:1" json:"shipping_profile_id"`
	ServiceProviderId    string               `gorm:"column:service_provider_id;type:text;index:IDX_shipping_option_service_provider_id,priority:1" json:"service_provider_id"`
	ShippingOptionTypeId string               `gorm:"column:shipping_option_type_id;type:text;uniqueIndex:shipping_option_shipping_option_type_id_unique,priority:1;index:IDX_shipping_option_shipping_option_type_id,priority:1" json:"shipping_option_type_id"`
	Data                 string               `gorm:"column:data;type:jsonb" json:"data"`
	ServiceZone          *ServiceZone         `gorm:"foreignKey:ServiceZoneId" json:"service_zone"`
	ShippingProfile      *ShippingProfile     `gorm:"foreignKey:ShippingProfileId" json:"shipping_profile"`
	ServiceProvider      *ServiceProvider     `gorm:"foreignKey:ServiceProviderId" json:"service_provider"`
	ShippingOptionType   *ShippingOptionType  `gorm:"foreignKey:ShippingOptionTypeId" json:"shipping_option_type"`
	Rules                []ShippingOptionRule `gorm:"foreignKey:ShippingOptionId" json:"rules"`
	Fulfillments         []Fulfillment        `gorm:"foreignKey:ShippingOptionId" json:"fulfillments"`
	DeletedAt            gorm.DeletedAt       `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_shipping_option_deleted_at,priority:1" json:"deleted_at"`
}

func (*ShippingOption) TableName() string {
	return "shipping_option"
}

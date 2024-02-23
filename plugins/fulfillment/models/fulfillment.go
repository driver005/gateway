package models

import (
	"time"

	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type Fulfillment struct {
	core.BaseModel

	LocationId        string             `gorm:"column:location_id;type:text;not null;index:IDX_fulfillment_location_id,priority:1" json:"location_id"`
	PackedAt          time.Time          `gorm:"column:packed_at;type:timestamp with time zone" json:"packed_at"`
	ShippedAt         time.Time          `gorm:"column:shipped_at;type:timestamp with time zone" json:"shipped_at"`
	DeliveredAt       time.Time          `gorm:"column:delivered_at;type:timestamp with time zone" json:"delivered_at"`
	CanceledAt        time.Time          `gorm:"column:canceled_at;type:timestamp with time zone" json:"canceled_at"`
	Data              string             `gorm:"column:data;type:jsonb" json:"data"`
	ProviderId        string             `gorm:"column:provider_id;type:text;not null;index:IDX_fulfillment_provider_id,priority:1" json:"provider_id"`
	ShippingOptionId  string             `gorm:"column:shipping_option_id;type:text;index:IDX_fulfillment_shipping_option_id,priority:1" json:"shipping_option_id"`
	DeliveryAddressId string             `gorm:"column:delivery_address_id;type:text;not null" json:"delivery_address_id"`
	ItemsId           string             `gorm:"column:items_id;type:text;not null" json:"items_id"`
	ShippingOption    *ShippingOption    `gorm:"foreignKey:ShippingOptionId;references:ID;constraint:OnDelete:CASCADE" json:"shipping_option"`
	Provider          *ServiceProvider   `gorm:"foreignKey:ProviderId;references:ID;constraint:OnDelete:CASCADE" json:"provider"`
	DeliveryAddress   *Address           `gorm:"foreignKey:DeliveryAddressId;references:ID;constraint:OnDelete:CASCADE" json:"delivery_address"`
	Items             []FulfillmentItem  `gorm:"foreignKey:FulfillmentId;references:ID;constraint:OnDelete:CASCADE" json:"items"`
	Labels            []FulfillmentLabel `gorm:"foreignKey:FulfillmentId;references:ID;constraint:OnDelete:CASCADE" json:"labels"`
	DeletedAt         gorm.DeletedAt     `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_fulfillment_deleted_at,priority:1" json:"deleted_at"`
}

func (*Fulfillment) TableName() string {
	return "fulfillment"
}

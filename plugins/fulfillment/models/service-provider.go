package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type ServiceProvider struct {
	core.BaseModel

	ShippingOptions []ShippingOption `gorm:"foreignKey:ServiceProviderId" json:"shipping_options"`
	DeletedAt       gorm.DeletedAt   `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_service_provider_deleted_at,priority:1" json:"deleted_at"`
}

func (*ServiceProvider) TableName() string {
	return "service_provider"
}

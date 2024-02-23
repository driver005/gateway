package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductShippingProfile struct {
	ProfileId uuid.NullUUID  `gorm:"column:profile_id;type:text;not null;uniqueIndex:idx_product_shipping_profile_profile_id_product_id_unique,priority:1;index:idx_product_shipping_profile_profile_id,priority:1" json:"profile_id"`
	ProductId uuid.NullUUID  `gorm:"column:product_id;type:text;not null;uniqueIndex:idx_product_shipping_profile_profile_id_product_id_unique,priority:2;index:idx_product_shipping_profile_product_id,priority:1" json:"product_id"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp with time zone;not null;default:now()" json:"updated_at"`
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp with time zone;not null;default:now()" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone" json:"deleted_at"`
}

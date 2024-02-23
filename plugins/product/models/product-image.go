package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type ProductImage struct {
	core.BaseModel

	ID        string         `gorm:"column:id;type:text;primaryKey" json:"id"`
	URL       string         `gorm:"column:url;type:text;not null;index:IDX_product_image_url,priority:1" json:"url"`
	Products  []Product      `gorm:"many2many:product_images;" json:"products"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_product_image_deleted_at,priority:1" json:"deleted_at"`
}

func (*ProductImage) TableName() string {
	return "image"
}

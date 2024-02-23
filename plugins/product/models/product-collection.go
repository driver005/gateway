package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type ProductCollection struct {
	core.BaseModel

	Title     string         `gorm:"column:title;type:text;not null" json:"title"`
	Handle    string         `gorm:"column:handle;type:text;not null;uniqueIndex:IDX_product_collection_handle_unique,priority:1" json:"handle"`
	Products  []Product      `gorm:"foreignKey:CollectionId" json:"products"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_product_collection_deleted_at,priority:1" json:"deleted_at"`
}

func (*ProductCollection) TableName() string {
	return "product_collection"
}

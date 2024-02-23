package models

import "github.com/google/uuid"

type ProductImage struct {
	ProductId uuid.NullUUID `gorm:"column:product_id;type:character varying;primaryKey;index:IDX_4f166bb8c2bfcef2498d97b406,priority:1" json:"product_id"`
	ImageId   uuid.NullUUID `gorm:"column:image_id;type:character varying;primaryKey;index:IDX_2212515ba306c79f42c46a99db,priority:1" json:"image_id"`
}

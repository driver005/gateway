// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameProductTag = "product_tag"

// ProductTag mapped from table <product_tag>
type ProductTag struct {
	ID        string         `gorm:"column:id;type:character varying;primaryKey" json:"id"`
	Value     string         `gorm:"column:value;type:character varying;not null" json:"value"`
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp with time zone;not null;default:now()" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp with time zone;not null;default:now()" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone" json:"deleted_at"`
	Metadata  string         `gorm:"column:metadata;type:jsonb" json:"metadata"`
}

// TableName ProductTag's table name
func (*ProductTag) TableName() string {
	return TableNameProductTag
}

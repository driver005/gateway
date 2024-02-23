package models

import (
	"time"

	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type PriceList struct {
	core.BaseModel

	Status      string         `gorm:"column:status;type:text;not null;default:draft" json:"status"`
	StartsAt    time.Time      `gorm:"column:starts_at;type:timestamp with time zone" json:"starts_at"`
	EndsAt      time.Time      `gorm:"column:ends_at;type:timestamp with time zone" json:"ends_at"`
	RulesCount  int32          `gorm:"column:rules_count;type:integer;not null" json:"rules_count"`
	Title       string         `gorm:"column:title;type:text;not null" json:"title"`
	Name        string         `gorm:"column:name;type:text" json:"name"`
	Description string         `gorm:"column:description;type:text;not null" json:"description"`
	Type        string         `gorm:"column:type;type:text;not null;default:sale" json:"type"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_price_list_deleted_at,priority:1" json:"deleted_at"`
}

func (*PriceList) TableName() string {
	return "price_list"
}

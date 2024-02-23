package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type Region struct {
	core.BaseModel

	Name         string         `gorm:"column:name;type:text;not null" json:"name"`
	CurrencyCode string         `gorm:"column:currency_code;type:text;not null;index:IDX_region_currency_code,priority:1" json:"currency_code"`
	Currency     *Currency      `gorm:"foreignKey:CurrencyCode" json:"currency"`
	Countries    []Country      `gorm:"foreignKey:RegionId" json:"countries"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_region_deleted_at,priority:1" json:"deleted_at"`
}

func (*Region) TableName() string {
	return "region"
}

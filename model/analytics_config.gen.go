// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameAnalyticsConfig = "analytics_config"

// AnalyticsConfig mapped from table <analytics_config>
type AnalyticsConfig struct {
	ID        string         `gorm:"column:id;type:character varying;primaryKey;uniqueIndex:IDX_379ca70338ce9991f3affdeedf,priority:1" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp with time zone;not null;default:now()" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp with time zone;not null;default:now()" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone" json:"deleted_at"`
	UserID    string         `gorm:"column:user_id;type:character varying;not null;uniqueIndex:IDX_379ca70338ce9991f3affdeedf,priority:2" json:"user_id"`
	OptOut    bool           `gorm:"column:opt_out;type:boolean;not null" json:"opt_out"`
	Anonymize bool           `gorm:"column:anonymize;type:boolean;not null" json:"anonymize"`
}

// TableName AnalyticsConfig's table name
func (*AnalyticsConfig) TableName() string {
	return TableNameAnalyticsConfig
}
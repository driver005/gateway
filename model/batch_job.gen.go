// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameBatchJob = "batch_job"

// BatchJob mapped from table <batch_job>
type BatchJob struct {
	ID             string         `gorm:"column:id;type:character varying;primaryKey" json:"id"`
	Type           string         `gorm:"column:type;type:text;not null" json:"type"`
	CreatedBy      string         `gorm:"column:created_by;type:character varying" json:"created_by"`
	Context        string         `gorm:"column:context;type:jsonb" json:"context"`
	Result         string         `gorm:"column:result;type:jsonb" json:"result"`
	DryRun         bool           `gorm:"column:dry_run;type:boolean;not null" json:"dry_run"`
	CreatedAt      time.Time      `gorm:"column:created_at;type:timestamp with time zone;not null;default:now()" json:"created_at"`
	PreProcessedAt time.Time      `gorm:"column:pre_processed_at;type:timestamp with time zone" json:"pre_processed_at"`
	ConfirmedAt    time.Time      `gorm:"column:confirmed_at;type:timestamp with time zone" json:"confirmed_at"`
	ProcessingAt   time.Time      `gorm:"column:processing_at;type:timestamp with time zone" json:"processing_at"`
	CompletedAt    time.Time      `gorm:"column:completed_at;type:timestamp with time zone" json:"completed_at"`
	FailedAt       time.Time      `gorm:"column:failed_at;type:timestamp with time zone" json:"failed_at"`
	CanceledAt     time.Time      `gorm:"column:canceled_at;type:timestamp with time zone" json:"canceled_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;type:timestamp with time zone;not null;default:now()" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone" json:"deleted_at"`
}

// TableName BatchJob's table name
func (*BatchJob) TableName() string {
	return TableNameBatchJob
}

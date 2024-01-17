package models

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/helper"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// BatchJob - A Batch Job.
type BatchJob struct {
	core.Model

	// The type of batch job.
	Type string `json:"type"`

	// The status of the batch job.
	Status BatchJobStatus `json:"status" gorm:"default:created"`

	// The unique identifier of the user that created the batch job.
	CreatedBy uuid.UUID `json:"created_by" gorm:"default:null"`

	// A user object. Available if the relation `created_by_user` is expanded.
	CreatedByUser *User `json:"created_by_user" gorm:"foreignKey:id"`

	// The context of the batch job, the type of the batch job determines what the context should contain.
	Context core.JSONB `json:"context" gorm:"default:null"`

	// Specify if the job must apply the modifications or not.
	DryRun bool `json:"dry_run" gorm:"default:null"`

	Result *BatchJobResult `json:"result" gorm:"default:null"`

	// The date from which the job has been pre processed.
	PreProcessedAt *time.Time `json:"pre_processed_at" gorm:"default:null"`

	// The date the job is processing at.
	ProcessingAt *time.Time `json:"processing_at" gorm:"default:null"`

	// The date when the confirmation has been done.
	ConfirmedAt *time.Time `json:"confirmed_at" gorm:"default:null"`

	// The date of the completion.
	CompletedAt *time.Time `json:"completed_at" gorm:"default:null"`

	// The date of the concellation.
	CanceledAt *time.Time `json:"canceled_at" gorm:"default:null"`

	// The date when the job failed.
	FailedAt *time.Time `json:"failed_at" gorm:"default:null"`
}

type BatchJobResultErrorsCode struct {
	Message string `json:"message,omitempty" gorm:"default:null"`

	Code string `json:"code,omitempty" gorm:"default:null"`
}

type BatchJobResultErrors struct {
	Message string `json:"message,omitempty" gorm:"default:null"`

	Code *BatchJobResultErrorsCode `json:"code,omitempty" gorm:"default:null"`

	Err []string `json:"err,omitempty" gorm:"default:null"`
}

type BatchJobResultStatDescriptors struct {
	Key string `json:"key,omitempty" gorm:"default:null"`

	Name string `json:"name,omitempty" gorm:"default:null"`

	Message string `json:"message,omitempty" gorm:"default:null"`
}

// BatchJobResult - The result of the batch job.
type BatchJobResult struct {
	Count float64 `json:"count,omitempty" gorm:"default:null"`

	AdvancementCount float64 `json:"advancement_count,omitempty" gorm:"default:null"`

	Progress float64 `json:"progress,omitempty" gorm:"default:null"`

	Errors *BatchJobResultErrors `json:"errors,omitempty" gorm:"default:null"`

	StatDescriptors *BatchJobResultStatDescriptors `json:"stat_descriptors,omitempty" gorm:"default:null"`

	FileKey string `json:"file_key,omitempty" gorm:"default:null"`

	FileSize float64 `json:"file_size,omitempty" gorm:"default:null"`
}

func (b BatchJobResult) Interface() interface{} {
	return BatchJobResult(b)
}

func (BatchJobResult) GormDataType() string {
	return "result"
}

func (b BatchJobResult) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL: "ROW(?, ?, ?, ?, ?, ?, ?, ?)::result",
		Vars: []interface{}{
			b.Count,
			b.AdvancementCount,
			b.Progress,
			b.Errors,
			b.StatDescriptors,
			b.FileKey,
			b.FileSize,
		},
	}
}

func (b *BatchJobResult) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	s := strings.Split(strings.Trim(string(bytes), "()"), ",")
	for i := 0; i < reflect.TypeOf(BatchJobResult{}).NumField(); i++ {
		helper.SetField(&b, reflect.TypeOf(BatchJobResult{}).Field(i).Name, s[i])
	}

	return nil
}

// The status of the Price List
type BatchJobStatus string

// Defines values for BatchJobStatus.
const (
	BatchJobStatusCreated      BatchJobStatus = "created"
	BatchJobStatusPreProcessed BatchJobStatus = "pre_processed"
	BatchJobStatusConfirmed    BatchJobStatus = "confirmed"
	BatchJobStatusProcessing   BatchJobStatus = "processing"
	BatchJobStatusCompleted    BatchJobStatus = "completed"
	BatchJobStatusCanceled     BatchJobStatus = "canceled"
	BatchJobStatusFailed       BatchJobStatus = "failed"
)

func (pl *BatchJobStatus) Scan(value interface{}) error {
	*pl = BatchJobStatus(value.([]byte))
	return nil
}

func (pl BatchJobStatus) Value() (driver.Value, error) {
	return string(pl), nil
}

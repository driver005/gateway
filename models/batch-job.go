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
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//
// @oas:schema:BatchJob
// title: "Batch Job"
// description: "A Batch Job indicates an asynchronus task stored in the Medusa backend. Its status determines whether it has been executed or not."
// type: object
// required:
//   - canceled_at
//   - completed_at
//   - confirmed_at
//   - context
//   - created_at
//   - created_by
//   - deleted_at
//   - dry_run
//   - failed_at
//   - id
//   - pre_processed_at
//   - processing_at
//   - result
//   - status
//   - type
//   - updated_at
// properties:
//  id:
//    description: The unique identifier for the batch job.
//    type: string
//    example: batch_01G8T782965PYFG0751G0Z38B4
//  type:
//    description: The type of batch job.
//    type: string
//    enum:
//      - product-import
//      - product-export
//  status:
//    description: The status of the batch job.
//    type: string
//    enum:
//      - created
//      - pre_processed
//      - confirmed
//      - processing
//      - completed
//      - canceled
//      - failed
//    default: created
//  created_by:
//    description: The unique identifier of the user that created the batch job.
//    nullable: true
//    type: string
//    example: usr_01G1G5V26F5TB3GPAPNJ8X1S3V
//  created_by_user:
//    description: The details of the user that created the batch job.
//    x-expandable: "created_by_user"
//    nullable: true
//    $ref: "#/components/schemas/User"
//  context:
//    description: The context of the batch job, the type of the batch job determines what the context should contain.
//    nullable: true
//    type: object
//    example:
//      shape:
//        prices:
//          - region: null
//            currency_code: "eur"
//        dynamicImageColumnCount: 4
//        dynamicOptionColumnCount: 2
//      list_config:
//        skip: 0
//        take: 50
//        order:
//          created_at: "DESC"
//        relations:
//          - variants
//          - variant.prices
//          - images
//  dry_run:
//    description: Specify if the job must apply the modifications or not.
//    type: boolean
//    default: false
//  result:
//    description: The result of the batch job.
//    nullable: true
//    allOf:
//    - type: object
//      example: {}
//    - type: object
//      properties:
//        count:
//          type: number
//        advancement_count:
//          type: number
//        progress:
//          type: number
//        errors:
//          type: object
//          properties:
//            message:
//              type: string
//            code:
//              oneOf:
//                - type: string
//                - type: number
//            err:
//              type: array
//        stat_descriptors:
//          type: object
//          properties:
//            key:
//              type: string
//            name:
//              type: string
//            message:
//              type: string
//        file_key:
//          type: string
//        file_size:
//          type: number
//    example:
//      errors:
//        - err: []
//          code: "unknown"
//          message: "Method not implemented."
//      stat_descriptors:
//        - key: "product-export-count"
//          name: "Product count to export"
//          message: "There will be 8 products exported by this action"
//  pre_processed_at:
//    description: The date from which the job has been pre-processed.
//    nullable: true
//    type: string
//    format: date-time
//  processing_at:
//    description: The date the job is processing at.
//    nullable: true
//    type: string
//    format: date-time
//  confirmed_at:
//    description: The date when the confirmation has been done.
//    nullable: true
//    type: string
//    format: date-time
//  completed_at:
//    description: The date of the completion.
//    nullable: true
//    type: string
//    format: date-time
//  canceled_at:
//    description: The date of the concellation.
//    nullable: true
//    type: string
//    format: date-time
//  failed_at:
//    description: The date when the job failed.
//    nullable: true
//    type: string
//    format: date-time
//  created_at:
//    description: The date with timezone at which the resource was created.
//    type: string
//    format: date-time
//  updated_at:
//    description: The date with timezone at which the resource was last updated.
//    type: string
//    format: date-time
//  deleted_at:
//    description: The date with timezone at which the resource was deleted.
//    nullable: true
//    type: string
//    format: date-time
//

type BatchJob struct {
	core.Model

	Type           string          `json:"type"`
	Status         BatchJobStatus  `json:"status" gorm:"default:created"`
	CreatedBy      uuid.UUID       `json:"created_by" gorm:"default:null"`
	CreatedByUser  *User           `json:"created_by_user" gorm:"foreignKey:id"`
	Context        core.JSONB      `json:"context" gorm:"default:null"`
	DryRun         bool            `json:"dry_run" gorm:"default:null"`
	Result         *BatchJobResult `json:"result" gorm:"default:null"`
	PreProcessedAt *time.Time      `json:"pre_processed_at" gorm:"default:null"`
	ProcessingAt   *time.Time      `json:"processing_at" gorm:"default:null"`
	ConfirmedAt    *time.Time      `json:"confirmed_at" gorm:"default:null"`
	CompletedAt    *time.Time      `json:"completed_at" gorm:"default:null"`
	CanceledAt     *time.Time      `json:"canceled_at" gorm:"default:null"`
	FailedAt       *time.Time      `json:"failed_at" gorm:"default:null"`
}

type BatchJobResultErrorsCode struct {
	Message string `json:"message,omitempty" gorm:"default:null"`
	Code    string `json:"code,omitempty" gorm:"default:null"`
}

type BatchJobResultErrors struct {
	Message string                    `json:"message,omitempty" gorm:"default:null"`
	Code    *BatchJobResultErrorsCode `json:"code,omitempty" gorm:"default:null"`
	Err     []string                  `json:"err,omitempty" gorm:"default:null"`
}

type BatchJobResultStatDescriptors struct {
	Key     string `json:"key,omitempty" gorm:"default:null"`
	Name    string `json:"name,omitempty" gorm:"default:null"`
	Message string `json:"message,omitempty" gorm:"default:null"`
}

// BatchJobResult - The result of the batch job.
type BatchJobResult struct {
	Count            float64                        `json:"count,omitempty" gorm:"default:null"`
	AdvancementCount float64                        `json:"advancement_count,omitempty" gorm:"default:null"`
	Progress         float64                        `json:"progress,omitempty" gorm:"default:null"`
	Errors           *BatchJobResultErrors          `json:"errors,omitempty" gorm:"default:null"`
	StatDescriptors  *BatchJobResultStatDescriptors `json:"stat_descriptors,omitempty" gorm:"default:null"`
	FileKey          string                         `json:"file_key,omitempty" gorm:"default:null"`
	FileSize         float64                        `json:"file_size,omitempty" gorm:"default:null"`
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
		return errors.New(fmt.Sprint("Failed to unmarshal core.JSONB value:", value))
	}
	s := strings.Split(strings.Trim(string(bytes), "()"), ",")
	for i := 0; i < reflect.TypeOf(BatchJobResult{}).NumField(); i++ {
		setField(&b, reflect.TypeOf(BatchJobResult{}).Field(i).Name, s[i])
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

func setField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := reflect.Indirect(structValue).FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("no such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		invalidTypeError := errors.New("provided value type didn't match obj field type")
		return invalidTypeError
	}

	structFieldValue.Set(val)
	return nil
}

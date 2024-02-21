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
	core.SoftDeletableModel

	Type           string          `json:"type" gorm:"column:type"`
	Status         BatchJobStatus  `json:"status" gorm:"column:status;default:'created'"`
	CreatedBy      uuid.NullUUID   `json:"created_by" gorm:"column:created_by"`
	CreatedByUser  *User           `json:"created_by_user" gorm:"foreignKey:CreatedBy"`
	Context        core.JSONB      `json:"context" gorm:"column:context"`
	DryRun         bool            `json:"dry_run" gorm:"column:dry_run;default:false"`
	Result         *BatchJobResult `json:"result" gorm:"column:result"`
	PreProcessedAt *time.Time      `json:"pre_processed_at" gorm:"column:pre_processed_at"`
	ProcessingAt   *time.Time      `json:"processing_at" gorm:"column:processing_at"`
	ConfirmedAt    *time.Time      `json:"confirmed_at" gorm:"column:confirmed_at"`
	CompletedAt    *time.Time      `json:"completed_at" gorm:"column:completed_at"`
	CanceledAt     *time.Time      `json:"canceled_at" gorm:"column:canceled_at"`
	FailedAt       *time.Time      `json:"failed_at" gorm:"column:failed_at"`
}

type BatchJobResultErrorsCode struct {
	Message string `json:"message" gorm:"column:message"`
	Code    string `json:"code" gorm:"column:code"`
}

type BatchJobResultErrors struct {
	Message string                    `json:"message" gorm:"column:message"`
	Code    *BatchJobResultErrorsCode `json:"code" gorm:"column:code"`
	Err     []string                  `json:"err" gorm:"column:err"`
}

type BatchJobResultStatDescriptors struct {
	Key     string `json:"key" gorm:"column:key"`
	Name    string `json:"name" gorm:"column:name"`
	Message string `json:"message" gorm:"column:message"`
}

// BatchJobResult - The result of the batch job.
type BatchJobResult struct {
	Count            float64                        `json:"count" gorm:"column:count"`
	AdvancementCount float64                        `json:"advancement_count" gorm:"column:advancement_count"`
	Progress         float64                        `json:"progress" gorm:"column:progress"`
	Errors           *BatchJobResultErrors          `json:"errors" gorm:"column:errors"`
	StatDescriptors  *BatchJobResultStatDescriptors `json:"stat_descriptors" gorm:"column:stat_descriptors"`
	FileKey          string                         `json:"file_key" gorm:"column:file_key"`
	FileSize         float64                        `json:"file_size" gorm:"column:file_size"`
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

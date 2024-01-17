package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type BatchJobUpdateProps struct {
	Context core.JSONB             `json:"context,omitempty" validate:"omitempty"`
	Result  *models.BatchJobResult `json:"result,omitempty" validate:"omitempty"`
}

type CreateBatchJobInput struct {
	Type    string     `json:"type"`
	Context core.JSONB `json:"context,omitempty" validate:"omitempty"`
	DryRun  bool       `json:"dry_run"`
}

type BatchJobResultError struct {
	Message string                           `json:"message"`
	Code    *models.BatchJobResultErrorsCode `json:"code"`
	// Extra   core.JSONB           `json:"extra"`
}

type BatchJobResultStatDescriptor struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

type BatchJobCreateProps struct {
	Context   core.JSONB `json:"context"`
	Type      string     `json:"type"`
	CreatedBy uuid.UUID  `json:"created_by"`
	DryRun    bool       `json:"dry_run"`
}

type FilterableBatchJob struct {
	core.FilterModel

	Status         []models.BatchJobStatus `json:"status,omitempty" validate:"omitempty"`
	Type           []string                `json:"type,omitempty" validate:"omitempty"`
	CreatedBy      uuid.UUIDs              `json:"created_at,omitempty" validate:"omitempty"`
	ConfirmedAt    *core.TimeModel         `json:"confirmed_at,omitempty" validate:"omitempty"`
	PreProcessedAt *core.TimeModel         `json:"pre_processed_at,omitempty" validate:"omitempty"`
	CompletedAt    *core.TimeModel         `json:"completed_at,omitempty" validate:"omitempty"`
	FailedAt       *core.TimeModel         `json:"failed_at,omitempty" validate:"omitempty"`
	CanceledAt     *core.TimeModel         `json:"canceled_at,omitempty" validate:"omitempty"`
}

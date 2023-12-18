package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
)

type FilterableBatchJob struct {
	core.FilterModel

	Status []models.BatchJobStatus `json:"status,omitempty"`
	Type   []string                `json:"type,omitempty"`
}

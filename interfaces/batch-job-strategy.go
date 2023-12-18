package interfaces

import (
	"net/http"

	"github.com/driver005/gateway/core"
	"github.com/gofrs/uuid"
)

type CreateBatchJobInput struct {
	Type    string
	Context core.JSONB
	DryRun  bool
}

type IBatchJobStrategy interface {
	PrepareBatchJobForProcessing(batchJobEntity CreateBatchJobInput, req *http.Request) (*CreateBatchJobInput, error)
	PreProcessBatchJob(batchJobId uuid.UUID) error
	ProcessJob(batchJobId uuid.UUID) error
	BuildTemplate() (string, error)
}

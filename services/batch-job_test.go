package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/sarulabs/di"
)

func TestNewBatchJobService(t *testing.T) {
	type args struct {
		container di.Container
		r         Registry
	}
	tests := []struct {
		name string
		args args
		want *BatchJobService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBatchJobService(tt.args.container, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBatchJobService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBatchJobService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *BatchJobService
		args args
		want *BatchJobService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BatchJobService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBatchJobService_ResolveBatchJobByType(t *testing.T) {
	type args struct {
		batchtype string
	}
	tests := []struct {
		name string
		s    *BatchJobService
		args args
		want interfaces.IBatchJobStrategy
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ResolveBatchJobByType(tt.args.batchtype); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BatchJobService.ResolveBatchJobByType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBatchJobService_Retrive(t *testing.T) {
	type args struct {
		batchJobId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *BatchJobService
		args  args
		want  *models.BatchJob
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrive(tt.args.batchJobId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BatchJobService.Retrive() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BatchJobService.Retrive() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBatchJobService_ListAndCount(t *testing.T) {
	type args struct {
		selector *types.FilterableBatchJob
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *BatchJobService
		args  args
		want  []models.BatchJob
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BatchJobService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BatchJobService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("BatchJobService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestBatchJobService_Create(t *testing.T) {
	type args struct {
		data *types.BatchJobCreateProps
	}
	tests := []struct {
		name  string
		s     *BatchJobService
		args  args
		want  *models.BatchJob
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BatchJobService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BatchJobService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBatchJobService_Update(t *testing.T) {
	type args struct {
		batchJobId uuid.UUID
		model      *models.BatchJob
		data       *types.BatchJobUpdateProps
	}
	tests := []struct {
		name  string
		s     *BatchJobService
		args  args
		want  *models.BatchJob
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.batchJobId, tt.args.model, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BatchJobService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BatchJobService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBatchJobService_Confirm(t *testing.T) {
	type args struct {
		batchJobId uuid.UUID
		model      *models.BatchJob
	}
	tests := []struct {
		name  string
		s     *BatchJobService
		args  args
		want  *models.BatchJob
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Confirm(tt.args.batchJobId, tt.args.model)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BatchJobService.Confirm() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BatchJobService.Confirm() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBatchJobService_Complete(t *testing.T) {
	type args struct {
		batchJobId uuid.UUID
		model      *models.BatchJob
	}
	tests := []struct {
		name  string
		s     *BatchJobService
		args  args
		want  *models.BatchJob
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Complete(tt.args.batchJobId, tt.args.model)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BatchJobService.Complete() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BatchJobService.Complete() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBatchJobService_Cancel(t *testing.T) {
	type args struct {
		batchJobId uuid.UUID
		model      *models.BatchJob
	}
	tests := []struct {
		name  string
		s     *BatchJobService
		args  args
		want  *models.BatchJob
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Cancel(tt.args.batchJobId, tt.args.model)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BatchJobService.Cancel() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BatchJobService.Cancel() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBatchJobService_SetPreProcessingDone(t *testing.T) {
	type args struct {
		batchJobId uuid.UUID
		model      *models.BatchJob
	}
	tests := []struct {
		name  string
		s     *BatchJobService
		args  args
		want  *models.BatchJob
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.SetPreProcessingDone(tt.args.batchJobId, tt.args.model)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BatchJobService.SetPreProcessingDone() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BatchJobService.SetPreProcessingDone() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBatchJobService_SetProcessing(t *testing.T) {
	type args struct {
		batchJobId uuid.UUID
		model      *models.BatchJob
	}
	tests := []struct {
		name  string
		s     *BatchJobService
		args  args
		want  *models.BatchJob
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.SetProcessing(tt.args.batchJobId, tt.args.model)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BatchJobService.SetProcessing() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BatchJobService.SetProcessing() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBatchJobService_SetFailed(t *testing.T) {
	type args struct {
		batchJobId   uuid.UUID
		model        *models.BatchJob
		errorMessage types.BatchJobResultError
	}
	tests := []struct {
		name  string
		s     *BatchJobService
		args  args
		want  *models.BatchJob
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.SetFailed(tt.args.batchJobId, tt.args.model, tt.args.errorMessage)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BatchJobService.SetFailed() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BatchJobService.SetFailed() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBatchJobService_PrepareBatchJobForProcessing(t *testing.T) {
	type args struct {
		data *types.CreateBatchJobInput
	}
	tests := []struct {
		name  string
		s     *BatchJobService
		args  args
		want  *types.CreateBatchJobInput
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.PrepareBatchJobForProcessing(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BatchJobService.PrepareBatchJobForProcessing() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BatchJobService.PrepareBatchJobForProcessing() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBatchJobService_UpdateStatus(t *testing.T) {
	type args struct {
		batchJobId uuid.UUID
		model      *models.BatchJob
		status     models.BatchJobStatus
	}
	tests := []struct {
		name  string
		s     *BatchJobService
		args  args
		want  *models.BatchJob
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.UpdateStatus(tt.args.batchJobId, tt.args.model, tt.args.status)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BatchJobService.UpdateStatus() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BatchJobService.UpdateStatus() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

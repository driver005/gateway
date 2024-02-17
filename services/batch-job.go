package services

import (
	"context"
	"reflect"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/sarulabs/di"
)

type BatchJobService struct {
	ctx       context.Context
	container di.Container
	r         Registry
}

func NewBatchJobService(container di.Container, r Registry) *BatchJobService {
	return &BatchJobService{
		context.Background(),
		container,
		r,
	}
}

func (s *BatchJobService) SetContext(context context.Context) *BatchJobService {
	s.ctx = context

	return s
}

func (s *BatchJobService) ResolveBatchJobByType(batchtype string) interfaces.IBatchJobStrategy {
	var provider interfaces.IBatchJobStrategy
	objectInterface, err := s.container.SafeGet("batchType_" + batchtype)
	if err != nil {
		panic("Unable to find a BatchJob strategy with the type " + batchtype)
	}
	provider, _ = objectInterface.(interfaces.IBatchJobStrategy)

	return provider
}

func (s *BatchJobService) Retrive(batchJobId uuid.UUID) (*models.BatchJob, *utils.ApplictaionError) {
	var model *models.BatchJob

	if err := s.r.BatchJobRepository().FindOne(s.ctx, model, sql.BuildQuery[models.BatchJob](
		models.BatchJob{
			Model: core.Model{Id: batchJobId},
		},
		&sql.Options{},
	)); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *BatchJobService) ListAndCount(selector *types.FilterableBatchJob, config *sql.Options) ([]models.BatchJob, *int64, *utils.ApplictaionError) {
	var res []models.BatchJob

	if reflect.DeepEqual(config, &sql.Options{}) {
		config.Skip = 0
		config.Take = 20
	}

	query := sql.BuildQuery(selector, config)
	count, err := s.r.BatchJobRepository().FindAndCount(s.ctx, &res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

func (s *BatchJobService) Create(data *types.BatchJobCreateProps) (*models.BatchJob, *utils.ApplictaionError) {
	model := &models.BatchJob{
		Type:      data.Type,
		Status:    models.BatchJobStatusCreated,
		Context:   data.Context,
		CreatedBy: data.CreatedBy,
		DryRun:    data.DryRun,
	}
	if err := s.r.BatchJobRepository().Save(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *BatchJobService) Update(batchJobId uuid.UUID, model *models.BatchJob, data *types.BatchJobUpdateProps) (*models.BatchJob, *utils.ApplictaionError) {
	if batchJobId != uuid.Nil {
		batch, err := s.Retrive(batchJobId)
		if err != nil {
			return nil, err
		}

		model = batch
	}

	if data.Context != nil {
		model.Context = utils.MergeMaps(model.Context, data.Context)
	}

	if data.Result != nil {
		model.Result = data.Result
	}

	if err := s.r.BatchJobRepository().Update(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *BatchJobService) Confirm(batchJobId uuid.UUID, model *models.BatchJob) (*models.BatchJob, *utils.ApplictaionError) {
	if batchJobId != uuid.Nil {
		model.Id = batchJobId

		if err := s.r.BatchJobRepository().FindOne(s.ctx, model, sql.Query{}); err != nil {
			return nil, err
		}
	}

	if model.Status != models.BatchJobStatusPreProcessed {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			`cannot confirm processing for a batch job that is not pre processed`,
			nil,
		)
	}

	return s.UpdateStatus(uuid.Nil, model, models.BatchJobStatusConfirmed)
}

func (s *BatchJobService) Complete(batchJobId uuid.UUID, model *models.BatchJob) (*models.BatchJob, *utils.ApplictaionError) {
	if batchJobId != uuid.Nil {
		model.Id = batchJobId

		if err := s.r.BatchJobRepository().FindOne(s.ctx, model, sql.Query{}); err != nil {
			return nil, err
		}
	}

	if model.Status != models.BatchJobStatusProcessing {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			`cannot complete a batch job with status "${batchJob.status}". the batch job must be processing`,
			nil,
		)
	}

	return s.UpdateStatus(uuid.Nil, model, models.BatchJobStatusCompleted)
}

func (s *BatchJobService) Cancel(batchJobId uuid.UUID, model *models.BatchJob) (*models.BatchJob, *utils.ApplictaionError) {
	if batchJobId != uuid.Nil {
		model.Id = batchJobId

		if err := s.r.BatchJobRepository().FindOne(s.ctx, model, sql.Query{}); err != nil {
			return nil, err
		}
	}

	if model.Status == models.BatchJobStatusCompleted {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			`cannot cancel completed batch job`,
			nil,
		)
	}

	return s.UpdateStatus(uuid.Nil, model, models.BatchJobStatusCanceled)
}

func (s *BatchJobService) SetPreProcessingDone(batchJobId uuid.UUID, model *models.BatchJob) (*models.BatchJob, *utils.ApplictaionError) {
	if batchJobId != uuid.Nil {
		model.Id = batchJobId

		if err := s.r.BatchJobRepository().FindOne(s.ctx, model, sql.Query{}); err != nil {
			return nil, err
		}
	}

	if model.Status == models.BatchJobStatusPreProcessed {
		return model, nil
	}

	if model.Status != models.BatchJobStatusCreated {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			`cannot mark a batch job as pre processed if it is not in created status`,
			nil,
		)
	}

	res, err := s.UpdateStatus(uuid.Nil, model, models.BatchJobStatusPreProcessed)
	if err != nil {
		return nil, err
	}

	return s.Confirm(uuid.Nil, res)
}

func (s *BatchJobService) SetProcessing(batchJobId uuid.UUID, model *models.BatchJob) (*models.BatchJob, *utils.ApplictaionError) {
	if batchJobId != uuid.Nil {
		model.Id = batchJobId

		if err := s.r.BatchJobRepository().FindOne(s.ctx, model, sql.Query{}); err != nil {
			return nil, err
		}
	}

	if model.Status != models.BatchJobStatusConfirmed {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			`cannot mark a batch job as processing if the status is different that confirmed`,
			nil,
		)
	}

	return s.UpdateStatus(uuid.Nil, model, models.BatchJobStatusProcessing)
}

func (s *BatchJobService) SetFailed(batchJobId uuid.UUID, model *models.BatchJob, errorMessage types.BatchJobResultError) (*models.BatchJob, *utils.ApplictaionError) {
	if batchJobId != uuid.Nil {
		model.Id = batchJobId

		if err := s.r.BatchJobRepository().FindOne(s.ctx, model, sql.Query{}); err != nil {
			return nil, err
		}
	}

	res, err := s.Update(uuid.Nil, model, &types.BatchJobUpdateProps{
		Result: &models.BatchJobResult{
			Errors: &models.BatchJobResultErrors{
				Message: errorMessage.Message,
				Code:    errorMessage.Code,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return s.UpdateStatus(uuid.Nil, res, models.BatchJobStatusFailed)
}

func (s *BatchJobService) PrepareBatchJobForProcessing(data *types.CreateBatchJobInput) (*types.CreateBatchJobInput, *utils.ApplictaionError) {
	// s.r.BatchJobStrategy().PrepareBatchJobForProcessing()
	return nil, nil
}

func (s *BatchJobService) UpdateStatus(batchJobId uuid.UUID, model *models.BatchJob, status models.BatchJobStatus) (*models.BatchJob, *utils.ApplictaionError) {
	if batchJobId != uuid.Nil {
		model.Id = batchJobId

		if err := s.r.BatchJobRepository().FindOne(s.ctx, model, sql.Query{}); err != nil {
			return nil, err
		}
	}

	now := time.Now()

	if status == models.BatchJobStatusCreated {
		model.CreatedAt = now
		model.Status = models.BatchJobStatusCreated
	} else if status == models.BatchJobStatusPreProcessed {
		model.PreProcessedAt = &now
		model.Status = models.BatchJobStatusPreProcessed
	} else if status == models.BatchJobStatusConfirmed {
		model.ConfirmedAt = &now
		model.Status = models.BatchJobStatusConfirmed
	} else if status == models.BatchJobStatusProcessing {
		model.PreProcessedAt = &now
		model.Status = models.BatchJobStatusProcessing
	} else if status == models.BatchJobStatusCompleted {
		model.CompletedAt = &now
		model.Status = models.BatchJobStatusCompleted
	} else if status == models.BatchJobStatusCanceled {
		model.CanceledAt = &now
		model.Status = models.BatchJobStatusCanceled
	} else if status == models.BatchJobStatusFailed {
		model.FailedAt = &now
		model.Status = models.BatchJobStatusFailed
	}

	if err := s.r.BatchJobRepository().Upsert(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

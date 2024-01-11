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
	"github.com/icza/gox/gox"
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
		sql.Options{},
	)); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *BatchJobService) ListAndCount(selector types.FilterableBatchJob, config sql.Options) ([]models.BatchJob, *int64, *utils.ApplictaionError) {
	var res []models.BatchJob

	if reflect.DeepEqual(config, sql.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(20)
	}

	query := sql.BuildQuery[types.FilterableBatchJob](selector, config)
	count, err := s.r.BatchJobRepository().FindAndCount(s.ctx, res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

func (s *BatchJobService) Create(model *models.BatchJob) (*models.BatchJob, *utils.ApplictaionError) {
	if err := s.r.BatchJobRepository().Save(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *BatchJobService) Update(batchJobId uuid.UUID, model *models.BatchJob) (*models.BatchJob, *utils.ApplictaionError) {
	if batchJobId != uuid.Nil {
		model.Id = batchJobId

		if err := s.r.BatchJobRepository().FindOne(s.ctx, model, sql.Query{}); err != nil {
			return nil, err
		}
	}

	if err := s.r.BatchJobRepository().Upsert(s.ctx, model); err != nil {
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
			"500",
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
			"500",
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
			"500",
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
			"500",
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
			"500",
			nil,
		)
	}

	return s.UpdateStatus(uuid.Nil, model, models.BatchJobStatusProcessing)
}

func (s *BatchJobService) SetFailed(batchJobId uuid.UUID, model *models.BatchJob, errorMessage models.BatchJobResultErrors) (*models.BatchJob, *utils.ApplictaionError) {
	if batchJobId != uuid.Nil {
		model.Id = batchJobId

		if err := s.r.BatchJobRepository().FindOne(s.ctx, model, sql.Query{}); err != nil {
			return nil, err
		}
	}

	model.Result = models.BatchJobResult{
		Errors: errorMessage,
	}

	res, err := s.Update(uuid.Nil, model)
	if err != nil {
		return nil, err
	}

	return s.UpdateStatus(uuid.Nil, res, models.BatchJobStatusFailed)
}

func (s *BatchJobService) PrepareBatchJobForProcessing(data interface{}) (*models.BatchJob, *utils.ApplictaionError) {

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

package services

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/repository"
	"github.com/driver005/gateway/types"
	"github.com/google/uuid"
	"github.com/icza/gox/gox"
)

type BatchJobService struct {
	ctx  context.Context
	repo *repository.BatchJobRepo
}

func NewBatchJobService(ctx context.Context, repo *repository.BatchJobRepo) *BatchJobService {
	return &BatchJobService{
		ctx,
		repo,
	}
}

func (s *BatchJobService) Retrive(batchJobId uuid.UUID) (*models.BatchJob, error) {
	var model *models.BatchJob

	if err := s.repo.FindOne(s.ctx, model, repository.BuildQuery[models.BatchJob](
		models.BatchJob{
			Model: core.Model{Id: batchJobId},
		},
		repository.Options{},
	)); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *BatchJobService) ListAndCount(selector types.FilterableBatchJob, config repository.Options) ([]models.BatchJob, *int64, error) {
	var res []models.BatchJob

	if reflect.DeepEqual(config, repository.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(20)
	}

	query := repository.BuildQuery[types.FilterableBatchJob](selector, config)
	count, err := s.repo.FindAndCount(s.ctx, res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

func (s *BatchJobService) Create(model *models.BatchJob) (*models.BatchJob, error) {
	if err := s.repo.Save(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *BatchJobService) Update(batchJobId uuid.UUID, model *models.BatchJob) (*models.BatchJob, error) {
	if batchJobId != uuid.Nil {
		model.Id = batchJobId

		if err := s.repo.FindOne(s.ctx, model, repository.Query{}); err != nil {
			return nil, err
		}
	}

	if err := s.repo.Upsert(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *BatchJobService) Confirm(batchJobId uuid.UUID, model *models.BatchJob) (*models.BatchJob, error) {
	if batchJobId != uuid.Nil {
		model.Id = batchJobId

		if err := s.repo.FindOne(s.ctx, model, repository.Query{}); err != nil {
			return nil, err
		}
	}

	if model.Status != models.BatchJobStatusPreProcessed {
		return nil, errors.New(`cannot confirm processing for a batch job that is not pre processed`)
	}

	return s.UpdateStatus(uuid.Nil, model, models.BatchJobStatusConfirmed)
}

func (s *BatchJobService) Complete(batchJobId uuid.UUID, model *models.BatchJob) (*models.BatchJob, error) {
	if batchJobId != uuid.Nil {
		model.Id = batchJobId

		if err := s.repo.FindOne(s.ctx, model, repository.Query{}); err != nil {
			return nil, err
		}
	}

	if model.Status != models.BatchJobStatusProcessing {
		return nil, errors.New(`cannot complete a batch job with status "${batchJob.status}". the batch job must be processing`)
	}

	return s.UpdateStatus(uuid.Nil, model, models.BatchJobStatusCompleted)
}

func (s *BatchJobService) Cancel(batchJobId uuid.UUID, model *models.BatchJob) (*models.BatchJob, error) {
	if batchJobId != uuid.Nil {
		model.Id = batchJobId

		if err := s.repo.FindOne(s.ctx, model, repository.Query{}); err != nil {
			return nil, err
		}
	}

	if model.Status == models.BatchJobStatusCompleted {
		return nil, errors.New(`cannot cancel completed batch job`)
	}

	return s.UpdateStatus(uuid.Nil, model, models.BatchJobStatusCanceled)
}

func (s *BatchJobService) SetPreProcessingDone(batchJobId uuid.UUID, model *models.BatchJob) (*models.BatchJob, error) {
	if batchJobId != uuid.Nil {
		model.Id = batchJobId

		if err := s.repo.FindOne(s.ctx, model, repository.Query{}); err != nil {
			return nil, err
		}
	}

	if model.Status == models.BatchJobStatusPreProcessed {
		return model, nil
	}

	if model.Status != models.BatchJobStatusCreated {
		return nil, errors.New(`cannot mark a batch job as pre processed if it is not in created status`)
	}

	res, err := s.UpdateStatus(uuid.Nil, model, models.BatchJobStatusPreProcessed)
	if err != nil {
		return nil, err
	}

	return s.Confirm(uuid.Nil, res)
}

func (s *BatchJobService) SetProcessing(batchJobId uuid.UUID, model *models.BatchJob) (*models.BatchJob, error) {
	if batchJobId != uuid.Nil {
		model.Id = batchJobId

		if err := s.repo.FindOne(s.ctx, model, repository.Query{}); err != nil {
			return nil, err
		}
	}

	if model.Status != models.BatchJobStatusConfirmed {
		return nil, errors.New(`cannot mark a batch job as processing if the status is different that confirmed`)
	}

	return s.UpdateStatus(uuid.Nil, model, models.BatchJobStatusProcessing)
}

func (s *BatchJobService) SetFailed(batchJobId uuid.UUID, model *models.BatchJob, errorMessage models.BatchJobResultErrors) (*models.BatchJob, error) {
	if batchJobId != uuid.Nil {
		model.Id = batchJobId

		if err := s.repo.FindOne(s.ctx, model, repository.Query{}); err != nil {
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

func (s *BatchJobService) PrepareBatchJobForProcessing(data interface{}) (*models.BatchJob, error) {

	return nil, nil
}

func (s *BatchJobService) UpdateStatus(batchJobId uuid.UUID, model *models.BatchJob, status models.BatchJobStatus) (*models.BatchJob, error) {
	if batchJobId != uuid.Nil {
		model.Id = batchJobId

		if err := s.repo.FindOne(s.ctx, model, repository.Query{}); err != nil {
			return nil, err
		}
	}

	if status == models.BatchJobStatusCreated {
		model.CreatedAt = time.Now()
		model.Status = models.BatchJobStatusCreated
	} else if status == models.BatchJobStatusPreProcessed {
		model.PreProcessedAt = time.Now()
		model.Status = models.BatchJobStatusPreProcessed
	} else if status == models.BatchJobStatusConfirmed {
		model.ConfirmedAt = time.Now()
		model.Status = models.BatchJobStatusConfirmed
	} else if status == models.BatchJobStatusProcessing {
		model.PreProcessedAt = time.Now()
		model.Status = models.BatchJobStatusProcessing
	} else if status == models.BatchJobStatusCompleted {
		model.CompletedAt = time.Now()
		model.Status = models.BatchJobStatusCompleted
	} else if status == models.BatchJobStatusCanceled {
		model.CanceledAt = time.Now()
		model.Status = models.BatchJobStatusCanceled
	} else if status == models.BatchJobStatusFailed {
		model.FailedAt = time.Now()
		model.Status = models.BatchJobStatusFailed
	}

	if err := s.repo.Upsert(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

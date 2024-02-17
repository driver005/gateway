package services

import (
	"context"
	"reflect"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type StagedJobService struct {
	ctx context.Context
	r   Registry
}

func NewStagedJobService(r Registry) *StagedJobService {
	return &StagedJobService{
		context.Background(),
		r,
	}
}

func (s *StagedJobService) SetContext(context context.Context) *StagedJobService {
	s.ctx = context
	return s
}

func (s *StagedJobService) List(selector models.StagedJob, config *sql.Options) ([]models.StagedJob, *utils.ApplictaionError) {
	var res []models.StagedJob

	if reflect.DeepEqual(config, &sql.Options{}) {
		config.Skip = 0
		config.Take = 50
		config.Order = "created_at DESC"
	}

	query := sql.BuildQuery(selector, config)

	if err := s.r.StagedJobRepository().Find(s.ctx, &res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *StagedJobService) Delete(stagedJobIds uuid.UUIDs) *utils.ApplictaionError {
	var data []models.StagedJob
	query := sql.BuildQuery(models.StagedJob{}, &sql.Options{
		Specification: []sql.Specification{sql.In("id", stagedJobIds)},
	})
	if err := s.r.StagedJobRepository().Find(s.ctx, &data, query); err != nil {
		return err
	}

	if err := s.r.StagedJobRepository().RemoveSlice(s.ctx, data); err != nil {
		return err
	}

	return nil
}

func (s *StagedJobService) Create(data []types.EmitData) *utils.ApplictaionError {
	var stagedJobs []models.StagedJob

	for i, job := range data {
		stagedJobs[i] = models.StagedJob{
			EventName: job.EventName,
			Data:      job.Data,
			Options:   job.Options,
		}
	}

	if err := s.r.StagedJobRepository().InsertSlice(s.ctx, &stagedJobs); err != nil {
		return err
	}

	return nil
}

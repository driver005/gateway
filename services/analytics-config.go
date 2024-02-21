package services

import (
	"context"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type AnalyticsConfigService struct {
	ctx context.Context
	r   Registry
}

func NewAnalyticsConfigService(
	r Registry,
) *AnalyticsConfigService {
	return &AnalyticsConfigService{
		context.Background(),
		r,
	}
}

func (s *AnalyticsConfigService) SetContext(context context.Context) *AnalyticsConfigService {
	s.ctx = context
	return s
}

func (s *AnalyticsConfigService) Retrive(userId uuid.UUID) (*models.AnalyticsConfig, *utils.ApplictaionError) {
	var model *models.AnalyticsConfig = &models.AnalyticsConfig{}
	if err := s.r.AnalyticsConfigRepository().FindOne(s.ctx, model, sql.BuildQuery[models.AnalyticsConfig](
		models.AnalyticsConfig{
			UserId: uuid.NullUUID{UUID: userId},
		},
		&sql.Options{},
	)); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *AnalyticsConfigService) Create(userId uuid.UUID, data *types.CreateAnalyticsConfig) (*models.AnalyticsConfig, *utils.ApplictaionError) {
	model := &models.AnalyticsConfig{
		UserId:    uuid.NullUUID{UUID: userId},
		OptOut:    data.OptOut,
		Anonymize: data.Anonymize,
	}
	if err := s.r.AnalyticsConfigRepository().Save(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *AnalyticsConfigService) Update(userId uuid.UUID, data *types.UpdateAnalyticsConfig) (*models.AnalyticsConfig, *utils.ApplictaionError) {
	model := &models.AnalyticsConfig{
		UserId:    uuid.NullUUID{UUID: userId},
		OptOut:    data.OptOut,
		Anonymize: data.Anonymize,
	}
	if err := s.r.AnalyticsConfigRepository().Upsert(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *AnalyticsConfigService) Delete(userId uuid.UUID) *utils.ApplictaionError {
	data, err := s.Retrive(userId)
	if err != nil {
		return err
	}

	if err := s.r.AnalyticsConfigRepository().SoftRemove(s.ctx, data); err != nil {
		return err
	}

	return nil
}

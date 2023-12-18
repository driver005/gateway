package services

import (
	"context"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/repository"
	"github.com/google/uuid"
)

type AnalyticsConfigService struct {
	ctx  context.Context
	repo *repository.AnalyticsConfigRepo
}

func NewAnalyticsConfigService(
	ctx context.Context,
	repo *repository.AnalyticsConfigRepo,
) *AnalyticsConfigService {
	return &AnalyticsConfigService{
		ctx,
		repo,
	}
}

func (s *AnalyticsConfigService) Retrive(userId uuid.UUID) (*models.AnalyticsConfig, error) {
	var model *models.AnalyticsConfig
	if err := s.repo.FindOne(s.ctx, model, repository.BuildQuery[models.AnalyticsConfig](
		models.AnalyticsConfig{
			UserId: userId,
		},
		repository.Options{},
	)); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *AnalyticsConfigService) Create(model *models.AnalyticsConfig) (*models.AnalyticsConfig, error) {
	if err := s.repo.Save(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *AnalyticsConfigService) Update(model *models.AnalyticsConfig) (*models.AnalyticsConfig, error) {
	if err := s.repo.Upsert(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *AnalyticsConfigService) Delete(userId uuid.UUID) error {
	data, err := s.Retrive(userId)
	if err != nil {
		return err
	}

	if err := s.repo.SoftRemove(s.ctx, data); err != nil {
		return err
	}

	return nil
}

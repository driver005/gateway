package services

import (
	"context"
	"errors"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/repository"
	"github.com/google/uuid"
)

type CustomShippingOptionService struct {
	ctx  context.Context
	repo *repository.CustomShippingOptionRepo
}

func NewCustomShippingOptionService(
	ctx context.Context,
	repo *repository.CustomShippingOptionRepo,
) *CustomShippingOptionService {
	return &CustomShippingOptionService{
		ctx,
		repo,
	}
}

func (s *CustomShippingOptionService) Retrieve(id uuid.UUID, config repository.Options) (*models.CustomShippingOption, error) {
	var res *models.CustomShippingOption
	query := repository.BuildQuery(models.CustomShippingOption{Model: core.Model{Id: id}}, config)
	if err := s.repo.FindOne(s.ctx, res, query); err != nil {
		return nil, errors.New("Custom shipping option with id: " + id.String() + " was not found.")
	}
	return res, nil
}

func (s *CustomShippingOptionService) List(selector models.CustomShippingOption, config repository.Options) ([]models.CustomShippingOption, error) {
	var res []models.CustomShippingOption
	query := repository.BuildQuery(selector, config)
	if err := s.repo.Find(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CustomShippingOptionService) Create(data []models.CustomShippingOption) ([]models.CustomShippingOption, error) {
	if err := s.repo.SaveSlice(s.ctx, data); err != nil {
		return nil, err
	}

	return data, nil
}

package services

import (
	"context"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type CustomShippingOptionService struct {
	ctx context.Context
	r   Registry
}

func NewCustomShippingOptionService(
	r Registry,
) *CustomShippingOptionService {
	return &CustomShippingOptionService{
		context.Background(),
		r,
	}
}

func (s *CustomShippingOptionService) SetContext(context context.Context) *CustomShippingOptionService {
	s.ctx = context
	return s
}

func (s *CustomShippingOptionService) Retrieve(id uuid.UUID, config sql.Options) (*models.CustomShippingOption, *utils.ApplictaionError) {
	if id == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"id" must be defined`,
			"500",
			nil,
		)
	}
	var res *models.CustomShippingOption
	query := sql.BuildQuery(models.CustomShippingOption{Model: core.Model{Id: id}}, config)
	if err := s.r.CustomShippingOptionRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CustomShippingOptionService) List(selector models.CustomShippingOption, config sql.Options) ([]models.CustomShippingOption, *utils.ApplictaionError) {
	var res []models.CustomShippingOption
	query := sql.BuildQuery(selector, config)
	if err := s.r.CustomShippingOptionRepository().Find(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CustomShippingOptionService) Create(data []models.CustomShippingOption) ([]models.CustomShippingOption, *utils.ApplictaionError) {
	if err := s.r.CustomShippingOptionRepository().SaveSlice(s.ctx, data); err != nil {
		return nil, err
	}

	return data, nil
}

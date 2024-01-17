package services

import (
	"context"
	"reflect"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/icza/gox/gox"
)

type ProductTypeService struct {
	ctx context.Context
	r   Registry
}

func NewProductTypeService(
	r Registry,
) *ProductTypeService {
	return &ProductTypeService{
		context.Background(),
		r,
	}
}

func (s *ProductTypeService) SetContext(context context.Context) *ProductTypeService {
	s.ctx = context
	return s
}

func (s *ProductTypeService) Retrieve(id uuid.UUID, config *sql.Options) (*models.ProductType, *utils.ApplictaionError) {
	if id == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"id" must be defined`,
			nil,
		)
	}

	var res *models.ProductType

	query := sql.BuildQuery(models.ProductType{Model: core.Model{Id: id}}, config)

	if err := s.r.ProductTypeRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ProductTypeService) List(selector models.ProductType, config *sql.Options, q *string) ([]models.ProductType, *utils.ApplictaionError) {
	productTypes, _, err := s.ListAndCount(selector, config, q)
	if err != nil {
		return nil, err
	}
	return productTypes, nil
}

func (s *ProductTypeService) ListAndCount(selector models.ProductType, config *sql.Options, q *string) ([]models.ProductType, *int64, *utils.ApplictaionError) {
	var res []models.ProductType

	if reflect.DeepEqual(config, &sql.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(20)
	}

	if q != nil {
		v := sql.ILike(*q)
		selector.Value = v
	}

	query := sql.BuildQuery(selector, config)

	count, err := s.r.ProductTypeRepository().FindAndCount(s.ctx, res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

package services

import (
	"context"
	"reflect"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
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

	var res *models.ProductType = &models.ProductType{}

	query := sql.BuildQuery(models.ProductType{SoftDeletableModel: core.SoftDeletableModel{Id: id}}, config)

	if err := s.r.ProductTypeRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ProductTypeService) List(selector *types.FilterableProductType, config *sql.Options) ([]models.ProductType, *utils.ApplictaionError) {
	productTypes, _, err := s.ListAndCount(selector, config)
	if err != nil {
		return nil, err
	}
	return productTypes, nil
}

func (s *ProductTypeService) ListAndCount(selector *types.FilterableProductType, config *sql.Options) ([]models.ProductType, *int64, *utils.ApplictaionError) {
	var res []models.ProductType

	if reflect.DeepEqual(config, &sql.Options{}) {
		config.Skip = 0
		config.Take = 20
	}

	if !reflect.ValueOf(config.Q).IsZero() {
		v := sql.ILike(config.Q)
		selector.Value = v
	}

	query := sql.BuildQuery(selector, config)

	count, err := s.r.ProductTypeRepository().FindAndCount(s.ctx, &res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

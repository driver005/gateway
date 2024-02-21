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

type ProductTagService struct {
	ctx context.Context
	r   Registry
}

func NewProductTagService(
	r Registry,
) *ProductTagService {
	return &ProductTagService{
		context.Background(),
		r,
	}
}

func (s *ProductTagService) SetContext(context context.Context) *ProductTagService {
	s.ctx = context
	return s
}

func (s *ProductTagService) Retrieve(id uuid.UUID, config *sql.Options) (*models.ProductTag, *utils.ApplictaionError) {
	if id == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"id" must be defined`,
			nil,
		)
	}

	var res *models.ProductTag = &models.ProductTag{}

	query := sql.BuildQuery(models.ProductTag{SoftDeletableModel: core.SoftDeletableModel{Id: id}}, config)

	if err := s.r.ProductTagRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ProductTagService) Create(tag *models.ProductTag) (*models.ProductTag, *utils.ApplictaionError) {
	if err := s.r.ProductTagRepository().Save(s.ctx, tag); err != nil {
		return nil, err
	}
	return tag, nil
}

func (s *ProductTagService) List(selector *types.FilterableProductTag, config *sql.Options) ([]models.ProductTag, *utils.ApplictaionError) {
	tags, _, err := s.ListAndCount(selector, config)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (s *ProductTagService) ListAndCount(selector *types.FilterableProductTag, config *sql.Options) ([]models.ProductTag, *int64, *utils.ApplictaionError) {
	var res []models.ProductTag

	if reflect.DeepEqual(config, &sql.Options{}) {
		config.Skip = 0
		config.Take = 20
	}

	if !reflect.ValueOf(config.Q).IsZero() {
		v := sql.ILike(config.Q)
		selector.Value = v
	}

	query := sql.BuildQuery(selector, config)

	if selector.DiscountConditionId != uuid.Nil {
		return s.r.ProductTagRepository().FindAndCountByDiscountConditionID(selector.DiscountConditionId, query)
	}

	count, err := s.r.ProductTagRepository().FindAndCount(s.ctx, &res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

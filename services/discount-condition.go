package services

import (
	"context"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type DiscountConditionService struct {
	ctx context.Context
	r   Registry
}

func NewDiscountConditionService(
	r Registry,
) *DiscountConditionService {
	return &DiscountConditionService{
		context.Background(),
		r,
	}
}

func (s *DiscountConditionService) SetContext(context context.Context) *DiscountConditionService {
	s.ctx = context
	return s
}

func (s *DiscountConditionService) Retrieve(conditionId uuid.UUID, config *sql.Options) (*models.DiscountCondition, *utils.ApplictaionError) {
	if conditionId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"conditionId" must be defined`,
			nil,
		)
	}

	var res *models.DiscountCondition = &models.DiscountCondition{}
	query := sql.BuildQuery[models.DiscountCondition](models.DiscountCondition{Model: core.Model{Id: conditionId}}, &sql.Options{})

	if err := s.r.DiscountConditionRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *DiscountConditionService) ResolveConditionType(data *types.DiscountConditionInput) (*models.DiscountCondition, *utils.ApplictaionError) {
	var model *models.DiscountCondition = &models.DiscountCondition{}
	switch {
	case data.Products != nil:
		model.Type = models.DiscountConditionTypeProducts
		return model, nil
	case data.ProductCollections != nil:
		model.Type = models.DiscountConditionTypeProductCollections
		return model, nil
	case data.ProductTypes != nil:
		model.Type = models.DiscountConditionTypeProductTypes
		return model, nil
	case data.ProductTags != nil:
		model.Type = models.DiscountConditionTypeProductTags
		return model, nil
	case data.CustomerGroups != nil:
		model.Type = models.DiscountConditionTypeCustomerGroups
		return model, nil
	default:
		return nil, nil
	}
}

func (s *DiscountConditionService) UpsertCondition(data *types.DiscountConditionInput, overrideExisting bool) ([]models.DiscountCondition, *utils.ApplictaionError) {
	resolvedConditionType, err := s.ResolveConditionType(data)
	if err != nil {
		return nil, err
	}

	if resolvedConditionType == nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`Missing one of products, collections, tags, types or customer groups in data`,
			nil,
		)
	}

	if data.Id != uuid.Nil {
		resolvedCondition, err := s.Retrieve(data.Id, &sql.Options{})
		if err != nil {
			return nil, err
		}

		if data.Operator != resolvedCondition.Operator {
			resolvedCondition.Operator = data.Operator
			err = s.r.DiscountConditionRepository().Save(s.ctx, resolvedCondition)
			if err != nil {
				return nil, err
			}
		}

		return s.r.DiscountConditionRepository().AddConditionResources(s.ctx, data.Id, resolvedConditionType, overrideExisting)
	}

	created := &models.DiscountCondition{
		Model: core.Model{
			Id: data.Id,
		},
		Operator: data.Operator,
		Type:     resolvedConditionType.Type,
	}

	if err := s.r.DiscountConditionRepository().Save(s.ctx, created); err != nil {
		return nil, err
	}

	return s.r.DiscountConditionRepository().AddConditionResources(s.ctx, created.Id, resolvedConditionType, overrideExisting)
}

func (s *DiscountConditionService) RemoveResources(data *types.DiscountConditionInput) *utils.ApplictaionError {
	resolvedConditionType, err := s.ResolveConditionType(data)
	if err != nil {
		return err
	}

	if resolvedConditionType == nil {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			`Missing one of products, collections, tags, types or customer groups in data`,
			nil,
		)
	}

	resolvedCondition, err := s.Retrieve(data.Id, &sql.Options{})
	if err != nil {
		return err
	}

	if data.Operator != resolvedCondition.Operator {
		resolvedCondition.Operator = data.Operator
		if err := s.r.DiscountConditionRepository().Save(s.ctx, resolvedCondition); err != nil {
			return err
		}
	}

	return s.r.DiscountConditionRepository().RemoveConditionResources(s.ctx, data.Id, resolvedConditionType)
}

func (s *DiscountConditionService) Delete(conditionId uuid.UUID) *utils.ApplictaionError {
	condition, err := s.Retrieve(conditionId, &sql.Options{})
	if err != nil {
		return err
	}

	if err := s.r.DiscountConditionRepository().Remove(s.ctx, condition); err != nil {
		return err
	}

	return nil
}

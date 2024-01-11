package services

import (
	"context"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
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

func (s *DiscountConditionService) Retrieve(conditionId uuid.UUID, config sql.Options) (*models.DiscountCondition, *utils.ApplictaionError) {
	if conditionId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"conditionId" must be defined`,
			"500",
			nil,
		)
	}

	var res *models.DiscountCondition
	query := sql.BuildQuery[models.DiscountCondition](models.DiscountCondition{Model: core.Model{Id: conditionId}}, sql.Options{})

	if err := s.r.DiscountConditionRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *DiscountConditionService) ResolveConditionType(data *models.DiscountCondition) (*models.DiscountCondition, *utils.ApplictaionError) {
	switch {
	case data.Products != nil:
		data.Type = models.DiscountConditionTypeRroducts
		return data, nil
	case data.ProductCollections != nil:
		data.Type = models.DiscountConditionTypeProductCollections
		return data, nil
	case data.ProductTypes != nil:
		data.Type = models.DiscountConditionTypeProductTypes
		return data, nil
	case data.ProductTags != nil:
		data.Type = models.DiscountConditionTypeProductTags
		return data, nil
	case data.CustomerGroups != nil:
		data.Type = models.DiscountConditionTypeCustomerGroups
		return data, nil
	default:
		return nil, nil
	}
}

func (s *DiscountConditionService) UpsertCondition(data *models.DiscountCondition, overrideExisting bool) ([]models.DiscountCondition, *utils.ApplictaionError) {
	resolvedConditionType, err := s.ResolveConditionType(data)
	if err != nil {
		return nil, err
	}

	if resolvedConditionType == nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`Missing one of products, collections, tags, types or customer groups in data`,
			"500",
			nil,
		)
	}

	if data.Id != uuid.Nil {
		resolvedCondition, err := s.Retrieve(data.Id, sql.Options{})
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
		Type:     data.Type,
	}

	if err := s.r.DiscountConditionRepository().Save(s.ctx, created); err != nil {
		return nil, err
	}

	return s.r.DiscountConditionRepository().AddConditionResources(s.ctx, created.Id, resolvedConditionType, overrideExisting)
}

func (s *DiscountConditionService) RemoveResources(data *models.DiscountCondition) *utils.ApplictaionError {
	resolvedConditionType, err := s.ResolveConditionType(data)
	if err != nil {
		return err
	}

	if resolvedConditionType == nil {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			`Missing one of products, collections, tags, types or customer groups in data`,
			"500",
			nil,
		)
	}

	resolvedCondition, err := s.Retrieve(data.Id, sql.Options{})
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
	condition, err := s.Retrieve(conditionId, sql.Options{})
	if err != nil {
		return err
	}

	if err := s.r.DiscountConditionRepository().Remove(s.ctx, condition); err != nil {
		return err
	}

	return nil
}

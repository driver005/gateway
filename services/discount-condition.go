package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/repository"
	"github.com/google/uuid"
)

type DiscountConditionService struct {
	ctx  context.Context
	repo *repository.DiscountConditionRepo
}

func NewDiscountConditionService(
	ctx context.Context,
	repo *repository.DiscountConditionRepo,
) *DiscountConditionService {
	return &DiscountConditionService{
		ctx,
		repo,
	}
}

func (s *DiscountConditionService) Retrieve(conditionId uuid.UUID, config repository.Options) (*models.DiscountCondition, error) {
	if conditionId == uuid.Nil {
		return nil, errors.New(`"conditionId" must be defined`)
	}

	var res *models.DiscountCondition
	query := repository.BuildQuery[models.DiscountCondition](models.DiscountCondition{Model: core.Model{Id: conditionId}}, repository.Options{})

	if err := s.repo.FindOne(s.ctx, res, query); err != nil {
		return nil, fmt.Errorf("DiscountCondition with id %s was not found", conditionId)
	}

	return res, nil
}

func (s *DiscountConditionService) ResolveConditionType(data *models.DiscountCondition) (*models.DiscountCondition, error) {
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

func (s *DiscountConditionService) UpsertCondition(data *models.DiscountCondition, overrideExisting bool) ([]models.DiscountCondition, error) {
	resolvedConditionType, err := s.ResolveConditionType(data)
	if err != nil {
		return nil, err
	}

	if resolvedConditionType == nil {
		return nil, errors.New(`Missing one of products, collections, tags, types or customer groups in data`)
	}

	if data.Id != uuid.Nil {
		resolvedCondition, err := s.Retrieve(data.Id, repository.Options{})
		if err != nil {
			return nil, err
		}

		if data.Operator != resolvedCondition.Operator {
			resolvedCondition.Operator = data.Operator
			err = s.repo.Save(s.ctx, resolvedCondition)
			if err != nil {
				return nil, err
			}
		}

		return s.repo.AddConditionResources(s.ctx, data.Id, resolvedConditionType, overrideExisting)
	}

	created := &models.DiscountCondition{
		Model: core.Model{
			Id: data.Id,
		},
		Operator: data.Operator,
		Type:     data.Type,
	}

	if err := s.repo.Save(s.ctx, created); err != nil {
		return nil, err
	}

	return s.repo.AddConditionResources(s.ctx, created.Id, resolvedConditionType, overrideExisting)
}

func (s *DiscountConditionService) RemoveResources(data *models.DiscountCondition) error {
	resolvedConditionType, err := s.ResolveConditionType(data)
	if err != nil {
		return err
	}

	if resolvedConditionType == nil {
		return errors.New(`Missing one of products, collections, tags, types or customer groups in data`)
	}

	resolvedCondition, err := s.Retrieve(data.Id, repository.Options{})
	if err != nil {
		return err
	}

	if data.Operator != resolvedCondition.Operator {
		resolvedCondition.Operator = data.Operator
		if err := s.repo.Save(s.ctx, resolvedCondition); err != nil {
			return err
		}
	}

	return s.repo.RemoveConditionResources(s.ctx, data.Id, resolvedConditionType)
}

func (s *DiscountConditionService) Delete(conditionId uuid.UUID) error {
	condition, err := s.Retrieve(conditionId, repository.Options{})
	if err != nil {
		return err
	}

	if err := s.repo.Remove(s.ctx, condition); err != nil {
		return err
	}

	return nil
}

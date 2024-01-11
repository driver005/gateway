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

type OrderItemChangeService struct {
	ctx context.Context
	r   Registry
}

func NewOrderItemChangeService(
	r Registry,
) *OrderItemChangeService {
	return &OrderItemChangeService{
		context.Background(),
		r,
	}
}

func (s *OrderItemChangeService) SetContext(context context.Context) *OrderItemChangeService {
	s.ctx = context
	return s
}

func (s *OrderItemChangeService) Retrieve(id uuid.UUID, config sql.Options) (*models.OrderItemChange, *utils.ApplictaionError) {
	if id == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"id" must be defined`,
			"500",
			nil,
		)
	}
	var res *models.OrderItemChange

	query := sql.BuildQuery(models.OrderItemChange{Model: core.Model{Id: id}}, config)

	if err := s.r.OrderItemChangeRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *OrderItemChangeService) List(selector models.OrderItemChange, config sql.Options) ([]models.OrderItemChange, *utils.ApplictaionError) {
	var res []models.OrderItemChange

	if reflect.DeepEqual(config, sql.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(50)
		config.Order = gox.NewString("created_at DESC")
	}

	query := sql.BuildQuery(selector, config)

	if err := s.r.OrderItemChangeRepository().Find(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *OrderItemChangeService) Create(data *models.OrderItemChange) (*models.OrderItemChange, *utils.ApplictaionError) {
	if err := s.r.OrderItemChangeRepository().Save(s.ctx, data); err != nil {
		return nil, err
	}
	// err = s.eventBus_.Emit(OrderItemChangeService.Events.CREATED, map[string]interface{}{"id": change.id})
	// if err != nil {
	// 	return nil, err
	// }
	return data, nil
}

func (s *OrderItemChangeService) Delete(itemChangeIds uuid.UUIDs) *utils.ApplictaionError {
	if len(itemChangeIds) == 0 {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			"itemChangeIds cannot be empty",
			"500",
			nil,
		)
	}

	var changes []models.OrderItemChange

	query := sql.BuildQuery(models.OrderItemChange{}, sql.Options{
		Specification: []sql.Specification{sql.In("id", itemChangeIds)},
	})

	if err := s.r.OrderItemChangeRepository().Find(s.ctx, changes, query); err != nil {
		return err
	}

	var lineItemToRemove []models.OrderItemChange
	var lineItemIdsToRemove uuid.UUIDs
	for _, change := range changes {
		if change.LineItemId.UUID != uuid.Nil {
			lineItemToRemove = append(lineItemToRemove, change)
			lineItemIdsToRemove = append(lineItemIdsToRemove, change.LineItemId.UUID)
		}
	}
	if err := s.r.OrderItemChangeRepository().DeleteSlice(s.ctx, lineItemToRemove); err != nil {
		return err
	}
	for _, id := range lineItemIdsToRemove {
		if err := s.r.LineItemService().SetContext(s.ctx).Delete(id); err != nil {
			return err
		}
	}
	if err := s.r.TaxProviderService().SetContext(s.ctx).ClearLineItemsTaxLines(lineItemIdsToRemove); err != nil {
		return err
	}
	// err = s.eventBus_.Emit(OrderItemChangeService.Events.DELETED, map[string]interface{}{"ids": itemChangeIds})
	// if err != nil {
	// 	return err
	// }
	return nil
}

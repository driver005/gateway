package services

import (
	"context"
	"fmt"
	"reflect"
	"slices"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type OrderEditService struct {
	ctx context.Context
	r   Registry
}

func NewOrderEditService(
	r Registry,
) *OrderEditService {
	return &OrderEditService{
		context.Background(),
		r,
	}
}

func (s *OrderEditService) SetContext(context context.Context) *OrderEditService {
	s.ctx = context
	return s
}

func (s *OrderEditService) Retrieve(id uuid.UUID, config *sql.Options) (*models.OrderEdit, *utils.ApplictaionError) {
	if id == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"id" must be defined`,
			nil,
		)
	}
	var res *models.OrderEdit
	query := sql.BuildQuery(models.OrderEdit{Model: core.Model{Id: id}}, config)
	if err := s.r.OrderEditRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *OrderEditService) ListAndCount(selector *types.FilterableOrderEdit, config *sql.Options) ([]models.OrderEdit, *int64, *utils.ApplictaionError) {
	var res []models.OrderEdit

	if !reflect.ValueOf(config.Q).IsZero() {
		v := sql.ILike(config.Q)
		selector.InternalNote = v
	}

	query := sql.BuildQuery(selector, config)

	count, err := s.r.OrderEditRepository().FindAndCount(s.ctx, &res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

func (s *OrderEditService) List(selector *types.FilterableOrderEdit, config *sql.Options) ([]models.OrderEdit, *utils.ApplictaionError) {
	orderEdits, _, err := s.ListAndCount(selector, config)
	if err != nil {
		return nil, err
	}

	return orderEdits, nil
}

func (s *OrderEditService) Create(data *types.CreateOrderEditInput, createdBy uuid.UUID) (*models.OrderEdit, *utils.ApplictaionError) {
	activeOrderEdit, err := s.RetrieveActive(data.OrderId, &sql.Options{})
	if err != nil {
		return nil, err
	}
	if activeOrderEdit != nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			fmt.Sprintf("An active order edit already exists for the order %s", data.OrderId),
			nil,
		)
	}

	orderEdit := &models.OrderEdit{
		OrderId:      uuid.NullUUID{UUID: data.OrderId},
		InternalNote: data.InternalNote,
		CreatedBy:    createdBy,
	}
	if err := s.r.OrderEditRepository().Save(s.ctx, orderEdit); err != nil {
		return nil, err
	}

	orderLineItems, err := s.r.LineItemService().SetContext(s.ctx).List(models.LineItem{OrderId: uuid.NullUUID{UUID: data.OrderId}}, &sql.Options{Selects: []string{"id"}})
	if err != nil {
		return nil, err
	}
	var lineItemIds uuid.UUIDs
	for _, item := range orderLineItems {
		lineItemIds = append(lineItemIds, item.Id)
	}
	_, err = s.r.LineItemService().SetContext(s.ctx).CloneTo(lineItemIds, &models.LineItem{OrderEditId: uuid.NullUUID{UUID: orderEdit.Id}}, map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	// err = s.eventBusService_.WithTransaction(transactionManager).Emit(OrderEditServiceEvents.CREATED, map[string]interface{}{"id": orderEdit.Id})
	// if err != nil {
	// 	return nil, err
	// }

	return orderEdit, nil

}

func (s *OrderEditService) Update(id uuid.UUID, data *models.OrderEdit) (*models.OrderEdit, *utils.ApplictaionError) {
	orderEdit, err := s.Retrieve(id, &sql.Options{})
	if err != nil {
		return nil, err
	}

	data.Id = orderEdit.Id

	if err := s.r.OrderEditRepository().Save(s.ctx, data); err != nil {
		return nil, err
	}

	// err = s.eventBusService.Emit(OrderEditServiceEvents.UPDATED, map[string]interface{}{"id": result.Id})
	// if err != nil {
	// 	return nil, err
	// }

	return data, nil
}

func (s *OrderEditService) Delete(id uuid.UUID) *utils.ApplictaionError {
	edit, err := s.Retrieve(id, &sql.Options{})
	if err != nil {
		return nil
	}

	if edit.Status != models.OrderEditStatusCreated {
		return utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			fmt.Sprintf("Cannot delete order edit with status %s", edit.Status),
			nil,
		)
	}

	err = s.DeleteClonedItems(id)
	if err != nil {
		return err
	}

	if err := s.r.OrderEditRepository().Remove(s.ctx, edit); err != nil {
		return err
	}

	return nil
}

func (s *OrderEditService) Decline(id uuid.UUID, declinedBy uuid.UUID, declinedReason string) (*models.OrderEdit, *utils.ApplictaionError) {
	orderEdit, err := s.Retrieve(id, &sql.Options{})
	if err != nil {
		return nil, err
	}

	if orderEdit.Status == models.OrderEditStatusDeclined {
		return orderEdit, nil
	}

	if orderEdit.Status != models.OrderEditStatusRequested {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			fmt.Sprintf("Cannot decline an order edit with status %s.", orderEdit.Status),
			nil,
		)
	}

	now := time.Now()

	orderEdit.DeclinedAt = &now
	orderEdit.DeclinedBy = declinedBy
	orderEdit.DeclinedReason = declinedReason
	if err := s.r.OrderEditRepository().Save(s.ctx, orderEdit); err != nil {
		return nil, err
	}

	// err = s.eventBusService.Emit(OrderEditServiceEvents.DECLINED, map[string]interface{}{"id": result.Id})
	// if err != nil {
	// 	return nil, err
	// }

	return orderEdit, nil
}

func (s *OrderEditService) UpdateLineItem(id uuid.UUID, itemId uuid.UUID, quantity int) *utils.ApplictaionError {
	orderEdit, err := s.Retrieve(id, &sql.Options{Selects: []string{"id", "order_id", "created_at", "requested_at", "confirmed_at", "declined_at", "canceled_at"}})
	if err != nil {
		return err
	}

	isOrderEditActive := s.isOrderEditActive(orderEdit)
	if !isOrderEditActive {
		return utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			fmt.Sprintf("Can not update an item on the order edit %s with the status %s", id, orderEdit.Status),
			nil,
		)
	}

	lineItem, err := s.r.LineItemService().SetContext(s.ctx).Retrieve(itemId, &sql.Options{Selects: []string{"id", "order_edit_id", "original_item_id"}})
	if err != nil {
		return err
	}

	if lineItem.OrderEditId.UUID != id {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			fmt.Sprintf("Invalid line item id %s it does not belong to the same order edit %s.", itemId, orderEdit.OrderId.UUID),
			nil,
		)
	}

	var change *models.OrderItemChange
	changes, err := s.r.OrderItemChangeService().List(models.OrderItemChange{LineItemId: uuid.NullUUID{UUID: itemId}}, &sql.Options{Selects: []string{"line_item_id", "original_line_item_id"}})
	if err != nil {
		return err
	}
	if len(changes) > 0 {
		change = &changes[len(changes)-1]
	} else {
		change, err = s.r.OrderItemChangeService().Create(&types.CreateOrderEditItemChangeInput{
			Type:               models.OrderEditStatusItemUpdate,
			OrderEditId:        id,
			OriginalLineItemId: lineItem.OriginalItemId.UUID,
			LineItemId:         itemId,
		})
		if err != nil {
			return err
		}
	}

	_, err = s.r.LineItemService().SetContext(s.ctx).Update(change.LineItemId.UUID, nil, &models.LineItem{Quantity: quantity}, &sql.Options{})
	if err != nil {
		return err
	}

	err = s.RefreshAdjustments(id, true)
	if err != nil {
		return err
	}

	return nil
}

func (s *OrderEditService) RemoveLineItem(id uuid.UUID, lineItemId uuid.UUID) *utils.ApplictaionError {
	orderEdit, err := s.Retrieve(id, &sql.Options{Selects: []string{"id", "created_at", "requested_at", "confirmed_at", "declined_at", "canceled_at"}})
	if err != nil {
		return err
	}

	isOrderEditActive := s.isOrderEditActive(orderEdit)
	if !isOrderEditActive {
		return utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			fmt.Sprintf("Can not update an item on the order edit %s with the status %s", id, orderEdit.Status),
			nil,
		)
	}

	lineItem, err := s.r.LineItemService().SetContext(s.ctx).Retrieve(lineItemId, &sql.Options{Selects: []string{"id", "order_edit_id", "original_item_id"}})
	if err != nil {
		return nil
	}

	if lineItem.OrderEditId.UUID != id || lineItem.OriginalItemId.UUID == uuid.Nil {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			fmt.Sprintf("Invalid line item id %s it does not belong to the same order edit %s.", lineItemId, orderEdit.OrderId.UUID),
			nil,
		)
	}

	err = s.r.LineItemService().SetContext(s.ctx).DeleteWithTaxLines(lineItem.Id)
	if err != nil {
		return err
	}

	err = s.RefreshAdjustments(id, false)
	if err != nil {
		return err
	}

	_, err = s.r.OrderItemChangeService().Create(&types.CreateOrderEditItemChangeInput{
		OriginalLineItemId: lineItem.OriginalItemId.UUID,
		Type:               models.OrderEditStatusItemRemove,
		OrderEditId:        orderEdit.Id,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *OrderEditService) RefreshAdjustments(id uuid.UUID, preserveCustomAdjustments bool) *utils.ApplictaionError {
	orderEdit, err := s.Retrieve(id, &sql.Options{Relations: []string{"items", "items.variant", "items.adjustments", "items.tax_lines", "order", "order.customer", "order.discounts", "order.discounts.rule", "order.gift_cards", "order.region", "order.shipping_address", "order.shipping_methods"}})
	if err != nil {
		return err
	}

	var clonedItemAdjustmentIds uuid.UUIDs
	for _, item := range orderEdit.Items {
		if len(item.Adjustments) > 0 {
			for _, adjustment := range item.Adjustments {
				if preserveCustomAdjustments {
					clonedItemAdjustmentIds = append(clonedItemAdjustmentIds, adjustment.Id)
				}
			}
		}
	}

	err = s.r.LineItemAdjustmentService().SetContext(s.ctx).DeleteSlice(clonedItemAdjustmentIds, nil)
	if err != nil {
		return err
	}

	localCart := &models.Cart{
		Items: orderEdit.Items,
	}

	copier.CopyWithOption(&localCart, orderEdit.Order, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	_, _, err = s.r.LineItemAdjustmentService().SetContext(s.ctx).CreateAdjustments(localCart, nil)
	if err != nil {
		return err
	}

	return nil
}

func (s *OrderEditService) DecorateTotals(orderEdit *models.OrderEdit) (*models.OrderEdit, *utils.ApplictaionError) {
	edit, err := s.Retrieve(orderEdit.Id, &sql.Options{Selects: []string{"id", "order_id", "items"}, Relations: []string{"items", "items.tax_lines", "items.adjustments", "items.variant"}})
	if err != nil {
		return nil, err
	}

	order, err := s.r.OrderService().SetContext(s.ctx).RetrieveById(edit.OrderId.UUID, &sql.Options{Relations: []string{"discounts", "discounts.rule", "gift_cards", "region", "items", "items.tax_lines", "items.adjustments", "items.variant", "region.tax_rates", "shipping_methods", "shipping_methods.shipping_option", "shipping_methods.tax_lines"}})
	if err != nil {
		return nil, err
	}

	order.Items = edit.Items

	computedOrder, err := s.r.OrderService().SetContext(s.ctx).decorateTotals(order, nil, types.TotalsContext{})
	if err != nil {
		return nil, err
	}

	orderEdit.Items = computedOrder.Items
	orderEdit.DiscountTotal = computedOrder.DiscountTotal
	orderEdit.GiftCardTotal = computedOrder.GiftCardTotal
	orderEdit.GiftCardTaxTotal = computedOrder.GiftCardTaxTotal
	orderEdit.ShippingTotal = computedOrder.ShippingTotal
	orderEdit.Subtotal = computedOrder.Subtotal
	orderEdit.TaxTotal = computedOrder.TaxTotal
	orderEdit.Total = computedOrder.Total
	orderEdit.DifferenceDue = computedOrder.Total - order.Total

	return orderEdit, nil
}

func (s *OrderEditService) AddLineItem(id uuid.UUID, data *types.AddOrderEditLineItemInput) *utils.ApplictaionError {
	orderEdit, err := s.Retrieve(id, &sql.Options{Relations: []string{"order", "order.region"}})
	if err != nil {
		return err
	}

	if !s.isOrderEditActive(orderEdit) {
		return utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			fmt.Sprintf("Can not add an item to the edit with status %s", orderEdit.Status),
			nil,
		)
	}

	lineItemData, err := s.r.LineItemService().SetContext(s.ctx).Generate(data.VariantId, nil, orderEdit.Order.RegionId.UUID, data.Quantity, types.GenerateLineItemContext{})
	if err != nil {
		return err
	}

	item, err := s.r.LineItemService().SetContext(s.ctx).Create(lineItemData)
	if err != nil {
		return err
	}

	lineItem, err := s.r.LineItemService().SetContext(s.ctx).Retrieve(item[0].Id, &sql.Options{Relations: []string{"variant.product.profiles"}})
	if err != nil {
		return err
	}

	err = s.RefreshAdjustments(id, false)
	if err != nil {
		return err
	}

	_, err = s.r.OrderItemChangeService().Create(&types.CreateOrderEditItemChangeInput{
		Type:        models.OrderEditStatusItemAdd,
		LineItemId:  lineItem.Id,
		OrderEditId: id,
	})
	if err != nil {
		return err
	}

	calcContext, err := s.r.TotalsService().GetCalculationContext(types.CalculationContextData{
		Discounts:       orderEdit.Order.Discounts,
		Items:           []models.LineItem{*lineItem},
		Customer:        orderEdit.Order.Customer,
		Region:          orderEdit.Order.Region,
		ShippingAddress: orderEdit.Order.ShippingAddress,
		Swaps:           orderEdit.Order.Swaps,
		Claims:          orderEdit.Order.Claims,
		ShippingMethods: orderEdit.Order.ShippingMethods,
	}, CalculationContextOptions{ExcludeShipping: true})
	if err != nil {
		return err
	}

	_, _, err = s.r.TaxProviderService().SetContext(s.ctx).CreateTaxLines(nil, []models.LineItem{*lineItem}, calcContext)
	if err != nil {
		return err
	}

	return nil
}

func (s *OrderEditService) DeleteItemChange(id uuid.UUID, itemChangeId uuid.UUID) *utils.ApplictaionError {
	itemChange, err := s.r.OrderItemChangeService().Retrieve(itemChangeId, &sql.Options{Selects: []string{"id", "order_edit_id"}})
	if err != nil {
		return err
	}

	orderEdit, err := s.Retrieve(id, &sql.Options{Selects: []string{"id", "confirmed_at", "canceled_at"}})
	if err != nil {
		return err
	}

	if orderEdit.Id != itemChange.OrderEditId.UUID {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			fmt.Sprintf("The item change you are trying to delete doesn't belong to the models.OrderEdit with id: %s.", id),
			nil,
		)
	}

	if orderEdit.ConfirmedAt != nil || orderEdit.CanceledAt != nil {
		return utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			fmt.Sprintf("Cannot delete and item change from a %s order edit", orderEdit.Status),
			nil,
		)
	}

	return s.r.OrderItemChangeService().Delete(uuid.UUIDs{itemChangeId})
}

func (s *OrderEditService) RequestConfirmation(id uuid.UUID, requestedBy uuid.UUID) (*models.OrderEdit, *utils.ApplictaionError) {
	orderEdit, err := s.Retrieve(id, &sql.Options{Relations: []string{"changes", "changes.original_line_item", "changes.original_line_item.variant"}, Selects: []string{"id", "order_id", "requested_at"}})
	if err != nil {
		return nil, err
	}

	if len(orderEdit.Changes) == 0 {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Cannot request a confirmation on an edit with no changes",
			nil,
		)
	}

	if orderEdit.RequestedAt != nil {
		return orderEdit, nil
	}

	now := time.Now()
	orderEdit.RequestedAt = &now
	orderEdit.RequestedBy = requestedBy
	if err := s.r.OrderEditRepository().Save(s.ctx, orderEdit); err != nil {
		return nil, err
	}

	// err = s.eventBusService.Emit(OrderEditServiceEvents.REQUESTED, map[string]interface{}{"id": id})
	// if err != nil {
	// 	return nil, err
	// }

	return orderEdit, nil
}

func (s *OrderEditService) Cancel(id uuid.UUID, canceledBy uuid.UUID) (*models.OrderEdit, *utils.ApplictaionError) {
	orderEdit, err := s.Retrieve(id, &sql.Options{})
	if err != nil {
		return nil, err
	}

	if orderEdit.Status == models.OrderEditStatusCanceled {
		return orderEdit, nil
	}

	if slices.ContainsFunc([]models.OrderEditStatus{models.OrderEditStatusConfirmed, models.OrderEditStatusDeclined}, func(edit models.OrderEditStatus) bool {
		return edit == orderEdit.Status
	}) {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			fmt.Sprintf("Cannot cancel order edit with status %s", orderEdit.Status),
			nil,
		)
	}

	now := time.Now()
	orderEdit.CanceledAt = &now
	orderEdit.CanceledBy = canceledBy
	if err := s.r.OrderEditRepository().Save(s.ctx, orderEdit); err != nil {
		return nil, err
	}

	// err = s.eventBusService.Emit(OrderEditServiceEvents.CANCELED, map[string]interface{}{"id": id})
	// if err != nil {
	// 	return nil, err
	// }

	return orderEdit, nil
}

func (s *OrderEditService) Confirm(id uuid.UUID, confirmedBy uuid.UUID) (*models.OrderEdit, *utils.ApplictaionError) {
	orderEdit, err := s.Retrieve(id, &sql.Options{})
	if err != nil {
		return nil, err
	}

	if slices.ContainsFunc([]models.OrderEditStatus{models.OrderEditStatusCanceled, models.OrderEditStatusDeclined}, func(edit models.OrderEditStatus) bool {
		return edit == orderEdit.Status
	}) {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			fmt.Sprintf("Cannot confirm an order edit with status %s", orderEdit.Status),
			nil,
		)
	}

	if orderEdit.Status == models.OrderEditStatusConfirmed {
		return orderEdit, nil
	}

	var originalOrderLineItems []models.LineItem
	original, err := s.r.LineItemService().SetContext(s.ctx).Update(
		id,
		&models.LineItem{OrderId: orderEdit.OrderId},
		&models.LineItem{OrderId: uuid.NullUUID{}},
		&sql.Options{
			Specification: []sql.Specification{
				sql.Not(sql.Equal("order_edit_id", id)),
				sql.IsNull("order_edit_id"),
			},
		},
	)
	if err != nil {
		return nil, err
	}
	originalOrderLineItems = append(originalOrderLineItems, *original)
	lineItem, err := s.r.LineItemService().SetContext(s.ctx).Update(
		uuid.Nil,
		&models.LineItem{OrderEditId: uuid.NullUUID{UUID: orderEdit.Id}},
		&models.LineItem{OrderId: orderEdit.OrderId},
		&sql.Options{},
	)
	if err != nil {
		return nil, err
	}
	originalOrderLineItems = append(originalOrderLineItems, *lineItem)

	now := time.Now()
	orderEdit.ConfirmedAt = &now
	orderEdit.ConfirmedBy = confirmedBy
	if err := s.r.OrderEditRepository().Save(s.ctx, orderEdit); err != nil {
		return nil, err
	}

	if s.r.InventoryService() != nil {
		var itemIds uuid.UUIDs
		for _, item := range originalOrderLineItems {
			itemIds = append(itemIds, item.Id)
		}
		err = s.r.InventoryService().DeleteReservationItemsByLineItem(s.ctx, itemIds)
		if err != nil {
			return nil, err
		}
	}

	// err = s.eventBusService.Emit(OrderEditServiceEvents.CONFIRMED, map[string]interface{}{"id": id})
	// if err != nil {
	// 	return nil, err
	// }

	return orderEdit, nil
}

func (s *OrderEditService) RetrieveActive(orderId uuid.UUID, config *sql.Options) (*models.OrderEdit, *utils.ApplictaionError) {
	if orderId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"orderId" must be defined`,
			nil,
		)
	}
	var res *models.OrderEdit
	query := sql.BuildQuery(models.OrderEdit{
		OrderId:     uuid.NullUUID{UUID: orderId},
		ConfirmedAt: &time.Time{},
		CanceledAt:  &time.Time{},
		DeclinedAt:  &time.Time{},
	}, config)
	if err := s.r.OrderEditRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *OrderEditService) DeleteClonedItems(id uuid.UUID) *utils.ApplictaionError {
	clonedLineItems, err := s.r.LineItemService().SetContext(s.ctx).List(models.LineItem{
		OrderEditId: uuid.NullUUID{UUID: id},
	}, &sql.Options{
		Selects:   []string{"id", "tax_lines", "adjustments"},
		Relations: []string{"tax_lines", "adjustments"},
	})
	if err != nil {
		return err
	}
	var clonedItemIds uuid.UUIDs
	for _, item := range clonedLineItems {
		clonedItemIds = append(clonedItemIds, item.Id)
	}
	orderEdit, err := s.Retrieve(id, &sql.Options{
		Selects:   []string{"id", "changes"},
		Relations: []string{"changes", "changes.original_line_item", "changes.original_line_item.variant"},
	})
	if err != nil {
		return err
	}
	var changeIds uuid.UUIDs
	for _, change := range orderEdit.Changes {
		changeIds = append(changeIds, change.Id)
	}

	err = s.r.OrderItemChangeService().Delete(changeIds)
	if err != nil {
		return err
	}

	for _, id := range clonedItemIds {
		if err := s.r.LineItemAdjustmentService().SetContext(s.ctx).Delete(uuid.Nil, &models.LineItemAdjustment{
			ItemId: uuid.NullUUID{UUID: id},
		}, &sql.Options{}); err != nil {
			return err
		}

	}
	if err := s.r.TaxProviderService().SetContext(s.ctx).ClearLineItemsTaxLines(clonedItemIds); err != nil {
		return err
	}
	for _, id := range clonedItemIds {
		if err = s.r.LineItemService().SetContext(s.ctx).Delete(id); err != nil {
			return err
		}
	}
	return nil
}

func (s *OrderEditService) isOrderEditActive(orderEdit *models.OrderEdit) bool {
	return !(orderEdit.Status == models.OrderEditStatusConfirmed || orderEdit.Status == models.OrderEditStatusCanceled || orderEdit.Status == models.OrderEditStatusDeclined)
}

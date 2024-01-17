package services

import (
	"context"
	"reflect"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/icza/gox/gox"
)

type DraftOrderService struct {
	ctx context.Context
	r   Registry
}

func NewDraftOrderService(
	r Registry,
) *DraftOrderService {
	return &DraftOrderService{
		context.Background(),
		r,
	}
}

func (s *DraftOrderService) SetContext(context context.Context) *DraftOrderService {
	s.ctx = context
	return s
}

func (s *DraftOrderService) Retrieve(draftOrderId uuid.UUID, config *sql.Options) (*models.DraftOrder, *utils.ApplictaionError) {
	if draftOrderId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"draftOrderId" must be defined`,
			nil,
		)
	}

	var draftOrder *models.DraftOrder

	query := sql.BuildQuery(models.DraftOrder{Model: core.Model{Id: draftOrderId}}, config)
	if err := s.r.DraftOrderRepository().FindOne(s.ctx, draftOrder, query); err != nil {
		return nil, err
	}

	return draftOrder, nil
}

func (s *DraftOrderService) RetrieveByCartId(cartId uuid.UUID, config *sql.Options) (*models.DraftOrder, *utils.ApplictaionError) {
	var draftOrder *models.DraftOrder

	query := sql.BuildQuery(models.DraftOrder{CartId: uuid.NullUUID{UUID: cartId}}, config)
	if err := s.r.DraftOrderRepository().FindOne(s.ctx, draftOrder, query); err != nil {
		return nil, err
	}

	return draftOrder, nil
}

func (s *DraftOrderService) Delete(draftOrderId uuid.UUID) *utils.ApplictaionError {
	data, err := s.Retrieve(draftOrderId, &sql.Options{})
	if err != nil {
		return err
	}

	if err := s.r.DraftOrderRepository().SoftRemove(s.ctx, data); err != nil {
		return err
	}

	return nil
}

func (s *DraftOrderService) ListAndCount(selector *models.DraftOrder, config *sql.Options) ([]models.DraftOrder, *int64, *utils.ApplictaionError) {
	var res []models.DraftOrder

	if reflect.DeepEqual(config, &sql.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(20)
	}

	query := sql.BuildQuery(selector, config)
	count, err := s.r.DraftOrderRepository().FindAndCount(s.ctx, res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

func (s *DraftOrderService) List(selector *models.DraftOrder, config *sql.Options) ([]models.DraftOrder, *utils.ApplictaionError) {
	var res []models.DraftOrder
	query := sql.BuildQuery(selector, config)
	if err := s.r.DraftOrderRepository().Find(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *DraftOrderService) Create(data types.DraftOrderCreate) (*models.DraftOrder, *utils.ApplictaionError) {
	if data.RegionId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"region_id is required to Create a draft order",
			nil,
		)
	}
	shipping_methods := data.ShippingMethods
	no_notification_order := data.NoNotificationOrder
	items := data.Items
	idempotency_key := data.IdempotencyKey
	discounts := data.Discounts

	createdCart, err := s.r.CartService().SetContext(s.ctx).Create(&types.CartCreateProps{Type: models.CartClaim})
	if err != nil {
		return nil, err
	}

	var draftOrder *models.DraftOrder
	draftOrder.CartId = uuid.NullUUID{UUID: createdCart.Id}
	draftOrder.NoNotificationOrder = no_notification_order
	draftOrder.IdempotencyKey = idempotency_key

	if err := s.r.DraftOrderRepository().Save(s.ctx, draftOrder); err != nil {
		return nil, err
	}

	//TODO: ADD
	// this.eventBus_.withTransaction(transactionManager).emit(DraftOrderService.Events.CREATED, Event{
	// 	ID: result.ID,
	// })

	itemsToGenerate := []types.GenerateInputData{}
	itemsToCreate := []models.LineItem{}
	for _, item := range items {
		if item.VariantId != uuid.Nil {
			itemsToGenerate = append(itemsToGenerate, types.GenerateInputData{
				VariantId: item.VariantId,
				Quantity:  item.Quantity,
				Metadata:  item.Metadata,
				UnitPrice: item.UnitPrice,
			})
			continue
		}
		var price float64
		if item.UnitPrice < 0 {
			price = 0
		} else {
			price = item.UnitPrice
		}
		itemsToCreate = append(itemsToCreate, models.LineItem{
			Model: core.Model{
				Metadata: item.Metadata,
			},
			CartId:         uuid.NullUUID{UUID: createdCart.Id},
			HasShipping:    true,
			Title:          item.Title,
			AllowDiscounts: false,
			UnitPrice:      price,
			Quantity:       item.Quantity,
		})
	}

	if len(itemsToGenerate) > 0 {
		generatedLines, err := s.r.LineItemService().SetContext(s.ctx).Generate(uuid.Nil, itemsToGenerate, data.RegionId, 0, types.GenerateLineItemContext{})
		if err != nil {
			return nil, err
		}
		toCreate := []models.LineItem{}
		for _, line := range generatedLines {
			line.CartId = uuid.NullUUID{UUID: createdCart.Id}
			toCreate = append(toCreate, line)
		}
		s.r.LineItemService().SetContext(s.ctx).Create(toCreate)
	}
	if len(itemsToCreate) > 0 {
		s.r.LineItemService().SetContext(s.ctx).Create(itemsToCreate)
	}
	shippingMethodToCreate := []types.CreateCustomShippingOptionInput{}
	for _, method := range shipping_methods {
		shippingMethodToCreate = append(shippingMethodToCreate, types.CreateCustomShippingOptionInput{
			ShippingOptionId: method.OptionId,
			CartId:           createdCart.Id,
			Price:            method.Price,
		})
	}
	if len(shippingMethodToCreate) > 0 {
		_, err := s.r.CustomShippingOptionService().SetContext(s.ctx).Create(shippingMethodToCreate)
		if err != nil {
			return nil, err
		}
	}
	createdCart, err = s.r.CartService().SetContext(s.ctx).RetrieveWithTotals(createdCart.Id, &sql.Options{
		Relations: []string{
			"shipping_methods",
			"shipping_methods.shipping_option",
			"items.variant.product.profiles",
			"payment_sessions",
		},
	}, TotalsConfig{})
	if err != nil {
		return nil, err
	}
	for _, method := range shipping_methods {
		_, err := s.r.CartService().SetContext(s.ctx).AddShippingMethod(uuid.Nil, createdCart, method.OptionId, method.Data)
		if err != nil {
			return nil, err
		}
	}

	if len(discounts) > 0 {
		_, err := s.r.CartService().SetContext(s.ctx).Update(createdCart.Id, nil, &types.CartUpdateProps{
			Discounts: discounts,
		})
		if err != nil {
			return nil, err
		}
	}
	return draftOrder, nil
}

func (s *DraftOrderService) RegisterCartCompletion(id uuid.UUID, orderId uuid.UUID) (*models.DraftOrder, *utils.ApplictaionError) {
	var draftOrder *models.DraftOrder

	query := sql.BuildQuery(models.DraftOrder{Model: core.Model{Id: id}}, &sql.Options{})
	if err := s.r.DraftOrderRepository().FindOne(s.ctx, draftOrder, query); err != nil {
		return nil, err
	}

	now := time.Now()
	draftOrder.Status = "completed"
	draftOrder.CompletedAt = &now
	draftOrder.OrderId = uuid.NullUUID{UUID: orderId}

	if err := s.r.DraftOrderRepository().Update(s.ctx, draftOrder); err != nil {
		return nil, err
	}

	return draftOrder, nil
}

func (s *DraftOrderService) Update(id uuid.UUID, data *models.DraftOrder) (*models.DraftOrder, *utils.ApplictaionError) {
	var draftOrder *models.DraftOrder

	query := sql.BuildQuery(models.DraftOrder{Model: core.Model{Id: id}}, &sql.Options{})
	if err := s.r.DraftOrderRepository().FindOne(s.ctx, draftOrder, query); err != nil {
		return nil, err
	}

	if draftOrder.Status == "completed" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`Can't Update a draft order which is complete`,
			nil,
		)
	}

	if data.NoNotificationOrder {
		draftOrder.NoNotificationOrder = data.NoNotificationOrder
	}

	if err := s.r.DraftOrderRepository().Update(s.ctx, draftOrder); err != nil {
		return nil, err
	}

	return draftOrder, nil
}

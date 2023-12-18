package services

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/repository"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/icza/gox/gox"
)

type DraftOrderService struct {
	ctx                         context.Context
	repo                        *repository.DraftOrderRepo
	paymentRepository           *repository.PaymentRepo
	orderRepository             *repository.OrderRepo
	eventBusService             *EventBus
	cartService                 *CartService
	lineItemService             *LineItemService
	productVariantService       *ProductVariantService
	shippingOptionService       *ShippingOptionService
	customShippingOptionService *CustomShippingOptionService
}

func NewDraftOrderService(
	ctx context.Context,
	repo *repository.DraftOrderRepo,
	paymentRepository *repository.PaymentRepo,
	orderRepository *repository.OrderRepo,
	eventBusService *EventBus,
	cartService *CartService,
	lineItemService *LineItemService,
	productVariantService *ProductVariantService,
	shippingOptionService *ShippingOptionService,
	customShippingOptionService *CustomShippingOptionService,
) *DraftOrderService {
	return &DraftOrderService{
		ctx:                         ctx,
		repo:                        repo,
		paymentRepository:           paymentRepository,
		orderRepository:             orderRepository,
		eventBusService:             eventBusService,
		cartService:                 cartService,
		lineItemService:             lineItemService,
		productVariantService:       productVariantService,
		shippingOptionService:       shippingOptionService,
		customShippingOptionService: customShippingOptionService,
	}
}

func (s *DraftOrderService) Retrieve(draftOrderId uuid.UUID, config repository.Options) (*models.DraftOrder, error) {
	if draftOrderId == uuid.Nil {
		return nil, errors.New(`"draftOrderId" must be defined`)
	}

	var draftOrder *models.DraftOrder

	query := repository.BuildQuery(models.DraftOrder{Model: core.Model{Id: draftOrderId}}, config)
	if err := s.repo.FindOne(s.ctx, draftOrder, query); err != nil {
		return nil, err
	}

	return draftOrder, nil
}

func (s *DraftOrderService) RetrieveByCartId(cartId uuid.UUID, config repository.Options) (*models.DraftOrder, error) {
	var draftOrder *models.DraftOrder

	query := repository.BuildQuery(models.DraftOrder{CartId: uuid.NullUUID{UUID: cartId}}, config)
	if err := s.repo.FindOne(s.ctx, draftOrder, query); err != nil {
		return nil, err
	}

	return draftOrder, nil
}

func (s *DraftOrderService) Delete(draftOrderId uuid.UUID) error {
	data, err := s.Retrieve(draftOrderId, repository.Options{})
	if err != nil {
		return err
	}

	if err := s.repo.SoftRemove(s.ctx, data); err != nil {
		return err
	}

	return nil
}

func (s *DraftOrderService) ListAndCount(selector *models.DraftOrder, config repository.Options) ([]models.DraftOrder, *int64, error) {
	var res []models.DraftOrder

	if reflect.DeepEqual(config, repository.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(20)
	}

	query := repository.BuildQuery(selector, config)
	count, err := s.repo.FindAndCount(s.ctx, res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

func (s *DraftOrderService) List(selector *models.DraftOrder, config repository.Options) ([]models.DraftOrder, error) {
	var res []models.DraftOrder
	query := repository.BuildQuery(selector, config)
	if err := s.repo.Find(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *DraftOrderService) Create(data types.DraftOrderCreate) (*models.DraftOrder, error) {
	if data.RegionId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"region_id is required to create a draft order",
			"500",
			nil,
		)
	}
	shipping_methods := data.ShippingMethods
	no_notification_order := data.NoNotificationOrder
	items := data.Items
	idempotency_key := data.IdempotencyKey
	discounts := data.Discounts
	rawCart := data.RawCart

	createdCart, err := cartServiceTx.create(models.Cart{
		Type: CartType.DRAFT_ORDER,
	})
	if err != nil {
		return nil, err
	}

	var draftOrder *models.DraftOrder
	draftOrder.CartId = createdCart.Id
	draftOrder.NoNotificationOrder = no_notification_order
	draftOrder.IdempotencyKey = idempotency_key

	if err := s.repo.Save(s.ctx, draftOrder); err != nil {
		return nil, err
	}

	//TODO: ADD
	// this.eventBus_.withTransaction(transactionManager).emit(DraftOrderService.Events.CREATED, Event{
	// 	ID: result.ID,
	// })

	itemsToGenerate := []GenerateInputData{}
	itemsToCreate := []LineItem{}
	for _, item := range items {
		if item.VariantID != nil {
			itemsToGenerate = append(itemsToGenerate, GenerateInputData{
				VariantID: item.VariantID,
				Quantity:  item.Quantity,
				Metadata:  item.Metadata,
				UnitPrice: item.UnitPrice,
			})
			continue
		}
		var price int
		if item.UnitPrice == nil || *item.UnitPrice < 0 {
			price = 0
		} else {
			price = *item.UnitPrice
		}
		itemsToCreate = append(itemsToCreate, LineItem{
			CartID:         createdCart.ID,
			HasShipping:    true,
			Title:          item.Title,
			AllowDiscounts: false,
			UnitPrice:      price,
			Quantity:       item.Quantity,
			Metadata:       item.Metadata,
		})
	}
	promises := []Promise{}
	if len(itemsToGenerate) > 0 {
		generatedLines, err := lineItemServiceTx.generate(itemsToGenerate, GenerateOptions{
			RegionID: data.RegionID,
		})
		if err != nil {
			return nil, err
		}
		toCreate := []LineItem{}
		for _, line := range generatedLines {
			toCreate = append(toCreate, LineItem{
				// ...line,
				CartID: createdCart.ID,
			})
		}
		promises = append(promises, lineItemServiceTx.create(toCreate))
	}
	if len(itemsToCreate) > 0 {
		promises = append(promises, lineItemServiceTx.create(itemsToCreate))
	}
	shippingMethodToCreate := []ShippingMethod{}
	for _, method := range shipping_methods {
		if method.Price != nil {
			shippingMethodToCreate = append(shippingMethodToCreate, ShippingMethod{
				ShippingOptionID: method.OptionID,
				CartID:           createdCart.ID,
				Price:            *method.Price,
			})
			continue
		}
	}
	if len(shippingMethodToCreate) > 0 {
		err := this.customShippingOptionService_.withTransaction(transactionManager).create(shippingMethodToCreate)
		if err != nil {
			return nil, err
		}
	}
	createdCart, err = cartServiceTx.retrieveWithTotals(createdCart.ID, RetrieveOptions{
		Relations: []string{
			"shipping_methods",
			"shipping_methods.shipping_option",
			"items.variant.product.profiles",
			"payment_sessions",
		},
	})
	if err != nil {
		return nil, err
	}
	for _, method := range shipping_methods {
		promises = append(promises, cartServiceTx.addShippingMethod(createdCart, method.OptionID, method.Data))
	}

	if len(discounts) > 0 {
		err := cartServiceTx.update(createdCart.ID, UpdateOptions{
			Discounts: discounts,
		})
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (s *DraftOrderService) RegisterCartCompletion(id uuid.UUID, orderId uuid.UUID) (*models.DraftOrder, error) {
	var draftOrder *models.DraftOrder

	query := repository.BuildQuery(models.DraftOrder{Model: core.Model{Id: id}}, repository.Options{})
	if err := s.repo.FindOne(s.ctx, draftOrder, query); err != nil {
		return nil, err
	}

	draftOrder.Status = "completed"
	draftOrder.CompletedAt = time.Now()
	draftOrder.OrderId = uuid.NullUUID{UUID: orderId}

	if err := s.repo.Update(s.ctx, draftOrder); err != nil {
		return nil, err
	}

	return draftOrder, nil
}

func (s *DraftOrderService) Update(id uuid.UUID, data *models.DraftOrder) (*models.DraftOrder, error) {
	var draftOrder *models.DraftOrder

	query := repository.BuildQuery(models.DraftOrder{Model: core.Model{Id: id}}, repository.Options{})
	if err := s.repo.FindOne(s.ctx, draftOrder, query); err != nil {
		return nil, err
	}

	if draftOrder.Status == "completed" {
		return nil, errors.New(`Can't update a draft order which is complete`)
	}

	if data.NoNotificationOrder {
		draftOrder.NoNotificationOrder = data.NoNotificationOrder
	}

	if err := s.repo.Update(s.ctx, draftOrder); err != nil {
		return nil, err
	}

	return draftOrder, nil
}

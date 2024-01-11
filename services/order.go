package services

import (
	"context"
	"fmt"
	"math"
	"reflect"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/icza/gox/gox"
	"github.com/jinzhu/copier"
	"github.com/samber/lo"
)

type OrderService struct {
	ctx context.Context
	r   Registry
}

func NewOrderService(
	r Registry,
) *OrderService {
	return &OrderService{
		context.Background(),
		r,
	}
}

func (s *OrderService) SetContext(context context.Context) *OrderService {
	s.ctx = context
	return s
}

func (s *OrderService) List(selector models.Order, config sql.Options, q *string) ([]models.Order, *utils.ApplictaionError) {
	orders, _, err := s.ListAndCount(selector, config, q)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *OrderService) ListAndCount(selector models.Order, config sql.Options, q *string) ([]models.Order, *int64, *utils.ApplictaionError) {
	var res []models.Order

	if reflect.DeepEqual(config, sql.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(50)
		config.Order = gox.NewString("created_at DESC")
	}
	var specification []sql.Specification

	if q != nil {
		if config.Relations != nil {
			config.Relations = append(config.Relations, "shipping_address", "customer")
		} else {
			config.Relations = []string{"shipping_address", "customer"}
		}
		v := sql.ILike(*q)

		specification = append(specification, sql.Not(sql.IsNull("customer.id")))
		specification = append(specification, sql.Not(sql.IsNull("shipping_address.id")))

		selector.Email = v
		selector.DisplayId = v
		selector.ShippingAddress.FirstName = v
		selector.Customer.FirstName = v
		selector.Customer.LastName = v
		selector.Customer.Phone = v
	}

	config.Specification = append(config.Specification, specification...)
	query := sql.BuildQuery(selector, config)

	selects, relations, totalsToSelect := s.transformQueryForTotals(config)

	query.Selects = selects
	query.Relations = relations

	count, err := s.r.OrderRepository().FindAndCount(s.ctx, res, query)
	if err != nil {
		return nil, nil, err
	}

	var orders []models.Order
	for _, r := range res {
		order, err := s.decorateTotals(&r, totalsToSelect, types.TotalsContext{})
		if err != nil {
			return nil, nil, err
		}
		orders = append(orders, *order)
	}
	return orders, count, nil
}

func (s *OrderService) Retrieve(selector models.Order, config sql.Options) (*models.Order, *utils.ApplictaionError) {
	var res *models.Order
	query := sql.BuildQuery(selector, config)

	if err := s.r.OrderRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *OrderService) RetrieveWithTotals(selector models.Order, config sql.Options, context types.TotalsContext) (*models.Order, *utils.ApplictaionError) {
	var res *models.Order
	query := sql.BuildQuery(selector, config)

	selects, relations, totalsToSelect := s.transformQueryForTotals(config)
	query.Selects = selects
	query.Relations = relations

	if err := s.r.OrderRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return s.decorateTotals(res, totalsToSelect, context)
}

func (s *OrderService) RetrieveById(id uuid.UUID, config sql.Options) (*models.Order, *utils.ApplictaionError) {
	if id == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"id" must be defined`,
			"500",
			nil,
		)
	}

	return s.Retrieve(models.Order{Model: core.Model{Id: id}}, config)
}

func (s *OrderService) RetrieveLegacy(id uuid.UUID, selector models.Order, config sql.Options) (*models.Order, *utils.ApplictaionError) {
	if id != uuid.Nil {
		selector.Id = id
	}

	return s.Retrieve(selector, config)
}

func (s *OrderService) RetrieveByIdWithTotals(id uuid.UUID, config sql.Options, context types.TotalsContext) (*models.Order, *utils.ApplictaionError) {
	if id == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"id" must be defined`,
			"500",
			nil,
		)
	}

	return s.RetrieveWithTotals(models.Order{Model: core.Model{Id: id}}, config, context)
}

func (s *OrderService) RetrieveByCartId(cartId uuid.UUID, config sql.Options) (*models.Order, *utils.ApplictaionError) {
	if cartId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"cartId" must be defined`,
			"500",
			nil,
		)
	}
	return s.Retrieve(models.Order{CartId: uuid.NullUUID{UUID: cartId}}, config)
}

func (s *OrderService) RetrieveByCartIdWithTotals(cartId uuid.UUID, config sql.Options) (*models.Order, *utils.ApplictaionError) {
	if cartId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"cartId" must be defined`,
			"500",
			nil,
		)
	}
	return s.RetrieveWithTotals(models.Order{CartId: uuid.NullUUID{UUID: cartId}}, config, types.TotalsContext{})
}

func (s *OrderService) RetrieveByExternalId(externalId uuid.UUID, config sql.Options) (*models.Order, *utils.ApplictaionError) {
	if externalId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"externalId" must be defined`,
			"500",
			nil,
		)
	}

	return s.RetrieveWithTotals(models.Order{ExternalId: uuid.NullUUID{UUID: externalId}}, config, types.TotalsContext{})
}

func (s *OrderService) CompleteOrder(id uuid.UUID) (*models.Order, *utils.ApplictaionError) {
	order, err := s.RetrieveById(id, sql.Options{})
	if err != nil {
		return nil, err
	}
	if order.Status == "canceled" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"A canceled order cannot be completed",
			"500",
			nil,
		)
	}
	// s.eventBus_.emit(OrderService.Events.COMPLETED, map[string]interface{}{"id": id, "no_notification": order.NoNotification})
	order.Status = models.OrderStatusCompleted
	if err := s.r.OrderRepository().Save(s.ctx, order); err != nil {
		return nil, err
	}
	return order, nil
}

func (s *OrderService) CreateFromCart(id uuid.UUID, data *models.Cart) (*models.Order, *utils.ApplictaionError) {
	if id != uuid.Nil {
		data.Id = id
	}
	_, err := s.RetrieveByCartId(id, sql.Options{Selects: []string{"id"}})
	if err != nil {
		return nil, err
	}

	var cart *models.Cart
	if id != uuid.Nil {
		cart, err = s.r.CartService().SetContext(s.ctx).RetrieveWithTotals(id, sql.Options{Relations: []string{"region", "payment", "items"}}, TotalsConfig{})
		if err != nil {
			return nil, err
		}
	} else {
		cart = data
	}
	if len(cart.Items) == 0 {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Cannot create order from empty cart",
			"500",
			nil,
		)
	}
	if cart.CustomerId.UUID == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Cannot create an order from the cart without a customer",
			"500",
			nil,
		)
	}
	payment := cart.Payment
	region := cart.Region
	total := cart.Total
	if total != 0 {
		if payment == nil {
			return nil, utils.NewApplictaionError(
				utils.INVALID_DATA,
				"Cart does not contain a payment method",
				"500",
				nil,
			)
		}
		paymentStatus, err := s.r.PaymentProviderService().SetContext(s.ctx).GetStatus(payment)
		if err != nil {
			return nil, err
		}
		if *paymentStatus != models.PaymentSessionStatusAuthorized {
			return nil, utils.NewApplictaionError(
				utils.INVALID_DATA,
				"Payment method is not authorized",
				"500",
				nil,
			)
		}
	}

	var shippingMethods []models.ShippingMethod
	for _, method := range cart.ShippingMethods {
		method.TaxLines = []models.ShippingMethodTaxLine{}
		shippingMethods = append(shippingMethods, method)
	}
	order := &models.Order{
		Model:             core.Model{Metadata: cart.Metadata},
		PaymentStatus:     "awaiting",
		Discounts:         cart.Discounts,
		GiftCards:         cart.GiftCards,
		ShippingMethods:   shippingMethods,
		ShippingAddressId: cart.ShippingAddressId,
		BillingAddressId:  cart.BillingAddressId,
		RegionId:          cart.RegionId,
		Email:             cart.Email,
		CustomerId:        cart.CustomerId,
		CartId:            uuid.NullUUID{UUID: cart.Id},
		CurrencyCode:      region.CurrencyCode,
	}

	feature := true
	featurev2 := true
	if cart.SalesChannelId.UUID != uuid.Nil && feature && !featurev2 {
		order.SalesChannelId = cart.SalesChannelId
	}
	if cart.Type == models.CartDraftOrder {
		draft, err := s.r.DraftOrderService().SetContext(s.ctx).RetrieveByCartId(cart.Id, sql.Options{})
		if err != nil {
			return nil, err
		}
		order.DraftOrderId = uuid.NullUUID{UUID: draft.Id}
		order.NoNotification = draft.NoNotificationOrder
	}

	if err := s.r.OrderRepository().Save(s.ctx, order); err != nil {
		return nil, err
	}
	// if feature && featurev2 {
	// err = s.remoteLink_.create(map[string]interface{}{"orderService": map[string]interface{}{"order_id": order.Id}, "salesChannelService": map[string]interface{}{"sales_channel_id": cart.SalesChannelId}})
	// if err != nil {
	// 	return nil, err
	// }
	// }
	if total != 0 && payment != nil {
		_, err = s.r.PaymentProviderService().SetContext(s.ctx).UpdatePayment(payment.Id, &models.Payment{OrderId: uuid.NullUUID{UUID: order.Id}})
		if err != nil {
			return nil, err
		}
	}
	if cart.Subtotal == 0 || cart.DiscountTotal == 0 {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Unable to compute gift cardable amount during order creation from cart. The cart is missing the subtotal and/or discount_total",
			"500",
			nil,
		)
	}
	giftCardableAmountBalance := cart.GiftCardTotal
	orderedGiftCards := cart.GiftCards

	sort.Slice(orderedGiftCards, func(i, j int) bool {
		var aEnd time.Time
		var bEnd time.Time
		if orderedGiftCards[i].EndsAt != nil {
			aEnd = *orderedGiftCards[i].EndsAt
		} else {
			aEnd = time.Now()
		}

		if orderedGiftCards[j].EndsAt != nil {
			bEnd = *orderedGiftCards[j].EndsAt
		} else {
			bEnd = time.Now()
		}
		return aEnd.Before(bEnd) || (aEnd.Equal(bEnd) && orderedGiftCards[i].Balance < orderedGiftCards[j].Balance)
	})
	for _, giftCard := range orderedGiftCards {
		newGiftCardBalance := math.Max(0, giftCard.Balance-giftCardableAmountBalance)
		giftCardBalanceUsed := giftCard.Balance - newGiftCardBalance
		_, err := s.r.GiftCardService().SetContext(s.ctx).Update(giftCard.Id, &models.GiftCard{Balance: newGiftCardBalance, IsDisabled: newGiftCardBalance == 0})
		if err != nil {
			return nil, err
		}

		_, err = s.r.GiftCardService().SetContext(s.ctx).CreateTransaction(&models.GiftCardTransaction{GiftCardId: uuid.NullUUID{UUID: giftCard.Id}, OrderId: uuid.NullUUID{UUID: order.Id}, Amount: giftCardBalanceUsed, IsTaxable: giftCard.TaxRate != 0, TaxRate: giftCard.TaxRate})
		if err != nil {
			return nil, err
		}
		giftCardableAmountBalance -= giftCardBalanceUsed
		if giftCardableAmountBalance == 0 {
			break
		}
	}

	for _, lineItem := range cart.Items {
		_, err := s.r.LineItemService().SetContext(s.ctx).Update(lineItem.Id, nil, &models.LineItem{OrderId: uuid.NullUUID{UUID: order.Id}}, sql.Options{})
		if err != nil {
			return nil, err
		}
		if lineItem.IsGiftcard {
			giftCard, err := s.createGiftCardsFromLineItem(order, lineItem)
			if err != nil {
				return nil, err
			}
			cart.GiftCards = append(cart.GiftCards, giftCard...)
		}
	}
	for _, method := range cart.ShippingMethods {
		method.TaxLines = []models.ShippingMethodTaxLine{}
		_, err = s.r.ShippingOptionService().UpdateShippingMethod(method.Id, &models.ShippingMethod{OrderId: uuid.NullUUID{UUID: order.Id}})
		if err != nil {
			return nil, err
		}
	}
	// err = s.eventBus_.emit(OrderService.Events.PLACED, map[string]interface{}{"id": order.Id, "no_notification": order.NoNotification})
	// if err != nil {
	// 	return nil, err
	// }
	now := time.Now()
	_, err = s.r.CartService().SetContext(s.ctx).Update(cart.Id, nil, &models.Cart{CompletedAt: &now})
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *OrderService) createGiftCardsFromLineItem(order *models.Order, lineItem models.LineItem) ([]models.GiftCard, *utils.ApplictaionError) {
	var createGiftCardResults []models.GiftCard
	if lineItem.Subtotal == 0 || lineItem.Quantity == 0 {
		return createGiftCardResults, nil
	}
	taxExclusivePrice := lineItem.Subtotal / float64(lineItem.Quantity)
	giftCardTaxRate := 0.0
	for _, taxLine := range lineItem.TaxLines {
		giftCardTaxRate += taxLine.Rate
	}

	for qty := 0; qty < lineItem.Quantity; qty++ {
		res, err := s.r.GiftCardService().SetContext(s.ctx).Create(&models.GiftCard{
			Model: core.Model{
				Metadata: lineItem.Metadata,
			},
			RegionId: order.RegionId,
			OrderId:  uuid.NullUUID{UUID: order.Id},
			Value:    taxExclusivePrice,
			Balance:  taxExclusivePrice,
			TaxRate:  giftCardTaxRate,
		})

		if err != nil {
			return nil, err
		}

		createGiftCardResults = append(createGiftCardResults, *res)
	}
	return createGiftCardResults, nil
}

func (s *OrderService) CreateShipment(
	orderId uuid.UUID,
	fulfillmentId uuid.UUID,
	trackingLinks []models.TrackingLink,
	config struct {
		NoNotification bool
		Metadata       map[string]interface{}
	},
) (*models.Order, *utils.ApplictaionError) {
	metadata := config.Metadata
	noNotification := config.NoNotification

	order, err := s.RetrieveById(orderId, sql.Options{
		Relations: []string{"items"},
	})
	if err != nil {
		return nil, err
	}
	shipment, err := s.r.FulfillmentService().SetContext(s.ctx).Retrieve(fulfillmentId, sql.Options{})
	if err != nil {
		return nil, err
	}
	if order.Status == "canceled" {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"A canceled order cannot be fulfilled as shipped",
			"500",
			nil,
		)
	}
	if shipment == nil || shipment.OrderId.UUID != orderId {
		return nil, utils.NewApplictaionError(
			utils.NOT_FOUND,
			"Could not find fulfillment",
			"500",
			nil,
		)
	}
	evaluatedNoNotification := noNotification
	if !noNotification {
		evaluatedNoNotification = shipment.NoNotification
	}
	shipmentRes, err := s.r.FulfillmentService().SetContext(s.ctx).CreateShipment(fulfillmentId, trackingLinks, &models.Fulfillment{
		Model: core.Model{
			Metadata: metadata,
		},
		NoNotification: evaluatedNoNotification,
	})
	if err != nil {
		return nil, err
	}

	order.FulfillmentStatus = models.FulfillmentStatusShipped
	for _, item := range order.Items {
		shipped, ok := lo.Find(shipmentRes.Items, func(i models.FulfillmentItem) bool {
			return i.ItemId.UUID == item.Id
		})
		if ok {
			shippedQty := item.ShippedQuantity + shipped.Quantity
			if shippedQty != item.Quantity {
				order.FulfillmentStatus = models.FulfillmentStatusPartiallyShipped
			}
			_, err := s.r.LineItemService().SetContext(s.ctx).Update(item.Id, nil, &models.LineItem{
				ShippedQuantity: shippedQty,
			}, sql.Options{})
			if err != nil {
				return nil, err
			}
		} else {
			if item.ShippedQuantity != item.Quantity {
				order.FulfillmentStatus = models.FulfillmentStatusPartiallyShipped
			}
		}
	}

	if err := s.r.OrderRepository().Save(s.ctx, order); err != nil {
		return nil, err
	}
	// err = s.eventBus_.Emit(OrderService_Events_SHIPMENT_CREATED, &struct {
	// 	ID             string
	// 	FulfillmentID  string
	// 	NoNotification bool
	// }{
	// 	ID:             orderId,
	// 	FulfillmentID:  shipmentRes.ID,
	// 	NoNotification: evaluatedNoNotification,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	return order, nil
}

func (s *OrderService) UpdateBillingAddress(
	order *models.Order,
	address *models.Address,
) *utils.ApplictaionError {
	if address.CountryCode != "" {
		countryCode := strings.ToLower(address.CountryCode)
		address.CountryCode = countryCode
	}

	region, err := s.r.RegionService().SetContext(s.ctx).Retrieve(order.RegionId.UUID, sql.Options{
		Relations: []string{"countries"},
	})
	if err != nil {
		return err
	}
	found := false
	for _, country := range region.Countries {
		if address.CountryCode != "" && strings.ToLower(country.Iso2) == address.CountryCode {
			found = true
			break
		}
	}
	if !found {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Shipping country must be in the order region",
			"500",
			nil,
		)
	}
	var addr *models.Address
	if order.BillingAddressId.UUID != uuid.Nil {
		query := sql.BuildQuery(models.Address{Model: core.Model{Id: order.BillingAddressId.UUID}}, sql.Options{})

		if err := s.r.AddressRepository().FindOne(s.ctx, addr, query); err != nil {
			return err
		}

		address.Id = addr.Id

		if err := s.r.AddressRepository().Save(s.ctx, address); err != nil {
			return err
		}
	} else {
		order.BillingAddress = address
	}
	return nil
}

func (s *OrderService) UpdateShippingAddress(
	order *models.Order,
	address *models.Address,
) *utils.ApplictaionError {
	if address.CountryCode != "" {
		countryCode := strings.ToLower(address.CountryCode)
		address.CountryCode = countryCode
	}

	region, err := s.r.RegionService().SetContext(s.ctx).Retrieve(order.RegionId.UUID, sql.Options{
		Relations: []string{"countries"},
	})
	if err != nil {
		return err
	}
	found := false
	for _, country := range region.Countries {
		if address.CountryCode != "" && strings.ToLower(country.Iso2) == address.CountryCode {
			found = true
			break
		}
	}
	if !found {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Shipping country must be in the order region",
			"500",
			nil,
		)
	}
	var addr *models.Address
	if order.ShippingAddressId.UUID != uuid.Nil {
		query := sql.BuildQuery(models.Address{Model: core.Model{Id: order.ShippingAddressId.UUID}}, sql.Options{})

		if err := s.r.AddressRepository().FindOne(s.ctx, addr, query); err != nil {
			return err
		}

		address.Id = addr.Id

		if err := s.r.AddressRepository().Save(s.ctx, address); err != nil {
			return err
		}
	} else {
		order.ShippingAddress = address
	}
	return nil
}

func (s *OrderService) AddShippingMethod(
	orderId uuid.UUID,
	optionId uuid.UUID,
	data map[string]interface{},
	config *models.ShippingMethod,
) (*models.Order, *utils.ApplictaionError) {
	order, err := s.RetrieveByIdWithTotals(
		orderId,
		sql.Options{
			Relations: []string{
				"shipping_methods",
				"shipping_methods.shipping_option",
				"items.variant.product.profiles",
			},
		},
		types.TotalsContext{},
	)
	if err != nil {
		return nil, err
	}
	if order.Status == "canceled" {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"A shipping method cannot be added to a canceled order",
			"500",
			nil,
		)
	}

	config.Order = order
	newMethod, err := s.r.ShippingOptionService().CreateShippingMethod(optionId, data, config)
	if err != nil {
		return nil, err
	}

	methods := []models.ShippingMethod{*newMethod}
	if len(order.ShippingMethods) > 0 {
		for _, sm := range order.ShippingMethods {
			if sm.ShippingOption.ProfileId == newMethod.ShippingOption.ProfileId {
				err := s.r.ShippingOptionService().DeleteShippingMethods([]models.ShippingMethod{sm})
				if err != nil {
					return nil, err
				}
			} else {
				methods = append(methods, sm)
			}
		}
	}
	result, err := s.RetrieveById(orderId, sql.Options{})
	if err != nil {
		return nil, err
	}
	// err = s.eventBus_.Emit(OrderService_Events_UPDATED, &struct {
	// 	ID string
	// }{
	// 	ID: result.ID,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	return result, nil
}

func (s *OrderService) Update(orderId uuid.UUID, update *models.Order) (*models.Order, *utils.ApplictaionError) {
	order, err := s.RetrieveById(orderId, sql.Options{})
	if err != nil {
		return nil, err
	}
	if order.Status == "canceled" {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"A canceled order cannot be updated",
			"500",
			nil,
		)
	}
	if (update.Payments != nil || update.Items != nil) && (order.FulfillmentStatus != "not_fulfilled" || order.PaymentStatus != "awaiting" || order.Status != "pending") {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Can't update shipping, billing, items and payment method when order is processed",
			"500",
			nil,
		)
	}
	if update.Status != "null" || update.FulfillmentStatus != "null" || update.PaymentStatus != "null" {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Can't update order statuses. This will happen automatically. Use metadata in order for additional statuses",
			"500",
			nil,
		)
	}

	if update.ShippingAddress != nil {
		err := s.UpdateShippingAddress(order, update.ShippingAddress)
		if err != nil {
			return nil, err
		}
	}
	if update.BillingAddress != nil {
		err := s.UpdateBillingAddress(order, update.BillingAddress)
		if err != nil {
			return nil, err
		}
	}

	if update.Items != nil {
		for _, item := range update.Items {
			item.OrderId = uuid.NullUUID{UUID: orderId}
			_, err := s.r.LineItemService().SetContext(s.ctx).Create([]models.LineItem{item})
			if err != nil {
				return nil, err
			}
		}
	}

	update.Id = order.Id

	if err := s.r.OrderRepository().Update(s.ctx, update); err != nil {
		return nil, err
	}
	// err = s.eventBus_.Emit(OrderService_Events_UPDATED, &struct {
	// 	ID             string
	// 	NoNotification bool
	// }{
	// 	ID:             orderId,
	// 	NoNotification: order.NoNotification,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	return update, nil
}

func (s *OrderService) Cancel(orderId uuid.UUID) (*models.Order, *utils.ApplictaionError) {
	order, err := s.RetrieveById(orderId, sql.Options{
		Relations: []string{
			"refunds",
			"fulfillments",
			"payments",
			"returns",
			"claims",
			"swaps",
			"items",
		},
	})
	if err != nil {
		return nil, err
	}
	if len(order.Refunds) > 0 {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"models.Order with refund(s) cannot be canceled",
			"500",
			nil,
		)
	}
	throwErrorIf := func(arr interface{}, pred func(obj interface{}) bool, typ string) *utils.ApplictaionError {
		if arr != nil {
			arrVal := reflect.ValueOf(arr)
			for i := 0; i < arrVal.Len(); i++ {
				obj := arrVal.Index(i).Interface()
				if pred(obj) {
					return utils.NewApplictaionError(
						utils.NOT_ALLOWED,
						fmt.Sprintf("All %s must be canceled before canceling an order", typ),
						"500",
						nil,
					)
				}
			}
		}
		return nil
	}
	notCanceled := func(o interface{}) bool {
		return o.(models.Fulfillment).CanceledAt != nil
	}
	if err := throwErrorIf(order.Fulfillments, notCanceled, "fulfillments"); err != nil {
		return nil, err
	}
	if err := throwErrorIf(order.Returns, func(r interface{}) bool {
		return r.(models.Return).Status != "canceled"
	}, "returns"); err != nil {
		return nil, err
	}
	if err := throwErrorIf(order.Claims, notCanceled, "claims"); err != nil {
		return nil, err
	}
	if err := throwErrorIf(order.Swaps, notCanceled, "swaps"); err != nil {
		return nil, err
	}

	for _, item := range order.Items {
		if item.VariantId.UUID != uuid.Nil {
			err := s.r.ProductVariantInventoryService().SetContext(s.ctx).DeleteReservationsByLineItem(item.Id, item.VariantId.UUID, item.Quantity)
			if err != nil {
				return nil, err
			}
		}
	}

	for _, p := range order.Payments {
		_, err := s.r.PaymentProviderService().SetContext(s.ctx).CancelPayment(&p)
		if err != nil {
			return nil, err
		}
	}
	order.Status = models.OrderStatusCanceled
	order.FulfillmentStatus = models.FulfillmentStatusCanceled
	order.PaymentStatus = models.PaymentStatusCanceled
	now := time.Now()
	order.CanceledAt = &now
	if err := s.r.OrderRepository().Save(s.ctx, order); err != nil {
		return nil, err
	}

	// err = s.eventBus_.Emit(OrderService_Events_CANCELED, &struct {
	// 	ID             string
	// 	NoNotification bool
	// }{
	// 	ID:             order.ID,
	// 	NoNotification: order.NoNotification,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	return order, nil
}

func (s *OrderService) CapturePayment(orderId uuid.UUID) (*models.Order, *utils.ApplictaionError) {
	order, err := s.RetrieveById(orderId, sql.Options{
		Relations: []string{"payments"},
	})
	if err != nil {
		return nil, err
	}
	if order.Status == "canceled" {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"A canceled order cannot capture payment",
			"500",
			nil,
		)
	}

	var payments []models.Payment
	for _, p := range order.Payments {
		if p.CapturedAt == nil {
			result, err := s.r.PaymentProviderService().SetContext(s.ctx).CapturePayment(&p)
			if err != nil {
				// err := s.eventBus_.Emit(OrderService_Events_PAYMENT_CAPTURE_FAILED, &struct {
				// 	ID             string
				// 	PaymentID      string
				// 	Error          *utils.ApplictaionError
				// 	NoNotification bool
				// }{
				// 	ID:             orderId,
				// 	PaymentID:      p.ID,
				// 	Error:          err,
				// 	NoNotification: order.NoNotification,
				// })
				// if err != nil {
				// 	return nil, err
				// }
				payments = append(payments, p)
			} else {
				payments = append(payments, *result)
			}
		} else {
			payments = append(payments, p)
		}
	}
	order.Payments = payments
	if paymentsAreCaptured(payments) {
		order.PaymentStatus = models.PaymentStatusCaptured
	} else {
		order.PaymentStatus = models.PaymentStatusRequiresAction
	}
	if err := s.r.OrderRepository().Save(s.ctx, order); err != nil {
		return nil, err
	}
	// if order.PaymentStatus == models.PaymentStatusCaptured {
	// err := s.eventBus_.Emit(OrderService_Events_PAYMENT_CAPTURED, &struct {
	// 	ID             string
	// 	NoNotification bool
	// }{
	// 	ID:             result.ID,
	// 	NoNotification: order.NoNotification,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// }
	return order, nil
}

func (s *OrderService) ValidateFulfillmentLineItem(
	item *models.LineItem,
	quantity int,
) *models.LineItem {
	if item == nil {
		return nil
	}
	if quantity > item.Quantity-item.FulfilledQuantity {
		panic(utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Cannot fulfill more items than have been purchased",
			"500",
			nil,
		))
	}

	item.Quantity = quantity

	return item
}

func paymentsAreCaptured(payments []models.Payment) bool {
	for _, p := range payments {
		if p.CapturedAt == nil {
			return false
		}
	}
	return true
}

func (s *OrderService) CreateFulfillment(id uuid.UUID, itemsToFulfill []types.FulFillmentItemType, config map[string]interface{}) (*models.Order, *utils.ApplictaionError) {
	metadata := config["metadata"].(map[string]interface{})
	no_notification := config["no_notification"].(bool)
	location_id := config["location_id"].(string)

	order, err := s.RetrieveById(id, sql.Options{
		Selects: []string{
			"subtotal",
			"shipping_total",
			"discount_total",
			"tax_total",
			"gift_card_total",
			"total",
		},
		Relations: []string{
			"discounts",
			"discounts.rule",
			"region",
			"fulfillments",
			"shipping_address",
			"billing_address",
			"shipping_methods",
			"shipping_methods.shipping_option",
			"items.adjustments",
			"items.variant.product.profiles",
			"payments",
		},
	})
	if err != nil {
		return nil, err
	}

	if order.Status == "CANCELED" {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"A canceled order cannot be fulfilled",
			"500",
			nil,
		)
	}

	if len(order.ShippingMethods) == 0 {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Cannot fulfill an order that lacks shipping methods",
			"500",
			nil,
		)
	}

	fulfillmentOrder := &types.CreateFulfillmentOrder{}

	copier.CopyWithOption(&fulfillmentOrder, order, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	fulfillments, err := s.r.FulfillmentService().SetContext(s.ctx).CreateFulfillment(
		fulfillmentOrder,
		itemsToFulfill,
		models.Fulfillment{
			Model: core.Model{
				Metadata: metadata,
			},
			NoNotification: no_notification,
			OrderId:        uuid.NullUUID{UUID: order.Id},
			LocationId:     location_id,
		},
	)
	if err != nil {
		return nil, err
	}

	var successfullyFulfilled []models.FulfillmentItem
	for _, f := range fulfillments {
		successfullyFulfilled = append(successfullyFulfilled, f.Items...)
	}

	order.FulfillmentStatus = models.FulfillmentStatusFulfilled

	for _, item := range order.Items {
		fulfillmentItem, ok := lo.Find(successfullyFulfilled, func(i models.FulfillmentItem) bool {
			return i.ItemId.UUID == item.Id
		})
		if ok {
			fulfilledQuantity := item.FulfilledQuantity + fulfillmentItem.Quantity
			_, err := s.r.LineItemService().SetContext(s.ctx).Update(item.Id, nil, &models.LineItem{
				FulfilledQuantity: fulfilledQuantity,
			}, sql.Options{})
			if err != nil {
				return nil, err
			}
			if item.Quantity != fulfilledQuantity {
				order.FulfillmentStatus = models.FulfillmentStatusPartiallyFulfilled
			}
		} else {
			if item.Quantity != item.FulfilledQuantity {
				order.FulfillmentStatus = models.FulfillmentStatusPartiallyFulfilled
			}
		}
	}

	order.Fulfillments = append(order.Fulfillments, fulfillments...)

	if err := s.r.OrderRepository().Save(s.ctx, order); err != nil {
		return nil, err
	}

	// evaluatedNoNotification := order.NoNotification
	// if no_notification {
	// 	evaluatedNoNotification = no_notification
	// }

	// eventsToEmit := []map[string]interface{}{}
	// for _, fulfillment := range fulfillments {
	// 	event := map[string]interface{}{
	// 		"eventName": OrderService.Events.FULFILLMENT_CREATED,
	// 		"data": map[string]interface{}{
	// 			"id":              id,
	// 			"fulfillment_id":  fulfillment.ID,
	// 			"no_notification": evaluatedNoNotification,
	// 		},
	// 	}
	// 	eventsToEmit = append(eventsToEmit, event)
	// }

	// err = eventBus.emit(eventsToEmit)
	// if err != nil {
	// 	return nil, err
	// }

	return order, nil
}

func (s *OrderService) CancelFulfillment(fulfillmentId uuid.UUID) (*models.Order, *utils.ApplictaionError) {
	canceled, err := s.r.FulfillmentService().SetContext(s.ctx).CancelFulfillment(fulfillmentId, nil)
	if err != nil {
		return nil, err
	}

	if canceled.OrderId.UUID == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Fufillment not related to an order",
			"500",
			nil,
		)
	}

	order, err := s.RetrieveById(canceled.OrderId.UUID, sql.Options{})
	if err != nil {
		return nil, err
	}

	order.FulfillmentStatus = models.FulfillmentStatusCanceled

	if err := s.r.OrderRepository().Save(s.ctx, order); err != nil {
		return nil, err
	}

	// err = eventBus.emit(OrderService.Events.FULFILLMENT_CANCELED, map[string]interface{}{
	// 	"id":              order.ID,
	// 	"fulfillment_id":  canceled.ID,
	// 	"no_notification": canceled.NoNotification,
	// })
	// if err != nil {
	// 	return nil, err
	// }

	return order, nil
}

func (s *OrderService) GetFulfillmentItems(order models.Order, items []types.FulFillmentItemType, transformer func(item models.LineItem, quantity int) interface{}) ([]models.LineItem, *utils.ApplictaionError) {
	lineItems := []models.LineItem{}
	for _, item := range items {
		lineItem, ok := lo.Find(order.Items, func(i models.LineItem) bool {
			return i.Id == item.ItemId
		})
		if ok {
			result := transformer(lineItem, item.Quantity)
			if result != nil {
				lineItems = append(lineItems, result.(models.LineItem))
			}
		}
	}
	return lineItems, nil
}

func (s *OrderService) Archive(id uuid.UUID) (*models.Order, *utils.ApplictaionError) {
	order, err := s.RetrieveById(id, sql.Options{})
	if err != nil {
		return nil, err
	}

	if order.Status != models.OrderStatusCompleted && order.Status != models.OrderStatusRefunded {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Can't archive an unprocessed order",
			"500",
			nil,
		)
	}

	order.Status = models.OrderStatusArchived

	if err := s.r.OrderRepository().Save(s.ctx, order); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) CreateRefund(id uuid.UUID, refundAmount float64, reason string, note *string, noNotification *bool) (*models.Order, *utils.ApplictaionError) {
	order, err := s.RetrieveById(id, sql.Options{
		Selects: []string{
			"refundable_amount",
			"total",
			"refunded_total",
		},
		Relations: []string{
			"payments",
		},
	})
	if err != nil {
		return nil, err
	}

	if order.Status == models.OrderStatusCanceled {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"A canceled order cannot be refunded",
			"500",
			nil,
		)
	}

	if refundAmount > order.RefundableAmount {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Cannot refund more than the original order amount",
			"500",
			nil,
		)
	}

	_, err = s.r.PaymentProviderService().SetContext(s.ctx).RefundPayments(order.Payments, refundAmount, reason, note)
	if err != nil {
		return nil, err
	}

	result, err := s.RetrieveByIdWithTotals(id, sql.Options{
		Relations: []string{
			"payments",
		},
	}, types.TotalsContext{})
	if err != nil {
		return nil, err
	}

	if result.RefundedTotal > 0 && result.RefundableAmount > 0 {
		result.PaymentStatus = models.PaymentStatusPartiallyRefunded
		if err := s.r.OrderRepository().Save(s.ctx, result); err != nil {
			return nil, err
		}
	}

	if result.PaidTotal > 0 && result.RefundedTotal == result.PaidTotal {
		result.PaymentStatus = models.PaymentStatusRefunded
		if err := s.r.OrderRepository().Save(s.ctx, result); err != nil {
			return nil, err
		}
	}

	// evaluatedNoNotification := noNotification or order.NoNotification

	// err = eventBus.emit(OrderService.Events.REFUND_CREATED, map[string]interface{}{
	// 	"id":              result.ID,
	// 	"refund_id":       refund.ID,
	// 	"no_notification": evaluatedNoNotification,
	// })
	// if err != nil {
	// 	return nil, err
	// }

	return result, nil
}

func (s *OrderService) decorateTotalsLegacy(order *models.Order, totalsFields []string) (*models.Order, *utils.ApplictaionError) {
	if slices.Contains(totalsFields, "subtotal") || slices.Contains(totalsFields, "total") {
		calculationContext, err := s.r.TotalsService().GetCalculationContext(nil, order, CalculationContextOptions{ExcludeShipping: true})
		if err != nil {
			return nil, err
		}

		for i, item := range order.Items {
			itemTotals, err := s.r.TotalsService().GetLineItemTotals(item, nil, order, &LineItemTotalsOptions{
				IncludeTax:         true,
				CalculationContext: calculationContext,
			})
			if err != nil {
				return nil, err
			}
			order.Items[i] = *itemTotals
		}
	}

	for _, totalField := range totalsFields {
		switch totalField {
		case "shipping_total":
			total, err := s.r.TotalsService().GetShippingTotal(nil, order)
			if err != nil {
				return nil, err
			}
			order.ShippingTotal = total
		case "gift_card_total":
			giftCardBreakdown, err := s.r.TotalsService().GetGiftCardTotal(nil, order, nil)
			if err != nil {
				return nil, err
			}
			order.GiftCardTotal = giftCardBreakdown.Total
			order.GiftCardTaxTotal = giftCardBreakdown.TaxTotal
		case "discount_total":
			total, err := s.r.TotalsService().GetDiscountTotal(nil, order)
			if err != nil {
				return nil, err
			}
			order.DiscountTotal = total
		case "tax_total":
			total, err := s.r.TotalsService().GetTaxTotal(nil, order, false)
			if err != nil {
				return nil, err
			}
			order.TaxTotal = total
		case "subtotal":
			total, err := s.r.TotalsService().GetSubtotal(nil, order, types.SubtotalOptions{})
			if err != nil {
				return nil, err
			}
			order.Subtotal = total
		case "total":
			total, err := s.r.TotalsService().GetTotal(nil, order, GetTotalsOptions{})
			if err != nil {
				return nil, err
			}
			order.Total = total
		case "refunded_total":
			order.RefundedTotal = s.r.TotalsService().GetRefundedTotal(order)
		case "paid_total":
			order.PaidTotal = s.r.TotalsService().GetPaidTotal(order)
		case "refundable_amount":
			paidTotal := s.r.TotalsService().GetPaidTotal(order)
			refundedTotal := s.r.TotalsService().GetRefundedTotal(order)
			order.RefundableAmount = paidTotal - refundedTotal
		case "items.refundable":
			items := []models.LineItem{}
			for _, item := range order.Items {
				refundable, err := s.r.TotalsService().GetLineItemRefund(order, item)
				if err != nil {
					return nil, err
				}
				item.Refundable = refundable
				items = append(items, item)
			}
			order.Items = items
		case "swaps.additional_items.refundable":
			for _, swap := range order.Swaps {
				items := []models.LineItem{}
				for _, item := range swap.AdditionalItems {
					refundable, err := s.r.TotalsService().GetLineItemRefund(order, item)
					if err != nil {
						return nil, err
					}
					item.Refundable = refundable
					items = append(items, item)
				}
				swap.AdditionalItems = items
			}
		case "claims.additional_items.refundable":
			for _, claim := range order.Claims {
				items := []models.LineItem{}
				for _, item := range claim.AdditionalItems {
					refundable, err := s.r.TotalsService().GetLineItemRefund(order, item)
					if err != nil {
						return nil, err
					}
					item.Refundable = refundable
					items = append(items, item)
				}
				claim.AdditionalItems = items
			}
		}
	}

	return order, nil
}

func (s *OrderService) decorateTotals(order *models.Order, totalsFields []string, context types.TotalsContext) (*models.Order, *utils.ApplictaionError) {
	if len(totalsFields) > 0 {
		return s.decorateTotalsLegacy(order, totalsFields)
	}

	calculationContext, err := s.r.TotalsService().GetCalculationContext(nil, order, CalculationContextOptions{})
	if err != nil {
		return nil, err
	}

	var returnableItems []models.LineItem
	if context.ReturnableItems {
		returnableItems = []models.LineItem{}
	} else {
		returnableItems = nil
	}

	isReturnableItem := func(item models.LineItem) bool {
		return context.ReturnableItems && item.ReturnedQuantity < item.ShippedQuantity
	}

	var items []models.LineItem
	if context.ReturnableItems {
		for _, swap := range order.Swaps {
			items = append(items, swap.AdditionalItems...)
		}
		for _, claim := range order.Claims {
			items = append(items, claim.AdditionalItems...)
		}
	}
	allItems := append(order.Items, items...)

	itemsTotals, err := s.r.NewTotalsService().SetContext(s.ctx).GetLineItemTotals(allItems, true, calculationContext, &order.TaxRate)
	if err != nil {
		return nil, err
	}

	shippingTotals, err := s.r.NewTotalsService().SetContext(s.ctx).GetShippingMethodTotals(order.ShippingMethods, true, order.Discounts, &order.TaxRate, calculationContext)
	if err != nil {
		return nil, err
	}

	order.Subtotal = 0
	order.DiscountTotal = 0
	order.ShippingTotal = 0
	order.RefundedTotal = lo.Reduce(order.Refunds, func(agg float64, item models.Refund, index int) float64 {
		return agg + item.Amount
	}, 0)
	order.PaidTotal = lo.Reduce(order.Payments, func(agg float64, item models.Payment, index int) float64 {
		return agg + item.Amount
	}, 0)
	order.RefundableAmount = order.PaidTotal - order.RefundedTotal
	itemTaxTotal := 0.0
	shippingTaxTotal := 0.0

	for i, item := range order.Items {
		item.Quantity = item.Quantity - item.ReturnedQuantity
		refundable, err := s.r.NewTotalsService().SetContext(s.ctx).GetLineItemRefund(item, calculationContext, &order.TaxRate)
		if err != nil {
			return nil, err
		}
		copier.CopyWithOption(&item, itemsTotals[item.Id], copier.Option{IgnoreEmpty: true, DeepCopy: true})
		item.Refundable = refundable
		order.Subtotal += item.Subtotal
		order.DiscountTotal += item.RawDiscountTotal
		itemTaxTotal += item.TaxTotal
		if isReturnableItem(item) {
			returnableItems = append(returnableItems, item)
		}
		order.Items[i] = item
	}

	for i, shippingMethod := range order.ShippingMethods {
		copier.CopyWithOption(&shippingMethod, shippingTotals[shippingMethod.Id], copier.Option{IgnoreEmpty: true, DeepCopy: true})
		order.ShippingTotal += shippingMethod.Subtotal
		shippingTaxTotal += shippingMethod.TaxTotal
		order.ShippingMethods[i] = shippingMethod
	}

	order.ItemTaxTotal = itemTaxTotal
	order.ShippingTaxTotal = shippingTaxTotal
	order.TaxTotal = itemTaxTotal + shippingTaxTotal

	giftCardableAmount := s.r.NewTotalsService().SetContext(s.ctx).GetGiftCardableAmount(
		order.Region.GiftCardsTaxable,
		order.Subtotal,
		order.ShippingTotal,
		order.DiscountTotal,
		order.TaxTotal,
	)
	giftCardTotal, err := s.r.NewTotalsService().SetContext(s.ctx).GetGiftCardTotals(
		giftCardableAmount,
		order.GiftCardTransactions,
		order.Region,
		order.GiftCards,
	)
	if err != nil {
		return nil, err
	}
	order.GiftCardTotal = giftCardTotal.Total
	order.GiftCardTaxTotal = giftCardTotal.TaxTotal
	order.TaxTotal -= order.GiftCardTaxTotal

	for _, swap := range order.Swaps {
		for i, item := range swap.AdditionalItems {
			item.Quantity = item.Quantity - item.ReturnedQuantity
			refundable, err := s.r.NewTotalsService().SetContext(s.ctx).GetLineItemRefund(
				item,
				calculationContext,
				&order.TaxRate,
			)
			if err != nil {
				return nil, err
			}
			item.Refundable = refundable
			copier.CopyWithOption(&item, itemsTotals[item.Id], copier.Option{IgnoreEmpty: true, DeepCopy: true})
			if isReturnableItem(item) {
				returnableItems = append(returnableItems, item)
			}
			swap.AdditionalItems[i] = item
		}
	}

	for _, claim := range order.Claims {
		for i, item := range claim.AdditionalItems {
			item.Quantity = item.Quantity - item.ReturnedQuantity
			refundable, err := s.r.NewTotalsService().SetContext(s.ctx).GetLineItemRefund(
				item,
				calculationContext,
				&order.TaxRate,
			)
			if err != nil {
				return nil, err
			}
			item.Refundable = refundable
			copier.CopyWithOption(&item, itemsTotals[item.Id], copier.Option{IgnoreEmpty: true, DeepCopy: true})
			if isReturnableItem(item) {
				returnableItems = append(returnableItems, item)
			}
			claim.AdditionalItems[i] = item
		}
	}

	order.RawDiscountTotal = order.DiscountTotal
	order.DiscountTotal = math.Round(order.DiscountTotal)
	order.Total = order.Subtotal + order.ShippingTotal + order.TaxTotal - (order.GiftCardTotal + order.DiscountTotal)
	order.ReturnableItems = returnableItems

	return order, nil
}

func (s *OrderService) RegisterReturnReceived(id uuid.UUID, receivedReturn *models.Return, customRefundAmount *float64) (*models.Order, *utils.ApplictaionError) {
	order, err := s.RetrieveById(id, sql.Options{
		Selects:   []string{"total", "refunded_total", "refundable_amount"},
		Relations: []string{"items", "returns", "payments"},
	})
	if err != nil {
		return nil, err
	}
	if order.Status == "canceled" {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"A canceled order cannot be registered as received",
			"500",
			nil,
		)
	}
	if receivedReturn == nil || receivedReturn.OrderId.UUID != id {
		return nil, utils.NewApplictaionError(
			utils.NOT_FOUND,
			"Received return does not exist",
			"500",
			nil,
		)
	}
	refundAmount := receivedReturn.RefundAmount
	if customRefundAmount != nil {
		refundAmount = *customRefundAmount
	}

	if refundAmount > order.RefundableAmount {
		order.FulfillmentStatus = models.FulfillmentStatusRequiresAction
		if err := s.r.OrderRepository().Save(s.ctx, order); err != nil {
			return nil, err
		}
		// err = s.eventBus_.Emit(OrderServiceEvents.RETURN_ACTION_REQUIRED, map[string]interface{}{
		// 	"id":              result.ID,
		// 	"return_id":       receivedReturn.ID,
		// 	"no_notification": receivedReturn.NoNotification,
		// })
		// if err != nil {
		// 	return nil, err
		// }
		return order, nil
	}
	isFullReturn := true
	for _, i := range order.Items {
		if i.ReturnedQuantity != i.Quantity {
			isFullReturn = false
			break
		}
	}
	if refundAmount > 0 {
		refund, err := s.r.PaymentProviderService().SetContext(s.ctx).RefundPayments(order.Payments, refundAmount, "return", nil)
		if err != nil {
			return nil, err
		}
		order.Refunds = append(order.Refunds, *refund)
	}
	if isFullReturn {
		order.FulfillmentStatus = models.FulfillmentStatusReturned
	} else {
		order.FulfillmentStatus = models.FulfillmentStatusPartiallyReturned
	}
	if err := s.r.OrderRepository().Save(s.ctx, order); err != nil {
		return nil, err
	}
	// err = s.eventBus_.Emit(OrderServiceEvents.ITEMS_RETURNED, map[string]interface{}{
	// 	"id":              order.ID,
	// 	"return_id":       receivedReturn.ID,
	// 	"no_notification": receivedReturn.NoNotification,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	return order, nil
}

func (s *OrderService) transformQueryForTotals(config sql.Options) ([]string, []string, []string) {
	selects, relations := config.Selects, config.Relations
	if selects == nil {
		return selects, relations, []string{}
	}
	totalFields := []string{
		"subtotal",
		"tax_total",
		"shipping_total",
		"discount_total",
		"gift_card_total",
		"total",
		"paid_total",
		"refunded_total",
		"refundable_amount",
		"items.refundable",
		"swaps.additional_items.refundable",
		"claims.additional_items.refundable",
	}
	totalsToSelect := []string{}
	for _, v := range selects {
		if slices.Contains(totalFields, v) {
			totalsToSelect = append(totalsToSelect, v)
		}
	}
	if len(totalsToSelect) > 0 {
		relationSet := []string{
			"items.tax_lines",
			"items.adjustments",
			"items.variant.product.profiles",
			"swaps",
			"swaps.additional_items",
			"swaps.additional_items.tax_lines",
			"swaps.additional_items.adjustments",
			"claims",
			"claims.additional_items",
			"claims.additional_items.tax_lines",
			"claims.additional_items.adjustments",
			"discounts",
			"discounts.rule",
			"gift_cards",
			"gift_card_transactions",
			"gift_card_transactions.gift_card",
			"refunds",
			"shipping_methods",
			"shipping_methods.tax_lines",
			"region",
		}

		for _, r := range relations {
			if !slices.Contains(relationSet, r) {
				relationSet = append(relationSet, r)
			}
		}

		relations = []string{}
		relations = append(relations, relationSet...)

		selects = []string{}
		for _, v := range selects {
			if !slices.Contains(totalFields, v) {
				selects = append(selects, v)
			}
		}
	}
	toSelect := []string{}
	toSelect = append(toSelect, selects...)

	if len(toSelect) > 0 && !slices.Contains(toSelect, "tax_rate") {
		toSelect = append(toSelect, "tax_rate")
	}
	return toSelect, relations, totalsToSelect
}

func (s *OrderService) GetTotalsRelations(config sql.Options) []string {
	relationSet := []string{
		"items",
		"items.tax_lines",
		"items.adjustments",
		"items.variant",
		"swaps",
		"swaps.additional_items",
		"swaps.additional_items.tax_lines",
		"swaps.additional_items.adjustments",
		"claims",
		"claims.additional_items",
		"claims.additional_items.tax_lines",
		"claims.additional_items.adjustments",
		"discounts",
		"discounts.rule",
		"gift_cards",
		"gift_card_transactions",
		"refunds",
		"shipping_methods",
		"shipping_methods.tax_lines",
		"region",
		"payments",
	}

	for _, relation := range config.Relations {
		if slices.Contains(relationSet, relation) {
			relationSet = append(relationSet, relation)
		}
	}

	return relationSet
}

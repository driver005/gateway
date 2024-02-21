package services

import (
	"context"
	"reflect"
	"strings"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"golang.org/x/exp/slices"
)

type SwapService struct {
	ctx context.Context
	r   Registry
}

func NewSwapService(
	r Registry,
) *SwapService {
	return &SwapService{
		context.Background(),
		r,
	}
}

func (s *SwapService) SetContext(context context.Context) *SwapService {
	s.ctx = context
	return s
}

func (s *SwapService) transformQueryForCart(config *sql.Options) map[string]*sql.Options {
	selects := config.Selects
	relations := config.Relations
	var cartSelects []string
	var cartRelations []string

	if relations != nil && slices.Contains(relations, "cart") {
		swapRelations := make([]string, 0)
		cartRels := make([]string, 0)

		for _, next := range relations {
			if next == "cart" {
				continue
			}
			if strings.HasPrefix(next, "cart.") {
				rel := strings.Split(next, ".")[1:]
				cartRels = append(cartRels, strings.Join(rel, "."))
			} else {
				swapRelations = append(swapRelations, next)
			}
		}

		relations = swapRelations
		cartRelations = cartRels

		foundCartId := false
		if selects != nil {
			swapSelects := make([]string, 0)
			cartSels := make([]string, 0)

			for _, next := range selects {
				if strings.HasPrefix(next, "cart.") {
					rel := strings.Split(next, ".")[1:]
					cartSels = append(cartSels, strings.Join(rel, "."))
				} else {
					if next == "cart_id" {
						foundCartId = true
					}
					swapSelects = append(swapSelects, next)
				}
			}

			if foundCartId {
				selects = swapSelects
			} else {
				selects = append(swapSelects, "cart_id")
			}
			cartSelects = cartSels
		}
	}

	res := make(map[string]*sql.Options)

	res["swap"] = &sql.Options{
		Selects:   selects,
		Relations: relations,
	}
	res["cart"] = &sql.Options{
		Selects:   cartSelects,
		Relations: cartRelations,
	}

	return res
}

func (s *SwapService) Retrieve(swapId uuid.UUID, config *sql.Options) (*models.Swap, *utils.ApplictaionError) {
	if swapId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			`"swapId" must be defined`,
			nil,
		)
	}

	var swap *models.Swap = &models.Swap{}
	newConfig := s.transformQueryForCart(config)
	query := sql.BuildQuery(models.Swap{Model: core.Model{Id: swapId}}, newConfig["swap"])
	if err := s.r.SwapRepository().FindOne(s.ctx, swap, query); err != nil {
		return nil, err
	}

	if newConfig["cart"].Selects != nil || newConfig["cart"].Relations != nil {
		cart, err := s.r.CartService().SetContext(s.ctx).Retrieve(swap.CartId.UUID, &sql.Options{
			Selects:   newConfig["cart"].Selects,
			Relations: newConfig["cart"].Relations,
		}, TotalsConfig{})
		if err != nil {
			return nil, err
		}
		swap.Cart = cart
	}

	return swap, nil
}

func (s *SwapService) RetrieveByCartId(cartId uuid.UUID, relations []string) (*models.Swap, *utils.ApplictaionError) {
	var swap *models.Swap = &models.Swap{}
	query := sql.BuildQuery(models.Swap{CartId: uuid.NullUUID{UUID: cartId}}, &sql.Options{
		Relations: relations,
	})
	if err := s.r.SwapRepository().FindOne(s.ctx, swap, query); err != nil {
		return nil, err
	}

	return swap, nil
}

func (s *SwapService) List(selector *types.FilterableSwap, config *sql.Options) ([]models.Swap, *utils.ApplictaionError) {
	res, _, err := s.ListAndCount(selector, config)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SwapService) ListAndCount(selector *types.FilterableSwap, config *sql.Options) ([]models.Swap, *int64, *utils.ApplictaionError) {
	if reflect.DeepEqual(config, &sql.Options{}) {
		config.Skip = 0
		config.Take = 50
		config.Order = "created_at DESC"
	}

	var res []models.Swap

	query := sql.BuildQuery(selector, config)

	count, err := s.r.SwapRepository().FindAndCount(s.ctx, &res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

func (s *SwapService) Create(order *models.Order, returnItems []types.OrderReturnItem, additionalItems []types.CreateClaimItemAdditionalItemInput, returnShipping *types.CreateClaimReturnShippingInput, custom map[string]interface{}) (*models.Swap, *utils.ApplictaionError) {
	noNotification, _ := custom["no_notification"].(bool)
	idempotencyKey, _ := custom["idempotency_key"].(string)
	allowBackorder, _ := custom["allow_backorder"].(bool)
	locationId, _ := custom["location_id"].(uuid.UUID)

	if order.PaymentStatus != models.PaymentStatusCaptured {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			"Cannot swap an order that has not been captured",
			nil,
		)
	}
	if order.FulfillmentStatus == models.FulfillmentStatusNotFulfilled {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			"Cannot swap an order that has not been fulfilled",
			nil,
		)
	}

	areReturnItemsValid, _ := s.AreReturnItemsValid(returnItems)
	if !areReturnItemsValid {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			"Cannot Create a swap on a canceled item.",
			nil,
		)
	}

	var newItems []models.LineItem
	for _, item := range additionalItems {
		if item.VariantId == uuid.Nil {
			return nil, utils.NewApplictaionError(
				utils.CONFLICT,
				"You must include a variant when creating additional items on a swap",
				"500",
				nil,
			)
		}
		newItem, _ := s.r.LineItemService().SetContext(s.ctx).Generate(
			item.VariantId,
			nil,
			order.RegionId.UUID,
			item.Quantity,
			types.GenerateLineItemContext{
				RegionId: order.RegionId.UUID,
				Cart:     order.Cart,
			},
		)

		newItems = append(newItems, newItem...)
	}

	evaluatedNoNotification := noNotification
	if !evaluatedNoNotification {
		evaluatedNoNotification = order.NoNotification
	}

	res := &models.Swap{
		FulfillmentStatus: models.SwapFulfillmentNotFulfilled,
		PaymentStatus:     models.SwapPaymentNotPaid,
		OrderId:           uuid.NullUUID{UUID: order.Id},
		AdditionalItems:   newItems,
		NoNotification:    evaluatedNoNotification,
		IdempotencyKey:    idempotencyKey,
		AllowBackorder:    allowBackorder,
	}
	if err := s.r.SwapRepository().Save(s.ctx, res); err != nil {
		return nil, err
	}

	data := &types.CreateReturnInput{
		SwapId:         res.Id,
		OrderId:        order.Id,
		Items:          returnItems,
		ShippingMethod: returnShipping,
		NoNotification: evaluatedNoNotification,
		LocationId:     locationId,
	}
	_, err := s.r.ReturnService().Create(data)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SwapService) ProcessDifference(swapId uuid.UUID) (*models.Swap, *utils.ApplictaionError) {
	swap, err := s.Retrieve(swapId, &sql.Options{
		Relations: []string{"payment", "order", "order.payments"},
	})
	if err != nil {
		return nil, err
	}

	if swap.CanceledAt != nil {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			"Canceled swap cannot be processed",
			nil,
		)
	}
	if swap.ConfirmedAt == nil {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			"Cannot process a swap that hasn't been confirmed by the customer",
			nil,
		)
	}

	if swap.DifferenceDue < 0 {
		if swap.PaymentStatus == models.SwapPaymentDifferenceRefunded {
			return swap, nil
		}
		_, err := s.r.PaymentProviderService().SetContext(s.ctx).RefundPayments(swap.Order.Payments, -1*swap.DifferenceDue, "swap", nil)
		if err != nil {
			swap.PaymentStatus = models.SwapPaymentRequiresAction
			if err := s.r.SwapRepository().Save(s.ctx, swap); err != nil {
				return nil, err
			}
			return swap, nil
		}
		swap.PaymentStatus = models.SwapPaymentDifferenceRefunded
		if err := s.r.SwapRepository().Save(s.ctx, swap); err != nil {
			return nil, err
		}
		return swap, nil
	} else if swap.DifferenceDue == 0 {
		if swap.PaymentStatus == models.SwapPaymentDifferenceRefunded {
			return swap, nil
		}
		swap.PaymentStatus = models.SwapPaymentDifferenceRefunded
		if err := s.r.SwapRepository().Save(s.ctx, swap); err != nil {
			return nil, err
		}
		return swap, nil
	}

	if swap.PaymentStatus == models.SwapPaymentCaptured {
		return swap, nil
	}
	_, err = s.r.PaymentProviderService().SetContext(s.ctx).CapturePayment(swap.Payment)
	if err != nil {
		swap.PaymentStatus = models.SwapPaymentRequiresAction
		if err := s.r.SwapRepository().Save(s.ctx, swap); err != nil {
			return nil, err
		}
		return swap, nil
	}
	swap.PaymentStatus = models.SwapPaymentCaptured
	if err := s.r.SwapRepository().Save(s.ctx, swap); err != nil {
		return nil, err
	}
	return swap, nil
}

func (s *SwapService) Update(swapId uuid.UUID, Update *models.Swap) (*models.Swap, *utils.ApplictaionError) {
	swap, err := s.Retrieve(swapId, &sql.Options{})
	if err != nil {
		return nil, err
	}

	// if Update.ShippingMethods != nil {
	// TODO: Check this - calling method that doesn't exist
	// also it seems that Update swap isn't call anywhere
	// await this.updateShippingAddress_(swap, Update.shipping_address)
	// }

	Update.Id = swap.Id

	if err := s.r.SwapRepository().Upsert(s.ctx, Update); err != nil {
		return nil, err
	}
	return Update, nil
}

func (s *SwapService) CreateCart(swapId uuid.UUID, customShippingOptions []types.CreateCustomShippingOptionInput, context map[string]interface{}) (*models.Swap, *utils.ApplictaionError) {
	swap, err := s.Retrieve(swapId, &sql.Options{
		Relations: []string{
			"order.items.variant.product.profiles",
			"order.swaps",
			"order.swaps.additional_items",
			"order.discounts",
			"order.discounts.rule",
			"order.claims",
			"order.claims.additional_items",
			"additional_items",
			"additional_items.variant",
			"return_order",
			"return_order.items",
			"return_order.shipping_method",
			"return_order.shipping_method.shipping_option",
			"return_order.shipping_method.tax_lines",
		},
	})
	if err != nil {
		return nil, err
	}

	if swap.CanceledAt != nil {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			"Canceled swap cannot be used to Create a cart",
			nil,
		)
	}
	if swap.CartId.UUID != uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			"A cart has already been created for the swap",
			nil,
		)
	}

	order := swap.Order
	var discounts []types.Discount
	if order.Discounts != nil {
		for _, discount := range order.Discounts {
			if discount.Rule.Type != "free_shipping" {
				discounts = append(discounts, types.Discount{
					Code: discount.Code,
				})
			}
		}
	}

	salesChannelId, _ := context["sales_channel_id"].(uuid.UUID)
	cart, err := s.r.CartService().SetContext(s.ctx).Create(&types.CartCreateProps{
		Metadata: core.JSONB{
			"swap_id":         swap.Id,
			"parent_order_id": order.Id,
		},
		Discounts:         discounts,
		Email:             order.Email,
		BillingAddressId:  order.BillingAddressId.UUID,
		ShippingAddressId: order.ShippingAddressId.UUID,
		RegionId:          order.RegionId.UUID,
		CustomerId:        order.CustomerId.UUID,
		SalesChannelId:    salesChannelId,
		Type:              models.CartSwap,
	})
	if err != nil {
		return nil, err
	}

	_, err = s.r.CustomShippingOptionService().SetContext(s.ctx).Create(customShippingOptions)
	if err != nil {
		return nil, err
	}

	for _, item := range swap.AdditionalItems {
		_, err := s.r.LineItemService().SetContext(s.ctx).Update(item.Id, nil, &models.LineItem{
			CartId: uuid.NullUUID{UUID: cart.Id},
		}, &sql.Options{})
		if err != nil {
			return nil, err
		}
	}

	c, err := s.r.CartService().SetContext(s.ctx).Retrieve(cart.Id, &sql.Options{
		Relations: []string{
			"items",
			"items.variant",
			"region",
			"discounts",
			"discounts.rule",
		},
	}, TotalsConfig{})
	if err != nil {
		return nil, err
	}
	cart = c

	for _, item := range cart.Items {
		s.r.LineItemAdjustmentService().SetContext(s.ctx).CreateAdjustmentForLineItem(cart, &item)
	}

	if swap.ReturnOrder != nil && swap.ReturnOrder.ShippingMethod != nil {
		s.r.LineItemService().SetContext(s.ctx).Create([]models.LineItem{
			{
				CartId:         uuid.NullUUID{UUID: cart.Id},
				Title:          "Return shipping",
				Quantity:       1,
				HasShipping:    true,
				AllowDiscounts: false,
				UnitPrice:      swap.ReturnOrder.ShippingMethod.Price,
				IsReturn:       true,
				TaxLines: func() []models.LineItemTaxLine {
					var taxLines []models.LineItemTaxLine
					for _, tl := range swap.ReturnOrder.ShippingMethod.TaxLines {
						item, err := s.r.LineItemService().SetContext(s.ctx).CreateTaxLine(&models.LineItemTaxLine{
							Model: core.Model{
								Metadata: tl.Metadata,
							},
							Name: tl.Name,
							Code: tl.Code,
							Rate: tl.Rate,
						})

						if err != nil {
							return nil
						}

						taxLines = append(taxLines, *item)
					}
					return taxLines
				}(),
			},
		})
	}

	s.r.LineItemService().SetContext(s.ctx).CreateReturnLines(swap.ReturnOrder.Id, cart.Id)

	swap.CartId = uuid.NullUUID{UUID: cart.Id}
	if err := s.r.SwapRepository().Save(s.ctx, swap); err != nil {
		return nil, err
	}

	return swap, nil
}

func (s *SwapService) RegisterCartCompletion(swapId uuid.UUID) (*models.Swap, *utils.ApplictaionError) {
	swap, err := s.Retrieve(swapId, &sql.Options{
		Selects: []string{
			"id",
			"order_id",
			"no_notification",
			"allow_backorder",
			"canceled_at",
			"confirmed_at",
			"cart_id",
		},
	})
	if err != nil {
		return nil, err
	}
	if swap.ConfirmedAt != nil {
		return swap, nil
	}
	if swap.CanceledAt != nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Cart related to canceled swap cannot be completed",
			nil,
		)
	}
	cart, err := s.r.CartService().SetContext(s.ctx).RetrieveWithTotals(swap.CartId.UUID, &sql.Options{
		Relations: []string{"payment"},
	}, TotalsConfig{})
	if err != nil {
		return nil, err
	}
	payment := cart.Payment
	items := cart.Items
	total := cart.Total
	if total > 0 {
		if payment == nil {
			return nil, utils.NewApplictaionError(
				utils.INVALID_ARGUMENT,
				"Cart does not contain a payment",
				"500",
				nil,
			)
		}
		paymentStatus, err := s.r.PaymentProviderService().SetContext(s.ctx).GetStatus(payment)
		if err != nil {
			return nil, err
		}
		if *paymentStatus != models.PaymentSessionStatusAuthorized && *paymentStatus != models.PaymentSessionStatusSuccess {
			return nil, utils.NewApplictaionError(
				utils.INVALID_ARGUMENT,
				"Payment method is not authorized",
				"500",
				nil,
			)
		}
		_, err = s.r.PaymentProviderService().SetContext(s.ctx).UpdatePayment(payment.Id, &types.UpdatePaymentInput{
			SwapId:  swapId,
			OrderId: swap.OrderId.UUID,
		})
		if err != nil {
			return nil, err
		}
		for _, item := range items {
			if item.VariantId.UUID != uuid.Nil {
				_, err = s.r.ProductVariantInventoryService().SetContext(s.ctx).ReserveQuantity(item.VariantId.UUID, item.Quantity, ReserveQuantityContext{
					LineItemId:     item.Id,
					SalesChannelId: cart.SalesChannelId.UUID,
				})
				if err != nil {
					return nil, err
				}
			}
		}
	}
	swap.DifferenceDue = total
	swap.ShippingAddressId = cart.ShippingAddressId
	swap.ShippingMethods = make([]models.ShippingMethod, len(cart.ShippingMethods))
	for i, method := range cart.ShippingMethods {
		method.TaxLines = nil
		swap.ShippingMethods[i] = method
	}
	now := time.Now()
	swap.ConfirmedAt = &now
	if total == 0 {
		swap.PaymentStatus = models.SwapPaymentConfirmed
	} else {
		swap.PaymentStatus = models.SwapPaymentAwaiting
	}
	if err := s.r.SwapRepository().Save(s.ctx, swap); err != nil {
		return nil, err
	}

	for _, method := range cart.ShippingMethods {
		m, err := s.r.ShippingOptionService().SetContext(s.ctx).UpdateShippingMethod(method.Id, &types.ShippingMethodUpdate{
			SwapId: swap.Id,
		})
		if err != nil {
			return nil, err
		}
		method = *m
	}
	// err = s.eventBus_.Emit(SwapServiceEventsPaymentCompleted, &PaymentCompletedEvent{
	// 	ID:             swap.ID,
	// 	NoNotification: swap.NoNotification,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	_, err = s.r.CartService().SetContext(s.ctx).Update(cart.Id, nil, &types.CartUpdateProps{
		CompletedAt: &now,
	})
	if err != nil {
		return nil, err
	}
	return swap, nil
}

func (s *SwapService) Cancel(swapId uuid.UUID) (*models.Swap, *utils.ApplictaionError) {
	swap, err := s.Retrieve(swapId, &sql.Options{
		Relations: []string{"payment", "fulfillments", "return_order"},
	})
	if err != nil {
		return nil, err
	}
	if swap.PaymentStatus == models.SwapPaymentDifferenceRefunded || swap.PaymentStatus == models.SwapPaymentPartiallyRefunded || swap.PaymentStatus == models.SwapPaymentRefunded {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Swap with a refund cannot be canceled",
			nil,
		)
	}
	if swap.Fulfillments != nil {
		for _, f := range swap.Fulfillments {
			if f.CanceledAt == nil {
				return nil, utils.NewApplictaionError(
					utils.NOT_ALLOWED,
					"All fulfillments must be canceled before the swap can be canceled",
					"500",
					nil,
				)
			}
		}
	}
	if swap.ReturnOrder != nil && swap.ReturnOrder.Status != models.ReturnCanceled {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Return must be canceled before the swap can be canceled",
			nil,
		)
	}
	now := time.Now()
	swap.PaymentStatus = models.SwapPaymentCanceled
	swap.FulfillmentStatus = models.SwapFulfillmentCanceled
	swap.CanceledAt = &now
	if swap.Payment != nil {
		_, err := s.r.PaymentProviderService().SetContext(s.ctx).CancelPayment(swap.Payment)
		if err != nil {
			return nil, err
		}
	}
	if err := s.r.SwapRepository().Save(s.ctx, swap); err != nil {
		return nil, err
	}
	return swap, nil

}

func (s *SwapService) CreateFulfillment(swapId uuid.UUID, config *types.CreateShipmentConfig) (*models.Swap, *utils.ApplictaionError) {
	metadata := config.Metadata
	noNotification := config.NoNotification
	swap, err := s.Retrieve(swapId, &sql.Options{
		Relations: []string{
			"payment",
			"shipping_address",
			"additional_items.tax_lines",
			"additional_items.variant.product.profiles",
			"shipping_methods",
			"shipping_methods.shipping_option",
			"shipping_methods.tax_lines",
			"order",
			"order.region",
			"order.billing_address",
			"order.discounts",
			"order.discounts.rule",
			"order.payments",
		},
	})
	if err != nil {
		return nil, err
	}
	order := swap.Order
	if swap.CanceledAt != nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Canceled swap cannot be fulfilled",
			nil,
		)
	}
	if swap.FulfillmentStatus != models.SwapFulfillmentNotFulfilled && swap.FulfillmentStatus != models.SwapFulfillmentCanceled {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"The swap was already fulfilled",
			nil,
		)
	}
	if len(swap.ShippingMethods) == 0 {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Cannot fulfill an swap that doesn't have shipping methods",
			nil,
		)
	}
	evaluatedNoNotification := noNotification
	if !evaluatedNoNotification {
		evaluatedNoNotification = swap.NoNotification
	}

	fulfillmentOrder := &types.CreateFulfillmentOrder{}

	copier.CopyWithOption(&fulfillmentOrder, swap, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	fulfillmentOrder.Payments = []models.Payment{*swap.Payment}
	fulfillmentOrder.Email = order.Email
	fulfillmentOrder.Discounts = order.Discounts
	fulfillmentOrder.CurrencyCode = order.CurrencyCode
	fulfillmentOrder.TaxRate = order.TaxRate
	fulfillmentOrder.RegionId = order.RegionId.UUID
	fulfillmentOrder.Region = order.Region
	fulfillmentOrder.DisplayId = order.DisplayId
	fulfillmentOrder.BillingAddress = order.BillingAddress
	fulfillmentOrder.Items = swap.AdditionalItems
	fulfillmentOrder.ShippingMethods = swap.ShippingMethods
	fulfillmentOrder.IsSwap = true
	fulfillmentOrder.NoNotification = evaluatedNoNotification

	var items []types.FulFillmentItemType
	for _, item := range swap.AdditionalItems {
		items = append(items, types.FulFillmentItemType{
			ItemId:   item.Id,
			Quantity: item.Quantity,
		})
	}

	swap.Fulfillments, err = s.r.FulfillmentService().SetContext(s.ctx).CreateFulfillment(
		fulfillmentOrder,
		items,
		models.Fulfillment{
			Model: core.Model{
				Metadata: metadata,
			},
			SwapId:     uuid.NullUUID{UUID: swap.Id},
			LocationId: uuid.NullUUID{UUID: config.LocationId},
		},
	)

	if err != nil {
		return nil, err
	}
	var successfullyFulfilled []models.FulfillmentItem
	for _, f := range swap.Fulfillments {
		successfullyFulfilled = append(successfullyFulfilled, f.Items...)
	}
	swap.FulfillmentStatus = models.SwapFulfillmentFulfilled
	for _, item := range swap.AdditionalItems {
		var fulfillmentItem *models.FulfillmentItem
		for _, f := range successfullyFulfilled {
			if item.Id == f.ItemId.UUID {
				fulfillmentItem = &f
			}
		}
		if fulfillmentItem != nil {
			fulfilledQuantity := item.FulfilledQuantity + fulfillmentItem.Quantity
			_, err = s.r.LineItemService().SetContext(s.ctx).Update(item.Id, nil, &models.LineItem{
				FulfilledQuantity: fulfilledQuantity,
			}, &sql.Options{})
			if err != nil {
				return nil, err
			}
			if item.Quantity != fulfilledQuantity {
				swap.FulfillmentStatus = models.SwapFulfillmentRequiresAction
			}
		} else {
			if item.Quantity != item.FulfilledQuantity {
				swap.FulfillmentStatus = models.SwapFulfillmentRequiresAction
			}
		}
	}
	if err := s.r.SwapRepository().Save(s.ctx, swap); err != nil {
		return nil, err
	}
	// err = s.eventBus_.Emit(SwapServiceEventsFulfillmentCreated, &FulfillmentCreatedEvent{
	// 	ID:             swapId,
	// 	FulfillmentID:  result.ID,
	// 	NoNotification: evaluatedNoNotification,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	return swap, nil

}

func (s *SwapService) CancelFulfillment(fulfillmentId uuid.UUID) (*models.Swap, *utils.ApplictaionError) {
	canceled, err := s.r.FulfillmentService().SetContext(s.ctx).CancelFulfillment(fulfillmentId, nil)
	if err != nil {
		return nil, err
	}
	if canceled.SwapId.UUID == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Fufillment not related to a swap",
			nil,
		)
	}
	swap, err := s.Retrieve(canceled.SwapId.UUID, &sql.Options{})
	if err != nil {
		return nil, err
	}
	swap.FulfillmentStatus = models.SwapFulfillmentCanceled
	if err := s.r.SwapRepository().Save(s.ctx, swap); err != nil {
		return nil, err
	}
	return swap, nil

}

func (s *SwapService) CreateShipment(swapId uuid.UUID, fulfillmentId uuid.UUID, trackingLinks []models.TrackingLink, config *types.CreateShipmentConfig) (*models.Swap, *utils.ApplictaionError) {
	metadata := config.Metadata
	noNotification := config.NoNotification
	swap, err := s.Retrieve(swapId, &sql.Options{
		Relations: []string{"additional_items"},
	})
	if err != nil {
		return nil, err
	}
	if swap.CanceledAt != nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Canceled swap cannot be fulfilled as shipped",
			nil,
		)
	}
	evaluatedNoNotification := noNotification
	if !evaluatedNoNotification {
		evaluatedNoNotification = swap.NoNotification
	}
	shipment, err := s.r.FulfillmentService().SetContext(s.ctx).CreateShipment(fulfillmentId, trackingLinks, &types.CreateShipmentConfig{
		Metadata:       metadata,
		NoNotification: evaluatedNoNotification,
	})
	if err != nil {
		return nil, err
	}
	swap.FulfillmentStatus = models.SwapFulfillmentShipped
	for _, i := range swap.AdditionalItems {
		var shipped *models.FulfillmentItem
		for _, f := range shipment.Items {
			if i.Id == f.ItemId.UUID {
				shipped = &f
			}
		}
		if shipped != nil {
			shippedQty := i.ShippedQuantity + shipped.Quantity
			_, err = s.r.LineItemService().SetContext(s.ctx).Update(i.Id, nil, &models.LineItem{
				ShippedQuantity: shippedQty,
			}, &sql.Options{})
			if err != nil {
				return nil, err
			}
			if shippedQty != i.Quantity {
				swap.FulfillmentStatus = models.SwapFulfillmentPartiallyShipped
			}
		} else {
			if i.ShippedQuantity != i.Quantity {
				swap.FulfillmentStatus = models.SwapFulfillmentPartiallyShipped
			}
		}
	}
	if err := s.r.SwapRepository().Save(s.ctx, swap); err != nil {
		return nil, err
	}
	// err = s.eventBus_.Emit(SwapServiceEventsShipmentCreated, &ShipmentCreatedEvent{
	// 	ID:             swapId,
	// 	FulfillmentID:  shipment.ID,
	// 	NoNotification: swap.NoNotification,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	return swap, nil
}

func (s *SwapService) DeleteMetadata(swapId string, key string) (*models.Swap, *utils.ApplictaionError) {
	var swap *models.Swap = &models.Swap{}

	if err := swap.ParseUUID(swapId); err != nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			err.Error(),
			nil,
		)
	}

	query := sql.BuildQuery(swap, &sql.Options{})

	swap.Id = uuid.Nil

	if err := s.r.SwapRepository().FindOne(s.ctx, swap, query); err != nil {
		return nil, err
	}

	if err := s.r.SwapRepository().Save(s.ctx, swap); err != nil {
		return nil, err
	}
	// err = s.eventBus_.WithTransaction(transactionManager).Emit(CartServiceEventsUpdated, updatedSwap)
	// if err != nil {
	// 	return nil, err
	// }
	return swap, nil
}

func (s *SwapService) RegisterReceived(id uuid.UUID) (*models.Swap, *utils.ApplictaionError) {
	swap, err := s.Retrieve(id, &sql.Options{
		Relations: []string{"return_order", "return_order.items"},
	})
	if err != nil {
		return nil, err
	}
	if swap.CanceledAt != nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Canceled swap cannot be registered as received",
			nil,
		)
	}
	if swap.ReturnOrder.Status != models.ReturnReceived {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Swap is not received",
			nil,
		)
	}
	result, err := s.Retrieve(id, &sql.Options{})
	if err != nil {
		return nil, err
	}
	// err = s.eventBus_.Emit(SwapServiceEventsReceived, &ReceivedEvent{
	// 	ID:             id,
	// 	OrderID:        result.OrderID,
	// 	NoNotification: swap.NoNotification,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	return result, nil
}

func (s *SwapService) AreReturnItemsValid(returnItems []types.OrderReturnItem) (bool, *utils.ApplictaionError) {
	var itemId uuid.UUIDs

	for _, item := range returnItems {
		itemId = append(itemId, item.ItemId)
	}

	returnItemsEntities, err := s.r.LineItemService().SetContext(s.ctx).List(models.LineItem{}, &sql.Options{
		Relations:     []string{"order", "swap", "claim_order"},
		Specification: []sql.Specification{sql.In("id", itemId)},
	})
	if err != nil {
		return false, err
	}

	hasCanceledItem := false
	for _, item := range returnItemsEntities {
		if item.Order != nil && item.Order.CanceledAt != nil {
			hasCanceledItem = true
			break
		}
		if item.Swap != nil && item.Swap.CanceledAt != nil {
			hasCanceledItem = true
			break
		}
		if item.ClaimOrder != nil && item.ClaimOrder.CanceledAt != nil {
			hasCanceledItem = true
			break
		}
	}
	return !hasCanceledItem, nil
}

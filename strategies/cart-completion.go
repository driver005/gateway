package strategies

import (
	"context"
	"fmt"
	"reflect"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/services"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type CartCompletionStrategy struct {
	ctx context.Context
	r   Registry
}

func NewCartCompletionStrategy(
	r Registry,
) *CartCompletionStrategy {
	return &CartCompletionStrategy{
		context.Background(),
		r,
	}
}

func (s *CartCompletionStrategy) SetContext(context context.Context) *CartCompletionStrategy {
	s.ctx = context
	return s
}

func (s *CartCompletionStrategy) Complete(id uuid.UUID, idempotencyKey *models.IdempotencyKey, context types.RequestContext) (*interfaces.CartCompletionResponse, *utils.ApplictaionError) {
	var err *utils.ApplictaionError

	switch idempotencyKey.RecoveryPoint {
	case "started":
		idempotencyKey, err = s.r.IdempotencyKeyService().SetContext(s.ctx).WorkStage(idempotencyKey.IdempotencyKey, func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
			return s.handleCreateTaxLines(id)
		})
		if err != nil {
			break
		}
	case "tax_lines_created":
		idempotencyKey, err = s.r.IdempotencyKeyService().SetContext(s.ctx).WorkStage(idempotencyKey.IdempotencyKey, func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
			return s.handleTaxLineCreated(id, idempotencyKey, context)
		})
		if err != nil {
			break
		}
	case "payment_authorized":
		idempotencyKey, err = s.r.IdempotencyKeyService().SetContext(s.ctx).WorkStage(idempotencyKey.IdempotencyKey, func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
			return s.handlePaymentAuthorized(id)
		})
		if err != nil {
			break
		}
	default:
		idempotencyKey, err = s.r.IdempotencyKeyService().SetContext(s.ctx).Update(idempotencyKey.IdempotencyKey, &models.IdempotencyKey{
			RecoveryPoint: "finished",
			ResponseCode:  500,
			ResponseBody:  core.JSONB{"message": "Unknown recovery point"},
		})
		if err != nil {
			break
		}
	}

	if err != nil {
		if idempotencyKey.RecoveryPoint != "started" {
			_, errorObj := s.r.OrderService().SetContext(s.ctx).RetrieveByCartId(id, &sql.Options{})
			if errorObj != nil {
				if errorObj := s.r.CartService().SetContext(s.ctx).DeleteTaxLines(id); errorObj != nil {
					return nil, errorObj
				}
			}
		}
		return nil, err
	}

	return &interfaces.CartCompletionResponse{
		ResponseCode: idempotencyKey.ResponseCode,
		ResponseBody: idempotencyKey.ResponseBody,
	}, nil
}

func (s *CartCompletionStrategy) handleCreateTaxLines(id uuid.UUID) (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
	cart, err := s.r.CartService().SetContext(s.ctx).Retrieve(id, &sql.Options{
		Relations: []string{
			"customer",
			"discounts",
			"discounts.rule",
			"gift_cards",
			"items.variant.product.profiles",
			"items.adjustments",
			"region",
			"region.tax_rates",
			"shipping_address",
			"shipping_methods",
			"shipping_methods.shipping_option",
		},
	}, services.TotalsConfig{})
	if err != nil {
		return nil, err
	}
	if cart.CompletedAt != nil {
		if cart.Type == "swap" {
			swapId, ok := cart.Metadata["swap_id"].(uuid.UUID)
			if ok {
				swap, err := s.r.SwapService().SetContext(s.ctx).Retrieve(swapId, &sql.Options{
					Relations: []string{"shipping_address"},
				})
				if err != nil {
					return nil, err
				}
				return &types.IdempotencyCallbackResult{
					ResponseCode: 200,
					ResponseBody: map[string]interface{}{
						"data": swap,
						"type": "swap",
					},
				}, nil
			}
		}
		order, err := s.r.OrderService().SetContext(s.ctx).RetrieveByCartIdWithTotals(id, &sql.Options{
			Relations: []string{"shipping_address", "items", "payments"},
		})
		if err != nil {
			return nil, err
		}

		return &types.IdempotencyCallbackResult{
			ResponseCode: 200,
			ResponseBody: map[string]interface{}{
				"data": order,
				"type": "order",
			},
		}, nil
	}
	err = s.r.CartService().SetContext(s.ctx).CreateTaxLines(uuid.Nil, cart)
	if err != nil {
		return nil, err
	}
	return &types.IdempotencyCallbackResult{
		RecoveryPoint: "tax_lines_created",
	}, nil
}

func (s *CartCompletionStrategy) handleTaxLineCreated(id uuid.UUID, idempotencyKey *models.IdempotencyKey, context types.RequestContext) (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
	res, err := s.handleCreateTaxLines(id)
	if !reflect.ValueOf(res.ResponseCode).IsZero() {
		return nil, err
	}
	cart, err := s.r.CartService().AuthorizePayment(id, nil, map[string]interface{}{
		"idempotency_key": idempotencyKey,
	})
	if err != nil {
		return nil, err
	}
	if cart.PaymentSession != nil {
		if cart.PaymentSession.Status == models.PaymentSessionStatusRequiresMore || cart.PaymentSession.Status == models.PaymentSessionStatusPending {
			if err := s.r.CartService().DeleteTaxLines(id); err != nil {
				return nil, err
			}
			return &types.IdempotencyCallbackResult{
				ResponseCode: 200,
				ResponseBody: map[string]interface{}{
					"data":           cart,
					"payment_status": cart.PaymentSession.Status,
					"type":           "cart",
				},
			}, nil
		}
	}
	return &types.IdempotencyCallbackResult{
		RecoveryPoint: "payment_authorized",
	}, nil
}

func (s *CartCompletionStrategy) removeReservations(reservations [][]interfaces.ReservationItemDTO) *utils.ApplictaionError {
	if s.r.InventoryService() != nil {
		for _, reservationItemArr := range reservations {
			for _, reservation := range reservationItemArr {
				if err := s.r.InventoryService().DeleteReservationItem(s.ctx, uuid.UUIDs{reservation.Id}); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (s *CartCompletionStrategy) handlePaymentAuthorized(id uuid.UUID) (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
	res, err := s.handleCreateTaxLines(id)
	if !reflect.ValueOf(res.ResponseCode).IsZero() {
		return nil, err
	}

	cart, err := s.r.CartService().SetContext(s.ctx).RetrieveWithTotals(id, &sql.Options{
		Relations: []string{
			"region",
			"payment",
			"payment_sessions",
			"items.variant.product.profiles",
		},
	}, services.TotalsConfig{})
	if err != nil {
		return nil, err
	}
	allowBackorder := false
	if cart.Type == "swap" {
		swap, err := s.r.SwapService().SetContext(s.ctx).RetrieveByCartId(id, []string{})
		if err != nil {
			return nil, err
		}
		allowBackorder = swap.AllowBackorder
	}
	var reservations [][]interfaces.ReservationItemDTO
	if !allowBackorder {
		reservations = make([][]interfaces.ReservationItemDTO, len(cart.Items))
		for i, item := range cart.Items {
			if item.VariantId.UUID != uuid.Nil {
				inventoryConfirmed, err := s.r.ProductVariantInventoryService().SetContext(s.ctx).ConfirmInventory(item.VariantId.UUID, item.Quantity, map[string]interface{}{
					"salesChannelId": cart.SalesChannelId.UUID,
				})
				if err != nil {
					return nil, err
				}
				if !inventoryConfirmed {
					return nil, utils.NewApplictaionError(
						utils.INSUFFICIENT_INVENTORY,
						fmt.Sprintf("Variant with id: %s does not have the required inventory", item.VariantId.UUID),
					)
				}
				reservations[i], err = s.r.ProductVariantInventoryService().SetContext(s.ctx).ReserveQuantity(item.VariantId.UUID, item.Quantity, services.ReserveQuantityContext{
					LineItemId:     item.Id,
					SalesChannelId: cart.SalesChannelId.UUID,
				})
				if err != nil {
					return nil, err
				}
			}
		}
		if s.r.InventoryService() != nil {
			var ids uuid.UUIDs
			for _, reservationItemArr := range reservations {
				for _, item := range reservationItemArr {
					ids = append(ids, item.Id)
				}
			}
			// err = s.eventBusService_.Emit("reservation-items.bulk-created", map[string]interface{}{
			// 	"ids": ids,
			// })
			// if err != nil {
			// 	return res, err
			// }
		}
	}
	if cart.Type == "swap" {
		swapId, ok := cart.Metadata["swap_id"].(uuid.UUID)
		if ok {
			swap, err := s.r.SwapService().SetContext(s.ctx).RegisterCartCompletion(swapId)
			if err != nil {
				if err := s.removeReservations(reservations); err != nil {
					return nil, err
				}
				if err.Type == utils.INSUFFICIENT_INVENTORY {
					return &types.IdempotencyCallbackResult{
						ResponseCode: 409,
						ResponseBody: map[string]interface{}{
							"message": err.Message,
							"type":    err.Type,
							"code":    err.Code,
						},
					}, nil
				} else {
					return nil, err
				}
			}
			swap, err = s.r.SwapService().SetContext(s.ctx).Retrieve(swap.Id, &sql.Options{
				Relations: []string{"shipping_address"},
			})
			if err != nil {
				return nil, err
			}
			return &types.IdempotencyCallbackResult{
				ResponseCode: 200,
				ResponseBody: map[string]interface{}{
					"data": swap,
					"type": "swap",
				},
			}, nil
		}

	}
	if cart.Payment == nil && !reflect.ValueOf(cart.Total).IsZero() && cart.Total > 0 {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Cart payment not authorized",
		)
	}
	order, err := s.r.OrderService().SetContext(s.ctx).CreateFromCart(uuid.Nil, cart)
	if err != nil {
		err = s.removeReservations(reservations)
		if err != nil {
			return res, err
		}
		if err.Message == "Order from cart already exists" {
			order, err := s.r.OrderService().SetContext(s.ctx).RetrieveByCartIdWithTotals(id, &sql.Options{
				Relations: []string{"shipping_address", "items", "payments"},
			})
			if err != nil {
				return nil, err
			}
			return &types.IdempotencyCallbackResult{
				ResponseCode: 200,
				ResponseBody: map[string]interface{}{
					"data": order,
					"type": "order",
				},
			}, nil
		} else if err.Type == utils.INSUFFICIENT_INVENTORY {
			return &types.IdempotencyCallbackResult{
				ResponseCode: 409,
				ResponseBody: map[string]interface{}{
					"message": err.Message,
					"type":    err.Type,
					"code":    err.Code,
				},
			}, nil
		} else {
			return res, err
		}
	}
	order, err = s.r.OrderService().SetContext(s.ctx).RetrieveByIdWithTotals(order.Id, &sql.Options{
		Relations: []string{"shipping_address", "items", "payments"},
	}, types.TotalsContext{})
	if err != nil {
		return res, err
	}
	return &types.IdempotencyCallbackResult{
		ResponseCode: 200,
		ResponseBody: map[string]interface{}{
			"data": order,
			"type": "order",
		},
	}, nil
}

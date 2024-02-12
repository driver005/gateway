package store

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type Swap struct {
	r Registry
}

func NewSwap(r Registry) *Swap {
	m := Swap{r: r}
	return &m
}

func (m *Swap) SetRoutes(router fiber.Router) {
	route := router.Group("/swaps")
	route.Get("/:cart_id", m.GetByCart)
	route.Post("", m.Create)
}

func (m *Swap) GetByCart(context fiber.Ctx) error {
	id, err := utils.ParseUUID(context.Params("cart_id"))
	if err != nil {
		return err
	}

	result, err := m.r.SwapService().SetContext(context.Context()).RetrieveByCartId(id, []string{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Swap) Create(context fiber.Ctx) error {
	model, config, err := api.BindList[types.CreateSwap](context)
	if err != nil {
		return err
	}

	headerKey := context.Get("Idempotency-Key")

	idempotencyKey := &models.IdempotencyKey{}
	idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).InitializeRequest(headerKey, context.Method(), nil, context.Path())
	if err != nil {
		return err
	}

	context.Set("Access-Control-Expose-Headers", "Idempotency-Key")
	context.Set("Idempotency-Key", idempotencyKey.IdempotencyKey)

	if idempotencyKey.RecoveryPoint == "started" {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).WorkStage(idempotencyKey.IdempotencyKey, func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
			order, err := m.r.OrderService().SetContext(context.Context()).RetrieveById(model.OrderId, &sql.Options{
				Selects: []string{"refunded_total", "total"},
				Relations: []string{
					"items.variant",
					"items.tax_lines",
					"swaps.additional_items.variant.product.profiles",
					"swaps.additional_items.tax_lines",
				},
			})
			if err != nil {
				return nil, err
			}

			var returnShipping *types.CreateClaimReturnShippingInput
			if model.ReturnShippingOption != uuid.Nil {
				returnShipping.OptionId = model.ReturnShippingOption
			}

			swap, err := m.r.SwapService().SetContext(context.Context()).Create(
				order,
				model.ReturnItems,
				model.AdditionalItems,
				returnShipping,
				map[string]interface{}{
					"idempotency_key": idempotencyKey.IdempotencyKey,
					"no_notification": true,
				},
			)
			if err != nil {
				return nil, err
			}

			if _, err := m.r.SwapService().SetContext(context.Context()).CreateCart(swap.Id, []types.CreateCustomShippingOptionInput{}, map[string]interface{}{}); err != nil {
				return nil, err
			}

			returnOrder, err := m.r.ReturnService().SetContext(context.Context()).RetrieveBySwap(swap.Id, []string{})
			if err != nil {
				return nil, err
			}

			if _, err := m.r.ReturnService().SetContext(context.Context()).Fulfill(returnOrder.Id); err != nil {
				return nil, err
			}

			return &types.IdempotencyCallbackResult{RecoveryPoint: "swap_created"}, nil
		})
		if err != nil {
			return err
		}
	} else if idempotencyKey.RecoveryPoint == "swap_created" {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).WorkStage(idempotencyKey.IdempotencyKey, func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
			swaps, err := m.r.SwapService().SetContext(context.Context()).List(&types.FilterableSwap{IdempotencyKey: idempotencyKey.IdempotencyKey}, &sql.Options{})
			if err != nil {
				return nil, err
			}

			if len(swaps) == 0 {
				return nil, utils.NewApplictaionError(
					utils.INVALID_DATA,
					"Swap not found",
				)
			}

			result, err := m.r.OrderService().SetContext(context.Context()).RetrieveById(swaps[0].Id, config)
			if err != nil {
				return nil, err
			}

			return &types.IdempotencyCallbackResult{
				ResponseCode: 200,
				ResponseBody: core.JSONB{
					"data": result,
				},
			}, nil
		})
		if err != nil {
			return err
		}
	} else {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).Update(idempotencyKey.IdempotencyKey, &models.IdempotencyKey{
			RecoveryPoint: "finished",
			ResponseCode:  500,
			ResponseBody:  core.JSONB{"message": "Unknown recovery point"},
		})
		if err != nil {
			return err
		}
	}

	return context.Status(idempotencyKey.ResponseCode).JSON(idempotencyKey.ResponseBody)
}

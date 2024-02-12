package store

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
)

type Return struct {
	r Registry
}

func NewReturn(r Registry) *Return {
	m := Return{r: r}
	return &m
}

func (m *Return) SetRoutes(router fiber.Router) {
	route := router.Group("/returns")
	route.Post("", m.Create)
}

func (m *Return) Create(context fiber.Ctx) error {
	model, err := api.Bind[types.CreateReturn](context, m.r.Validator())
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
			data := &types.CreateReturnInput{
				OrderId: model.OrderId,
				Items:   model.Items,
			}

			if model.ReturnShipping != nil {
				data.ShippingMethod = &types.CreateClaimReturnShippingInput{
					OptionId: model.ReturnShipping.OptionId,
				}
			}

			createdReturn, err := m.r.ReturnService().SetContext(context.Context()).Create(data)
			if err != nil {
				return nil, err
			}

			if model.ReturnShipping != nil {
				if _, err := m.r.ReturnService().SetContext(context.Context()).Fulfill(createdReturn.Id); err != nil {
					return nil, err
				}
			}

			// await eventBus
			//         .withTransaction(manager)
			//         .emit("order.return_requested", {
			//           id: returnDto.order_id,
			//           return_id: createdReturn.id,
			//         })

			return &types.IdempotencyCallbackResult{RecoveryPoint: "return_requested"}, nil
		})
		if err != nil {
			return err
		}
	} else if idempotencyKey.RecoveryPoint == "return_requested" {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).WorkStage(idempotencyKey.IdempotencyKey, func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
			returnOrders, err := m.r.ReturnService().SetContext(context.Context()).List(&types.FilterableReturn{IdempotencyKey: idempotencyKey.IdempotencyKey}, &sql.Options{})
			if err != nil {
				return nil, err
			}

			if len(returnOrders) == 0 {
				return nil, utils.NewApplictaionError(
					utils.INVALID_DATA,
					"Return not found",
				)
			}

			return &types.IdempotencyCallbackResult{
				ResponseCode: 200,
				ResponseBody: core.JSONB{
					"data": returnOrders[0],
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

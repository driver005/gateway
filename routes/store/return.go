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

// @oas:path [post] /store/returns
// operationId: "PostReturns"
// summary: "Create Return"
// description: "Create a Return for an Order. If a return shipping method is specified, the return is automatically fulfilled."
// externalDocs:
//
//	description: "How to create a return in a storefront"
//	url: "https://docs.medusajs.com/modules/orders/storefront/create-return"
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/StorePostReturnsReq"
//
// x-codegen:
//
//	method: create
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.returns.create({
//     order_id,
//     items: [
//     {
//     item_id,
//     quantity: 1
//     }
//     ]
//     })
//     .then((data) => {
//     console.log(data.return.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useCreateReturn } from "medusa-react"
//
//     type CreateReturnData = {
//     items: {
//     item_id: string,
//     quantity: number
//     }[]
//     return_shipping: {
//     option_id: string
//     }
//     }
//
//     type Props = {
//     orderId: string
//     }
//
//     const CreateReturn = ({ orderId }: Props) => {
//     const createReturn = useCreateReturn()
//     // ...
//
//     const handleCreate = (data: CreateReturnData) => {
//     createReturn.mutate({
//     ...data,
//     order_id: orderId
//     }, {
//     onSuccess: ({ return: returnData }) => {
//     console.log(returnData.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default CreateReturn
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/store/returns' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "order_id": "asfasf",
//     "items": [
//     {
//     "item_id": "assfasf",
//     "quantity": 1
//     }
//     ]
//     }'
//
// tags:
//   - Returns
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreReturnsRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
//	"404":
//	  description: "Not Found"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/not_found_error"
//	"409":
//	  description: "Invalid State"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/invalid_state_error"
//	"422":
//	  description: "Invalid Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/invalid_request_error"
//	"500":
//	  description: "Internal Server"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/responses/500_error"
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

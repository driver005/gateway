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

// @oas:path [get] /store/swaps/{cart_id}
// operationId: GetSwapsSwapCartId
// summary: Get by Cart ID
// description: "Retrieve a Swap's details by the ID of its cart."
// parameters:
//   - (path) cart_id {string} The id of the Cart
//
// x-codegen:
//
//	method: retrieveByCartId
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.swaps.retrieveByCartId(cartId)
//     .then(({ swap }) => {
//     console.log(swap.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useCartSwap } from "medusa-react"
//     type Props = {
//     cartId: string
//     }
//
//     const Swap = ({ cartId }: Props) => {
//     const {
//     swap,
//     isLoading,
//     } = useCartSwap(cartId)
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {swap && <span>{swap.id}</span>}
//
//     </div>
//     )
//     }
//
//     export default Swap
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '{backend_url}/store/swaps/{cart_id}'
//
// tags:
//   - Swaps
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreSwapsRes"
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

// @oas:path [post] /store/swaps
// operationId: PostSwaps
// summary: Create a Swap
// description: |
//
//	Create a Swap for an Order. This will also create a return and associate it with the swap. If a return shipping option is specified, the return will automatically be fulfilled.
//	To complete the swap, you must use the Complete Cart API Route passing it the ID of the swap's cart.
//
//	An idempotency key will be generated if none is provided in the header `Idempotency-Key` and added to
//	the response. If an error occurs during swap creation or the request is interrupted for any reason, the swap creation can be retried by passing the idempotency
//	key in the `Idempotency-Key` header.
//
// externalDocs:
//
//	description: "How to create a swap"
//	url: "https://docs.medusajs.com/modules/orders/storefront/create-swap"
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/StorePostSwapsReq"
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
//     medusa.swaps.create({
//     order_id,
//     return_items: [
//     {
//     item_id,
//     quantity: 1
//     }
//     ],
//     additional_items: [
//     {
//     variant_id,
//     quantity: 1
//     }
//     ]
//     })
//     .then(({ swap }) => {
//     console.log(swap.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useCreateSwap } from "medusa-react"
//
//     type Props = {
//     orderId: string
//     }
//
//     type CreateData = {
//     return_items: {
//     item_id: string
//     quantity: number
//     }[]
//     additional_items: {
//     variant_id: string
//     quantity: number
//     }[]
//     return_shipping_option: string
//     }
//
//     const CreateSwap = ({
//     orderId
//     }: Props) => {
//     const createSwap = useCreateSwap()
//     // ...
//
//     const handleCreate = (
//     data: CreateData
//     ) => {
//     createSwap.mutate({
//     ...data,
//     order_id: orderId
//     }, {
//     onSuccess: ({ swap }) => {
//     console.log(swap.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default CreateSwap
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/store/swaps' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "order_id": "{order_id}",
//     "return_items": [
//     {
//     "item_id": "{item_id}",
//     "quantity": 1
//     }
//     ],
//     "additional_items": [
//     {
//     "variant_id": "{variant_id}",
//     "quantity": 1
//     }
//     ]
//     }'
//
// tags:
//   - Swaps
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreSwapsRes"
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

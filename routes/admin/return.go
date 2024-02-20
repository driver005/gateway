package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type Return struct {
	r    Registry
	name string
}

func NewReturn(r Registry) *Return {
	m := Return{r: r, name: "return"}
	return &m
}

func (m *Return) SetRoutes(router fiber.Router) {
	route := router.Group("/returns")
	route.Get("", m.List)

	route.Post("/:id/receive", m.Receive)
	route.Post("/:id/cancel", m.Cancel)
}

// @oas:path [get] /admin/returns
// operationId: "GetReturns"
// summary: "List Returns"
// description: "Retrieve a list of Returns. The returns can be paginated."
// parameters:
//   - (query) limit=50 {number} Limit the number of Returns returned.
//   - (query) offset=0 {number} The number of Returns to skip when retrieving the Returns.
//
// x-codegen:
//
//	method: list
//	queryParams: AdminGetReturnsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.returns.list()
//     .then(({ returns, limit, offset, count }) => {
//     console.log(returns.length)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminReturns } from "medusa-react"
//
//     const Returns = () => {
//     const { returns, isLoading } = useAdminReturns()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {returns && !returns.length && (
//     <span>No Returns</span>
//     )}
//     {returns && returns.length > 0 && (
//     <ul>
//     {returns.map((returnData) => (
//     <li key={returnData.id}>
//     {returnData.status}
//     </li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default Returns
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/returns' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Returns
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminReturnsListRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
//	"401":
//	  description: "Unauthorized"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/unauthorized"
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
func (m *Return) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableReturn](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.ReturnService().SetContext(context.Context()).ListAndCount(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"returns": result,
		"count":   count,
		"offset":  config.Skip,
		"limit":   config.Take,
	})
}

// @oas:path [post] /admin/returns/{id}/cancel
// operationId: "PostReturnsReturnCancel"
// summary: "Cancel a Return"
// description: "Registers a Return as canceled. The return can be associated with an order, claim, or swap."
// parameters:
//   - (path) id=* {string} The ID of the Return.
//
// x-codegen:
//
//	method: cancel
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.returns.cancel(returnId)
//     .then(({ order }) => {
//     console.log(order.id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCancelReturn } from "medusa-react"
//
//     type Props = {
//     returnId: string
//     }
//
//     const Return = ({ returnId }: Props) => {
//     const cancelReturn = useAdminCancelReturn(
//     returnId
//     )
//     // ...
//
//     const handleCancel = () => {
//     cancelReturn.mutate(void 0, {
//     onSuccess: ({ order }) => {
//     console.log(order.returns)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Return
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/returns/{id}/cancel' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Returns
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminReturnsCancelRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
//	"401":
//	  description: "Unauthorized"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/unauthorized"
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
func (m *Return) Cancel(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	model, err := m.r.ReturnService().SetContext(context.Context()).Cancel(id)
	if err != nil {
		return err
	}

	var orderId uuid.UUID

	if model.SwapId.UUID != uuid.Nil {
		data, err := m.r.SwapService().SetContext(context.Context()).Retrieve(model.SwapId.UUID, &sql.Options{})
		if err != nil {
			return err
		}

		orderId = data.OrderId.UUID
	} else if model.ClaimOrderId.UUID != uuid.Nil {
		data, err := m.r.ClaimService().SetContext(context.Context()).Retrieve(model.ClaimOrderId.UUID, &sql.Options{})
		if err != nil {
			return err
		}

		orderId = data.OrderId.UUID
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveById(orderId, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"order": result,
	})
}

// @oas:path [post] /admin/returns/{id}/receive
// operationId: "PostReturnsReturnReceive"
// summary: "Receive a Return"
// description: "Mark a Return as received. This also updates the status of associated order, claim, or swap accordingly."
// parameters:
//   - (path) id=* {string} The ID of the Return.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostReturnsReturnReceiveReq"
//
// x-codegen:
//
//	method: receive
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.returns.receive(returnId, {
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
//     import { useAdminReceiveReturn } from "medusa-react"
//
//     type ReceiveReturnData = {
//     items: {
//     item_id: string
//     quantity: number
//     }[]
//     }
//
//     type Props = {
//     returnId: string
//     }
//
//     const Return = ({ returnId }: Props) => {
//     const receiveReturn = useAdminReceiveReturn(
//     returnId
//     )
//     // ...
//
//     const handleReceive = (data: ReceiveReturnData) => {
//     receiveReturn.mutate(data, {
//     onSuccess: ({ return: dataReturn }) => {
//     console.log(dataReturn.status)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Return
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/returns/{id}/receive' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "items": [
//     {
//     "item_id": "asafg",
//     "quantity": 1
//     }
//     ]
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Returns
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminReturnsRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
//	"401":
//	  description: "Unauthorized"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/unauthorized"
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
func (m *Return) Receive(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.ReturnReceive](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	refundAmount := model.Refund

	if refundAmount < 0 {
		refundAmount = 0
	}

	receivedReturn, err := m.r.ReturnService().SetContext(context.Context()).Receive(id, model.Items, &refundAmount, true, map[string]interface{}{
		"locationId": model.LocationId,
	})
	if err != nil {
		return err
	}

	if receivedReturn.OrderId.UUID != uuid.Nil {
		if _, err := m.r.OrderService().SetContext(context.Context()).RegisterReturnReceived(receivedReturn.OrderId.UUID, receivedReturn, &refundAmount); err != nil {
			return err
		}
	}

	if receivedReturn.SwapId.UUID != uuid.Nil {
		if _, err := m.r.SwapService().SetContext(context.Context()).RegisterReceived(receivedReturn.SwapId.UUID); err != nil {
			return err
		}
	}

	result, err := m.r.ReturnService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

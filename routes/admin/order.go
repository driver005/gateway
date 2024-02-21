package admin

import (
	"fmt"
	"reflect"

	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/services"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

const includes bool = true

type Order struct {
	r    Registry
	name string
}

func NewOrder(r Registry) *Order {
	m := Order{r: r, name: "order"}
	return &m
}

func (m *Order) SetRoutes(router fiber.Router) {
	route := router.Group("/orders")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("/:id", m.Update)

	route.Post("/:id/complete", m.Complete)
	route.Post("/:id/cancel", m.Cancel)
	route.Post("/:id/archive", m.Archive)
	route.Post("/:id/refund", m.RefundPayment)
	route.Post("/:id/capture", m.CapturePayment)
	route.Post("/:id/shipment", m.CreateShipment)
	route.Post("/:id/return", m.RequestReturn)
	route.Post("/:id/shipping-methods", m.AddShippingMethod)

	route.Post("/:id/swaps", m.CreateSwap)
	route.Post("/:id/swaps/:swap_id/cancel", m.CancelSwap)
	route.Post("/:id/swaps/:swap_id/fulfillments", m.FulfillSwap)
	route.Post("/:id/swaps/:swap_id/shipments", m.CreateSwapShipment)
	route.Post("/:id/swaps/:swap_id/process-payment", m.ProcessSwapPayment)

	route.Post("/:id/claims", m.CreateClaim)
	route.Post("/:id/claims/:claim_id/cancel", m.CancelClaim)
	route.Post("/:id/claims/:claim_id", m.UpdateClaim)
	route.Post("/:id/claims/:claim_id/fulfillments", m.FulfillClaim)
	route.Post("/:id/claims/:claim_id/shipments", m.CreateClaimShippment)

	route.Post("/:id/reservations", m.GetReservations)
	route.Post("/:id/line-items/:line_item_id/reserve", m.CreateReservationForLineItem)

	route.Post("/:id/fulfillment", m.CreateFulfillment)
	route.Post("/:id/fulfillments/:fulfillment_id/cancel", m.CancelFullfillment)
	route.Post("/:id/swaps/:swap_id/fulfillments/:fulfillment_id/cancel", m.CancelFullfillmentSwap)
	route.Post("/:id/claims/:claim_id/fulfillments/:fulfillment_id/cancel", m.CancelFullfillmentClaim)
}

// @oas:path [get] /admin/orders/{id}
// operationId: "GetOrdersOrder"
// summary: "Get an Order"
// description: "Retrieve an Order's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Order.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order.
//
// x-codegen:
//
//	method: retrieve
//	queryParams: AdminGetOrdersOrderParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orders.retrieve(orderId)
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminOrder } from "medusa-react"
//
//     type Props = {
//     orderId: string
//     }
//
//     const Order = ({ orderId }: Props) => {
//     const {
//     order,
//     isLoading,
//     } = useAdminOrder(orderId)
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {order && <span>{order.display_id}</span>}
//
//     </div>
//     )
//     }
//
//     export default Order
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/orders/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrdersRes"
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
func (m *Order) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveById(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [get] /admin/orders
// operationId: "GetOrders"
// summary: "List Orders"
// description: "Retrieve a list of Orders. The orders can be filtered by fields such as `status` or `display_id`. The order can also be paginated."
// x-authenticated: true
// parameters:
//   - (query) q {string} term to search orders' shipping address, first name, email, and display ID
//   - (query) id {string} Filter by ID.
//   - in: query
//     name: status
//     style: form
//     explode: false
//     description: Filter by status
//     schema:
//     type: array
//     items:
//     type: string
//     enum: [pending, completed, archived, canceled, requires_action]
//   - in: query
//     name: fulfillment_status
//     style: form
//     explode: false
//     description: Filter by fulfillment status
//     schema:
//     type: array
//     items:
//     type: string
//     enum: [not_fulfilled, fulfilled, partially_fulfilled, shipped, partially_shipped, canceled, returned, partially_returned, requires_action]
//   - in: query
//     name: payment_status
//     style: form
//     explode: false
//     description: Filter by payment status
//     schema:
//     type: array
//     items:
//     type: string
//     enum: [captured, awaiting, not_paid, refunded, partially_refunded, canceled, requires_action]
//   - (query) display_id {string} Filter by display ID
//   - (query) cart_id {string} Filter by cart ID
//   - (query) customer_id {string} Filter by customer ID
//   - (query) email {string} Filter by email
//   - in: query
//     name: region_id
//     style: form
//     explode: false
//     description: Filter by region IDs.
//     schema:
//     oneOf:
//   - type: string
//     description: ID of a Region.
//   - type: array
//     items:
//     type: string
//     description: ID of a Region.
//   - in: query
//     name: currency_code
//     style: form
//     explode: false
//     description: Filter by currency codes.
//     schema:
//     type: string
//     externalDocs:
//     url: https://en.wikipedia.org/wiki/ISO_4217#Active_codes
//     description: See a list of codes.
//   - (query) tax_rate {string} Filter by tax rate.
//   - in: query
//     name: created_at
//     description: Filter by a creation date range.
//     schema:
//     type: object
//     properties:
//     lt:
//     type: string
//     description: filter by dates less than this date
//     format: date
//     gt:
//     type: string
//     description: filter by dates greater than this date
//     format: date
//     lte:
//     type: string
//     description: filter by dates less than or equal to this date
//     format: date
//     gte:
//     type: string
//     description: filter by dates greater than or equal to this date
//     format: date
//   - in: query
//     name: updated_at
//     description: Filter by an update date range.
//     schema:
//     type: object
//     properties:
//     lt:
//     type: string
//     description: filter by dates less than this date
//     format: date
//     gt:
//     type: string
//     description: filter by dates greater than this date
//     format: date
//     lte:
//     type: string
//     description: filter by dates less than or equal to this date
//     format: date
//     gte:
//     type: string
//     description: filter by dates greater than or equal to this date
//     format: date
//   - in: query
//     name: canceled_at
//     description: Filter by a cancelation date range.
//     schema:
//     type: object
//     properties:
//     lt:
//     type: string
//     description: filter by dates less than this date
//     format: date
//     gt:
//     type: string
//     description: filter by dates greater than this date
//     format: date
//     lte:
//     type: string
//     description: filter by dates less than or equal to this date
//     format: date
//     gte:
//     type: string
//     description: filter by dates greater than or equal to this date
//     format: date
//   - in: query
//     name: sales_channel_id
//     style: form
//     explode: false
//     description: Filter by Sales Channel IDs
//     schema:
//     type: array
//     items:
//     type: string
//     description: The ID of a Sales Channel
//   - (query) offset=0 {integer} The number of orders to skip when retrieving the orders.
//   - (query) limit=50 {integer} Limit the number of orders returned.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order.
//   - (query) order {string} A order field to sort-order the retrieved orders by.
//
// x-codegen:
//
//	method: list
//	queryParams: AdminGetOrdersParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orders.list()
//     .then(({ orders, limit, offset, count }) => {
//     console.log(orders.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminOrders } from "medusa-react"
//
//     const Orders = () => {
//     const { orders, isLoading } = useAdminOrders()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {orders && !orders.length && <span>No Orders</span>}
//     {orders && orders.length > 0 && (
//     <ul>
//     {orders.map((order) => (
//     <li key={order.id}>{order.display_id}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default Orders
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/orders' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrdersListRes"
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
func (m *Order) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableOrder](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.OrderService().SetContext(context.Context()).ListAndCount(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"orders": result,
		"count":  count,
		"offset": config.Skip,
		"limit":  config.Take,
	})
}

// @oas:path [post] /admin/orders/{id}
// operationId: "PostOrdersOrder"
// summary: "Update an Order"
// description: "Update and order's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Order.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostOrdersOrderReq"
//
// x-codegen:
//
//	method: update
//	params: AdminPostOrdersOrderParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orders.update(orderId, {
//     email: "user@example.com"
//     })
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdateOrder } from "medusa-react"
//
//     type Props = {
//     orderId: string
//     }
//
//     const Order = ({ orderId }: Props) => {
//     const updateOrder = useAdminUpdateOrder(
//     orderId
//     )
//
//     const handleUpdate = (
//     email: string
//     ) => {
//     updateOrder.mutate({
//     email,
//     }, {
//     onSuccess: ({ order }) => {
//     console.log(order.email)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Order
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/orders/adasda' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "email": "user@example.com"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrdersRes"
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
func (m *Order) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateOrderInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/orders/{id}/shipping-methods
// operationId: "PostOrdersOrderShippingMethods"
// summary: "Add a Shipping Method"
// description: "Add a Shipping Method to an Order. If another Shipping Method exists with the same Shipping Profile, the previous Shipping Method will be replaced."
// parameters:
//   - (path) id=* {string} The ID of the Order.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostOrdersOrderShippingMethodsReq"
//
// x-authenticated: true
// x-codegen:
//
//	method: addShippingMethod
//	params: AdminPostOrdersOrderShippingMethodsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orders.addShippingMethod(orderId, {
//     price: 1000,
//     option_id
//     })
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminAddShippingMethod } from "medusa-react"
//
//     type Props = {
//     orderId: string
//     }
//
//     const Order = ({ orderId }: Props) => {
//     const addShippingMethod = useAdminAddShippingMethod(
//     orderId
//     )
//     // ...
//
//     const handleAddShippingMethod = (
//     optionId: string,
//     price: number
//     ) => {
//     addShippingMethod.mutate({
//     option_id: optionId,
//     price
//     }, {
//     onSuccess: ({ order }) => {
//     console.log(order.shipping_methods)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Order
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/orders/{id}/shipping-methods' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "price": 1000,
//     "option_id": "{option_id}"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrdersRes"
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
func (m *Order) AddShippingMethod(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.OrderShippingMethod](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if _, err := m.r.OrderService().SetContext(context.Context()).AddShippingMethod(id, model.OptionId, model.Data, &types.CreateShippingMethodDto{
		CreateShippingMethod: types.CreateShippingMethod{
			Price: model.Price,
		},
	}); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/orders/{id}/archive
// operationId: "PostOrdersOrderArchive"
// summary: "Archive Order"
// description: "Archive an order and change its status."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Order.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order.
//
// x-codegen:
//
//	method: archive
//	params: AdminPostOrdersOrderArchiveParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orders.archive(orderId)
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminArchiveOrder } from "medusa-react"
//
//     type Props = {
//     orderId: string
//     }
//
//     const Order = ({ orderId }: Props) => {
//     const archiveOrder = useAdminArchiveOrder(
//     orderId
//     )
//     // ...
//
//     const handleArchivingOrder = () => {
//     archiveOrder.mutate(void 0, {
//     onSuccess: ({ order }) => {
//     console.log(order.status)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Order
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/orders/{id}/archive' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrdersRes"
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
func (m *Order) Archive(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	if _, err := m.r.OrderService().SetContext(context.Context()).Archive(id); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/orders/{id}/cancel
// operationId: "PostOrdersOrderCancel"
// summary: "Cancel an Order"
// description: "Cancel an order and change its status. This will also cancel any associated Fulfillments and Payments, and it may fail if the Payment or Fulfillment Provider is unable to cancel the Payment/Fulfillment."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Order.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order.
//
// x-codegen:
//
//	method: cancel
//	params: AdminPostOrdersOrderCancel
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orders.cancel(orderId)
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCancelOrder } from "medusa-react"
//
//     type Props = {
//     orderId: string
//     }
//
//     const Order = ({ orderId }: Props) => {
//     const cancelOrder = useAdminCancelOrder(
//     orderId
//     )
//     // ...
//
//     const handleCancel = () => {
//     cancelOrder.mutate(void 0, {
//     onSuccess: ({ order }) => {
//     console.log(order.status)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Order
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/orders/{id}/cancel' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrdersRes"
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
func (m *Order) Cancel(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	if _, err := m.r.OrderService().SetContext(context.Context()).Cancel(id); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})

}

// @oas:path [post] /admin/orders/{id}/swaps/{swap_id}/cancel
// operationId: "PostOrdersSwapCancel"
// summary: "Cancel a Swap"
// description: "Cancel a Swap and change its status."
// x-authenticated: true
// externalDocs:
//
//	description: Canceling a swap
//	url: https://docs.medusajs.com/modules/orders/swaps#canceling-a-swap
//
// parameters:
//   - (path) id=* {string} The ID of the Order the swap is associated with.
//   - (path) swap_id=* {string} The ID of the Swap.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order.
//
// x-codegen:
//
//	method: cancelSwap
//	params: AdminPostOrdersSwapCancelParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orders.cancelSwap(orderId, swapId)
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCancelSwap } from "medusa-react"
//
//     type Props = {
//     orderId: string,
//     swapId: string
//     }
//
//     const Swap = ({
//     orderId,
//     swapId
//     }: Props) => {
//     const cancelSwap = useAdminCancelSwap(
//     orderId
//     )
//     // ...
//
//     const handleCancel = () => {
//     cancelSwap.mutate(swapId, {
//     onSuccess: ({ order }) => {
//     console.log(order.swaps)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Swap
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/orders/{order_id}/swaps/{swap_id}/cancel' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrdersRes"
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
func (m *Order) CancelSwap(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	swapId, err := api.BindDelete(context, "swap_id")
	if err != nil {
		return err
	}

	swap, err := m.r.SwapService().SetContext(context.Context()).Retrieve(swapId, &sql.Options{})
	if err != nil {
		return err
	}

	if swap.OrderId.UUID != id {
		return utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintf("no swap was found with the id: %s related to order: %s", swapId, id),
		)
	}

	if _, err := m.r.SwapService().SetContext(context.Context()).Cancel(swapId); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/orders/{id}/claims/{claim_id}/cancel
// operationId: "PostOrdersClaimCancel"
// summary: "Cancel a Claim"
// description: "Cancel a Claim and change its status. A claim can't be canceled if it has a refund, if its fulfillments haven't been canceled, of if its associated return hasn't been canceled."
// x-authenticated: true
// externalDocs:
//
//	description: Canceling a claim
//	url: https://docs.medusajs.com/modules/orders/claims#cancel-a-claim
//
// parameters:
//   - (path) id=* {string} The ID of the order the claim is associated with.
//   - (path) claim_id=* {string} The ID of the Claim.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order.
//
// x-codegen:
//
//	method: cancelClaim
//	params: AdminPostOrdersClaimCancel
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orders.cancelClaim(orderId, claimId)
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCancelClaim } from "medusa-react"
//
//     type Props = {
//     orderId: string
//     claimId: string
//     }
//
//     const Claim = ({ orderId, claimId }: Props) => {
//     const cancelClaim = useAdminCancelClaim(orderId)
//     // ...
//
//     const handleCancel = () => {
//     cancelClaim.mutate(claimId)
//     }
//
//     // ...
//     }
//
//     export default Claim
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/orders/{id}/claims/{claim_id}/cancel' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrdersRes"
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
func (m *Order) CancelClaim(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	claimId, err := api.BindDelete(context, "claim_id")
	if err != nil {
		return err
	}

	claim, err := m.r.ClaimService().SetContext(context.Context()).Retrieve(claimId, &sql.Options{})
	if err != nil {
		return err
	}

	if claim.OrderId.UUID != id {
		return utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintf("no claim was found with the id: %s related to order: %s", claimId, id),
		)
	}

	if _, err := m.r.ClaimService().SetContext(context.Context()).Cancel(claimId); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/orders/{id}/claims/{claim_id}/fulfillments/{fulfillment_id}/cancel
// operationId: "PostOrdersClaimFulfillmentsCancel"
// summary: "Cancel Claim's Fulfillment"
// description: "Cancel a claim's fulfillment and change its fulfillment status to `canceled`."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the order the claim is associated with.
//   - (path) claim_id=* {string} The ID of the claim.
//   - (path) fulfillment_id=* {string} The ID of the fulfillment.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order.
//
// x-codegen:
//
//	method: cancelClaimFulfillment
//	params: AdminPostOrdersClaimFulfillmentsCancelParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orders.cancelClaimFulfillment(orderId, claimId, fulfillmentId)
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCancelClaimFulfillment } from "medusa-react"
//
//     type Props = {
//     orderId: string
//     claimId: string
//     }
//
//     const Claim = ({ orderId, claimId }: Props) => {
//     const cancelFulfillment = useAdminCancelClaimFulfillment(
//     orderId
//     )
//     // ...
//
//     const handleCancel = (fulfillmentId: string) => {
//     cancelFulfillment.mutate({
//     claim_id: claimId,
//     fulfillment_id: fulfillmentId,
//     }, {
//     onSuccess: ({ order }) => {
//     console.log(order.claims)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Claim
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/orders/{id}/claims/{claim_id}/fulfillments/{fulfillment_id}/cancel' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrdersRes"
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
func (m *Order) CancelFullfillmentClaim(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	claimId, err := api.BindDelete(context, "claim_id")
	if err != nil {
		return err
	}

	fulfillmentId, err := api.BindDelete(context, "fulfillment_id")
	if err != nil {
		return err
	}

	fulfillment, err := m.r.FulfillmentService().SetContext(context.Context()).Retrieve(fulfillmentId, &sql.Options{})
	if err != nil {
		return err
	}

	if fulfillment.ClaimOrderId.UUID != id {
		return utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintf("no fulfillment was found with the id: %s related to claim: %s", fulfillmentId, claimId),
		)
	}

	claim, err := m.r.ClaimService().SetContext(context.Context()).Retrieve(claimId, &sql.Options{})
	if err != nil {
		return err
	}

	if claim.OrderId.UUID != id {
		return utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintf("no claim was found with the id: %s related to order: %s", claimId, id),
		)
	}

	if _, err := m.r.ClaimService().SetContext(context.Context()).CancelFulfillment(fulfillmentId); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/orders/{id}/swaps/{swap_id}/fulfillments/{fulfillment_id}/cancel
// operationId: "PostOrdersSwapFulfillmentsCancel"
// summary: "Cancel Swap's Fulfilmment"
// description: "Cancel a swap's fulfillment and change its fulfillment status to `canceled`."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the order the swap is associated with.
//   - (path) swap_id=* {string} The ID of the swap.
//   - (path) fulfillment_id=* {string} The ID of the fulfillment.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order.
//
// x-codegen:
//
//	method: cancelSwapFulfillment
//	params: AdminPostOrdersSwapFulfillementsCancelParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orders.cancelSwapFulfillment(orderId, swapId, fulfillmentId)
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCancelSwapFulfillment } from "medusa-react"
//
//     type Props = {
//     orderId: string,
//     swapId: string
//     }
//
//     const Swap = ({
//     orderId,
//     swapId
//     }: Props) => {
//     const cancelFulfillment = useAdminCancelSwapFulfillment(
//     orderId
//     )
//     // ...
//
//     const handleCancelFulfillment = (
//     fulfillmentId: string
//     ) => {
//     cancelFulfillment.mutate({
//     swap_id: swapId,
//     fulfillment_id: fulfillmentId,
//     })
//     }
//
//     // ...
//     }
//
//     export default Swap
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/orders/{id}/swaps/{swap_id}/fulfillments/{fulfillment_id}/cancel' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrdersRes"
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
func (m *Order) CancelFullfillmentSwap(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	swapId, err := api.BindDelete(context, "swap_id")
	if err != nil {
		return err
	}

	fulfillmentId, err := api.BindDelete(context, "fulfillment_id")
	if err != nil {
		return err
	}

	fulfillment, err := m.r.FulfillmentService().SetContext(context.Context()).Retrieve(fulfillmentId, &sql.Options{})
	if err != nil {
		return err
	}

	if fulfillment.SwapId.UUID != swapId {
		return utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintf("no fulfillment was found with the id: %s related to swap: %s", fulfillmentId, swapId),
		)
	}

	swap, err := m.r.SwapService().SetContext(context.Context()).Retrieve(swapId, &sql.Options{})
	if err != nil {
		return err
	}

	if swap.OrderId.UUID != id {
		return utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintf("no swap was found with the id: %s related to order: %s", swapId, id),
		)
	}

	if _, err := m.r.SwapService().SetContext(context.Context()).CancelFulfillment(fulfillmentId); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/orders/{id}/fulfillments/{fulfillment_id}/cancel
// operationId: "PostOrdersOrderFulfillmentsCancel"
// summary: "Cancel a Fulfilmment"
// description: "Cancel an order's fulfillment and change its fulfillment status to `canceled`."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Order.
//   - (path) fulfillment_id=* {string} The ID of the Fulfillment.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order.
//
// x-codegen:
//
//	method: cancelFulfillment
//	params: AdminPostOrdersOrderFulfillementsCancelParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orders.cancelFulfillment(orderId, fulfillmentId)
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCancelFulfillment } from "medusa-react"
//
//     type Props = {
//     orderId: string
//     }
//
//     const Order = ({ orderId }: Props) => {
//     const cancelFulfillment = useAdminCancelFulfillment(
//     orderId
//     )
//     // ...
//
//     const handleCancel = (
//     fulfillmentId: string
//     ) => {
//     cancelFulfillment.mutate(fulfillmentId, {
//     onSuccess: ({ order }) => {
//     console.log(order.fulfillments)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Order
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/orders/{id}/fulfillments/{fulfillment_id}/cancel' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrdersRes"
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
func (m *Order) CancelFullfillment(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	fulfillmentId, err := api.BindDelete(context, "fulfillment_id")
	if err != nil {
		return err
	}

	fulfillment, err := m.r.FulfillmentService().SetContext(context.Context()).Retrieve(fulfillmentId, &sql.Options{})
	if err != nil {
		return err
	}

	if fulfillment.OrderId.UUID != id {
		return utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintf("no fulfillment was found with the id: %s related to order: %s", fulfillmentId, id),
		)
	}

	if _, err := m.r.OrderService().SetContext(context.Context()).CancelFulfillment(fulfillmentId); err != nil {
		return err
	}

	data, err := m.r.FulfillmentService().SetContext(context.Context()).Retrieve(fulfillmentId, &sql.Options{Relations: []string{"items", "items.item"}})
	if err != nil {
		return err
	}

	if data.LocationId.UUID != uuid.Nil && m.r.InventoryService() != nil {
		for _, item := range data.Items {
			if item.Item.VariantId.UUID != uuid.Nil {
				if err := m.r.ProductVariantInventoryService().SetContext(context.Context()).AdjustInventory(item.Item.VariantId.UUID, item.Fulfillment.LocationId.UUID, item.Quantity); err != nil {
					return err
				}
			}
		}
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/orders/{id}/capture
// operationId: "PostOrdersOrderCapture"
// summary: "Capture an Order's Payments"
// description: "Capture all the Payments associated with an Order. The payment of canceled orders can't be captured."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Order.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order.
//
// x-codegen:
//
//	method: capturePayment
//	params: AdminPostOrdersOrderCaptureParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orders.capturePayment(orderId)
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCapturePayment } from "medusa-react"
//
//     type Props = {
//     orderId: string
//     }
//
//     const Order = ({ orderId }: Props) => {
//     const capturePayment = useAdminCapturePayment(
//     orderId
//     )
//     // ...
//
//     const handleCapture = () => {
//     capturePayment.mutate(void 0, {
//     onSuccess: ({ order }) => {
//     console.log(order.status)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Order
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/orders/{id}/capture' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrdersRes"
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
func (m *Order) CapturePayment(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	if _, err := m.r.OrderService().SetContext(context.Context()).CapturePayment(id); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/orders/{id}/complete
// operationId: "PostOrdersOrderComplete"
// summary: "Complete an Order"
// description: "Complete an Order and change its status. A canceled order can't be completed."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Order.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order.
//
// x-codegen:
//
//	method: complete
//	params: AdminPostOrdersOrderCompleteParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orders.complete(orderId)
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCompleteOrder } from "medusa-react"
//
//     type Props = {
//     orderId: string
//     }
//
//     const Order = ({ orderId }: Props) => {
//     const completeOrder = useAdminCompleteOrder(
//     orderId
//     )
//     // ...
//
//     const handleComplete = () => {
//     completeOrder.mutate(void 0, {
//     onSuccess: ({ order }) => {
//     console.log(order.status)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Order
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/orders/{id}/complete' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrdersRes"
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
func (m *Order) Complete(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	if _, err := m.r.OrderService().SetContext(context.Context()).CompleteOrder(id); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/orders/{id}/claims/{claim_id}/shipments
// operationId: "PostOrdersOrderClaimsClaimShipments"
// summary: "Ship a Claim's Fulfillment"
// description: "Create a shipment for the claim and mark its fulfillment as shipped. This changes the claim's fulfillment status to either `partially_shipped` or `shipped`, depending on
//
//	whether all the items were shipped."
//
// x-authenticated: true
// externalDocs:
//
//	description: Fulfill a claim
//	url: https://docs.medusajs.com/modules/orders/claims#fulfill-a-claim
//
// parameters:
//   - (path) id=* {string} The ID of the Order the claim is associated with.
//   - (path) claim_id=* {string} The ID of the Claim.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostOrdersOrderClaimsClaimShipmentsReq"
//
// x-codegen:
//
//	method: createClaimShipment
//	params: AdminPostOrdersOrderClaimsClaimShipmentsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orders.createClaimShipment(orderId, claimId, {
//     fulfillment_id
//     })
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreateClaimShipment } from "medusa-react"
//
//     type Props = {
//     orderId: string
//     claimId: string
//     }
//
//     const Claim = ({ orderId, claimId }: Props) => {
//     const createShipment = useAdminCreateClaimShipment(orderId)
//     // ...
//
//     const handleCreateShipment = (fulfillmentId: string) => {
//     createShipment.mutate({
//     claim_id: claimId,
//     fulfillment_id: fulfillmentId,
//     }, {
//     onSuccess: ({ order }) => {
//     console.log(order.claims)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Claim
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/orders/{id}/claims/{claim_id}/shipments' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "fulfillment_id": "{fulfillment_id}"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrdersRes"
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
func (m *Order) CreateClaimShippment(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.OrderClaimShipments](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	claimId, err := api.BindDelete(context, "claim_id")
	if err != nil {
		return err
	}

	trakingLinks := []models.TrackingLink{}
	for _, link := range model.TrackingNumbers {
		trakingLinks = append(trakingLinks, models.TrackingLink{TrackingNumber: link})
	}

	if _, err := m.r.ClaimService().SetContext(context.Context()).CreateShipment(claimId, model.FulfillmentId, trakingLinks, false, nil); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

//
// @oas:path [post] /admin/orders/{id}/claims
// operationId: "PostOrdersOrderClaims"
// summary: "Create a Claim"
// description: "Create a Claim for an order. If a return shipping method is specified, a return will also be created and associated with the claim. If the claim's type is `refund`,
//  the refund is processed as well."
// externalDocs:
//   description: How are claims created
//   url: https://docs.medusajs.com/modules/orders/claims#how-are-claims-created
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Order.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order.
// requestBody:
//   content:
//     application/json:
//       schema:
//         $ref: "#/components/schemas/AdminPostOrdersOrderClaimsReq"
// x-codegen:
//   method: createClaim
//   params: AdminPostOrdersOrderClaimsParams
// x-codeSamples:
//   - lang: JavaScript
//     label: JS Client
//     source: |
//       import Medusa from "@medusajs/medusa-js"
//       const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//       // must be previously logged in or use api token
//       medusa.admin.orders.createClaim(orderId, {
//         type: 'refund',
//         claim_items: [
//           {
//             item_id,
//             quantity: 1
//           }
//         ]
//       })
//       .then(({ order }) => {
//         console.log(order.id);
//       })
//   - lang: tsx
//     label: Medusa React
//     source: |
//       import React from "react"
//       import { useAdminCreateClaim } from "medusa-react"
//
//       type Props = {
//         orderId: string
//       }
//

//	    const CreateClaim = ({ orderId }: Props) => {
//
//	    const CreateClaim = (orderId: string) => {
//	      const createClaim = useAdminCreateClaim(orderId)
//	      // ...
//
//	      const handleCreate = (itemId: string) => {
//	        createClaim.mutate({
//	          type: "refund",
//	          claim_items: [
//	            {
//	              item_id: itemId,
//	              quantity: 1,
//	            },
//	          ],
//	        }, {
//	          onSuccess: ({ order }) => {
//	            console.log(order.claims)
//	          }
//	        })
//	      }
//
//	      // ...
//	    }
//
//	    export default CreateClaim
//	- lang: Shell
//	  label: cURL
//	  source: |
//	    curl -X POST '"{backend_url}"/admin/orders/{id}/claims' \
//	    -H 'x-medusa-access-token: "{api_token}"' \
//	    -H 'Content-Type: application/json' \
//	    --data-raw '{
//	        "type": "refund",
//	        "claim_items": [
//	          {
//	            "item_id": "asdsd",
//	            "quantity": 1
//	          }
//	        ]
//	    }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrdersRes"
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
func (m *Order) CreateClaim(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.CreateClaimInput](context, "id", m.r.Validator())
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
			order, err := m.r.OrderService().SetContext(context.Context()).RetrieveById(id, &sql.Options{
				Relations: []string{
					"customer",
					"shipping_address",
					"region",
					"items",
					"items.tax_lines",
					"discounts",
					"discounts.rule",
					"claims",
					"claims.additional_items",
					"claims.additional_items.tax_lines",
					"swaps",
					"swaps.additional_items",
					"swaps.additional_items.tax_lines",
				},
			})
			if err != nil {
				return nil, err
			}

			model.Order = order
			// model.IdempotencyKey = idempotencyKey.IdempotencyKey

			if _, err := m.r.ClaimService().SetContext(context.Context()).Create(model); err != nil {
				return nil, err
			}

			return &types.IdempotencyCallbackResult{RecoveryPoint: "claim_created"}, nil
		})
		if err != nil {
			return err
		}
	} else if idempotencyKey.RecoveryPoint == "claim_created" {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).WorkStage(idempotencyKey.IdempotencyKey, func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
			claims, err := m.r.ClaimService().SetContext(context.Context()).List(&models.ClaimOrder{IdempotencyKey: idempotencyKey.IdempotencyKey}, &sql.Options{})
			if err != nil {
				return nil, err
			}

			claim := claims[0]

			if claim.Type == "refund" {
				if _, err := m.r.ClaimService().SetContext(context.Context()).ProcessRefund(claim.Id); err != nil {
					return nil, err
				}
			}

			return &types.IdempotencyCallbackResult{RecoveryPoint: "refund_handled"}, nil
		})
		if err != nil {
			return err
		}
	} else if idempotencyKey.RecoveryPoint == "refund_handled" {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).WorkStage(idempotencyKey.IdempotencyKey, func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
			claims, err := m.r.ClaimService().SetContext(context.Context()).List(&models.ClaimOrder{IdempotencyKey: idempotencyKey.IdempotencyKey}, &sql.Options{Relations: []string{"return_order"}})
			if err != nil {
				return nil, err
			}

			claim := claims[0]

			if !reflect.ValueOf(claim.ReturnOrder).IsZero() {
				if _, err := m.r.ReturnService().SetContext(context.Context()).Fulfill(claim.ReturnOrder.Id); err != nil {
					return nil, err
				}
			}

			result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
				ReturnableItems: includes,
			})
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

// @oas:path [post] /admin/orders/{id}/fulfillment
// operationId: "PostOrdersOrderFulfillments"
// summary: "Create a Fulfillment"
// description: "Create a Fulfillment of an Order using the fulfillment provider, and change the order's fulfillment status to either `partially_fulfilled` or `fulfilled`, depending on
//
//	whether all the items were fulfilled."
//
// x-authenticated: true
// externalDocs:
//
//	description: Fulfillments of orders
//	url: https://docs.medusajs.com/modules/orders/#fulfillments-in-orders
//
// parameters:
//   - (path) id=* {string} The ID of the Order.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostOrdersOrderFulfillmentsReq"
//
// x-codegen:
//
//	method: createFulfillment
//	params: AdminPostOrdersOrderFulfillmentsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orders.createFulfillment(orderId, {
//     items: [
//     {
//     item_id,
//     quantity: 1
//     }
//     ]
//     })
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreateFulfillment } from "medusa-react"
//
//     type Props = {
//     orderId: string
//     }
//
//     const Order = ({ orderId }: Props) => {
//     const createFulfillment = useAdminCreateFulfillment(
//     orderId
//     )
//     // ...
//
//     const handleCreateFulfillment = (
//     itemId: string,
//     quantity: number
//     ) => {
//     createFulfillment.mutate({
//     items: [
//     {
//     item_id: itemId,
//     quantity,
//     },
//     ],
//     }, {
//     onSuccess: ({ order }) => {
//     console.log(order.fulfillments)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Order
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/orders/{id}/fulfillment' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "items": [
//     {
//     "item_id": "{item_id}",
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
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrdersRes"
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
func (m *Order) CreateFulfillment(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.OrderFulfillments](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	fulfillments, err := m.r.OrderService().SetContext(context.Context()).RetrieveById(id, &sql.Options{Relations: []string{"fulfillments"}})
	if err != nil {
		return err
	}

	var fulfillmentIds uuid.UUIDs
	for _, fulfillment := range fulfillments.Fulfillments {
		fulfillmentIds = append(fulfillmentIds, fulfillment.Id)
	}

	if _, err := m.r.OrderService().SetContext(context.Context()).CreateFulfillment(id, model.Items, map[string]interface{}{
		"metadata":        model.Metadata,
		"no_notification": model.NoNotification,
		"location_id":     model.LocationId,
	}); err != nil {
		return err
	}

	if model.LocationId != uuid.Nil {
		order, err := m.r.OrderService().SetContext(context.Context()).RetrieveById(id, &sql.Options{Relations: []string{"fulfillments", "fulfillments.items", "fulfillments.items.item"}})
		if err != nil {
			return err
		}

		data := lo.Filter[models.Fulfillment](order.Fulfillments, func(item models.Fulfillment, index int) bool {
			for _, fulfillmentId := range fulfillmentIds {
				return item.Id == fulfillmentId
			}

			return false
		})

		for _, fulfillment := range data {
			items := []models.LineItem{}
			for _, item := range fulfillment.Items {
				lineItem := item.Item
				lineItem.Quantity = item.Quantity
				items = append(items, *lineItem)
			}
			if err := m.r.ProductVariantInventoryService().ValidateInventoryAtLocation(items, model.LocationId); err != nil {
				return err
			}

			for _, item := range fulfillment.Items {
				if item.Item.VariantId.UUID != uuid.Nil {
					break
				}

				if err := m.r.ProductVariantInventoryService().AdjustReservationsQuantityByLineItem(item.Item.Id, item.Item.VariantId.UUID, model.LocationId, -item.Quantity); err != nil {
					return err
				}

				if err := m.r.ProductVariantInventoryService().AdjustInventory(item.Item.VariantId.UUID, model.LocationId, -item.Quantity); err != nil {
					return err
				}
			}
		}

	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/orders/{id}/line-items/{line_item_id}/reserve
// operationId: "PostOrdersOrderLineItemReservations"
// summary: "Create a Reservation"
// description: "Create a Reservation for a line item at a specified location, optionally for a partial quantity."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Order.
//   - (path) line_item_id=* {string} The ID of the Line item.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminOrdersOrderLineItemReservationReq"
//
// x-codeSamples:
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/orders/{id}/line-items/{line_item_id}/reserve' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "location_id": "loc_1"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPostReservationsReq"
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
func (m *Order) CreateReservationForLineItem(context fiber.Ctx) error {
	model, err := api.Bind[types.OrderLineItemReservation](context, m.r.Validator())
	if err != nil {
		return err
	}

	lineItemId, err := api.BindDelete(context, "line_item_id")
	if err != nil {
		return err
	}

	lineItem, err := m.r.LineItemService().SetContext(context.Context()).Retrieve(lineItemId, &sql.Options{})
	if err != nil {
		return err
	}

	if lineItem.VariantId.UUID == uuid.Nil {
		return utils.NewApplictaionError(
			utils.NOT_FOUND,
			`Can't create a reservation for a Line Item wihtout a variant`,
		)
	}

	quantity := 0
	if !reflect.ValueOf(model.Quantity).IsZero() {
		quantity = model.Quantity
	} else {
		quantity = lineItem.Quantity
	}

	reservations, err := m.r.ProductVariantInventoryService().SetContext(context.Context()).ReserveQuantity(lineItem.VariantId.UUID, quantity, services.ReserveQuantityContext{LocationId: model.LocationId})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(reservations[0])
}

// @oas:path [post] /admin/orders/{id}/shipment
// operationId: "PostOrdersOrderShipment"
// summary: "Ship a Fulfillment"
// description: "Create a shipment and mark a fulfillment as shipped. This changes the order's fulfillment status to either `partially_shipped` or `shipped`, depending on
//
//	whether all the items were shipped."
//
// x-authenticated: true
// externalDocs:
//
//	description: Fulfillments of orders
//	url: https://docs.medusajs.com/modules/orders/#fulfillments-in-orders
//
// parameters:
//   - (path) id=* {string} The ID of the Order.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostOrdersOrderShipmentReq"
//
// x-codegen:
//
//	method: createShipment
//	params: AdminPostOrdersOrderShipmentParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orders.createShipment(order_id, {
//     fulfillment_id
//     })
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreateShipment } from "medusa-react"
//
//     type Props = {
//     orderId: string
//     }
//
//     const Order = ({ orderId }: Props) => {
//     const createShipment = useAdminCreateShipment(
//     orderId
//     )
//     // ...
//
//     const handleCreate = (
//     fulfillmentId: string
//     ) => {
//     createShipment.mutate({
//     fulfillment_id: fulfillmentId,
//     }, {
//     onSuccess: ({ order }) => {
//     console.log(order.fulfillment_status)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Order
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/orders/{id}/shipment' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "fulfillment_id": "{fulfillment_id}"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrdersRes"
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
func (m *Order) CreateShipment(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.CreateOrderShipment](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	trakingLinks := []models.TrackingLink{}
	for _, link := range model.TrackingNumbers {
		trakingLinks = append(trakingLinks, models.TrackingLink{TrackingNumber: link})
	}

	if _, err := m.r.OrderService().SetContext(context.Context()).CreateShipment(id, model.FulfillmentId, trakingLinks, struct {
		NoNotification bool
		Metadata       map[string]interface{}
	}{NoNotification: model.NoNotification}); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/orders/{id}/swaps/{swap_id}/shipments
// operationId: "PostOrdersOrderSwapsSwapShipments"
// summary: "Ship a Swap's Fulfillment"
// description: "Create a shipment for a swap and mark its fulfillment as shipped. This changes the swap's fulfillment status to either `partially_shipped` or `shipped`, depending on
//
//	whether all the items were shipped."
//
// x-authenticated: true
// externalDocs:
//
//	description: Handling swap fulfillments
//	url: https://docs.medusajs.com/modules/orders/swaps#handling-swap-fulfillment
//
// parameters:
//   - (path) id=* {string} The ID of the Order.
//   - (path) swap_id=* {string} The ID of the Swap.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostOrdersOrderSwapsSwapShipmentsReq"
//
// x-codegen:
//
//	method: createSwapShipment
//	params: AdminPostOrdersOrderSwapsSwapShipmentsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orders.createSwapShipment(orderId, swapId, {
//     fulfillment_id
//     })
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreateSwapShipment } from "medusa-react"
//
//     type Props = {
//     orderId: string,
//     swapId: string
//     }
//
//     const Swap = ({
//     orderId,
//     swapId
//     }: Props) => {
//     const createShipment = useAdminCreateSwapShipment(
//     orderId
//     )
//     // ...
//
//     const handleCreateShipment = (
//     fulfillmentId: string
//     ) => {
//     createShipment.mutate({
//     swap_id: swapId,
//     fulfillment_id: fulfillmentId,
//     }, {
//     onSuccess: ({ order }) => {
//     console.log(order.swaps)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Swap
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/orders/{id}/swaps/{swap_id}/shipments' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "fulfillment_id": "{fulfillment_id}"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrdersRes"
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
func (m *Order) CreateSwapShipment(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.CreateOrderShipment](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	swapId, err := api.BindDelete(context, "swap_id")
	if err != nil {
		return err
	}

	trakingLinks := []models.TrackingLink{}
	for _, link := range model.TrackingNumbers {
		trakingLinks = append(trakingLinks, models.TrackingLink{TrackingNumber: link})
	}

	if _, err := m.r.SwapService().SetContext(context.Context()).CreateShipment(swapId, model.FulfillmentId, trakingLinks, &types.CreateShipmentConfig{NoNotification: model.NoNotification}); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/orders/{id}/swaps
// operationId: "PostOrdersOrderSwaps"
// summary: "Create a Swap"
// description: "Create a Swap. This includes creating a return that is associated with the swap."
// x-authenticated: true
// externalDocs:
//
//	description: How are swaps created
//	url: https://docs.medusajs.com/modules/orders/swaps#how-are-swaps-created
//
// parameters:
//   - (path) id=* {string} The ID of the Order.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostOrdersOrderSwapsReq"
//
// x-codegen:
//
//	method: createSwap
//	queryParams: AdminPostOrdersOrderSwapsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orders.createSwap(orderId, {
//     return_items: [
//     {
//     item_id,
//     quantity: 1
//     }
//     ]
//     })
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreateSwap } from "medusa-react"
//
//     type Props = {
//     orderId: string
//     }
//
//     const CreateSwap = ({ orderId }: Props) => {
//     const createSwap = useAdminCreateSwap(orderId)
//     // ...
//
//     const handleCreate = (
//     returnItems: {
//     item_id: string,
//     quantity: number
//     }[]
//     ) => {
//     createSwap.mutate({
//     return_items: returnItems
//     }, {
//     onSuccess: ({ order }) => {
//     console.log(order.swaps)
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
//     curl -X POST '"{backend_url}"/admin/orders/{id}/swaps' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "return_items": [
//     {
//     "item_id": "asfasf",
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
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrdersRes"
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
func (m *Order) CreateSwap(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.OrderSwap](context, "id", m.r.Validator())
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
			order, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, &sql.Options{
				Relations: []string{
					"cart",
					"items",
					"items.variant",
					"items.tax_lines",
					"swaps",
					"swaps.additional_items",
					"swaps.additional_items.variant",
					"swaps.additional_items.tax_lines",
				},
			}, types.TotalsContext{})
			if err != nil {
				return nil, err
			}

			swap, err := m.r.SwapService().SetContext(context.Context()).Create(order, model.ReturnItems, model.AdditionalItems, &model.ReturnShipping, map[string]interface{}{
				// "idempotency_key": idempotencyKey.idempotency_key,
				"no_notification": model.NoNotification,
				"allow_backorder": model.AllowBackorder,
				"location_id":     model.ReturnLocationId,
			})
			if err != nil {
				return nil, err
			}

			_, err = m.r.SwapService().SetContext(context.Context()).CreateCart(swap.Id, model.CustomShippingOptions, map[string]interface{}{
				"sales_channel_id": model.SalesChannelId,
			})
			if err != nil {
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

			result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
				ReturnableItems: includes,
			})
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

// @oas:path [post] /admin/orders/{id}/claims/{claim_id}/fulfillments
// operationId: "PostOrdersOrderClaimsClaimFulfillments"
// summary: "Create a Claim Fulfillment"
// description: "Create a Fulfillment for a Claim, and change its fulfillment status to `partially_fulfilled` or `fulfilled` depending on whether all the items were fulfilled.
// It may also change the status to `requires_action` if any actions are required."
// x-authenticated: true
// externalDocs:
//
//	description: Fulfill a claim
//	url: https://docs.medusajs.com/modules/orders/claims#fulfill-a-claim
//
// parameters:
//   - (path) id=* {string} The ID of the Order the claim is associated with.
//   - (path) claim_id=* {string} The ID of the Claim.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostOrdersOrderClaimsClaimFulfillmentsReq"
//
// x-codegen:
//
//	method: fulfillClaim
//	params: AdminPostOrdersOrderClaimsClaimFulfillmentsReq
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orders.fulfillClaim(orderId, claimId, {
//     })
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminFulfillClaim } from "medusa-react"
//
//     type Props = {
//     orderId: string
//     claimId: string
//     }
//
//     const Claim = ({ orderId, claimId }: Props) => {
//     const fulfillClaim = useAdminFulfillClaim(orderId)
//     // ...
//
//     const handleFulfill = () => {
//     fulfillClaim.mutate({
//     claim_id: claimId,
//     }, {
//     onSuccess: ({ order }) => {
//     console.log(order.claims)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Claim
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/orders/{id}/claims/{claim_id}/fulfillments' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrdersRes"
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
func (m *Order) FulfillClaim(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.OrderClaimFulfillments](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	claimId, err := api.BindDelete(context, "claim_id")
	if err != nil {
		return err
	}

	fulfillments, err := m.r.ClaimService().SetContext(context.Context()).Retrieve(claimId, &sql.Options{Relations: []string{"fulfillments"}})
	if err != nil {
		return err
	}

	var fulfillmentIds uuid.UUIDs
	for _, fulfillment := range fulfillments.Fulfillments {
		fulfillmentIds = append(fulfillmentIds, fulfillment.Id)
	}

	if _, err := m.r.ClaimService().SetContext(context.Context()).CreateFulfillment(claimId, model.NoNotification, model.LocationId, model.Metadata); err != nil {
		return err
	}

	if model.LocationId != uuid.Nil {
		order, err := m.r.ClaimService().SetContext(context.Context()).Retrieve(claimId, &sql.Options{Relations: []string{"fulfillments", "fulfillments.items", "fulfillments.items.item"}})
		if err != nil {
			return err
		}

		data := lo.Filter[models.Fulfillment](order.Fulfillments, func(item models.Fulfillment, index int) bool {
			for _, fulfillmentId := range fulfillmentIds {
				return item.Id == fulfillmentId
			}

			return false
		})

		for _, fulfillment := range data {
			items := []models.LineItem{}
			for _, item := range fulfillment.Items {
				lineItem := item.Item
				lineItem.Quantity = item.Quantity
				items = append(items, *lineItem)
			}
			if err := m.r.ProductVariantInventoryService().ValidateInventoryAtLocation(items, model.LocationId); err != nil {
				return err
			}

			for _, item := range fulfillment.Items {
				if item.Item.VariantId.UUID != uuid.Nil {
					break
				}

				if err := m.r.ProductVariantInventoryService().AdjustReservationsQuantityByLineItem(item.Item.Id, item.Item.VariantId.UUID, model.LocationId, -item.Quantity); err != nil {
					return err
				}

				if err := m.r.ProductVariantInventoryService().AdjustInventory(item.Item.VariantId.UUID, model.LocationId, -item.Quantity); err != nil {
					return err
				}
			}
		}

	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

//
// @oas:path [post] /admin/orders/{id}/swaps/{swap_id}/fulfillments
// operationId: "PostOrdersOrderSwapsSwapFulfillments"
// summary: "Create a Swap Fulfillment"
// description: "Create a Fulfillment for a Swap and change its fulfillment status to `fulfilled`. If it requires any additional actions,
// its fulfillment status may change to `requires_action`."
// x-authenticated: true
// externalDocs:
//   description: Handling a swap's fulfillment
//   url: https://docs.medusajs.com/modules/orders/swaps#handling-swap-fulfillment
// parameters:
//   - (path) id=* {string} The ID of the Order the swap is associated with.
//   - (path) swap_id=* {string} The ID of the Swap.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order.
// requestBody:
//   content:
//     application/json:
//       schema:
//         $ref: "#/components/schemas/AdminPostOrdersOrderSwapsSwapFulfillmentsReq"
// x-codegen:
//   method: fulfillSwap
//   params: AdminPostOrdersOrderSwapsSwapFulfillmentsParams
// x-codeSamples:
//   - lang: JavaScript
//     label: JS Client
//     source: |
//       import Medusa from "@medusajs/medusa-js"
//       const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//       // must be previously logged in or use api token

//	    medusa.admin.orders.fulfillSwap(orderId, swapId, {
//
//	    })
//	    .then(({ order }) => {
//	      console.log(order.id);
//	    })
//	- lang: tsx
//	  label: Medusa React
//	  source: |
//	    import React from "react"
//	    import { useAdminFulfillSwap } from "medusa-react"
//
//	    type Props = {
//	      orderId: string,
//	      swapId: string
//	    }
//
//	    const Swap = ({
//	      orderId,
//	      swapId
//	    }: Props) => {
//	      const fulfillSwap = useAdminFulfillSwap(
//	        orderId
//	      )
//	      // ...
//
//	      const handleFulfill = () => {
//	        fulfillSwap.mutate({
//	          swap_id: swapId,
//	        }, {
//	          onSuccess: ({ order }) => {
//	            console.log(order.swaps)
//	          }
//	        })
//	      }
//
//	      // ...
//	    }
//
//	    export default Swap
//	- lang: Shell
//	  label: cURL
//	  source: |
//	    curl -X POST '"{backend_url}"/admin/orders/{id}/swaps/{swap_id}/fulfillments' \
//	    -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrdersRes"
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
func (m *Order) FulfillSwap(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.OrderClaimFulfillments](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	swapId, err := api.BindDelete(context, "swap_id")
	if err != nil {
		return err
	}

	fulfillments, err := m.r.SwapService().SetContext(context.Context()).Retrieve(swapId, &sql.Options{Relations: []string{"fulfillments"}})
	if err != nil {
		return err
	}

	var fulfillmentIds uuid.UUIDs
	for _, fulfillment := range fulfillments.Fulfillments {
		fulfillmentIds = append(fulfillmentIds, fulfillment.Id)
	}

	if _, err := m.r.SwapService().SetContext(context.Context()).CreateFulfillment(swapId, &types.CreateShipmentConfig{
		NoNotification: model.NoNotification,
		LocationId:     model.LocationId,
		Metadata:       model.Metadata,
	}); err != nil {
		return err
	}

	if model.LocationId != uuid.Nil {
		order, err := m.r.SwapService().SetContext(context.Context()).Retrieve(swapId, &sql.Options{Relations: []string{"fulfillments", "fulfillments.items", "fulfillments.items.item"}})
		if err != nil {
			return err
		}

		data := lo.Filter[models.Fulfillment](order.Fulfillments, func(item models.Fulfillment, index int) bool {
			for _, fulfillmentId := range fulfillmentIds {
				return item.Id == fulfillmentId
			}

			return false
		})

		for _, fulfillment := range data {
			items := []models.LineItem{}
			for _, item := range fulfillment.Items {
				lineItem := item.Item
				lineItem.Quantity = item.Quantity
				items = append(items, *lineItem)
			}
			if err := m.r.ProductVariantInventoryService().ValidateInventoryAtLocation(items, model.LocationId); err != nil {
				return err
			}

			for _, item := range fulfillment.Items {
				if item.Item.VariantId.UUID != uuid.Nil {
					break
				}

				if err := m.r.ProductVariantInventoryService().AdjustReservationsQuantityByLineItem(item.Item.Id, item.Item.VariantId.UUID, model.LocationId, -item.Quantity); err != nil {
					return err
				}

				if err := m.r.ProductVariantInventoryService().AdjustInventory(item.Item.VariantId.UUID, model.LocationId, -item.Quantity); err != nil {
					return err
				}
			}
		}

	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [get] /admin/orders/{id}/reservations
// operationId: "GetOrdersOrderReservations"
// summary: "Get Order Reservations"
// description: "Retrieve the list of reservations of an Order"
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Order.
//   - (query) offset=0 {integer} The number of reservations to skip when retrieving the reservations.
//   - (query) limit=20 {integer} Limit the number of reservations returned.
//
// x-codeSamples:
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/orders/{id}/reservations' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminReservationsListRes"
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
func (m *Order) GetReservations(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	order, err := m.r.OrderService().SetContext(context.Context()).RetrieveById(id, &sql.Options{Relations: []string{"items"}})
	if err != nil {
		return err
	}

	var itemIds uuid.UUIDs
	for _, item := range order.Items {
		itemIds = append(itemIds, item.Id)
	}

	result, count, err := m.r.InventoryService().ListReservationItems(context.Context(), interfaces.FilterableReservationItemProps{LineItemId: itemIds}, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   result,
		"count":  count,
		"offset": config.Skip,
		"limit":  config.Take,
	})
}

// @oas:path [post] /admin/orders/{id}/swaps/{swap_id}/process-payment
// operationId: "PostOrdersOrderSwapsSwapProcessPayment"
// summary: "Process a Swap Payment"
// description: "Process a swap's payment either by refunding or issuing a payment. This depends on the `difference_due` of the swap. If `difference_due` is negative, the amount is refunded.
//
//	If `difference_due` is positive, the amount is captured."
//
// x-authenticated: true
// externalDocs:
//
//	description: Handling a swap's payment
//	url: https://docs.medusajs.com/modules/orders/swaps#handling-swap-payment
//
// parameters:
//   - (path) id=* {string} The ID of the order the swap is associated with.
//   - (path) swap_id=* {string} The ID of the swap.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order.
//
// x-codegen:
//
//	method: processSwapPayment
//	params: AdminPostOrdersOrderSwapsSwapProcessPaymentParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orders.processSwapPayment(orderId, swapId)
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminProcessSwapPayment } from "medusa-react"
//
//     type Props = {
//     orderId: string,
//     swapId: string
//     }
//
//     const Swap = ({
//     orderId,
//     swapId
//     }: Props) => {
//     const processPayment = useAdminProcessSwapPayment(
//     orderId
//     )
//     // ...
//
//     const handleProcessPayment = () => {
//     processPayment.mutate(swapId, {
//     onSuccess: ({ order }) => {
//     console.log(order.swaps)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Swap
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/orders/{id}/swaps/{swap_id}/process-payment' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrdersRes"
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
func (m *Order) ProcessSwapPayment(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	swapId, err := api.BindDelete(context, "swap_id")
	if err != nil {
		return err
	}

	if _, err := m.r.SwapService().ProcessDifference(swapId); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/orders/{id}/refund
// operationId: "PostOrdersOrderRefunds"
// summary: "Create a Refund"
// description: "Refund an amount for an order. The amount must be less than or equal the `refundable_amount` of the order."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Order.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostOrdersOrderRefundsReq"
//
// x-codegen:
//
//	method: refundPayment
//	params: AdminPostOrdersOrderRefundsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orders.refundPayment(orderId, {
//     amount: 1000,
//     reason: "Do not like it"
//     })
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminRefundPayment } from "medusa-react"
//
//     type Props = {
//     orderId: string
//     }
//
//     const Order = ({ orderId }: Props) => {
//     const refundPayment = useAdminRefundPayment(
//     orderId
//     )
//     // ...
//
//     const handleRefund = (
//     amount: number,
//     reason: string
//     ) => {
//     refundPayment.mutate({
//     amount,
//     reason,
//     }, {
//     onSuccess: ({ order }) => {
//     console.log(order.refunds)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Order
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/orders/adasda/refund' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "amount": 1000,
//     "reason": "Do not like it"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrdersRes"
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
func (m *Order) RefundPayment(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.OrderRefunds](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if _, err := m.r.OrderService().SetContext(context.Context()).CreateRefund(id, model.Amount, model.Reason, &model.Note, &model.NoNotification); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/orders/{id}/return
// operationId: "PostOrdersOrderReturns"
// summary: "Request a Return"
// description: "Request and create a Return for items in an order. If the return shipping method is specified, it will be automatically fulfilled."
// x-authenticated: true
// externalDocs:
//
//	description: Return creation process
//	url: https://docs.medusajs.com/modules/orders/returns#returns-process
//
// parameters:
//   - (path) id=* {string} The ID of the Order.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostOrdersOrderReturnsReq"
//
// x-codegen:
//
//	method: requestReturn
//	params: AdminPostOrdersOrderReturnsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orders.requestReturn(orderId, {
//     items: [
//     {
//     item_id,
//     quantity: 1
//     }
//     ]
//     })
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminRequestReturn } from "medusa-react"
//
//     type Props = {
//     orderId: string
//     }
//
//     const Order = ({ orderId }: Props) => {
//     const requestReturn = useAdminRequestReturn(
//     orderId
//     )
//     // ...
//
//     const handleRequestingReturn = (
//     itemId: string,
//     quantity: number
//     ) => {
//     requestReturn.mutate({
//     items: [
//     {
//     item_id: itemId,
//     quantity
//     }
//     ]
//     }, {
//     onSuccess: ({ order }) => {
//     console.log(order.returns)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Order
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/orders/{id}/return' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "items": [
//     {
//     "item_id": "{item_id}",
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
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrdersRes"
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
func (m *Order) RequestReturn(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.OrderReturns](context, "id", m.r.Validator())
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
				OrderId: id,
				Items:   model.Items,
			}

			if !reflect.ValueOf(model.LocationId).IsZero() && (m.r.InventoryService() != nil || m.r.StockLocationService() != nil) {
				data.LocationId = model.LocationId
			}

			if !reflect.ValueOf(model.ReturnShipping).IsZero() {
				data.ShippingMethod = &model.ReturnShipping
			}

			if !reflect.ValueOf(model.Refund).IsZero() && model.Refund < 0 {
				data.RefundAmount = 0
			} else {
				if !reflect.ValueOf(model.Refund).IsZero() && model.Refund >= 0 {
					data.RefundAmount = model.Refund
				}
			}

			evaluatedNoNotification := model.NoNotification

			if !reflect.ValueOf(model.NoNotification).IsZero() {
				order, err := m.r.OrderService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
				if err != nil {
					return nil, err
				}

				evaluatedNoNotification = order.NoNotification
			}

			data.NoNotification = evaluatedNoNotification

			createdReturn, err := m.r.ReturnService().SetContext(context.Context()).Create(data)
			if err != nil {
				return nil, err
			}

			if !reflect.ValueOf(model.ReturnShipping).IsZero() {
				if _, err := m.r.ReturnService().SetContext(context.Context()).Fulfill(createdReturn.Id); err != nil {
					return nil, err
				}
			}

			// eventBus
			//             .withTransaction(manager)
			//             .emit(OrderService.Events.RETURN_REQUESTED, {
			//               id,
			//               return_id: createdReturn.id,
			//               no_notification: evaluatedNoNotification,
			//             })

			return &types.IdempotencyCallbackResult{RecoveryPoint: "return_requested"}, nil
		})
		if err != nil {
			return err
		}
	} else if idempotencyKey.RecoveryPoint == "return_requested" {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).WorkStage(idempotencyKey.IdempotencyKey, func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
			if model.ReceiveNow {
				returns, err := m.r.ReturnService().SetContext(context.Context()).List(&types.FilterableReturn{IdempotencyKey: idempotencyKey.IdempotencyKey}, &sql.Options{})
				if err != nil {
					return nil, err
				}

				returnOrder := returns[0]

				if _, err := m.r.ReturnService().SetContext(context.Context()).Receive(returnOrder.Id, model.Items, &model.Refund, false, map[string]interface{}{}); err != nil {
					return nil, err
				}
			}

			result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
				ReturnableItems: includes,
			})
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

// @oas:path [post] /admin/orders/{id}/claims/{claim_id}
// operationId: "PostOrdersOrderClaimsClaim"
// summary: "Update a Claim"
// description: "Update a Claim's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Order associated with the claim.
//   - (path) claim_id=* {string} The ID of the Claim.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostOrdersOrderClaimsClaimReq"
//
// x-codegen:
//
//	method: updateClaim
//	params: AdminPostOrdersOrderClaimsClaimParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orders.updateClaim(orderId, claimId, {
//     no_notification: true
//     })
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdateClaim } from "medusa-react"
//
//     type Props = {
//     orderId: string
//     claimId: string
//     }
//
//     const Claim = ({ orderId, claimId }: Props) => {
//     const updateClaim = useAdminUpdateClaim(orderId)
//     // ...
//
//     const handleUpdate = () => {
//     updateClaim.mutate({
//     claim_id: claimId,
//     no_notification: false
//     }, {
//     onSuccess: ({ order }) => {
//     console.log(order.claims)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Claim
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/orders/{id}/claims/{claim_id}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "no_notification": true
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrdersRes"
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
func (m *Order) UpdateClaim(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.UpdateClaimInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	claimId, err := api.BindDelete(context, "claim_id")
	if err != nil {
		return err
	}

	if _, err := m.r.ClaimService().SetContext(context.Context()).Update(claimId, model); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

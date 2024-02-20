package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type OrderEdit struct {
	r    Registry
	name string
}

func NewOrderEdit(r Registry) *OrderEdit {
	m := OrderEdit{r: r, name: "order_edit"}
	return &m
}

func (m *OrderEdit) SetRoutes(router fiber.Router) {
	route := router.Group("/order-edits")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)

	route.Post("/:id/cancel", m.Cancel)
	route.Post("/:id/confirm", m.Confirm)
	route.Post("/:id/request", m.RequestConfirmation)
	route.Post("/:id/items", m.AddLineItem)
	route.Post("/:id/items/:item_id", m.UpdateLineItem)
	route.Delete("/:id/items/:item_id", m.DeleteLineItem)
	route.Delete("/:id/changes/:change_id", m.DeleteItemChange)
}

// @oas:path [get] /admin/order-edits/{id}
// operationId: "GetOrderEditsOrderEdit"
// summary: "Get an Order Edit"
// description: "Retrieve an Order Edit's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the OrderEdit.
//   - (query) expand {string} Comma-separated relations that should be expanded in each returned order edit.
//   - (query) fields {string} Comma-separated fields that should be included in the returned order edit.
//
// x-codegen:
//
//	method: retrieve
//	queryParams: GetOrderEditsOrderEditParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orderEdits.retrieve(orderEditId)
//     .then(({ order_edit }) => {
//     console.log(order_edit.id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminOrderEdit } from "medusa-react"
//
//     type Props = {
//     orderEditId: string
//     }
//
//     const OrderEdit = ({ orderEditId }: Props) => {
//     const {
//     order_edit,
//     isLoading,
//     } = useAdminOrderEdit(orderEditId)
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {order_edit && <span>{order_edit.status}</span>}
//     </div>
//     )
//     }
//
//     export default OrderEdit
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/order-edits/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Order Edits
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrderEditsRes"
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
func (m *OrderEdit) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.OrderEditService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [get] /admin/order-edits
// operationId: "GetOrderEdits"
// summary: "List Order Edits"
// description: "Retrieve a list of order edits. The order edits can be filtered by fields such as `q` or `order_id`. The order edits can also be paginated."
// x-authenticated: true
// parameters:
//   - (query) q {string} term to search order edits' internal note.
//   - (query) order_id {string} Filter by order ID
//   - (query) limit=20 {number} Limit the number of order edits returned.
//   - (query) offset=0 {number} The number of order edits to skip when retrieving the order edits.
//   - (query) expand {string} Comma-separated relations that should be expanded in each returned order edit.
//   - (query) fields {string} Comma-separated fields that should be included in each returned order edit.
//
// x-codegen:
//
//	method: list
//	queryParams: GetOrderEditsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orderEdits.list()
//     .then(({ order_edits, count, limit, offset }) => {
//     console.log(order_edits.length)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminOrderEdits } from "medusa-react"
//
//     const OrderEdits = () => {
//     const { order_edits, isLoading } = useAdminOrderEdits()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {order_edits && !order_edits.length && (
//     <span>No Order Edits</span>
//     )}
//     {order_edits && order_edits.length > 0 && (
//     <ul>
//     {order_edits.map((orderEdit) => (
//     <li key={orderEdit.id}>
//     {orderEdit.status}
//     </li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default OrderEdits
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/order-edits' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Order Edits
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrderEditsListRes"
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
func (m *OrderEdit) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableOrderEdit](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.OrderEditService().SetContext(context.Context()).ListAndCount(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"order_edits": result,
		"count":       count,
		"offset":      config.Skip,
		"limit":       config.Take,
	})
}

// @oas:path [post] /admin/order-edits
// operationId: "PostOrderEdits"
// summary: "Create an OrderEdit"
// description: "Create an Order Edit."
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostOrderEditsReq"
//
// x-authenticated: true
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
//     // must be previously logged in or use api token
//     medusa.admin.orderEdits.create({ orderId })
//     .then(({ order_edit }) => {
//     console.log(order_edit.id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreateOrderEdit } from "medusa-react"
//
//     const CreateOrderEdit = () => {
//     const createOrderEdit = useAdminCreateOrderEdit()
//
//     const handleCreateOrderEdit = (orderId: string) => {
//     createOrderEdit.mutate({
//     order_id: orderId,
//     }, {
//     onSuccess: ({ order_edit }) => {
//     console.log(order_edit.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default CreateOrderEdit
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/order-edits' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{ "order_id": "my_order_id", "internal_note": "my_optional_note" }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Order Edits
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrderEditsRes"
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
func (m *OrderEdit) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateOrderEditInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	userId := api.GetUser(context)

	result, err := m.r.OrderEditService().SetContext(context.Context()).Create(model, userId)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/order-edits/{id}
// operationId: "PostOrderEditsOrderEdit"
// summary: "Update an Order Edit"
// description: "Update an Order Edit's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the OrderEdit.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostOrderEditsOrderEditReq"
//
// x-codegen:
//
//	method: update
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orderEdits.update(orderEditId, {
//     internal_note: "internal reason XY"
//     })
//     .then(({ order_edit }) => {
//     console.log(order_edit.id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdateOrderEdit } from "medusa-react"
//
//     type Props = {
//     orderEditId: string
//     }
//
//     const OrderEdit = ({ orderEditId }: Props) => {
//     const updateOrderEdit = useAdminUpdateOrderEdit(
//     orderEditId,
//     )
//
//     const handleUpdate = (
//     internalNote: string
//     ) => {
//     updateOrderEdit.mutate({
//     internal_note: internalNote
//     }, {
//     onSuccess: ({ order_edit }) => {
//     console.log(order_edit.internal_note)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default OrderEdit
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/order-edits/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "internal_note": "internal reason XY"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Order Edits
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrderEditsRes"
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
func (m *OrderEdit) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateOrderEditInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.OrderEditService().SetContext(context.Context()).Update(id, &models.OrderEdit{InternalNote: model.InternalNote})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/order-edits/{id}
// operationId: "DeleteOrderEditsOrderEdit"
// summary: "Delete an Order Edit"
// description: "Delete an Order Edit. Only order edits that have the status `created` can be deleted."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Order Edit to delete.
//
// x-codegen:
//
//	method: delete
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orderEdits.delete(orderEditId)
//     .then(({ id, object, deleted }) => {
//     console.log(id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteOrderEdit } from "medusa-react"
//
//     type Props = {
//     orderEditId: string
//     }
//
//     const OrderEdit = ({ orderEditId }: Props) => {
//     const deleteOrderEdit = useAdminDeleteOrderEdit(
//     orderEditId
//     )
//
//     const handleDelete = () => {
//     deleteOrderEdit.mutate(void 0, {
//     onSuccess: ({ id, object, deleted }) => {
//     console.log(id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default OrderEdit
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/order-edits/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Order Edits
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrderEditDeleteRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
func (m *OrderEdit) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.OrderEditService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "order-edit",
		"deleted": true,
	})
}

// @oas:path [post] /admin/order-edits/{id}/items
// operationId: "PostOrderEditsEditLineItems"
// summary: "Add a Line Item"
// description: "Create a line item change in the order edit that indicates adding an item in the original order. The item will not be added to the original order until the order edit is
//
//	confirmed."
//
// parameters:
//   - (path) id=* {string} The ID of the Order Edit.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostOrderEditsEditLineItemsReq"
//
// x-authenticated: true
// x-codegen:
//
//	method: addLineItem
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orderEdits.addLineItem(orderEditId, {
//     variant_id,
//     quantity
//     })
//     .then(({ order_edit }) => {
//     console.log(order_edit.id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminOrderEditAddLineItem } from "medusa-react"
//
//     type Props = {
//     orderEditId: string
//     }
//
//     const OrderEdit = ({ orderEditId }: Props) => {
//     const addLineItem = useAdminOrderEditAddLineItem(
//     orderEditId
//     )
//
//     const handleAddLineItem =
//     (quantity: number, variantId: string) => {
//     addLineItem.mutate({
//     quantity,
//     variant_id: variantId,
//     }, {
//     onSuccess: ({ order_edit }) => {
//     console.log(order_edit.changes)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default OrderEdit
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/order-edits/{id}/items' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{ "variant_id": "variant_01G1G5V2MRX2V3PVSR2WXYPFB6", "quantity": 3 }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Order Edits
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrderEditsRes"
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
func (m *OrderEdit) AddLineItem(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.AddOrderEditLineItemInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if err := m.r.OrderEditService().SetContext(context.Context()).AddLineItem(id, model); err != nil {
		return err
	}

	order, err := m.r.OrderEditService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err := m.r.OrderEditService().SetContext(context.Context()).DecorateTotals(order)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/order-edits/{id}/cancel
// operationId: "PostOrderEditsOrderEditCancel"
// summary: "Cancel an Order Edit"
// description: "Cancel an Order Edit."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the OrderEdit.
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
//     medusa.admin.orderEdits.cancel(orderEditId)
//     .then(({ order_edit }) => {
//     console.log(order_edit.id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import {
//     useAdminCancelOrderEdit
//     } from "medusa-react"
//
//     type Props = {
//     orderEditId: string
//     }
//
//     const OrderEdit = ({ orderEditId }: Props) => {
//     const cancelOrderEdit =
//     useAdminCancelOrderEdit(
//     orderEditId
//     )
//
//     const handleCancel = () => {
//     cancelOrderEdit.mutate(void 0, {
//     onSuccess: ({ order_edit }) => {
//     console.log(
//     order_edit.id
//     )
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default OrderEdit
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/order-edits/{id}/cancel' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Order Edits
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrderEditsRes"
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
//	"500":
//	  description: "Internal Server"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/responses/500_error"
func (m *OrderEdit) Cancel(context fiber.Ctx) error {
	id, err := utils.ParseUUID(context.Params("id"))
	if err != nil {
		return err
	}

	userId := api.GetUser(context)

	if _, err := m.r.OrderEditService().SetContext(context.Context()).Cancel(id, userId); err != nil {
		return err
	}

	order, err := m.r.OrderEditService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err := m.r.OrderEditService().SetContext(context.Context()).DecorateTotals(order)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/order-edits/{id}/confirm
// operationId: "PostOrderEditsOrderEditConfirm"
// summary: "Confirm an OrderEdit"
// description: "Confirm an Order Edit. This will reflect the changes in the order edit on the associated order."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the order edit.
//
// x-codegen:
//
//	method: confirm
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orderEdits.confirm(orderEditId)
//     .then(({ order_edit }) => {
//     console.log(order_edit.id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminConfirmOrderEdit } from "medusa-react"
//
//     type Props = {
//     orderEditId: string
//     }
//
//     const OrderEdit = ({ orderEditId }: Props) => {
//     const confirmOrderEdit = useAdminConfirmOrderEdit(
//     orderEditId
//     )
//
//     const handleConfirmOrderEdit = () => {
//     confirmOrderEdit.mutate(void 0, {
//     onSuccess: ({ order_edit }) => {
//     console.log(
//     order_edit.confirmed_at,
//     order_edit.confirmed_by
//     )
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default OrderEdit
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/order-edits/{id}/confirm' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Order Edits
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrderEditsRes"
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
//	"500":
//	  description: "Internal Server"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/responses/500_error"
func (m *OrderEdit) Confirm(context fiber.Ctx) error {
	id, err := utils.ParseUUID(context.Params("id"))
	if err != nil {
		return err
	}

	userId := api.GetUser(context)

	if _, err := m.r.OrderEditService().SetContext(context.Context()).Confirm(id, userId); err != nil {
		return err
	}

	order, err := m.r.OrderEditService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err := m.r.OrderEditService().SetContext(context.Context()).DecorateTotals(order)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/order-edits/{id}/request
// operationId: "PostOrderEditsOrderEditRequest"
// summary: "Request Confirmation"
// description: "Request customer confirmation of an Order Edit. This would emit the event `order-edit.requested` which Notification Providers listen to and send
//
//	a notification to the customer about the order edit."
//
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Order Edit.
//
// x-codegen:
//
//	method: requestConfirmation
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orderEdits.requestConfirmation(orderEditId)
//     .then({ order_edit }) => {
//     console.log(order_edit.id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import {
//     useAdminRequestOrderEditConfirmation,
//     } from "medusa-react"
//
//     type Props = {
//     orderEditId: string
//     }
//
//     const OrderEdit = ({ orderEditId }: Props) => {
//     const requestOrderConfirmation =
//     useAdminRequestOrderEditConfirmation(
//     orderEditId
//     )
//
//     const handleRequestConfirmation = () => {
//     requestOrderConfirmation.mutate(void 0, {
//     onSuccess: ({ order_edit }) => {
//     console.log(
//     order_edit.requested_at,
//     order_edit.requested_by
//     )
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default OrderEdit
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/order-edits/{id}/request' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Order Edits
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrderEditsRes"
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
//	"500":
//	  description: "Internal Server"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/responses/500_error"
func (m *OrderEdit) RequestConfirmation(context fiber.Ctx) error {
	model, id, err := api.BindWithUUID[types.OrderEditsRequestConfirmation](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	userId := api.GetUser(context)

	orderEdit, err := m.r.OrderEditService().SetContext(context.Context()).RequestConfirmation(id, userId)
	if err != nil {
		return err
	}

	total, err := m.r.OrderEditService().SetContext(context.Context()).DecorateTotals(orderEdit)
	if err != nil {
		return err
	}

	if total.DifferenceDue > 0 {
		order, err := m.r.OrderService().SetContext(context.Context()).RetrieveById(orderEdit.OrderId.UUID, &sql.Options{
			Selects: []string{"currency_code", "region_id"},
		})
		if err != nil {
			return err
		}

		paymentCollection, err := m.r.PaymentCollectionService().SetContext(context.Context()).Create(&types.CreatePaymentCollectionInput{
			Type:         models.PaymentCollectionTypeOrderEdit,
			Amount:       total.DifferenceDue,
			CurrencyCode: order.CurrencyCode,
			RegionId:     order.RegionId.UUID,
			Description:  model.PaymentCollectionDescription,
			CreatedBy:    userId,
		})
		if err != nil {
			return err
		}

		orderEdit.PaymentCollectionId = uuid.NullUUID{UUID: paymentCollection.Id}

		_, err = m.r.OrderEditService().SetContext(context.Context()).Update(orderEdit.Id, orderEdit)
		if err != nil {
			return err
		}
	}

	order, err := m.r.OrderEditService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err := m.r.OrderEditService().SetContext(context.Context()).DecorateTotals(order)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/order-edits/{id}/items/{item_id}
// operationId: "DeleteOrderEditsOrderEditLineItemsLineItem"
// summary: "Delete Line Item"
// description: "Create a line item change in the order edit that indicates deleting an item in the original order. The item in the original order will not be deleted until the order edit is
//
//	confirmed."
//
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Order Edit.
//   - (path) item_id=* {string} The ID of line item in the original order.
//
// x-codegen:
//
//	method: removeLineItem
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orderEdits.removeLineItem(orderEditId, lineItemId)
//     .then(({ order_edit }) => {
//     console.log(order_edit.id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminOrderEditDeleteLineItem } from "medusa-react"
//
//     type Props = {
//     orderEditId: string
//     itemId: string
//     }
//
//     const OrderEditLineItem = ({
//     orderEditId,
//     itemId
//     }: Props) => {
//     const removeLineItem = useAdminOrderEditDeleteLineItem(
//     orderEditId,
//     itemId
//     )
//
//     const handleRemoveLineItem = () => {
//     removeLineItem.mutate(void 0, {
//     onSuccess: ({ order_edit }) => {
//     console.log(order_edit.changes)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default OrderEditLineItem
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/order-edits/{id}/items/{item_id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Order Edits
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrderEditsRes"
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
func (m *OrderEdit) DeleteLineItem(context fiber.Ctx) error {
	id, err := utils.ParseUUID(context.Params("id"))
	if err != nil {
		return err
	}

	itemId, err := utils.ParseUUID(context.Params("item_id"))
	if err != nil {
		return err
	}

	if err := m.r.OrderEditService().SetContext(context.Context()).RemoveLineItem(id, itemId); err != nil {
		return err
	}

	order, err := m.r.OrderEditService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err := m.r.OrderEditService().SetContext(context.Context()).DecorateTotals(order)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/order-edits/{id}/changes/{change_id}
// operationId: "DeleteOrderEditsOrderEditItemChange"
// summary: "Delete a Line Item Change"
// description: "Delete a line item change that indicates the addition, deletion, or update of a line item in the original order."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Order Edit.
//   - (path) change_id=* {string} The ID of the Line Item Change to delete.
//
// x-codegen:
//
//	method: deleteItemChange
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orderEdits.deleteItemChange(orderEdit_id, itemChangeId)
//     .then(({ id, object, deleted }) => {
//     console.log(id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteOrderEditItemChange } from "medusa-react"
//
//     type Props = {
//     orderEditId: string
//     itemChangeId: string
//     }
//
//     const OrderEditItemChange = ({
//     orderEditId,
//     itemChangeId
//     }: Props) => {
//     const deleteItemChange = useAdminDeleteOrderEditItemChange(
//     orderEditId,
//     itemChangeId
//     )
//
//     const handleDeleteItemChange = () => {
//     deleteItemChange.mutate(void 0, {
//     onSuccess: ({ id, object, deleted }) => {
//     console.log(id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default OrderEditItemChange
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/order-edits/{id}/changes/{change_id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
// - api_token: []
// - cookie_auth: []
// - jwt_token: []
// tags:
// - Order Edits
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrderEditItemChangeDeleteRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
func (m *OrderEdit) DeleteItemChange(context fiber.Ctx) error {
	id, err := utils.ParseUUID(context.Params("id"))
	if err != nil {
		return err
	}

	changeId, err := utils.ParseUUID(context.Params("change_id"))
	if err != nil {
		return err
	}

	if err := m.r.OrderEditService().SetContext(context.Context()).DeleteItemChange(id, changeId); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      changeId,
		"object":  "item-change",
		"deleted": true,
	})
}

// @oas:path [post] /admin/order-edits/{id}/items/{item_id}
// operationId: "PostOrderEditsEditLineItemsLineItem"
// summary: "Upsert Line Item Change"
// description: "Create or update a line item change in the order edit that indicates addition, deletion, or update of a line item into an original order. Line item changes
// are only reflected on the original order after the order edit is confirmed."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Order Edit.
//   - (path) item_id=* {string} The ID of the line item in the original order.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostOrderEditsEditLineItemsLineItemReq"
//
// x-codegen:
//
//	method: updateLineItem
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.orderEdits.updateLineItem(orderEditId, lineItemId, {
//     quantity: 5
//     })
//     .then(({ order_edit }) => {
//     console.log(order_edit.id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminOrderEditUpdateLineItem } from "medusa-react"
//
//     type Props = {
//     orderEditId: string
//     itemId: string
//     }
//
//     const OrderEditItemChange = ({
//     orderEditId,
//     itemId
//     }: Props) => {
//     const updateLineItem = useAdminOrderEditUpdateLineItem(
//     orderEditId,
//     itemId
//     )
//
//     const handleUpdateLineItem = (quantity: number) => {
//     updateLineItem.mutate({
//     quantity,
//     }, {
//     onSuccess: ({ order_edit }) => {
//     console.log(order_edit.items)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default OrderEditItemChange
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/order-edits/{id}/items/{item_id}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{ "quantity": 5 }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Order Edits
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminOrderEditsRes"
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
func (m *OrderEdit) UpdateLineItem(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.OrderEditsEditLineItem](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	itemId, err := utils.ParseUUID(context.Params("item_id"))
	if err != nil {
		return err
	}

	if err := m.r.OrderEditService().SetContext(context.Context()).UpdateLineItem(id, itemId, model.Quantity); err != nil {
		return err
	}

	order, err := m.r.OrderEditService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err := m.r.OrderEditService().SetContext(context.Context()).DecorateTotals(order)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/services"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type DraftOrder struct {
	r    Registry
	name string
}

func NewDraftOrder(r Registry) *DraftOrder {
	m := DraftOrder{r: r, name: "draft_order"}
	return &m
}

func (m *DraftOrder) SetRoutes(router fiber.Router) {
	route := router.Group("/draft-orders")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)

	route.Delete("/:id/line-items/:line_id", m.DeleteLineItem)
	route.Post("/:id/line-items", m.CreateLineItem)
	route.Post("/:id/line-items/:line_id", m.UpdateLineItem)
	route.Post("/:id/pay", m.RegisterPayment)
}

// @oas:path [get] /admin/draft-orders/{id}
// operationId: "GetDraftOrdersDraftOrder"
// summary: "Get a Draft Order"
// description: "Retrieve a Draft Order's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Draft Order.
//
// x-codegen:
//
//	method: retrieve
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.draftOrders.retrieve(draftOrderId)
//     .then(({ draft_order }) => {
//     console.log(draft_order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDraftOrder } from "medusa-react"
//
//     type Props = {
//     draftOrderId: string
//     }
//
//     const DraftOrder = ({ draftOrderId }: Props) => {
//     const {
//     draft_order,
//     isLoading,
//     } = useAdminDraftOrder(draftOrderId)
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {draft_order && <span>{draft_order.display_id}</span>}
//
//     </div>
//     )
//     }
//
//     export default DraftOrder
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/draft-orders/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Draft Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDraftOrdersRes"
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
func (m *DraftOrder) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.DraftOrderService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [get] /admin/draft-orders
// operationId: "GetDraftOrders"
// summary: "List Draft Orders"
// description: "Retrieve an list of Draft Orders. The draft orders can be filtered by fields such as `q`. The draft orders can also paginated."
// x-authenticated: true
// parameters:
//   - (query) offset=0 {number} The number of draft orders to skip when retrieving the draft orders.
//   - (query) limit=50 {number} Limit the number of draft orders returned.
//   - (query) q {string} a term to search draft orders' display IDs and emails in the draft order's cart
//
// x-codegen:
//
//	method: list
//	queryParams: AdminGetDraftOrdersParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.draftOrders.list()
//     .then(({ draft_orders, limit, offset, count }) => {
//     console.log(draft_orders.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDraftOrders } from "medusa-react"
//
//     const DraftOrders = () => {
//     const { draft_orders, isLoading } = useAdminDraftOrders()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {draft_orders && !draft_orders.length && (
//     <span>No Draft Orders</span>
//     )}
//     {draft_orders && draft_orders.length > 0 && (
//     <ul>
//     {draft_orders.map((order) => (
//     <li key={order.id}>{order.display_id}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default DraftOrders
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/draft-orders' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Draft Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDraftOrdersListRes"
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
func (m *DraftOrder) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableDraftOrder](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.DraftOrderService().SetContext(context.Context()).ListAndCount(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"draft_orders": result,
		"count":        count,
		"offset":       config.Skip,
		"limit":        config.Take,
	})
}

// @oas:path [post] /admin/draft-orders
// operationId: "PostDraftOrders"
// summary: "Create a Draft Order"
// description: "Create a Draft Order. A draft order is not transformed into an order until payment is captured."
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostDraftOrdersReq"
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
//     // must be previously logged in or use api token
//     medusa.admin.draftOrders.create({
//     email: "user@example.com",
//     region_id,
//     items: [
//     {
//     quantity: 1
//     }
//     ],
//     shipping_methods: [
//     {
//     option_id
//     }
//     ],
//     })
//     .then(({ draft_order }) => {
//     console.log(draft_order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreateDraftOrder } from "medusa-react"
//
//     type DraftOrderData = {
//     email: string
//     region_id: string
//     items: {
//     quantity: number,
//     variant_id: string
//     }[]
//     shipping_methods: {
//     option_id: string
//     price: number
//     }[]
//     }
//
//     const CreateDraftOrder = () => {
//     const createDraftOrder = useAdminCreateDraftOrder()
//     // ...
//
//     const handleCreate = (data: DraftOrderData) => {
//     createDraftOrder.mutate(data, {
//     onSuccess: ({ draft_order }) => {
//     console.log(draft_order.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default CreateDraftOrder
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/draft-orders' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "email": "user@example.com",
//     "region_id": "{region_id}"
//     "items": [
//     {
//     "quantity": 1
//     }
//     ],
//     "shipping_methods": [
//     {
//     "option_id": "{option_id}"
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
//   - Draft Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDraftOrdersRes"
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
func (m *DraftOrder) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.DraftOrderCreate](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.DraftOrderService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/draft-orders/{id}
// operationId: PostDraftOrdersDraftOrder
// summary: Update a Draft Order
// description: "Update a Draft Order's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Draft Order.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostDraftOrdersDraftOrderReq"
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
//     medusa.admin.draftOrders.update(draftOrderId, {
//     email: "user@example.com"
//     })
//     .then(({ draft_order }) => {
//     console.log(draft_order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdateDraftOrder } from "medusa-react"
//
//     type Props = {
//     draftOrderId: string
//     }
//
//     const DraftOrder = ({ draftOrderId }: Props) => {
//     const updateDraftOrder = useAdminUpdateDraftOrder(
//     draftOrderId
//     )
//     // ...
//
//     const handleUpdate = (email: string) => {
//     updateDraftOrder.mutate({
//     email,
//     }, {
//     onSuccess: ({ draft_order }) => {
//     console.log(draft_order.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default DraftOrder
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/draft-orders/{id}' \
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
//   - Draft Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDraftOrdersRes"
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
func (m *DraftOrder) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.DraftOrderUpdate](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	draftOrder, err := m.r.DraftOrderService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	if draftOrder.Status == models.DraftOrderStatusCompleted {
		return utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"You are only allowed to update open draft orders",
		)
	}

	if _, err := m.r.DraftOrderService().SetContext(context.Context()).Update(id, &models.DraftOrder{NoNotificationOrder: model.NoNotificationOrder}); err != nil {
		return err
	}

	_, err = m.r.CartService().SetContext(context.Context()).Update(id, nil, &types.CartUpdateProps{
		RegionId:          model.RegionId,
		CountryCode:       model.CountryCode,
		Email:             model.Email,
		BillingAddressId:  model.BillingAddressId,
		BillingAddress:    model.BillingAddress,
		ShippingAddressId: model.ShippingAddressId,
		ShippingAddress:   model.ShippingAddress,
		Discounts:         model.Discounts,
		CustomerId:        model.CustomerId,
	})
	if err != nil {
		return err
	}

	result, err := m.r.DraftOrderService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/draft-orders/{id}
// operationId: DeleteDraftOrdersDraftOrder
// summary: Delete a Draft Order
// description: "Delete a Draft Order."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Draft Order.
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
//     medusa.admin.draftOrders.delete(draftOrderId)
//     .then(({ id, object, deleted }) => {
//     console.log(id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteDraftOrder } from "medusa-react"
//
//     type Props = {
//     draftOrderId: string
//     }
//
//     const DraftOrder = ({ draftOrderId }: Props) => {
//     const deleteDraftOrder = useAdminDeleteDraftOrder(
//     draftOrderId
//     )
//     // ...
//
//     const handleDelete = () => {
//     deleteDraftOrder.mutate(void 0, {
//     onSuccess: ({ id, object, deleted }) => {
//     console.log(id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default DraftOrder
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/draft-orders/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Draft Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDraftOrdersDeleteRes"
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
func (m *DraftOrder) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.DraftOrderService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "draft-order",
		"deleted": true,
	})
}

// @oas:path [post] /admin/draft-orders/{id}/line-items
// operationId: "PostDraftOrdersDraftOrderLineItems"
// summary: "Create a Line Item"
// description: "Create a Line Item in the Draft Order."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Draft Order.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostDraftOrdersDraftOrderLineItemsReq"
//
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
//     medusa.admin.draftOrders.addLineItem(draftOrderId, {
//     quantity: 1
//     })
//     .then(({ draft_order }) => {
//     console.log(draft_order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDraftOrderAddLineItem } from "medusa-react"
//
//     type Props = {
//     draftOrderId: string
//     }
//
//     const DraftOrder = ({ draftOrderId }: Props) => {
//     const addLineItem = useAdminDraftOrderAddLineItem(
//     draftOrderId
//     )
//     // ...
//
//     const handleAdd = (quantity: number) => {
//     addLineItem.mutate({
//     quantity,
//     }, {
//     onSuccess: ({ draft_order }) => {
//     console.log(draft_order.cart)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default DraftOrder
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/draft-orders/{id}/line-items' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "quantity": 1
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Draft Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDraftOrdersRes"
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
func (m *DraftOrder) CreateLineItem(context fiber.Ctx) error {
	model, id, err := api.BindWithUUID[types.Item](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	draftOrder, err := m.r.DraftOrderService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	if draftOrder.Status != models.DraftOrderStatusCompleted {
		return utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"You are only allowed to update open draft orders",
		)
	}

	if model.VariantId == uuid.Nil {
		line, err := m.r.LineItemService().SetContext(context.Context()).Generate(
			model.VariantId,
			nil,
			draftOrder.Cart.RegionId.UUID,
			model.Quantity,
			types.GenerateLineItemContext{
				Metadata:  model.Metadata,
				UnitPrice: model.UnitPrice,
			},
		)
		if err != nil {
			return err
		}

		if err := m.r.CartService().SetContext(context.Context()).AddOrUpdateLineItems(draftOrder.CartId.UUID, line, false); err != nil {
			return err
		}
	} else {
		_, err := m.r.LineItemService().SetContext(context.Context()).Create(
			[]models.LineItem{
				{
					BaseModel: core.BaseModel{
						Metadata: model.Metadata,
					},
					CartId:         draftOrder.CartId,
					HasShipping:    true,
					Title:          model.Title,
					AllowDiscounts: false,
					UnitPrice:      model.UnitPrice,
					Quantity:       model.Quantity,
				},
			},
		)
		if err != nil {
			return err
		}
	}

	cart, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(draftOrder.CartId.UUID, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	draftOrder.Cart = cart

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): draftOrder,
	})
}

// @oas:path [post] /admin/draft-orders/{id}/pay
// summary: "Mark Paid"
// operationId: "PostDraftOrdersDraftOrderRegisterPayment"
// description: "Capture the draft order's payment. This will also set the draft order's status to `completed` and create an Order from the draft order. The payment is captured through Medusa's system payment,
//
//	which is manual payment that isn't integrated with any third-party payment provider. It is assumed that the payment capturing is handled manually by the admin."
//
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The Draft Order ID.
//
// x-codegen:
//
//	method: markPaid
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.draftOrders.markPaid(draftOrderId)
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDraftOrderRegisterPayment } from "medusa-react"
//
//     type Props = {
//     draftOrderId: string
//     }
//
//     const DraftOrder = ({ draftOrderId }: Props) => {
//     const registerPayment = useAdminDraftOrderRegisterPayment(
//     draftOrderId
//     )
//     // ...
//
//     const handlePayment = () => {
//     registerPayment.mutate(void 0, {
//     onSuccess: ({ order }) => {
//     console.log(order.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default DraftOrder
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/draft-orders/{id}/pay' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Draft Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPostDraftOrdersDraftOrderRegisterPaymentRes"
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
func (m *DraftOrder) RegisterPayment(context fiber.Ctx) error {
	id, err := utils.ParseUUID(context.Params("id"))
	if err != nil {
		return err
	}

	draftOrder, err := m.r.DraftOrderService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	cart, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(draftOrder.CartId.UUID, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	_, err = m.r.PaymentProviderService().SetContext(context.Context()).CreateSession(uuid.Nil, &types.PaymentSessionInput{
		PaymentSessionId:   cart.PaymentSession.Id,
		ProviderId:         cart.PaymentSession.ProviderId.UUID,
		Cart:               cart,
		Customer:           cart.Customer,
		CurrencyCode:       cart.Payment.CurrencyCode,
		Amount:             cart.Payment.Amount,
		ResourceId:         cart.Id,
		PaymentSessionData: cart.PaymentSession.Data,
		Context:            cart.Context,
	})
	if err != nil {
		return err
	}

	if err := m.r.CartService().SetContext(context.Context()).SetPaymentSession(cart.Id, uuid.Nil); err != nil {
		return err
	}

	if err := m.r.CartService().SetContext(context.Context()).CreateTaxLines(cart.Id, nil); err != nil {
		return err
	}

	_, err = m.r.CartService().SetContext(context.Context()).AuthorizePayment(cart.Id, nil, map[string]interface{}{})
	if err != nil {
		return err
	}

	order, err := m.r.OrderService().SetContext(context.Context()).CreateFromCart(id, nil)
	if err != nil {
		return err
	}

	_, err = m.r.DraftOrderService().SetContext(context.Context()).RegisterCartCompletion(id, order.Id)
	if err != nil {
		return err
	}

	_, err = m.r.OrderService().SetContext(context.Context()).CapturePayment(order.Id)
	if err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, &sql.Options{}, types.TotalsContext{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/draft-orders/{id}/line-items/{line_id}
// operationId: "PostDraftOrdersDraftOrderLineItemsItem"
// summary: "Update a Line Item"
// description: "Update a Line Item in a Draft Order."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Draft Order.
//   - (path) line_id=* {string} The ID of the Line Item.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostDraftOrdersDraftOrderLineItemsItemReq"
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
//     medusa.admin.draftOrders.updateLineItem(draftOrderId, lineId, {
//     quantity: 1
//     })
//     .then(({ draft_order }) => {
//     console.log(draft_order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDraftOrderUpdateLineItem } from "medusa-react"
//
//     type Props = {
//     draftOrderId: string
//     }
//
//     const DraftOrder = ({ draftOrderId }: Props) => {
//     const updateLineItem = useAdminDraftOrderUpdateLineItem(
//     draftOrderId
//     )
//     // ...
//
//     const handleUpdate = (
//     itemId: string,
//     quantity: number
//     ) => {
//     updateLineItem.mutate({
//     item_id: itemId,
//     quantity,
//     })
//     }
//
//     // ...
//     }
//
//     export default DraftOrder
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/draft-orders/{id}/line-items/{line_id}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "quantity": 1
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Draft Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDraftOrdersRes"
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
func (m *DraftOrder) UpdateLineItem(context fiber.Ctx) error {
	model, id, err := api.BindWithUUID[types.Item](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	lineId, err := utils.ParseUUID(context.Params("line_id"))
	if err != nil {
		return err
	}

	draftOrder, err := m.r.DraftOrderService().SetContext(context.Context()).Retrieve(id, &sql.Options{Relations: []string{"cart", "cart.items"}})
	if err != nil {
		return err
	}

	if draftOrder.Status == models.DraftOrderStatusCompleted {
		return utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"You are only allowed to update open draft orders",
		)
	}

	if model.Quantity == 0 {
		if err := m.r.CartService().SetContext(context.Context()).RemoveLineItem(draftOrder.CartId.UUID, uuid.UUIDs{lineId}); err != nil {
			return err
		}
	} else {
		_, ok := lo.Find(draftOrder.Cart.Items, func(v models.LineItem) bool {
			return v.Id == lineId
		})

		if !ok {
			return utils.NewApplictaionError(
				utils.INVALID_DATA,
				"Could not find the line item",
			)
		}

		item := &types.LineItemUpdate{
			RegionId: draftOrder.Cart.RegionId.UUID,
		}

		item.Title = model.Title
		item.UnitPrice = model.UnitPrice
		item.VariantId = model.VariantId
		item.Quantity = model.Quantity
		item.Metadata = model.Metadata

		_, err := m.r.CartService().SetContext(context.Context()).UpdateLineItem(draftOrder.CartId.UUID, lineId, item)
		if err != nil {
			return err
		}
	}

	cart, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(draftOrder.CartId.UUID, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	draftOrder.Cart = cart

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): draftOrder,
	})
}

// @oas:path [delete] /admin/draft-orders/{id}/line-items/{line_id}
// operationId: DeleteDraftOrdersDraftOrderLineItemsItem
// summary: Delete a Line Item
// description: "Delete a Line Item from a Draft Order."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Draft Order.
//   - (path) line_id=* {string} The ID of the line item.
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
//     medusa.admin.draftOrders.removeLineItem(draftOrderId, itemId)
//     .then(({ draft_order }) => {
//     console.log(draft_order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDraftOrderRemoveLineItem } from "medusa-react"
//
//     type Props = {
//     draftOrderId: string
//     }
//
//     const DraftOrder = ({ draftOrderId }: Props) => {
//     const deleteLineItem = useAdminDraftOrderRemoveLineItem(
//     draftOrderId
//     )
//     // ...
//
//     const handleDelete = (itemId: string) => {
//     deleteLineItem.mutate(itemId, {
//     onSuccess: ({ draft_order }) => {
//     console.log(draft_order.cart)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default DraftOrder
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/draft-orders/{id}/line-items/{line_id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Draft Orders
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDraftOrdersRes"
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
func (m *DraftOrder) DeleteLineItem(context fiber.Ctx) error {
	id, err := utils.ParseUUID(context.Params("id"))
	if err != nil {
		return err
	}

	lineId, err := utils.ParseUUID(context.Params("line_id"))
	if err != nil {
		return err
	}

	draftOrder, err := m.r.DraftOrderService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	if draftOrder.Status == models.DraftOrderStatusCompleted {
		return utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"You are only allowed to update open draft orders",
		)
	}

	if err := m.r.CartService().SetContext(context.Context()).RemoveLineItem(draftOrder.CartId.UUID, uuid.UUIDs{lineId}); err != nil {
		return err
	}

	cart, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(draftOrder.CartId.UUID, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	draftOrder.Cart = cart

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): draftOrder,
	})
}

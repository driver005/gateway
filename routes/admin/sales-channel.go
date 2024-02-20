package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type SalesChannel struct {
	r    Registry
	name string
}

func NewSalesChannel(r Registry) *SalesChannel {
	m := SalesChannel{r: r, name: "sales_channel"}
	return &m
}

func (m *SalesChannel) SetRoutes(router fiber.Router) {
	route := router.Group("/sales-channels")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)

	route.Post("/stock-locations", m.AddStockLocation)
	route.Delete("/stock-locations", m.RemoveStockLocation)
	route.Post("/products/batch", m.AddProductsBatch)
	route.Delete("/products/batch", m.DeleteProductsBatch)
}

// @oas:path [get] /admin/sales-channels/{id}
// operationId: "GetSalesChannelsSalesChannel"
// summary: "Get a Sales Channel"
// description: "Retrieve a sales channel's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Sales channel.
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
//     medusa.admin.salesChannels.retrieve(salesChannelId)
//     .then(({ sales_channel }) => {
//     console.log(sales_channel.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminSalesChannel } from "medusa-react"
//
//     type Props = {
//     salesChannelId: string
//     }
//
//     const SalesChannel = ({ salesChannelId }: Props) => {
//     const {
//     sales_channel,
//     isLoading,
//     } = useAdminSalesChannel(salesChannelId)
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {sales_channel && <span>{sales_channel.name}</span>}
//     </div>
//     )
//     }
//
//     export default SalesChannel
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/sales-channels/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Sales Channels
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminSalesChannelsRes"
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
func (m *SalesChannel) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.SalesChannelService().SetContext(context.Context()).RetrieveById(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [get] /admin/sales-channels
// operationId: "GetSalesChannels"
// summary: "List Sales Channels"
// description: "Retrieve a list of sales channels. The sales channels can be filtered by fields such as `q` or `name`. The sales channels can also be sorted or paginated."
// x-authenticated: true
// parameters:
//   - (query) id {string} Filter by a sales channel ID.
//   - (query) name {string} Filter by name.
//   - (query) description {string} Filter by description.
//   - (query) q {string} term used to search sales channels' names and descriptions.
//   - (query) order {string} A sales-channel field to sort-order the retrieved sales channels by.
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
//     name: deleted_at
//     description: Filter by a deletion date range.
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
//   - (query) offset=0 {integer} The number of sales channels to skip when retrieving the sales channels.
//   - (query) limit=20 {integer} Limit the number of sales channels returned.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned sales channels.
//   - (query) fields {string} Comma-separated fields that should be included in the returned sales channels.
//
// x-codegen:
//
//	method: list
//	queryParams: AdminGetSalesChannelsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.salesChannels.list()
//     .then(({ sales_channels, limit, offset, count }) => {
//     console.log(sales_channels.length)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminSalesChannels } from "medusa-react"
//
//     const SalesChannels = () => {
//     const { sales_channels, isLoading } = useAdminSalesChannels()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {sales_channels && !sales_channels.length && (
//     <span>No Sales Channels</span>
//     )}
//     {sales_channels && sales_channels.length > 0 && (
//     <ul>
//     {sales_channels.map((salesChannel) => (
//     <li key={salesChannel.id}>{salesChannel.name}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default SalesChannels
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/sales-channels' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Sales Channels
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminSalesChannelsListRes"
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
func (m *SalesChannel) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableSalesChannel](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.SalesChannelService().SetContext(context.Context()).ListAndCount(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"sales_channels": result,
		"count":          count,
		"offset":         config.Skip,
		"limit":          config.Take,
	})
}

// @oas:path [post] /admin/sales-channels
// operationId: "PostSalesChannels"
// summary: "Create a Sales Channel"
// description: "Create a Sales Channel."
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostSalesChannelsReq"
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
//     medusa.admin.salesChannels.create({
//     name: "App",
//     description: "Mobile app"
//     })
//     .then(({ sales_channel }) => {
//     console.log(sales_channel.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreateSalesChannel } from "medusa-react"
//
//     const CreateSalesChannel = () => {
//     const createSalesChannel = useAdminCreateSalesChannel()
//     // ...
//
//     const handleCreate = (name: string, description: string) => {
//     createSalesChannel.mutate({
//     name,
//     description,
//     }, {
//     onSuccess: ({ sales_channel }) => {
//     console.log(sales_channel.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default CreateSalesChannel
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/sales-channels' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "name": "App"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Sales Channels
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminSalesChannelsRes"
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
func (m *SalesChannel) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateSalesChannelInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.SalesChannelService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/sales-channels/{id}
// operationId: "PostSalesChannelsSalesChannel"
// summary: "Update a Sales Channel"
// description: "Update a Sales Channel's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Sales Channel.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostSalesChannelsSalesChannelReq"
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
//     medusa.admin.salesChannels.update(salesChannelId, {
//     name: "App"
//     })
//     .then(({ sales_channel }) => {
//     console.log(sales_channel.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdateSalesChannel } from "medusa-react"
//
//     type Props = {
//     salesChannelId: string
//     }
//
//     const SalesChannel = ({ salesChannelId }: Props) => {
//     const updateSalesChannel = useAdminUpdateSalesChannel(
//     salesChannelId
//     )
//     // ...
//
//     const handleUpdate = (
//     is_disabled: boolean
//     ) => {
//     updateSalesChannel.mutate({
//     is_disabled,
//     }, {
//     onSuccess: ({ sales_channel }) => {
//     console.log(sales_channel.is_disabled)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default SalesChannel
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/sales-channels/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "name": "App"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Sales Channels
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminSalesChannelsRes"
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
func (m *SalesChannel) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateSalesChannelInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.SalesChannelService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/sales-channels/{id}
// operationId: "DeleteSalesChannelsSalesChannel"
// summary: "Delete a Sales Channel"
// description: "Delete a sales channel. Associated products, stock locations, and other resources are not deleted."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Sales channel.
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
//     medusa.admin.salesChannels.delete(salesChannelId)
//     .then(({ id, object, deleted }) => {
//     console.log(id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteSalesChannel } from "medusa-react"
//
//     type Props = {
//     salesChannelId: string
//     }
//
//     const SalesChannel = ({ salesChannelId }: Props) => {
//     const deleteSalesChannel = useAdminDeleteSalesChannel(
//     salesChannelId
//     )
//     // ...
//
//     const handleDelete = () => {
//     deleteSalesChannel.mutate(void 0, {
//     onSuccess: ({ id, object, deleted }) => {
//     console.log(id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default SalesChannel
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/sales-channels/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Sales Channels
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminSalesChannelsDeleteRes"
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
func (m *SalesChannel) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.SalesChannelService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "sales-channel",
		"deleted": true,
	})
}

// @oas:path [post] /admin/sales-channels/{id}/products/batch
// operationId: "PostSalesChannelsChannelProductsBatch"
// summary: "Add Products to Sales Channel"
// description: "Add a list of products to a sales channel."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Sales channel.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostSalesChannelsChannelProductsBatchReq"
//
// x-codegen:
//
//	method: addProducts
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.salesChannels.addProducts(salesChannelId, {
//     product_ids: [
//     {
//     id: productId
//     }
//     ]
//     })
//     .then(({ sales_channel }) => {
//     console.log(sales_channel.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminAddProductsToSalesChannel } from "medusa-react"
//
//     type Props = {
//     salesChannelId: string
//     }
//
//     const SalesChannel = ({ salesChannelId }: Props) => {
//     const addProducts = useAdminAddProductsToSalesChannel(
//     salesChannelId
//     )
//     // ...
//
//     const handleAddProducts = (productId: string) => {
//     addProducts.mutate({
//     product_ids: [
//     {
//     id: productId,
//     },
//     ],
//     }, {
//     onSuccess: ({ sales_channel }) => {
//     console.log(sales_channel.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default SalesChannel
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/sales-channels/{id}/products/batch' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "product_ids": [
//     {
//     "id": "{product_id}"
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
//   - Sales Channels
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminSalesChannelsRes"
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
func (m *SalesChannel) AddProductsBatch(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.AddProductsToCollectionInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	var productIds uuid.UUIDs
	for _, p := range model.ProductIds {
		productIds = append(productIds, p)
	}

	result, err := m.r.SalesChannelService().SetContext(context.Context()).AddProducts(id, productIds)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/sales-channels/{id}/products/batch
// operationId: "DeleteSalesChannelsChannelProductsBatch"
// summary: "Remove Products from Sales Channel"
// description: "Remove a list of products from a sales channel. This does not delete the product. It only removes the association between the product and the sales channel."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Sales Channel
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminDeleteSalesChannelsChannelProductsBatchReq"
//
// x-codegen:
//
//	method: removeProducts
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.salesChannels.removeProducts(salesChannelId, {
//     product_ids: [
//     {
//     id: productId
//     }
//     ]
//     })
//     .then(({ sales_channel }) => {
//     console.log(sales_channel.id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import {
//     useAdminDeleteProductsFromSalesChannel,
//     } from "medusa-react"
//
//     type Props = {
//     salesChannelId: string
//     }
//
//     const SalesChannel = ({ salesChannelId }: Props) => {
//     const deleteProducts = useAdminDeleteProductsFromSalesChannel(
//     salesChannelId
//     )
//     // ...
//
//     const handleDeleteProducts = (productId: string) => {
//     deleteProducts.mutate({
//     product_ids: [
//     {
//     id: productId,
//     },
//     ],
//     }, {
//     onSuccess: ({ sales_channel }) => {
//     console.log(sales_channel.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default SalesChannel
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/sales-channels/{id}/products/batch' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "product_ids": [
//     {
//     "id": "{product_id}"
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
//   - Sales Channels
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminSalesChannelsRes"
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
func (m *SalesChannel) DeleteProductsBatch(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.AddProductsToCollectionInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	var productIds uuid.UUIDs
	for _, p := range model.ProductIds {
		productIds = append(productIds, p)
	}

	result, err := m.r.SalesChannelService().SetContext(context.Context()).RemoveProducts(id, productIds)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/sales-channels/{id}/stock-locations
// operationId: "PostSalesChannelsSalesChannelStockLocation"
// summary: "Associate a Stock Location"
// description: "Associate a stock location with a Sales Channel."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Sales Channel.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostSalesChannelsChannelStockLocationsReq"
//
// x-codegen:
//
//	method: addLocation
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.salesChannels.addLocation(salesChannelId, {
//     location_id: "loc_123"
//     })
//     .then(({ sales_channel }) => {
//     console.log(sales_channel.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import {
//     useAdminAddLocationToSalesChannel
//     } from "medusa-react"
//
//     type Props = {
//     salesChannelId: string
//     }
//
//     const SalesChannel = ({ salesChannelId }: Props) => {
//     const addLocation = useAdminAddLocationToSalesChannel()
//     // ...
//
//     const handleAddLocation = (locationId: string) => {
//     addLocation.mutate({
//     sales_channel_id: salesChannelId,
//     location_id: locationId
//     }, {
//     onSuccess: ({ sales_channel }) => {
//     console.log(sales_channel.locations)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default SalesChannel
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/sales-channels/{id}/stock-locations' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "locaton_id": "loc_123"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Sales Channels
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminSalesChannelsRes"
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
func (m *SalesChannel) AddStockLocation(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.SalesChannelStockLocations](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if err := m.r.SalesChannelLocationService().SetContext(context.Context()).AssociateLocation(id, model.LocationId); err != nil {
		return err
	}

	result, err := m.r.SalesChannelService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/sales-channels/{id}/stock-locations
// operationId: "DeleteSalesChannelsSalesChannelStockLocation"
// summary: "Remove Stock Location from Sales Channels."
// description: "Remove a stock location from a Sales Channel. This only removes the association between the stock location and the sales channel. It does not delete the stock location."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Sales Channel.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminDeleteSalesChannelsChannelStockLocationsReq"
//
// x-codegen:
//
//	method: removeLocation
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.salesChannels.removeLocation(salesChannelId, {
//     location_id: "loc_id"
//     })
//     .then(({ sales_channel }) => {
//     console.log(sales_channel.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import {
//     useAdminRemoveLocationFromSalesChannel
//     } from "medusa-react"
//
//     type Props = {
//     salesChannelId: string
//     }
//
//     const SalesChannel = ({ salesChannelId }: Props) => {
//     const removeLocation = useAdminRemoveLocationFromSalesChannel()
//     // ...
//
//     const handleRemoveLocation = (locationId: string) => {
//     removeLocation.mutate({
//     sales_channel_id: salesChannelId,
//     location_id: locationId
//     }, {
//     onSuccess: ({ sales_channel }) => {
//     console.log(sales_channel.locations)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default SalesChannel
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/sales-channels/{id}/stock-locations' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "locaton_id": "loc_id"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Sales Channels
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminSalesChannelsDeleteLocationRes"
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
func (m *SalesChannel) RemoveStockLocation(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.SalesChannelStockLocations](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if err := m.r.SalesChannelLocationService().SetContext(context.Context()).RemoveLocation(id, model.LocationId); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "sales-channel",
		"deleted": true,
	})
}

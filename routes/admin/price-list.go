package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type PriceList struct {
	r Registry
}

func NewPriceList(r Registry) *PriceList {
	m := PriceList{r: r}
	return &m
}

func (m *PriceList) SetRoutes(router fiber.Router) {
	route := router.Group("/price-lists")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)

	route.Get("/:id/products", m.ListPriceListProducts)
	route.Delete("/:id/products/:product_id/prices", m.DeleteProductPrices)
	route.Delete("/:id/products/prices/batch", m.DeleteProductPricesBatch)
	route.Delete("/:id/variants/:variant_id/prices", m.DeleteVariantPrices)
	route.Delete("/:id/prices/batch", m.DeletePricesBatch)
	route.Post("/:id/prices/batch", m.AddPricesBatch)
}

// @oas:path [get] /admin/price-lists/{id}
// operationId: "GetPriceListsPriceList"
// summary: "Get a Price List"
// description: "Retrieve a Price List's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Price List.
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
//     medusa.admin.priceLists.retrieve(priceListId)
//     .then(({ price_list }) => {
//     console.log(price_list.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminPriceList } from "medusa-react"
//
//     type Props = {
//     priceListId: string
//     }
//
//     const PriceList = ({
//     priceListId
//     }: Props) => {
//     const {
//     price_list,
//     isLoading,
//     } = useAdminPriceList(priceListId)
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {price_list && <span>{price_list.name}</span>}
//     </div>
//     )
//     }
//
//     export default PriceList
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/price-lists/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Price Lists
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPriceListRes"
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
func (m *PriceList) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.PriceListService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [get] /admin/price-lists
// operationId: "GetPriceLists"
// summary: "List Price Lists"
// description: "Retrieve a list of price lists. The price lists can be filtered by fields such as `q` or `status`. The price lists can also be sorted or paginated."
// x-authenticated: true
// parameters:
//   - (query) limit=10 {number} Limit the number of price lists returned.
//   - (query) offset=0 {number} The number of price lists to skip when retrieving the price lists.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned price lists.
//   - (query) fields {string} Comma-separated fields that should be included in the returned price lists.
//   - (query) order {string} A price-list field to sort-order the retrieved price lists by.
//   - (query) id {string} Filter by ID
//   - (query) q {string} term to search price lists' description, name, and customer group's name.
//   - in: query
//     name: status
//     style: form
//     explode: false
//     description: Filter by status.
//     schema:
//     type: array
//     items:
//     type: string
//     enum: [active, draft]
//   - (query) name {string} Filter by name
//   - in: query
//     name: customer_groups
//     style: form
//     explode: false
//     description: Filter by customer-group IDs.
//     schema:
//     type: array
//     items:
//     type: string
//   - in: query
//     name: type
//     style: form
//     explode: false
//     description: Filter by type.
//     schema:
//     type: array
//     items:
//     type: string
//     enum: [sale, override]
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
//
// x-codegen:
//
//	method: list
//	queryParams: AdminGetPriceListPaginationParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.priceLists.list()
//     .then(({ price_lists, limit, offset, count }) => {
//     console.log(price_lists.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminPriceLists } from "medusa-react"
//
//     const PriceLists = () => {
//     const { price_lists, isLoading } = useAdminPriceLists()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {price_lists && !price_lists.length && (
//     <span>No Price Lists</span>
//     )}
//     {price_lists && price_lists.length > 0 && (
//     <ul>
//     {price_lists.map((price_list) => (
//     <li key={price_list.id}>{price_list.name}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default PriceLists
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/price-lists' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Price Lists
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPriceListsListRes"
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
func (m *PriceList) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterablePriceList](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.PriceListService().SetContext(context.Context()).ListAndCount(model, config)
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

// @oas:path [post] /admin/price-lists
// operationId: "PostPriceListsPriceList"
// summary: "Create a Price List"
// description: "Create a Price List."
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostPriceListsPriceListReq"
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
//     import { PriceListType } from "@medusajs/medusa"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.priceLists.create({
//     name: "New Price List",
//     description: "A new price list",
//     type: PriceListType.SALE,
//     prices: [
//     {
//     amount: 1000,
//     variant_id,
//     currency_code: "eur"
//     }
//     ]
//     })
//     .then(({ price_list }) => {
//     console.log(price_list.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import {
//     PriceListStatus,
//     PriceListType,
//     } from "@medusajs/medusa"
//     import { useAdminCreatePriceList } from "medusa-react"
//
//     type CreateData = {
//     name: string
//     description: string
//     type: PriceListType
//     status: PriceListStatus
//     prices: {
//     amount: number
//     variant_id: string
//     currency_code: string
//     max_quantity: number
//     }[]
//     }
//
//     const CreatePriceList = () => {
//     const createPriceList = useAdminCreatePriceList()
//     // ...
//
//     const handleCreate = (
//     data: CreateData
//     ) => {
//     createPriceList.mutate(data, {
//     onSuccess: ({ price_list }) => {
//     console.log(price_list.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default CreatePriceList
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/price-lists' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "name": "New Price List",
//     "description": "A new price list",
//     "type": "sale",
//     "prices": [
//     {
//     "amount": 1000,
//     "variant_id": "afafa",
//     "currency_code": "eur"
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
//   - Price Lists
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPriceListRes"
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
func (m *PriceList) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreatePriceListInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.PriceListService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /admin/price-lists/{id}
// operationId: "PostPriceListsPriceListPriceList"
// summary: "Update a Price List"
// description: "Update a Price List's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Price List.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostPriceListsPriceListPriceListReq"
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
//     medusa.admin.priceLists.update(priceListId, {
//     name: "New Price List"
//     })
//     .then(({ price_list }) => {
//     console.log(price_list.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdatePriceList } from "medusa-react"
//
//     type Props = {
//     priceListId: string
//     }
//
//     const PriceList = ({
//     priceListId
//     }: Props) => {
//     const updatePriceList = useAdminUpdatePriceList(priceListId)
//     // ...
//
//     const handleUpdate = (
//     endsAt: Date
//     ) => {
//     updatePriceList.mutate({
//     ends_at: endsAt,
//     }, {
//     onSuccess: ({ price_list }) => {
//     console.log(price_list.ends_at)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default PriceList
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/price-lists/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "name": "New Price List"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Price Lists
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPriceListRes"
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
func (m *PriceList) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdatePriceListInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.PriceListService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [delete] /admin/price-lists/{id}
// operationId: "DeletePriceListsPriceList"
// summary: "Delete a Price List"
// description: "Delete a Price List and its associated prices."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Price List.
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
//     medusa.admin.priceLists.delete(priceListId)
//     .then(({ id, object, deleted }) => {
//     console.log(id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeletePriceList } from "medusa-react"
//
//     type Props = {
//     priceListId: string
//     }
//
//     const PriceList = ({
//     priceListId
//     }: Props) => {
//     const deletePriceList = useAdminDeletePriceList(priceListId)
//     // ...
//
//     const handleDelete = () => {
//     deletePriceList.mutate(void 0, {
//     onSuccess: ({ id, object, deleted }) => {
//     console.log(id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default PriceList
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/price-lists/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Price Lists
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPriceListDeleteRes"
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
func (m *PriceList) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.PriceListService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "price-list",
		"deleted": true,
	})
}

// @oas:path [post] /admin/price-lists/{id}/prices/batch
// operationId: "PostPriceListsPriceListPricesBatch"
// summary: "Add or Update Prices"
// description: "Add or update a list of prices in a Price List."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Price List.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostPriceListPricesPricesReq"
//
// x-codegen:
//
//	method: addPrices
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.priceLists.addPrices(priceListId, {
//     prices: [
//     {
//     amount: 1000,
//     variant_id,
//     currency_code: "eur"
//     }
//     ]
//     })
//     .then(({ price_list }) => {
//     console.log(price_list.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreatePriceListPrices } from "medusa-react"
//
//     type PriceData = {
//     amount: number
//     variant_id: string
//     currency_code: string
//     }
//
//     type Props = {
//     priceListId: string
//     }
//
//     const PriceList = ({
//     priceListId
//     }: Props) => {
//     const addPrices = useAdminCreatePriceListPrices(priceListId)
//     // ...
//
//     const handleAddPrices = (prices: PriceData[]) => {
//     addPrices.mutate({
//     prices
//     }, {
//     onSuccess: ({ price_list }) => {
//     console.log(price_list.prices)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default PriceList
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/price-lists/{id}/prices/batch' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "prices": [
//     {
//     "amount": 100,
//     "variant_id": "afasfa",
//     "currency_code": "eur"
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
//   - Price Lists
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPriceListRes"
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
func (m *PriceList) AddPricesBatch(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.AddPriceListPrices](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if _, err := m.r.PriceListService().SetContext(context.Context()).AddPrices(id, model.Prices, model.Override); err != nil {
		return err
	}

	result, err := m.r.PriceListService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [delete] /admin/price-lists/{id}/prices/batch
// operationId: "DeletePriceListsPriceListPricesBatch"
// summary: "Delete Prices"
// description: "Delete a list of prices in a Price List"
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Price List
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminDeletePriceListPricesPricesReq"
//
// x-codegen:
//
//	method: deletePrices
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.priceLists.deletePrices(priceListId, {
//     price_ids: [
//     price_id
//     ]
//     })
//     .then(({ ids, object, deleted }) => {
//     console.log(ids.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeletePriceListPrices } from "medusa-react"
//
//     const PriceList = (
//     priceListId: string
//     ) => {
//     const deletePrices = useAdminDeletePriceListPrices(priceListId)
//     // ...
//
//     const handleDeletePrices = (priceIds: string[]) => {
//     deletePrices.mutate({
//     price_ids: priceIds
//     }, {
//     onSuccess: ({ ids, deleted, object }) => {
//     console.log(ids)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default PriceList
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/price-lists/{id}/prices/batch' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "price_ids": [
//     "adasfa"
//     ]
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Price Lists
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPriceListDeleteBatchRes"
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
func (m *PriceList) DeletePricesBatch(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.DeletePriceListPrices](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if err := m.r.PriceListService().SetContext(context.Context()).DeletePrices(id, model.PriceIds); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"ids":     model.PriceIds,
		"object":  "money-amount",
		"deleted": true,
	})
}

// @oas:path [delete] /admin/price-lists/{id}/products/{product_id}/prices
// operationId: "DeletePriceListsPriceListProductsProductPrices"
// summary: "Delete a Product's Prices"
// description: "Delete all the prices related to a specific product in a price list."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Price List.
//   - (path) product_id=* {string} The ID of the product from which the prices will be deleted.
//
// x-codegen:
//
//	method: deleteProductPrices
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.priceLists.deleteProductPrices(priceListId, productId)
//     .then(({ ids, object, deleted }) => {
//     console.log(ids.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import {
//     useAdminDeletePriceListProductPrices
//     } from "medusa-react"
//
//     type Props = {
//     priceListId: string
//     productId: string
//     }
//
//     const PriceListProduct = ({
//     priceListId,
//     productId
//     }: Props) => {
//     const deleteProductPrices = useAdminDeletePriceListProductPrices(
//     priceListId,
//     productId
//     )
//     // ...
//
//     const handleDeleteProductPrices = () => {
//     deleteProductPrices.mutate(void 0, {
//     onSuccess: ({ ids, deleted, object }) => {
//     console.log(ids)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default PriceListProduct
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/price-lists/{id}/products/{product_id}/prices' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Price Lists
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPriceListDeleteProductPricesRes"
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
func (m *PriceList) DeleteProductPrices(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	productId, err := api.BindDelete(context, "product_id")
	if err != nil {
		return err
	}

	deletedIds, _, err := m.r.PriceListService().SetContext(context.Context()).DeleteProductPrices(id, uuid.UUIDs{productId})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"ids":     deletedIds,
		"object":  "money-amount",
		"deleted": true,
	})
}

// @oas:path [delete] /admin/price-lists/{id}/products/prices/batch
// operationId: "DeletePriceListsPriceListProductsPricesBatch"
// summary: "Delete Product Prices"
// description: "Delete all the prices associated with multiple products in a price list."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Price List
//
// x-codegen:
//
//	method: deleteProductsPrices
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.priceLists.deleteProductsPrices(priceListId, {
//     product_ids: [
//     productId1,
//     productId2,
//     ]
//     })
//     .then(({ ids, object, deleted }) => {
//     console.log(ids.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeletePriceListProductsPrices } from "medusa-react"
//
//     type Props = {
//     priceListId: string
//     }
//
//     const PriceList = ({
//     priceListId
//     }: Props) => {
//     const deleteProductsPrices = useAdminDeletePriceListProductsPrices(
//     priceListId
//     )
//     // ...
//
//     const handleDeleteProductsPrices = (productIds: string[]) => {
//     deleteProductsPrices.mutate({
//     product_ids: productIds
//     }, {
//     onSuccess: ({ ids, deleted, object }) => {
//     console.log(ids)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default PriceList
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/price-lists/{id}/products/prices/batch' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "product_ids": [
//     "prod_1",
//     "prod_2"
//     ]
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Price Lists
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPriceListDeleteProductPricesRes"
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
func (m *PriceList) DeleteProductPricesBatch(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.DeletePriceListPricesBatch](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	deletedIds, _, err := m.r.PriceListService().SetContext(context.Context()).DeleteProductPrices(id, model.ProductIds)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"ids":     deletedIds,
		"object":  "money-amount",
		"deleted": true,
	})
}

// @oas:path [delete] /admin/price-lists/{id}/variants/{variant_id}/prices
// operationId: "DeletePriceListsPriceListVariantsVariantPrices"
// summary: "Delete a Variant's Prices"
// description: "Delete all the prices related to a specific variant in a price list."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Price List.
//   - (path) variant_id=* {string} The ID of the variant.
//
// x-codegen:
//
//	method: deleteVariantPrices
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.priceLists.deleteVariantPrices(priceListId, variantId)
//     .then(({ ids, object, deleted }) => {
//     console.log(ids);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import {
//     useAdminDeletePriceListVariantPrices
//     } from "medusa-react"
//
//     type Props = {
//     priceListId: string
//     variantId: string
//     }
//
//     const PriceListVariant = ({
//     priceListId,
//     variantId
//     }: Props) => {
//     const deleteVariantPrices = useAdminDeletePriceListVariantPrices(
//     priceListId,
//     variantId
//     )
//     // ...
//
//     const handleDeleteVariantPrices = () => {
//     deleteVariantPrices.mutate(void 0, {
//     onSuccess: ({ ids, deleted, object }) => {
//     console.log(ids)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default PriceListVariant
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/price-lists/{id}/variants/{variant_id}/prices' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Price Lists
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPriceListDeleteVariantPricesRes"
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
func (m *PriceList) DeleteVariantPrices(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	variantId, err := api.BindDelete(context, "variant_id")
	if err != nil {
		return err
	}

	deletedIds, _, err := m.r.PriceListService().SetContext(context.Context()).DeleteVariantPrices(id, uuid.UUIDs{variantId})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"ids":     deletedIds,
		"object":  "money-amount",
		"deleted": true,
	})
}

// @oas:path [get] /admin/price-lists/{id}/products
// operationId: "GetPriceListsPriceListProducts"
// summary: "List Products"
// description: "Retrieve a price list's products. The products can be filtered by fields such as `q` or `status`. The products can also be sorted or paginated."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} ID of the price list.
//   - (query) q {string} term used to search products' title, description, product variant's title and sku, and product collection's title.
//   - (query) id {string} Filter by product ID
//   - in: query
//     name: status
//     description: Filter by product status
//     style: form
//     explode: false
//     schema:
//     type: array
//     items:
//     type: string
//     enum: [draft, proposed, published, rejected]
//   - in: query
//     name: collection_id
//     description: Filter by product collection ID. Only products in the specified collections are retrieved.
//     style: form
//     explode: false
//     schema:
//     type: array
//     items:
//     type: string
//   - in: query
//     name: tags
//     description: Filter by tag IDs. Only products having the specified tags are retrieved.
//     style: form
//     explode: false
//     schema:
//     type: array
//     items:
//     type: string
//   - (query) title {string} Filter by title
//   - (query) description {string} Filter by description
//   - (query) handle {string} Filter by handle
//   - (query) is_giftcard {string} A boolean value to filter by whether the product is a gift card or not.
//   - (query) type {string} Filter product type.
//   - (query) order {string} A product field to sort-order the retrieved products by.
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
//   - (query) offset=0 {integer} The number of products to skip when retrieving the products.
//   - (query) limit=50 {integer} Limit the number of products returned.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned products.
//   - (query) fields {string} Comma-separated fields that should be included in the returned products.
//
// x-codegen:
//
//	method: listProducts
//	queryParams: AdminGetPriceListsPriceListProductsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.priceLists.listProducts(priceListId)
//     .then(({ products, limit, offset, count }) => {
//     console.log(products.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminPriceListProducts } from "medusa-react"
//
//     type Props = {
//     priceListId: string
//     }
//
//     const PriceListProducts = ({
//     priceListId
//     }: Props) => {
//     const { products, isLoading } = useAdminPriceListProducts(
//     priceListId
//     )
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {products && !products.length && (
//     <span>No Price Lists</span>
//     )}
//     {products && products.length > 0 && (
//     <ul>
//     {products.map((product) => (
//     <li key={product.id}>{product.title}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default PriceListProducts
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/price-lists/{id}/products' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Price Lists
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPriceListsProductsListRes"
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
func (m *PriceList) ListPriceListProducts(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	model, config, err := api.BindList[types.FilterableProduct](context)
	if err != nil {
		return err
	}

	model.PriceListId = uuid.UUIDs{id}

	result, count, err := m.r.PriceListService().SetContext(context.Context()).ListProducts(id, model, config, false)
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

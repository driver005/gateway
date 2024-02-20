package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type ShippingOption struct {
	r    Registry
	name string
}

func NewShippingOption(r Registry) *ShippingOption {
	m := ShippingOption{r: r, name: "shipping_option"}
	return &m
}

func (m *ShippingOption) SetRoutes(router fiber.Router) {
	route := router.Group("/shipping-options")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)
}

// @oas:path [get] /admin/shipping-options/{id}
// operationId: "GetShippingOptionsOption"
// summary: "Get a Shipping Option"
// description: "Retrieve a Shipping Option's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Shipping Option.
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
//     medusa.admin.shippingOptions.retrieve(optionId)
//     .then(({ shipping_option }) => {
//     console.log(shipping_option.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminShippingOption } from "medusa-react"
//
//     type Props = {
//     shippingOptionId: string
//     }
//
//     const ShippingOption = ({ shippingOptionId }: Props) => {
//     const {
//     shipping_option,
//     isLoading
//     } = useAdminShippingOption(
//     shippingOptionId
//     )
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {shipping_option && <span>{shipping_option.name}</span>}
//     </div>
//     )
//     }
//
//     export default ShippingOption
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/shipping-options/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Shipping Options
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminShippingOptionsRes"
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
func (m *ShippingOption) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.ShippingOptionService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [get] /admin/shipping-options
// operationId: "GetShippingOptions"
// summary: "List Shipping Options"
// description: "Retrieve a list of Shipping Options. The shipping options can be filtered by fields such as `region_id` or `is_return`. The shipping options can also be sorted or paginated."
// x-authenticated: true
// parameters:
//   - (query) name {string} Filter by name.
//   - (query) region_id {string} Filter by the ID of the region the shipping options belong to.
//   - (query) is_return {boolean} Filter by whether the shipping options are return shipping options.
//   - (query) admin_only {boolean} Filter by whether the shipping options are available for admin users only.
//   - (query) q {string} Term used to search shipping options' name.
//   - (query) order {string} A shipping option field to sort-order the retrieved shipping options by.
//   - in: query
//     name: id
//     style: form
//     explode: false
//     description: Filter by shipping option IDs.
//     schema:
//     oneOf:
//   - type: string
//     description: ID of the shipping option.
//   - type: array
//     items:
//     type: string
//     description: ID of a shipping option.
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
//   - (query) offset=0 {integer} The number of users to skip when retrieving the shipping options.
//   - (query) limit=20 {integer} Limit the number of shipping options returned.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned shipping options.
//   - (query) fields {string} Comma-separated fields that should be included in the returned shipping options.
//
// x-codegen:
//
//	method: list
//	queryParams: AdminGetShippingOptionsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.shippingOptions.list()
//     .then(({ shipping_options, count }) => {
//     console.log(shipping_options.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminShippingOptions } from "medusa-react"
//
//     const ShippingOptions = () => {
//     const {
//     shipping_options,
//     isLoading
//     } = useAdminShippingOptions()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {shipping_options && !shipping_options.length && (
//     <span>No Shipping Options</span>
//     )}
//     {shipping_options && shipping_options.length > 0 && (
//     <ul>
//     {shipping_options.map((option) => (
//     <li key={option.id}>{option.name}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default ShippingOptions
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/shipping-options' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Shipping Options
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminShippingOptionsListRes"
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
func (m *ShippingOption) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableShippingOption](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.ShippingOptionService().SetContext(context.Context()).ListAndCount(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"shipping_options": result,
		"count":            count,
		"offset":           config.Skip,
		"limit":            config.Take,
	})
}

// @oas:path [post] /admin/shipping-options
// operationId: "PostShippingOptions"
// summary: "Create Shipping Option"
// description: "Create a Shipping Option."
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostShippingOptionsReq"
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
//     medusa.admin.shippingOptions.create({
//     name: "PostFake",
//     region_id,
//     provider_id,
//     data: {
//     },
//     price_type: "flat_rate"
//     })
//     .then(({ shipping_option }) => {
//     console.log(shipping_option.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreateShippingOption } from "medusa-react"
//
//     type CreateShippingOption = {
//     name: string
//     provider_id: string
//     data: Record<string, unknown>
//     price_type: string
//     amount: number
//     }
//
//     type Props = {
//     regionId: string
//     }
//
//     const Region = ({ regionId }: Props) => {
//     const createShippingOption = useAdminCreateShippingOption()
//     // ...
//
//     const handleCreate = (
//     data: CreateShippingOption
//     ) => {
//     createShippingOption.mutate({
//     ...data,
//     region_id: regionId
//     }, {
//     onSuccess: ({ shipping_option }) => {
//     console.log(shipping_option.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Region
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/shipping-options' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "name": "PostFake",
//     "region_id": "afasf",
//     "provider_id": "manual",
//     "data": {},
//     "price_type": "flat_rate"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Shipping Options
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminShippingOptionsRes"
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
func (m *ShippingOption) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateShippingOptionInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ShippingOptionService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/shipping-options/{id}
// operationId: "PostShippingOptionsOption"
// summary: "Update Shipping Option"
// description: "Update a Shipping Option's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Shipping Option.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostShippingOptionsOptionReq"
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
//     medusa.admin.shippingOptions.update(optionId, {
//     name: "PostFake",
//     requirements: [
//     {
//     id,
//     type: "max_subtotal",
//     amount: 1000
//     }
//     ]
//     })
//     .then(({ shipping_option }) => {
//     console.log(shipping_option.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdateShippingOption } from "medusa-react"
//
//     type Props = {
//     shippingOptionId: string
//     }
//
//     const ShippingOption = ({ shippingOptionId }: Props) => {
//     const updateShippingOption = useAdminUpdateShippingOption(
//     shippingOptionId
//     )
//     // ...
//
//     const handleUpdate = (
//     name: string,
//     requirements: {
//     id: string,
//     type: string,
//     amount: number
//     }[]
//     ) => {
//     updateShippingOption.mutate({
//     name,
//     requirements
//     }, {
//     onSuccess: ({ shipping_option }) => {
//     console.log(shipping_option.requirements)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default ShippingOption
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/shipping-options/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "requirements": [
//     {
//     "type": "max_subtotal",
//     "amount": 1000
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
//   - Shipping Options
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminShippingOptionsRes"
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
func (m *ShippingOption) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateShippingOptionInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ShippingOptionService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/shipping-options/{id}
// operationId: "DeleteShippingOptionsOption"
// summary: "Delete Shipping Option"
// description: "Delete a Shipping Option. Once deleted, it can't be used when creating orders or returns."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Shipping Option.
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
//     medusa.admin.shippingOptions.delete(optionId)
//     .then(({ id, object, deleted }) => {
//     console.log(id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteShippingOption } from "medusa-react"
//
//     type Props = {
//     shippingOptionId: string
//     }
//
//     const ShippingOption = ({ shippingOptionId }: Props) => {
//     const deleteShippingOption = useAdminDeleteShippingOption(
//     shippingOptionId
//     )
//     // ...
//
//     const handleDelete = () => {
//     deleteShippingOption.mutate(void 0, {
//     onSuccess: ({ id, object, deleted }) => {
//     console.log(id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default ShippingOption
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/shipping-options/{option_id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Shipping Options
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminShippingOptionsDeleteRes"
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
func (m *ShippingOption) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.ShippingOptionService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "shipping-option",
		"deleted": true,
	})
}

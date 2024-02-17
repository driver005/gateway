package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type TaxRate struct {
	r Registry
}

func NewTaxRate(r Registry) *TaxRate {
	m := TaxRate{r: r}
	return &m
}

func (m *TaxRate) SetRoutes(router fiber.Router) {
	route := router.Group("/store")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)

	route.Post("/:id/products/batch", m.AddToProducts)
	route.Post("/:id/product-types/batch", m.AddProductTypes)
	route.Post("/:id/shipping-options/batch", m.AddToShippingOptions)
	route.Delete("/:id/products/batch", m.RemoveFromProducts)
	route.Delete("/:id/product-types/batch", m.RemoveFromProductTypes)
	route.Delete("/:id/shipping-options/batch", m.RemoveFromShippingOptions)
}

// @oas:path [get] /admin/tax-rates/{id}
// operationId: "GetTaxRatesTaxRate"
// summary: "Get a Tax Rate"
// description: "Retrieve a Tax Rate's details."
// parameters:
//   - (path) id=* {string} ID of the tax rate.
//   - in: query
//     name: fields
//     description: "Comma-separated fields that should be included in the returned tax rate."
//     style: form
//     explode: false
//     schema:
//     type: array
//     items:
//     type: string
//   - in: query
//     name: expand
//     description: "Comma-separated relations that should be expanded in the returned tax rate."
//     style: form
//     explode: false
//     schema:
//     type: array
//     items:
//     type: string
//
// x-authenticated: true
// x-codegen:
//
//	method: retrieve
//	queryParams: AdminGetTaxRatesTaxRateParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.taxRates.retrieve(taxRateId)
//     .then(({ tax_rate }) => {
//     console.log(tax_rate.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminTaxRate } from "medusa-react"
//
//     type Props = {
//     taxRateId: string
//     }
//
//     const TaxRate = ({ taxRateId }: Props) => {
//     const { tax_rate, isLoading } = useAdminTaxRate(taxRateId)
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {tax_rate && <span>{tax_rate.code}</span>}
//     </div>
//     )
//     }
//
//     export default TaxRate
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/tax-rates/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Tax Rates
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminTaxRatesRes"
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
func (m *TaxRate) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.TaxRateService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [get] /admin/tax-rates
// operationId: "GetTaxRates"
// summary: "List Tax Rates"
// description: "Retrieve a list of Tax Rates. The tax rates can be filtered by fields such as `name` or `rate`. The tax rates can also be paginated."
// x-authenticated: true
// parameters:
//   - (query) name {string} Filter by name.
//   - in: query
//     name: region_id
//     style: form
//     explode: false
//     description: Filter by Region IDs
//     schema:
//     oneOf:
//   - type: string
//   - type: array
//     items:
//     type: string
//   - (query) code {string} Filter by code.
//   - in: query
//     name: rate
//     style: form
//     explode: false
//     description: Filter by Rate
//     schema:
//     oneOf:
//   - type: number
//   - type: object
//     properties:
//     lt:
//     type: number
//     description: filter by rates less than this number
//     gt:
//     type: number
//     description: filter by rates greater than this number
//     lte:
//     type: number
//     description: filter by rates less than or equal to this number
//     gte:
//     type: number
//     description: filter by rates greater than or equal to this number
//   - (query) offset=0 {integer} The number of tax rates to skip when retrieving the tax rates.
//   - (query) limit=50 {integer} Limit the number of tax rates returned.
//   - in: query
//     name: fields
//     description: "Comma-separated fields that should be included in the returned tax rate."
//     style: form
//     explode: false
//     schema:
//     type: array
//     items:
//     type: string
//   - in: query
//     name: expand
//     description: "Comma-separated relations that should be expanded in the returned tax rate."
//     style: form
//     explode: false
//     schema:
//     type: array
//     items:
//     type: string
//
// x-codegen:
//
//	method: list
//	queryParams: AdminGetTaxRatesParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.taxRates.list()
//     .then(({ tax_rates, limit, offset, count }) => {
//     console.log(tax_rates.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminTaxRates } from "medusa-react"
//
//     const TaxRates = () => {
//     const {
//     tax_rates,
//     isLoading
//     } = useAdminTaxRates()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {tax_rates && !tax_rates.length && (
//     <span>No Tax Rates</span>
//     )}
//     {x_rates && tax_rates.length > 0 && (
//     <ul>
//     {tax_rates.map((tax_rate) => (
//     <li key={tax_rate.id}>{tax_rate.code}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default TaxRates
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/tax-rates' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Tax Rates
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminTaxRatesListRes"
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
func (m *TaxRate) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableTaxRate](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.TaxRateService().SetContext(context.Context()).ListAndCount(model, config)
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

// @oas:path [post] /admin/tax-rates
// operationId: "PostTaxRates"
// summary: "Create a Tax Rate"
// description: "Create a Tax Rate."
// parameters:
//   - in: query
//     name: fields
//     description: "Comma-separated fields that should be included in the returned tax rate."
//     style: form
//     explode: false
//     schema:
//     type: array
//     items:
//     type: string
//   - in: query
//     name: expand
//     description: "Comma-separated relations that should be expanded in the returned tax rate."
//     style: form
//     explode: false
//     schema:
//     type: array
//     items:
//     type: string
//
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostTaxRatesReq"
//
// x-codegen:
//
//	method: create
//	queryParams: AdminPostTaxRatesParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.taxRates.create({
//     code: "TEST",
//     name: "New Tax Rate",
//     region_id
//     })
//     .then(({ tax_rate }) => {
//     console.log(tax_rate.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreateTaxRate } from "medusa-react"
//
//     type Props = {
//     regionId: string
//     }
//
//     const CreateTaxRate = ({ regionId }: Props) => {
//     const createTaxRate = useAdminCreateTaxRate()
//     // ...
//
//     const handleCreate = (
//     code: string,
//     name: string,
//     rate: number
//     ) => {
//     createTaxRate.mutate({
//     code,
//     name,
//     region_id: regionId,
//     rate,
//     }, {
//     onSuccess: ({ tax_rate }) => {
//     console.log(tax_rate.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default CreateTaxRate
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/tax-rates' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "code": "TEST",
//     "name": "New Tax Rate",
//     "region_id": "{region_id}"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Tax Rates
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminTaxRatesRes"
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
func (m *TaxRate) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateTaxRateInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.TaxRateService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /admin/tax-rates/{id}
// operationId: "PostTaxRatesTaxRate"
// summary: "Update a Tax Rate"
// description: "Update a Tax Rate's details."
// parameters:
//   - (path) id=* {string} ID of the tax rate.
//   - in: query
//     name: fields
//     description: "Comma-separated fields that should be included in the returned tax rate."
//     style: form
//     explode: false
//     schema:
//     type: array
//     items:
//     type: string
//   - in: query
//     name: expand
//     description: "Comma-separated relations that should be expanded in the returned tax rate."
//     style: form
//     explode: false
//     schema:
//     type: array
//     items:
//     type: string
//
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostTaxRatesTaxRateReq"
//
// x-codegen:
//
//	method: update
//	queryParams: AdminPostTaxRatesTaxRateParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.taxRates.update(taxRateId, {
//     name: "New Tax Rate"
//     })
//     .then(({ tax_rate }) => {
//     console.log(tax_rate.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdateTaxRate } from "medusa-react"
//
//     type Props = {
//     taxRateId: string
//     }
//
//     const TaxRate = ({ taxRateId }: Props) => {
//     const updateTaxRate = useAdminUpdateTaxRate(taxRateId)
//     // ...
//
//     const handleUpdate = (
//     name: string
//     ) => {
//     updateTaxRate.mutate({
//     name
//     }, {
//     onSuccess: ({ tax_rate }) => {
//     console.log(tax_rate.name)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default TaxRate
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/tax-rates/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "name": "New Tax Rate"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Tax Rates
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminTaxRatesRes"
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
func (m *TaxRate) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateTaxRateInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.TaxRateService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [delete] /admin/tax-rates/{id}
// operationId: "DeleteTaxRatesTaxRate"
// summary: "Delete a Tax Rate"
// description: "Delete a Tax Rate. Resources associated with the tax rate, such as products or product types, are not deleted."
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
//     medusa.admin.taxRates.delete(taxRateId)
//     .then(({ id, object, deleted }) => {
//     console.log(id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteTaxRate } from "medusa-react"
//
//     type Props = {
//     taxRateId: string
//     }
//
//     const TaxRate = ({ taxRateId }: Props) => {
//     const deleteTaxRate = useAdminDeleteTaxRate(taxRateId)
//     // ...
//
//     const handleDelete = () => {
//     deleteTaxRate.mutate(void 0, {
//     onSuccess: ({ id, object, deleted }) => {
//     console.log(id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default TaxRate
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/tax-rates/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Tax Rates
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminTaxRatesDeleteRes"
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
func (m *TaxRate) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.TaxRateService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "tax-rate",
		"deleted": true,
	})
}

// @oas:path [post] /admin/tax-rates/{id}/product-types/batch
// operationId: "PostTaxRatesTaxRateProductTypes"
// summary: "Add to Product Types"
// description: "Add Product Types to a Tax Rate."
// parameters:
//   - (path) id=* {string} ID of the tax rate.
//   - in: query
//     name: fields
//     description: "Comma-separated fields that should be included in the returned tax rate."
//     style: form
//     explode: false
//     schema:
//     type: array
//     items:
//     type: string
//   - in: query
//     name: expand
//     description: "Comma-separated relations that should be expanded in the returned tax rate."
//     style: form
//     explode: false
//     schema:
//     type: array
//     items:
//     type: string
//
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostTaxRatesTaxRateProductTypesReq"
//
// x-codegen:
//
//	method: addProductTypes
//	queryParams: AdminPostTaxRatesTaxRateProductTypesParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.taxRates.addProductTypes(taxRateId, {
//     product_types: [
//     productTypeId
//     ]
//     })
//     .then(({ tax_rate }) => {
//     console.log(tax_rate.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import {
//     useAdminCreateProductTypeTaxRates,
//     } from "medusa-react"
//
//     type Props = {
//     taxRateId: string
//     }
//
//     const TaxRate = ({ taxRateId }: Props) => {
//     const addProductTypes = useAdminCreateProductTypeTaxRates(
//     taxRateId
//     )
//     // ...
//
//     const handleAddProductTypes = (productTypeIds: string[]) => {
//     addProductTypes.mutate({
//     product_types: productTypeIds,
//     }, {
//     onSuccess: ({ tax_rate }) => {
//     console.log(tax_rate.product_types)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default TaxRate
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/tax-rates/{id}/product-types/batch' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "product_types": [
//     {product_type_id}"
//     ]
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Tax Rates
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminTaxRatesRes"
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
func (m *TaxRate) AddProductTypes(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.TaxRateProductTypes](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if _, err := m.r.TaxRateService().SetContext(context.Context()).AddToProductType(id, model.ProductTypes, false); err != nil {
		return err
	}

	result, err := m.r.TaxRateService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /admin/tax-rates/{id}/products/batch
// operationId: "PostTaxRatesTaxRateProducts"
// summary: "Add to Products"
// description: "Add products to a tax rate."
// parameters:
//   - (path) id=* {string} ID of the tax rate.
//   - in: query
//     name: fields
//     description: "Comma-separated fields that should be included in the returned tax rate."
//     style: form
//     explode: false
//     schema:
//     type: array
//     items:
//     type: string
//   - in: query
//     name: expand
//     description: "Comma-separated relations that should be expanded in the returned tax rate."
//     style: form
//     explode: false
//     schema:
//     type: array
//     items:
//     type: string
//
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostTaxRatesTaxRateProductsReq"
//
// x-codegen:
//
//	method: addProducts
//	queryParams: AdminPostTaxRatesTaxRateProductsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.taxRates.addProducts(taxRateId, {
//     products: [
//     productId
//     ]
//     })
//     .then(({ tax_rate }) => {
//     console.log(tax_rate.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreateProductTaxRates } from "medusa-react"
//
//     type Props = {
//     taxRateId: string
//     }
//
//     const TaxRate = ({ taxRateId }: Props) => {
//     const addProduct = useAdminCreateProductTaxRates(taxRateId)
//     // ...
//
//     const handleAddProduct = (productIds: string[]) => {
//     addProduct.mutate({
//     products: productIds,
//     }, {
//     onSuccess: ({ tax_rate }) => {
//     console.log(tax_rate.products)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default TaxRate
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/tax-rates/{id}/products/batch' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "products": [
//     {product_id}"
//     ]
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Tax Rates
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminTaxRatesRes"
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
func (m *TaxRate) AddToProducts(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.TaxRateProducts](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if _, err := m.r.TaxRateService().SetContext(context.Context()).AddToProduct(id, model.Products, false); err != nil {
		return err
	}

	result, err := m.r.TaxRateService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /admin/tax-rates/{id}/shipping-options/batch
// operationId: "PostTaxRatesTaxRateShippingOptions"
// summary: "Add to Shipping Options"
// description: "Add Shipping Options to a Tax Rate."
// parameters:
//   - (path) id=* {string} ID of the tax rate.
//   - in: query
//     name: fields
//     description: "Comma-separated fields that should be included in the returned tax rate."
//     style: form
//     explode: false
//     schema:
//     type: array
//     items:
//     type: string
//   - in: query
//     name: expand
//     description: "Comma-separated relations that should be expanded in the returned tax rate."
//     style: form
//     explode: false
//     schema:
//     type: array
//     items:
//     type: string
//
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostTaxRatesTaxRateShippingOptionsReq"
//
// x-codegen:
//
//	method: addShippingOptions
//	queryParams: AdminPostTaxRatesTaxRateShippingOptionsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.taxRates.addShippingOptions(taxRateId, {
//     shipping_options: [
//     shippingOptionId
//     ]
//     })
//     .then(({ tax_rate }) => {
//     console.log(tax_rate.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreateShippingTaxRates } from "medusa-react"
//
//     type Props = {
//     taxRateId: string
//     }
//
//     const TaxRate = ({ taxRateId }: Props) => {
//     const addShippingOption = useAdminCreateShippingTaxRates(
//     taxRateId
//     )
//     // ...
//
//     const handleAddShippingOptions = (
//     shippingOptionIds: string[]
//     ) => {
//     addShippingOption.mutate({
//     shipping_options: shippingOptionIds,
//     }, {
//     onSuccess: ({ tax_rate }) => {
//     console.log(tax_rate.shipping_options)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default TaxRate
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/tax-rates/{id}/shipping-options/batch' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "shipping_options": [
//     {shipping_option_id}"
//     ]
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Tax Rates
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminTaxRatesRes"
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
func (m *TaxRate) AddToShippingOptions(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.TaxRateShippingOptions](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if _, err := m.r.TaxRateService().SetContext(context.Context()).AddToShippingOption(id, model.ShippingOptions, false); err != nil {
		return err
	}

	result, err := m.r.TaxRateService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [delete] /admin/tax-rates/{id}/product-types/batch
// operationId: "DeleteTaxRatesTaxRateProductTypes"
// summary: "Remove Product Types from Rate"
// description: "Remove product types from a tax rate. This only removes the association between the product types and the tax rate. It does not delete the product types."
// parameters:
//   - (path) id=* {string} ID of the tax rate.
//   - in: query
//     name: fields
//     description: "Comma-separated fields that should be included in the returned tax rate."
//     style: form
//     explode: false
//     schema:
//     type: array
//     items:
//     type: string
//   - in: query
//     name: expand
//     description: "Comma-separated relations that should be expanded in the returned tax rate."
//     style: form
//     explode: false
//     schema:
//     type: array
//     items:
//     type: string
//
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminDeleteTaxRatesTaxRateProductTypesReq"
//
// x-codegen:
//
//	method: removeProductTypes
//	queryParams: AdminDeleteTaxRatesTaxRateProductTypesParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.taxRates.removeProductTypes(taxRateId, {
//     product_types: [
//     productTypeId
//     ]
//     })
//     .then(({ tax_rate }) => {
//     console.log(tax_rate.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import {
//     useAdminDeleteProductTypeTaxRates,
//     } from "medusa-react"
//
//     type Props = {
//     taxRateId: string
//     }
//
//     const TaxRate = ({ taxRateId }: Props) => {
//     const removeProductTypes = useAdminDeleteProductTypeTaxRates(
//     taxRateId
//     )
//     // ...
//
//     const handleRemoveProductTypes = (
//     productTypeIds: string[]
//     ) => {
//     removeProductTypes.mutate({
//     product_types: productTypeIds,
//     }, {
//     onSuccess: ({ tax_rate }) => {
//     console.log(tax_rate.product_types)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default TaxRate
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/tax-rates/{id}/product-types/batch' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "product_types": [
//     {product_type_id}"
//     ]
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Tax Rates
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminTaxRatesRes"
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
func (m *TaxRate) RemoveFromProductTypes(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.TaxRateProductTypes](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if err := m.r.TaxRateService().SetContext(context.Context()).RemoveFromProductType(id, model.ProductTypes); err != nil {
		return err
	}

	result, err := m.r.TaxRateService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [delete] /admin/tax-rates/{id}/products/batch
// operationId: "DeleteTaxRatesTaxRateProducts"
// summary: "Remove Products from Rate"
// description: "Remove products from a tax rate. This only removes the association between the products and the tax rate. It does not delete the products."
// parameters:
//   - (path) id=* {string} ID of the tax rate.
//   - in: query
//     name: fields
//     description: "Comma-separated fields that should be included in the returned tax rate."
//     style: form
//     explode: false
//     schema:
//     type: array
//     items:
//     type: string
//   - in: query
//     name: expand
//     description: "Comma-separated relations that should be expanded in the returned tax rate."
//     style: form
//     explode: false
//     schema:
//     type: array
//     items:
//     type: string
//
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminDeleteTaxRatesTaxRateProductsReq"
//
// x-codegen:
//
//	method: removeProducts
//	queryParams: AdminDeleteTaxRatesTaxRateProductsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.taxRates.removeProducts(taxRateId, {
//     products: [
//     productId
//     ]
//     })
//     .then(({ tax_rate }) => {
//     console.log(tax_rate.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteProductTaxRates } from "medusa-react"
//
//     type Props = {
//     taxRateId: string
//     }
//
//     const TaxRate = ({ taxRateId }: Props) => {
//     const removeProduct = useAdminDeleteProductTaxRates(taxRateId)
//     // ...
//
//     const handleRemoveProduct = (productIds: string[]) => {
//     removeProduct.mutate({
//     products: productIds,
//     }, {
//     onSuccess: ({ tax_rate }) => {
//     console.log(tax_rate.products)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default TaxRate
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/tax-rates/{id}/products/batch' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "products": [
//     {product_id}"
//     ]
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Tax Rates
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminTaxRatesRes"
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
func (m *TaxRate) RemoveFromProducts(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.TaxRateProducts](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if err := m.r.TaxRateService().SetContext(context.Context()).RemoveFromProduct(id, model.Products); err != nil {
		return err
	}

	result, err := m.r.TaxRateService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [delete] /admin/tax-rates/{id}/shipping-options/batch
// operationId: "DeleteTaxRatesTaxRateShippingOptions"
// summary: "Remove Shipping Options from Rate"
// description: "Remove shipping options from a tax rate. This only removes the association between the shipping options and the tax rate. It does not delete the shipping options."
// parameters:
//   - (path) id=* {string} ID of the tax rate.
//   - in: query
//     name: fields
//     description: "Comma-separated fields that should be included in the returned tax rate."
//     style: form
//     explode: false
//     schema:
//     type: array
//     items:
//     type: string
//   - in: query
//     name: expand
//     description: "Comma-separated relations that should be expanded in the returned tax rate."
//     style: form
//     explode: false
//     schema:
//     type: array
//     items:
//     type: string
//
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminDeleteTaxRatesTaxRateShippingOptionsReq"
//
// x-codegen:
//
//	method: removeShippingOptions
//	queryParams: AdminDeleteTaxRatesTaxRateShippingOptionsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.taxRates.removeShippingOptions(taxRateId, {
//     shipping_options: [
//     shippingOptionId
//     ]
//     })
//     .then(({ tax_rate }) => {
//     console.log(tax_rate.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteShippingTaxRates } from "medusa-react"
//
//     type Props = {
//     taxRateId: string
//     }
//
//     const TaxRate = ({ taxRateId }: Props) => {
//     const removeShippingOptions = useAdminDeleteShippingTaxRates(
//     taxRateId
//     )
//     // ...
//
//     const handleRemoveShippingOptions = (
//     shippingOptionIds: string[]
//     ) => {
//     removeShippingOptions.mutate({
//     shipping_options: shippingOptionIds,
//     }, {
//     onSuccess: ({ tax_rate }) => {
//     console.log(tax_rate.shipping_options)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default TaxRate
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/tax-rates/{id}/shipping-options/batch' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "shipping_options": [
//     {shipping_option_id}"
//     ]
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Tax Rates
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminTaxRatesRes"
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
func (m *TaxRate) RemoveFromShippingOptions(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.TaxRateShippingOptions](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if err := m.r.TaxRateService().SetContext(context.Context()).RemoveFromShippingOption(id, model.ShippingOptions); err != nil {
		return err
	}

	result, err := m.r.TaxRateService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

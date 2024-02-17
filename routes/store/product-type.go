package store

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type ProductType struct {
	r Registry
}

func NewProductType(r Registry) *ProductType {
	m := ProductType{r: r}
	return &m
}

func (m *ProductType) SetRoutes(router fiber.Router) {
	route := router.Group("/product-types")
	route.Get("", m.List)
}

// @oas:path [get] /store/product-types
// operationId: "GetProductTypes"
// summary: "List Product Types"
// description: "Retrieve a list of product types. The product types can be filtered by fields such as `value` or `q`. The product types can also be sorted or paginated."
// x-authenticated: true
// parameters:
//   - (query) limit=20 {integer} Limit the number of product types returned.
//   - (query) offset=0 {integer} The number of product types to skip when retrieving the product types.
//   - (query) order {string} A product-type field to sort-order the retrieved product types by.
//   - (query) discount_condition_id {string} Filter by the ID of a discount condition. When provided, only types that the discount condition applies for will be retrieved.
//   - in: query
//     name: value
//     style: form
//     explode: false
//     description: Filter by type values.
//     schema:
//     type: array
//     items:
//     type: string
//   - in: query
//     name: id
//     style: form
//     explode: false
//     description: Filter by IDs.
//     schema:
//     type: array
//     items:
//     type: string
//   - (query) q {string} term to search product type's value.
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
//
// x-codegen:
//
//	method: list
//	queryParams: StoreGetProductTypesParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.productTypes.list()
//     .then(({ product_types }) => {
//     console.log(product_types.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useProductTypes } from "medusa-react"
//
//     function Types() {
//     const {
//     product_types,
//     isLoading,
//     } = useProductTypes()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {product_types && !product_types.length && (
//     <span>No Product Types</span>
//     )}
//     {product_types && product_types.length > 0 && (
//     <ul>
//     {product_types.map(
//     (type) => (
//     <li key={type.id}>{type.value}</li>
//     )
//     )}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default Types
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '{backend_url}/store/product-types'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Product Types
//
// responses:
//
//	"200":
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreProductTypesListRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
//	"401":
//	  $ref: "#/components/responses/unauthorized"
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
func (m *ProductType) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableProductType](context)
	if err != nil {
		return err
	}

	result, count, err := m.r.ProductTypeService().SetContext(context.Context()).ListAndCount(model, config)
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

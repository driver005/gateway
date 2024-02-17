package store

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type ProductTag struct {
	r Registry
}

func NewProductTag(r Registry) *ProductTag {
	m := ProductTag{r: r}
	return &m
}

func (m *ProductTag) SetRoutes(router fiber.Router) {
	route := router.Group("/product-tags")
	route.Get("", m.List)
}

// @oas:path [get] /store/product-tags
// operationId: "GetProductTags"
// summary: "List Product Tags"
// description: "Retrieve a list of product tags. The product tags can be filtered by fields such as `id` or `q`. The product tags can also be sorted or paginated."
// x-authenticated: true
// x-codegen:
//
//	method: list
//	queryParams: StoreGetProductTagsParams
//
// parameters:
//   - (query) limit=20 {integer} Limit the number of product tags returned.
//   - (query) offset=0 {integer} The number of product tags to skip when retrieving the product tags.
//   - (query) order {string} A product-tag field to sort-order the retrieved product tags by.
//   - (query) discount_condition_id {string} Filter by the ID of a discount condition. When provided, only tags that the discount condition applies for will be retrieved.
//   - in: query
//     name: value
//     style: form
//     explode: false
//     description: Filter by tag values.
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
//   - (query) q {string} term to search product tag's value.
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
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.productTags.list()
//     .then(({ product_tags }) => {
//     console.log(product_tags.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useProductTags } from "medusa-react"
//
//     function Tags() {
//     const {
//     product_tags,
//     isLoading,
//     } = useProductTags()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {product_tags && !product_tags.length && (
//     <span>No Product Tags</span>
//     )}
//     {product_tags && product_tags.length > 0 && (
//     <ul>
//     {product_tags.map(
//     (tag) => (
//     <li key={tag.id}>{tag.value}</li>
//     )
//     )}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default Tags
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '{backend_url}/store/product-tags'
//
// tags:
//   - Product Tags
//
// responses:
//
//	"200":
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreProductTagsListRes"
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
func (m *ProductTag) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableProductTag](context)
	if err != nil {
		return err
	}

	result, count, err := m.r.ProductTagService().SetContext(context.Context()).ListAndCount(model, config)
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

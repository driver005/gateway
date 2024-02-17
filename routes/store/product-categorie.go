package store

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type ProductCategory struct {
	r Registry
}

func NewProductCategory(r Registry) *ProductCategory {
	m := ProductCategory{r: r}
	return &m
}

func (m *ProductCategory) SetRoutes(router fiber.Router) {
	route := router.Group("/product-categories")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
}

// @oas:path [get] /store/product-categories/{id}
// operationId: "GetProductCategoriesCategory"
// summary: "Get a Product Category"
// description: "Retrieve a Product Category's details."
// x-featureFlag: "product_categories"
// parameters:
//   - (path) id=* {string} The ID of the Product Category
//   - (query) fields {string} Comma-separated fields that should be expanded in the returned product category.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned product category.
//
// x-codegen:
//
//	method: retrieve
//	queryParams: StoreGetProductCategoriesCategoryParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.productCategories.retrieve(productCategoryId)
//     .then(({ product_category }) => {
//     console.log(product_category.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useProductCategory } from "medusa-react"
//
//     type Props = {
//     categoryId: string
//     }
//
//     const Category = ({ categoryId }: Props) => {
//     const { product_category, isLoading } = useProductCategory(
//     categoryId
//     )
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {product_category && <span>{product_category.name}</span>}
//     </div>
//     )
//     }
//
//     export default Category
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '{backend_url}/store/product-categories/{id}' \
//     -H 'x-medusa-access-token: {api_token}'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Product Categories
//
// responses:
//
//	"200":
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreGetProductCategoriesCategoryRes"
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
func (m *ProductCategory) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	result, err := m.r.ProductCategoryService().SetContext(context.Context()).RetrieveById(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [get] /store/product-categories
// operationId: "GetProductCategories"
// summary: "List Product Categories"
// description: "Retrieve a list of product categories. The product categories can be filtered by fields such as `handle` or `q`. The product categories can also be paginated.
//
//	This API Route can also be used to retrieve a product category by its handle."
//
// x-featureFlag: "product_categories"
// externalDocs:
//
//	description: "How to retrieve a product category by its handle"
//	url: "https://docs.medusajs.com/modules/products/storefront/use-categories#get-a-category-by-its-handle"
//
// parameters:
//   - (query) q {string} term used to search product category's names and handles.
//   - (query) handle {string} Filter by handle.
//   - (query) parent_category_id {string} Filter by the ID of a parent category. Only children of the provided parent category are retrieved.
//   - (query) include_descendants_tree {boolean} Whether all nested categories inside a category should be retrieved.
//   - (query) offset=0 {integer} The number of product categories to skip when retrieving the product categories.
//   - (query) limit=100 {integer} Limit the number of product categories returned.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned product categories.
//   - (query) fields {string} Comma-separated fields that should be included in the returned product categories.
//
// x-codegen:
//
//	method: list
//	queryParams: StoreGetProductCategoriesParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.productCategories.list()
//     .then(({ product_categories, limit, offset, count }) => {
//     console.log(product_categories.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useProductCategories } from "medusa-react"
//
//     function Categories() {
//     const {
//     product_categories,
//     isLoading,
//     } = useProductCategories()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {product_categories && !product_categories.length && (
//     <span>No Categories</span>
//     )}
//     {product_categories && product_categories.length > 0 && (
//     <ul>
//     {product_categories.map(
//     (category) => (
//     <li key={category.id}>{category.name}</li>
//     )
//     )}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default Categories
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '{backend_url}/store/product-categories' \
//     -H 'x-medusa-access-token: {api_token}'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Product Categories
//
// responses:
//
//	"200":
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreGetProductCategoriesRes"
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
func (m *ProductCategory) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableProductCategory](context)
	if err != nil {
		return err
	}

	result, count, err := m.r.ProductCategoryService().SetContext(context.Context()).ListAndCount(model, config)
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

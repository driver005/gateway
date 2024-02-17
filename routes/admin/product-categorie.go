package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
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
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)

	route.Post("/:id/products/batch", m.AddProductsBatch)
	route.Delete("/:id/products/batch", m.DeleteProductsBatch)
}

// @oas:path [get] /admin/product-categories/{id}
// operationId: "GetProductCategoriesCategory"
// summary: "Get a Product Category"
// description: "Retrieve a Product Category's details."
// x-authenticated: true
// x-featureFlag: "product_categories"
// parameters:
//   - (path) id=* {string} The ID of the Product Category
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned product category.
//   - (query) fields {string} Comma-separated fields that should be included in the returned product category.
//
// x-codegen:
//
//	method: retrieve
//	queryParams: AdminGetProductCategoryParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.productCategories.retrieve(productCategoryId)
//     .then(({ product_category }) => {
//     console.log(product_category.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminProductCategory } from "medusa-react"
//
//     type Props = {
//     productCategoryId: string
//     }
//
//     const Category = ({
//     productCategoryId
//     }: Props) => {
//     const {
//     product_category,
//     isLoading,
//     } = useAdminProductCategory(productCategoryId)
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {product_category && (
//     <span>{product_category.name}</span>
//     )}
//
//     </div>
//     )
//     }
//
//     export default Category
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/product-categories/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
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
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminProductCategoriesCategoryRes"
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

// @oas:path [get] /admin/product-categories
// operationId: "GetProductCategories"
// summary: "List Product Categories"
// description: "Retrieve a list of product categories. The product categories can be filtered by fields such as `q` or `handle`. The product categories can also be paginated."
// x-authenticated: true
// x-featureFlag: "product_categories"
// parameters:
//   - (query) q {string} term to search product categories' names and handles.
//   - (query) handle {string} Filter by handle.
//   - (query) is_internal {boolean} Filter by whether the category is internal or not.
//   - (query) is_active {boolean} Filter by whether the category is active or not.
//   - (query) include_descendants_tree {boolean} If set to `true`, all nested descendants of a category are included in the response.
//   - (query) parent_category_id {string} Filter by the ID of a parent category.
//   - (query) offset=0 {integer} The number of product categories to skip when retrieving the product categories.
//   - (query) limit=100 {integer} Limit the number of product categories returned.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned product categories.
//   - (query) fields {string} Comma-separated fields that should be included in the returned product categories.
//
// x-codegen:
//
//	method: list
//	queryParams: AdminGetProductCategoriesParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.productCategories.list()
//     .then(({ product_categories, limit, offset, count }) => {
//     console.log(product_categories.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminProductCategories } from "medusa-react"
//
//     function Categories() {
//     const {
//     product_categories,
//     isLoading
//     } = useAdminProductCategories()
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
//     curl '"{backend_url}"/admin/product-categories' \
//     -H 'x-medusa-access-token: "{api_token}"'
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
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminProductCategoriesListRes"
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

// @oas:path [post] /admin/product-categories
// operationId: "PostProductCategories"
// summary: "Create a Product Category"
// description: "Create a Product Category."
// x-authenticated: true
// x-featureFlag: "product_categories"
// parameters:
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned product category.
//   - (query) fields {string} Comma-separated fields that should be included in the returned product category.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostProductCategoriesReq"
//
// x-codegen:
//
//	method: create
//	queryParams: AdminPostProductCategoriesParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.productCategories.create({
//     name: "Skinny Jeans",
//     })
//     .then(({ product_category }) => {
//     console.log(product_category.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreateProductCategory } from "medusa-react"
//
//     const CreateCategory = () => {
//     const createCategory = useAdminCreateProductCategory()
//     // ...
//
//     const handleCreate = (
//     name: string
//     ) => {
//     createCategory.mutate({
//     name,
//     }, {
//     onSuccess: ({ product_category }) => {
//     console.log(product_category.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default CreateCategory
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/product-categories' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "name": "Skinny Jeans"
//     }'
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
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminProductCategoriesCategoryRes"
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
func (m *ProductCategory) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateProductCategoryInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ProductCategoryService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /admin/product-categories/{id}
// operationId: "PostProductCategoriesCategory"
// summary: "Update a Product Category"
// description: "Updates a Product Category."
// x-authenticated: true
// x-featureFlag: "product_categories"
// parameters:
//   - (path) id=* {string} The ID of the Product Category.
//   - (query) expand {string} (Comma separated) Which fields should be expanded in each product category.
//   - (query) fields {string} (Comma separated) Which fields should be retrieved in each product category.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostProductCategoriesCategoryReq"
//
// x-codegen:
//
//	method: update
//	queryParams: AdminPostProductCategoriesCategoryParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.productCategories.update(productCategoryId, {
//     name: "Skinny Jeans"
//     })
//     .then(({ product_category }) => {
//     console.log(product_category.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdateProductCategory } from "medusa-react"
//
//     type Props = {
//     productCategoryId: string
//     }
//
//     const Category = ({
//     productCategoryId
//     }: Props) => {
//     const updateCategory = useAdminUpdateProductCategory(
//     productCategoryId
//     )
//     // ...
//
//     const handleUpdate = (
//     name: string
//     ) => {
//     updateCategory.mutate({
//     name,
//     }, {
//     onSuccess: ({ product_category }) => {
//     console.log(product_category.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Category
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/product-categories/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "name": "Skinny Jeans"
//     }'
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
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminProductCategoriesCategoryRes"
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
func (m *ProductCategory) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateProductCategoryInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ProductCategoryService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [delete] /admin/product-categories/{id}
// operationId: "DeleteProductCategoriesCategory"
// summary: "Delete a Product Category"
// description: "Delete a Product Category. This does not delete associated products."
// x-authenticated: true
// x-featureFlag: "product_categories"
// parameters:
//   - (path) id=* {string} The ID of the Product Category
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
//     medusa.admin.productCategories.delete(productCategoryId)
//     .then(({ id, object, deleted }) => {
//     console.log(id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteProductCategory } from "medusa-react"
//
//     type Props = {
//     productCategoryId: string
//     }
//
//     const Category = ({
//     productCategoryId
//     }: Props) => {
//     const deleteCategory = useAdminDeleteProductCategory(
//     productCategoryId
//     )
//     // ...
//
//     const handleDelete = () => {
//     deleteCategory.mutate(void 0, {
//     onSuccess: ({ id, object, deleted }) => {
//     console.log(id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Category
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/product-categories/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
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
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminProductCategoriesCategoryDeleteRes"
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
func (m *ProductCategory) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.ProductCategoryService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "product-category",
		"deleted": true,
	})
}

// @oas:path [post] /admin/product-categories/{id}/products/batch
// operationId: "PostProductCategoriesCategoryProductsBatch"
// summary: "Add Products to a Category"
// description: "Add a list of products to a product category."
// x-authenticated: true
// x-featureFlag: "product_categories"
// parameters:
//   - (path) id=* {string} The ID of the Product Category.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned product category.
//   - (query) fields {string} Comma-separated fields that should be included in the returned product category.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostProductCategoriesCategoryProductsBatchReq"
//
// x-codegen:
//
//	method: addProducts
//	queryParams: AdminPostProductCategoriesCategoryProductsBatchParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.productCategories.addProducts(productCategoryId, {
//     product_ids: [
//     {
//     id: productId
//     }
//     ]
//     })
//     .then(({ product_category }) => {
//     console.log(product_category.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminAddProductsToCategory } from "medusa-react"
//
//     type ProductsData = {
//     id: string
//     }
//
//     type Props = {
//     productCategoryId: string
//     }
//
//     const Category = ({
//     productCategoryId
//     }: Props) => {
//     const addProducts = useAdminAddProductsToCategory(
//     productCategoryId
//     )
//     // ...
//
//     const handleAddProducts = (
//     productIds: ProductsData[]
//     ) => {
//     addProducts.mutate({
//     product_ids: productIds
//     }, {
//     onSuccess: ({ product_category }) => {
//     console.log(product_category.products)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Category
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/product-categories/{id}/products/batch' \
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
//   - Product Categories
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminProductCategoriesCategoryRes"
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
func (m *ProductCategory) AddProductsBatch(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.AddProductsToCollectionInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	var productIds uuid.UUIDs
	for _, p := range model.ProductIds {
		productIds = append(productIds, p)
	}

	if err := m.r.ProductCategoryService().SetContext(context.Context()).AddProducts(id, productIds); err != nil {
		return err
	}

	result, err := m.r.ProductCategoryService().SetContext(context.Context()).RetrieveById(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [delete] /admin/product-categories/{id}/products/batch
// operationId: "DeleteProductCategoriesCategoryProductsBatch"
// summary: "Remove Products from Category"
// description: "Remove a list of products from a product category."
// x-authenticated: true
// x-featureFlag: "product_categories"
// parameters:
//   - (path) id=* {string} The ID of the Product Category.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned product category.
//   - (query) fields {string} Comma-separated fields that should be included in the returned product category.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminDeleteProductCategoriesCategoryProductsBatchReq"
//
// x-codegen:
//
//	method: removeProducts
//	queryParams: AdminDeleteProductCategoriesCategoryProductsBatchParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.productCategories.removeProducts(productCategoryId, {
//     product_ids: [
//     {
//     id: productId
//     }
//     ]
//     })
//     .then(({ product_category }) => {
//     console.log(product_category.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteProductsFromCategory } from "medusa-react"
//
//     type ProductsData = {
//     id: string
//     }
//
//     type Props = {
//     productCategoryId: string
//     }
//
//     const Category = ({
//     productCategoryId
//     }: Props) => {
//     const deleteProducts = useAdminDeleteProductsFromCategory(
//     productCategoryId
//     )
//     // ...
//
//     const handleDeleteProducts = (
//     productIds: ProductsData[]
//     ) => {
//     deleteProducts.mutate({
//     product_ids: productIds
//     }, {
//     onSuccess: ({ product_category }) => {
//     console.log(product_category.products)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Category
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/product-categories/{id}/products/batch' \
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
//   - Product Categories
//
// responses:
//
//	"200":
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminProductCategoriesCategoryRes"
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
func (m *ProductCategory) DeleteProductsBatch(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.AddProductsToCollectionInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	var productIds uuid.UUIDs
	for _, p := range model.ProductIds {
		productIds = append(productIds, p)
	}

	if err := m.r.ProductCategoryService().SetContext(context.Context()).RemoveProducts(id, productIds); err != nil {
		return err
	}

	result, err := m.r.ProductCategoryService().SetContext(context.Context()).RetrieveById(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

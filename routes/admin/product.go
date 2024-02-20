package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type Product struct {
	r    Registry
	name string
}

func NewProduct(r Registry) *Product {
	m := Product{r: r, name: "product"}
	return &m
}

func (m *Product) SetRoutes(router fiber.Router) {
	route := router.Group("/products")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)

	route.Get("/types", m.ListTypes)
	route.Get("/tag-usage", m.ListTagUsageCount)
	route.Get("/:id/variants", m.ListVariants)
	route.Post("/:id/variants", m.CreateVariant)
	route.Delete("/:id/variants/:variant_id", m.DeletVariant)
	route.Post("/:id/variants/:variant_id", m.UpdateVariant)
	route.Post("/:id/options/:option_id", m.UpdateOption)
	route.Delete("/:id/options/:option_id", m.DeletOption)
	route.Post("/:id/options/", m.AddOption)
	route.Post("/:id/metadata", m.SetMetadata)
}

// @oas:path [get] /admin/products/{id}
// operationId: "GetProductsProduct"
// summary: "Get a Product"
// description: "Retrieve a Product's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Product.
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
//     medusa.admin.products.retrieve(productId)
//     .then(({ product }) => {
//     console.log(product.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminProduct } from "medusa-react"
//
//     type Props = {
//     productId: string
//     }
//
//     const Product = ({ productId }: Props) => {
//     const {
//     product,
//     isLoading,
//     } = useAdminProduct(productId)
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {product && <span>{product.title}</span>}
//
//     </div>
//     )
//     }
//
//     export default Product
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/products/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Products
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminProductsRes"
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
func (m *Product) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.ProductService().SetContext(context.Context()).RetrieveById(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [get] /admin/products
// operationId: "GetProducts"
// summary: "List Products"
// description: "Retrieve a list of products. The products can be filtered by fields such as `q` or `status`. The products can also be sorted or paginated."
// x-authenticated: true
// parameters:
//   - (query) q {string} term to search products' title, description, variants' title and sku, and collections' title.
//   - (query) discount_condition_id {string} Filter by the ID of a discount condition. Only products that this discount condition is applied to will be retrieved.
//   - in: query
//     name: id
//     style: form
//     explode: false
//     description: Filter by product IDs.
//     schema:
//     oneOf:
//   - type: string
//     description: ID of the product.
//   - type: array
//     items:
//     type: string
//     description: ID of a product.
//   - in: query
//     name: status
//     style: form
//     explode: false
//     description: Filter by status.
//     schema:
//     type: array
//     items:
//     type: string
//     enum: [draft, proposed, published, rejected]
//   - in: query
//     name: collection_id
//     style: form
//     explode: false
//     description: Filter by product collection IDs. Only products that are associated with the specified collections will be retrieved.
//     schema:
//     type: array
//     items:
//     type: string
//   - in: query
//     name: tags
//     style: form
//     explode: false
//     description: Filter by product tag IDs. Only products that are associated with the specified tags will be retrieved.
//     schema:
//     type: array
//     items:
//     type: string
//   - in: query
//     name: price_list_id
//     style: form
//     explode: false
//     description: Filter by IDs of price lists. Only products that these price lists are applied to will be retrieved.
//     schema:
//     type: array
//     items:
//     type: string
//   - in: query
//     name: sales_channel_id
//     style: form
//     explode: false
//     description: Filter by sales channel IDs. Only products that are available in the specified sales channels will be retrieved.
//     schema:
//     type: array
//     items:
//     type: string
//   - in: query
//     name: type_id
//     style: form
//     explode: false
//     description: Filter by product type IDs. Only products that are associated with the specified types will be retrieved.
//     schema:
//     type: array
//     items:
//     type: string
//   - in: query
//     name: category_id
//     style: form
//     explode: false
//     description: Filter by product category IDs. Only products that are associated with the specified categories will be retrieved.
//     schema:
//     type: array
//     x-featureFlag: "product_categories"
//     items:
//     type: string
//   - in: query
//     name: include_category_children
//     style: form
//     explode: false
//     description: whether to include product category children when filtering by `category_id`
//     schema:
//     type: boolean
//     x-featureFlag: "product_categories"
//   - (query) title {string} Filter by title.
//   - (query) description {string} Filter by description.
//   - (query) handle {string} Filter by handle.
//   - (query) is_giftcard {boolean} Whether to retrieve gift cards or regular products.
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
//   - (query) order {string} A product field to sort-order the retrieved products by.
//
// x-codegen:
//
//	method: list
//	queryParams: AdminGetProductsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.products.list()
//     .then(({ products, limit, offset, count }) => {
//     console.log(products.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminProducts } from "medusa-react"
//
//     const Products = () => {
//     const { products, isLoading } = useAdminProducts()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {products && !products.length && <span>No Products</span>}
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
//     export default Products
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/products' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Products
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminProductsListRes"
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
func (m *Product) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableProduct](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.ProductService().SetContext(context.Context()).ListAndCount(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"products": result,
		"count":    count,
		"offset":   config.Skip,
		"limit":    config.Take,
	})
}

// @oas:path [post] /admin/products
// operationId: "PostProducts"
// summary: "Create a Product"
// x-authenticated: true
// description: "Create a new Product. This API Route can also be used to create a gift card if the `is_giftcard` field is set to `true`."
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostProductsReq"
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
//     medusa.admin.products.create({
//     title: "Shirt",
//     is_giftcard: false,
//     discountable: true
//     })
//     .then(({ product }) => {
//     console.log(product.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreateProduct } from "medusa-react"
//
//     type CreateProductData = {
//     title: string
//     is_giftcard: boolean
//     discountable: boolean
//     options: {
//     title: string
//     }[]
//     variants: {
//     title: string
//     prices: {
//     amount: number
//     currency_code :string
//     }[]
//     options: {
//     value: string
//     }[]
//     }[],
//     collection_id: string
//     categories: {
//     id: string
//     }[]
//     type: {
//     value: string
//     }
//     tags: {
//     value: string
//     }[]
//     }
//
//     const CreateProduct = () => {
//     const createProduct = useAdminCreateProduct()
//     // ...
//
//     const handleCreate = (productData: CreateProductData) => {
//     createProduct.mutate(productData, {
//     onSuccess: ({ product }) => {
//     console.log(product.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default CreateProduct
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/products' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "title": "Shirt"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Products
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminProductsRes"
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
func (m *Product) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateProductInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ProductService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/products/{id}
// operationId: "PostProductsProduct"
// summary: "Update a Product"
// description: "Update a Product's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Product.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostProductsProductReq"
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
//     medusa.admin.products.update(productId, {
//     title: "Shirt",
//     })
//     .then(({ product }) => {
//     console.log(product.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdateProduct } from "medusa-react"
//
//     type Props = {
//     productId: string
//     }
//
//     const Product = ({ productId }: Props) => {
//     const updateProduct = useAdminUpdateProduct(
//     productId
//     )
//     // ...
//
//     const handleUpdate = (
//     title: string
//     ) => {
//     updateProduct.mutate({
//     title,
//     }, {
//     onSuccess: ({ product }) => {
//     console.log(product.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Product
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/products/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "title": "Size"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Products
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminProductsRes"
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
func (m *Product) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateProductInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ProductService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/products/{id}
// operationId: "DeleteProductsProduct"
// summary: "Delete a Product"
// description: "Delete a Product and its associated product variants and options."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Product.
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
//     medusa.admin.products.delete(productId)
//     .then(({ id, object, deleted }) => {
//     console.log(id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteProduct } from "medusa-react"
//
//     type Props = {
//     productId: string
//     }
//
//     const Product = ({ productId }: Props) => {
//     const deleteProduct = useAdminDeleteProduct(
//     productId
//     )
//     // ...
//
//     const handleDelete = () => {
//     deleteProduct.mutate(void 0, {
//     onSuccess: ({ id, object, deleted}) => {
//     console.log(id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Product
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/products/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Products
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminProductsDeleteRes"
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
func (m *Product) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.ProductService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "product",
		"deleted": true,
	})
}

// @oas:path [post] /admin/products/{id}/options
// operationId: "PostProductsProductOptions"
// summary: "Add a Product Option"
// description: "Add a Product Option to a Product."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Product.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostProductsProductOptionsReq"
//
// x-codegen:
//
//	method: addOption
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.products.addOption(productId, {
//     title: "Size"
//     })
//     .then(({ product }) => {
//     console.log(product.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreateProductOption } from "medusa-react"
//
//     type Props = {
//     productId: string
//     }
//
//     const CreateProductOption = ({ productId }: Props) => {
//     const createOption = useAdminCreateProductOption(
//     productId
//     )
//     // ...
//
//     const handleCreate = (
//     title: string
//     ) => {
//     createOption.mutate({
//     title
//     }, {
//     onSuccess: ({ product }) => {
//     console.log(product.options)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default CreateProductOption
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/products/{id}/options' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "title": "Size"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Products
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminProductsRes"
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
func (m *Product) AddOption(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.CreateProductProductOption](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if _, err := m.r.ProductService().SetContext(context.Context()).AddOption(id, model.Title); err != nil {
		return err
	}

	result, err := m.r.ProductService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/products/{id}/options/{option_id}
// operationId: "DeleteProductsProductOptionsOption"
// summary: "Delete a Product Option"
// description: "Delete a Product Option. If there are product variants that use this product option, they must be deleted before deleting the product option."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Product.
//   - (path) option_id=* {string} The ID of the Product Option.
//
// x-codegen:
//
//	method: deleteOption
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.products.deleteOption(productId, optionId)
//     .then(({ option_id, object, deleted, product }) => {
//     console.log(product.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteProductOption } from "medusa-react"
//
//     type Props = {
//     productId: string
//     optionId: string
//     }
//
//     const ProductOption = ({
//     productId,
//     optionId
//     }: Props) => {
//     const deleteOption = useAdminDeleteProductOption(
//     productId
//     )
//     // ...
//
//     const handleDelete = () => {
//     deleteOption.mutate(optionId, {
//     onSuccess: ({ option_id, object, deleted, product }) => {
//     console.log(product.options)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default ProductOption
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/products/{id}/options/{option_id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Products
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminProductsDeleteOptionRes"
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
func (m *Product) DeletOption(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	optionId, err := api.BindDelete(context, "option_id")
	if err != nil {
		return err
	}

	if _, err := m.r.ProductService().SetContext(context.Context()).DeleteOption(id, optionId); err != nil {
		return err
	}

	result, err := m.r.ProductService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"option_id": optionId,
		"object":    "option",
		"deleted":   true,
		"product":   result,
	})
}

// @oas:path [post] /admin/products/{id}/variants
// operationId: "PostProductsProductVariants"
// summary: "Create a Product Variant"
// description: "Create a Product Variant associated with a Product. Each product variant must have a unique combination of Product Option values."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Product.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostProductsProductVariantsReq"
//
// x-codegen:
//
//	method: createVariant
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.products.createVariant(productId, {
//     title: "Color",
//     prices: [
//     {
//     amount: 1000,
//     currency_code: "eur"
//     }
//     ],
//     options: [
//     {
//     option_id,
//     value: "S"
//     }
//     ],
//     inventory_quantity: 100
//     })
//     .then(({ product }) => {
//     console.log(product.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreateVariant } from "medusa-react"
//
//     type CreateVariantData = {
//     title: string
//     prices: {
//     amount: number
//     currency_code: string
//     }[]
//     options: {
//     option_id: string
//     value: string
//     }[]
//     }
//
//     type Props = {
//     productId: string
//     }
//
//     const CreateProductVariant = ({ productId }: Props) => {
//     const createVariant = useAdminCreateVariant(
//     productId
//     )
//     // ...
//
//     const handleCreate = (
//     variantData: CreateVariantData
//     ) => {
//     createVariant.mutate(variantData, {
//     onSuccess: ({ product }) => {
//     console.log(product.variants)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default CreateProductVariant
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/products/{id}/variants' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "title": "Color",
//     "prices": [
//     {
//     "amount": 1000,
//     "currency_code": "eur"
//     }
//     ],
//     "options": [
//     {
//     "option_id": "asdasf",
//     "value": "S"
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
//   - Products
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminProductsRes"
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
func (m *Product) CreateVariant(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.CreateProductVariantInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if m.r.InventoryService() != nil {
		variants, err := m.r.ProductVariantService().SetContext(context.Context()).Create(id, nil, []types.CreateProductVariantInput{*model})
		if err != nil {
			return err
		}

		for _, variant := range variants {
			if !variant.ManageInventory {
				continue
			}

			inventoryItem, err := m.r.InventoryService().CreateInventoryItem(context.Context(), interfaces.CreateInventoryItemInput{
				SKU:           variant.Sku,
				OriginCountry: variant.OriginCountry,
				HsCode:        variant.HsCode,
				MidCode:       variant.MIdCode,
				Material:      variant.Material,
				Weight:        variant.Weight,
				Length:        variant.Length,
				Height:        variant.Height,
				Width:         variant.Width,
			})
			if err != nil {
				return err
			}

			if _, err := m.r.ProductVariantInventoryService().AttachInventoryItem([]models.ProductVariantInventoryItem{
				{
					VariantId:       uuid.NullUUID{UUID: variant.Id},
					InventoryItemId: uuid.NullUUID{UUID: inventoryItem.Id},
				},
			}); err != nil {
				return err
			}
		}
	} else {
		_, err := m.r.ProductVariantService().SetContext(context.Context()).Create(id, nil, []types.CreateProductVariantInput{*model})
		if err != nil {
			return err
		}
	}

	raw, err := m.r.ProductService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err := m.r.PricingService().SetContext(context.Context()).SetAdminProductPricing([]models.Product{*raw})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/products/{id}/variants/{variant_id}
// operationId: "DeleteProductsProductVariantsVariant"
// summary: "Delete a Product Variant"
// description: "Delete a Product Variant."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Product.
//   - (path) variant_id=* {string} The ID of the Product Variant.
//
// x-codegen:
//
//	method: deleteVariant
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.products.deleteVariant(productId, variantId)
//     .then(({ variant_id, object, deleted, product }) => {
//     console.log(product.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteVariant } from "medusa-react"
//
//     type Props = {
//     productId: string
//     variantId: string
//     }
//
//     const ProductVariant = ({
//     productId,
//     variantId
//     }: Props) => {
//     const deleteVariant = useAdminDeleteVariant(
//     productId
//     )
//     // ...
//
//     const handleDelete = () => {
//     deleteVariant.mutate(variantId, {
//     onSuccess: ({ variant_id, object, deleted, product }) => {
//     console.log(product.variants)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default ProductVariant
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/products/{id}/variants/{variant_id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Products
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminProductsDeleteVariantRes"
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
func (m *Product) DeletVariant(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	variantId, err := api.BindDelete(context, "variant_id")
	if err != nil {
		return err
	}

	if err := m.r.ProductVariantService().SetContext(context.Context()).Delete(uuid.UUIDs{variantId}); err != nil {
		return err
	}

	raw, err := m.r.ProductService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err := m.r.PricingService().SetContext(context.Context()).SetAdminProductPricing([]models.Product{*raw})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"option_id": variantId,
		"object":    "product-variant",
		"deleted":   true,
		"product":   result,
	})
}

// @oas:path [get] /admin/products/tag-usage
// operationId: "GetProductsTagUsage"
// summary: "List Tags Usage Number"
// description: "Retrieve a list of Product Tags with how many times each is used in products."
// x-authenticated: true
// x-codegen:
//
//	method: listTags
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.products.listTags()
//     .then(({ tags }) => {
//     console.log(tags.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminProductTagUsage } from "medusa-react"
//
//     const ProductTags = (productId: string) => {
//     const { tags, isLoading } = useAdminProductTagUsage()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {tags && !tags.length && <span>No Product Tags</span>}
//     {tags && tags.length > 0 && (
//     <ul>
//     {tags.map((tag) => (
//     <li key={tag.id}>{tag.value} - {tag.usage_count}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default ProductTags
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/products/tag-usage' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Products
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminProductsListTagsRes"
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
func (m *Product) ListTagUsageCount(context fiber.Ctx) error {
	result, err := m.r.ProductService().SetContext(context.Context()).ListTagsByUsage(-1)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"tags": result,
	})
}

// @oas:path [get] /admin/products/types
// deprecated: true
// operationId: "GetProductsTypes"
// summary: "List Product Types"
// description: "Retrieve a list of Product Types."
// x-authenticated: true
// x-codegen:
//
//	method: listTypes
//
// x-codeSamples:
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.products.listTypes()
//     .then(({ types }) => {
//     console.log(types.length);
//     })
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/products/types' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Products
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminProductsListTypesRes"
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
func (m *Product) ListTypes(context fiber.Ctx) error {
	result, err := m.r.ProductService().SetContext(context.Context()).ListTypes()
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"types": result,
	})
}

// @oas:path [get] /admin/products/{id}/variants
// operationId: "GetProductsProductVariants"
// summary: "List a Product's Variants"
// description: |
//
//	Retrieve a list of Product Variants associated with a Product. The variants can be paginated.
//
//	By default, each variant will only have the `id` and `variant_id` fields. You can use the `expand` and `fields` request parameters to retrieve more fields or relations.
//
// x-authenticated: true
// parameters:
//   - (path) id=* {string} ID of the product.
//   - (query) fields {string} Comma-separated fields that should be included in the returned product variants.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned product variants.
//   - (query) offset=0 {integer} The number of product variants to skip when retrieving the product variants.
//   - (query) limit=100 {integer} Limit the number of product variants returned.
//
// x-codegen:
//
//	method: listVariants
//	queryParams: AdminGetProductsVariantsParams
//
// x-codeSamples:
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/products/{id}/variants' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Products
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminProductsListVariantsRes"
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
func (m *Product) ListVariants(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableProductVariant](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.ProductVariantService().SetContext(context.Context()).ListAndCount(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"variants": result,
		"count":    count,
		"offset":   config.Skip,
		"limit":    config.Take,
	})
}

// @oas:path [post] /admin/products/{id}/metadata
// operationId: "PostProductsProductMetadata"
// summary: "Set Metadata"
// description: "Set the metadata of a Product. It can be any key-value pair, which allows adding custom data to a product."
// externalDocs:
//
//	description: "Learn about the metadata attribute, and how to delete and update it."
//	url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
//
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Product.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostProductsProductMetadataReq"
//
// x-codegen:
//
//	method: setMetadata
//
// x-codeSamples:
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.products.setMetadata(productId, {
//     key: "test",
//     value: "true"
//     })
//     .then(({ product }) => {
//     console.log(product.id);
//     })
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/products/{id}/metadata' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "key": "test",
//     "value": "true"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Products
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminProductsRes"
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
func (m *Product) SetMetadata(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.Metadata](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	data := &types.UpdateProductInput{}
	data.Metadata = data.Metadata.Add(model.Key, model.Value)

	if _, err := m.r.ProductService().SetContext(context.Context()).Update(id, data); err != nil {
		return err
	}

	raw, err := m.r.ProductService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err := m.r.PricingService().SetContext(context.Context()).SetAdminProductPricing([]models.Product{*raw})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/products/{id}/options/{option_id}
// operationId: "PostProductsProductOptionsOption"
// summary: "Update a Product Option"
// description: "Update a Product Option's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Product.
//   - (path) option_id=* {string} The ID of the Product Option.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostProductsProductOptionsOption"
//
// x-codegen:
//
//	method: updateOption
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.products.updateOption(productId, optionId, {
//     title: "Size"
//     })
//     .then(({ product }) => {
//     console.log(product.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdateProductOption } from "medusa-react"
//
//     type Props = {
//     productId: string
//     optionId: string
//     }
//
//     const ProductOption = ({
//     productId,
//     optionId
//     }: Props) => {
//     const updateOption = useAdminUpdateProductOption(
//     productId
//     )
//     // ...
//
//     const handleUpdate = (
//     title: string
//     ) => {
//     updateOption.mutate({
//     option_id: optionId,
//     title,
//     }, {
//     onSuccess: ({ product }) => {
//     console.log(product.options)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default ProductOption
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/products/{id}/options/{option_id}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "title": "Size"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Products
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminProductsRes"
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
func (m *Product) UpdateOption(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.CreateProductProductOption](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	optionId, err := api.BindDelete(context, "option_id")
	if err != nil {
		return err
	}

	if _, err := m.r.ProductService().SetContext(context.Context()).UpdateOption(id, optionId, &types.ProductOptionInput{
		Title: model.Title,
	}); err != nil {
		return err
	}

	raw, err := m.r.ProductService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err := m.r.PricingService().SetContext(context.Context()).SetAdminProductPricing([]models.Product{*raw})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/products/{id}/variants/{variant_id}
// operationId: "PostProductsProductVariantsVariant"
// summary: "Update a Product Variant"
// description: "Update a Product Variant's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Product.
//   - (path) variant_id=* {string} The ID of the Product Variant.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostProductsProductVariantsVariantReq"
//
// x-codegen:
//
//	method: updateVariant
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.products.updateVariant(productId, variantId, {
//     title: "Color",
//     prices: [
//     {
//     amount: 1000,
//     currency_code: "eur"
//     }
//     ],
//     options: [
//     {
//     option_id,
//     value: "S"
//     }
//     ],
//     inventory_quantity: 100
//     })
//     .then(({ product }) => {
//     console.log(product.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdateVariant } from "medusa-react"
//
//     type Props = {
//     productId: string
//     variantId: string
//     }
//
//     const ProductVariant = ({
//     productId,
//     variantId
//     }: Props) => {
//     const updateVariant = useAdminUpdateVariant(
//     productId
//     )
//     // ...
//
//     const handleUpdate = (title: string) => {
//     updateVariant.mutate({
//     variant_id: variantId,
//     title,
//     }, {
//     onSuccess: ({ product }) => {
//     console.log(product.variants)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default ProductVariant
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/products/{id}/variants/{variant_id}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "title": "Color",
//     "prices": [
//     {
//     "amount": 1000,
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
//   - Products
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminProductsRes"
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
func (m *Product) UpdateVariant(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateProductVariantInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	variantId, err := api.BindDelete(context, "variant_id")
	if err != nil {
		return err
	}

	model.ProductId = id

	if _, err := m.r.ProductVariantService().SetContext(context.Context()).Update(variantId, nil, model); err != nil {
		return err
	}

	raw, err := m.r.ProductService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err := m.r.PricingService().SetContext(context.Context()).SetAdminProductPricing([]models.Product{*raw})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})

}

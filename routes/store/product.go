package store

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/services"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Product struct {
	r Registry
}

func NewProduct(r Registry) *Product {
	m := Product{r: r}
	return &m
}

func (m *Product) SetRoutes(router fiber.Router) {
	route := router.Group("/regions")
	route.Get("/:id", m.Get)
	route.Get("", m.List)

	route.Post("", m.Search)
}

// @oas:path [get] /store/products/{id}
// operationId: GetProductsProduct
// summary: Get a Product
// description: |
//
//	Retrieve a Product's details. For accurate and correct pricing of the product based on the customer's context, it's highly recommended to pass fields such as
//	`region_id`, `currency_code`, and `cart_id` when available.
//
//	Passing `sales_channel_id` ensures retrieving only products available in the current sales channel.
//	You can alternatively use a publishable API key in the request header instead of passing a `sales_channel_id`.
//
// externalDocs:
//
//	description: "How to pass product pricing parameters"
//	url: "https://docs.medusajs.com/modules/products/storefront/show-products#product-pricing-parameters"
//
// parameters:
//   - (path) id=* {string} The ID of the Product.
//   - (query) sales_channel_id {string} The ID of the sales channel the customer is viewing the product from.
//   - (query) cart_id {string} The ID of the cart. This is useful for accurate pricing based on the cart's context.
//   - (query) region_id {string} The ID of the region. This is useful for accurate pricing based on the selected region.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned product.
//   - (query) fields {string} Comma-separated fields that should be included in the returned product.
//   - in: query
//     name: currency_code
//     style: form
//     explode: false
//     description: A 3 character ISO currency code. This is useful for accurate pricing based on the selected currency.
//     schema:
//     type: string
//     externalDocs:
//     url: https://en.wikipedia.org/wiki/ISO_4217#Active_codes
//     description: See a list of codes.
//
// x-codegen:
//
//	method: retrieve
//	queryParams: StoreGetProductsProductParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.products.retrieve(productId)
//     .then(({ product }) => {
//     console.log(product.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useProduct } from "medusa-react"
//
//     type Props = {
//     productId: string
//     }
//
//     const Product = ({ productId }: Props) => {
//     const { product, isLoading } = useProduct(productId)
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {product && <span>{product.title}</span>}
//     </div>
//     )
//     }
//
//     export default Product
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '{backend_url}/store/products/{id}'
//
// tags:
//   - Products
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreProductsRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
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
	model, id, config, err := api.BindAll[types.ProductVariantVariant](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	customerId := context.Locals("customer_id").(uuid.UUID)

	product, err := m.r.ProductService().SetContext(context.Context()).RetrieveById(id, config)
	if err != nil {
		return err
	}

	salesChannelId := model.SalesChannelId
	publishableApiKeyScopes, ok := context.Locals("publishableApiKeyScopes").(types.PublishableApiKeyScopes)
	if ok {
		salesChannelId = publishableApiKeyScopes.SalesChannelIds[0]
	}

	regionId := model.RegionId
	currencyCode := model.CurrencyCode
	if model.CartId != uuid.Nil {
		cart, err := m.r.CartService().SetContext(context.Context()).Retrieve(model.CartId, &sql.Options{Selects: []string{"id", "region_id"}}, services.TotalsConfig{})
		if err != nil {
			return err
		}
		region, err := m.r.RegionService().SetContext(context.Context()).Retrieve(cart.RegionId.UUID, &sql.Options{Selects: []string{"id", "currency_code"}})
		if err != nil {
			return err
		}
		regionId = region.Id
		currencyCode = region.CurrencyCode
	}

	result := []models.Product{*product}

	if lo.ContainsBy(config.Relations, func(item string) bool {
		return item == "variants" || item == "variants.prices"
	}) {
		result, err = m.r.PricingService().SetContext(context.Context()).SetProductPrices(result, &interfaces.PricingContext{
			CartId:                model.CartId,
			CustomerId:            customerId,
			RegionId:              regionId,
			CurrencyCode:          currencyCode,
			IncludeDiscountPrices: true,
		})
		if err != nil {
			return err
		}
	}

	if lo.ContainsBy(config.Relations, func(item string) bool {
		return item == "variants"
	}) {
		result, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetProductAvailability(result, uuid.UUIDs{salesChannelId})
		if err != nil {
			return err
		}
	}

	//TODO: Result only variant not list
	return context.Status(fiber.StatusOK).JSON(result[0])
}

// @oas:path [get] /store/products
// operationId: GetProducts
// summary: List Products
// description: |
//
//	Retrieves a list of products. The products can be filtered by fields such as `id` or `q`. The products can also be sorted or paginated.
//	This API Route can also be used to retrieve a product by its handle.
//
//	For accurate and correct pricing of the products based on the customer's context, it's highly recommended to pass fields such as
//	`region_id`, `currency_code`, and `cart_id` when available.
//
//	Passing `sales_channel_id` ensures retrieving only products available in the specified sales channel.
//	You can alternatively use a publishable API key in the request header instead of passing a `sales_channel_id`.
//
// externalDocs:
//
//	description: "How to retrieve a product by its handle"
//	url: "https://docs.medusajs.com/modules/products/storefront/show-products#retrieve-product-by-handle"
//
// parameters:
//   - (query) q {string} term used to search products' title, description, variant's title, variant's sku, and collection's title.
//   - in: query
//     name: id
//     style: form
//     explode: false
//     description: Filter by IDs.
//     schema:
//     oneOf:
//   - type: string
//   - type: array
//     items:
//     type: string
//   - in: query
//     name: sales_channel_id
//     style: form
//     explode: false
//     description: "Filter by sales channel IDs. When provided, only products available in the selected sales channels are retrieved. Alternatively, you can pass a
//     publishable API key in the request header and this will have the same effect."
//     schema:
//     type: array
//     items:
//     type: string
//   - in: query
//     name: collection_id
//     style: form
//     explode: false
//     description: Filter by product collection IDs. When provided, only products that belong to the specified product collections are retrieved.
//     schema:
//     type: array
//     items:
//     type: string
//   - in: query
//     name: type_id
//     style: form
//     explode: false
//     description: Filter by product type IDs. When provided, only products that belong to the specified product types are retrieved.
//     schema:
//     type: array
//     items:
//     type: string
//   - in: query
//     name: tags
//     style: form
//     explode: false
//     description: Filter by product tag IDs. When provided, only products that belong to the specified product tags are retrieved.
//     schema:
//     type: array
//     items:
//     type: string
//   - (query) title {string} Filter by title.
//   - (query) description {string} Filter by description
//   - (query) handle {string} Filter by handle.
//   - (query) is_giftcard {boolean} Whether to retrieve regular products or gift-card products.
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
//     name: category_id
//     style: form
//     explode: false
//     description: Filter by product category IDs. When provided, only products that belong to the specified product categories are retrieved.
//     schema:
//     type: array
//     x-featureFlag: "product_categories"
//     items:
//     type: string
//   - in: query
//     name: include_category_children
//     style: form
//     explode: false
//     description: Whether to include child product categories when filtering using the `category_id` field.
//     schema:
//     type: boolean
//     x-featureFlag: "product_categories"
//   - (query) offset=0 {integer} The number of products to skip when retrieving the products.
//   - (query) limit=100 {integer} Limit the number of products returned.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned products.
//   - (query) fields {string} Comma-separated fields that should be included in the returned products.
//   - (query) order {string} A product field to sort-order the retrieved products by.
//   - (query) cart_id {string} The ID of the cart. This is useful for accurate pricing based on the cart's context.
//   - (query) region_id {string} The ID of the region. This is useful for accurate pricing based on the selected region.
//   - in: query
//     name: currency_code
//     style: form
//     explode: false
//     description: A 3 character ISO currency code. This is useful for accurate pricing based on the selected currency.
//     schema:
//     type: string
//     externalDocs:
//     url: https://en.wikipedia.org/wiki/ISO_4217#Active_codes
//     description: See a list of codes.
//
// x-codegen:
//
//	method: list
//	queryParams: StoreGetProductsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.products.list()
//     .then(({ products, limit, offset, count }) => {
//     console.log(products.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useProducts } from "medusa-react"
//
//     const Products = () => {
//     const { products, isLoading } = useProducts()
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
//     curl '{backend_url}/store/products'
//
// tags:
//   - Products
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreProductsListRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
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

	customerId := context.Locals("customer_id").(uuid.UUID)
	regionId := model.RegionId
	currencyCode := model.CurrencyCode

	model.Status = []models.ProductStatus{models.ProductStatusPublished}
	publishableApiKeyScopes, ok := context.Locals("publishableApiKeyScopes").(types.PublishableApiKeyScopes)
	if ok {
		if model.SalesChannelId == nil {
			model.SalesChannelId = publishableApiKeyScopes.SalesChannelIds
		}

		if !lo.Contains(config.Relations, "listConfig.relations") {
			config.Relations = append(config.Relations, "sales_channels")
		}
	}

	result, count, err := m.r.ProductService().SetContext(context.Context()).ListAndCount(model, config)
	if err != nil {
		return err
	}

	cart := &models.Cart{}

	if model.CartId != uuid.Nil {
		cart, err = m.r.CartService().SetContext(context.Context()).Retrieve(model.CartId, &sql.Options{Selects: []string{"id", "region_id"}, Relations: []string{"region"}}, services.TotalsConfig{})
		if err != nil {
			return err
		}

		regionId = cart.RegionId.UUID
		currencyCode = cart.Region.CurrencyCode
	}

	if lo.ContainsBy(config.Relations, func(item string) bool {
		return item == "variants" || item == "variants.prices"
	}) {
		result, err = m.r.PricingService().SetContext(context.Context()).SetProductPrices(result, &interfaces.PricingContext{
			CartId:                model.CartId,
			CustomerId:            customerId,
			RegionId:              regionId,
			CurrencyCode:          currencyCode,
			IncludeDiscountPrices: true,
		})
		if err != nil {
			return err
		}
	}

	if lo.ContainsBy(config.Relations, func(item string) bool {
		return item == "variants"
	}) {
		result, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetProductAvailability(result, model.SalesChannelId)
		if err != nil {
			return err
		}
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   result,
		"count":  count,
		"offset": config.Skip,
		"limit":  config.Take,
	})
}

// @oas:path [post] /store/products/search
// operationId: PostProductsSearch
// summary: Search Products
// description: "Run a search query on products using the search service installed on the Medusa backend. The searching is handled through the search service, so the returned data's
//
//	format depends on the search service you're using."
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/StorePostSearchReq"
//
// x-codegen:
//
//	method: search
//
// x-codeSamples:
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.products.search({
//     q: "Shirt"
//     })
//     .then(({ hits }) => {
//     console.log(hits.length);
//     })
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/store/products/search' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "q": "Shirt"
//     }'
//
// tags:
//   - Products
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StorePostSearchRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
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
func (m *Product) Search(context fiber.Ctx) error {
	model, err := api.Bind[types.ProductSearch](context, m.r.Validator())
	if err != nil {
		return err
	}

	result := m.r.DefaultSearchService().SetContext(context.Context()).Search("product", model.Q, map[string]interface{}{
		"limit":  model.Limit,
		"offset": model.Offset,
		"filter": model.Filter,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

package store

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/services"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type Variant struct {
	r Registry
}

func NewVariant(r Registry) *Variant {
	m := Variant{r: r}
	return &m
}

func (m *Variant) SetRoutes(router fiber.Router) {
	route := router.Group("/variants")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
}

// @oas:path [get] /store/variants/{id}
// operationId: GetVariantsVariant
// summary: Get a Product Variant
// description: |
//
//	Retrieve a Product Variant's details. For accurate and correct pricing of the product variant based on the customer's context, it's highly recommended to pass fields such as
//	`region_id`, `currency_code`, and `cart_id` when available.
//
//	Passing `sales_channel_id` ensures retrieving only variants of products available in the current sales channel.
//	You can alternatively use a publishable API key in the request header instead of passing a `sales_channel_id`.
//
// externalDocs:
//
//	description: "How to pass product pricing parameters"
//	url: "https://docs.medusajs.com/modules/products/storefront/show-products#product-pricing-parameters"
//
// parameters:
//   - (path) id=* {string} The ID of the Product Variant.
//   - (query) sales_channel_id {string} The ID of the sales channel the customer is viewing the product variant from.
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
//	method: retrieve
//	queryParams: StoreGetVariantsVariantParams
//
// x-codeSamples:
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.product.variants.retrieve(productVariantId)
//     .then(({ variant }) => {
//     console.log(variant.id);
//     })
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '{backend_url}/store/variants/{id}'
//
// tags:
//   - Product Variants
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreVariantsRes"
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
func (m *Variant) Get(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.ProductVariantVariant](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	customerId := context.Locals("customer_id").(uuid.UUID)

	variant, err := m.r.ProductVariantService().SetContext(context.Context()).Retrieve(id, config)
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

	prices, err := m.r.PricingService().SetContext(context.Context()).SetVariantPrices([]models.ProductVariant{*variant}, &interfaces.PricingContext{
		CartId:                model.CartId,
		CustomerId:            customerId,
		RegionId:              regionId,
		CurrencyCode:          currencyCode,
		IncludeDiscountPrices: true,
	})
	if err != nil {
		return err
	}

	result, err := m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(prices, uuid.UUIDs{salesChannelId}, &services.AvailabilityContext{})
	if err != nil {
		return err
	}

	//TODO: Result only variant not list
	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [get] /store/variants
// operationId: GetVariants
// summary: Get Product Variants
// description: |
//
//	Retrieves a list of product variants. The product variants can be filtered by fields such as `id` or `title`. The product variants can also be paginated.
//
//	For accurate and correct pricing of the product variants based on the customer's context, it's highly recommended to pass fields such as
//	`region_id`, `currency_code`, and `cart_id` when available.
//
//	Passing `sales_channel_id` ensures retrieving only variants of products available in the specified sales channel.
//	You can alternatively use a publishable API key in the request header instead of passing a `sales_channel_id`.
//
// externalDocs:
//
//	description: "How to pass product pricing parameters"
//	url: "https://docs.medusajs.com/modules/products/storefront/show-products#product-pricing-parameters"
//
// parameters:
//   - (query) ids {string} Filter by a comma-separated list of IDs. If supplied, it overrides the `id` parameter.
//   - in: query
//     name: id
//     style: form
//     explode: false
//     description: Filter by one or more IDs. If `ids` is supplied, it's overrides the value of this parameter.
//     schema:
//     oneOf:
//   - type: string
//     description: Filter by an ID.
//   - type: array
//     description: Filter by IDs.
//     items:
//     type: string
//   - (query) sales_channel_id {string} "Filter by sales channel IDs. When provided, only products available in the selected sales channels are retrieved. Alternatively, you can pass a
//     publishable API key in the request header and this will have the same effect."
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned product variants.
//   - (query) fields {string} Comma-separated fields that should be included in the returned product variants.
//   - (query) offset=0 {number} The number of products to skip when retrieving the product variants.
//   - (query) limit=100 {number} Limit the number of product variants returned.
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
//   - in: query
//     name: title
//     style: form
//     explode: false
//     description: Filter by title
//     schema:
//     oneOf:
//   - type: string
//     description: a single title to filter by
//   - type: array
//     description: multiple titles to filter by
//     items:
//     type: string
//   - in: query
//     name: inventory_quantity
//     description: Filter by available inventory quantity
//     schema:
//     oneOf:
//   - type: number
//     description: A specific number to filter by.
//   - type: object
//     description: Filter using less and greater than comparisons.
//     properties:
//     lt:
//     type: number
//     description: Filter by inventory quantity less than this number
//     gt:
//     type: number
//     description: Filter by inventory quantity greater than this number
//     lte:
//     type: number
//     description: Filter by inventory quantity less than or equal to this number
//     gte:
//     type: number
//     description: Filter by inventory quantity greater than or equal to this number
//
// x-codegen:
//
//	method: list
//	queryParams: StoreGetVariantsParams
//
// x-codeSamples:
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.product.variants.list()
//     .then(({ variants }) => {
//     console.log(variants.length);
//     })
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '{backend_url}/store/variants'
//
// tags:
//   - Product Variants
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreVariantsListRes"
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
func (m *Variant) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.ProductVariantParams](context)
	if err != nil {
		return err
	}

	customerId := context.Locals("customer_id").(uuid.UUID)
	publishableApiKeyScopes, ok := context.Locals("publishableApiKeyScopes").(types.PublishableApiKeyScopes)
	if ok {
		model.SalesChannelId = publishableApiKeyScopes.SalesChannelIds[0]
	}

	variants, err := m.r.ProductVariantService().SetContext(context.Context()).List(&types.FilterableProductVariant{
		FilterModel: core.FilterModel{
			Id: []uuid.UUID{model.Id},
		},
		Title:             model.Title,
		InventoryQuantity: model.InventoryQuantity,
	}, config)
	if err != nil {
		return err
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

	prices, err := m.r.PricingService().SetContext(context.Context()).SetVariantPrices(variants, &interfaces.PricingContext{
		CartId:                model.CartId,
		CustomerId:            customerId,
		RegionId:              regionId,
		CurrencyCode:          currencyCode,
		IncludeDiscountPrices: true,
	})
	if err != nil {
		return err
	}

	result, err := m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(prices, uuid.UUIDs{model.SalesChannelId}, &services.AvailabilityContext{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

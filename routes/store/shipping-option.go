package store

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/services"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"golang.org/x/exp/maps"
)

type ShippingOption struct {
	r Registry
}

func NewShippingOption(r Registry) *ShippingOption {
	m := ShippingOption{r: r}
	return &m
}

func (m *ShippingOption) SetRoutes(router fiber.Router) {
	route := router.Group("/shipping-options")
	route.Get("", m.ListOptions)
	route.Get("/:cart_id", m.List)
}

// @oas:path [get] /store/shipping-options/{cart_id}
// operationId: GetShippingOptionsCartId
// summary: List for Cart
// description: "Retrieve a list of Shipping Options available for a cart."
// externalDocs:
//
//	description: "How to implement shipping step in checkout"
//	url: "https://docs.medusajs.com/modules/carts-and-checkout/storefront/implement-checkout-flow#shipping-step"
//
// parameters:
//   - (path) cart_id {string} The ID of the Cart.
//
// x-codegen:
//
//	method: listCartOptions
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.shippingOptions.listCartOptions(cartId)
//     .then(({ shipping_options }) => {
//     console.log(shipping_options.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useCartShippingOptions } from "medusa-react"
//
//     type Props = {
//     cartId: string
//     }
//
//     const ShippingOptions = ({ cartId }: Props) => {
//     const { shipping_options, isLoading } =
//     useCartShippingOptions(cartId)
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {shipping_options && !shipping_options.length && (
//     <span>No shipping options</span>
//     )}
//     {shipping_options && (
//     <ul>
//     {shipping_options.map(
//     (shipping_option) => (
//     <li key={shipping_option.id}>
//     {shipping_option.name}
//     </li>
//     )
//     )}
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
//     curl '{backend_url}/store/shipping-options/{cart_id}'
//
// tags:
//   - Shipping Options
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCartShippingOptionsListRes"
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
func (m *ShippingOption) List(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "cart_id")
	if err != nil {
		return err
	}
	cart, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	options, err := m.r.ShippingProfileService().SetContext(context.Context()).FetchCartOptions(cart)
	if err != nil {
		return err
	}

	result, err := m.r.PricingService().SetContext(context.Context()).SetShippingOptionPrices(options, &interfaces.PricingContext{
		CartId: id,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"shipping_options": result,
	})
}

// @oas:path [get] /store/shipping-options
// operationId: GetShippingOptions
// summary: Get Shipping Options
// description: "Retrieve a list of Shipping Options."
// parameters:
//   - (query) is_return {boolean} Whether return shipping options should be included. By default, all shipping options are returned.
//   - (query) product_ids {string} "Comma-separated list of Product IDs to filter Shipping Options by. If provided, only shipping options that can be used with the provided products are retrieved."
//   - (query) region_id {string} "The ID of the region that the shipping options belong to. If not provided, all shipping options are retrieved."
//
// x-codegen:
//
//	method: list
//	queryParams: StoreGetShippingOptionsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.shippingOptions.list()
//     .then(({ shipping_options }) => {
//     console.log(shipping_options.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useShippingOptions } from "medusa-react"
//
//     const ShippingOptions = () => {
//     const {
//     shipping_options,
//     isLoading,
//     } = useShippingOptions()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {shipping_options?.length &&
//     shipping_options?.length > 0 && (
//     <ul>
//     {shipping_options?.map((shipping_option) => (
//     <li key={shipping_option.id}>
//     {shipping_option.id}
//     </li>
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
//     curl '{backend_url}/store/shipping-options'
//
// tags:
//   - Shipping Options
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreShippingOptionsListRes"
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
func (m *ShippingOption) ListOptions(context fiber.Ctx) error {
	model, config, err := api.BindList[types.ShippingOptionParams](context)
	if err != nil {
		return err
	}

	query := &types.FilterableShippingOption{}

	if context.Query("is_return") != "" {
		query.IsReturn = model.IsReturn == "true"
	}

	if model.RegionId != uuid.Nil {
		query.RegionId = model.RegionId
	}

	query.AdminOnly = false

	if len(model.ProductIds) > 0 {
		productShippinProfileMap, err := m.r.ShippingProfileService().SetContext(context.Context()).GetMapProfileIdsByProductIds(model.ProductIds)

		query.ProfileId = maps.Values(*productShippinProfileMap)
		if err != nil {
			return err
		}
	}

	result, err := m.r.ShippingOptionService().SetContext(context.Context()).List(query, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"shipping_options": result,
	})
}

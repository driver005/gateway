package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/services"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type ResponseInventoryItem struct {
	interfaces.InventoryItemDTO
	LocationLevels   []interfaces.InventoryLevelDTO `json:"location_levels,omitempty"`
	StockedQuantity  int
	ReservedQuantity int
}

type SalesChannelAvailability struct {
	ChannelId         uuid.UUID `json:"channel_id"`
	ChannelName       string    `json:"channel_name"`
	AvailableQuantity int       `json:"available_quantity"`
}

type VariantInventory struct {
	Id                       uuid.UUID                  `json:"id"`
	Inventory                []ResponseInventoryItem    `json:"inventory"`
	SalesChannelAvailability []SalesChannelAvailability `json:"sales_channel_availability"`
}

type AdminGetVariantsParams struct {
	types.FilterableProductVariant
	CartId       uuid.UUID `json:"cart_id,omitempty" validate:"omitempty"`
	RegionId     uuid.UUID `json:"region_id,omitempty" validate:"omitempty"`
	CurrencyCode string    `json:"currency_code,omitempty" validate:"omitempty"`
	CustomerId   uuid.UUID `json:"customer_id,omitempty" validate:"omitempty"`
}

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

	route.Post("/:id/inventory", m.GetInventory)
}

// @oas:path [get] /admin/variants/{id}
// operationId: "GetVariantsVariant"
// summary: "Get a Product variant"
// description: "Retrieve a product variant's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the product variant.
//   - (query) expand {string} "Comma-separated relations that should be expanded in the returned product variant."
//   - (query) fields {string} "Comma-separated fields that should be included in the returned product variant."
//
// x-codegen:
//
//	method: retrieve
//	queryParams: AdminGetVariantParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.variants.retrieve(variantId)
//     .then(({ variant }) => {
//     console.log(variant.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminVariant } from "medusa-react"
//
//     type Props = {
//     variantId: string
//     }
//
//     const Variant = ({ variantId }: Props) => {
//     const { variant, isLoading } = useAdminVariant(
//     variantId
//     )
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {variant && <span>{variant.title}</span>}
//     </div>
//     )
//     }
//
//     export default Variant
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/variants/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Product Variants
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminVariantsRes"
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
func (m *Variant) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	rawVariant, err := m.r.ProductVariantService().Retrieve(id, config)
	if err != nil {
		return err
	}

	variant, err := m.r.PricingService().SetAdminVariantPricing([]models.ProductVariant{*rawVariant}, &interfaces.PricingContext{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(variant)
}

// @oas:path [get] /admin/variants
// operationId: "GetVariants"
// summary: "List Product Variants"
// description: "Retrieve a list of Product Variants. The product variant can be filtered by fields such as `id` or `title`. The product variant can also be paginated."
// x-authenticated: true
// parameters:
//   - in: query
//     name: id
//     style: form
//     explode: false
//     description: Filter by product variant IDs.
//     schema:
//     oneOf:
//   - type: string
//     description: A product variant ID.
//   - type: array
//     description: An array of product variant IDs.
//     items:
//     type: string
//   - (query) expand {string} "Comma-separated relations that should be expanded in the returned product variants."
//   - (query) fields {string} "Comma-separated fields that should be included in the returned product variants."
//   - (query) offset=0 {number} The number of product variants to skip when retrieving the product variants.
//   - (query) limit=100 {number} Limit the number of product variants returned.
//   - in: query
//     name: cart_id
//     style: form
//     explode: false
//     description: The ID of the cart to use for the price selection context.
//     schema:
//     type: string
//   - in: query
//     name: region_id
//     style: form
//     explode: false
//     description: The ID of the region to use for the price selection context.
//     schema:
//     type: string
//     externalDocs:
//     description: "Price selection context overview"
//     url: "https://docs.medusajs.com/modules/price-lists/price-selection-strategy#context-object"
//   - in: query
//     name: currency_code
//     style: form
//     explode: false
//     description: The 3 character ISO currency code to use for the price selection context.
//     schema:
//     type: string
//     externalDocs:
//     description: "Price selection context overview"
//     url: "https://docs.medusajs.com/modules/price-lists/price-selection-strategy#context-object"
//   - in: query
//     name: customer_id
//     style: form
//     explode: false
//     description: The ID of the customer to use for the price selection context.
//     schema:
//     type: string
//     externalDocs:
//     description: "Price selection context overview"
//     url: "https://docs.medusajs.com/modules/price-lists/price-selection-strategy#context-object"
//   - in: query
//     name: title
//     style: form
//     explode: false
//     description: Filter by title.
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
//     description: a specific number to filter by.
//   - type: object
//     description: filter using less and greater than comparisons.
//     properties:
//     lt:
//     type: number
//     description: filter by inventory quantity less than this number
//     gt:
//     type: number
//     description: filter by inventory quantity greater than this number
//     lte:
//     type: number
//     description: filter by inventory quantity less than or equal to this number
//     gte:
//     type: number
//     description: filter by inventory quantity greater than or equal to this number
//
// x-codegen:
//
//	method: list
//	queryParams: AdminGetVariantsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.variants.list()
//     .then(({ variants, limit, offset, count }) => {
//     console.log(variants.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminVariants } from "medusa-react"
//
//     const Variants = () => {
//     const { variants, isLoading } = useAdminVariants()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {variants && !variants.length && (
//     <span>No Variants</span>
//     )}
//     {variants && variants.length > 0 && (
//     <ul>
//     {variants.map((variant) => (
//     <li key={variant.id}>{variant.title}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default Variants
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/variants' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Product Variants
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminVariantsListRes"
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
func (m *Variant) List(context fiber.Ctx) error {
	model, config, err := api.BindList[AdminGetVariantsParams](context)
	if err != nil {
		return err
	}

	rawVariants, count, err := m.r.ProductVariantService().ListAndCount(&types.FilterableProductVariant{
		Title:             model.Title,
		ProductId:         model.ProductId,
		Product:           model.Product,
		SKU:               model.SKU,
		Barcode:           model.Barcode,
		EAN:               model.EAN,
		UPC:               model.UPC,
		InventoryQuantity: model.InventoryQuantity,
		AllowBackorder:    model.AllowBackorder,
		ManageInventory:   model.ManageInventory,
		HSCode:            model.HSCode,
		OriginCountry:     model.OriginCountry,
		MidCode:           model.MidCode,
		Material:          model.Material,
		Weight:            model.Weight,
		Length:            model.Length,
		Height:            model.Height,
		Width:             model.Width,
	}, config)
	if err != nil {
		return err
	}

	regionId := model.RegionId
	currencyCode := model.CurrencyCode
	if model.CartId != uuid.Nil {
		cart, err := m.r.CartService().Retrieve(model.CartId, &sql.Options{
			Selects: []string{"id", "region_id"},
		}, services.TotalsConfig{})
		if err != nil {
			return err
		}
		region, err := m.r.RegionService().Retrieve(cart.RegionId.UUID, &sql.Options{
			Selects: []string{"id", "currency_code"},
		})
		if err != nil {
			return err
		}
		regionId = region.Id
		currencyCode = region.CurrencyCode
	}

	variants, err := m.r.PricingService().SetAdminVariantPricing(rawVariants, &interfaces.PricingContext{
		CartId:                model.CartId,
		RegionId:              regionId,
		CurrencyCode:          currencyCode,
		CustomerId:            model.CustomerId,
		IncludeDiscountPrices: true,
		IgnoreCache:           true,
	})
	if err != nil {
		return err
	}

	if m.r.InventoryService() != nil {
		salesChannelsIds, err := m.r.SalesChannelService().List(&types.FilterableSalesChannel{}, &sql.Options{
			Selects: []string{"id"},
		})
		if err != nil {
			return err
		}
		v, err := m.r.ProductVariantInventoryService().SetVariantAvailability(variants, func() uuid.UUIDs {
			var ids uuid.UUIDs
			for _, salesChannel := range salesChannelsIds {
				ids = append(ids, salesChannel.Id)
			}
			return ids
		}(), &services.AvailabilityContext{})
		if err != nil {
			return err
		}

		variants = v
	}

	return context.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"variants": variants,
		"count":    count,
		"offset":   config.Skip,
		"limit":    config.Take,
	})

}

// @oas:path [get] /admin/variants/{id}/inventory
// operationId: "GetVariantsVariantInventory"
// summary: "Get Variant's Inventory"
// description: "Retrieve the available inventory of a Product Variant."
// x-authenticated: true
// parameters:
//   - (path) id {string} The Product Variant ID.
//
// x-codegen:
//
//	method: getInventory
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.variants.getInventory(variantId)
//     .then(({ variant }) => {
//     console.log(variant.inventory, variant.sales_channel_availability)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminVariantsInventory } from "medusa-react"
//
//     type Props = {
//     variantId: string
//     }
//
//     const VariantInventory = ({ variantId }: Props) => {
//     const { variant, isLoading } = useAdminVariantsInventory(
//     variantId
//     )
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {variant && variant.inventory.length === 0 && (
//     <span>Variant doesn't have inventory details</span>
//     )}
//     {variant && variant.inventory.length > 0 && (
//     <ul>
//     {variant.inventory.map((inventory) => (
//     <li key={inventory.id}>{inventory.title}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default VariantInventory
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/variants/{id}/inventory' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Product Variants
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminGetVariantsVariantInventoryRes"
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
func (m *Variant) GetInventory(context fiber.Ctx) error {
	id, err := utils.ParseUUID(context.Params("id"))
	if err != nil {
		return err
	}

	variant, _ := m.r.ProductVariantService().Retrieve(id, &sql.Options{Selects: []string{"id"}})
	responseVariant := VariantInventory{
		Id:                       variant.Id,
		Inventory:                []ResponseInventoryItem{},
		SalesChannelAvailability: []SalesChannelAvailability{},
	}

	rawChannels, _, err := m.r.SalesChannelService().ListAndCount(&types.FilterableSalesChannel{}, &sql.Options{})
	if err != nil {
		return err
	}
	var channels []models.SalesChannel
	for _, channel := range rawChannels {
		locationIds, _ := m.r.SalesChannelLocationService().ListLocationIds(uuid.UUIDs{channel.Id})
		channels = append(channels, models.SalesChannel{
			Model: core.Model{
				Id: channel.Id,
			},
			Name:        channel.Name,
			LocationIds: locationIds,
		})
	}

	variantInventoryItems, err := m.r.ProductVariantInventoryService().ListByVariant(uuid.UUIDs{variant.Id})
	if err != nil {
		return err
	}
	inventory, _ := m.r.ProductVariantInventoryService().ListInventoryItemsByVariant(variant.Id)
	in, err := joinLevels(inventory, uuid.UUIDs{}, m.r.InventoryService())
	if err != nil {
		return err
	}
	responseVariant.Inventory = in

	if len(inventory) > 0 {
		for _, channel := range channels {
			if len(channel.LocationIds) == 0 {
				responseVariant.SalesChannelAvailability = append(responseVariant.SalesChannelAvailability, SalesChannelAvailability{
					ChannelName:       channel.Name,
					ChannelId:         channel.Id,
					AvailableQuantity: 0,
				})
			} else {
				quantity, err := m.r.ProductVariantInventoryService().GetVariantQuantityFromVariantInventoryItems(variantInventoryItems, channel.Id)
				if err != nil {
					return err
				}
				responseVariant.SalesChannelAvailability = append(responseVariant.SalesChannelAvailability, SalesChannelAvailability{
					ChannelName:       channel.Name,
					ChannelId:         channel.Id,
					AvailableQuantity: *quantity,
				})
			}
		}
	}

	return context.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"variant": responseVariant,
	})
}

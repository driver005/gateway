package store

import (
	"reflect"
	"strings"

	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/services"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Cart struct {
	r Registry
}

func NewCart(r Registry) *Cart {
	m := Cart{r: r}
	return &m
}

func (m *Cart) SetRoutes(router fiber.Router) {
	route := router.Group("/gift-cards")
	route.Get("/:id", m.Get)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)

	route.Post("/:id/complete", m.Complete)
	route.Post("/:id/line-items", m.CreateLineItem)
	route.Post("/:id/line-items/:line_id", m.UpdateLineItem)
	route.Delete("/:id/line-items/:line_id", m.DeleteLineItem)
	route.Post("/:id/payment-session", m.SetPaymentSession)
	route.Post("/:id/payment-sessions", m.CreatePaymentSession)
	route.Post("/:id/payment-sessions/:provider_id", m.UpdatePaymentSession)
	route.Delete("/:id/payment-sessions/:provider_id", m.DeletePaymentSession)
	route.Post("/:id/payment-sessions/:provider_id/refresh", m.RefreshPaymentSession)
	route.Post("/:id/shipping-methods", m.AddShippingMethod)
	route.Post("/:id/taxes", m.CalculateTaxes)
	route.Delete("/:id/discounts/:code", m.DeleteDiscount)
}

// @oas:path [get] /store/carts/{id}
// operationId: "GetCartsCart"
// summary: "Get a Cart"
// description: "Retrieve a Cart's details. This includes recalculating its totals."
// parameters:
//   - (path) id=* {string} The ID of the Cart.
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
//     medusa.carts.retrieve(cartId)
//     .then(({ cart }) => {
//     console.log(cart.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useGetCart } from "medusa-react"
//
//     type Props = {
//     cartId: string
//     }
//
//     const Cart = ({ cartId }: Props) => {
//     const { cart, isLoading } = useGetCart(cartId)
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {cart && cart.items.length === 0 && (
//     <span>Cart is empty</span>
//     )}
//     {cart && cart.items.length > 0 && (
//     <ul>
//     {cart.items.map((item) => (
//     <li key={item.id}>{item.title}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default Cart
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '{backend_url}/store/carts/{id}'
//
// tags:
//   - Carts
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCartsRes"
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
func (m *Cart) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	cart, err := m.r.CartService().SetContext(context.Context()).Retrieve(id, &sql.Options{Selects: []string{"id", "customer_id"}}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	customerId := context.Locals("customer_id").(uuid.UUID)

	if context.Locals("user") != nil && customerId != uuid.Nil {
		if cart.CustomerId.UUID == uuid.Nil || cart.Email == "" || cart.CustomerId.UUID != customerId {
			if _, err := m.r.CartService().SetContext(context.Context()).Update(id, nil, &types.CartUpdateProps{CustomerId: customerId}); err != nil {
				return err
			}
		}
	}

	config.Selects = append(config.Selects, "sales_channel_id")

	result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, config, services.TotalsConfig{})
	if err != nil {
		return err
	}

	if lo.Contains(config.Relations, "variant") {
		var variants []models.ProductVariant
		for _, item := range result.Items {
			variants = append(variants, *item.Variant)
		}
		if _, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(variants, uuid.UUIDs{result.SalesChannelId.UUID}, &services.AvailabilityContext{}); err != nil {
			return err
		}
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /store/carts
// operationId: "PostCart"
// summary: "Create a Cart"
// description: |
//
//	Create a Cart. Although optional, specifying the cart's region and sales channel can affect the cart's pricing and
//	the products that can be added to the cart respectively. So, make sure to set those early on and change them if necessary, such as when the customer changes their region.
//
//	If a customer is logged in, make sure to pass its ID or email within the cart's details so that the cart is attached to the customer.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/StorePostCartReq"
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
//     medusa.carts.create()
//     .then(({ cart }) => {
//     console.log(cart.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useCreateCart } from "medusa-react"
//
//     type Props = {
//     regionId: string
//     }
//
//     const Cart = ({ regionId }: Props) => {
//     const createCart = useCreateCart()
//
//     const handleCreate = () => {
//     createCart.mutate({
//     region_id: regionId
//     // creates an empty cart
//     }, {
//     onSuccess: ({ cart }) => {
//     console.log(cart.items)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Cart
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/store/carts'
//
// tags:
//   - Carts
//
// responses:
//
//	200:
//	  description: "Successfully created a new Cart"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCartsRes"
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
func (m *Cart) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateCart](context, m.r.Validator())
	if err != nil {
		return err
	}

	var regionId uuid.UUID
	if reflect.ValueOf(model.RegionId).IsZero() {
		regionId = model.RegionId
	} else {
		regions, err := m.r.CartService().SetContext(context.Context()).List(types.FilterableCartProps{}, &sql.Options{})
		if err != nil {
			return err
		}

		if len(regions) == 0 {
			return utils.NewApplictaionError(
				utils.INVALID_DATA,
				"A region is required to create a cart",
			)
		}

		regionId = regions[0].RegionId.UUID
	}

	model.Context = utils.MergeMaps(model.Context, map[string]interface{}{
		"ip":         context.IP(),
		"user_agent": context.Get("user-agent"),
	})

	data := &types.CartCreateProps{
		RegionId:       regionId,
		SalesChannelId: model.SalesChannelId,
		Context:        model.Context,
	}

	customerId := context.Locals("customer_id").(uuid.UUID)

	if customerId != uuid.Nil {
		customer, err := m.r.CustomerService().SetContext(context.Context()).RetrieveById(customerId, &sql.Options{})
		if err != nil {
			return err
		}

		data.CustomerId = customer.Id
		data.Email = customer.Email
	}

	if model.CountryCode != "" {
		data.ShippingAddress = &types.AddressPayload{
			CountryCode: strings.ToLower(model.CountryCode),
		}
	}

	publishableApiKeyScopes, ok := context.Locals("publishableApiKeyScopes").(types.PublishableApiKeyScopes)
	if data.SalesChannelId != uuid.Nil && ok {
		if len(publishableApiKeyScopes.SalesChannelIds) > 1 {
			return utils.NewApplictaionError(
				utils.UNEXPECTED_STATE,
				"The PublishableApiKey provided in the request header has multiple associated sales channels.",
			)
		}

		data.SalesChannelId = publishableApiKeyScopes.SalesChannelIds[0]
	}

	cart, err := m.r.CartService().SetContext(context.Context()).Create(data)
	if err != nil {
		return err
	}

	if len(model.Items) != 0 {
		var generateInputData []types.GenerateInputData
		for _, item := range model.Items {
			generateInputData = append(generateInputData, types.GenerateInputData{
				VariantId: item.VariantId,
				Quantity:  item.Quantity,
			})
		}

		//TODO: Check Quantity
		generatedLineItems, err := m.r.LineItemService().SetContext(context.Context()).Generate(uuid.Nil, generateInputData, uuid.Nil, 0, types.GenerateLineItemContext{
			RegionId:   regionId,
			CustomerId: customerId,
		})
		if err != nil {
			return err
		}

		if err := m.r.CartService().SetContext(context.Context()).AddOrUpdateLineItems(cart.Id, generatedLineItems, true); err != nil {
			return err
		}
	}

	result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(cart.Id, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	var variants []models.ProductVariant
	for _, item := range result.Items {
		variants = append(variants, *item.Variant)
	}
	if _, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(variants, uuid.UUIDs{result.SalesChannelId.UUID}, &services.AvailabilityContext{}); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /store/carts/{id}
// operationId: PostCartsCart
// summary: Update a Cart
// description: "Update a Cart's details. If the cart has payment sessions and the region was not changed, the payment sessions are updated. The cart's totals are also recalculated."
// parameters:
//   - (path) id=* {string} The ID of the Cart.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/StorePostCartsCartReq"
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
//     medusa.carts.update(cartId, {
//     email: "user@example.com"
//     })
//     .then(({ cart }) => {
//     console.log(cart.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useUpdateCart } from "medusa-react"
//
//     type Props = {
//     cartId: string
//     }
//
//     const Cart = ({ cartId }: Props) => {
//     const updateCart = useUpdateCart(cartId)
//
//     const handleUpdate = (
//     email: string
//     ) => {
//     updateCart.mutate({
//     email
//     }, {
//     onSuccess: ({ cart }) => {
//     console.log(cart.email)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Cart
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/store/carts/{id}' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "email": "user@example.com"
//     }'
//
// tags:
//   - Carts
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCartsRes"
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
func (m *Cart) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.CartUpdateProps](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	customerId := context.Locals("customer_id").(uuid.UUID)

	if customerId != uuid.Nil {
		model.CustomerId = customerId
	}

	if _, err := m.r.CartService().SetContext(context.Context()).Update(id, nil, model); err != nil {
		return err
	}

	updated, err := m.r.CartService().SetContext(context.Context()).Retrieve(id, &sql.Options{Relations: []string{"payment_sessions", "shipping_methods"}}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	if len(updated.PaymentSessions) != 0 && model.RegionId != uuid.Nil {
		if err := m.r.CartService().SetContext(context.Context()).SetPaymentSessions(id, nil); err != nil {
			return err
		}
	}

	result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	var variants []models.ProductVariant
	for _, item := range result.Items {
		variants = append(variants, *item.Variant)
	}
	if _, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(variants, uuid.UUIDs{result.SalesChannelId.UUID}, &services.AvailabilityContext{}); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /store/carts/{id}/complete
// summary: "Complete a Cart"
// operationId: "PostCartsCartComplete"
// description: |
//
//	Complete a cart and place an order or create a swap, based on the cart's type. This includes attempting to authorize the cart's payment.
//	If authorizing the payment requires more action, the cart will not be completed and the order will not be placed or the swap will not be created.
//
//	An idempotency key will be generated if none is provided in the header `Idempotency-Key` and added to
//	the response. If an error occurs during cart completion or the request is interrupted for any reason, the cart completion can be retried by passing the idempotency
//	key in the `Idempotency-Key` header.
//
// externalDocs:
//
//	description: "Cart completion overview"
//	url: "https://docs.medusajs.com/modules/carts-and-checkout/cart#cart-completion"
//
// parameters:
//   - (path) id=* {String} The Cart ID.
//
// x-codegen:
//
//	method: complete
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.carts.complete(cartId)
//     .then(({ cart }) => {
//     console.log(cart.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useCompleteCart } from "medusa-react"
//
//     type Props = {
//     cartId: string
//     }
//
//     const Cart = ({ cartId }: Props) => {
//     const completeCart = useCompleteCart(cartId)
//
//     const handleComplete = () => {
//     completeCart.mutate(void 0, {
//     onSuccess: ({ data, type }) => {
//     console.log(data.id, type)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Cart
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/store/carts/{id}/complete'
//
// tags:
//   - Carts
//
// responses:
//
//	200:
//	  description: "If the payment of the cart was successfully authorized, but requires further
//	    action from the customer, the response body will contain the cart with an
//	    updated payment session. Otherwise, if the payment was authorized and the cart was successfully completed, the
//	    response body will contain either the newly created order or swap, depending on what the cart was created for."
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCompleteCartRes"
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
func (m *Cart) Complete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	headerKey := context.Get("Idempotency-Key")

	idempotencyKey := &models.IdempotencyKey{}
	idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).InitializeRequest(headerKey, context.Method(), nil, context.Path())
	if err != nil {
		return err
	}

	context.Set("Access-Control-Expose-Headers", "Idempotency-Key")
	context.Set("Idempotency-Key", idempotencyKey.IdempotencyKey)

	//TODO: add req.request_context
	result, err := m.r.CartCompletionStrategy().Complete(id, idempotencyKey, types.RequestContext{})
	if err != nil {
		return err
	}

	return context.Status(result.ResponseCode).JSON(result.ResponseBody)
}

func (m *Cart) CreateLineItem(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.AddOrderEditLineItemInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	customerId := context.Locals("customer_id").(uuid.UUID)

	headerKey := context.Get("Idempotency-Key")

	idempotencyKey := &models.IdempotencyKey{}
	idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).InitializeRequest(headerKey, context.Method(), nil, context.Path())
	if err != nil {
		return err
	}

	context.Set("Access-Control-Expose-Headers", "Idempotency-Key")
	context.Set("Idempotency-Key", idempotencyKey.IdempotencyKey)

	if idempotencyKey.RecoveryPoint == "started" {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).WorkStage(idempotencyKey.IdempotencyKey, func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
			cart, err := m.r.CartService().SetContext(context.Context()).Retrieve(id, &sql.Options{Selects: []string{"id", "region_id", "customer_id"}}, services.TotalsConfig{})
			if err != nil {
				return nil, err
			}

			cusId := customerId
			if cusId == uuid.Nil {
				cusId = cart.CustomerId.UUID
			}

			line, err := m.r.LineItemService().SetContext(context.Context()).Generate(id, nil, cart.RegionId.UUID, model.Quantity, types.GenerateLineItemContext{
				CustomerId: cusId,
				Metadata:   model.Metadata,
			})
			if err != nil {
				return nil, err
			}

			if err := m.r.CartService().SetContext(context.Context()).AddOrUpdateLineItems(cart.Id, line, true); err != nil {
				return nil, err
			}

			return &types.IdempotencyCallbackResult{RecoveryPoint: "set-payment-sessions"}, nil
		})
		if err != nil {
			return err
		}
	} else if idempotencyKey.RecoveryPoint == "set-payment-sessions" {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).WorkStage(idempotencyKey.IdempotencyKey, func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
			//TODO: add defaultStoreCartRelations
			cart, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{Relations: []string{"region.tax_rates", "customer"}}, services.TotalsConfig{})
			if err != nil {
				return nil, err
			}

			if len(cart.PaymentSessions) > 0 {
				if err := m.r.CartService().SetContext(context.Context()).SetPaymentSessions(uuid.Nil, cart); err != nil {
					return nil, err
				}
			}

			//TODO: add defaultStoreCartRelations
			result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{Relations: []string{"region.tax_rates", "customer"}}, services.TotalsConfig{})
			if err != nil {
				return nil, err
			}

			var variants []models.ProductVariant
			for _, item := range result.Items {
				variants = append(variants, *item.Variant)
			}

			if len(variants) > 0 {
				if _, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(variants, uuid.UUIDs{result.SalesChannelId.UUID}, &services.AvailabilityContext{}); err != nil {
					return nil, err
				}
			}

			return &types.IdempotencyCallbackResult{ResponseCode: 200, ResponseBody: core.JSONB{"cart": result}}, nil
		})
		if err != nil {
			return err
		}
	} else {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).Update(idempotencyKey.IdempotencyKey, &models.IdempotencyKey{
			RecoveryPoint: "finished",
			ResponseCode:  500,
			ResponseBody:  core.JSONB{"message": "Unknown recovery point"},
		})
		if err != nil {
			return err
		}
	}

	return context.Status(idempotencyKey.ResponseCode).JSON(idempotencyKey.ResponseBody)
}

// @oas:path [post] /store/carts/{id}/line-items/{line_id}
// operationId: PostCartsCartLineItemsItem
// summary: Update a Line Item
// description: "Update a line item's quantity."
// parameters:
//   - (path) id=* {string} The ID of the Cart.
//   - (path) line_id=* {string} The ID of the Line Item.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/StorePostCartsCartLineItemsItemReq"
//
// x-codegen:
//
//	method: updateLineItem
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.carts.lineItems.update(cartId, lineId, {
//     quantity: 1
//     })
//     .then(({ cart }) => {
//     console.log(cart.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useUpdateLineItem } from "medusa-react"
//
//     type Props = {
//     cartId: string
//     }
//
//     const Cart = ({ cartId }: Props) => {
//     const updateLineItem = useUpdateLineItem(cartId)
//
//     const handleUpdateItem = (
//     lineItemId: string,
//     quantity: number
//     ) => {
//     updateLineItem.mutate({
//     lineId: lineItemId,
//     quantity,
//     }, {
//     onSuccess: ({ cart }) => {
//     console.log(cart.items)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Cart
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/store/carts/{id}/line-items/{line_id}' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "quantity": 1
//     }'
//
// tags:
//   - Carts
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCartsRes"
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
func (m *Cart) UpdateLineItem(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateLineItem](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	lineId, err := api.BindDelete(context, "line_id")
	if err != nil {
		return err
	}

	if model.Quantity == 0 {
		if err = m.r.CartService().SetContext(context.Context()).RemoveLineItem(id, uuid.UUIDs{lineId}); err != nil {
			return err
		}
	} else {
		cart, err := m.r.CartService().SetContext(context.Context()).Retrieve(id, &sql.Options{Relations: []string{"items", "items.variant", "shipping_methods"}}, services.TotalsConfig{})
		if err != nil {
			return err
		}

		existing, ok := lo.Find(cart.Items, func(item models.LineItem) bool {
			return item.Id == lineId
		})

		if !ok {
			return utils.NewApplictaionError(
				utils.INVALID_DATA,
				"Could not find the line item",
			)
		}

		data := &types.LineItemUpdate{
			VariantId:             existing.Variant.Id,
			RegionId:              cart.RegionId.UUID,
			Quantity:              model.Quantity,
			Metadata:              model.Metadata,
			ShouldCalculatePrices: true,
		}

		if _, err = m.r.CartService().SetContext(context.Context()).UpdateLineItem(id, lineId, data); err != nil {
			return err
		}
	}

	updated, err := m.r.CartService().SetContext(context.Context()).Retrieve(id, &sql.Options{Relations: []string{"payment_sessions"}}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	if len(updated.PaymentSessions) > 0 {
		if err := m.r.CartService().SetContext(context.Context()).SetPaymentSessions(id, nil); err != nil {
			return err
		}
	}

	result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	var variants []models.ProductVariant
	for _, item := range result.Items {
		variants = append(variants, *item.Variant)
	}
	if _, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(variants, uuid.UUIDs{result.SalesChannelId.UUID}, &services.AvailabilityContext{}); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [delete] /store/carts/{id}/line-items/{line_id}
// operationId: DeleteCartsCartLineItemsItem
// summary: Delete a Line Item
// description: "Delete a Line Item from a Cart. The payment sessions will be updated and the totals will be recalculated."
// parameters:
//   - (path) id=* {string} The ID of the Cart.
//   - (path) line_id=* {string} The ID of the Line Item.
//
// x-codegen:
//
//	method: deleteLineItem
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.carts.lineItems.delete(cartId, lineId)
//     .then(({ cart }) => {
//     console.log(cart.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useDeleteLineItem } from "medusa-react"
//
//     type Props = {
//     cartId: string
//     }
//
//     const Cart = ({ cartId }: Props) => {
//     const deleteLineItem = useDeleteLineItem(cartId)
//
//     const handleDeleteItem = (
//     lineItemId: string
//     ) => {
//     deleteLineItem.mutate({
//     lineId: lineItemId,
//     }, {
//     onSuccess: ({ cart }) => {
//     console.log(cart.items)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Cart
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '{backend_url}/store/carts/{id}/line-items/{line_id}'
//
// tags:
//   - Carts
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCartsRes"
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
func (m *Cart) DeleteLineItem(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	lineId, err := api.BindDelete(context, "line_id")
	if err != nil {
		return err
	}

	if err := m.r.CartService().SetContext(context.Context()).RemoveLineItem(id, uuid.UUIDs{lineId}); err != nil {
		return err
	}

	updated, err := m.r.CartService().SetContext(context.Context()).Retrieve(id, &sql.Options{Relations: []string{"payment_sessions"}}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	if len(updated.PaymentSessions) > 0 {
		if err := m.r.CartService().SetContext(context.Context()).SetPaymentSessions(id, nil); err != nil {
			return err
		}
	}

	result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	var variants []models.ProductVariant
	for _, item := range result.Items {
		variants = append(variants, *item.Variant)
	}
	if _, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(variants, uuid.UUIDs{result.SalesChannelId.UUID}, &services.AvailabilityContext{}); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /store/carts/{id}/payment-sessions
// operationId: "PostCartsCartPaymentSessions"
// summary: "Create Payment Sessions"
// description: "Create Payment Sessions for each of the available Payment Providers in the Cart's Region. If there's only one payment session created,
//
//	it will be selected by default. The creation of the payment session uses the payment provider and may require sending requests to third-party services."
//
// parameters:
//   - (path) id=* {string} The ID of the Cart.
//
// x-codegen:
//
//	method: createPaymentSessions
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.carts.createPaymentSessions(cartId)
//     .then(({ cart }) => {
//     console.log(cart.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useCreatePaymentSession } from "medusa-react"
//
//     type Props = {
//     cartId: string
//     }
//
//     const Cart = ({ cartId }: Props) => {
//     const createPaymentSession = useCreatePaymentSession(cartId)
//
//     const handleComplete = () => {
//     createPaymentSession.mutate(void 0, {
//     onSuccess: ({ cart }) => {
//     console.log(cart.payment_sessions)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Cart
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/store/carts/{id}/payment-sessions'
//
// tags:
//   - Carts
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCartsRes"
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
func (m *Cart) CreatePaymentSession(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	headerKey := context.Get("Idempotency-Key")

	idempotencyKey := &models.IdempotencyKey{}
	idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).InitializeRequest(headerKey, context.Method(), nil, context.Path())
	if err != nil {
		return err
	}

	context.Set("Access-Control-Expose-Headers", "Idempotency-Key")
	context.Set("Idempotency-Key", idempotencyKey.IdempotencyKey)

	if idempotencyKey.RecoveryPoint == "started" {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).WorkStage(idempotencyKey.IdempotencyKey, func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
			//TODO: add defaultStoreCartRelations
			cart, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{Relations: []string{"region.tax_rates", "customer"}}, services.TotalsConfig{})
			if err != nil {
				return nil, err
			}

			if len(cart.PaymentSessions) > 0 {
				if err := m.r.CartService().SetContext(context.Context()).SetPaymentSessions(uuid.Nil, cart); err != nil {
					return nil, err
				}
			}

			//TODO: add defaultStoreCartRelations
			result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{Relations: []string{"region.tax_rates", "customer"}}, services.TotalsConfig{})
			if err != nil {
				return nil, err
			}

			var variants []models.ProductVariant
			for _, item := range result.Items {
				variants = append(variants, *item.Variant)
			}

			if len(variants) > 0 {
				if _, err := m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(variants, uuid.UUIDs{result.SalesChannelId.UUID}, &services.AvailabilityContext{}); err != nil {
					return nil, err
				}
			}

			return &types.IdempotencyCallbackResult{ResponseCode: 200, ResponseBody: core.JSONB{"cart": result}}, nil
		})
		if err != nil {
			return err
		}
	} else {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).Update(idempotencyKey.IdempotencyKey, &models.IdempotencyKey{
			RecoveryPoint: "finished",
			ResponseCode:  500,
			ResponseBody:  core.JSONB{"message": "Unknown recovery point"},
		})
		if err != nil {
			return err
		}
	}

	return context.Status(idempotencyKey.ResponseCode).JSON(idempotencyKey.ResponseBody)
}

// @oas:path [post] /store/carts/{id}/payment-sessions/{provider_id}
// operationId: PostCartsCartPaymentSessionUpdate
// summary: Update a Payment Session
// description: "Update a Payment Session with additional data. This can be useful depending on the payment provider used.
//
//	All payment sessions are updated and cart totals are recalculated afterwards."
//
// parameters:
//   - (path) id=* {string} The ID of the Cart.
//   - (path) provider_id=* {string} The ID of the payment provider.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/StorePostCartsCartPaymentSessionUpdateReq"
//
// x-codegen:
//
//	method: updatePaymentSession
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.carts.updatePaymentSession(cartId, "manual", {
//     data: {
//
//     }
//     })
//     .then(({ cart }) => {
//     console.log(cart.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useUpdatePaymentSession } from "medusa-react"
//
//     type Props = {
//     cartId: string
//     }
//
//     const Cart = ({ cartId }: Props) => {
//     const updatePaymentSession = useUpdatePaymentSession(cartId)
//
//     const handleUpdate = (
//     providerId: string,
//     data: Record<string, unknown>
//     ) => {
//     updatePaymentSession.mutate({
//     provider_id: providerId,
//     data
//     }, {
//     onSuccess: ({ cart }) => {
//     console.log(cart.payment_session)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Cart
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/store/carts/{id}/payment-sessions/manual' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "data": {}
//     }'
//
// tags:
//   - Carts
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCartsRes"
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
func (m *Cart) UpdatePaymentSession(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdatePaymentSession](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	providerId, err := api.BindDelete(context, "provider_id")
	if err != nil {
		return err
	}

	if err := m.r.CartService().SetContext(context.Context()).SetPaymentSession(id, providerId); err != nil {
		return err
	}

	if _, err := m.r.CartService().SetContext(context.Context()).UpdatePaymentSession(id, model.Data); err != nil {
		return err
	}

	result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	var variants []models.ProductVariant
	for _, item := range result.Items {
		variants = append(variants, *item.Variant)
	}
	if _, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(variants, uuid.UUIDs{result.SalesChannelId.UUID}, &services.AvailabilityContext{}); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [delete] /store/carts/{id}/payment-sessions/{provider_id}
// operationId: DeleteCartsCartPaymentSessionsSession
// summary: "Delete a Payment Session"
// description: "Delete a Payment Session in a Cart. May be useful if a payment has failed. The totals will be recalculated."
// parameters:
//   - (path) id=* {string} The ID of the Cart.
//   - (path) provider_id=* {string} The ID of the Payment Provider used to create the Payment Session to be deleted.
//
// x-codegen:
//
//	method: deletePaymentSession
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.carts.deletePaymentSession(cartId, "manual")
//     .then(({ cart }) => {
//     console.log(cart.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useDeletePaymentSession } from "medusa-react"
//
//     type Props = {
//     cartId: string
//     }
//
//     const Cart = ({ cartId }: Props) => {
//     const deletePaymentSession = useDeletePaymentSession(cartId)
//
//     const handleDeletePaymentSession = (
//     providerId: string
//     ) => {
//     deletePaymentSession.mutate({
//     provider_id: providerId,
//     }, {
//     onSuccess: ({ cart }) => {
//     console.log(cart.payment_sessions)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Cart
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '{backend_url}/store/carts/{id}/payment-sessions/{provider_id}'
//
// tags:
//   - Carts
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCartsRes"
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
func (m *Cart) DeletePaymentSession(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	providerId, err := api.BindDelete(context, "provider_id")
	if err != nil {
		return err
	}

	if err := m.r.CartService().SetContext(context.Context()).DeletePaymentSession(id, providerId); err != nil {
		return err
	}

	result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	var variants []models.ProductVariant
	for _, item := range result.Items {
		variants = append(variants, *item.Variant)
	}
	if _, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(variants, uuid.UUIDs{result.SalesChannelId.UUID}, &services.AvailabilityContext{}); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /store/carts/{id}/payment-sessions/{provider_id}/refresh
// operationId: PostCartsCartPaymentSessionsSession
// summary: Refresh a Payment Session
// description: "Refresh a Payment Session to ensure that it is in sync with the Cart. This is usually not necessary, but is provided for edge cases."
// parameters:
//   - (path) id=* {string} The ID of the Cart.
//   - (path) provider_id=* {string} The ID of the Payment Provider that created the Payment Session to be refreshed.
//
// x-codegen:
//
//	method: refreshPaymentSession
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.carts.refreshPaymentSession(cartId, "manual")
//     .then(({ cart }) => {
//     console.log(cart.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useRefreshPaymentSession } from "medusa-react"
//
//     type Props = {
//     cartId: string
//     }
//
//     const Cart = ({ cartId }: Props) => {
//     const refreshPaymentSession = useRefreshPaymentSession(cartId)
//
//     const handleRefresh = (
//     providerId: string
//     ) => {
//     refreshPaymentSession.mutate({
//     provider_id: providerId,
//     }, {
//     onSuccess: ({ cart }) => {
//     console.log(cart.payment_sessions)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Cart
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/store/carts/{id}/payment-sessions/{provider_id}/refresh'
//
// tags:
//   - Carts
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCartsRes"
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
func (m *Cart) RefreshPaymentSession(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	providerId, err := api.BindDelete(context, "provider_id")
	if err != nil {
		return err
	}

	if err := m.r.CartService().SetContext(context.Context()).RefreshPaymentSession(id, providerId); err != nil {
		return err
	}

	result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{
		Relations: []string{
			"region",
			"region.countries",
			"region.payment_providers",
			"shipping_methods",
			"payment_sessions",
			"shipping_methods.shipping_option",
		},
	}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /store/carts/{id}/payment-session
// operationId: PostCartsCartPaymentSession
// summary: Select a Payment Session
// description: "Select the Payment Session that will be used to complete the cart. This is typically used when the customer chooses their preferred payment method during checkout.
//
//	The totals of the cart will be recalculated."
//
// parameters:
//   - (path) id=* {string} The ID of the Cart.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/StorePostCartsCartPaymentSessionReq"
//
// x-codegen:
//
//	method: setPaymentSession
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.carts.setPaymentSession(cartId, {
//     provider_id: "manual"
//     })
//     .then(({ cart }) => {
//     console.log(cart.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useSetPaymentSession } from "medusa-react"
//
//     type Props = {
//     cartId: string
//     }
//
//     const Cart = ({ cartId }: Props) => {
//     const setPaymentSession = useSetPaymentSession(cartId)
//
//     const handleSetPaymentSession = (
//     providerId: string
//     ) => {
//     setPaymentSession.mutate({
//     provider_id: providerId,
//     }, {
//     onSuccess: ({ cart }) => {
//     console.log(cart.payment_session)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Cart
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/store/carts/{id}/payment-sessions' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "provider_id": "manual"
//     }'
//
// tags:
//   - Carts
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCartsRes"
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
func (m *Cart) SetPaymentSession(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.SessionsInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if err := m.r.CartService().SetContext(context.Context()).SetPaymentSession(id, model.ProviderId); err != nil {
		return err
	}

	result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	var variants []models.ProductVariant
	for _, item := range result.Items {
		variants = append(variants, *item.Variant)
	}
	if _, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(variants, uuid.UUIDs{result.SalesChannelId.UUID}, &services.AvailabilityContext{}); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /store/carts/{id}/shipping-methods
// operationId: "PostCartsCartShippingMethod"
// summary: "Add Shipping Method"
// description: "Add a Shipping Method to the Cart. The validation of the `data` field is handled by the fulfillment provider of the chosen shipping option."
// parameters:
//   - (path) id=* {string} The cart ID.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/StorePostCartsCartShippingMethodReq"
//
// x-codegen:
//
//	method: addShippingMethod
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.carts.addShippingMethod(cartId, {
//     option_id
//     })
//     .then(({ cart }) => {
//     console.log(cart.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAddShippingMethodToCart } from "medusa-react"
//
//     type Props = {
//     cartId: string
//     }
//
//     const Cart = ({ cartId }: Props) => {
//     const addShippingMethod = useAddShippingMethodToCart(cartId)
//
//     const handleAddShippingMethod = (
//     optionId: string
//     ) => {
//     addShippingMethod.mutate({
//     option_id: optionId,
//     }, {
//     onSuccess: ({ cart }) => {
//     console.log(cart.shipping_methods)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Cart
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/store/carts/{id}/shipping-methods' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "option_id": "{option_id}",
//     }'
//
// tags:
//   - Carts
//
// responses:
//
//	"200":
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCartsRes"
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
func (m *Cart) AddShippingMethod(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.AddShippingMethod](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if _, err := m.r.CartService().SetContext(context.Context()).AddShippingMethod(id, nil, model.OptionId, model.Data); err != nil {
		return err
	}

	updated, err := m.r.CartService().SetContext(context.Context()).Retrieve(id, &sql.Options{Relations: []string{"payment_sessions"}}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	if len(updated.PaymentSessions) > 0 {
		if err := m.r.CartService().SetContext(context.Context()).SetPaymentSessions(id, nil); err != nil {
			return err
		}
	}

	result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	var variants []models.ProductVariant
	for _, item := range result.Items {
		variants = append(variants, *item.Variant)
	}
	if _, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(variants, uuid.UUIDs{result.SalesChannelId.UUID}, &services.AvailabilityContext{}); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /store/carts/{id}/taxes
// operationId: "PostCartsCartTaxes"
// summary: "Calculate Cart Taxes"
// description: "Calculate the taxes for a cart. This is useful if the `automatic_taxes` field of the cart's region is set to `false`. If the cart's region uses a tax provider other than
//
//	Medusa's system provider, this may lead to sending requests to third-party services."
//
// externalDocs:
//
//	description: "How to calculate taxes manually during checkout"
//	url: "https://docs.medusajs.com/modules/taxes/storefront/manual-calculation"
//
// parameters:
//   - (path) id=* {String} The Cart ID.
//
// x-codegen:
//
//	method: calculateTaxes
//
// x-codeSamples:
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/store/carts/{id}/taxes'
//
// tags:
//   - Carts
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCartsRes"
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
func (m *Cart) CalculateTaxes(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	headerKey := context.Get("Idempotency-Key")

	idempotencyKey := &models.IdempotencyKey{}
	idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).InitializeRequest(headerKey, context.Method(), nil, context.Path())
	if err != nil {
		return err
	}

	context.Set("Access-Control-Expose-Headers", "Idempotency-Key")
	context.Set("Idempotency-Key", idempotencyKey.IdempotencyKey)

	if idempotencyKey.RecoveryPoint == "started" {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).WorkStage(idempotencyKey.IdempotencyKey, func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
			result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{}, services.TotalsConfig{ForceTaxes: true})
			if err != nil {
				return nil, err
			}

			return &types.IdempotencyCallbackResult{ResponseCode: 200, ResponseBody: core.JSONB{"cart": result}}, nil
		})
		if err != nil {
			return err
		}
	} else {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).Update(idempotencyKey.IdempotencyKey, &models.IdempotencyKey{
			RecoveryPoint: "finished",
			ResponseCode:  500,
			ResponseBody:  core.JSONB{"message": "Unknown recovery point"},
		})
		if err != nil {
			return err
		}
	}

	return context.Status(idempotencyKey.ResponseCode).JSON(idempotencyKey.ResponseBody)
}

// @oas:path [delete] /store/carts/{id}/discounts/{code}
// operationId: DeleteCartsCartDiscountsDiscount
// summary: "Remove Discount"
// description: "Remove a Discount from a Cart. This only removes the application of the discount, and not completely deletes it. The totals will be re-calculated and the payment sessions
//
//	will be refreshed after the removal."
//
// parameters:
//   - (path) id=* {string} The ID of the Cart.
//   - (path) code=* {string} The unique discount code.
//
// x-codegen:
//
//	method: deleteDiscount
//
// x-codeSamples:
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.carts.deleteDiscount(cartId, code)
//     .then(({ cart }) => {
//     console.log(cart.id);
//     })
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '{backend_url}/store/carts/{id}/discounts/{code}'
//
// tags:
//   - Carts
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCartsRes"
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
func (m *Cart) DeleteDiscount(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	code := context.Params("code")

	if _, err := m.r.CartService().SetContext(context.Context()).RemoveDiscount(id, code); err != nil {
		return err
	}

	updated, err := m.r.CartService().SetContext(context.Context()).Retrieve(id, &sql.Options{Relations: []string{"payment_sessions"}}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	if len(updated.PaymentSessions) > 0 {
		if err := m.r.CartService().SetContext(context.Context()).SetPaymentSessions(id, nil); err != nil {
			return err
		}
	}

	result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	var variants []models.ProductVariant
	for _, item := range result.Items {
		variants = append(variants, *item.Variant)
	}
	if _, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(variants, uuid.UUIDs{result.SalesChannelId.UUID}, &services.AvailabilityContext{}); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)

}

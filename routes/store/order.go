package store

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type Order struct {
	r Registry
}

func NewOrder(r Registry) *Order {
	m := Order{r: r}
	return &m
}

func (m *Order) SetRoutes(router fiber.Router) {
	route := router.Group("/orders")
	route.Get("", m.Lookup)
	route.Get("/:id", m.Get)

	route.Get("/cart/:cart_id", m.GetByCart)
	route.Post("/customer/confirm", m.ConfirmRequest)
	route.Post("/batch/customer/token", m.Request)
}

// @oas:path [get] /store/orders/{id}
// operationId: GetOrdersOrder
// summary: Get an Order
// description: "Retrieve an Order's details."
// parameters:
//   - (path) id=* {string} The ID of the Order.
//   - (query) fields {string} Comma-separated fields that should be expanded in the returned order.
//   - (query) expand {string} Comma-separated relations that should be included in the returned order.
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
//     medusa.orders.retrieve(orderId)
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useOrder } from "medusa-react"
//
//     type Props = {
//     orderId: string
//     }
//
//     const Order = ({ orderId }: Props) => {
//     const {
//     order,
//     isLoading,
//     } = useOrder(orderId)
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {order && <span>{order.display_id}</span>}
//
//     </div>
//     )
//     }
//
//     export default Order
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '{backend_url}/store/orders/{id}'
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreOrdersRes"
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
func (m *Order) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [get] /store/orders/cart/{cart_id}
// operationId: GetOrdersOrderCartId
// summary: Get by Cart ID
// description: "Retrieve an Order's details by the ID of the Cart that was used to create the Order."
// parameters:
//   - (path) cart_id=* {string} The ID of Cart.
//
// x-codegen:
//
//	method: retrieveByCartId
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.orders.retrieveByCartId(cartId)
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useCartOrder } from "medusa-react"
//
//     type Props = {
//     cartId: string
//     }
//
//     const Order = ({ cartId }: Props) => {
//     const {
//     order,
//     isLoading,
//     } = useCartOrder(cartId)
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {order && <span>{order.display_id}</span>}
//
//     </div>
//     )
//     }
//
//     export default Order
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '{backend_url}/store/orders/cart/{cart_id}'
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreOrdersRes"
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
func (m *Order) GetByCart(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "cart_id")
	if err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByCartIdWithTotals(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [get] /store/orders
// operationId: "GetOrders"
// summary: "Look Up an Order"
// description: "Look up an order using filters. If the filters don't narrow down the results to a single order, a 404 response is returned with no orders."
// parameters:
//   - (query) display_id=* {number} Filter by ID.
//   - (query) fields {string} Comma-separated fields that should be expanded in the returned order.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned order.
//   - in: query
//     name: email
//     style: form
//     explode: false
//     description: Filter by email.
//     required: true
//     schema:
//     type: string
//     format: email
//   - in: query
//     name: shipping_address
//     style: form
//     explode: false
//     description: Filter by the shipping address's postal code.
//     schema:
//     type: object
//     properties:
//     postal_code:
//     type: string
//     description: The postal code of the shipping address
//
// x-codegen:
//
//	method: lookupOrder
//	queryParams: StoreGetOrdersParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.orders.lookupOrder({
//     display_id: 1,
//     email: "user@example.com"
//     })
//     .then(({ order }) => {
//     console.log(order.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useOrders } from "medusa-react"
//
//     type Props = {
//     displayId: number
//     email: string
//     }
//
//     const Order = ({
//     displayId,
//     email
//     }: Props) => {
//     const {
//     order,
//     isLoading,
//     } = useOrders({
//     display_id: displayId,
//     email,
//     })
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {order && <span>{order.display_id}</span>}
//
//     </div>
//     )
//     }
//
//     export default Order
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '{backend_url}/store/orders?display_id=1&email=user@example.com'
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreOrdersRes"
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
func (m *Order) Lookup(context fiber.Ctx) error {
	model, err := api.Bind[types.OrderLookup](context, m.r.Validator())
	if err != nil {
		return err
	}

	orders, err := m.r.OrderService().SetContext(context.Context()).List(&types.FilterableOrder{
		DisplayId: model.DisplayId,
		Email:     model.Email,
	}, &sql.Options{})
	if err != nil {
		return err
	}

	if len(orders) != 1 {
		return context.SendStatus(fiber.StatusNotFound)
	}

	return context.Status(fiber.StatusOK).JSON(orders[0])
}

// @oas:path [post] /store/orders/customer/confirm
// operationId: "PostOrdersCustomerOrderClaimsCustomerOrderClaimAccept"
// summary: "Verify Order Claim"
// description: "Verify the claim order token provided to the customer when they request ownership of an order."
// externalDocs:
//
//	description: "How to implement claim-order flow in a storefront"
//	url: "https://docs.medusajs.com/modules/orders/storefront/implement-claim-order"
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/StorePostCustomersCustomerAcceptClaimReq"
//
// x-codegen:
//
//	method: confirmRequest
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.orders.confirmRequest(
//     token,
//     )
//     .then(() => {
//     // successful
//     })
//     .catch(() => {
//     // an error occurred
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useGrantOrderAccess } from "medusa-react"
//
//     const ClaimOrder = () => {
//     const confirmOrderRequest = useGrantOrderAccess()
//
//     const handleOrderRequestConfirmation = (
//     token: string
//     ) => {
//     confirmOrderRequest.mutate({
//     token
//     }, {
//     onSuccess: () => {
//     // successful
//     },
//     onError: () => {
//     // an error occurred.
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default ClaimOrder
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/store/orders/customer/confirm' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "token": "{token}",
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: OK
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
func (m *Order) ConfirmRequest(context fiber.Ctx) error {
	model, err := api.Bind[types.CustomerAcceptClaim](context, m.r.Validator())
	if err != nil {
		return err
	}

	_, claims, errObj := m.r.TockenService().SetContext(context.Context()).VerifyToken(model.Token)
	if errObj != nil {
		return utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Invalid token",
		)
	}

	customerId := claims["claimingCustomerId"].(uuid.UUID)

	customer, err := m.r.CustomerService().SetContext(context.Context()).RetrieveById(customerId, &sql.Options{})
	if err != nil {
		return err
	}

	orders, err := m.r.OrderService().SetContext(context.Context()).List(&types.FilterableOrder{FilterModel: core.FilterModel{Id: claims["orders"].(uuid.UUIDs)}}, &sql.Options{})
	if err != nil {
		return err
	}

	for _, order := range orders {
		if _, err := m.r.OrderService().SetContext(context.Context()).Update(order.Id, &types.UpdateOrderInput{
			CustomerId: customerId,
			Email:      customer.Email,
		}); err != nil {
			return err
		}
	}

	return context.SendStatus(fiber.StatusOK)
}

// @oas:path [post] /store/orders/batch/customer/token
// operationId: "PostOrdersCustomerOrderClaim"
// summary: "Claim Order"
// description: "Allow the logged-in customer to claim ownership of one or more orders. This generates a token that can be used later on to verify the claim using the Verify Order Claim API Route.
//
//	This also emits the event `order-update-token.created`. So, if you have a notification provider installed that handles this event and sends the customer a notification, such as an email,
//	the customer should receive instructions on how to finalize their claim ownership."
//
// externalDocs:
//
//	description: "How to implement claim-order flow in a storefront"
//	url: "https://docs.medusajs.com/modules/orders/storefront/implement-claim-order"
//
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/StorePostCustomersCustomerOrderClaimReq"
//
// x-codegen:
//
//	method: requestCustomerOrders
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.orders.requestCustomerOrders({
//     order_ids,
//     })
//     .then(() => {
//     // successful
//     })
//     .catch(() => {
//     // an error occurred
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useRequestOrderAccess } from "medusa-react"
//
//     const ClaimOrder = () => {
//     const claimOrder = useRequestOrderAccess()
//
//     const handleClaimOrder = (
//     orderIds: string[]
//     ) => {
//     claimOrder.mutate({
//     order_ids: orderIds
//     }, {
//     onSuccess: () => {
//     // successful
//     },
//     onError: () => {
//     // an error occurred.
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default ClaimOrder
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/store/batch/customer/token' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "order_ids": ["id"],
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Orders
//
// responses:
//
//	200:
//	  description: OK
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
func (m *Order) Request(context fiber.Ctx) error {
	model, err := api.Bind[types.CustomerOrderClaim](context, m.r.Validator())
	if err != nil {
		return err
	}

	customerId := context.Locals("customer_id").(uuid.UUID)

	customer, err := m.r.CustomerService().SetContext(context.Context()).RetrieveById(customerId, &sql.Options{})
	if err != nil {
		return err
	}

	if !customer.HasAccount {
		utils.NewApplictaionError(
			utils.UNAUTHORIZED,
			"Customer does not have an account",
		)
	}

	orders, err := m.r.OrderService().SetContext(context.Context()).List(&types.FilterableOrder{FilterModel: core.FilterModel{Id: model.OrderIds}}, &sql.Options{Selects: []string{"id", "email"}})
	if err != nil {
		return err
	}

	emailOrderMapping := make(map[string]uuid.UUIDs)
	for _, order := range orders {
		emailOrderMapping[order.Email] = append(emailOrderMapping[order.Email], order.Id)
	}

	// 1. email
	for _, ids := range emailOrderMapping {
		// 1. token
		_, errObj := m.r.TockenService().SetContext(context.Context()).SignToken(map[string]interface{}{
			"claimingCustomerId": customerId,
			"orders":             ids,
		})
		if errObj != nil {
			return utils.NewApplictaionError(
				utils.NOT_ALLOWED,
				"Invalid token",
			)
		}

		// err = eventBusService.Emit(TokenEventsOrderUpdateTokenCreated, TokenEventPayload{
		// 	OldEmail:      email,
		// 	NewCustomerID: customer.ID,
		// 	Orders:        ids,
		// 	Token:         token,
		// })
		// if err != nil {
		// 	return err
		// }
	}

	return context.SendStatus(fiber.StatusOK)
}

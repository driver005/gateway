package store

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/fatih/structs"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type Customer struct {
	r Registry
}

func NewCustomer(r Registry) *Customer {
	m := Customer{r: r}
	return &m
}

func (m *Customer) SetRoutes(router fiber.Router) {
	route := router.Group("/customers")
	route.Post("", m.Create)

	route.Post("/password-tocken", m.ResetPasswordTocken)
	route.Post("/reste-password", m.ResetPassword)

	route.Use(utils.ConvertMiddleware(m.r.Middleware().AuthenticateCustomer())...)

	route.Get("/me", m.Get)
	route.Post("/me", m.Update)

	route.Post("/me/orders", m.ListOrders)
	route.Post("/me/addresses", m.CreateAddress)
	route.Post("/me/addresses/:address_id", m.UpdateAddress)
	route.Delete("/me/addresses/:address_id", m.DeleteAdress)
	route.Get("/me/payment-methods", m.GetPaymnetMethods)

}

// @oas:path [get] /store/customers/me
// operationId: GetCustomersCustomer
// summary: Get a Customer
// description: "Retrieve the logged-in Customer's details."
// x-authenticated: true
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
//     // must be previously logged
//     medusa.customers.retrieve()
//     .then(({ customer }) => {
//     console.log(customer.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useMeCustomer } from "medusa-react"
//
//     const Customer = () => {
//     const { customer, isLoading } = useMeCustomer()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {customer && (
//     <span>{customer.first_name} {customer.last_name}</span>
//     )}
//     </div>
//     )
//     }
//
//     export default Customer
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '{backend_url}/store/customers/me' \
//     -H 'Authorization: Bearer {access_token}'
//
// security:
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Customers
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCustomersRes"
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
func (m *Customer) Get(context fiber.Ctx) error {
	id := context.Locals("customer_id").(uuid.UUID)

	result, err := m.r.CustomerService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /store/customers
// operationId: PostCustomers
// summary: Create a Customer
// description: "Register a new customer. This will also automatically authenticate the customer and set their login session in the response Cookie header.
//
//	The cookie session can be used in subsequent requests to authenticate the customer.
//	When using Medusa's JS or Medusa React clients, the cookie is automatically attached to subsequent requests."
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/PostCustomersReq"
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
//     medusa.customers.create({
//     first_name: "Alec",
//     last_name: "Reynolds",
//     email: "user@example.com",
//     password: "supersecret"
//     })
//     .then(({ customer }) => {
//     console.log(customer.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useCreateCustomer } from "medusa-react"
//
//     const RegisterCustomer = () => {
//     const createCustomer = useCreateCustomer()
//     // ...
//
//     const handleCreate = (
//     customerData: {
//     first_name: string
//     last_name: string
//     email: string
//     password: string
//     }
//     ) => {
//     // ...
//     createCustomer.mutate(customerData, {
//     onSuccess: ({ customer }) => {
//     console.log(customer.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default RegisterCustomer
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/store/customers' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "first_name": "Alec",
//     "last_name": "Reynolds",
//     "email": "user@example.com",
//     "password": "supersecret"
//     }'
//
// tags:
//   - Customers
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCustomersRes"
//	422:
//	  description: A customer with the same email exists
//	  content:
//	    application/json:
//	      schema:
//	        type: object
//	        properties:
//	          code:
//	            type: string
//	            description: The error code
//	          type:
//	            type: string
//	            description: The type of error
//	          message:
//	            type: string
//	            description: Human-readable message with details about the error
//	      example:
//	        code: "invalid_request_error"
//	        type: "duplicate_error"
//	        message: "A customer with the given email already has an account. Log in instead"
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
//	"500":
//	  description: "Internal Server"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/responses/500_error"
func (m *Customer) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateCustomerInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.CustomerService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /store/customers/me
// operationId: PostCustomersCustomer
// summary: Update Customer
// description: "Update the logged-in customer's details."
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/PostCustomersCustomerReq"
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
//     // must be previously logged
//     medusa.customers.update({
//     first_name: "Laury"
//     })
//     .then(({ customer }) => {
//     console.log(customer.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useUpdateMe } from "medusa-react"
//
//     type Props = {
//     customerId: string
//     }
//
//     const Customer = ({ customerId }: Props) => {
//     const updateCustomer = useUpdateMe()
//     // ...
//
//     const handleUpdate = (
//     firstName: string
//     ) => {
//     // ...
//     updateCustomer.mutate({
//     id: customerId,
//     first_name: firstName,
//     }, {
//     onSuccess: ({ customer }) => {
//     console.log(customer.first_name)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Customer
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/store/customers/me' \
//     -H 'Authorization: Bearer {access_token}' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "first_name": "Laury"
//     }'
//
// security:
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Customers
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCustomersRes"
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
func (m *Customer) Update(context fiber.Ctx) error {
	id := context.Locals("customer_id").(uuid.UUID)

	model, err := api.Bind[types.UpdateCustomerInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.CustomerService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [get] /store/customers/me/payment-methods
// operationId: GetCustomersCustomerPaymentMethods
// summary: Get Saved Payment Methods
// description: "Retrieve the logged-in customer's saved payment methods. This API Route only works with payment providers created with the deprecated Payment Service interface.
//
//	The payment methods are saved using the Payment Service's third-party service, and not on the Medusa backend. So, they're retrieved from the third-party service."
//
// x-authenticated: true
// deprecated: true
// x-codegen:
//
//	method: listPaymentMethods
//
// x-codeSamples:
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged
//     medusa.customers.paymentMethods.list()
//     .then(({ payment_methods }) => {
//     console.log(payment_methods.length);
//     })
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '{backend_url}/store/customers/me/payment-methods' \
//     -H 'Authorization: Bearer {access_token}'
//
// security:
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Customers
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCustomersListPaymentMethodsRes"
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
func (m *Customer) GetPaymnetMethods(context fiber.Ctx) error {
	id := context.Locals("customer_id").(uuid.UUID)

	customer, err := m.r.CustomerService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	paymentProviders, err := m.r.PaymentProviderService().SetContext(context.Context()).List()
	if err != nil {
		return err
	}

	var methods []types.PaymentMethod
	for _, paymentProvider := range paymentProviders {
		provider, err := m.r.PaymentProviderService().SetContext(context.Context()).RetrieveProvider(paymentProvider.Id)
		if err != nil {
			return err
		}

		pMethods := provider.RetrieveSavedMethods(customer)
		for _, pMethod := range pMethods {
			methods = append(methods, types.PaymentMethod{
				ProviderId: paymentProvider.Id,
				Data:       structs.Map(pMethod),
			})
		}
	}

	return context.Status(fiber.StatusOK).JSON(methods)
}

// @oas:path [get] /store/customers/me/orders
// operationId: GetCustomersCustomerOrders
// summary: List Orders
// description: "Retrieve a list of the logged-in Customer's Orders. The orders can be filtered by fields such as `status` or `fulfillment_status`. The orders can also be paginated."
// x-authenticated: true
// parameters:
//   - (query) q {string} term to search orders' display ID, email, shipping address's first name, customer's first name, customer's last name, and customer's phone number.
//   - (query) id {string} Filter by ID.
//   - in: query
//     name: status
//     style: form
//     explode: false
//     description: Filter by status.
//     schema:
//     type: array
//     items:
//     type: string
//     enum: [pending, completed, archived, canceled, requires_action]
//   - in: query
//     name: fulfillment_status
//     style: form
//     explode: false
//     description: Fulfillment status to search for.
//     schema:
//     type: array
//     items:
//     type: string
//     enum: [not_fulfilled, partially_fulfilled, fulfilled, partially_shipped, shipped, partially_returned, returned, canceled, requires_action]
//   - in: query
//     name: payment_status
//     style: form
//     explode: false
//     description: Payment status to search for.
//     schema:
//     type: array
//     items:
//     type: string
//     enum: [not_paid, awaiting, captured, partially_refunded, refunded, canceled, requires_action]
//   - (query) display_id {string} Filter by display ID.
//   - (query) cart_id {string} Filter by cart ID.
//   - (query) email {string} Filter by email.
//   - (query) region_id {string} Filter by region ID.
//   - in: query
//     name: currency_code
//     style: form
//     explode: false
//     description: Filter by the 3 character ISO currency code of the order.
//     schema:
//     type: string
//     externalDocs:
//     url: https://en.wikipedia.org/wiki/ISO_4217#Active_codes
//     description: See a list of codes.
//   - (query) tax_rate {string} Filter by tax rate.
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
//     name: canceled_at
//     description: Filter by a cancelation date range.
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
//   - (query) limit=10 {integer} Limit the number of orders returned.
//   - (query) offset=0 {integer} The number of orders to skip when retrieving the orders.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned orders.
//   - (query) fields {string} Comma-separated fields that should be included in the returned orders.
//
// x-codegen:
//
//	method: listOrders
//	queryParams: StoreGetCustomersCustomerOrdersParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged
//     medusa.customers.listOrders()
//     .then(({ orders, limit, offset, count }) => {
//     console.log(orders);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useCustomerOrders } from "medusa-react"
//
//     const Orders = () => {
//     // refetch a function that can be used to
//     // re-retrieve orders after the customer logs in
//     const { orders, isLoading } = useCustomerOrders()
//
//     return (
//     <div>
//     {isLoading && <span>Loading orders...</span>}
//     {orders?.length && (
//     <ul>
//     {orders.map((order) => (
//     <li key={order.id}>{order.display_id}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default Orders
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '{backend_url}/store/customers/me/orders' \
//     -H 'Authorization: Bearer {access_token}'
//
// security:
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Customers
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCustomersListOrdersRes"
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
func (m *Customer) ListOrders(context fiber.Ctx) error {
	id := context.Locals("customer_id").(uuid.UUID)

	model, config, err := api.BindList[types.FilterableOrder](context)
	if err != nil {
		return err
	}

	model.CustomerId = id

	result, count, err := m.r.OrderService().SetContext(context.Context()).ListAndCount(model, config)
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

// @oas:path [post] /store/customers/me/addresses
// operationId: PostCustomersCustomerAddresses
// summary: "Add a Shipping Address"
// description: "Add a Shipping Address to a Customer's saved addresses."
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/StorePostCustomersCustomerAddressesReq"
//
// x-codegen:
//
//	method: addAddress
//
// x-codeSamples:
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged
//     medusa.customers.addresses.addAddress({
//     address: {
//     first_name: "Celia",
//     last_name: "Schumm",
//     address_1: "225 Bednar Curve",
//     city: "Danielville",
//     country_code: "US",
//     postal_code: "85137",
//     phone: "981-596-6748 x90188",
//     company: "Wyman LLC",
//     province: "Georgia",
//     }
//     })
//     .then(({ customer }) => {
//     console.log(customer.id);
//     })
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/store/customers/me/addresses' \
//     -H 'Authorization: Bearer {access_token}' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "address": {
//     "first_name": "Celia",
//     "last_name": "Schumm",
//     "address_1": "225 Bednar Curve",
//     "city": "Danielville",
//     "country_code": "US",
//     "postal_code": "85137"
//     }
//     }'
//
// security:
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Customers
//
// responses:
//
//	"200":
//	  description: "A successful response"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCustomersRes"
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
func (m *Customer) CreateAddress(context fiber.Ctx) error {
	id := context.Locals("customer_id").(uuid.UUID)

	model, err := api.BindCreate[types.CustomerAddAddress](context, m.r.Validator())
	if err != nil {
		return err
	}

	if _, _, err := m.r.CustomerService().SetContext(context.Context()).AddAddress(id, utils.CreateToAddress(model.Address)); err != nil {
		return err
	}

	result, err := m.r.CustomerService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /store/customers/me/addresses/{address_id}
// operationId: PostCustomersCustomerAddressesAddress
// summary: "Update a Shipping Address"
// description: "Update the logged-in customer's saved Shipping Address's details."
// x-authenticated: true
// parameters:
//   - (path) address_id=* {String} The ID of the Address.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/StorePostCustomersCustomerAddressesAddressReq"
//
// x-codegen:
//
//	method: updateAddress
//
// x-codeSamples:
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged
//     medusa.customers.addresses.updateAddress(addressId, {
//     first_name: "Gina"
//     })
//     .then(({ customer }) => {
//     console.log(customer.id);
//     })
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/store/customers/me/addresses/{address_id}' \
//     -H 'Authorization: Bearer {access_token}' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "first_name": "Gina"
//     }'
//
// security:
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Customers
//
// responses:
//
//	"200":
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCustomersRes"
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
func (m *Customer) UpdateAddress(context fiber.Ctx) error {
	id := context.Locals("customer_id").(uuid.UUID)

	model, addressId, err := api.BindUpdate[types.AddressPayload](context, "address_id", m.r.Validator())
	if err != nil {
		return err
	}

	if _, err := m.r.CustomerService().SetContext(context.Context()).UpdateAddress(id, addressId, utils.ToAddress(model)); err != nil {
		return err
	}

	result, err := m.r.CustomerService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [delete] /store/customers/me/addresses/{address_id}
// operationId: DeleteCustomersCustomerAddressesAddress
// summary: Delete an Address
// description: "Delete an Address from the Customer's saved addresses."
// x-authenticated: true
// parameters:
//   - (path) address_id=* {string} The id of the Address to remove.
//
// x-codegen:
//
//	method: deleteAddress
//
// x-codeSamples:
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged
//     medusa.customers.addresses.deleteAddress(addressId)
//     .then(({ customer }) => {
//     console.log(customer.id);
//     })
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '{backend_url}/store/customers/me/addresses/{address_id}' \
//     -H 'Authorization: Bearer {access_token}'
//
// security:
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Customers
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCustomersRes"
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
func (m *Customer) DeleteAdress(context fiber.Ctx) error {
	id := context.Locals("customer_id").(uuid.UUID)

	addressId, err := api.BindDelete(context, "address_id")
	if err != nil {
		return err
	}

	if err := m.r.CustomerService().SetContext(context.Context()).RemoveAddress(id, addressId); err != nil {
		return err
	}

	result, err := m.r.CustomerService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /store/customers/password-reset
// operationId: PostCustomersResetPassword
// summary: Reset Password
// description: "Reset a Customer's password using a password token created by a previous request to the Request Password Reset API Route. If the password token expired,
//
//	you must create a new one."
//
// externalDocs:
//
//	description: "How to reset password"
//	url: "https://docs.medusajs.com/modules/customers/storefront/implement-customer-profiles#reset-password"
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/ResetPasswordRequest"
//
// x-codegen:
//
//	method: resetPassword
//
// x-codeSamples:
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.customers.resetPassword({
//     email: "user@example.com",
//     password: "supersecret",
//     token: "supersecrettoken"
//     })
//     .then(({ customer }) => {
//     console.log(customer.id);
//     })
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/store/customers/password-reset' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "email": "user@example.com",
//     "password": "supersecret",
//     "token": "supersecrettoken"
//     }'
//
// tags:
//   - Customers
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCustomersResetPasswordRes"
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
func (m *Customer) ResetPassword(context fiber.Ctx) error {
	model, err := api.BindCreate[types.UserResetPasswordRequest](context, m.r.Validator())
	if err != nil {
		return err
	}

	customer, err := m.r.CustomerService().SetContext(context.Context()).RetrieveRegisteredByEmail(model.Email, &sql.Options{Selects: []string{"id", "password_hash"}})
	if err != nil {
		return err
	}

	tocken, claims, er := m.r.TockenService().VerifyTokenWithSecret(model.Token, []byte(customer.PasswordHash))
	if er != nil {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			er.Error(),
		)
	}

	if tocken == nil || claims["customer_id"] != customer.Id {
		return utils.NewApplictaionError(
			utils.UNAUTHORIZED,
			"Invalid or expired password reset token",
		)
	}

	result, err := m.r.CustomerService().SetContext(context.Context()).Update(customer.Id, &types.UpdateCustomerInput{
		Password: model.Password,
	})
	if err != nil {
		return err
	}

	//TODO: Check If working
	result.PasswordHash = ""
	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /store/customers/password-token
// operationId: PostCustomersCustomerPasswordToken
// summary: Request Password Reset
// description: "Create a reset password token to be used in a subsequent Reset Password API Route. This emits the event `customer.password_reset`. If a notification provider is
//
//	installed in the Medusa backend and is configured to handle this event, a notification to the customer, such as an email, may be sent with reset instructions."
//
// externalDocs:
//
//	description: "How to reset password"
//	url: "https://docs.medusajs.com/modules/customers/storefront/implement-customer-profiles#reset-password"
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/ResetPasswordTokenRequest"
//
// x-codegen:
//
//	method: generatePasswordToken
//
// x-codeSamples:
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.customers.generatePasswordToken({
//     email: "user@example.com"
//     })
//     .then(() => {
//     // successful
//     })
//     .catch(() => {
//     // failed
//     })
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/store/customers/password-token' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "email": "user@example.com"
//     }'
//
// tags:
//   - Customers
//
// responses:
//
//	204:
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
func (m *Customer) ResetPasswordTocken(context fiber.Ctx) error {
	model, err := api.BindCreate[types.UserResetPasswordToken](context, m.r.Validator())
	if err != nil {
		return err
	}

	customer, err := m.r.CustomerService().SetContext(context.Context()).RetrieveRegisteredByEmail(model.Email, &sql.Options{})
	if err != nil {
		return err
	}

	if customer != nil {
		if _, err := m.r.CustomerService().SetContext(context.Context()).GenerateResetPasswordToken(customer.Id); err != nil {
			return err
		}
	}

	return context.SendStatus(204)
}

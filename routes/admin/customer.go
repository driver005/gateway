package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type Customer struct {
	r    Registry
	name string
}

func NewCustomer(r Registry) *Customer {
	m := Customer{r: r, name: "customer"}
	return &m
}

func (m *Customer) SetRoutes(router fiber.Router) {
	route := router.Group("/customers")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
}

// @oas:path [get] /admin/customers/{id}
// operationId: "GetCustomersCustomer"
// summary: "Get a Customer"
// description: "Retrieve the details of a customer."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Customer.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned customer.
//   - (query) fields {string} Comma-separated fields that should be included in the returned customer.
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
//     medusa.admin.customers.retrieve(customerId)
//     .then(({ customer }) => {
//     console.log(customer.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCustomer } from "medusa-react"
//
//     type Props = {
//     customerId: string
//     }
//
//     const Customer = ({ customerId }: Props) => {
//     const { customer, isLoading } = useAdminCustomer(
//     customerId
//     )
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {customer && <span>{customer.first_name}</span>}
//     </div>
//     )
//     }
//
//     export default Customer
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/customers/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Customers
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminCustomersRes"
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
func (m *Customer) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.CustomerService().SetContext(context.Context()).RetrieveById(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [get] /admin/customers
// operationId: "GetCustomers"
// summary: "List Customers"
// description: "Retrieve a list of Customers. The customers can be filtered by fields such as `q` or `groups`. The customers can also be paginated."
// x-authenticated: true
// parameters:
//   - (query) limit=50 {integer} The number of customers to return.
//   - (query) offset=0 {integer} The number of customers to skip when retrieving the customers.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned customer.
//   - (query) q {string} term to search customers' email, first_name, and last_name fields.
//   - in: query
//     name: groups
//     style: form
//     explode: false
//     description: Filter by customer group IDs.
//     schema:
//     type: array
//     items:
//     type: string
//
// x-codegen:
//
//	method: list
//	queryParams: AdminGetCustomersParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.customers.list()
//     .then(({ customers, limit, offset, count }) => {
//     console.log(customers.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCustomers } from "medusa-react"
//
//     const Customers = () => {
//     const { customers, isLoading } = useAdminCustomers()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {customers && !customers.length && (
//     <span>No customers</span>
//     )}
//     {customers && customers.length > 0 && (
//     <ul>
//     {customers.map((customer) => (
//     <li key={customer.id}>{customer.first_name}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default Customers
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/customers' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Customers
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminCustomersListRes"
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
func (m *Customer) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableCustomer](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.CustomerService().SetContext(context.Context()).ListAndCount(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"customers": result,
		"count":     count,
		"offset":    config.Skip,
		"limit":     config.Take,
	})
}

// @oas:path [post] /admin/customers
// operationId: "PostCustomers"
// summary: "Create a Customer"
// description: "Create a customer as an admin."
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostCustomersReq"
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
//     medusa.admin.customers.create({
//     email: "user@example.com",
//     first_name: "Caterina",
//     last_name: "Yost",
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
//     import { useAdminCreateCustomer } from "medusa-react"
//
//     type CustomerData = {
//     first_name: string
//     last_name: string
//     email: string
//     password: string
//     }
//
//     const CreateCustomer = () => {
//     const createCustomer = useAdminCreateCustomer()
//     // ...
//
//     const handleCreate = (customerData: CustomerData) => {
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
//     export default CreateCustomer
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/customers' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "email": "user@example.com",
//     "first_name": "Caterina",
//     "last_name": "Yost",
//     "password": "supersecret"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Customers
//
// responses:
//
//	201:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminCustomersRes"
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
func (m *Customer) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateCustomerInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.CustomerService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/customers/{id}
// operationId: "PostCustomersCustomer"
// summary: "Update a Customer"
// description: "Update a Customer's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Customer.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned customer.
//   - (query) fields {string} Comma-separated fields that should be retrieved in the returned customer.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostCustomersCustomerReq"
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
//     medusa.admin.customers.update(customerId, {
//     first_name: "Dolly"
//     })
//     .then(({ customer }) => {
//     console.log(customer.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdateCustomer } from "medusa-react"
//
//     type CustomerData = {
//     first_name: string
//     last_name: string
//     email: string
//     password: string
//     }
//
//     type Props = {
//     customerId: string
//     }
//
//     const Customer = ({ customerId }: Props) => {
//     const updateCustomer = useAdminUpdateCustomer(customerId)
//     // ...
//
//     const handleUpdate = (customerData: CustomerData) => {
//     updateCustomer.mutate(customerData)
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
//     curl -X POST '"{backend_url}"/admin/customers/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "first_name": "Dolly"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Customers
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminCustomersRes"
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
func (m *Customer) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateCustomerInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.CustomerService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

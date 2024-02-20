package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type CustomerGroup struct {
	r    Registry
	name string
}

func NewCustomerGroup(r Registry) *CustomerGroup {
	m := CustomerGroup{r: r, name: "customer_group"}
	return &m
}

func (m *CustomerGroup) SetRoutes(router fiber.Router) {
	route := router.Group("/customer-groups")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)

	route.Get("/:id/customers", m.GetBatch)
	route.Post("/:id/customers/batch", m.AddCustomers)
	route.Delete("/:id/customers/batch", m.DeleteCustomers)
}

// @oas:path [get] /admin/customer-groups/{id}
// operationId: "GetCustomerGroupsGroup"
// summary: "Get a Customer Group"
// description: "Retrieve a Customer Group by its ID. You can expand the customer group's relations or select the fields that should be returned."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Customer Group.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned customer group.
//   - (query) fields {string} Comma-separated fields that should be included in the returned customer group.
//
// x-codegen:
//
//	method: retrieve
//	queryParams: AdminGetCustomerGroupsGroupParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.customerGroups.retrieve(customerGroupId)
//     .then(({ customer_group }) => {
//     console.log(customer_group.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCustomerGroup } from "medusa-react"
//
//     type Props = {
//     customerGroupId: string
//     }
//
//     const CustomerGroup = ({ customerGroupId }: Props) => {
//     const { customer_group, isLoading } = useAdminCustomerGroup(
//     customerGroupId
//     )
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {customer_group && <span>{customer_group.name}</span>}
//     </div>
//     )
//     }
//
//     export default CustomerGroup
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/customer-groups/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Customer Groups
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminCustomerGroupsRes"
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
func (m *CustomerGroup) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.CustomerGroupService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [get] /admin/customer-groups
// operationId: "GetCustomerGroups"
// summary: "List Customer Groups"
// description: "Retrieve a list of customer groups. The customer groups can be filtered by fields such as `name` or `id. The customer groups can also be sorted or paginated."
// x-authenticated: true
// parameters:
//   - (query) q {string} term to search customer groups by name.
//   - (query) offset=0 {integer} The number of customer groups to skip when retrieving the customer groups.
//   - (query) order {string} A field to sort order the retrieved customer groups by.
//   - (query) discount_condition_id {string} Filter by discount condition ID.
//   - in: query
//     name: id
//     style: form
//     explode: false
//     description: Filter by the customer group ID
//     schema:
//     oneOf:
//   - type: string
//     description: customer group ID
//   - type: array
//     description: an array of customer group IDs
//     items:
//     type: string
//   - type: object
//     properties:
//     lt:
//     type: string
//     description: filter by IDs less than this ID
//     gt:
//     type: string
//     description: filter by IDs greater than this ID
//     lte:
//     type: string
//     description: filter by IDs less than or equal to this ID
//     gte:
//     type: string
//     description: filter by IDs greater than or equal to this ID
//   - in: query
//     name: name
//     style: form
//     explode: false
//     description: Filter by the customer group name
//     schema:
//     type: array
//     description: an array of customer group names
//     items:
//     type: string
//     description: customer group name
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
//   - (query) limit=10 {integer} The number of customer groups to return.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned customer groups.
//
// x-codegen:
//
//	method: list
//	queryParams: AdminGetCustomerGroupsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.customerGroups.list()
//     .then(({ customer_groups, limit, offset, count }) => {
//     console.log(customer_groups.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCustomerGroups } from "medusa-react"
//
//     const CustomerGroups = () => {
//     const {
//     customer_groups,
//     isLoading,
//     } = useAdminCustomerGroups()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {customer_groups && !customer_groups.length && (
//     <span>No Customer Groups</span>
//     )}
//     {customer_groups && customer_groups.length > 0 && (
//     <ul>
//     {customer_groups.map(
//     (customerGroup) => (
//     <li key={customerGroup.id}>
//     {customerGroup.name}
//     </li>
//     )
//     )}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default CustomerGroups
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/customer-groups' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Customer Groups
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminCustomerGroupsListRes"
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
func (m *CustomerGroup) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableCustomerGroup](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.CustomerGroupService().SetContext(context.Context()).ListAndCount(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"customer_groups": result,
		"count":           count,
		"offset":          config.Skip,
		"limit":           config.Take,
	})
}

// @oas:path [post] /admin/customer-groups
// operationId: "PostCustomerGroups"
// summary: "Create a Customer Group"
// description: "Create a Customer Group."
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostCustomerGroupsReq"
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
//     medusa.admin.customerGroups.create({
//     name: "VIP"
//     })
//     .then(({ customer_group }) => {
//     console.log(customer_group.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreateCustomerGroup } from "medusa-react"
//
//     const CreateCustomerGroup = () => {
//     const createCustomerGroup = useAdminCreateCustomerGroup()
//     // ...
//
//     const handleCreate = (name: string) => {
//     createCustomerGroup.mutate({
//     name,
//     })
//     }
//
//     // ...
//     }
//
//     export default CreateCustomerGroup
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/customer-groups' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "name": "VIP"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Customer Groups
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminCustomerGroupsRes"
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
func (m *CustomerGroup) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateCustomerGroup](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.CustomerGroupService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/customer-groups/{id}
// operationId: "PostCustomerGroupsGroup"
// summary: "Update a Customer Group"
// description: "Update a Customer Group's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the customer group.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostCustomerGroupsGroupReq"
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
//     medusa.admin.customerGroups.update(customerGroupId, {
//     name: "VIP"
//     })
//     .then(({ customer_group }) => {
//     console.log(customer_group.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdateCustomerGroup } from "medusa-react"
//
//     type Props = {
//     customerGroupId: string
//     }
//
//     const CustomerGroup = ({ customerGroupId }: Props) => {
//     const updateCustomerGroup = useAdminUpdateCustomerGroup(
//     customerGroupId
//     )
//     // ..
//
//     const handleUpdate = (name: string) => {
//     updateCustomerGroup.mutate({
//     name,
//     })
//     }
//
//     // ...
//     }
//
//     export default CustomerGroup
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/customer-groups/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "name": "VIP"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Customer Groups
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminCustomerGroupsRes"
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
func (m *CustomerGroup) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateCustomerGroup](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.CustomerGroupService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/customer-groups/{id}
// operationId: "DeleteCustomerGroupsCustomerGroup"
// summary: "Delete a Customer Group"
// description: "Delete a customer group. This doesn't delete the customers associated with the customer group."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Customer Group
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
//     medusa.admin.customerGroups.delete(customerGroupId)
//     .then(({ id, object, deleted }) => {
//     console.log(id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteCustomerGroup } from "medusa-react"
//
//     type Props = {
//     customerGroupId: string
//     }
//
//     const CustomerGroup = ({ customerGroupId }: Props) => {
//     const deleteCustomerGroup = useAdminDeleteCustomerGroup(
//     customerGroupId
//     )
//     // ...
//
//     const handleDeleteCustomerGroup = () => {
//     deleteCustomerGroup.mutate()
//     }
//
//     // ...
//     }
//
//     export default CustomerGroup
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/customer-groups/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Customer Groups
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminCustomerGroupsDeleteRes"
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
func (m *CustomerGroup) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.CustomerGroupService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "cusomer-group",
		"deleted": true,
	})
}

// @oas:path [post] /admin/customer-groups/{id}/customers/batch
// operationId: "PostCustomerGroupsGroupCustomersBatch"
// summary: "Add Customers to Group"
// description: "Add a list of customers to a customer group."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the customer group.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostCustomerGroupsGroupCustomersBatchReq"
//
// x-codegen:
//
//	method: addCustomers
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.customerGroups.addCustomers(customerGroupId, {
//     customer_ids: [
//     {
//     id: customerId
//     }
//     ]
//     })
//     .then(({ customer_group }) => {
//     console.log(customer_group.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import {
//     useAdminAddCustomersToCustomerGroup,
//     } from "medusa-react"
//
//     type Props = {
//     customerGroupId: string
//     }
//
//     const CustomerGroup = ({ customerGroupId }: Props) => {
//     const addCustomers = useAdminAddCustomersToCustomerGroup(
//     customerGroupId
//     )
//     // ...
//
//     const handleAddCustomers= (customerId: string) => {
//     addCustomers.mutate({
//     customer_ids: [
//     {
//     id: customerId,
//     },
//     ],
//     })
//     }
//
//     // ...
//     }
//
//     export default CustomerGroup
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/customer-groups/{id}/customers/batch' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "customer_ids": [
//     {
//     "id": "cus_01G2Q4BS9GAHDBMDEN4ZQZCJB2"
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
//   - Customer Groups
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminCustomerGroupsRes"
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
func (m *CustomerGroup) AddCustomers(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.CustomersToCustomerGroup](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.CustomerGroupService().SetContext(context.Context()).AddCustomers(id, model.CustomerIds)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [get] /admin/customer-groups/{id}/customers
// operationId: "GetCustomerGroupsGroupCustomers"
// summary: "List Customers"
// description: "Retrieve a list of customers in a customer group. The customers can be filtered by the `q` field. The customers can also be paginated."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the customer group.
//   - (query) limit=50 {integer} The number of customers to return.
//   - (query) offset=0 {integer} The number of customers to skip when retrieving the customers.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned customers.
//   - (query) q {string} a term to search customers by email, first_name, and last_name.
//
// x-codegen:
//
//	method: listCustomers
//	queryParams: AdminGetGroupsGroupCustomersParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.customerGroups.listCustomers(customerGroupId)
//     .then(({ customers }) => {
//     console.log(customers.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCustomerGroupCustomers } from "medusa-react"
//
//     type Props = {
//     customerGroupId: string
//     }
//
//     const CustomerGroup = ({ customerGroupId }: Props) => {
//     const {
//     customers,
//     isLoading,
//     } = useAdminCustomerGroupCustomers(
//     customerGroupId
//     )
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
//     export default CustomerGroup
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/customer-groups/{id}/customers' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Customer Groups
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
func (m *CustomerGroup) GetBatch(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableCustomerGroup](context)
	if err != nil {
		return err
	}
	groups, count, err := m.r.CustomerGroupService().SetContext(context.Context()).ListAndCount(model, config)
	if err != nil {
		return err
	}

	var result []models.Customer
	for _, group := range groups {
		result = append(result, group.Customers...)
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   result,
		"count":  count,
		"offset": config.Skip,
		"limit":  config.Take,
	})
}

// @oas:path [delete] /admin/customer-groups/{id}/customers/batch
// operationId: "DeleteCustomerGroupsGroupCustomerBatch"
// summary: "Remove Customers from Group"
// description: "Remove a list of customers from a customer group. This doesn't delete the customer, only the association between the customer and the customer group."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the customer group.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminDeleteCustomerGroupsGroupCustomerBatchReq"
//
// x-codegen:
//
//	method: removeCustomers
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.customerGroups.removeCustomers(customerGroupId, {
//     customer_ids: [
//     {
//     id: customerId
//     }
//     ]
//     })
//     .then(({ customer_group }) => {
//     console.log(customer_group.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import {
//     useAdminRemoveCustomersFromCustomerGroup,
//     } from "medusa-react"
//
//     type Props = {
//     customerGroupId: string
//     }
//
//     const CustomerGroup = ({ customerGroupId }: Props) => {
//     const removeCustomers =
//     useAdminRemoveCustomersFromCustomerGroup(
//     customerGroupId
//     )
//     // ...
//
//     const handleRemoveCustomer = (customerId: string) => {
//     removeCustomers.mutate({
//     customer_ids: [
//     {
//     id: customerId,
//     },
//     ],
//     })
//     }
//
//     // ...
//     }
//
//     export default CustomerGroup
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/customer-groups/{id}/customers/batch' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "customer_ids": [
//     {
//     "id": "cus_01G2Q4BS9GAHDBMDEN4ZQZCJB2"
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
//   - Customer Groups
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminCustomerGroupsRes"
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
func (m *CustomerGroup) DeleteCustomers(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.CustomersToCustomerGroup](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.CustomerGroupService().SetContext(context.Context()).RemoveCustomer(id, model.CustomerIds)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

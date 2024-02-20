package admin

import (
	"fmt"

	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
)

type Discount struct {
	r    Registry
	name string
}

func NewDiscount(r Registry) *Discount {
	m := Discount{r: r, name: "discount"}
	return &m
}

func (m *Discount) SetRoutes(router fiber.Router) {
	route := router.Group("/discounts")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)

	route.Get("/code/:code", m.GetDiscountByCode)

	// Dynamic codes
	route.Post("/:id/dynamic-codes", m.CreateDynamicCode)
	route.Delete("/:id/dynamic-codes/:code", m.DeleteDynamicCode)

	// Discount region management
	route.Post("/:id/regions/:region_id", m.AddRegion)
	route.Delete("/:id/regions/:region_id", m.RemoveRegion)

	// Discount condition management
	route.Post("/:id/conditions", m.CreateConditon)
	route.Delete("/:id/conditions/:condition_id", m.DeleteConditon)

	route.Get("/:id/conditions/:condition_id", m.GetConditon)
	route.Post("/:id/conditions/:condition_id", m.UpdateConditon)
	route.Post("/:id/conditions/:condition_id/batch", m.AddResourcesToConditionBatch)
	route.Delete("/:id/conditions/:condition_id/batch", m.DeleteResourcesToConditionBatch)
}

// @oas:path [get] /admin/discounts/{id}
// operationId: "GetDiscountsDiscount"
// summary: "Get a Discount"
// description: "Retrieve a Discount."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Discount
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned discount.
//   - (query) fields {string} Comma-separated fields that should be included in the returned discount.
//
// x-codegen:
//
//	method: retrieve
//	queryParams: AdminGetDiscountParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.discounts.retrieve(discountId)
//     .then(({ discount }) => {
//     console.log(discount.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDiscount } from "medusa-react"
//
//     type Props = {
//     discountId: string
//     }
//
//     const Discount = ({ discountId }: Props) => {
//     const { discount, isLoading } = useAdminDiscount(
//     discountId
//     )
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {discount && <span>{discount.code}</span>}
//     </div>
//     )
//     }
//
//     export default Discount
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/discounts/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Discounts
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDiscountsRes"
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
func (m *Discount) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.DiscountService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [get] /admin/discounts
// operationId: "GetDiscounts"
// summary: "List Discounts"
// x-authenticated: true
// description: "Retrieve a list of Discounts. The discounts can be filtered by fields such as `rule` or `is_dynamic`. The discounts can also be paginated."
// parameters:
//   - (query) q {string} term to search discounts' code field.
//   - in: query
//     name: rule
//     description: Filter discounts by rule fields.
//     schema:
//     type: object
//     properties:
//     type:
//     type: string
//     enum: [fixed, percentage, free_shipping]
//     description: "Filter discounts by type."
//     allocation:
//     type: string
//     enum: [total, item]
//     description: "Filter discounts by allocation type."
//   - (query) is_dynamic {boolean} Filter discounts by whether they're dynamic or not.
//   - (query) is_disabled {boolean} Filter discounts by whether they're disabled or not.
//   - (query) limit=20 {number} The number of discounts to return
//   - (query) offset=0 {number} The number of discounts to skip when retrieving the discounts.
//   - (query) expand {string} Comma-separated relations that should be expanded in each returned discount.
//
// x-codegen:
//
//	method: list
//	queryParams: AdminGetDiscountsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.discounts.list()
//     .then(({ discounts, limit, offset, count }) => {
//     console.log(discounts.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDiscounts } from "medusa-react"
//
//     const Discounts = () => {
//     const { discounts, isLoading } = useAdminDiscounts()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {discounts && !discounts.length && (
//     <span>No customers</span>
//     )}
//     {discounts && discounts.length > 0 && (
//     <ul>
//     {discounts.map((discount) => (
//     <li key={discount.id}>{discount.code}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default Discounts
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/discounts' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Discounts
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDiscountsListRes"
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
func (m *Discount) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableDiscount](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.DiscountService().SetContext(context.Context()).ListAndCount(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"discounts": result,
		"count":     count,
		"offset":    config.Skip,
		"limit":     config.Take,
	})
}

//
// @oas:path [post] /admin/discounts
// operationId: "PostDiscounts"
// summary: "Create a Discount"
// x-authenticated: true
// description: "Create a Discount with a given set of rules that defines how the Discount is applied."
// parameters:
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned discount.
//   - (query) fields {string} Comma-separated fields that should be retrieved in the returned discount.
// requestBody:
//   content:
//     application/json:
//       schema:
//         $ref: "#/components/schemas/AdminPostDiscountsReq"
// x-codegen:
//   method: create
//   queryParams: AdminPostDiscountsParams
// x-codeSamples:
//   - lang: JavaScript
//     label: JS Client
//     source: |
//       import Medusa from "@medusajs/medusa-js"
//       import { AllocationType, DiscountRuleType } from "@medusajs/medusa"
//       const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//       // must be previously logged in or use api token
//       medusa.admin.discounts.create({
//         code: "TEST",
//         rule: {
//           type: DiscountRuleType.FIXED,
//           value: 10,
//           allocation: AllocationType.ITEM
//         },
//         regions: ["reg_XXXXXXXX"],
//         is_dynamic: false,
//         is_disabled: false
//       })
//       .then(({ discount }) => {
//         console.log(discount.id);
//       })
//   - lang: tsx
//     label: Medusa React
//     source: |
//       import React from "react"
//       import {
//         useAdminCreateDiscount,
//       } from "medusa-react"
//       import {
//         AllocationType,
//         DiscountRuleType,
//       } from "@medusajs/medusa"
//
//       const CreateDiscount = () => {
//         const createDiscount = useAdminCreateDiscount()
//         // ...
//
//         const handleCreate = (
//           currencyCode: string,
//           regionId: string
//         ) => {
//           // ...
//           createDiscount.mutate({
//             code: currencyCode,
//             rule: {
//               type: DiscountRuleType.FIXED,
//               value: 10,
//               allocation: AllocationType.ITEM,
//             },
//             regions: [
//                 regionId,
//             ],
//             is_dynamic: false,
//             is_disabled: false,
//           })
//         }
//
//         // ...
//       }
//
//       export default CreateDiscount
//   - lang: Shell
//     label: cURL
//     source: |
//       curl -X POST '"{backend_url}"/admin/discounts' \
//       -H 'x-medusa-access-token: "{api_token}"' \
//       -H 'Content-Type: application/json' \
//       --data-raw '{
//           "code": "TEST",
//           "rule": {
//              "type": "fixed",
//              "value": 10,
//              "allocation": "item"
//           },
//           "regions": ["reg_XXXXXXXX"]
//       }'
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
// tags:
//   - Discounts
// responses:
//   200:
//     description: "OK"
//     content:
//       application/json:
//         schema:
//           $ref: "#/components/schemas/AdminDiscountsRes"
//   "400":
//     $ref: "#/components/responses/400_error"
//   "401":
//     $ref: "#/components/responses/unauthorized"
//   "404":
//     $ref: "#/components/responses/not_found_error"
//   "409":
//     $ref: "#/components/responses/invalid_state_error"
//   "422":
//     $ref: "#/components/responses/invalid_request_error"
//   "500":
//     $ref: "#/components/responses/500_error"
//

func (m *Discount) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateDiscountInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.DiscountService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/discounts/{id}
// operationId: "PostDiscountsDiscount"
// summary: "Update a Discount"
// description: "Update a Discount with a given set of rules that define how the Discount is applied."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Discount.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned discount.
//   - (query) fields {string} Comma-separated fields that should be retrieved in the returned discount.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostDiscountsDiscountReq"
//
// x-codegen:
//
//	method: update
//	queryParams: AdminPostDiscountsDiscountParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.discounts.update(discountId, {
//     code: "TEST"
//     })
//     .then(({ discount }) => {
//     console.log(discount.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdateDiscount } from "medusa-react"
//
//     type Props = {
//     discountId: string
//     }
//
//     const Discount = ({ discountId }: Props) => {
//     const updateDiscount = useAdminUpdateDiscount(discountId)
//     // ...
//
//     const handleUpdate = (isDisabled: boolean) => {
//     updateDiscount.mutate({
//     is_disabled: isDisabled,
//     })
//     }
//
//     // ...
//     }
//
//     export default Discount
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/discounts/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "code": "TEST"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Discounts
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDiscountsRes"
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
func (m *Discount) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateDiscountInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.DiscountService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/discounts/{id}
// operationId: "DeleteDiscountsDiscount"
// summary: "Delete a Discount"
// description: "Delete a Discount. Deleting the discount will make it unavailable for customers to use."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Discount
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
//     medusa.admin.discounts.delete(discountId)
//     .then(({ id, object, deleted }) => {
//     console.log(id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteDiscount } from "medusa-react"
//
//     const Discount = () => {
//     const deleteDiscount = useAdminDeleteDiscount(discount_id)
//     // ...
//
//     const handleDelete = () => {
//     deleteDiscount.mutate()
//     }
//
//     // ...
//     }
//
//     export default Discount
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/discounts/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Discounts
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDiscountsDeleteRes"
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
func (m *Discount) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.DiscountService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "disocunt",
		"deleted": true,
	})
}

// @oas:path [get] /admin/discounts/{discount_id}/conditions/{condition_id}
// operationId: "GetDiscountsDiscountConditionsCondition"
// summary: "Get a Condition"
// description: "Retrieve a Discount Condition's details."
// x-authenticated: true
// parameters:
//   - (path) discount_id=* {string} The ID of the Discount.
//   - (path) condition_id=* {string} The ID of the Discount Condition.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned discount condition.
//   - (query) fields {string} Comma-separated fields that should be included in the returned discount condition.
//
// x-codegen:
//
//	method: getCondition
//	queryParams: AdminGetDiscountsDiscountConditionsConditionParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.discounts.getCondition(discountId, conditionId)
//     .then(({ discount_condition }) => {
//     console.log(discount_condition.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminGetDiscountCondition } from "medusa-react"
//
//     type Props = {
//     discountId: string
//     discountConditionId: string
//     }
//
//     const DiscountCondition = ({
//     discountId,
//     discountConditionId
//     }: Props) => {
//     const {
//     discount_condition,
//     isLoading
//     } = useAdminGetDiscountCondition(
//     discountId,
//     discountConditionId
//     )
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {discount_condition && (
//     <span>{discount_condition.type}</span>
//     )}
//     </div>
//     )
//     }
//
//     export default DiscountCondition
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/discounts/{id}/conditions/{condition_id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Discounts
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDiscountConditionsRes"
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
func (m *Discount) GetConditon(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "condition_id")
	if err != nil {
		return err
	}

	result, err := m.r.DiscountConditionService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [get] /admin/discounts/code/{code}
// operationId: "GetDiscountsDiscountCode"
// summary: "Get Discount by Code"
// description: "Retrieve a Discount's details by its discount code"
// x-authenticated: true
// parameters:
//   - (path) code=* {string} The code of the Discount
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned discount.
//   - (query) fields {string} Comma-separated fields that should be included in the returned discount.
//
// x-codegen:
//
//	method: retrieveByCode
//	queryParams: AdminGetDiscountsDiscountCodeParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.discounts.retrieveByCode(code)
//     .then(({ discount }) => {
//     console.log(discount.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminGetDiscountByCode } from "medusa-react"
//
//     type Props = {
//     discountCode: string
//     }
//
//     const Discount = ({ discountCode }: Props) => {
//     const { discount, isLoading } = useAdminGetDiscountByCode(
//     discountCode
//     )
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {discount && <span>{discount.code}</span>}
//     </div>
//     )
//     }
//
//     export default Discount
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/discounts/code/{code}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Discounts
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDiscountsRes"
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
func (m *Discount) GetDiscountByCode(context fiber.Ctx) error {
	config, err := api.BindConfig(context, m.r.Validator())
	if err != nil {
		return err
	}

	code := context.Params("code")

	result, err := m.r.DiscountService().SetContext(context.Context()).RetrieveByCode(code, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/discounts/{id}/regions/{region_id}
// operationId: "PostDiscountsDiscountRegionsRegion"
// summary: "Add Region to Discount"
// description: "Add a Region to the list of Regions a Discount can be used in."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Discount.
//   - (path) region_id=* {string} The ID of the Region.
//
// x-codegen:
//
//	method: addRegion
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.discounts.addRegion(discountId, regionId)
//     .then(({ discount }) => {
//     console.log(discount.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDiscountAddRegion } from "medusa-react"
//
//     type Props = {
//     discountId: string
//     }
//
//     const Discount = ({ discountId }: Props) => {
//     const addRegion = useAdminDiscountAddRegion(discountId)
//     // ...
//
//     const handleAdd = (regionId: string) => {
//     addRegion.mutate(regionId, {
//     onSuccess: ({ discount }) => {
//     console.log(discount.regions)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Discount
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/discounts/{id}/regions/{region_id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Discounts
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDiscountsRes"
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
func (m *Discount) AddRegion(context fiber.Ctx) error {
	id, err := utils.ParseUUID(context.Params("id"))
	if err != nil {
		return err
	}

	regionId, err := utils.ParseUUID(context.Params("region_id"))
	if err != nil {
		return err
	}

	result, err := m.r.DiscountService().SetContext(context.Context()).AddRegion(id, regionId)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/discounts/{discount_id}/conditions/{condition_id}/batch
// operationId: "PostDiscountsDiscountConditionsConditionBatch"
// summary: "Add Batch Resources"
// description: "Add a batch of resources to a discount condition. The type of resource depends on the type of discount condition. For example, if the discount condition's type is `products`,
// the resources being added should be products."
// x-authenticated: true
// parameters:
//   - (path) discount_id=* {string} The ID of the discount the condition belongs to.
//   - (path) condition_id=* {string} The ID of the discount condition on which to add the item.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned discount.
//   - (query) fields {string} Comma-separated fields that should be included in the returned discount.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostDiscountsDiscountConditionsConditionBatchReq"
//
// x-codegen:
//
//	method: addConditionResourceBatch
//	queryParams: AdminPostDiscountsDiscountConditionsConditionBatchParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.discounts.addConditionResourceBatch(discountId, conditionId, {
//     resources: [{ id: itemId }]
//     })
//     .then(({ discount }) => {
//     console.log(discount.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import {
//     useAdminAddDiscountConditionResourceBatch
//     } from "medusa-react"
//
//     type Props = {
//     discountId: string
//     conditionId: string
//     }
//
//     const DiscountCondition = ({
//     discountId,
//     conditionId
//     }: Props) => {
//     const addConditionResources = useAdminAddDiscountConditionResourceBatch(
//     discountId,
//     conditionId
//     )
//     // ...
//
//     const handleAdd = (itemId: string) => {
//     addConditionResources.mutate({
//     resources: [
//     {
//     id: itemId
//     }
//     ]
//     }, {
//     onSuccess: ({ discount }) => {
//     console.log(discount.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default DiscountCondition
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/discounts/{id}/conditions/{condition_id}/batch' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "resources": [{ "id": "item_id" }]
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Discounts
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDiscountsRes"
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
func (m *Discount) AddResourcesToConditionBatch(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.AddResourcesToConditionBatch](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	conditionId, err := utils.ParseUUID(context.Params("condition_id"))
	if err != nil {
		return err
	}

	condition, err := m.r.DiscountConditionService().SetContext(context.Context()).Retrieve(conditionId, &sql.Options{Selects: []string{"id", "type", "discount_rule_id"}})
	if err != nil {
		return err
	}

	input := &types.DiscountConditionInput{Id: conditionId, RuleId: condition.DiscountRuleId.UUID}
	if condition.Type == models.DiscountConditionTypeProducts {
		input.Products = model.Resources
	} else if condition.Type == models.DiscountConditionTypeProductTypes {
		input.ProductTypes = model.Resources
	} else if condition.Type == models.DiscountConditionTypeProductTags {
		input.ProductTags = model.Resources
	} else if condition.Type == models.DiscountConditionTypeProductCollections {
		input.ProductCollections = model.Resources
	} else if condition.Type == models.DiscountConditionTypeCustomerGroups {
		input.CustomerGroups = model.Resources
	}

	_, err = m.r.DiscountConditionService().SetContext(context.Context()).UpsertCondition(input, false)
	if err != nil {
		return err
	}

	result, err := m.r.DiscountService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/discounts/{discount_id}/conditions
// operationId: "PostDiscountsDiscountConditions"
// summary: "Create a Condition"
// description: "Create a Discount Condition. Only one of `products`, `product_types`, `product_collections`, `product_tags`, and `customer_groups` should be provided, based on the type of discount condition.
//
//	For example, if the discount condition's type is `products`, the `products` field should be provided in the request body."
//
// x-authenticated: true
// parameters:
//   - (path) discount_id=* {string} The ID of the discount.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned discount.
//   - (query) fields {string} Comma-separated fields that should be included in the returned discount.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostDiscountsDiscountConditions"
//
// x-codegen:
//
//	method: createCondition
//	queryParams: AdminPostDiscountsDiscountConditionsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     import { DiscountConditionOperator } from "@medusajs/medusa"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.discounts.createCondition(discountId, {
//     operator: DiscountConditionOperator.IN,
//     products: [productId]
//     })
//     .then(({ discount }) => {
//     console.log(discount.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { DiscountConditionOperator } from "@medusajs/medusa"
//     import { useAdminDiscountCreateCondition } from "medusa-react"
//
//     type Props = {
//     discountId: string
//     }
//
//     const Discount = ({ discountId }: Props) => {
//     const createCondition = useAdminDiscountCreateCondition(discountId)
//     // ...
//
//     const handleCreateCondition = (
//     operator: DiscountConditionOperator,
//     products: string[]
//     ) => {
//     createCondition.mutate({
//     operator,
//     products
//     }, {
//     onSuccess: ({ discount }) => {
//     console.log(discount.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Discount
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/discounts/{id}/conditions' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "operator": "in"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Discounts
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDiscountsRes"
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
func (m *Discount) CreateConditon(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.CreateConditon](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	discount, err := m.r.DiscountService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	_, err = m.r.DiscountConditionService().SetContext(context.Context()).UpsertCondition(&types.DiscountConditionInput{Operator: model.Operator, RuleId: discount.RuleId.UUID}, false)
	if err != nil {
		return err
	}

	result, err := m.r.DiscountService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/discounts/{id}/dynamic-codes
// operationId: "PostDiscountsDiscountDynamicCodes"
// summary: "Create a Dynamic Code"
// description: "Create a dynamic unique code that can map to a parent Discount. This is useful if you want to automatically generate codes with the same rules and conditions."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Discount to create the dynamic code for."
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostDiscountsDiscountDynamicCodesReq"
//
// x-codegen:
//
//	method: createDynamicCode
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.discounts.createDynamicCode(discountId, {
//     code: "TEST",
//     usage_limit: 1
//     })
//     .then(({ discount }) => {
//     console.log(discount.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreateDynamicDiscountCode } from "medusa-react"
//
//     type Props = {
//     discountId: string
//     }
//
//     const Discount = ({ discountId }: Props) => {
//     const createDynamicDiscount = useAdminCreateDynamicDiscountCode(discountId)
//     // ...
//
//     const handleCreate = (
//     code: string,
//     usageLimit: number
//     ) => {
//     createDynamicDiscount.mutate({
//     code,
//     usage_limit: usageLimit
//     }, {
//     onSuccess: ({ discount }) => {
//     console.log(discount.is_dynamic)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Discount
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/discounts/{id}/dynamic-codes' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "code": "TEST"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Discounts
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDiscountsRes"
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
func (m *Discount) CreateDynamicCode(context fiber.Ctx) error {
	model, id, err := api.BindWithUUID[types.CreateDynamicDiscountInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	_, err = m.r.DiscountService().SetContext(context.Context()).CreateDynamicCode(id, model)
	if err != nil {
		return err
	}

	result, err := m.r.DiscountService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/discounts/{discount_id}/conditions/{condition_id}
// operationId: "PostDiscountsDiscountConditionsCondition"
// summary: "Update a Condition"
// description: "Update a Discount Condition. Only one of `products`, `product_types`, `product_collections`, `product_tags`, and `customer_groups` should be provided, based on the type of discount condition.
//
//	For example, if the discount condition's type is `products`, the `products` field should be provided in the request body."
//
// x-authenticated: true
// parameters:
//   - (path) discount_id=* {string} The ID of the Discount.
//   - (path) condition_id=* {string} The ID of the Discount Condition.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned discount.
//   - (query) fields {string} Comma-separated fields that should be included in the returned discount.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostDiscountsDiscountConditionsCondition"
//
// x-codegen:
//
//	method: updateCondition
//	queryParams: AdminPostDiscountsDiscountConditionsConditionParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.discounts.updateCondition(discountId, conditionId, {
//     products: [
//     productId
//     ]
//     })
//     .then(({ discount }) => {
//     console.log(discount.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDiscountUpdateCondition } from "medusa-react"
//
//     type Props = {
//     discountId: string
//     conditionId: string
//     }
//
//     const DiscountCondition = ({
//     discountId,
//     conditionId
//     }: Props) => {
//     const update = useAdminDiscountUpdateCondition(
//     discountId,
//     conditionId
//     )
//     // ...
//
//     const handleUpdate = (
//     products: string[]
//     ) => {
//     update.mutate({
//     products
//     }, {
//     onSuccess: ({ discount }) => {
//     console.log(discount.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default DiscountCondition
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/discounts/{id}/conditions/{condition}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "products": [
//     "prod_01G1G5V2MBA328390B5AXJ610F"
//     ]
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Discounts
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDiscountsRes"
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
func (m *Discount) UpdateConditon(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.AdminUpsertConditionsReq](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	conditionId, err := utils.ParseUUID(context.Params("condition_id"))
	if err != nil {
		return err
	}

	condition, err := m.r.DiscountConditionService().SetContext(context.Context()).Retrieve(conditionId, &sql.Options{Selects: []string{"id"}})
	if err != nil {
		return err
	}

	discount, err := m.r.DiscountService().SetContext(context.Context()).Retrieve(id, &sql.Options{Selects: []string{}})
	if err != nil {
		return err
	}

	_, err = m.r.DiscountConditionService().SetContext(context.Context()).UpsertCondition(&types.DiscountConditionInput{
		Id:                 condition.Id,
		RuleId:             discount.RuleId.UUID,
		Products:           model.Products,
		ProductCollections: model.ProductCollections,
		ProductTypes:       model.ProductTypes,
		ProductTags:        model.ProductTags,
		CustomerGroups:     model.CustomerGroups,
	}, false)
	if err != nil {
		return err
	}

	result, err := m.r.DiscountService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/discounts/{discount_id}/conditions/{condition_id}
// operationId: "DeleteDiscountsDiscountConditionsCondition"
// summary: "Delete a Condition"
// description: "Delete a Discount Condition. This does not delete resources associated to the discount condition."
// x-authenticated: true
// parameters:
//   - (path) discount_id=* {string} The ID of the Discount
//   - (path) condition_id=* {string} The ID of the Discount Condition
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned discount.
//   - (query) fields {string} Comma-separated fields that should be included in the returned discount.
//
// x-codegen:
//
//	method: deleteCondition
//	queryParams: AdminDeleteDiscountsDiscountConditionsConditionParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.discounts.deleteCondition(discountId, conditionId)
//     .then(({ id, object, deleted }) => {
//     console.log(id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import {
//     useAdminDiscountRemoveCondition
//     } from "medusa-react"
//
//     type Props = {
//     discountId: string
//     }
//
//     const Discount = ({ discountId }: Props) => {
//     const deleteCondition = useAdminDiscountRemoveCondition(
//     discountId
//     )
//     // ...
//
//     const handleDelete = (
//     conditionId: string
//     ) => {
//     deleteCondition.mutate(conditionId, {
//     onSuccess: ({ id, object, deleted }) => {
//     console.log(deleted)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Discount
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/discounts/{id}/conditions/{condition_id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Discounts
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDiscountConditionsDeleteRes"
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
func (m *Discount) DeleteConditon(context fiber.Ctx) error {
	config, err := api.BindConfig(context, m.r.Validator())
	if err != nil {
		return err
	}

	id, err := utils.ParseUUID(context.Params("id"))
	if err != nil {
		return err
	}

	conditionId, err := utils.ParseUUID(context.Params("condition_id"))
	if err != nil {
		return err
	}

	condition, err := m.r.DiscountConditionService().SetContext(context.Context()).Retrieve(conditionId, config)
	if err != nil {
		discount, err := m.r.DiscountService().SetContext(context.Context()).Retrieve(id, config)
		if err != nil {
			return err
		}

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"id":       conditionId,
			"object":   "disocunt-condition",
			"deleted":  true,
			"discount": discount,
		})
	}

	discount, err := m.r.DiscountService().SetContext(context.Context()).Retrieve(id, &sql.Options{Selects: []string{"id", "rule_id"}})
	if err != nil {
		return err
	}

	if condition.DiscountRuleId.UUID != discount.RuleId.UUID {
		return utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintf(`Condition with id %s does not belong to Discount with id %s`, conditionId, id),
		)
	}

	if err := m.r.DiscountConditionService().SetContext(context.Context()).Delete(conditionId); err != nil {
		return err
	}

	result, err := m.r.DiscountService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":       conditionId,
		"object":   "disocunt-condition",
		"deleted":  true,
		"discount": result,
	})
}

// @oas:path [delete] /admin/discounts/{id}/dynamic-codes/{code}
// operationId: "DeleteDiscountsDiscountDynamicCodesCode"
// summary: "Delete a Dynamic Code"
// description: "Delete a dynamic code from a Discount."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Discount
//   - (path) code=* {string} The dynamic code to delete
//
// x-codegen:
//
//	method: deleteDynamicCode
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.discounts.deleteDynamicCode(discountId, code)
//     .then(({ discount }) => {
//     console.log(discount.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteDynamicDiscountCode } from "medusa-react"
//
//     type Props = {
//     discountId: string
//     }
//
//     const Discount = ({ discountId }: Props) => {
//     const deleteDynamicDiscount = useAdminDeleteDynamicDiscountCode(discountId)
//     // ...
//
//     const handleDelete = (code: string) => {
//     deleteDynamicDiscount.mutate(code, {
//     onSuccess: ({ discount }) => {
//     console.log(discount.is_dynamic)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Discount
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/discounts/{id}/dynamic-codes/{code}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Discounts
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDiscountsRes"
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
func (m *Discount) DeleteDynamicCode(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	code := context.Params("code")

	if err := m.r.DiscountService().SetContext(context.Context()).DeleteDynamicCode(id, code); err != nil {
		return err
	}

	result, err := m.r.DiscountService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/discounts/{discount_id}/conditions/{condition_id}/batch
// operationId: "DeleteDiscountsDiscountConditionsConditionBatch"
// summary: "Remove Batch Resources"
// description: "Remove a batch of resources from a discount condition. This will only remove the association between the resource and the discount condition, not the resource itself."
// x-authenticated: true
// parameters:
//   - (path) discount_id=* {string} The ID of the discount.
//   - (path) condition_id=* {string} The ID of the condition to remove the resources from.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned discount.
//   - (query) fields {string} Comma-separated fields that should be included in the returned discount.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminDeleteDiscountsDiscountConditionsConditionBatchReq"
//
// x-codegen:
//
//	method: deleteConditionResourceBatch
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.discounts.deleteConditionResourceBatch(discountId, conditionId, {
//     resources: [{ id: itemId }]
//     })
//     .then(({ discount }) => {
//     console.log(discount.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import {
//     useAdminDeleteDiscountConditionResourceBatch
//     } from "medusa-react"
//
//     type Props = {
//     discountId: string
//     conditionId: string
//     }
//
//     const DiscountCondition = ({
//     discountId,
//     conditionId
//     }: Props) => {
//     const deleteConditionResource = useAdminDeleteDiscountConditionResourceBatch(
//     discountId,
//     conditionId,
//     )
//     // ...
//
//     const handleDelete = (itemId: string) => {
//     deleteConditionResource.mutate({
//     resources: [
//     {
//     id: itemId
//     }
//     ]
//     }, {
//     onSuccess: ({ discount }) => {
//     console.log(discount.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default DiscountCondition
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/discounts/{id}/conditions/{condition_id}/batch' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "resources": [{ "id": "item_id" }]
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Discounts
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDiscountsRes"
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
func (m *Discount) DeleteResourcesToConditionBatch(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.AddResourcesToConditionBatch](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	conditionId, err := utils.ParseUUID(context.Params("condition_id"))
	if err != nil {
		return err
	}

	condition, err := m.r.DiscountConditionService().SetContext(context.Context()).Retrieve(conditionId, &sql.Options{Selects: []string{"id", "type", "discount_rule_id"}})
	if err != nil {
		return err
	}

	input := &types.DiscountConditionInput{Id: conditionId, RuleId: condition.DiscountRuleId.UUID}
	if condition.Type == models.DiscountConditionTypeProducts {
		input.Products = model.Resources
	} else if condition.Type == models.DiscountConditionTypeProductTypes {
		input.ProductTypes = model.Resources
	} else if condition.Type == models.DiscountConditionTypeProductTags {
		input.ProductTags = model.Resources
	} else if condition.Type == models.DiscountConditionTypeProductCollections {
		input.ProductCollections = model.Resources
	} else if condition.Type == models.DiscountConditionTypeCustomerGroups {
		input.CustomerGroups = model.Resources
	}

	if err := m.r.DiscountConditionService().SetContext(context.Context()).RemoveResources(input); err != nil {
		return err
	}

	result, err := m.r.DiscountService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/discounts/{id}/regions/{region_id}
// operationId: "DeleteDiscountsDiscountRegionsRegion"
// summary: "Remove Region"
// x-authenticated: true
// description: "Remove a Region from the list of Regions that a Discount can be used in. This does not delete a region, only the association between it and the discount."
// parameters:
//   - (path) id=* {string} The ID of the Discount.
//   - (path) region_id=* {string} The ID of the Region.
//
// x-codegen:
//
//	method: removeRegion
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.discounts.removeRegion(discountId, regionId)
//     .then(({ discount }) => {
//     console.log(discount.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDiscountRemoveRegion } from "medusa-react"
//
//     type Props = {
//     discountId: string
//     }
//
//     const Discount = ({ discountId }: Props) => {
//     const deleteRegion = useAdminDiscountRemoveRegion(discountId)
//     // ...
//
//     const handleDelete = (regionId: string) => {
//     deleteRegion.mutate(regionId, {
//     onSuccess: ({ discount }) => {
//     console.log(discount.regions)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Discount
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/discounts/{id}/regions/{region_id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Discounts
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDiscountsRes"
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
func (m *Discount) RemoveRegion(context fiber.Ctx) error {
	id, err := utils.ParseUUID(context.Params("id"))
	if err != nil {
		return err
	}

	regionId, err := utils.ParseUUID(context.Params("region_id"))
	if err != nil {
		return err
	}

	result, err := m.r.DiscountService().SetContext(context.Context()).RemoveRegion(id, regionId)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

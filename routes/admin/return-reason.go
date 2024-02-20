package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type ReturnReason struct {
	r    Registry
	name string
}

func NewReturnReason(r Registry) *ReturnReason {
	m := ReturnReason{r: r, name: "return_reason"}
	return &m
}

func (m *ReturnReason) SetRoutes(router fiber.Router) {
	route := router.Group("/return-reasons")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)
}

// @oas:path [get] /admin/return-reasons/{id}
// operationId: "GetReturnReasonsReason"
// summary: "Get a Return Reason"
// description: "Retrieve a Return Reason's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Return Reason.
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
//     medusa.admin.returnReasons.retrieve(returnReasonId)
//     .then(({ return_reason }) => {
//     console.log(return_reason.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminReturnReason } from "medusa-react"
//
//     type Props = {
//     returnReasonId: string
//     }
//
//     const ReturnReason = ({ returnReasonId }: Props) => {
//     const { return_reason, isLoading } = useAdminReturnReason(
//     returnReasonId
//     )
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {return_reason && <span>{return_reason.label}</span>}
//     </div>
//     )
//     }
//
//     export default ReturnReason
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/return-reasons/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Return Reasons
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminReturnReasonsRes"
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
func (m *ReturnReason) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.ReturnReasonService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [get] /admin/return-reasons
// operationId: "GetReturnReasons"
// summary: "List Return Reasons"
// description: "Retrieve a list of Return Reasons."
// x-authenticated: true
// x-codegen:
//
//	method: list
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.returnReasons.list()
//     .then(({ return_reasons }) => {
//     console.log(return_reasons.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminReturnReasons } from "medusa-react"
//
//     const ReturnReasons = () => {
//     const { return_reasons, isLoading } = useAdminReturnReasons()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {return_reasons && !return_reasons.length && (
//     <span>No Return Reasons</span>
//     )}
//     {return_reasons && return_reasons.length > 0 && (
//     <ul>
//     {return_reasons.map((reason) => (
//     <li key={reason.id}>
//     {reason.label}: {reason.value}
//     </li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default ReturnReasons
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/return-reasons' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Return Reasons
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminReturnReasonsListRes"
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
func (m *ReturnReason) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableReturnReason](context)
	if err != nil {
		return err
	}
	result, err := m.r.ReturnReasonService().SetContext(context.Context()).List(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"return_reasons": result,
	})
}

// @oas:path [post] /admin/return-reasons
// operationId: "PostReturnReasons"
// summary: "Create a Return Reason"
// description: "Create a Return Reason."
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostReturnReasonsReq"
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
//     medusa.admin.returnReasons.create({
//     label: "Damaged",
//     value: "damaged"
//     })
//     .then(({ return_reason }) => {
//     console.log(return_reason.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreateReturnReason } from "medusa-react"
//
//     const CreateReturnReason = () => {
//     const createReturnReason = useAdminCreateReturnReason()
//     // ...
//
//     const handleCreate = (
//     label: string,
//     value: string
//     ) => {
//     createReturnReason.mutate({
//     label,
//     value,
//     }, {
//     onSuccess: ({ return_reason }) => {
//     console.log(return_reason.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default CreateReturnReason
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/return-reasons' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "label": "Damaged",
//     "value": "damaged"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Return Reasons
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminReturnReasonsRes"
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
func (m *ReturnReason) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateReturnReason](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ReturnReasonService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/return-reasons/{id}
// operationId: "PostReturnReasonsReason"
// summary: "Update a Return Reason"
// description: "Update a Return Reason's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Return Reason.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostReturnReasonsReasonReq"
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
//     medusa.admin.returnReasons.update(returnReasonId, {
//     label: "Damaged"
//     })
//     .then(({ return_reason }) => {
//     console.log(return_reason.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdateReturnReason } from "medusa-react"
//
//     type Props = {
//     returnReasonId: string
//     }
//
//     const ReturnReason = ({ returnReasonId }: Props) => {
//     const updateReturnReason = useAdminUpdateReturnReason(
//     returnReasonId
//     )
//     // ...
//
//     const handleUpdate = (
//     label: string
//     ) => {
//     updateReturnReason.mutate({
//     label,
//     }, {
//     onSuccess: ({ return_reason }) => {
//     console.log(return_reason.label)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default ReturnReason
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/return-reasons/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "label": "Damaged"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Return Reasons
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminReturnReasonsRes"
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
func (m *ReturnReason) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateReturnReason](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ReturnReasonService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/return-reasons/{id}
// operationId: "DeleteReturnReason"
// summary: "Delete a Return Reason"
// description: "Delete a return reason."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the return reason
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
//     medusa.admin.returnReasons.delete(returnReasonId)
//     .then(({ id, object, deleted }) => {
//     console.log(id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteReturnReason } from "medusa-react"
//
//     type Props = {
//     returnReasonId: string
//     }
//
//     const ReturnReason = ({ returnReasonId }: Props) => {
//     const deleteReturnReason = useAdminDeleteReturnReason(
//     returnReasonId
//     )
//     // ...
//
//     const handleDelete = () => {
//     deleteReturnReason.mutate(void 0, {
//     onSuccess: ({ id, object, deleted }) => {
//     console.log(id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default ReturnReason
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/return-reasons/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Return Reasons
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminReturnReasonsDeleteRes"
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
func (m *ReturnReason) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.ReturnReasonService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "return-reason",
		"deleted": true,
	})
}

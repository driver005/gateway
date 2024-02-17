package store

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type ReturnReason struct {
	r Registry
}

func NewReturnReason(r Registry) *ReturnReason {
	m := ReturnReason{r: r}
	return &m
}

func (m *ReturnReason) SetRoutes(router fiber.Router) {
	route := router.Group("/return-reasons")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
}

// @oas:path [get] /store/return-reasons/{id}
// operationId: "GetReturnReasonsReason"
// summary: "Get a Return Reason"
// description: "Retrieve a Return Reason's details."
// parameters:
//   - (path) id=* {string} The id of the Return Reason.
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
//     medusa.returnReasons.retrieve(reasonId)
//     .then(({ return_reason }) => {
//     console.log(return_reason.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useReturnReason } from "medusa-react"
//
//     type Props = {
//     returnReasonId: string
//     }
//
//     const ReturnReason = ({ returnReasonId }: Props) => {
//     const {
//     return_reason,
//     isLoading
//     } = useReturnReason(
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
//     curl '{backend_url}/store/return-reasons/{id}'
//
// tags:
//   - Return Reasons
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreReturnReasonsRes"
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
func (m *ReturnReason) Get(context fiber.Ctx) error {
	id, err := utils.ParseUUID(context.Params("id"))
	if err != nil {
		return err
	}

	result, err := m.r.ReturnReasonService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [get] /store/return-reasons
// operationId: "GetReturnReasons"
// summary: "List Return Reasons"
// description: "Retrieve a list of Return Reasons. This is useful when implementing a Create Return flow in the storefront."
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
//     medusa.returnReasons.list()
//     .then(({ return_reasons }) => {
//     console.log(return_reasons.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useReturnReasons } from "medusa-react"
//
//     const ReturnReasons = () => {
//     const {
//     return_reasons,
//     isLoading
//     } = useReturnReasons()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {return_reasons?.length && (
//     <ul>
//     {return_reasons.map((returnReason) => (
//     <li key={returnReason.id}>
//     {returnReason.label}
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
//     curl '{backend_url}/store/return-reasons'
//
// tags:
//   - Return Reasons
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreReturnReasonsListRes"
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
func (m *ReturnReason) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableReturnReason](context)
	if err != nil {
		return err
	}

	model.ParentReturnReasonId = uuid.Nil

	result, err := m.r.ReturnReasonService().SetContext(context.Context()).List(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

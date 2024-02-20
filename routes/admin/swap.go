package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type Swap struct {
	r    Registry
	name string
}

func NewSwap(r Registry) *Swap {
	m := Swap{r: r, name: "swap"}
	return &m
}

func (m *Swap) SetRoutes(router fiber.Router) {
	route := router.Group("/swaps")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
}

// @oas:path [get] /admin/swaps/{id}
// operationId: "GetSwapsSwap"
// summary: "Get a Swap"
// description: "Retrieve a Swap's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Swap.
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
//     medusa.admin.swaps.retrieve(swapId)
//     .then(({ swap }) => {
//     console.log(swap.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminSwap } from "medusa-react"
//
//     type Props = {
//     swapId: string
//     }
//
//     const Swap = ({ swapId }: Props) => {
//     const { swap, isLoading } = useAdminSwap(swapId)
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {swap && <span>{swap.id}</span>}
//     </div>
//     )
//     }
//
//     export default Swap
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/swaps/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Swaps
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminSwapsRes"
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
func (m *Swap) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.SwapService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [get] /admin/swaps
// operationId: "GetSwaps"
// summary: "List Swaps"
// description: "Retrieve a list of Swaps. The swaps can be paginated."
// parameters:
//   - (query) limit=50 {number} Limit the number of swaps returned.
//   - (query) offset=0 {number} The number of swaps to skip when retrieving the swaps.
//
// x-authenticated: true
// x-codegen:
//
//	method: list
//	queryParams: AdminGetSwapsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.swaps.list()
//     .then(({ swaps }) => {
//     console.log(swaps.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminSwaps } from "medusa-react"
//
//     const Swaps = () => {
//     const { swaps, isLoading } = useAdminSwaps()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {swaps && !swaps.length && <span>No Swaps</span>}
//     {swaps && swaps.length > 0 && (
//     <ul>
//     {swaps.map((swap) => (
//     <li key={swap.id}>{swap.payment_status}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default Swaps
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/swaps' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Swaps
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminSwapsListRes"
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
func (m *Swap) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableSwap](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.SwapService().SetContext(context.Context()).ListAndCount(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"swaps":  result,
		"count":  count,
		"offset": config.Skip,
		"limit":  config.Take,
	})
}

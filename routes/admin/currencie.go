package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type Currencie struct {
	r Registry
}

func NewCurrencie(r Registry) *Currencie {
	m := Currencie{r: r}
	return &m
}

func (m *Currencie) SetRoutes(router fiber.Router) {
	route := router.Group("/currencies")
	route.Get("", m.List)
	route.Post("/:id", m.Update)
}

// @oas:path [get] /admin/currencies
// operationId: "GetCurrencies"
// summary: "List Currency"
// description: "Retrieve a list of currencies. The currencies can be filtered by fields such as `code`. The currencies can also be sorted or paginated."
// x-authenticated: true
// parameters:
//   - (query) code {string} filter by currency code.
//   - in: query
//     name: includes_tax
//     description: filter currencies by whether they include taxes or not.
//     schema:
//     type: boolean
//     x-featureFlag: "tax_inclusive_pricing"
//   - (query) order {string} A field to sort order the retrieved currencies by.
//   - (query) q {string} Term used to search currencies' name and code.
//   - (query) offset=0 {number} The number of currencies to skip when retrieving the currencies.
//   - (query) limit=20 {number} The number of currencies to return.
//
// x-codegen:
//
//	method: list
//	queryParams: AdminGetCurrenciesParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.currencies.list()
//     .then(({ currencies, count, offset, limit }) => {
//     console.log(currencies.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCurrencies } from "medusa-react"
//
//     const Currencies = () => {
//     const { currencies, isLoading } = useAdminCurrencies()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {currencies && !currencies.length && (
//     <span>No Currencies</span>
//     )}
//     {currencies && currencies.length > 0 && (
//     <ul>
//     {currencies.map((currency) => (
//     <li key={currency.code}>{currency.name}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default Currencies
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/currencies' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Currencies
//
// responses:
//
//	"200":
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminCurrenciesListRes"
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
func (m *Currencie) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableCurrencyProps](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.CurrencyService().SetContext(context.Context()).ListAndCount(model, config)
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

// @oas:path [post] /admin/currencies/{code}
// operationId: "PostCurrenciesCurrency"
// summary: "Update a Currency"
// description: "Update a Currency's details."
// x-authenticated: true
// parameters:
//   - (path) code=* {string} The code of the Currency.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostCurrenciesCurrencyReq"
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
//     medusa.admin.currencies.update(code, {
//     includes_tax: true
//     })
//     .then(({ currency }) => {
//     console.log(currency.code);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdateCurrency } from "medusa-react"
//
//     type Props = {
//     currencyCode: string
//     }
//
//     const Currency = ({ currencyCode }: Props) => {
//     const updateCurrency = useAdminUpdateCurrency(currencyCode)
//     // ...
//
//     const handleUpdate = (includes_tax: boolean) => {
//     updateCurrency.mutate({
//     includes_tax,
//     }, {
//     onSuccess: ({ currency }) => {
//     console.log(currency)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Currency
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/currencies/{code}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "includes_tax": true
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Currencies
//
// responses:
//
//	"200":
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminCurrenciesRes"
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
func (m *Currencie) Update(context fiber.Ctx) error {
	model, code, err := api.BindWithString[types.UpdateCurrencyInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.CurrencyService().SetContext(context.Context()).Update(code, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

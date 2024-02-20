package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
)

type Store struct {
	r    Registry
	name string
}

func NewStore(r Registry) *Store {
	m := Store{r: r, name: "store"}
	return &m
}

func (m *Store) SetRoutes(router fiber.Router) {
	route := router.Group("/store")
	route.Get("", m.Get)
	route.Post("", m.Update)

	route.Get("/payment-providers", m.ListPaymentProviders)
	route.Get("/tax-providers", m.ListTaxProviders)
	route.Post("/currencies/:currency_code", m.AddCurrency)
	route.Delete("/currencies/:currency_code", m.RemoveCurrency)
}

// @oas:path [get] /admin/store
// operationId: "GetStore"
// summary: "Get Store details"
// description: "Retrieve the Store's details."
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
//     // must be previously logged in or use api token
//     medusa.admin.store.retrieve()
//     .then(({ store }) => {
//     console.log(store.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminStore } from "medusa-react"
//
//     const Store = () => {
//     const {
//     store,
//     isLoading
//     } = useAdminStore()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {store && <span>{store.name}</span>}
//     </div>
//     )
//     }
//
//     export default Store
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/store' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Store
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminExtendedStoresRes"
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
func (m *Store) Get(context fiber.Ctx) error {
	var config *sql.Options
	if err := context.Bind().Query(config); err != nil {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Invalid query parameters",
			nil,
		)
	}

	result, err := m.r.StoreService().SetContext(context.Context()).Retrieve(config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/store
// operationId: "PostStore"
// summary: "Update Store Details"
// description: "Update the Store's details."
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostStoreReq"
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
//     medusa.admin.store.update({
//     name: "Medusa Store"
//     })
//     .then(({ store }) => {
//     console.log(store.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdateStore } from "medusa-react"
//
//     function Store() {
//     const updateStore = useAdminUpdateStore()
//     // ...
//
//     const handleUpdate = (
//     name: string
//     ) => {
//     updateStore.mutate({
//     name
//     }, {
//     onSuccess: ({ store }) => {
//     console.log(store.name)
//     }
//     })
//     }
//     }
//
//     export default Store
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/store' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "name": "Medusa Store"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Store
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminStoresRes"
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
func (m *Store) Update(context fiber.Ctx) error {
	model, err := api.Bind[types.UpdateStoreInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.StoreService().SetContext(context.Context()).Update(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/store/currencies/{code}
// operationId: "PostStoreCurrenciesCode"
// summary: "Add a Currency Code"
// description: "Add a Currency Code to the available currencies in a store. This does not create new currencies, as currencies are defined within the Medusa backend.
// To create a currency, you can create a migration that inserts the currency into the database."
// x-authenticated: true
// parameters:
//   - in: path
//     name: code
//     required: true
//     description: The 3 character ISO currency code.
//     schema:
//     type: string
//     externalDocs:
//     url: https://en.wikipedia.org/wiki/ISO_4217#Active_codes
//     description: See a list of codes.
//
// x-codegen:
//
//	method: addCurrency
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.store.addCurrency("eur")
//     .then(({ store }) => {
//     console.log(store.currencies);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminAddStoreCurrency } from "medusa-react"
//
//     const Store = () => {
//     const addCurrency = useAdminAddStoreCurrency()
//     // ...
//
//     const handleAdd = (code: string) => {
//     addCurrency.mutate(code, {
//     onSuccess: ({ store }) => {
//     console.log(store.currencies)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Store
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/store/currencies/{currency_code}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Store
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminStoresRes"
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
func (m *Store) AddCurrency(context fiber.Ctx) error {
	currencyCode := context.Params("currency_code")

	result, err := m.r.StoreService().SetContext(context.Context()).AddCurrency(currencyCode)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/store/currencies/{code}
// operationId: "DeleteStoreCurrenciesCode"
// summary: "Remove a Currency"
// description: "Remove a Currency Code from the available currencies in a store. This does not completely delete the currency and it can be added again later to the store."
// x-authenticated: true
// parameters:
//   - in: path
//     name: code
//     required: true
//     description: The 3 character ISO currency code.
//     schema:
//     type: string
//     externalDocs:
//     url: https://en.wikipedia.org/wiki/ISO_4217#Active_codes
//     description: See a list of codes.
//
// x-codegen:
//
//	method: deleteCurrency
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.store.deleteCurrency("eur")
//     .then(({ store }) => {
//     console.log(store.currencies);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteStoreCurrency } from "medusa-react"
//
//     const Store = () => {
//     const deleteCurrency = useAdminDeleteStoreCurrency()
//     // ...
//
//     const handleAdd = (code: string) => {
//     deleteCurrency.mutate(code, {
//     onSuccess: ({ store }) => {
//     console.log(store.currencies)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Store
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/store/currencies/{currency_code}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Store
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminStoresRes"
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
func (m *Store) RemoveCurrency(context fiber.Ctx) error {
	currencyCode := context.Params("currency_code")

	result, err := m.r.StoreService().SetContext(context.Context()).RemoveCurrency(currencyCode)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [get] /admin/store/payment-providers
// operationId: "GetStorePaymentProviders"
// summary: "List Payment Providers"
// description: "Retrieve a list of available Payment Providers in a store."
// x-authenticated: true
// x-codegen:
//
//	method: listPaymentProviders
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.store.listPaymentProviders()
//     .then(({ payment_providers }) => {
//     console.log(payment_providers.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminStorePaymentProviders } from "medusa-react"
//
//     const PaymentProviders = () => {
//     const {
//     payment_providers,
//     isLoading
//     } = useAdminStorePaymentProviders()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {payment_providers && !payment_providers.length && (
//     <span>No Payment Providers</span>
//     )}
//     {payment_providers &&
//     payment_providers.length > 0 &&(
//     <ul>
//     {payment_providers.map((provider) => (
//     <li key={provider.id}>{provider.id}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default PaymentProviders
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/store/payment-providers' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Store
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPaymentProvidersList"
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
func (m *Store) ListPaymentProviders(context fiber.Ctx) error {
	result, err := m.r.PaymentProviderService().SetContext(context.Context()).List()
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"payment_providers": result,
	})
}

// @oas:path [get] /admin/store/tax-providers
// operationId: "GetStoreTaxProviders"
// summary: "List Tax Providers"
// description: "Retrieve a list of available Tax Providers in a store."
// x-authenticated: true
// x-codegen:
//
//	method: listTaxProviders
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.store.listTaxProviders()
//     .then(({ tax_providers }) => {
//     console.log(tax_providers.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminStoreTaxProviders } from "medusa-react"
//
//     const TaxProviders = () => {
//     const {
//     tax_providers,
//     isLoading
//     } = useAdminStoreTaxProviders()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {tax_providers && !tax_providers.length && (
//     <span>No Tax Providers</span>
//     )}
//     {tax_providers &&
//     tax_providers.length > 0 &&(
//     <ul>
//     {tax_providers.map((provider) => (
//     <li key={provider.id}>{provider.id}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default TaxProviders
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/store/tax-providers' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Store
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminTaxProvidersList"
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
func (m *Store) ListTaxProviders(context fiber.Ctx) error {
	result, err := m.r.TaxProviderService().SetContext(context.Context()).List(&models.TaxProvider{}, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"tax_providers": result,
	})
}

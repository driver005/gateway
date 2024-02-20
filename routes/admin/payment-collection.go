package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type PaymentCollection struct {
	r    Registry
	name string
}

func NewPaymentCollection(r Registry) *PaymentCollection {
	m := PaymentCollection{r: r, name: "payment_collection"}
	return &m
}

func (m *PaymentCollection) SetRoutes(router fiber.Router) {
	route := router.Group("/payment-collections")
	route.Get("/:id", m.Get)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)

	route.Post("/:id/authorize", m.MarkAuthorized)
}

// @oas:path [get] /admin/payment-collections/{id}
// operationId: "GetPaymentCollectionsPaymentCollection"
// summary: "Get a Payment Collection"
// description: "Retrieve a Payment Collection's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Payment Collection.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned payment collection.
//   - (query) fields {string} Comma-separated fields that should be included in the returned payment collection.
//
// x-codegen:
//
//	method: retrieve
//	queryParams: AdminGetPaymentCollectionsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.paymentCollections.retrieve(paymentCollectionId)
//     .then(({ payment_collection }) => {
//     console.log(payment_collection.id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminPaymentCollection } from "medusa-react"
//
//     type Props = {
//     paymentCollectionId: string
//     }
//
//     const PaymentCollection = ({ paymentCollectionId }: Props) => {
//     const {
//     payment_collection,
//     isLoading,
//     } = useAdminPaymentCollection(paymentCollectionId)
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {payment_collection && (
//     <span>{payment_collection.status}</span>
//     )}
//
//     </div>
//     )
//     }
//
//     export default PaymentCollection
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/payment-collections/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Payment Collections
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPaymentCollectionsRes"
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
func (m *PaymentCollection) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.PaymentCollectionService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/payment-collections/{id}
// operationId: "PostPaymentCollectionsPaymentCollection"
// summary: "Update Payment Collection"
// description: "Update a Payment Collection's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Payment Collection.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminUpdatePaymentCollectionsReq"
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
//     medusa.admin.paymentCollections.update(paymentCollectionId, {
//     description
//     })
//     .then(({ payment_collection }) => {
//     console.log(payment_collection.id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdatePaymentCollection } from "medusa-react"
//
//     type Props = {
//     paymentCollectionId: string
//     }
//
//     const PaymentCollection = ({ paymentCollectionId }: Props) => {
//     const updateCollection = useAdminUpdatePaymentCollection(
//     paymentCollectionId
//     )
//     // ...
//
//     const handleUpdate = (
//     description: string
//     ) => {
//     updateCollection.mutate({
//     description
//     }, {
//     onSuccess: ({ payment_collection }) => {
//     console.log(payment_collection.description)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default PaymentCollection
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/payment-collections/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "description": "Description of payment collection"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Payment Collections
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPaymentCollectionsRes"
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
func (m *PaymentCollection) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdatePaymentCollectionInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	data := &models.PaymentCollection{}
	data.Description = model.Description
	data.Metadata = model.Metadata

	result, err := m.r.PaymentCollectionService().SetContext(context.Context()).Update(id, data)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/payment-collections/{id}
// operationId: "DeletePaymentCollectionsPaymentCollection"
// summary: "Delete a Payment Collection"
// description: "Delete a Payment Collection. Only payment collections with the statuses `canceled` or `not_paid` can be deleted."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Payment Collection.
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
//     medusa.admin.paymentCollections.delete(paymentCollectionId)
//     .then(({ id, object, deleted }) => {
//     console.log(id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeletePaymentCollection } from "medusa-react"
//
//     type Props = {
//     paymentCollectionId: string
//     }
//
//     const PaymentCollection = ({ paymentCollectionId }: Props) => {
//     const deleteCollection = useAdminDeletePaymentCollection(
//     paymentCollectionId
//     )
//     // ...
//
//     const handleDelete = () => {
//     deleteCollection.mutate(void 0, {
//     onSuccess: ({ id, object, deleted }) => {
//     console.log(id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default PaymentCollection
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/payment-collections/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Payment Collections
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPaymentCollectionDeleteRes"
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
func (m *PaymentCollection) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.PaymentCollectionService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "payment-collection",
		"deleted": true,
	})
}

// @oas:path [post] /admin/payment-collections/{id}/authorize
// operationId: "PostPaymentCollectionsPaymentCollectionAuthorize"
// summary: "Mark Authorized"
// description: "Set the status of a Payment Collection as `authorized`. This will also change the `authorized_amount` of the payment collection."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Payment Collection.
//
// x-codegen:
//
//	method: markAsAuthorized
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.paymentCollections.markAsAuthorized(paymentCollectionId)
//     .then(({ payment_collection }) => {
//     console.log(payment_collection.id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminMarkPaymentCollectionAsAuthorized } from "medusa-react"
//
//     type Props = {
//     paymentCollectionId: string
//     }
//
//     const PaymentCollection = ({ paymentCollectionId }: Props) => {
//     const markAsAuthorized = useAdminMarkPaymentCollectionAsAuthorized(
//     paymentCollectionId
//     )
//     // ...
//
//     const handleAuthorization = () => {
//     markAsAuthorized.mutate(void 0, {
//     onSuccess: ({ payment_collection }) => {
//     console.log(payment_collection.status)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default PaymentCollection
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/payment-collections/{id}/authorize' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Payment Collections
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPaymentCollectionsRes"
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
func (m *PaymentCollection) MarkAuthorized(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	result, err := m.r.PaymentCollectionService().SetContext(context.Context()).MarkAsAuthorized(id)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

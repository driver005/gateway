package store

import (
	"fmt"

	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type PaymentCollection struct {
	r Registry
}

func NewPaymentCollection(r Registry) *PaymentCollection {
	m := PaymentCollection{r: r}
	return &m
}

func (m *PaymentCollection) SetRoutes(router fiber.Router) {
	route := router.Group("/payment-collections")
	route.Get("/:id", m.Get)

	route.Post("/:id/sessions/batch", m.PaymentSessionManageBatch)
	route.Post("/:id/sessions/batch/authorize", m.PaymentSessionAuthorizeBatch)
	route.Post("/:id/sessions", m.PaymentSessionManage)
	route.Post("/:id/sessions/:session_id", m.PaymentSessionRefresh)
	route.Post("/:id/sessions/:session_id/authorize", m.PaymentSessionAuthorize)
}

// @oas:path [get] /store/payment-collections/{id}
// operationId: "GetPaymentCollectionsPaymentCollection"
// summary: "Get a PaymentCollection"
// description: "Retrieve a Payment Collection's details."
// x-authenticated: false
// parameters:
//   - (path) id=* {string} The ID of the PaymentCollection.
//   - (query) fields {string} Comma-separated fields that should be expanded in the returned payment collection.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned payment collection.
//
// x-codegen:
//
//	method: retrieve
//	queryParams: StoreGetPaymentCollectionsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.paymentCollections.retrieve(paymentCollectionId)
//     .then(({ payment_collection }) => {
//     console.log(payment_collection.id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { usePaymentCollection } from "medusa-react"
//
//     type Props = {
//     paymentCollectionId: string
//     }
//
//     const PaymentCollection = ({
//     paymentCollectionId
//     }: Props) => {
//     const {
//     payment_collection,
//     isLoading
//     } = usePaymentCollection(
//     paymentCollectionId
//     )
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {payment_collection && (
//     <span>{payment_collection.status}</span>
//     )}
//     </div>
//     )
//     }
//
//     export default PaymentCollection
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '{backend_url}/store/payment-collections/{id}'
//
// security:
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Payment Collections
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StorePaymentCollectionsRes"
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
func (m *PaymentCollection) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	result, err := m.r.PaymentCollectionService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /store/payment-collections/{id}/sessions/batch/authorize
// operationId: "PostPaymentCollectionsSessionsBatchAuthorize"
// summary: "Authorize Payment Sessions"
// description: "Authorize the Payment Sessions of a Payment Collection."
// x-authenticated: false
// parameters:
//   - (path) id=* {string} The ID of the Payment Collections.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/StorePostPaymentCollectionsBatchSessionsAuthorizeReq"
//
// x-codegen:
//
//	method: authorizePaymentSessionsBatch
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.paymentCollections.authorize(paymentId)
//     .then(({ payment_collection }) => {
//     console.log(payment_collection.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAuthorizePaymentSessionsBatch } from "medusa-react"
//
//     type Props = {
//     paymentCollectionId: string
//     }
//
//     const PaymentCollection = ({
//     paymentCollectionId
//     }: Props) => {
//     const authorizePaymentSessions = useAuthorizePaymentSessionsBatch(
//     paymentCollectionId
//     )
//     // ...
//
//     const handleAuthorizePayments = (paymentSessionIds: string[]) => {
//     authorizePaymentSessions.mutate({
//     session_ids: paymentSessionIds
//     }, {
//     onSuccess: ({ payment_collection }) => {
//     console.log(payment_collection.payment_sessions)
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
//     curl -X POST '{backend_url}/store/payment-collections/{id}/sessions/batch/authorize'
//
// security:
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Payment Collections
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StorePaymentCollectionsRes"
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
func (m *PaymentCollection) PaymentSessionAuthorizeBatch(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.PaymentCollectionsAuthorizeBatch](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.PaymentCollectionService().SetContext(context.Context()).AuthorizePaymentSessions(id, model.SessionIds, context.Locals("request_context").(map[string]interface{}))
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /store/payment-collections/{id}/sessions/{session_id}/authorize
// operationId: "PostPaymentCollectionsSessionsSessionAuthorize"
// summary: "Authorize Payment Session"
// description: "Authorize a Payment Session of a Payment Collection."
// x-authenticated: false
// parameters:
//   - (path) id=* {string} The ID of the Payment Collection.
//   - (path) session_id=* {string} The ID of the Payment Session.
//
// x-codegen:
//
//	method: authorizePaymentSession
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.paymentCollections.authorize(paymentId, sessionId)
//     .then(({ payment_collection }) => {
//     console.log(payment_collection.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAuthorizePaymentSession } from "medusa-react"
//
//     type Props = {
//     paymentCollectionId: string
//     }
//
//     const PaymentCollection = ({
//     paymentCollectionId
//     }: Props) => {
//     const authorizePaymentSession = useAuthorizePaymentSession(
//     paymentCollectionId
//     )
//     // ...
//
//     const handleAuthorizePayment = (paymentSessionId: string) => {
//     authorizePaymentSession.mutate(paymentSessionId, {
//     onSuccess: ({ payment_collection }) => {
//     console.log(payment_collection.payment_sessions)
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
//     curl -X POST '{backend_url}/store/payment-collections/{id}/sessions/{session_id}/authorize'
//
// security:
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Payment Collections
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StorePaymentCollectionsSessionRes"
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
func (m *PaymentCollection) PaymentSessionAuthorize(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	sessionId, err := api.BindDelete(context, "session_id")
	if err != nil {
		return err
	}

	paymentCollection, err := m.r.PaymentCollectionService().SetContext(context.Context()).AuthorizePaymentSessions(id, uuid.UUIDs{sessionId}, context.Locals("request_context").(map[string]interface{}))
	if err != nil {
		return err
	}

	result, ok := lo.Find(paymentCollection.PaymentSessions, func(item models.PaymentSession) bool {
		return item.Id == sessionId
	})

	if !ok {
		return utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintf("Could not find Payment Session with id %s", id),
		)
	}

	if result.Status != models.PaymentSessionStatusAuthorized {
		return utils.NewApplictaionError(
			utils.PAYMENT_AUTHORIZATION_ERROR,
			fmt.Sprintf("Failed to authorize Payment Session id %s", id),
		)
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /store/payment-collections/{id}/sessions/batch
// operationId: "PostPaymentCollectionsPaymentCollectionSessionsBatch"
// summary: "Manage Payment Sessions"
// description: "Create, update, or delete a list of payment sessions of a Payment Collections. If a payment session is not provided in the `sessions` array, it's deleted."
// x-authenticated: false
// parameters:
//   - (path) id=* {string} The ID of the Payment Collection.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/StorePostPaymentCollectionsBatchSessionsReq"
//
// x-codegen:
//
//	method: managePaymentSessionsBatch
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//
//     // Total amount = 10000
//
//     // Example 1: Adding two new sessions
//     medusa.paymentCollections.managePaymentSessionsBatch(paymentId, {
//     sessions: [
//     {
//     provider_id: "stripe",
//     amount: 5000,
//     },
//     {
//     provider_id: "manual",
//     amount: 5000,
//     },
//     ]
//     })
//     .then(({ payment_collection }) => {
//     console.log(payment_collection.id);
//     })
//
//     // Example 2: Updating one session and removing the other
//     medusa.paymentCollections.managePaymentSessionsBatch(paymentId, {
//     sessions: [
//     {
//     provider_id: "stripe",
//     amount: 10000,
//     session_id: "ps_123456"
//     },
//     ]
//     })
//     .then(({ payment_collection }) => {
//     console.log(payment_collection.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useManageMultiplePaymentSessions } from "medusa-react"
//
//     type Props = {
//     paymentCollectionId: string
//     }
//
//     const PaymentCollection = ({
//     paymentCollectionId
//     }: Props) => {
//     const managePaymentSessions = useManageMultiplePaymentSessions(
//     paymentCollectionId
//     )
//
//     const handleManagePaymentSessions = () => {
//     // Total amount = 10000
//
//     // Example 1: Adding two new sessions
//     managePaymentSessions.mutate({
//     sessions: [
//     {
//     provider_id: "stripe",
//     amount: 5000,
//     },
//     {
//     provider_id: "manual",
//     amount: 5000,
//     },
//     ]
//     }, {
//     onSuccess: ({ payment_collection }) => {
//     console.log(payment_collection.payment_sessions)
//     }
//     })
//
//     // Example 2: Updating one session and removing the other
//     managePaymentSessions.mutate({
//     sessions: [
//     {
//     provider_id: "stripe",
//     amount: 10000,
//     session_id: "ps_123456"
//     },
//     ]
//     }, {
//     onSuccess: ({ payment_collection }) => {
//     console.log(payment_collection.payment_sessions)
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
//     curl -X POST '{backend_url}/store/payment-collections/{id}/sessions/batch' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "sessions": [
//     {
//     "provider_id": "stripe",
//     "amount": 5000
//     },
//     {
//     "provider_id": "manual",
//     "amount": 5000
//     }
//     ]
//     }'
//
// security:
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Payment Collections
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StorePaymentCollectionsRes"
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
func (m *PaymentCollection) PaymentSessionManageBatch(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.PaymentCollectionsSessionsBatch](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	customerId := context.Locals("customer_id").(uuid.UUID)

	result, err := m.r.PaymentCollectionService().SetContext(context.Context()).SetPaymentSessionsBatch(id, nil, model.Sessions, customerId)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /store/payment-collections/{id}/sessions
// operationId: "PostPaymentCollectionsSessions"
// summary: "Create a Payment Session"
// description: "Create a Payment Session for a payment provider in a Payment Collection."
// x-authenticated: false
// parameters:
//   - (path) id=* {string} The ID of the Payment Collection.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/StorePostCartsCartPaymentSessionReq"
//
// x-codegen:
//
//	method: managePaymentSession
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.paymentCollections.managePaymentSession(payment_id, { provider_id: "stripe" })
//     .then(({ payment_collection }) => {
//     console.log(payment_collection.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useManagePaymentSession } from "medusa-react"
//
//     type Props = {
//     paymentCollectionId: string
//     }
//
//     const PaymentCollection = ({
//     paymentCollectionId
//     }: Props) => {
//     const managePaymentSession = useManagePaymentSession(
//     paymentCollectionId
//     )
//
//     const handleManagePaymentSession = (
//     providerId: string
//     ) => {
//     managePaymentSession.mutate({
//     provider_id: providerId
//     }, {
//     onSuccess: ({ payment_collection }) => {
//     console.log(payment_collection.payment_sessions)
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
//     curl -X POST '{backend_url}/store/payment-collections/{id}/sessions' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "provider_id": "stripe"
//     }'
//
// security:
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Payment Collections
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StorePaymentCollectionsRes"
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
func (m *PaymentCollection) PaymentSessionManage(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.SessionsInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	customerId := context.Locals("customer_id").(uuid.UUID)

	result, err := m.r.PaymentCollectionService().SetContext(context.Context()).SetPaymentSession(id, model, customerId)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /store/payment-collections/{id}/sessions/{session_id}
// operationId: PostPaymentCollectionsPaymentCollectionPaymentSessionsSession
// summary: "Refresh a Payment Session"
// description: "Refresh a Payment Session's data to ensure that it is in sync with the Payment Collection."
// x-authenticated: false
// parameters:
//   - (path) id=* {string} The id of the PaymentCollection.
//   - (path) session_id=* {string} The id of the Payment Session to be refreshed.
//
// x-codegen:
//
//	method: refreshPaymentSession
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.paymentCollections.refreshPaymentSession(paymentCollectionId, sessionId)
//     .then(({ payment_session }) => {
//     console.log(payment_session.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { usePaymentCollectionRefreshPaymentSession } from "medusa-react"
//
//     type Props = {
//     paymentCollectionId: string
//     }
//
//     const PaymentCollection = ({
//     paymentCollectionId
//     }: Props) => {
//     const refreshPaymentSession = usePaymentCollectionRefreshPaymentSession(
//     paymentCollectionId
//     )
//     // ...
//
//     const handleRefreshPaymentSession = (paymentSessionId: string) => {
//     refreshPaymentSession.mutate(paymentSessionId, {
//     onSuccess: ({ payment_session }) => {
//     console.log(payment_session.status)
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
//     curl -X POST '{backend_url}/store/payment-collections/{id}/sessions/{session_id}'
//
// security:
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Payment Collections
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StorePaymentCollectionsSessionRes"
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
func (m *PaymentCollection) PaymentSessionRefresh(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	sessionId, err := api.BindDelete(context, "session_id")
	if err != nil {
		return err
	}

	customerId := context.Locals("customer_id").(uuid.UUID)

	result, err := m.r.PaymentCollectionService().SetContext(context.Context()).RefreshPaymentSession(id, sessionId, customerId)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

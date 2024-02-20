package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type Payment struct {
	r    Registry
	name string
}

func NewPayment(r Registry) *Payment {
	m := Payment{r: r, name: "payment"}
	return &m
}

func (m *Payment) SetRoutes(router fiber.Router) {
	route := router.Group("/payments")
	route.Get("/:id", m.Get)

	route.Post("/:id/capture", m.Capture)
	route.Post("/:id/refund", m.Refund)
}

// @oas:path [get] /admin/payments/{id}
// operationId: "GetPaymentsPayment"
// summary: "Get Payment details"
// description: "Retrieve a Payment's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Payment.
//
// x-codegen:
//
//	method: retrieve
//	queryParams: GetPaymentsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.payments.retrieve(paymentId)
//     .then(({ payment }) => {
//     console.log(payment.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminPayment } from "medusa-react"
//
//     type Props = {
//     paymentId: string
//     }
//
//     const Payment = ({ paymentId }: Props) => {
//     const {
//     payment,
//     isLoading,
//     } = useAdminPayment(paymentId)
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {payment && <span>{payment.amount}</span>}
//
//     </div>
//     )
//     }
//
//     export default Payment
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/payments/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Payments
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPaymentRes"
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
func (m *Payment) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.PaymentService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/payments/{id}/capture
// operationId: "PostPaymentsPaymentCapture"
// summary: "Capture a Payment"
// description: "Capture a Payment."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Payment.
//
// x-codegen:
//
//	method: capturePayment
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.payments.capturePayment(paymentId)
//     .then(({ payment }) => {
//     console.log(payment.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminPaymentsCapturePayment } from "medusa-react"
//
//     type Props = {
//     paymentId: string
//     }
//
//     const Payment = ({ paymentId }: Props) => {
//     const capture = useAdminPaymentsCapturePayment(
//     paymentId
//     )
//     // ...
//
//     const handleCapture = () => {
//     capture.mutate(void 0, {
//     onSuccess: ({ payment }) => {
//     console.log(payment.amount)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Payment
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/payments/{id}/capture' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Payments
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPaymentRes"
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
func (m *Payment) Capture(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	result, err := m.r.PaymentService().SetContext(context.Context()).Capture(id, nil)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/payments/{id}/refund
// operationId: "PostPaymentsPaymentRefunds"
// summary: "Refund Payment"
// description: "Refund a payment. The payment must be captured first."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Payment.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostPaymentRefundsReq"
//
// x-codegen:
//
//	method: refundPayment
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.payments.refundPayment(paymentId, {
//     amount: 1000,
//     reason: "return",
//     note: "Do not like it",
//     })
//     .then(({ payment }) => {
//     console.log(payment.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { RefundReason } from "@medusajs/medusa"
//     import { useAdminPaymentsRefundPayment } from "medusa-react"
//
//     type Props = {
//     paymentId: string
//     }
//
//     const Payment = ({ paymentId }: Props) => {
//     const refund = useAdminPaymentsRefundPayment(
//     paymentId
//     )
//     // ...
//
//     const handleRefund = (
//     amount: number,
//     reason: RefundReason,
//     note: string
//     ) => {
//     refund.mutate({
//     amount,
//     reason,
//     note
//     }, {
//     onSuccess: ({ refund }) => {
//     console.log(refund.amount)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Payment
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/payments/pay_123/refund' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "amount": 1000,
//     "reason": "return",
//     "note": "Do not like it"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Payments
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminRefundRes"
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
func (m *Payment) Refund(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.PaymentRefund](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.PaymentService().SetContext(context.Context()).Refund(id, nil, model.Amount, model.Reason, &model.Note)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

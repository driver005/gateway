package store

import (
	"fmt"

	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/samber/lo"
)

type OrderEdit struct {
	r Registry
}

func NewOrderEdit(r Registry) *OrderEdit {
	m := OrderEdit{r: r}
	return &m
}

func (m *OrderEdit) SetRoutes(router fiber.Router) {
	route := router.Group("/order-edits")
	route.Get("/:id", m.Get)

	route.Post("/:id/decline", m.Decline)
	route.Post("/:id/complete", m.Complete)
}

// @oas:path [get] /store/order-edits/{id}
// operationId: "GetOrderEditsOrderEdit"
// summary: "Retrieve an Order Edit"
// description: "Retrieve an Order Edit's details."
// parameters:
//   - (path) id=* {string} The ID of the OrderEdit.
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
//     medusa.orderEdits.retrieve(orderEditId)
//     .then(({ order_edit }) => {
//     console.log(order_edit.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useOrderEdit } from "medusa-react"
//
//     type Props = {
//     orderEditId: string
//     }
//
//     const OrderEdit = ({ orderEditId }: Props) => {
//     const { order_edit, isLoading } = useOrderEdit(orderEditId)
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {order_edit && (
//     <ul>
//     {order_edit.changes.map((change) => (
//     <li key={change.id}>{change.type}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default OrderEdit
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '{backend_url}/store/order-edits/{id}'
//
// tags:
//   - Order Edits
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreOrderEditsRes"
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
func (m *OrderEdit) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	result, err := m.r.OrderEditService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	result, err = m.r.OrderEditService().SetContext(context.Context()).DecorateTotals(result)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /store/order-edits/{id}/complete
// operationId: "PostOrderEditsOrderEditComplete"
// summary: "Complete an Order Edit"
// description: "Complete an Order Edit and reflect its changes on the original order. Any additional payment required must be authorized first using the Payment Collection API Routes."
// externalDocs:
//
//	description: "How to handle order edits in a storefront"
//	url: "https://docs.medusajs.com/modules/orders/storefront/handle-order-edits"
//
// parameters:
//   - (path) id=* {string} The ID of the Order Edit.
//
// x-codegen:
//
//	method: complete
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.orderEdits.complete(orderEditId)
//     .then(({ order_edit }) => {
//     console.log(order_edit.id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useCompleteOrderEdit } from "medusa-react"
//
//     type Props = {
//     orderEditId: string
//     }
//
//     const OrderEdit = ({ orderEditId }: Props) => {
//     const completeOrderEdit = useCompleteOrderEdit(
//     orderEditId
//     )
//     // ...
//
//     const handleCompleteOrderEdit = () => {
//     completeOrderEdit.mutate(void 0, {
//     onSuccess: ({ order_edit }) => {
//     console.log(order_edit.confirmed_at)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default OrderEdit
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/store/order-edits/{id}/complete'
//
// tags:
//   - Order Edits
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreOrderEditsRes"
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
//	"500":
//	  description: "Internal Server"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/responses/500_error"
func (m *OrderEdit) Complete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	userId := api.GetUserStore(context)

	orderEdit, err := m.r.OrderEditService().SetContext(context.Context()).Retrieve(id, &sql.Options{Relations: []string{"payment_collection", "payment_collection.payments"}})
	if err != nil {
		return err
	}

	allowedStatus := []models.OrderEditStatus{models.OrderEditStatusConfirmed, models.OrderEditStatusRequested}
	if orderEdit.PaymentCollection != nil && lo.Contains(allowedStatus, orderEdit.Status) {
		if orderEdit.PaymentCollection.Status != models.PaymentCollectionStatusAuthorized {
			return utils.NewApplictaionError(
				utils.NOT_ALLOWED,
				"Unable to complete an order edit if the payment is not authorized",
			)
		}

		if orderEdit.PaymentCollection != nil {
			for _, payment := range orderEdit.PaymentCollection.Payments {
				if payment.OrderId != orderEdit.OrderId {
					if _, err := m.r.PaymentProviderService().SetContext(context.Context()).UpdatePayment(payment.Id, &types.UpdatePaymentInput{
						OrderId: orderEdit.OrderId.UUID,
					}); err != nil {
						return err
					}

				}
			}
		}

		if orderEdit.Status != models.OrderEditStatusConfirmed {
			if orderEdit.Status != models.OrderEditStatusRequested {
				return utils.NewApplictaionError(
					utils.NOT_ALLOWED,
					fmt.Sprintf("Cannot complete an order edit with status %s", orderEdit.Status),
				)
			}

			if _, err := m.r.OrderEditService().SetContext(context.Context()).Confirm(id, userId); err != nil {
				return err
			}
		}
	}

	result, err := m.r.OrderEditService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err = m.r.OrderEditService().SetContext(context.Context()).DecorateTotals(result)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /store/order-edits/{id}/decline
// operationId: "PostOrderEditsOrderEditDecline"
// summary: "Decline an Order Edit"
// description: "Decline an Order Edit. The changes are not reflected on the original order."
// parameters:
//   - (path) id=* {string} The ID of the OrderEdit.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/StorePostOrderEditsOrderEditDecline"
//
// x-codegen:
//
//	method: decline
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.orderEdits.decline(orderEditId)
//     .then(({ order_edit }) => {
//     console.log(order_edit.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useDeclineOrderEdit } from "medusa-react"
//
//     type Props = {
//     orderEditId: string
//     }
//
//     const OrderEdit = ({ orderEditId }: Props) => {
//     const declineOrderEdit = useDeclineOrderEdit(orderEditId)
//     // ...
//
//     const handleDeclineOrderEdit = (
//     declinedReason: string
//     ) => {
//     declineOrderEdit.mutate({
//     declined_reason: declinedReason,
//     }, {
//     onSuccess: ({ order_edit }) => {
//     console.log(order_edit.declined_at)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default OrderEdit
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/store/order-edits/{id}/decline'
//
// tags:
//   - Order Edits
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreOrderEditsRes"
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
//	"500":
//	  description: "Internal Server"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/responses/500_error"
func (m *OrderEdit) Decline(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.OrderEditsDecline](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	userId := api.GetUserStore(context)

	if _, err := m.r.OrderEditService().SetContext(context.Context()).Decline(id, userId, model.DeclinedReason); err != nil {
		return err
	}

	result, err := m.r.OrderEditService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err = m.r.OrderEditService().SetContext(context.Context()).DecorateTotals(result)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

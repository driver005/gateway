package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
)

type Invite struct {
	r    Registry
	name string
}

func NewInvite(r Registry) *Invite {
	m := Invite{r: r, name: "invites"}
	return &m
}

func (m *Invite) UnauthenticatedInviteRoutes(router fiber.Router) {
	route := router.Group("/invites")
	route.Post("/accept", m.Accept)
}

func (m *Invite) SetRoutes(router fiber.Router) {
	route := router.Group("/invites")
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Delete("/:id", m.Delete)

	route.Post("/:id/resend", m.Resend)
}

// @oas:path [get] /admin/invites
// operationId: "GetInvites"
// summary: "Lists Invites"
// description: "Retrieve a list of invites."
// x-authenticated: true
// x-codegen:
//
//	method: list
//
// x-codeSamples:
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.invites.list()
//     .then(({ invites }) => {
//     console.log(invites.length);
//     })
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/invites' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Invites
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminListInvitesRes"
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
func (m *Invite) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableInvite](context)
	if err != nil {
		return err
	}
	result, err := m.r.InviteService().SetContext(context.Context()).List(model, config)
	if err != nil {
		return err
	}

	// return context.Status(fiber.StatusOK).JSON(fiber.Map{
	// 	"data":   result,
	// 	"count":  count,
	// 	"offset": config.Skip,
	// 	"limit":  config.Take,
	// })
	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/invites
// operationId: "PostInvites"
// summary: "Create an Invite"
// description: "Create an Invite. This will generate a token associated with the invite and trigger an `invite.created` event. If you have a Notification Provider installed that handles this
//
//	event, a notification should be sent to the email associated with the invite to allow them to accept the invite."
//
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostInvitesReq"
//
// x-codegen:
//
//	method: create
//
// x-codeSamples:
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.invites.create({
//     user: "user@example.com",
//     role: "admin"
//     })
//     .then(() => {
//     // successful
//     })
//     .catch(() => {
//     // an error occurred
//     })
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/invites' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "user": "user@example.com",
//     "role": "admin"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Invites
//
// responses:
//
//	200:
//	  description: "OK"
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
func (m *Invite) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateInviteInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	err = m.r.InviteService().SetContext(context.Context()).Create(model, -1)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).Send(nil)
}

// @oas:path [delete] /admin/invites/{invite_id}
// operationId: "DeleteInvitesInvite"
// summary: "Delete an Invite"
// description: "Delete an Invite. Only invites that weren't accepted can be deleted."
// x-authenticated: true
// parameters:
//   - (path) invite_id=* {string} The ID of the Invite
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
//     medusa.admin.invites.delete(inviteId)
//     .then(({ id, object, deleted }) => {
//     console.log(id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteInvite } from "medusa-react"
//
//     type Props = {
//     inviteId: string
//     }
//
//     const DeleteInvite = ({ inviteId }: Props) => {
//     const deleteInvite = useAdminDeleteInvite(inviteId)
//     // ...
//
//     const handleDelete = () => {
//     deleteInvite.mutate(void 0, {
//     onSuccess: ({ id, object, deleted }) => {
//     console.log(id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Invite
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/invites/{invite_id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Invites
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminInviteDeleteRes"
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
func (m *Invite) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.InviteService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "invite",
		"deleted": true,
	})
}

// @oas:path [post] /admin/invites/accept
// operationId: "PostInvitesInviteAccept"
// summary: "Accept an Invite"
// description: "Accept an Invite. This will also delete the invite and create a new user that can log in and perform admin functionalities. The user will have the email associated with the invite, and the password
//
//	provided in the request body."
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostInvitesInviteAcceptReq"
//
// x-codegen:
//
//	method: accept
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.invites.accept({
//     token,
//     user: {
//     first_name: "Brigitte",
//     last_name: "Collier",
//     password: "supersecret"
//     }
//     })
//     .then(() => {
//     // successful
//     })
//     .catch(() => {
//     // an error occurred
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminAcceptInvite } from "medusa-react"
//
//     const AcceptInvite = () => {
//     const acceptInvite = useAdminAcceptInvite()
//     // ...
//
//     const handleAccept = (
//     token: string,
//     firstName: string,
//     lastName: string,
//     password: string
//     ) => {
//     acceptInvite.mutate({
//     token,
//     user: {
//     first_name: firstName,
//     last_name: lastName,
//     password,
//     },
//     }, {
//     onSuccess: () => {
//     // invite accepted successfully.
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default AcceptInvite
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/invites/accept' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "token": "{token}",
//     "user": {
//     "first_name": "Brigitte",
//     "last_name": "Collier",
//     "password": "supersecret"
//     }
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Invites
//
// responses:
//
//	200:
//	  description: "OK"
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
func (m *Invite) Accept(context fiber.Ctx) error {
	model, err := api.BindCreate[types.AcceptInvite](context, m.r.Validator())
	if err != nil {
		return err
	}

	_, err = m.r.InviteService().SetContext(context.Context()).Accept(model.Token, model.User)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).Send(nil)
}

// @oas:path [post] /admin/invites/{invite_id}/resend
// operationId: "PostInvitesInviteResend"
// summary: "Resend an Invite"
// description: "Resend an Invite. This renews the expiry date by 7 days and generates a new token for the invite. It also triggers the `invite.created` event, so if you have a Notification Provider installed that handles this
//
//	event, a notification should be sent to the email associated with the invite to allow them to accept the invite."
//
// x-authenticated: true
// parameters:
//   - (path) invite_id=* {string} The ID of the Invite
//
// x-codegen:
//
//	method: resend
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.invites.resend(inviteId)
//     .then(() => {
//     // successful
//     })
//     .catch(() => {
//     // an error occurred
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminResendInvite } from "medusa-react"
//
//     type Props = {
//     inviteId: string
//     }
//
//     const ResendInvite = ({ inviteId }: Props) => {
//     const resendInvite = useAdminResendInvite(inviteId)
//     // ...
//
//     const handleResend = () => {
//     resendInvite.mutate(void 0, {
//     onSuccess: () => {
//     // invite resent successfully
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default ResendInvite
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/invites/{invite_id}/resend' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Invites
//
// responses:
//
//	200:
//	  description: "OK"
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
func (m *Invite) Resend(context fiber.Ctx) error {
	id, err := utils.ParseUUID(context.Params("id"))
	if err != nil {
		return err
	}

	if err := m.r.InviteService().SetContext(context.Context()).Resend(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).Send(nil)
}

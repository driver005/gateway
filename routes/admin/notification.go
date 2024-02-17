package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type Notification struct {
	r Registry
}

func NewNotification(r Registry) *Notification {
	m := Notification{r: r}
	return &m
}

func (m *Notification) SetRoutes(router fiber.Router) {
	route := router.Group("/notifications")
	route.Get("", m.List)

	route.Post("/:id/resend", m.Resend)
}

// @oas:path [get] /admin/notifications
// operationId: "GetNotifications"
// summary: "List Notifications"
// description: "Retrieve a list of notifications. The notifications can be filtered by fields such as `event_name` or `resource_type`. The notifications can also be paginated."
// x-authenticated: true
// parameters:
//   - (query) offset=0 {integer} The number of inventory items to skip when retrieving the inventory items.
//   - (query) limit=50 {integer} Limit the number of notifications returned.
//   - (query) fields {string} Comma-separated fields that should be included in each returned notification.
//   - (query) expand {string} Comma-separated relations that should be expanded in each returned notification.
//   - (query) event_name {string} Filter by the name of the event that triggered sending this notification.
//   - (query) resource_type {string} Filter by the resource type.
//   - (query) resource_id {string} Filter by the resource ID.
//   - (query) to {string} Filter by the address that the Notification was sent to. This will usually be an email address, but it can also represent other addresses such as a chat bot user id.
//   - (query) include_resends {string} A boolean indicating whether the result set should include resent notifications or not
//
// x-codegen:
//
//	method: list
//	queryParams: AdminGetNotificationsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.notifications.list()
//     .then(({ notifications }) => {
//     console.log(notifications.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminNotifications } from "medusa-react"
//
//     const Notifications = () => {
//     const { notifications, isLoading } = useAdminNotifications()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {notifications && !notifications.length && (
//     <span>No Notifications</span>
//     )}
//     {notifications && notifications.length > 0 && (
//     <ul>
//     {notifications.map((notification) => (
//     <li key={notification.id}>{notification.to}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default Notifications
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/notifications' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Notifications
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminNotificationsListRes"
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
func (m *Notification) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableNotification](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.NotificationService().SetContext(context.Context()).ListAndCount(model, config)
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

// @oas:path [post] /admin/notifications/{id}/resend
// operationId: "PostNotificationsNotificationResend"
// summary: "Resend Notification"
// description: "Resend a previously sent notifications, with the same data but optionally to a different address."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Notification
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostNotificationsNotificationResendReq"
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
//     medusa.admin.notifications.resend(notificationId)
//     .then(({ notification }) => {
//     console.log(notification.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminResendNotification } from "medusa-react"
//
//     type Props = {
//     notificationId: string
//     }
//
//     const Notification = ({ notificationId }: Props) => {
//     const resendNotification = useAdminResendNotification(
//     notificationId
//     )
//     // ...
//
//     const handleResend = () => {
//     resendNotification.mutate({}, {
//     onSuccess: ({ notification }) => {
//     console.log(notification.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Notification
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/notifications/{id}/resend' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Notifications
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminNotificationsRes"
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
func (m *Notification) Resend(context fiber.Ctx) error {
	_, id, config, err := api.BindAll[types.ResendNotification](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.NotificationService().SetContext(context.Context()).Resend(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

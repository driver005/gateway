package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FilterableNotification struct {
	core.FilterModel

	EventName      string    `json:"event_name,omitempty"`
	ResourceType   string    `json:"resource_type,omitempty"`
	ResourceId     uuid.UUID `json:"resource_id,omitempty"`
	To             string    `json:"to,omitempty"`
	IncludeResends bool      `json:"include_resends,omitempty"`
}

// @oas:schema:AdminPostNotificationsNotificationResendReq
// type: object
// description: "The resend details."
// properties:
//
//	to:
//	  description: >-
//	    A new address or user identifier that the Notification should be sent to. If not provided, the previous `to` field of the notification will be used.
//	  type: string
type ResendNotification struct {
	To string `json:"to,omitempty" validate:"omitempty"`
}

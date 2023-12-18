package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// Notification - Notifications a communications sent via Notification Providers as a reaction to internal events such as `order.placed`. Notifications can be used to show a chronological timeline for communications sent to a Customer regarding an Order, and enables resends.
type Notification struct {
	core.Model

	// The name of the event that the notification was sent for.
	EventName string `json:"event_name" gorm:"default:null"`

	// The type of resource that the Notification refers to.
	ResourceType string `json:"resource_type"`

	// The ID of the resource that the Notification refers to.
	ResourceId uuid.NullUUID `json:"resource_id"`

	// The ID of the Customer that the Notification was sent to.
	CustomerId uuid.NullUUID `json:"customer_id" gorm:"default:null"`

	// A customer object. Available if the relation `customer` is expanded.
	Customer *Customer `json:"customer" gorm:"foreignKey:id;references:customer_id"`

	// The address that the Notification was sent to. This will usually be an email address, but represent other addresses such as a chat bot user id
	To string `json:"to"`

	// The data that the Notification was sent with. This contains all the data necessary for the Notification Provider to initiate a resend.
	Data JSONB `json:"data" gorm:"default:null"`

	// The id of the Notification Provider that handles the Notification.
	ParentId uuid.NullUUID `json:"parent_id" gorm:"default:null"`

	ParentNotification *NotificationProvider `json:"parent_notification" gorm:"foreignKey:id;references:parent_id"`

	// The resends that have been completed after the original Notification.
	Resends []Notification `json:"resends" gorm:"foreignKey:id"`

	// The id of the Notification Provider that handles the Notification.
	ProviderId uuid.NullUUID `json:"provider_id" gorm:"default:null"`

	Provider *NotificationProvider `json:"provider" gorm:"foreignKey:id;references:provider_id"`
}

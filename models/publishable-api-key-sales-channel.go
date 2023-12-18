package models

import "github.com/google/uuid"

// This represents the association between the Publishable API keys and Sales Channels
type PublishableApiKeySalesChannel struct {

	// The sales channel's ID
	SalesChannelId uuid.NullUUID `json:"sales_channel_id"`

	// The publishable API key's ID
	PublishableKeyId uuid.NullUUID `json:"publishable_key_id"`
}

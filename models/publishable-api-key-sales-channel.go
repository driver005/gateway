package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:PublishableApiKeySalesChannel
// title: "Publishable API Key Sales Channel"
// description: "This represents the association between the Publishable API keys and Sales Channels"
// type: object
// required:
//   - publishable_key_id
//   - sales_channel_id
//   - created_at
//   - updated_at
//   - deleted_at
//
// properties:
//
//	id:
//	 description: The relation's ID
//	 type: string
//	 example: pksc_01G8X9A7ESKAJXG2H0E6F1MW7A
//	sales_channel_id:
//	  description: The sales channel's ID
//	  type: string
//	  example: sc_01G1G5V21KADXNGH29BJMAJ4B4
//	publishable_key_id:
//	  description: The publishable API key's ID
//	  type: string
//	  example: pak_01G1G5V21KADXNGH29BJMAJ4B4
//	created_at:
//	  description: The date with timezone at which the resource was created.
//	  type: string
//	  format: date-time
//	updated_at:
//	  description: The date with timezone at which the resource was updated.
//	  type: string
//	  format: date-time
//	deleted_at:
//	  description: The date with timezone at which the resource was deleted.
//	  nullable: true
//	  type: string
//	  format: date-time
type PublishableApiKeySalesChannel struct {
	core.BaseModel

	SalesChannelId   uuid.NullUUID `json:"sales_channel_id" gorm:"column:sales_channel_id;primarykey"`
	PublishableKeyId uuid.NullUUID `json:"publishable_key_id" gorm:"column:publishable_key_id;primarykey"`
}

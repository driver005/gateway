package models

import (
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:PublishableApiKey
// title: "Publishable API key"
// description: "A Publishable API key defines scopes that resources are available in. Then, it can be used in request to infer the resources without having to directly pass them. For example, a publishable API key can be associated with one or more sales channels. Then, when the publishable API key is passed in the header of a request, it is inferred what sales channel is being used without having to pass the sales channel as a query or body parameter of the request. Publishable API keys can only be used with sales channels, at the moment."
// type: object
// required:
//   - created_at
//   - created_by
//   - id
//   - revoked_by
//   - revoked_at
//   - title
//   - updated_at
//
// properties:
//
//	id:
//	  description: The key's ID
//	  type: string
//	  example: pk_01G1G5V27GYX4QXNARRQCW1N8T
//	created_by:
//	 description: The unique identifier of the user that created the key.
//	 nullable: true
//	 type: string
//	 example: usr_01G1G5V26F5TB3GPAPNJ8X1S3V
//	revoked_by:
//	  description: The unique identifier of the user that revoked the key.
//	  nullable: true
//	  type: string
//	  example: usr_01G1G5V26F5TB3GPAPNJ8X1S3V
//	revoked_at:
//	  description: The date with timezone at which the key was revoked.
//	  nullable: true
//	  type: string
//	  format: date-time
//	title:
//	  description: The key's title.
//	  type: string
//	created_at:
//	  description: The date with timezone at which the resource was created.
//	  type: string
//	  format: date-time
//	updated_at:
//	  description: The date with timezone at which the resource was updated.
//	  type: string
//	  format: date-time
type PublishableApiKey struct {
	core.BaseModel

	CreatedBy uuid.NullUUID `json:"created_by" gorm:"column:created_by"`
	RevokedBy uuid.NullUUID `json:"revoked_by" gorm:"column:revoked_by"`
	RevokedAt *time.Time    `json:"revoked_at" gorm:"column:revoked_at"`
	Title     string        `json:"title" gorm:"column:title"`
}

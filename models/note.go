package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:Note
// title: "Note"
// description: "A Note is an element that can be used in association with different resources to allow admin users to describe additional information. For example, they can be used to add additional information about orders."
// type: object
// required:
//   - author_id
//   - created_at
//   - deleted_at
//   - id
//   - metadata
//   - resource_id
//   - resource_type
//   - updated_at
//   - value
//
// properties:
//
//	id:
//	  description: The note's ID
//	  type: string
//	  example: note_01G8TM8ENBMC7R90XRR1G6H26Q
//	resource_type:
//	  description: The type of resource that the Note refers to.
//	  type: string
//	  example: order
//	resource_id:
//	  description: The ID of the resource that the Note refers to.
//	  type: string
//	  example: order_01G8TJSYT9M6AVS5N4EMNFS1EK
//	value:
//	  description: The contents of the note.
//	  type: string
//	  example: This order must be fulfilled on Monday
//	author_id:
//	  description: The ID of the user that created the note.
//	  nullable: true
//	  type: string
//	  example: usr_01G1G5V26F5TB3GPAPNJ8X1S3V
//	author:
//	  description: The details of the user that created the note.
//	  x-expandable: "author"
//	  nullable: true
//	  $ref: "#/components/schemas/User"
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
//	metadata:
//	  description: An optional key-value map with additional details
//	  nullable: true
//	  type: object
//	  example: {car: "white"}
type Note struct {
	core.Model

	ResourceType string        `json:"resource_type"  gorm:"column:resource_type"`
	ResourceId   uuid.NullUUID `json:"resource_id"  gorm:"column:resource_id"`
	Value        string        `json:"value"  gorm:"column:value"`
	AuthorId     uuid.NullUUID `json:"author_id"  gorm:"column:author_id"`
	Author       *User         `json:"author"  gorm:"column:author;foreignKey:AuthorId"`
}

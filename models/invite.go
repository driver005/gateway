package models

import (
	"time"

	"github.com/driver005/gateway/core"
)

// @oas:schema:Invite
// title: "Invite"
// description: "An invite is created when an admin user invites a new user to join the store's team. Once the invite is accepted, it's deleted."
// type: object
// required:
//   - accepted
//   - created_at
//   - deleted_at
//   - expires_at
//   - id
//   - metadata
//   - role
//   - token
//   - updated_at
//   - user_email
//
// properties:
//
//	id:
//	  type: string
//	  description: The invite's ID
//	  example: invite_01G8TKE4XYCTHSCK2GDEP47RE1
//	user_email:
//	  description: The email of the user being invited.
//	  type: string
//	  format: email
//	role:
//	  description: The user's role. These roles don't change the privileges of the user.
//	  nullable: true
//	  type: string
//	  enum:
//	    - admin
//	    - member
//	    - developer
//	  default: member
//	accepted:
//	  description: Whether the invite was accepted or not.
//	  type: boolean
//	  default: false
//	token:
//	  description: The token used to accept the invite.
//	  type: string
//	expires_at:
//	  description: The date the invite expires at.
//	  type: string
//	  format: date-time
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
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type Invite struct {
	core.SoftDeletableModel

	UserEmail string     `json:"user_email" gorm:"column:user_email"`
	Role      UserRole   `json:"role" gorm:"column:role;default:'member'"`
	Accepted  bool       `json:"accepted" gorm:"column:accepted;default:false"`
	Token     string     `json:"token" gorm:"column:token"`
	ExpiresAt *time.Time `json:"expores_at" gorm:"column:expores_at"`
}

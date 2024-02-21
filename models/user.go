package models

import (
	"database/sql/driver"

	"github.com/driver005/gateway/core"
)

// @oas:schema:User
// title: "User"
// description: "A User is an administrator who can manage store settings and data."
// type: object
// required:
//   - api_token
//   - created_at
//   - deleted_at
//   - email
//   - first_name
//   - id
//   - last_name
//   - metadata
//   - role
//   - updated_at
//
// properties:
//
//	id:
//	  description: The user's ID
//	  type: string
//	  example: usr_01G1G5V26F5TB3GPAPNJ8X1S3V
//	role:
//	  description: The user's role. These roles don't provide any different privileges.
//	  type: string
//	  enum:
//	    - admin
//	    - member
//	    - developer
//	  default: member
//	email:
//	  description: The email of the User
//	  type: string
//	  format: email
//	first_name:
//	  description: The first name of the User
//	  nullable: true
//	  type: string
//	  example: Levi
//	last_name:
//	  description: The last name of the User
//	  nullable: true
//	  type: string
//	  example: Bogan
//	api_token:
//	  description: An API token associated with the user.
//	  nullable: true
//	  type: string
//	  example: null
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
type User struct {
	core.Model

	Role         UserRole `json:"role" gorm:"column:role;default:'member'"`
	Email        string   `json:"email" gorm:"column:email"`
	FirstName    string   `json:"first_name" gorm:"column:first_name"`
	LastName     string   `json:"last_name" gorm:"column:last_name"`
	PasswordHash string   `json:"password_hash" gorm:"column:password_hash"`
	ApiToken     string   `json:"api_token" gorm:"column:api_token"`
}

type UserRole string

// Defines values for UserRole.
const (
	UserRoleAdmin     UserRole = "admin"
	UserRoleMember    UserRole = "member"
	UserRoleDeveloper UserRole = "developer"
)

func (sp *UserRole) Scan(value interface{}) error {
	*sp = UserRole(value.([]byte))
	return nil
}

func (sp UserRole) Value() (driver.Value, error) {
	return string(sp), nil
}

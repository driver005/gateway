package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
)

// @oas:schema:AdminCreateUserRequest
// type: object
// required:
//   - email
//   - password
//
// properties:
//
//	email:
//	  description: "The User's email."
//	  type: string
//	  format: email
//	first_name:
//	  description: "The first name of the User."
//	  type: string
//	last_name:
//	  description: "The last name of the User."
//	  type: string
//	role:
//	  description: "The role assigned to the user. These roles don't provide any different privileges."
//	  type: string
//	  enum: [admin, member, developer]
//	password:
//	  description: "The User's password."
//	  type: string
//	  format: password
type CreateUserInput struct {
	// Id        uuid.UUID       `json:"id,omitempty" validate:"omitempty"`
	Email     string          `json:"email"`
	FirstName string          `json:"first_name,omitempty" validate:"omitempty"`
	LastName  string          `json:"last_name,omitempty" validate:"omitempty"`
	APIToken  string          `json:"api_token,omitempty" validate:"omitempty"`
	Password  string          `json:"password,omitempty" validate:"omitempty"`
	Role      models.UserRole `json:"role,omitempty" validate:"omitempty"`
	Metadata  core.JSONB      `json:"metadata,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminUpdateUserRequest
// type: object
// properties:
//
//	first_name:
//	  description: "The first name of the User."
//	  type: string
//	last_name:
//	  description: "The last name of the User."
//	  type: string
//	role:
//	  description: "The role assigned to the user. These roles don't provide any different privileges."
//	  type: string
//	  enum: [admin, member, developer]
//	api_token:
//	  description: "The API token of the User."
//	  type: string
//	metadata:
//	  description: An optional set of key-value pairs with additional information.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type UpdateUserInput struct {
	Email        string          `json:"email,omitempty" validate:"omitempty"`
	FirstName    string          `json:"first_name,omitempty" validate:"omitempty"`
	LastName     string          `json:"last_name,omitempty" validate:"omitempty"`
	PasswordHash string          `json:"password_hash,omitempty" validate:"omitempty"`
	APIToken     string          `json:"api_token,omitempty" validate:"omitempty"`
	Role         models.UserRole `json:"role,omitempty" validate:"omitempty"`
	Metadata     core.JSONB      `json:"metadata,omitempty" validate:"omitempty"`
}

type FilterableUser struct {
	core.FilterModel

	Email     string `json:"email,omitempty" validate:"omitempty"`
	FirstName string `json:"first_name,omitempty" validate:"omitempty"`
	LastName  string `json:"last_name,omitempty" validate:"omitempty"`
}

// @oas:schema:ResetPasswordTokenRequest
// type: object
// description: "The details of the password reset token request."
// required:
//   - email
//
// properties:
//
//	email:
//	  description: "The User's email."
//	  type: string
//	  format: email
type UserResetPasswordToken struct {
	Email string `json:"email"`
}

// @oas:schema:ResetPasswordRequest
// type: object
// description: "The details of the password reset request."
// required:
//   - token
//   - password
//
// properties:
//
//	email:
//	  description: "The User's email."
//	  type: string
//	  format: email
//	token:
//	  description: "The password-reset token generated when the password reset was requested."
//	  type: string
//	password:
//	  description: "The User's new password."
//	  type: string
//	  format: password
type UserResetPasswordRequest struct {
	Email    string `json:"email,omitempty" validate:"omitempty"`
	Token    string `json:"token"`
	Password string `json:"password"`
}

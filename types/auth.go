package types

import "github.com/driver005/gateway/models"

type AuthenticateResult struct {
	Success  bool             `json:"success"`
	User     *models.User     `json:"user,omitempty" validate:"omitempty"`
	Customer *models.Customer `json:"customer,omitempty" validate:"omitempty"`
	Error    string           `json:"error,omitempty" validate:"omitempty"`
}

// @oas:schema:PostAuthReq
// type: object
// description: The admin's credentials used to log in.
// required:
//   - email
//   - password
//
// properties:
//
//	email:
//	  type: string
//	  description: The user's email.
//	  format: email
//	password:
//	  type: string
//	  description: The user's password.
//	  format: password
type CreateAuth struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

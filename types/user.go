package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
)

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

package types

import "github.com/driver005/gateway/models"

type AuthenticateResult struct {
	Success  bool             `json:"success"`
	User     *models.User     `json:"user,omitempty" validate:"omitempty"`
	Customer *models.Customer `json:"customer,omitempty" validate:"omitempty"`
	Error    string           `json:"error,omitempty" validate:"omitempty"`
}

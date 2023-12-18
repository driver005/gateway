package types

import "github.com/driver005/gateway/models"

type AuthenticateResult struct {
	Success  bool
	User     *models.User
	Customer *models.Customer
	Error    string
}

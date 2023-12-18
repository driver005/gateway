package types

import "github.com/driver005/gateway/core"

type FilterableUser struct {
	core.FilterModel

	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

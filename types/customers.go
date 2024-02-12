package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type Group struct {
	Id uuid.UUID `json:"id,omitempty" validate:"omitempty"`
}

// AdminListCustomerSelector filters used to filter retrieved customers.
type FilterableCustomer struct {
	core.FilterModel
	Email      string     `json:"email,omitempty" validate:"omitempty"`
	FirstName  string     `json:"first_name,omitempty" validate:"omitempty"`
	LastName   string     `json:"last_name,omitempty" validate:"omitempty"`
	Phone      string     `json:"phone,omitempty" validate:"omitempty"`
	HasAccount bool       `json:"has_account,omitempty" validate:"omitempty"`
	Groups     uuid.UUIDs `json:"groups,omitempty" validate:"omitempty"`
}

// CreateCustomerInput represents the input for creating a customer.
type CreateCustomerInput struct {
	Email        string     `json:"email"`
	Password     string     `json:"password,omitempty" validate:"omitempty"`
	PasswordHash string     `json:"password_hash,omitempty" validate:"omitempty"`
	HasAccount   bool       `json:"has_account,omitempty" validate:"omitempty"`
	FirstName    string     `json:"first_name,omitempty" validate:"omitempty"`
	LastName     string     `json:"last_name,omitempty" validate:"omitempty"`
	Phone        string     `json:"phone,omitempty" validate:"omitempty"`
	Metadata     core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

// UpdateCustomerInput represents the input for updating a customer.
type UpdateCustomerInput struct {
	Password         string          `json:"password,omitempty" validate:"omitempty"`
	Metadata         core.JSONB      `json:"metadata,omitempty" validate:"omitempty"`
	BillingAddress   *AddressPayload `json:"billing_address,omitempty" validate:"omitempty"`
	BillingAddressId uuid.UUID       `json:"billing_address_id,omitempty" validate:"omitempty"`
	Groups           []Group         `json:"groups,omitempty" validate:"omitempty"`
	Email            string          `json:"email,omitempty" validate:"omitempty"`
	FirstName        string          `json:"first_name,omitempty" validate:"omitempty"`
	LastName         string          `json:"last_name,omitempty" validate:"omitempty"`
	Phone            string          `json:"phone,omitempty" validate:"omitempty"`
}

type CustomerAddAddress struct {
	Address *AddressCreatePayload `json:"address"`
}

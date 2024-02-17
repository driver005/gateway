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

// @oas:schema:PostCustomersReq
// type: object
// description: "The details of the customer to create."
// required:
//   - email
//   - first_name
//   - last_name
//   - password
//
// properties:
//
//	email:
//	  type: string
//	  description: The customer's email.
//	  format: email
//	first_name:
//	  type: string
//	  description: The customer's first name.
//	last_name:
//	  type: string
//	  description: The customer's last name.
//	password:
//	  type: string
//	  description: The customer's password.
//	  format: password
//	phone:
//	  type: string
//	  description: The customer's phone number.
//	metadata:
//	  description: An optional set of key-value pairs to hold additional information.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
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

// @oas:schema:PostCustomersCustomerReq
// type: object
// description: "The details of the customer to update."
// properties:
//
//	email:
//	  type: string
//	  description: The Customer's email. You can't update the email of a registered customer.
//	  format: email
//	first_name:
//	  type: string
//	  description:  The Customer's first name.
//	last_name:
//	  type: string
//	  description:  The Customer's last name.
//	phone:
//	  type: string
//	  description: The Customer's phone number.
//	password:
//	  type: string
//	  description: The Customer's password.
//	  format: password
//	groups:
//	  type: array
//	  description: A list of customer groups to which the customer belongs.
//	  items:
//	    type: object
//	    required:
//	      - id
//	    properties:
//	      id:
//	        description: The ID of a customer group
//	        type: string
//	metadata:
//	  description: An optional set of key-value pairs to hold additional information.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
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

// @oas:schema:StorePostCustomersCustomerAddressesReq
// type: object
// required:
//   - address
//
// properties:
//
//	address:
//	  description: "The Address to add to the Customer's saved addresses."
//	  $ref: "#/components/schemas/AddressCreatePayload"
type CustomerAddAddress struct {
	Address *AddressCreatePayload `json:"address"`
}

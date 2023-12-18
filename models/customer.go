package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// Customer - Represents a customer
type Customer struct {
	core.Model

	// The customer's email
	Email string `json:"email"`

	// The customer's first name
	FirstName string `json:"first_name" gorm:"default:null"`

	// The customer's first name
	LastName string `json:"last_name" gorm:"default:null"`

	// The customer's billing address ID
	BillingAddressId uuid.NullUUID `json:"billing_address_id" gorm:"default:null"`

	BillingAddress *Address `json:"billing_address" gorm:"foreignKey:id;references:billing_address_id"`

	// Available if the relation `shipping_addresses` is expanded.
	ShippingAddresses []Address `json:"shipping_addresses" gorm:"foreignKey:id"`

	// Password of the customer
	Password string `json:"password" gorm:"-" sql:"-"`

	// Password hash of the customer
	PasswordHash string `json:"password_hash" gorm:"default:null"`

	// The customer's phone number
	Phone string `json:"phone" gorm:"default:null"`

	// Whether the customer has an account or not
	HasAccount bool `json:"has_account" gorm:"default:null"`

	// Available if the relation `orders` is expanded.
	Orders []Order `json:"orders" gorm:"foreignKey:id"`

	// The customer groups the customer belongs to. Available if the relation `groups` is expanded.
	Groups []CustomerGroup `json:"groups" gorm:"many2many:customer_group"`
}

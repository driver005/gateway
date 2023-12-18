package models

import (
	"github.com/driver005/gateway/core"

	"github.com/google/uuid"
)

// An address.
type Address struct {
	core.Model

	// Address line 1
	Address1 string `json:"address_1" gorm:"default:null"`

	// Address line 2
	Address2 string `json:"address_2" gorm:"default:null"`

	// City
	City string `json:"city" gorm:"default:null"`

	// Company name
	Company string `json:"company" gorm:"default:null"`

	// A country object. Available if the relation `country` is expanded.
	Country *Country `json:"country" gorm:"foreignKey:id;references:country_code"`

	// The 2 character ISO code of the country in lower case
	CountryCode string `json:"country_code" gorm:"default:null"`

	// Available if the relation `customer` is expanded.
	Customer *Customer `json:"customer" gorm:"foreignKey:id;references:customer_id"`

	// ID of the customer this address belongs to
	CustomerId uuid.NullUUID `json:"customer_id" gorm:"default:null"`

	// First name
	FirstName string `json:"first_name" gorm:"default:null"`

	// Last name
	LastName string `json:"last_name" gorm:"default:null"`

	// Phone Number
	Phone string `json:"phone" gorm:"default:null"`

	// Postal Code
	PostalCode string `json:"postal_code" gorm:"default:null"`

	// Province
	Province string `json:"province" gorm:"default:null"`
}

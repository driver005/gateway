package models

import (
	"github.com/driver005/gateway/core"

	"github.com/google/uuid"
)

// @oas:schema:Address
// title: "Address"
// description: "An address is used across the Medusa backend within other schemas and object types. For example, a customer's billing and shipping addresses both use the Address entity."
// type: object
// required:
//   - address_1
//   - address_2
//   - city
//   - company
//   - country_code
//   - created_at
//   - customer_id
//   - deleted_at
//   - first_name
//   - id
//   - last_name
//   - metadata
//   - phone
//   - postal_code
//   - province
//   - updated_at
//
// properties:
//
//	id:
//	  type: string
//	  description: ID of the address
//	  example: addr_01G8ZC9VS1XVE149MGH2J7QSSH
//	customer_id:
//	  description: ID of the customer this address belongs to
//	  nullable: true
//	  type: string
//	  example: cus_01G2SG30J8C85S4A5CHM2S1NS2
//	customer:
//	  description: Available if the relation `customer` is expanded.
//	  nullable: true
//	  $ref: "#/components/schemas/Customer"
//	company:
//	  description: Company name
//	  nullable: true
//	  type: string
//	  example: Acme
//	first_name:
//	  description: First name
//	  nullable: true
//	  type: string
//	  example: Arno
//	last_name:
//	  description: Last name
//	  nullable: true
//	  type: string
//	  example: Willms
//	address_1:
//	  description: Address line 1
//	  nullable: true
//	  type: string
//	  example: 14433 Kemmer Court
//	address_2:
//	  description: Address line 2
//	  nullable: true
//	  type: string
//	  example: Suite 369
//	city:
//	  description: City
//	  nullable: true
//	  type: string
//	  example: South Geoffreyview
//	country_code:
//	  description: The 2 character ISO code of the country in lower case
//	  nullable: true
//	  type: string
//	  externalDocs:
//	    url: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements
//	    description: See a list of codes.
//	  example: st
//	country:
//	  description: A country object.
//	  x-expandable: "country"
//	  nullable: true
//	  $ref: "#/components/schemas/Country"
//	province:
//	  description: Province
//	  nullable: true
//	  type: string
//	  example: Kentucky
//	postal_code:
//	  description: Postal Code
//	  nullable: true
//	  type: string
//	  example: 72093
//	phone:
//	  description: Phone Number
//	  nullable: true
//	  type: string
//	  example: 16128234334802
//	created_at:
//	  type: string
//	  description: "The date with timezone at which the resource was created."
//	  format: date-time
//	updated_at:
//	  type: string
//	  description: "The date with timezone at which the resource was updated."
//	  format: date-time
//	deleted_at:
//	  description: The date with timezone at which the resource was deleted.
//	  nullable: true
//	  type: string
//	  format: date-time
//	metadata:
//	  description: An optional key-value map with additional details
//	  nullable: true
//	  type: object
//	  example: {car: "white"}
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type Address struct {
	core.Model

	Address1    string        `json:"address_1" gorm:"default:null"`
	Address2    string        `json:"address_2" gorm:"default:null"`
	City        string        `json:"city" gorm:"default:null"`
	Company     string        `json:"company" gorm:"default:null"`
	Country     *Country      `json:"country" gorm:"foreignKey:id;references:country_code"`
	CountryCode string        `json:"country_code" gorm:"default:null"`
	Customer    *Customer     `json:"customer" gorm:"foreignKey:id;references:customer_id"`
	CustomerId  uuid.NullUUID `json:"customer_id" gorm:"default:null"`
	FirstName   string        `json:"first_name" gorm:"default:null"`
	LastName    string        `json:"last_name" gorm:"default:null"`
	Phone       string        `json:"phone" gorm:"default:null"`
	PostalCode  string        `json:"postal_code" gorm:"default:null"`
	Province    string        `json:"province" gorm:"default:null"`
}

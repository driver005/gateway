package models

import (
	"database/sql/driver"
	"time"

	"github.com/driver005/gateway/core"
)

// @oas:schema:PriceList
// title: "Price List"
// description: "A Price List represents a set of prices that override the default price for one or more product variants."
// type: object
// required:
//   - created_at
//   - deleted_at
//   - description
//   - ends_at
//   - id
//   - name
//   - starts_at
//   - status
//   - type
//   - updated_at
//
// properties:
//
//	id:
//	  description: The price list's ID
//	  type: string
//	  example: pl_01G8X3CKJXCG5VXVZ87H9KC09W
//	name:
//	  description: The price list's name
//	  type: string
//	  example: VIP Prices
//	description:
//	  description: The price list's description
//	  type: string
//	  example: Prices for VIP customers
//	type:
//	  description: The type of Price List. This can be one of either `sale` or `override`.
//	  type: string
//	  enum:
//	    - sale
//	    - override
//	  default: sale
//	status:
//	  description: The status of the Price List
//	  type: string
//	  enum:
//	    - active
//	    - draft
//	  default: draft
//	starts_at:
//	  description: The date with timezone that the Price List starts being valid.
//	  nullable: true
//	  type: string
//	  format: date-time
//	ends_at:
//	  description: The date with timezone that the Price List stops being valid.
//	  nullable: true
//	  type: string
//	  format: date-time
//	customer_groups:
//	  description: The details of the customer groups that the Price List can apply to.
//	  type: array
//	  x-expandable: "customer_groups"
//	  items:
//	    $ref: "#/components/schemas/CustomerGroup"
//	prices:
//	  description: The prices that belong to the price list, represented as a Money Amount.
//	  type: array
//	  x-expandable: "prices"
//	  items:
//	    $ref: "#/components/schemas/MoneyAmount"
//	includes_tax:
//	  description: "Whether the price list prices include tax"
//	  type: boolean
//	  x-featureFlag: "tax_inclusive_pricing"
//	  default: false
//	created_at:
//	  description: The date with timezone at which the resource was created.
//	  type: string
//	  format: date-time
//	updated_at:
//	  description: The date with timezone at which the resource was updated.
//	  type: string
//	  format: date-time
//	deleted_at:
//	  description: The date with timezone at which the resource was deleted.
//	  nullable: true
//	  type: string
//	  format: date-time
type PriceList struct {
	core.SoftDeletableModel

	CustomerGroups []CustomerGroup `json:"customer_groups" gorm:"many2many:price_list_customer_groups"`
	Description    string          `json:"description" gorm:"column:description"`
	EndsAt         *time.Time      `json:"ends_at" gorm:"column:ends_at"`
	IncludesTax    bool            `json:"includes_tax" gorm:"column:includes_tax"`
	Name           string          `json:"name" gorm:"column:name"`
	Prices         []MoneyAmount   `json:"prices" gorm:"foreignKey:Id"`
	StartsAt       *time.Time      `json:"starts_at" gorm:"column:starts_at"`
	Status         PriceListStatus `json:"status" gorm:"column:status;default:'draft'"`
	Type           PriceListType   `json:"type" gorm:"column:type;default:'sale'"`
}

// The status of the Price List
type PriceListStatus string

// Defines values for PriceListStatus.
const (
	PriceListStatusActive PriceListStatus = "active"
	PriceListStatusDraft  PriceListStatus = "draft"
)

func (pl *PriceListStatus) Scan(value interface{}) error {
	*pl = PriceListStatus(value.([]byte))
	return nil
}

func (pl PriceListStatus) Value() (driver.Value, error) {
	return string(pl), nil
}

// The type of Price List. This can be one of either `sale` or `override`.
type PriceListType string

const (
	Override PriceListType = "override"
	Sale     PriceListType = "sale"
)

func (pl *PriceListType) Scan(value interface{}) error {
	*pl = PriceListType(value.([]byte))
	return nil
}

func (pl PriceListType) Value() (driver.Value, error) {
	return string(pl), nil
}

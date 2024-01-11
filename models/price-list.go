package models

import (
	"database/sql/driver"
	"time"

	"github.com/driver005/gateway/core"
)

// Price Lists represents a set of prices that overrides the default price for one or more product variants.
type PriceList struct {
	core.Model

	// The Customer Groups that the Price List applies to. Available if the relation `customer_groups` is expanded.
	CustomerGroups []CustomerGroup `json:"customer_groups" gorm:"foreignKey:id"`

	// The price list's description
	Description string `json:"description"`

	// The date with timezone that the Price List stops being valid.
	EndsAt *time.Time `json:"ends_at" gorm:"default:null"`

	// [EXPERIMENTAL] Does the price list prices include tax
	IncludesTax bool `json:"includes_tax" gorm:"default:null"`

	// The price list's name
	Name string `json:"name"`

	// The Money Amounts that are associated with the Price List. Available if the relation `prices` is expanded.
	Prices []MoneyAmount `json:"prices" gorm:"foreignKey:id"`

	// The date with timezone that the Price List starts being valid.
	StartsAt *time.Time `json:"starts_at" gorm:"default:null"`

	// The status of the Price List
	Status PriceListStatus `json:"status" gorm:"default:null"`

	// The type of Price List. This can be one of either `sale` or `override`.
	Type PriceListType `json:"type" gorm:"default:null"`
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

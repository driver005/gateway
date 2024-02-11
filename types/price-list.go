package types

import (
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type FilterablePriceList struct {
	core.FilterModel

	Q              string                   `json:"q,omitempty" validate:"omitempty"`
	Status         []models.PriceListStatus `json:"status,omitempty" validate:"omitempty"`
	Name           string                   `json:"name,omitempty" validate:"omitempty"`
	CustomerGroups []string                 `json:"customer_groups,omitempty" validate:"omitempty"`
	Description    string                   `json:"description,omitempty" validate:"omitempty"`
	Type           []models.PriceListType   `json:"type,omitempty" validate:"omitempty"`
}

type AdminPriceListPricesUpdateReq struct {
	Id           uuid.UUID `json:"id,omitempty" validate:"omitempty"`
	RegionId     uuid.UUID `json:"region_id,omitempty" validate:"omitempty"`
	CurrencyCode string    `json:"currency_code,omitempty" validate:"omitempty"`
	VariantId    uuid.UUID `json:"variant_id"`
	Amount       float64   `json:"amount"`
	MinQuantity  int       `json:"min_quantity,omitempty" validate:"omitempty"`
	MaxQuantity  int       `json:"max_quantity,omitempty" validate:"omitempty"`
}

type AdminPriceListPricesCreateReq struct {
	RegionId     uuid.UUID `json:"region_id,omitempty" validate:"omitempty"`
	CurrencyCode string    `json:"currency_code,omitempty" validate:"omitempty"`
	Amount       float64   `json:"amount"`
	VariantId    uuid.UUID `json:"variant_id"`
	MinQuantity  int       `json:"min_quantity,omitempty" validate:"omitempty"`
	MaxQuantity  int       `json:"max_quantity,omitempty" validate:"omitempty"`
}

type CustomerGroups struct {
	Id uuid.UUID `json:"id"`
}

type CreatePriceListInput struct {
	Name           string                      `json:"name"`
	Description    string                      `json:"description"`
	Type           models.PriceListType        `json:"type"`
	Status         models.PriceListStatus      `json:"status,omitempty" validate:"omitempty"`
	Prices         []PriceListPriceCreateInput `json:"prices"`
	CustomerGroups []CustomerGroups            `json:"customer_groups,omitempty" validate:"omitempty"`
	StartsAt       *time.Time                  `json:"starts_at,omitempty" validate:"omitempty"`
	EndsAt         *time.Time                  `json:"ends_at,omitempty" validate:"omitempty"`
	IncludesTax    bool                        `json:"includes_tax,omitempty" validate:"omitempty"`
}

type UpdatePriceListInput struct {
	Name           string                      `json:"name,omitempty" validate:"omitempty"`
	Description    string                      `json:"description,omitempty" validate:"omitempty"`
	StartsAt       *time.Time                  `json:"starts_at,omitempty" validate:"omitempty"`
	EndsAt         *time.Time                  `json:"ends_at,omitempty" validate:"omitempty"`
	Status         models.PriceListStatus      `json:"status,omitempty" validate:"omitempty"`
	Type           models.PriceListType        `json:"type,omitempty" validate:"omitempty"`
	IncludesTax    bool                        `json:"includes_tax,omitempty" validate:"omitempty"`
	Prices         []PriceListPriceCreateInput `json:"prices,omitempty" validate:"omitempty"`
	CustomerGroups []CustomerGroups            `json:"customer_groups,omitempty" validate:"omitempty"`
}

type PriceListPriceUpdateInput struct {
	Id           uuid.UUID `json:"id,omitempty" validate:"omitempty"`
	VariantId    uuid.UUID `json:"variant_id,omitempty" validate:"omitempty"`
	RegionId     uuid.UUID `json:"region_id,omitempty" validate:"omitempty"`
	CurrencyCode string    `json:"currency_code,omitempty" validate:"omitempty"`
	Amount       float64   `json:"amount,omitempty" validate:"omitempty"`
	MinQuantity  int       `json:"min_quantity,omitempty" validate:"omitempty"`
	MaxQuantity  int       `json:"max_quantity,omitempty" validate:"omitempty"`
}

type PriceListPriceCreateInput struct {
	RegionId     uuid.UUID `json:"region_id,omitempty" validate:"omitempty"`
	CurrencyCode string    `json:"currency_code,omitempty" validate:"omitempty"`
	VariantId    uuid.UUID `json:"variant_id"`
	Amount       float64   `json:"amount"`
	MinQuantity  int       `json:"min_quantity,omitempty" validate:"omitempty"`
	MaxQuantity  int       `json:"max_quantity,omitempty" validate:"omitempty"`
}

type PriceListLoadConfig struct {
	IncludeDiscountPrices bool      `json:"include_discount_prices,omitempty" validate:"omitempty"`
	CustomerId            uuid.UUID `json:"customer_id,omitempty" validate:"omitempty"`
	CartId                uuid.UUID `json:"cart_id,omitempty" validate:"omitempty"`
	RegionId              uuid.UUID `json:"region_id,omitempty" validate:"omitempty"`
	CurrencyCode          string    `json:"currency_code,omitempty" validate:"omitempty"`
}

type AddPriceListPrices struct {
	Prices   []PriceListPriceCreateInput `json:"prices"`
	Override bool                        `json:"override,omitempty" validate:"omitempty"`
}

type DeletePriceListPrices struct {
	PriceIds uuid.UUIDs `json:"price_ids"`
}

type DeletePriceListPricesBatch struct {
	ProductIds []uuid.UUID `json:"product_ids"`
}

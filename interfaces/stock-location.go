package interfaces

import (
	"context"

	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type StringComparisonOperator string

type StockLocationAddressDTO struct {
	Id          uuid.UUID              `json:"id,omitempty"`
	Address1    string                 `json:"address_1"`
	Address2    *string                `json:"address_2,omitempty"`
	Company     *string                `json:"company,omitempty"`
	City        *string                `json:"city,omitempty"`
	CountryCode string                 `json:"country_code"`
	Phone       *string                `json:"phone,omitempty"`
	PostalCode  *string                `json:"postal_code,omitempty"`
	Province    *string                `json:"province,omitempty"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
	DeletedAt   *string                `json:"deleted_at,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type StockLocationDTO struct {
	Id        uuid.UUID                `json:"id"`
	Name      string                   `json:"name"`
	Metadata  map[string]interface{}   `json:"metadata,omitempty"`
	AddressID string                   `json:"address_id"`
	Address   *StockLocationAddressDTO `json:"address,omitempty"`
	CreatedAt string                   `json:"created_at"`
	UpdatedAt string                   `json:"updated_at"`
	DeletedAt *string                  `json:"deleted_at,omitempty"`
}

type StockLocationExpandedDTO struct {
	StockLocationDTO
	SalesChannels []interface{} `json:"sales_channels,omitempty"`
}

type FilterableStockLocation struct {
	Id   uuid.UUID `json:"id,omitempty"`
	Name string    `json:"name,omitempty"`
}

type StockLocationAddressInput struct {
	Address1    string                 `json:"address_1"`
	Address2    *string                `json:"address_2,omitempty"`
	City        *string                `json:"city,omitempty"`
	CountryCode string                 `json:"country_code"`
	Phone       *string                `json:"phone,omitempty"`
	PostalCode  *string                `json:"postal_code,omitempty"`
	Province    *string                `json:"province,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type CreateStockLocationInput struct {
	Name      string                     `json:"name"`
	AddressID *string                    `json:"address_id,omitempty"`
	Address   *StockLocationAddressInput `json:"address,omitempty"`
	Metadata  map[string]interface{}     `json:"metadata,omitempty"`
}

type UpdateStockLocationInput struct {
	Name      *string                    `json:"name,omitempty"`
	AddressID *string                    `json:"address_id,omitempty"`
	Address   *StockLocationAddressInput `json:"address,omitempty"`
	Metadata  map[string]interface{}     `json:"metadata,omitempty"`
}

type IStockLocationService interface {
	List(context context.Context, selector FilterableStockLocation, config sql.Options) ([]StockLocationDTO, *utils.ApplictaionError)
	ListAndCount(context context.Context, selector FilterableStockLocation, config sql.Options) ([]StockLocationDTO, *int64, *utils.ApplictaionError)
	Retrieve(context context.Context, id uuid.UUID, config sql.Options) (StockLocationDTO, *utils.ApplictaionError)
	Create(context context.Context, input CreateStockLocationInput) (StockLocationDTO, *utils.ApplictaionError)
	Update(context context.Context, id uuid.UUID, input UpdateStockLocationInput) (StockLocationDTO, *utils.ApplictaionError)
	Delete(context context.Context, id uuid.UUID) *utils.ApplictaionError
}

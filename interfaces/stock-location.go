package interfaces

import (
	"context"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type StringComparisonOperator string

type StockLocationAddressDTO struct {
	Id          uuid.UUID  `json:"id,omitempty"`
	Address1    string     `json:"address_1"`
	Address2    *string    `json:"address_2,omitempty"`
	Company     *string    `json:"company,omitempty"`
	City        *string    `json:"city,omitempty"`
	CountryCode string     `json:"country_code"`
	Phone       *string    `json:"phone,omitempty"`
	PostalCode  *string    `json:"postal_code,omitempty"`
	Province    *string    `json:"province,omitempty"`
	CreatedAt   string     `json:"created_at"`
	UpdatedAt   string     `json:"updated_at"`
	DeletedAt   *string    `json:"deleted_at,omitempty"`
	Metadata    core.JSONB `json:"metadata,omitempty"`
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

/**
 * @schema AdminPostStockLocationsReqAddress
 * type: object
 * required:
 *   - address_1
 *   - country_code
 * properties:
 *   address_1:
 *     type: string
 *     description: Stock location address
 *     example: 35, Jhon Doe Ave
 *   address_2:
 *     type: string
 *     description: Stock location address' complement
 *     example: apartment 4432
 *   company:
 *     type: string
 *     description: Stock location address' company
 *   city:
 *     type: string
 *     description: Stock location address' city
 *     example: Mexico city
 *   country_code:
 *     description: "The two character ISO code for the country."
 *     type: string
 *     externalDocs:
 *       url: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements
 *       description: See a list of codes.
 *   phone:
 *     type: string
 *     description: Stock location address' phone number
 *     example: +1 555 61646
 *   postal_code:
 *     type: string
 *     description: Stock location address' postal code
 *     example: HD3-1G8
 *   province:
 *     type: string
 *     description: Stock location address' province
 *     example: Sinaloa
 */
type StockLocationAddressInput struct {
	Address1    string     `json:"address_1"`
	Address2    *string    `json:"address_2,omitempty"`
	City        *string    `json:"city,omitempty"`
	CountryCode string     `json:"country_code"`
	Phone       *string    `json:"phone,omitempty"`
	PostalCode  *string    `json:"postal_code,omitempty"`
	Province    *string    `json:"province,omitempty"`
	Metadata    core.JSONB `json:"metadata,omitempty"`
}

/**
 * @schema AdminPostStockLocationsReq
 * type: object
 * description: "The details of the stock location to create."
 * required:
 *   - name
 * properties:
 *   name:
 *     description: the name of the stock location
 *     type: string
 *   address_id:
 *     description: the ID of an existing stock location address to associate with the stock location. Only required if `address` is not provided.
 *     type: string
 *   metadata:
 *     type: object
 *     description: An optional key-value map with additional details
 *     example: {car: "white"}
 *     externalDocs:
 *       description: "Learn about the metadata attribute, and how to delete and update it."
 *       url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
 *   address:
 *     description: A new stock location address to create and associate with the stock location. Only required if `address_id` is not provided.
 *     $ref: "#/components/schemas/StockLocationAddressInput"
 */
type CreateStockLocationInput struct {
	Name      string                     `json:"name"`
	AddressID *string                    `json:"address_id,omitempty"`
	Address   *StockLocationAddressInput `json:"address,omitempty"`
	Metadata  map[string]interface{}     `json:"metadata,omitempty"`
}

/**
 * @schema AdminPostStockLocationsLocationReq
 * type: object
 * description: "The details to update of the stock location."
 * properties:
 *   name:
 *     description: the name of the stock location
 *     type: string
 *   address_id:
 *     description: the stock location address ID
 *     type: string
 *   metadata:
 *     type: object
 *     description: An optional key-value map with additional details
 *     example: {car: "white"}
 *     externalDocs:
 *       description: "Learn about the metadata attribute, and how to delete and update it."
 *       url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
 *   address:
 *     description: The data of an associated address to create or update.
 *     $ref: "#/components/schemas/StockLocationAddressInput"
 */
type UpdateStockLocationInput struct {
	Name      *string                    `json:"name,omitempty"`
	AddressID *string                    `json:"address_id,omitempty"`
	Address   *StockLocationAddressInput `json:"address,omitempty"`
	Metadata  map[string]interface{}     `json:"metadata,omitempty"`
}

type IStockLocationService interface {
	List(context context.Context, selector FilterableStockLocation, config *sql.Options) ([]StockLocationDTO, *utils.ApplictaionError)
	ListAndCount(context context.Context, selector FilterableStockLocation, config *sql.Options) ([]StockLocationDTO, *int64, *utils.ApplictaionError)
	Retrieve(context context.Context, id uuid.UUID, config *sql.Options) (StockLocationDTO, *utils.ApplictaionError)
	Create(context context.Context, input CreateStockLocationInput) (StockLocationDTO, *utils.ApplictaionError)
	Update(context context.Context, id uuid.UUID, input UpdateStockLocationInput) (StockLocationDTO, *utils.ApplictaionError)
	Delete(context context.Context, id uuid.UUID) *utils.ApplictaionError
}

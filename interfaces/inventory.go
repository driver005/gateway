package interfaces

import (
	"context"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type InventoryItemDTO struct {
	core.Model
	SKU              string
	OriginCountry    string
	HsCode           string
	RequiresShipping bool
	MidCode          string
	Material         string
	Weight           float64
	Length           float64
	Height           float64
	Width            float64
	Title            string
	Description      string
	Thumbnail        string
}

type ReservationItemDTO struct {
	core.Model
	LocationId      uuid.UUID
	InventoryItemId uuid.UUID
	Quantity        int
	LineItemId      uuid.UUID
	Description     string
	CreatedBy       string
}

type InventoryLevelDTO struct {
	core.Model
	InventoryItemId  uuid.UUID
	LocationId       uuid.UUID
	StockedQuantity  int
	ReservedQuantity int
	IncomingQuantity int
}

type FilterableReservationItemProps struct {
	ID              string
	Type            string
	LineItemId      uuid.UUID
	InventoryItemId uuid.UUID
	LocationId      uuid.UUID
	Description     string
	CreatedBy       string
	Quantity        int
}

type FilterableInventoryItemProps struct {
	ID               string
	LocationId       uuid.UUID
	Q                string
	SKU              string
	OriginCountry    string
	HsCode           string
	RequiresShipping bool
}

type CreateInventoryItemInput struct {
	SKU              string
	OriginCountry    string
	MidCode          string
	Material         string
	Weight           float64
	Length           float64
	Height           float64
	Width            float64
	Title            string
	Description      string
	Thumbnail        string
	Metadata         map[string]interface{}
	HsCode           string
	RequiresShipping bool
}

type CreateReservationItemInput struct {
	LineItemId      uuid.UUID
	InventoryItemId uuid.UUID
	LocationId      uuid.UUID
	Quantity        int
	Description     string
	CreatedBy       string
	ExternalId      string
	Metadata        map[string]interface{}
}

type FilterableInventoryLevelProps struct {
	InventoryItemId  uuid.UUIDs
	LocationId       uuid.UUIDs
	StockedQuantity  int
	ReservedQuantity int
	IncomingQuantity int
}

type CreateInventoryLevelInput struct {
	InventoryItemIDd uuid.UUID
	LocationId       uuid.UUID
	StockedQuantity  int
	ReservedQuantity int
	IncomingQuantity int
}

type UpdateInventoryLevelInput struct {
	StockedQuantity  int
	IncomingQuantity int
}

type BulkUpdateInventoryLevelInput struct {
	InventoryItemIDd uuid.UUID
	LocationId       uuid.UUID
	StockedQuantity  int
	IncomingQuantity int
}

type UpdateReservationItemInput struct {
	Quantity    int
	LocationID  string
	Description string
	Metadata    map[string]interface{}
}

type ReserveQuantityContext struct {
	LocationID     string
	LineItemID     string
	SalesChannelID string
}

type ModuleJoinerConfig struct {
	JoinerID   string
	JoinerType string
}

type SharedContext struct {
	ContextID   string
	ContextType string
}

type IInventoryService interface {
	ListInventoryItems(context context.Context, selector FilterableInventoryItemProps, config sql.Options) ([]InventoryItemDTO, *int64, *utils.ApplictaionError)
	ListReservationItems(context context.Context, selector FilterableReservationItemProps, config sql.Options) ([]ReservationItemDTO, *int64, *utils.ApplictaionError)
	ListInventoryLevels(context context.Context, selector FilterableInventoryLevelProps, config sql.Options) ([]InventoryLevelDTO, *int64, *utils.ApplictaionError)
	RetrieveInventoryItem(context context.Context, inventoryItemId uuid.UUID, config sql.Options) (InventoryItemDTO, *utils.ApplictaionError)
	RetrieveInventoryLevel(context context.Context, inventoryItemId uuid.UUID, locationId string) (InventoryLevelDTO, *utils.ApplictaionError)
	RetrieveReservationItem(context context.Context, reservationId uuid.UUID) (ReservationItemDTO, *utils.ApplictaionError)
	CreateReservationItem(context context.Context, input CreateReservationItemInput) (ReservationItemDTO, *utils.ApplictaionError)
	CreateReservationItems(context context.Context, input []CreateReservationItemInput) ([]ReservationItemDTO, *utils.ApplictaionError)
	CreateInventoryItem(context context.Context, input CreateInventoryItemInput) (InventoryItemDTO, *utils.ApplictaionError)
	CreateInventoryItems(context context.Context, input []CreateInventoryItemInput) ([]InventoryItemDTO, *utils.ApplictaionError)
	CreateInventoryLevel(context context.Context, data CreateInventoryLevelInput) (InventoryLevelDTO, *utils.ApplictaionError)
	CreateInventoryLevels(context context.Context, data []CreateInventoryLevelInput) ([]InventoryLevelDTO, *utils.ApplictaionError)
	UpdateInventoryLevels(context context.Context, updates []BulkUpdateInventoryLevelInput) ([]InventoryLevelDTO, *utils.ApplictaionError)
	UpdateInventoryLevel(context context.Context, inventoryItemId uuid.UUID, locationId string, update UpdateInventoryLevelInput) (InventoryLevelDTO, *utils.ApplictaionError)
	UpdateInventoryItem(context context.Context, inventoryItemId uuid.UUID, input CreateInventoryItemInput) (InventoryItemDTO, *utils.ApplictaionError)
	UpdateReservationItem(context context.Context, reservationItemId uuid.UUID, input UpdateReservationItemInput) (ReservationItemDTO, *utils.ApplictaionError)
	DeleteReservationItemsByLineItem(context context.Context, lineItemId uuid.UUIDs) *utils.ApplictaionError
	DeleteReservationItem(context context.Context, reservationItemId uuid.UUIDs) *utils.ApplictaionError
	DeleteInventoryItem(context context.Context, inventoryItemId uuid.UUID) *utils.ApplictaionError
	RestoreInventoryItem(context context.Context, inventoryItemId uuid.UUID) *utils.ApplictaionError
	DeleteInventoryItemLevelByLocationID(context context.Context, locationId string) *utils.ApplictaionError
	DeleteReservationItemByLocationID(context context.Context, locationId string) *utils.ApplictaionError
	DeleteInventoryLevel(context context.Context, inventoryItemId uuid.UUID, locationId string) *utils.ApplictaionError
	AdjustInventory(context context.Context, inventoryItemId uuid.UUID, locationId string, adjustment int) (InventoryLevelDTO, *utils.ApplictaionError)
	ConfirmInventory(context context.Context, inventoryItemId uuid.UUID, locationIds uuid.UUIDs, quantity int) (bool, *utils.ApplictaionError)
	RetrieveAvailableQuantity(context context.Context, inventoryItemId uuid.UUID, locationIds uuid.UUIDs) (int, *utils.ApplictaionError)
	RetrieveStockedQuantity(context context.Context, inventoryItemId uuid.UUID, locationIds uuid.UUIDs) (int, *utils.ApplictaionError)
	RetrieveReservedQuantity(context context.Context, inventoryItemId uuid.UUID, locationIds uuid.UUIDs) (int, *utils.ApplictaionError)
}

package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
)

type ShippingRequirement struct {
	Type   models.ShippingOptionRequirementType
	Amount int
	ID     string
}

type ShippingMethodUpdate struct {
	Data         interface{}
	Price        float64
	ReturnID     string
	SwapID       string
	OrderID      string
	ClaimOrderID string
}

type CreateShippingMethod struct {
	Data             interface{}
	ShippingOptionID string
	Price            float64
	ReturnID         string
	SwapID           string
	CartID           string
	OrderID          string
	DraftOrderID     string
	ClaimOrderID     string
}

type CreateShippingMethodDto struct {
	CreateShippingMethod
	Cart  *models.Cart
	Order *models.Order
}

type CreateShippingOptionInput struct {
	PriceType    *models.ShippingOptionPriceType
	Name         string
	RegionID     string
	ProfileID    string
	ProviderID   string
	Data         core.JSONB
	IncludesTax  bool
	Amount       float64
	AdminOnly    bool
	IsReturn     bool
	Metadata     core.JSONB
	Requirements []models.ShippingOptionRequirement
}

type CreateCustomShippingOptionInput struct {
	Price            float64
	ShippingOptionID string
	CartID           *string
	Metadata         core.JSONB
}

type UpdateShippingOptionInput struct {
	Metadata     core.JSONB
	PriceType    *models.ShippingOptionPriceType
	Amount       float64
	Name         string
	AdminOnly    bool
	IsReturn     bool
	Requirements []models.ShippingOptionRequirement
	RegionID     string
	ProviderID   string
	ProfileID    string
	Data         string
	IncludesTax  *bool
}

type ValidatePriceTypeAndAmountInput struct {
	Amount    float64
	PriceType *models.ShippingOptionPriceType
}

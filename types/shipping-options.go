package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type FilterableShippingOption struct {
	core.FilterModel
	RegionId  uuid.UUID `json:"region_id,omitempty" validate:"omitempty"`
	IsReturn  bool      `json:"is_return,omitempty" validate:"omitempty"`
	AdminOnly bool      `json:"admin_only,omitempty" validate:"omitempty"`

	ProfileId uuid.UUIDs `json:"profile_id,omitempty" validate:"omitempty"`
}

type ShippingRequirement struct {
	Type   models.ShippingOptionRequirementType `json:"type"`
	Amount int                                  `json:"amount"`
	Id     uuid.UUID                            `json:"id"`
}

type ShippingMethodUpdate struct {
	Data         core.JSONB `json:"data,omitempty" validate:"omitempty"`
	Price        float64    `json:"price,omitempty" validate:"omitempty"`
	ReturnId     uuid.UUID  `json:"return_id,omitempty" validate:"omitempty"`
	SwapId       uuid.UUID  `json:"swap_id,omitempty" validate:"omitempty"`
	OrderId      uuid.UUID  `json:"order_id,omitempty" validate:"omitempty"`
	ClaimOrderId uuid.UUID  `json:"claim_order_id,omitempty" validate:"omitempty"`
}

type CreateShippingMethod struct {
	Data             core.JSONB `json:"data,omitempty" validate:"omitempty"`
	ShippingOptionId uuid.UUID  `json:"shipping_option_id,omitempty" validate:"omitempty"`
	Price            float64    `json:"price,omitempty" validate:"omitempty"`
	ReturnId         uuid.UUID  `json:"return_id,omitempty" validate:"omitempty"`
	SwapId           uuid.UUID  `json:"swap_id,omitempty" validate:"omitempty"`
	CartId           uuid.UUID  `json:"cart_id,omitempty" validate:"omitempty"`
	OrderId          uuid.UUID  `json:"order_id,omitempty" validate:"omitempty"`
	DraftOrderId     uuid.UUID  `json:"draft_order_id,omitempty" validate:"omitempty"`
	ClaimOrderId     uuid.UUID  `json:"claim_order_id,omitempty" validate:"omitempty"`
}

type CreateShippingMethodDto struct {
	CreateShippingMethod
	Cart  *models.Cart  `json:"cart,omitempty" validate:"omitempty"`
	Order *models.Order `json:"order,omitempty" validate:"omitempty"`
}

type RequirementInput struct {
	Type   models.ShippingOptionRequirementType `json:"type,omitempty" validate:"omitempty"`
	Amount float64                              `json:"amount,omitempty" validate:"omitempty"`
}

type ValidateRequirementTypeInput struct {
	Id     uuid.UUID                            `json:"id,omitempty" validate:"omitempty"`
	Type   models.ShippingOptionRequirementType `json:"type"`
	Amount float64                              `json:"amount"`
}

type CreateShippingOptionInput struct {
	PriceType    models.ShippingOptionPriceType `json:"price_type"`
	Name         string                         `json:"name"`
	RegionId     uuid.UUID                      `json:"region_id"`
	ProfileId    uuid.UUID                      `json:"profile_id"`
	ProviderId   uuid.UUID                      `json:"provider_id"`
	Data         core.JSONB                     `json:"data"`
	IncludesTax  bool                           `json:"includes_tax,omitempty" validate:"omitempty"`
	Amount       float64                        `json:"amount,omitempty" validate:"omitempty"`
	AdminOnly    bool                           `json:"admin_only,omitempty" validate:"omitempty"`
	IsReturn     bool                           `json:"is_return,omitempty" validate:"omitempty"`
	Metadata     core.JSONB                     `json:"metadata,omitempty" validate:"omitempty"`
	Requirements []RequirementInput             `json:"requirements,omitempty" validate:"omitempty"`
}

type CreateCustomShippingOptionInput struct {
	Price            float64    `json:"price,omitempty" validate:"omitempty"`
	ShippingOptionId uuid.UUID  `json:"shipping_option_id,omitempty" validate:"omitempty"`
	CartId           uuid.UUID  `json:"cart_id,omitempty" validate:"omitempty"`
	Metadata         core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

type UpdateShippingOptionInput struct {
	Metadata     core.JSONB                     `json:"metadata,omitempty" validate:"omitempty"`
	PriceType    models.ShippingOptionPriceType `json:"price_type,omitempty" validate:"omitempty"`
	Amount       float64                        `json:"amount,omitempty" validate:"omitempty"`
	Name         string                         `json:"name,omitempty" validate:"omitempty"`
	AdminOnly    bool                           `json:"admin_only,omitempty" validate:"omitempty"`
	IsReturn     bool                           `json:"is_return,omitempty" validate:"omitempty"`
	Requirements []RequirementInput             `json:"requirements,omitempty" validate:"omitempty"`
	RegionId     uuid.UUID                      `json:"region_id,omitempty" validate:"omitempty"`
	ProviderId   uuid.UUID                      `json:"provider_id,omitempty" validate:"omitempty"`
	ProfileId    uuid.UUID                      `json:"profile_id,omitempty" validate:"omitempty"`
	Data         string                         `json:"data,omitempty" validate:"omitempty"`
	IncludesTax  bool                           `json:"includes_tax,omitempty" validate:"omitempty"`
}

type ValidatePriceTypeAndAmountInput struct {
	Amount    float64                        `json:"amount,omitempty" validate:"omitempty"`
	PriceType models.ShippingOptionPriceType `json:"price_type,omitempty" validate:"omitempty"`
}

type ShippingOptionParams struct {
	ProductIds []uuid.UUID `json:"product_ids,omitempty" validate:"omitempty"`
	RegionId   uuid.UUID   `json:"region_id,omitempty" validate:"omitempty"`
	IsReturn   string      `json:"is_return,omitempty" validate:"omitempty"`
}

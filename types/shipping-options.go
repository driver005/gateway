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

// @oas:schema:AdminPostShippingOptionsReq
// type: object
// description: "The details of the shipping option to create."
// required:
//   - name
//   - region_id
//   - provider_id
//   - data
//   - price_type
//
// properties:
//
//	name:
//	  description: "The name of the Shipping Option"
//	  type: string
//	region_id:
//	  description: "The ID of the Region in which the Shipping Option will be available."
//	  type: string
//	provider_id:
//	  description: "The ID of the Fulfillment Provider that handles the Shipping Option."
//	  type: string
//	profile_id:
//	  description: "The ID of the Shipping Profile to add the Shipping Option to."
//	  type: number
//	data:
//	  description: "The data needed for the Fulfillment Provider to handle shipping with this Shipping Option."
//	  type: object
//	price_type:
//	  description: >-
//	    The type of the Shipping Option price. `flat_rate` indicates fixed pricing, whereas `calculated` indicates that the price will be calculated each time by the fulfillment provider.
//	  type: string
//	  enum:
//	    - flat_rate
//	    - calculated
//	amount:
//	  description: >-
//	    The amount to charge for the Shipping Option. If the `price_type` is set to `calculated`, this amount will not actually be used.
//	  type: integer
//	requirements:
//	  description: "The requirements that must be satisfied for the Shipping Option to be available."
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - type
//	      - amount
//	    properties:
//	      type:
//	        description: The type of the requirement
//	        type: string
//	        enum:
//	          - max_subtotal
//	          - min_subtotal
//	      amount:
//	        description: The amount to compare with.
//	        type: integer
//	is_return:
//	  description: Whether the Shipping Option can be used for returns or during checkout.
//	  type: boolean
//	  default: false
//	admin_only:
//	  description: >-
//	    If set to `true`, the shipping option can only be used when creating draft orders.
//	  type: boolean
//	  default: false
//	metadata:
//	  description: An optional set of key-value pairs with additional information.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
//	includes_tax:
//	  description: "Tax included in prices of shipping option"
//	  x-featureFlag: "tax_inclusive_pricing"
//	  type: boolean
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

// @oas:schema:AdminPostShippingOptionsOptionReq
// type: object
// description: "The details to update of the shipping option."
// required:
//   - requirements
//
// properties:
//
//	name:
//	  description: "The name of the Shipping Option"
//	  type: string
//	amount:
//	  description: >-
//	    The amount to charge for the Shipping Option. If the `price_type` of the shipping option is `calculated`, this amount will not actually be used.
//	  type: integer
//	admin_only:
//	  description: >-
//	    If set to `true`, the shipping option can only be used when creating draft orders.
//	  type: boolean
//	metadata:
//	  description: "An optional set of key-value pairs with additional information."
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
//	requirements:
//	  description: "The requirements that must be satisfied for the Shipping Option to be available."
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - type
//	      - amount
//	    properties:
//	      id:
//	        description: The ID of an existing requirement. If an ID is passed, the existing requirement's details are updated. Otherwise, a new requirement is created.
//	        type: string
//	      type:
//	        description: The type of the requirement
//	        type: string
//	        enum:
//	          - max_subtotal
//	          - min_subtotal
//	      amount:
//	        description: The amount to compare with.
//	        type: integer
//	includes_tax:
//	  description: "Tax included in prices of shipping option"
//	  x-featureFlag: "tax_inclusive_pricing"
//	  type: boolean
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

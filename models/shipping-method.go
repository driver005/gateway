package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:ShippingMethod
// title: "Shipping Method"
// description: "A Shipping Method represents a way in which an Order or Return can be shipped. Shipping Methods are created from a Shipping Option, but may contain additional details that can be necessary for the Fulfillment Provider to handle the shipment. If the shipping method is created for a return, it may be associated with a claim or a swap that the return is part of."
// type: object
// required:
//   - cart_id
//   - claim_order_id
//   - data
//   - id
//   - order_id
//   - price
//   - return_id
//   - shipping_option_id
//   - swap_id
//
// properties:
//
//	id:
//	  description: The shipping method's ID
//	  type: string
//	  example: sm_01F0YET7DR2E7CYVSDHM593QG2
//	shipping_option_id:
//	  description: The ID of the Shipping Option that the Shipping Method is built from.
//	  type: string
//	  example: so_01G1G5V27GYX4QXNARRQCW1N8T
//	order_id:
//	  description: The ID of the order that the shipping method is used in.
//	  nullable: true
//	  type: string
//	  example: order_01G8TJSYT9M6AVS5N4EMNFS1EK
//	order:
//	  description: The details of the order that the shipping method is used in.
//	  x-expandable: "order"
//	  nullable: true
//	  $ref: "#/components/schemas/Order"
//	claim_order_id:
//	  description: The ID of the claim that the shipping method is used in.
//	  nullable: true
//	  type: string
//	  example: null
//	claim_order:
//	  description: The details of the claim that the shipping method is used in.
//	  x-expandable: "claim_order"
//	  nullable: true
//	  $ref: "#/components/schemas/ClaimOrder"
//	cart_id:
//	  description: The ID of the cart that the shipping method is used in.
//	  nullable: true
//	  type: string
//	  example: cart_01G8ZH853Y6TFXWPG5EYE81X63
//	cart:
//	  description: The details of the cart that the shipping method is used in.
//	  x-expandable: "cart"
//	  nullable: true
//	  $ref: "#/components/schemas/Cart"
//	swap_id:
//	  description: The ID of the swap that the shipping method is used in.
//	  nullable: true
//	  type: string
//	  example: null
//	swap:
//	  description: The details of the swap that the shipping method is used in.
//	  x-expandable: "swap"
//	  nullable: true
//	  $ref: "#/components/schemas/Swap"
//	return_id:
//	  description: The ID of the return that the shipping method is used in.
//	  nullable: true
//	  type: string
//	  example: null
//	return_order:
//	  description: The details of the return that the shipping method is used in.
//	  x-expandable: "return_order"
//	  nullable: true
//	  $ref: "#/components/schemas/Return"
//	shipping_option:
//	  description: The details of the shipping option the method was created from.
//	  x-expandable: "shipping_option"
//	  nullable: true
//	  $ref: "#/components/schemas/ShippingOption"
//	tax_lines:
//	  description: The details of the tax lines applied on the shipping method.
//	  type: array
//	  x-expandable: "tax_lines"
//	  items:
//	    $ref: "#/components/schemas/ShippingMethodTaxLine"
//	price:
//	  description: The amount to charge for the Shipping Method. The currency of the price is defined by the Region that the Order that the Shipping Method belongs to is a part of.
//	  type: integer
//	  example: 200
//	data:
//	  description: Additional data that the Fulfillment Provider needs to fulfill the shipment. This is used in combination with the Shipping Options data, and may contain information such as a drop point id.
//	  type: object
//	  example: {}
//	includes_tax:
//	  description: "Whether the shipping method price include tax"
//	  type: boolean
//	  x-featureFlag: "tax_inclusive_pricing"
//	  default: false
//	subtotal:
//	  description: The subtotal of the shipping
//	  type: integer
//	  example: 8000
//	total:
//	  description: The total amount of the shipping
//	  type: integer
//	  example: 8200
//	tax_total:
//	  description: The total of tax
//	  type: integer
//	  example: 0
type ShippingMethod struct {
	core.SoftDeletableModel

	ShippingOptionId uuid.NullUUID           `json:"shipping_option_id" gorm:"column:shipping_option_id"`
	ShippingOption   *ShippingOption         `json:"shipping_option" gorm:"foreignKey:ShippingOptionId"`
	OrderId          uuid.NullUUID           `json:"order_id" gorm:"column:order_id"`
	Order            *Order                  `json:"order" gorm:"foreignKey:OrderId"`
	ReturnId         uuid.NullUUID           `json:"return_id" gorm:"column:return_id"`
	ReturnOrder      *Return                 `json:"return_order" gorm:"foreignKey:ReturnId"`
	SwapId           uuid.NullUUID           `json:"swap_id" gorm:"column:swap_id"`
	Swap             *Swap                   `json:"swap" gorm:"foreignKey:SwapId"`
	CartId           uuid.NullUUID           `json:"cart_id" gorm:"column:cart_id"`
	Cart             *Cart                   `json:"cart" gorm:"foreignKey:CartId"`
	ClaimOrderId     uuid.NullUUID           `json:"claim_order_id" gorm:"column:claim_order_id"`
	ClaimOrder       *ClaimOrder             `json:"claim_order" gorm:"foreignKey:ClaimOrderId"`
	TaxLines         []ShippingMethodTaxLine `json:"tax_lines" gorm:"foreignKey:Id"`
	Price            float64                 `json:"price" gorm:"column:price"`
	Subtotal         float64                 `json:"subtotal" gorm:"column:subtotal"`
	TaxTotal         float64                 `json:"tax_total" gorm:"column:tax_total"`
	Total            float64                 `json:"total" gorm:"column:total"`
	OriginalTotal    float64                 `json:"original_total" gorm:"column:original_total"`
	OriginalTaxTotal float64                 `json:"original_tax_total" gorm:"column:original_tax_total"`
	Data             core.JSONB              `json:"data" gorm:"column:data"`
	IncludesTax      bool                    `json:"includes_tax" gorm:"column:includes_tax"`
}

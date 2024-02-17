package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:LineItem
// title: "Line Item"
// description: "Line Items are created when a product is added to a Cart. When Line Items are purchased they will get copied to the resulting order, swap, or claim, and can eventually be referenced in Fulfillments and Returns. Line items may also be used for order edits."
// type: object
// required:
//   - allow_discounts
//   - cart_id
//   - claim_order_id
//   - created_at
//   - description
//   - fulfilled_quantity
//   - has_shipping
//   - id
//   - is_giftcard
//   - is_return
//   - metadata
//   - order_edit_id
//   - order_id
//   - original_item_id
//   - quantity
//   - returned_quantity
//   - shipped_quantity
//   - should_merge
//   - swap_id
//   - thumbnail
//   - title
//   - unit_price
//   - updated_at
//   - variant_id
//
// properties:
//
//	id:
//	  description: The line item's ID
//	  type: string
//	  example: item_01G8ZC9GWT6B2GP5FSXRXNFNGN
//	cart_id:
//	  description: The ID of the cart that the line item may belongs to.
//	  nullable: true
//	  type: string
//	  example: cart_01G8ZH853Y6TFXWPG5EYE81X63
//	cart:
//	  description: The details of the cart that the line item may belongs to.
//	  x-expandable: "cart"
//	  nullable: true
//	  $ref: "#/components/schemas/Cart"
//	order_id:
//	  description: The ID of the order that the line item may belongs to.
//	  nullable: true
//	  type: string
//	  example: order_01G8TJSYT9M6AVS5N4EMNFS1EK
//	order:
//	  description: The details of the order that the line item may belongs to.
//	  x-expandable: "order"
//	  nullable: true
//	  $ref: "#/components/schemas/Order"
//	swap_id:
//	  description: The ID of the swap that the line item may belong to.
//	  nullable: true
//	  type: string
//	  example: null
//	swap:
//	  description: The details of the swap that the line item may belong to.
//	  x-expandable: "swap"
//	  nullable: true
//	  $ref: "#/components/schemas/Swap"
//	claim_order_id:
//	  description: The ID of the claim that the line item may belong to.
//	  nullable: true
//	  type: string
//	  example: null
//	claim_order:
//	  description: The details of the claim that the line item may belong to.
//	  x-expandable: "claim_order"
//	  nullable: true
//	  $ref: "#/components/schemas/ClaimOrder"
//	tax_lines:
//	  description: The details of the item's tax lines.
//	  x-expandable: "tax_lines"
//	  type: array
//	  items:
//	    $ref: "#/components/schemas/LineItemTaxLine"
//	adjustments:
//	  description: The details of the item's adjustments, which are available when a discount is applied on the item.
//	  x-expandable: "adjustments"
//	  type: array
//	  items:
//	    $ref: "#/components/schemas/LineItemAdjustment"
//	original_item_id:
//	  description: The ID of the original line item. This is useful if the line item belongs to a resource that references an order, such as a return or an order edit.
//	  nullable: true
//	  type: string
//	order_edit_id:
//	  description: The ID of the order edit that the item may belong to.
//	  nullable: true
//	  type: string
//	order_edit:
//	  description: The details of the order edit.
//	  x-expandable: "order_edit"
//	  nullable: true
//	  $ref: "#/components/schemas/OrderEdit"
//	title:
//	  description: The title of the Line Item.
//	  type: string
//	  example: Medusa Coffee Mug
//	description:
//	  description: A more detailed description of the contents of the Line Item.
//	  nullable: true
//	  type: string
//	  example: One Size
//	thumbnail:
//	  description: A URL string to a small image of the contents of the Line Item.
//	  nullable: true
//	  type: string
//	  format: uri
//	  example: https://medusa-public-images.s3.eu-west-1.amazonaws.com/coffee-mug.png
//	is_return:
//	  description: Is the item being returned
//	  type: boolean
//	  default: false
//	is_giftcard:
//	  description: Flag to indicate if the Line Item is a Gift Card.
//	  type: boolean
//	  default: false
//	should_merge:
//	  description: Flag to indicate if new Line Items with the same variant should be merged or added as an additional Line Item.
//	  type: boolean
//	  default: true
//	allow_discounts:
//	  description: Flag to indicate if the Line Item should be included when doing discount calculations.
//	  type: boolean
//	  default: true
//	has_shipping:
//	  description: Flag to indicate if the Line Item has fulfillment associated with it.
//	  nullable: true
//	  type: boolean
//	  example: false
//	unit_price:
//	  description: The price of one unit of the content in the Line Item. This should be in the currency defined by the Cart/Order/Swap/Claim that the Line Item belongs to.
//	  type: integer
//	  example: 8000
//	variant_id:
//	  description: The id of the Product Variant contained in the Line Item.
//	  nullable: true
//	  type: string
//	  example: variant_01G1G5V2MRX2V3PVSR2WXYPFB6
//	variant:
//	  description: The details of the product variant that this item was created from.
//	  x-expandable: "variant"
//	  nullable: true
//	  $ref: "#/components/schemas/ProductVariant"
//	quantity:
//	  description: The quantity of the content in the Line Item.
//	  type: integer
//	  example: 1
//	fulfilled_quantity:
//	  description: The quantity of the Line Item that has been fulfilled.
//	  nullable: true
//	  type: integer
//	  example: 0
//	returned_quantity:
//	  description: The quantity of the Line Item that has been returned.
//	  nullable: true
//	  type: integer
//	  example: 0
//	shipped_quantity:
//	  description: The quantity of the Line Item that has been shipped.
//	  nullable: true
//	  type: integer
//	  example: 0
//	refundable:
//	  description: The amount that can be refunded from the given Line Item. Takes taxes and discounts into consideration.
//	  type: integer
//	  example: 0
//	subtotal:
//	  description: The subtotal of the line item
//	  type: integer
//	  example: 8000
//	tax_total:
//	  description: The total of tax of the line item
//	  type: integer
//	  example: 0
//	total:
//	  description: The total amount of the line item
//	  type: integer
//	  example: 8000
//	original_total:
//	  description: The original total amount of the line item
//	  type: integer
//	  example: 8000
//	original_tax_total:
//	  description: The original tax total amount of the line item
//	  type: integer
//	  example: 0
//	discount_total:
//	  description: The total of discount of the line item rounded
//	  type: integer
//	  example: 0
//	raw_discount_total:
//	  description: The total of discount of the line item
//	  type: integer
//	  example: 0
//	gift_card_total:
//	  description: The total of the gift card of the line item
//	  type: integer
//	  example: 0
//	includes_tax:
//	  description: "Indicates if the line item unit_price include tax"
//	  x-featureFlag: "tax_inclusive_pricing"
//	  type: boolean
//	  default: false
//	created_at:
//	  description: The date with timezone at which the resource was created.
//	  type: string
//	  format: date-time
//	updated_at:
//	  description: The date with timezone at which the resource was updated.
//	  type: string
//	  format: date-time
//	metadata:
//	  description: An optional key-value map with additional details
//	  nullable: true
//	  type: object
//	  example: {car: "white"}
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type LineItem struct {
	core.Model

	CartId            uuid.NullUUID        `json:"cart_id" gorm:"default:null"`
	Cart              *Cart                `json:"cart" gorm:"foreignKey:id;references:cart_id"`
	OrderId           uuid.NullUUID        `json:"order_id" gorm:"default:null"`
	Order             *Order               `json:"order" gorm:"foreignKey:id;references:order_id"`
	SwapId            uuid.NullUUID        `json:"swap_id" gorm:"default:null"`
	Swap              *Swap                `json:"swap" gorm:"foreignKey:id;references:swap_id"`
	ClaimOrderId      uuid.NullUUID        `json:"claim_order_id" gorm:"default:null"`
	ClaimOrder        *ClaimOrder          `json:"claim_order" gorm:"foreignKey:id;references:claim_order_id"`
	TaxLines          []LineItemTaxLine    `json:"tax_lines" gorm:"foreignKey:id"`
	Adjustments       []LineItemAdjustment `json:"adjustments" gorm:"foreignKey:id"`
	OriginalItemId    uuid.NullUUID        `json:"original_item_id" gorm:"default:null"`
	OrderEditId       uuid.NullUUID        `json:"order_edit_id" gorm:"default:null"`
	OrderEdit         *OrderEdit           `json:"order_edit" gorm:"foreignKey:id;references:order_edit_id"`
	Title             string               `json:"title"`
	Description       string               `json:"description" gorm:"default:null"`
	Thumbnail         string               `json:"thumbnail" gorm:"default:null"`
	IsReturn          bool                 `json:"is_return" gorm:"default:null"`
	IsGiftcard        bool                 `json:"is_giftcard" gorm:"default:null"`
	ShouldMerge       bool                 `json:"should_merge" gorm:"default:null"`
	AllowDiscounts    bool                 `json:"allow_discounts" gorm:"default:null"`
	HasShipping       bool                 `json:"has_shipping" gorm:"default:null"`
	UnitPrice         float64              `json:"unit_price"`
	ProductId         uuid.NullUUID        `json:"product_id" gorm:"default:null"`
	VariantId         uuid.NullUUID        `json:"variant_id" gorm:"default:null"`
	Variant           *ProductVariant      `json:"variant" gorm:"foreignKey:id;references:variant_id"`
	Quantity          int                  `json:"quantity"`
	FulfilledQuantity int                  `json:"fulfilled_quantity" gorm:"default:null"`
	ReturnedQuantity  int                  `json:"returned_quantity" gorm:"default:null"`
	ShippedQuantity   int                  `json:"shipped_quantity" gorm:"default:null"`
	Refundable        float64              `json:"refundable" gorm:"default:null"`
	Subtotal          float64              `json:"subtotal" gorm:"default:null"`
	TaxTotal          float64              `json:"tax_total" gorm:"default:null"`
	Total             float64              `json:"total" gorm:"default:null"`
	OriginalTotal     float64              `json:"original_total" gorm:"default:null"`
	OriginalTaxTotal  float64              `json:"original_tax_total" gorm:"default:null"`
	DiscountTotal     float64              `json:"discount_total" gorm:"default:null"`
	GiftCardTotal     float64              `json:"gift_card_total" gorm:"default:null"`
	RawDiscountTotal  float64              `json:"raw_discount_total" gorm:"default:null"`
	IncludesTax       bool                 `json:"includes_tax" gorm:"default:null"`
}

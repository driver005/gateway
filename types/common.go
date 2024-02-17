package types

type TotalField string

const (
	TotalFieldShippingTotal    TotalField = "shipping_total"
	TotalFieldDiscountTotal    TotalField = "discount_total"
	TotalFieldTaxTotal         TotalField = "tax_total"
	TotalFieldRefundedTotal    TotalField = "refunded_total"
	TotalFieldTotal            TotalField = "total"
	TotalFieldSubtotal         TotalField = "subtotal"
	TotalFieldRefundableAmount TotalField = "refundable_amount"
	TotalFieldGiftCardTotal    TotalField = "gift_card_total"
	TotalFieldGiftCardTaxTotal TotalField = "gift_card_tax_total"
)

// FindParams is used to configure how data is retrieved.
type FindParams struct {
	Expand string `json:"expand,omitempty" validate:"omitempty"`
	Fields string `json:"fields,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostProductsProductMetadataReq
// type: object
// required:
//   - key
//   - value
//
// properties:
//
//	key:
//	  description: The metadata key
//	  type: string
//	value:
//	  description: The metadata value
//	  type: string
type Metadata struct {
	Key   string `json:"key,omitempty" validate:"omitempty"`
	Value string `json:"value,omitempty" validate:"omitempty"`
}

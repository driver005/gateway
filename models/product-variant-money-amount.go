package models

import "github.com/driver005/gateway/core"

type ProductVariantMoneyAmount struct {
	core.Model

	MoneyAmountID string `gorm:"index:idx_product_variant_money_amount_money_amount_id_unique;unique"`

	VariantID string `gorm:"index:idx_product_variant_money_amount_variant_id"`
}

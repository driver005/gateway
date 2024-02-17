package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type ProductVariantMoneyAmount struct {
	core.Model

	MoneyAmountId uuid.NullUUID `gorm:"index:idx_product_variant_money_amount_money_amount_id_unique;unique"`
	VariantId     uuid.NullUUID `gorm:"index:idx_product_variant_money_amount_variant_id"`
}

package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type MoneyAmount struct {
	core.BaseModel

	CurrencyCode        string               `gorm:"column:currency_code;type:text;index:IDX_money_amount_currency_code,priority:1" json:"currency_code"`
	Amount              float64              `gorm:"column:amount;type:numeric" json:"amount"`
	MinQuantity         float64              `gorm:"column:min_quantity;type:numeric" json:"min_quantity"`
	MaxQuantity         float64              `gorm:"column:max_quantity;type:numeric" json:"max_quantity"`
	PriceSetMoneyAmount *PriceSetMoneyAmount `gorm:"foreignKey:MoneyAmountId" json:"price_set_money_amount"`
	DeletedAt           gorm.DeletedAt       `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_money_amount_deleted_at,priority:1" json:"deleted_at"`
}

func (*MoneyAmount) TableName() string {
	return "money_amount"
}

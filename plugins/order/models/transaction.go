package models

import "github.com/driver005/gateway/core"

type Transaction struct {
	core.BaseModel

	OrderId      string  `gorm:"column:order_id;type:text;not null;index:IDX_order_transaction_order_id,priority:1" json:"order_id"`
	Order        *Order  `gorm:"foreignkey:OrderId" json:"order"`
	Amount       float64 `gorm:"column:amount;type:numeric;not null" json:"amount"`
	RawAmount    string  `gorm:"column:raw_amount;type:jsonb;not null" json:"raw_amount"`
	CurrencyCode string  `gorm:"column:currency_code;type:text;not null;index:IDX_order_transaction_currency_code,priority:1" json:"currency_code"`
	Reference    string  `gorm:"column:reference;type:text" json:"reference"`
	ReferenceId  string  `gorm:"column:reference_id;type:text;index:IDX_order_transaction_reference_id,priority:1" json:"reference_id"`
}

func (*Transaction) TableName() string {
	return "order_transaction"
}

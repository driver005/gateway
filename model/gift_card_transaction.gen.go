// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGiftCardTransaction = "gift_card_transaction"

// GiftCardTransaction mapped from table <gift_card_transaction>
type GiftCardTransaction struct {
	ID         string    `gorm:"column:id;type:character varying;primaryKey" json:"id"`
	GiftCardID string    `gorm:"column:gift_card_id;type:character varying;not null;uniqueIndex:gcuniq,priority:1" json:"gift_card_id"`
	OrderID    string    `gorm:"column:order_id;type:character varying;not null;uniqueIndex:gcuniq,priority:2;index:IDX_d7d441b81012f87d4265fa57d2,priority:1" json:"order_id"`
	Amount     int32     `gorm:"column:amount;type:integer;not null" json:"amount"`
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp with time zone;not null;default:now()" json:"created_at"`
	IsTaxable  bool      `gorm:"column:is_taxable;type:boolean" json:"is_taxable"`
	TaxRate    float32   `gorm:"column:tax_rate;type:real" json:"tax_rate"`
}

// TableName GiftCardTransaction's table name
func (*GiftCardTransaction) TableName() string {
	return TableNameGiftCardTransaction
}

// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameStoreCurrency = "store_currencies"

// StoreCurrency mapped from table <store_currencies>
type StoreCurrency struct {
	StoreID      string `gorm:"column:store_id;type:character varying;primaryKey;index:IDX_b4f4b63d1736689b7008980394,priority:1" json:"store_id"`
	CurrencyCode string `gorm:"column:currency_code;type:character varying;primaryKey;index:IDX_82a6bbb0b527c20a0002ddcbd6,priority:1" json:"currency_code"`
}

// TableName StoreCurrency's table name
func (*StoreCurrency) TableName() string {
	return TableNameStoreCurrency
}

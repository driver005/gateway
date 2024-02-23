package models

import "github.com/google/uuid"

type StoreCurrency struct {
	StoreId      uuid.NullUUID `gorm:"column:store_id;type:character varying;primaryKey;index:IDX_b4f4b63d1736689b7008980394,priority:1" json:"store_id"`
	CurrencyCode string        `gorm:"column:currency_code;type:character varying;primaryKey;index:IDX_82a6bbb0b527c20a0002ddcbd6,priority:1" json:"currency_code"`
}

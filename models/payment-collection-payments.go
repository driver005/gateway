package models

import "github.com/google/uuid"

type PaymentCollectionPayment struct {
	PaymentCollectionId uuid.NullUUID `gorm:"column:payment_collection_id;type:character varying;primaryKey;index:IDX_payment_collection_payments_payment_collection_id,priority:1" json:"payment_collection_id"`
	PaymentId           uuid.NullUUID `gorm:"column:payment_id;type:character varying;primaryKey;index:IDX_payment_collection_payments_payment_id,priority:1" json:"payment_id"`
}

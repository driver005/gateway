package models

import "github.com/google/uuid"

type PaymentCollectionSession struct {
	PaymentCollectionId uuid.NullUUID `gorm:"column:payment_collection_id;type:character varying;primaryKey;index:IDX_payment_collection_sessions_payment_collection_id,priority:1" json:"payment_collection_id"`
	PaymentSessionId    uuid.NullUUID `gorm:"column:payment_session_id;type:character varying;primaryKey;index:IDX_payment_collection_sessions_payment_session_id,priority:1" json:"payment_session_id"`
}

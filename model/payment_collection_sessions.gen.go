// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePaymentCollectionSession = "payment_collection_sessions"

// PaymentCollectionSession mapped from table <payment_collection_sessions>
type PaymentCollectionSession struct {
	PaymentCollectionID string `gorm:"column:payment_collection_id;type:character varying;primaryKey;index:IDX_payment_collection_sessions_payment_collection_id,priority:1" json:"payment_collection_id"`
	PaymentSessionID    string `gorm:"column:payment_session_id;type:character varying;primaryKey;index:IDX_payment_collection_sessions_payment_session_id,priority:1" json:"payment_session_id"`
}

// TableName PaymentCollectionSession's table name
func (*PaymentCollectionSession) TableName() string {
	return TableNamePaymentCollectionSession
}
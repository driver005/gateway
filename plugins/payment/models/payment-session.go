package models

import (
	"time"

	"github.com/driver005/gateway/core"
)

type PaymentSession struct {
	core.BaseModel

	CurrencyCode        string    `gorm:"column:currency_code;type:text;not null" json:"currency_code"`
	Amount              float64   `gorm:"column:amount;type:numeric;not null" json:"amount"`
	ProviderId          string    `gorm:"column:provider_id;type:text;not null" json:"provider_id"`
	Data                string    `gorm:"column:data;type:jsonb;not null" json:"data"`
	Status              string    `gorm:"column:status;type:text;not null;default:pending" json:"status"`
	AuthorizedAt        time.Time `gorm:"column:authorized_at;type:timestamp with time zone" json:"authorized_at"`
	PaymentCollectionId string    `gorm:"column:payment_collection_id;type:text;not null;index:IDX_payment_session_payment_collection_id,priority:1" json:"payment_collection_id"`
	Payment             *Payment  `gorm:"foreignKey:PaymentSessionId;references:ID;constraint:OnDelete:CASCADE;"`
}

func (*PaymentSession) TableName() string {
	return "payment_session"
}

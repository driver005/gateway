package models

import (
	"time"

	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type Payment struct {
	core.BaseModel

	Amount              float64        `gorm:"column:amount;type:numeric;not null" json:"amount"`
	AuthorizedAmount    float64        `gorm:"column:authorized_amount;type:numeric" json:"authorized_amount"`
	CurrencyCode        string         `gorm:"column:currency_code;type:text;not null" json:"currency_code"`
	ProviderId          string         `gorm:"column:provider_id;type:text;not null" json:"provider_id"`
	CartId              string         `gorm:"column:cart_id;type:text" json:"cart_id"`
	OrderId             string         `gorm:"column:order_id;type:text" json:"order_id"`
	OrderEditId         string         `gorm:"column:order_edit_id;type:text" json:"order_edit_id"`
	CustomerId          string         `gorm:"column:customer_id;type:text" json:"customer_id"`
	Data                string         `gorm:"column:data;type:jsonb" json:"data"`
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_payment_deleted_at,priority:1" json:"deleted_at"`
	CapturedAt          time.Time      `gorm:"column:captured_at;type:timestamp with time zone" json:"captured_at"`
	CanceledAt          time.Time      `gorm:"column:canceled_at;type:timestamp with time zone" json:"canceled_at"`
	PaymentCollectionId string         `gorm:"column:payment_collection_id;type:text;not null;index:IDX_payment_payment_collection_id,priority:1" json:"payment_collection_id"`
	SessionId           string         `gorm:"column:session_id;type:text;not null;uniqueIndex:payment_session_id_unique,priority:1" json:"session_id"`
	PaymentSession      PaymentSession `gorm:"foreignKey:SessionId" json:"payment_session"`
	Refunds             []Refund       `gorm:"foreignKey:PaymentId" json:"refunds"`
	Captures            []Capture      `gorm:"foreignKey:PaymentId" json:"captures"`
}

func (*Payment) TableName() string {
	return "payment"
}

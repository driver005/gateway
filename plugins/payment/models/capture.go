package models

import "github.com/driver005/gateway/core"

type Capture struct {
	core.BaseModel

	Amount    float64  `gorm:"column:amount;type:numeric;not null" json:"amount"`
	PaymentId string   `gorm:"column:payment_id;type:text;not null;index:IDX_capture_payment_id,priority:1" json:"payment_id"`
	Payment   *Payment `gorm:"foreignKey:PaymentId" json:"payment"`
	CreatedBy string   `gorm:"column:created_by;type:text" json:"created_by"`
}

func (*Capture) TableName() string {
	return "capture"
}

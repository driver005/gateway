package models

import (
	"time"

	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type PaymentCollection struct {
	core.BaseModel

	CurrencyCode     string            `gorm:"column:currency_code;type:text;not null" json:"currency_code"`
	Amount           float64           `gorm:"column:amount;type:numeric;not null" json:"amount"`
	RegionId         string            `gorm:"column:region_id;type:text;not null;index:IDX_payment_collection_region_id,priority:1" json:"region_id"`
	DeletedAt        gorm.DeletedAt    `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_payment_collection_deleted_at,priority:1" json:"deleted_at"`
	CompletedAt      time.Time         `gorm:"column:completed_at;type:timestamp with time zone" json:"completed_at"`
	Status           string            `gorm:"column:status;type:text;not null;default:not_paid" json:"status"`
	PaymentProviders []PaymentProvider `gorm:"many2many:payment_collection_payment_provider;" json:"payment_providers"`
	PaymentSessions  []PaymentSession  `gorm:"foreignKey:PaymentCollectionId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"payment_sessions"`
	Payments         []Payment         `gorm:"foreignKey:PaymentCollectionId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"payments"`
}

func (*PaymentCollection) TableName() string {
	return "payment_collection"
}

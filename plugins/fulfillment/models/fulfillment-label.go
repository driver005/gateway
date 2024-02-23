package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type FulfillmentLabel struct {
	core.BaseModel

	TrackingNumber string         `gorm:"column:tracking_number;type:text;not null" json:"tracking_number"`
	TrackingURL    string         `gorm:"column:tracking_url;type:text;not null" json:"tracking_url"`
	LabelURL       string         `gorm:"column:label_url;type:text;not null" json:"label_url"`
	FulfillmentId  string         `gorm:"column:fulfillment_id;type:text;not null;index:IDX_fulfillment_label_fulfillment_id,priority:1" json:"fulfillment_id"`
	Fulfillment    *Fulfillment   `gorm:"foreignKey:FulfillmentId;references:ID;constraint:OnDelete:CASCADE" json:"fulfillment"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_fulfillment_label_deleted_at,priority:1" json:"deleted_at"`
}

func (*FulfillmentLabel) TableName() string {
	return "fulfillment_label"
}

package models

import "github.com/driver005/gateway/core"

type PaymentMethodToken struct {
	core.BaseModel

	ProviderId        string `gorm:"column:provider_id;type:text;not null" json:"provider_id"`
	Data              string `gorm:"column:data;type:jsonb" json:"data"`
	Name              string `gorm:"column:name;type:text;not null" json:"name"`
	TypeDetail        string `gorm:"column:type_detail;type:text" json:"type_detail"`
	DescriptionDetail string `gorm:"column:description_detail;type:text" json:"description_detail"`
}

func (*PaymentMethodToken) TableName() string {
	return "payment_method_token"
}

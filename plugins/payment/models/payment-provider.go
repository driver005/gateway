package models

type PaymentProvider struct {
	Id        string `gorm:"column:id;type:text;primaryKey" json:"id"`
	IsEnabled bool   `gorm:"column:is_enabled;type:boolean;not null;default:true" json:"is_enabled"`
}

func (*PaymentProvider) TableName() string {
	return "payment_provider"
}

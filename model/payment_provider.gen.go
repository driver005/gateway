// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePaymentProvider = "payment_provider"

// PaymentProvider mapped from table <payment_provider>
type PaymentProvider struct {
	ID          string `gorm:"column:id;type:character varying;primaryKey" json:"id"`
	IsInstalled bool   `gorm:"column:is_installed;type:boolean;not null;default:true" json:"is_installed"`
}

// TableName PaymentProvider's table name
func (*PaymentProvider) TableName() string {
	return TableNamePaymentProvider
}
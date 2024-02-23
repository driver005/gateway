package models

import "time"

type TaxRateRule struct {
	TaxRateId   string    `gorm:"column:tax_rate_id;type:text;primaryKey;index:IDX_tax_rate_rule_tax_rate_id,priority:1" json:"tax_rate_id"`
	ReferenceId string    `gorm:"column:reference_id;type:text;primaryKey" json:"reference_id"`
	Reference   string    `gorm:"column:reference;type:text;not null" json:"reference"`
	Metadata    string    `gorm:"column:metadata;type:jsonb" json:"metadata"`
	TaxRate     *TaxRate  `gorm:"foreignKey:TaxRateId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"tax_rate"`
	CreatedBy   string    `gorm:"column:created_by;type:text" json:"created_by"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp with time zone;not null;default:now()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp with time zone;not null;default:now()" json:"updated_at"`
}

func (*TaxRateRule) TableName() string {
	return "tax_rate_rule"
}

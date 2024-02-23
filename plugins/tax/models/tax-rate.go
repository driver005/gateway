package models

import "github.com/driver005/gateway/core"

type TaxRate struct {
	core.BaseModel

	Rate         float32       `gorm:"column:rate;type:real" json:"rate"`
	Code         string        `gorm:"column:code;type:text" json:"code"`
	Name         string        `gorm:"column:name;type:text;not null" json:"name"`
	IsDefault    bool          `gorm:"column:is_default;type:boolean;not null" json:"is_default"`
	IsCombinable bool          `gorm:"column:is_combinable;type:boolean;not null" json:"is_combinable"`
	TaxRegionId  string        `gorm:"column:tax_region_id;type:text;not null;uniqueIndex:IDX_single_default_region,priority:1;index:IDX_tax_rate_tax_region_id,priority:1" json:"tax_region_id"`
	CreatedBy    string        `gorm:"column:created_by;type:text" json:"created_by"`
	TaxRegion    *TaxRegion    `gorm:"foreignKey:TaxRegionId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"tax_region"`
	Rules        []TaxRateRule `gorm:"foreignKey:TaxRateId" json:"rules"`
}

func (*TaxRate) TableName() string {
	return "tax_rate"
}

package models

import (
	"github.com/driver005/gateway/core"
)

type TaxRegion struct {
	core.BaseModel

	CountryCode  string     `gorm:"column:country_code;type:text;not null;uniqueIndex:IDX_tax_region_unique_country_province,priority:1" json:"country_code"`
	ProvinceCode string     `gorm:"column:province_code;type:text;uniqueIndex:IDX_tax_region_unique_country_province,priority:2" json:"province_code"`
	ParentId     string     `gorm:"column:parent_id;type:text;index:IDX_tax_region_parent_id,priority:1" json:"parent_id"`
	CreatedBy    string     `gorm:"column:created_by;type:text" json:"created_by"`
	Parent       *TaxRegion `gorm:"foreignKey:ParentId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"parent"`
	TaxRates     []TaxRate  `gorm:"foreignKey:TaxRegionId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"tax_rates"`
}

func (*TaxRegion) TableName() string {
	return "tax_region"
}

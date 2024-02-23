package models

import "github.com/driver005/gateway/core"

type Country struct {
	core.BaseModel

	//ID          string `gorm:"column:id;type:text;primaryKey" json:"id"`
	Iso2        string  `gorm:"column:iso_2;type:text;not null;uniqueIndex:IDX_region_country_region_id_iso_2_unique,priority:1" json:"iso_2"`
	Iso3        string  `gorm:"column:iso_3;type:text;not null" json:"iso_3"`
	NumCode     int32   `gorm:"column:num_code;type:integer;not null" json:"num_code"`
	Name        string  `gorm:"column:name;type:text;not null" json:"name"`
	DisplayName string  `gorm:"column:display_name;type:text;not null" json:"display_name"`
	RegionId    string  `gorm:"column:region_id;type:text;uniqueIndex:IDX_region_country_region_id_iso_2_unique,priority:2" json:"region_id"`
	Region      *Region `gorm:"foreignKey:RegionId" json:"region"`
}

func (*Country) TableName() string {
	return "region_country"
}

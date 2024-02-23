package models

import (
	"github.com/driver005/gateway/core"
)

type ProductCategory struct {
	core.BaseModel

	Name             string            `gorm:"column:name;type:text;not null" json:"name"`
	Description      string            `gorm:"column:description;type:text;not null" json:"description"`
	Handle           string            `gorm:"column:handle;type:text;not null;uniqueIndex:IDX_product_category_handle,priority:1" json:"handle"`
	Mpath            string            `gorm:"column:mpath;type:text;not null;index:IDX_product_category_path,priority:1" json:"mpath"`
	IsActive         bool              `gorm:"column:is_active;type:boolean;not null" json:"is_active"`
	IsInternal       bool              `gorm:"column:is_internal;type:boolean;not null" json:"is_internal"`
	Rank             float64           `gorm:"column:rank;type:numeric;not null" json:"rank"`
	ParentCategoryId string            `gorm:"column:parent_category_id;type:text" json:"parent_category_id"`
	ParentCategory   *ProductCategory  `gorm:"foreignKey:ParentCategoryId" json:"parent_category"`
	CategoryChildren []ProductCategory `gorm:"foreignKey:ParentCategoryId" json:"category_children"`
	Products         []Product         `gorm:"many2many:product_categories;" json:"products"`
}

func (*ProductCategory) TableName() string {
	return "product_category"
}

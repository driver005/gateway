package models

import "github.com/google/uuid"

type ProductCategoryProduct struct {
	ProductCategoryId uuid.NullUUID `gorm:"column:product_category_id;type:character varying;not null;uniqueIndex:IDX_upcp_product_id_product_category_id,priority:1;index:IDX_pcp_product_category_id,priority:1" json:"product_category_id"`
	ProductId         uuid.NullUUID `gorm:"column:product_id;type:character varying;not null;uniqueIndex:IDX_upcp_product_id_product_category_id,priority:2;index:IDX_pcp_product_id,priority:1" json:"product_id"`
}

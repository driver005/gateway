// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameProductCategoryProduct = "product_category_product"

// ProductCategoryProduct mapped from table <product_category_product>
type ProductCategoryProduct struct {
	ProductCategoryID string `gorm:"column:product_category_id;type:character varying;not null;uniqueIndex:IDX_upcp_product_id_product_category_id,priority:1;index:IDX_pcp_product_category_id,priority:1" json:"product_category_id"`
	ProductID         string `gorm:"column:product_id;type:character varying;not null;uniqueIndex:IDX_upcp_product_id_product_category_id,priority:2;index:IDX_pcp_product_id,priority:1" json:"product_id"`
}

// TableName ProductCategoryProduct's table name
func (*ProductCategoryProduct) TableName() string {
	return TableNameProductCategoryProduct
}
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameDiscountRegion = "discount_regions"

// DiscountRegion mapped from table <discount_regions>
type DiscountRegion struct {
	DiscountID string `gorm:"column:discount_id;type:character varying;primaryKey;index:IDX_f4194aa81073f3fab8aa86906f,priority:1" json:"discount_id"`
	RegionID   string `gorm:"column:region_id;type:character varying;primaryKey;index:IDX_a21a7ffbe420d492eb46c305fe,priority:1" json:"region_id"`
}

// TableName DiscountRegion's table name
func (*DiscountRegion) TableName() string {
	return TableNameDiscountRegion
}

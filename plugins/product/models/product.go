package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type Product struct {
	core.BaseModel
	Title         string            `gorm:"column:title;type:text;not null" json:"title"`
	Handle        string            `gorm:"column:handle;type:text;not null;uniqueIndex:IDX_product_handle_unique,priority:1" json:"handle"`
	Subtitle      string            `gorm:"column:subtitle;type:text" json:"subtitle"`
	Description   string            `gorm:"column:description;type:text" json:"description"`
	IsGiftcard    bool              `gorm:"column:is_giftcard;type:boolean;not null" json:"is_giftcard"`
	Status        string            `gorm:"column:status;type:text;not null" json:"status"`
	Thumbnail     string            `gorm:"column:thumbnail;type:text" json:"thumbnail"`
	Options       []ProductOption   `gorm:"foreignKey:ProductID" json:"options"`
	Variants      []ProductVariant  `gorm:"foreignKey:ProductID" json:"variants"`
	Weight        string            `gorm:"column:weight;type:text" json:"weight"`
	Length        string            `gorm:"column:length;type:text" json:"length"`
	Height        string            `gorm:"column:height;type:text" json:"height"`
	Width         string            `gorm:"column:width;type:text" json:"width"`
	OriginCountry string            `gorm:"column:origin_country;type:text" json:"origin_country"`
	HsCode        string            `gorm:"column:hs_code;type:text" json:"hs_code"`
	MidCode       string            `gorm:"column:mid_code;type:text" json:"mid_code"`
	Material      string            `gorm:"column:material;type:text" json:"material"`
	CollectionId  string            `gorm:"column:collection_id;type:text" json:"collection_id"`
	Collection    ProductCollection `gorm:"foreignKey:CollectionId" json:"collection"`
	TypeId        string            `gorm:"column:type_id;type:text;index:IDX_product_type_id,priority:1" json:"type_id"`
	Type          ProductType       `gorm:"foreignKey:TypeId" json:"type"`
	Tags          []ProductTag      `gorm:"many2many:product_tags;" json:"tags"`
	Images        []ProductImage    `gorm:"many2many:product_images;" json:"images"`
	Categories    []ProductCategory `gorm:"many2many:product_category_product;" json:"categories"`
	Discountable  bool              `gorm:"column:discountable;type:boolean;not null;default:true" json:"discountable"`
	ExternalId    string            `gorm:"column:external_id;type:text" json:"external_id"`
	DeletedAt     gorm.DeletedAt    `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_product_deleted_at,priority:1" json:"deleted_at"`
}

func (*Product) TableName() string {
	return "product"
}

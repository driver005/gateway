package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type ProductVariantInventoryItem struct {
	sql.Repository[models.ProductVariantInventoryItem]
}

func ProductVariantInventoryItemRepository(db *gorm.DB) *ProductVariantInventoryItem {
	return &ProductVariantInventoryItem{*sql.NewRepository[models.ProductVariantInventoryItem](db)}
}

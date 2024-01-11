package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepo struct {
	sql.Repository[models.Product]
}

func ProductRepository(db *gorm.DB) *ProductRepo {
	return &ProductRepo{*sql.NewRepository[models.Product](db)}
}

type CategoryQueryParams struct {
	Value uuid.UUIDs
}

func (r *ProductRepo) BulkAddToCollection(productIds uuid.UUIDs, collectionId uuid.UUID) ([]models.Product, *utils.ApplictaionError) {
	var products []models.Product
	if err := r.Db().Model(&models.Product{}).Where("id IN ?", productIds).Find(&products).Error; err != nil {
		return nil, r.HandleDBError(err)
	}
	// Update products to add collectionId
	for i := range products {
		products[i].CollectionId = uuid.NullUUID{UUID: collectionId}
	}

	if err := r.Db().Save(&products).Error; err != nil {
		return nil, r.HandleDBError(err)
	}

	return products, nil
}

func (r *ProductRepo) BulkRemoveFromCollection(productIds uuid.UUIDs, collectionId uuid.UUID) *utils.ApplictaionError {
	if err := r.Db().Model(&models.Product{}).Where("id IN ?", productIds).Update("collection_id", nil).Error; err != nil {
		return r.HandleDBError(err)
	}
	return nil
}
func (r *ProductRepo) GetFreeTextSearchResultsAndCount(q *string, options sql.Query, relations []string) ([]models.Product, *int64, *utils.ApplictaionError) {
	// option := options
	// product := "product"
	// prices := "prices"
	// variants := "variants"
	// collection := "collection"
	// tags := "tags"

	// if _, ok := option.Where["Description"]; ok {
	// 	delete(option.Where, "Description")
	// }
	// if _, ok := option.Where["Title"]; ok {
	// 	delete(option.Where, "Title")
	// }

	// tags := option.Where.Tags
	// delete(option.Where, "Tags")
	// priceLists := option.Where.PriceListID
	// delete(option.Where, "PriceListID")
	// salesChannels := option.Where.SalesChannelID
	// delete(option.Where, "SalesChannelID")
	// discountConditionID := option.Where.DiscountConditionID
	// delete(option.Where, "DiscountConditionID")
	// categoryId := option.Where.CategoryID
	// delete(option.Where, "CategoryID")
	// includeCategoryChildrenQ := option.Where.IncludeCategoryChildren
	// delete(option.Where, "IncludeCategoryChildren")
	// categoriesQueryQ := option.Where.Categories
	// delete(option.Where, "Categories")

	// db := r.Db().Model(&models.Product{}).
	// 	Joins("left join variants on variants.product_id = product.id").
	// 	Joins("left join collection on collection.id = product.collection_id").
	// 	Select("product.id").
	// 	Where(option.Where).
	// 	Where(
	// 		gorm.Expr("product.description ILIKE ? OR product.title ILIKE ? OR variants.title ILIKE ? OR variants.sku ILIKE ? OR collection.title ILIKE ?",
	// 			"%"+*q+"%", "%"+*q+"%", "%"+*q+"%", "%"+*q+"%", "%"+*q+"%"),
	// 	)

	// if discountConditionID != "" {
	// 	db = db.Joins("inner join discount_condition_product dc_product on dc_product.product_id = product.id and dc_product.condition_id = ?", discountConditionID)
	// }

	// if tags != "" {
	// 	db = db.Joins("left join product_tags on product_tags.product_id = product.id").Where("product_tags.tag_id IN ?", tags)
	// }

	// if priceLists != nil {
	// 	db = db.Joins("left join variant_prices on variant_prices.variant_id = variants.id").
	// 		Joins("left join prices on prices.id = variant_prices.price_id").
	// 		Where("prices.price_list_id IN ?", priceLists)
	// }

	// if salesChannels != nil {
	// 	db = db.Joins("inner join product_channels on product_channels.product_id = product.id").
	// 		Where("product_channels.channel_id IN ?", salesChannels)
	// }

	// if categoryId != "" || categoriesQueryQ != nil {
	// 	var categoryIDs []string
	// 	if categoryId != "" {
	// 		categoryIDs = append(categoryIDs, categoryId)
	// 		if includeCategoryChildren {
	// 			// get all children category ids
	// 		}
	// 	}
	// 	if categoriesQuery != nil {
	// 		// append category ids from query
	// 	}
	// 	db = db.Joins("inner join product_categories on product_categories.product_id = product.id").
	// 		Where("product_categories.category_id IN ?", categoryIDs)
	// }

	// if categoryId != "" || categoriesQuery != nil {
	// 	categoryIDs := r.GetCategoryIDsFromInput(categoryId, includeCategoryChildren)
	// 	if len(categoryIDs) > 0 {
	// 		db = db.Joins("INNER JOIN product_categories ON product_categories.product_id = product.id").
	// 			Where("product_categories.category_id IN ?", categoryIDs)
	// 	}
	// }

	// // Apply ordering

	// if option.WithDeleted {
	// 	db = db.Unscoped()
	// }

	// var products []models.Product
	// var count *int64
	// db.Offset(*option.Skip).Limit(*option.Take).Find(&products).Count(count)

	// return products, count, nil
	return nil, nil, nil
}

func (r *ProductRepo) GetCategoryIDsFromInput(categoryID CategoryQueryParams, includeCategoryChildren bool) uuid.UUIDs {
	categoryIDs := categoryID.Value
	if len(categoryIDs) == 0 {
		return nil
	}

	if includeCategoryChildren {
		var categories []models.ProductCategory
		r.Db().Where("id IN ?", categoryIDs).Find(&categories)
		for _, category := range categories {
			var children []models.ProductCategory
			r.Db().Model(&category).Association("Children").Find(&children)
			for _, child := range children {
				categoryIDs = append(categoryIDs, child.Id)
			}
		}
	}

	return categoryIDs
}

func GetCategoryIDsRecursively(productCategory *models.ProductCategory) []string {
	var result []string
	result = append(result, productCategory.Id.String())

	for _, child := range productCategory.CategoryChildren {
		childIDs := GetCategoryIDsRecursively(&child)
		result = append(result, childIDs...)
	}

	return result
}

func (r *ProductRepo) IsProductInSalesChannels(id uuid.UUID, salesChannelIDs uuid.UUIDs) bool {
	var count int64
	r.Db().Model(&models.Product{}).
		Joins("SalesChannels").
		Where("sales_channels.id IN ?", salesChannelIDs).
		Where("products.id = ?", id).
		Count(&count)
	return count > 0
}

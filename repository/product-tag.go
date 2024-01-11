package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductTagRepo struct {
	sql.Repository[models.ProductTag]
}

func ProductTagRepository(db *gorm.DB) *ProductTagRepo {
	return &ProductTagRepo{*sql.NewRepository[models.ProductTag](db)}
}

func (r *ProductTagRepo) InsertBulk(data []models.ProductTag) ([]models.ProductTag, *utils.ApplictaionError) {
	err := r.Db().Create(&data).Error
	if err != nil {
		return nil, r.HandleDBError(err)
	}
	return data, nil
}

func (r *ProductTagRepo) ListTagsByUsage(take int) ([]models.ProductTag, *utils.ApplictaionError) {
	var tags []models.ProductTag
	err := r.Db().Table("product_tags").
		Select("id, COUNT(pts.product_tag_id) as usage_count, value").
		Joins("LEFT JOIN product_tags pts ON product_tags.id = pts.product_tag_id").
		Group("id").
		Order("usage_count DESC").
		Limit(take).
		Scan(&tags).Error
	if err != nil {
		return nil, r.HandleDBError(err)
	}
	return tags, nil
}

func (r *ProductTagRepo) UpsertTags(tags []models.ProductTag) ([]models.ProductTag, *utils.ApplictaionError) {
	var tagValues []string
	for _, tag := range tags {
		tagValues = append(tagValues, tag.Value)
	}

	var existingTags []models.ProductTag
	err := r.Db().Where("value IN ?", tagValues).Find(&existingTags).Error
	if err != nil {
		return nil, r.HandleDBError(err)
	}

	existingTagsMap := make(map[string]models.ProductTag)
	for _, tag := range existingTags {
		existingTagsMap[tag.Value] = tag
	}

	var upsertedTags []models.ProductTag
	var tagsToCreate []models.ProductTag
	for _, tag := range tags {
		aTag, ok := existingTagsMap[tag.Value]
		if ok {
			upsertedTags = append(upsertedTags, aTag)
		} else {
			tagsToCreate = append(tagsToCreate, models.ProductTag{
				Value: tag.Value,
			})
		}
	}

	if len(tagsToCreate) > 0 {
		r.InsertBulk(tagsToCreate)
		upsertedTags = append(upsertedTags, tagsToCreate...)
	}

	return upsertedTags, nil
}

func (r *ProductTagRepo) FindAndCountByDiscountConditionID(conditionID uuid.UUID, query sql.Query) ([]models.ProductTag, *int64, *utils.ApplictaionError) {
	var tags []models.ProductTag
	var count *int64
	err := r.Db().Model(&models.ProductTag{}).
		Where(query.Where).
		Scopes(func(db *gorm.DB) *gorm.DB {
			return db.Set("gorm:query_option", query)
		}).
		Joins("INNER JOIN discount_condition_product_tag dc_pt ON dc_pt.product_tag_id = product_tags.id AND dc_pt.condition_id = ?", conditionID).
		Find(&tags).
		Count(count).Error
	if err != nil {
		return nil, nil, r.HandleDBError(err)
	}
	return tags, count, nil
}

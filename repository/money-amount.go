package repository

import (
	"strings"
	"time"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MoneyAmountRepo struct {
	sql.Repository[models.MoneyAmount]
}

func MoneyAmountRepository(db *gorm.DB) *MoneyAmountRepo {
	return &MoneyAmountRepo{*sql.NewRepository[models.MoneyAmount](db)}
}

func (r *MoneyAmountRepo) InsertBulk(data []models.MoneyAmount) ([]models.MoneyAmount, *utils.ApplictaionError) {
	if err := r.Db().Create(&data).Error; err != nil {
		return nil, r.HandleDBError(err)
	}

	var variantMoneyAmounts []models.ProductVariantMoneyAmount
	for _, d := range data {
		if d.Variant != nil {
			variantMoneyAmounts = append(variantMoneyAmounts, models.ProductVariantMoneyAmount{
				VariantId:     uuid.NullUUID{UUID: d.Variant.Id},
				MoneyAmountId: uuid.NullUUID{UUID: d.Id},
			})
		}
	}
	err := r.Db().Create(&variantMoneyAmounts).Error
	if err != nil {
		return nil, r.HandleDBError(err)
	}
	return data, nil
}

func (r *MoneyAmountRepo) FindVariantPricesNotIn(variantId uuid.UUID, prices []models.MoneyAmount) ([]models.MoneyAmount, *utils.ApplictaionError) {
	var pricesNotInPricesPayload []models.MoneyAmount
	queryBuilder := r.Db().Model(&models.MoneyAmount{}).
		Joins("LEFT JOIN product_variant_money_amount pvma ON pvma.money_amount_id = ma.id").
		Where("pvma.variant_id = ?", variantId).
		Where("price_list_id IS NULL").
		Where("currency_code NOT IN ?", prices).
		Or("region_id NOT IN ?", prices).
		Find(&pricesNotInPricesPayload)
	if err := queryBuilder.Error; err != nil {
		return nil, r.HandleDBError(err)
	}
	return pricesNotInPricesPayload, nil
}

func (r *MoneyAmountRepo) DeleteVariantPricesNotIn(variantId uuid.UUID, prices []models.MoneyAmount) *utils.ApplictaionError {
	var maIdsForVariant []struct {
		ID        string
		VariantID string
	}
	queryBuilder := r.Db().Model(&models.MoneyAmount{}).
		Joins("LEFT JOIN product_variant_money_amount pvma ON pvma.money_amount_id = ma.id").
		Select("pvma.variant_id, pvma.money_amount_id").
		Where("pvma.variant_id = ?", variantId).
		Find(&maIdsForVariant)
	if err := queryBuilder.Error; err != nil {
		return r.HandleDBError(err)
	}
	var where []interface{}
	for _, price := range prices {
		if price.CurrencyCode != "" {
			where = append(where, []interface{}{
				map[string]interface{}{"currency_code": price.CurrencyCode},
				map[string]interface{}{"region_id": price.RegionId},
			})
		}
		if price.RegionId.UUID != uuid.Nil {
			where = append(where, map[string]interface{}{"region_id": price.RegionId})
		}
	}
	queryBuilder = r.Db().Model(&models.MoneyAmount{}).
		Where("id IN ?", maIdsForVariant).
		Where("price_list_id IS NULL").
		Where(where).
		Delete(&models.MoneyAmount{})
	if err := queryBuilder.Error; err != nil {
		return r.HandleDBError(err)
	}
	queryBuilder = r.Db().Model(&models.ProductVariantMoneyAmount{}).
		Where("money_amount_id IN ?", maIdsForVariant).
		Delete(&models.ProductVariantMoneyAmount{})
	if err := queryBuilder.Error; err != nil {
		return r.HandleDBError(err)
	}
	return nil
}

func (r *MoneyAmountRepo) UpsertVariantCurrencyPrice(variantId uuid.UUID, price models.MoneyAmount) (*models.MoneyAmount, *utils.ApplictaionError) {
	var moneyAmount *models.MoneyAmount
	queryBuilder := r.Db().Model(&models.MoneyAmount{}).
		Joins("LEFT JOIN product_variant_money_amount pvma ON pvma.money_amount_id = ma.id").
		Where("pvma.variant_id = ?", variantId.String()).
		Where("ma.currency_code = ?", price.CurrencyCode).
		Where("ma.region_id IS NULL").
		Where("ma.price_list_id IS NULL").
		First(&moneyAmount)
	if queryBuilder.Error != nil && queryBuilder.Error != gorm.ErrRecordNotFound {
		return nil, r.HandleDBError(queryBuilder.Error)
	}
	created := false
	if queryBuilder.Error == gorm.ErrRecordNotFound {
		moneyAmount = &models.MoneyAmount{
			Amount:       price.Amount,
			CurrencyCode: strings.ToLower(price.CurrencyCode),
			VariantId:    uuid.NullUUID{UUID: variantId},
		}
		created = true
	} else {
		moneyAmount.Amount = price.Amount
	}

	if err := r.Db().Save(&moneyAmount).Error; err != nil {
		return nil, r.HandleDBError(err)
	}
	if created {
		variantMoneyAmount := models.ProductVariantMoneyAmount{
			VariantId:     uuid.NullUUID{UUID: variantId},
			MoneyAmountId: uuid.NullUUID{UUID: moneyAmount.Id},
		}

		if err := r.Db().Create(&variantMoneyAmount).Error; err != nil {
			return nil, r.HandleDBError(err)
		}
	}
	return moneyAmount, nil
}

func (r *MoneyAmountRepo) AddPriceListPrices(priceListId uuid.UUID, prices []models.MoneyAmount, overrideExisting bool) ([]models.MoneyAmount, *utils.ApplictaionError) {
	var toInsert []models.MoneyAmount
	var joinTableValues []models.ProductVariantMoneyAmount
	for _, price := range prices {
		joinTableValue := models.ProductVariantMoneyAmount{
			VariantId:     price.VariantId,
			MoneyAmountId: uuid.NullUUID{UUID: price.Id},
		}
		toInsert = append(toInsert, price)
		joinTableValues = append(joinTableValues, joinTableValue)
	}
	queryBuilder := r.Db().Model(&models.MoneyAmount{}).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"amount", "currency_code", "region_id", "price_list_id"}),
		}).
		Create(&toInsert)
	if err := queryBuilder.Error; err != nil {
		return nil, r.HandleDBError(err)
	}
	if overrideExisting {
		queryBuilder = r.Db().Model(&models.MoneyAmount{}).
			Where("price_list_id = ?", priceListId).
			Where("id NOT IN ?", toInsert).
			Delete(&models.MoneyAmount{})
		if err := queryBuilder.Error; err != nil {
			return nil, r.HandleDBError(err)
		}
		queryBuilder = r.Db().Model(&models.ProductVariantMoneyAmount{}).
			Where("money_amount_id IN ?", toInsert).
			Delete(&models.ProductVariantMoneyAmount{})
		if err := queryBuilder.Error; err != nil {
			return nil, r.HandleDBError(err)
		}
	}

	if err := r.Db().Create(&joinTableValues).Error; err != nil {
		return nil, r.HandleDBError(err)
	}
	return toInsert, nil
}

func (r *MoneyAmountRepo) DeletePriceListPrices(priceListId uuid.UUID, moneyAmountIds uuid.UUIDs) *utils.ApplictaionError {
	queryBuilder := r.Db().Model(&models.MoneyAmount{}).
		Where("price_list_id = ?", priceListId.String()).
		Where("id IN ?", moneyAmountIds.Strings()).
		Delete(&models.MoneyAmount{})
	if err := queryBuilder.Error; err != nil {
		return r.HandleDBError(err)
	}
	return nil
}

func (r *MoneyAmountRepo) FindManyForVariantInPriceList(variantId uuid.UUID, priceListId uuid.UUID, requiresPriceList bool) ([]models.MoneyAmount, *int64, *utils.ApplictaionError) {
	var prices []models.MoneyAmount
	queryBuilder := r.Db().Model(&models.MoneyAmount{}).
		Joins("LEFT JOIN price_list ON price_list.id = ma.price_list_id").
		Joins("LEFT JOIN product_variant_money_amount pvma ON pvma.money_amount_id = ma.id").
		Where("pvma.variant_id = ?", variantId.String())
	if requiresPriceList {
		queryBuilder = queryBuilder.Where("ma.price_list_id = ?", priceListId.String())
	} else {
		queryBuilder = queryBuilder.Where("ma.price_list_id IS NULL OR ma.price_list_id = ?", priceListId.String())
	}
	if err := queryBuilder.Find(&prices).Error; err != nil {
		return nil, nil, r.HandleDBError(err)
	}
	return prices, &queryBuilder.RowsAffected, nil
}

func (r *MoneyAmountRepo) FindManyForVariantInRegion(variantId uuid.UUID, regionId uuid.UUID, currencyCode string, customerId uuid.UUID, includeDiscountPrices bool, includeTaxInclusivePricing bool) ([]models.MoneyAmount, *int64, *utils.ApplictaionError) {
	var prices []models.MoneyAmount
	queryBuilder := r.Db().Model(&models.MoneyAmount{}).
		Joins("LEFT JOIN price_list ON price_list.id = ma.price_list_id").
		Joins("LEFT JOIN product_variant_money_amount pvma ON pvma.money_amount_id = ma.id").
		Where("pvma.variant_id = ?", variantId.String()).
		Where("ma.region_id = ?", regionId.String()).
		Where("ma.currency_code = ?", currencyCode).
		Where("ma.price_list_id IS NULL")
	if includeTaxInclusivePricing {
		queryBuilder = queryBuilder.
			Joins("LEFT JOIN currency ON currency.code = ma.currency_code").
			Joins("LEFT JOIN region ON region.id = ma.region_id").
			Select("currency.includes_tax, region.includes_tax")
	}
	if customerId != uuid.Nil {
		queryBuilder = queryBuilder.
			Joins("LEFT JOIN customer_group_customers cgc ON cgc.customer_id = ?", customerId.String()).
			Joins("LEFT JOIN customer_groups cgroup ON cgroup.id = cgc.customer_group_id").
			Where("cgroup.id IS NULL OR cgc.customer_id = ?", customerId.String())
	} else {
		queryBuilder = queryBuilder.
			Joins("LEFT JOIN customer_groups cgroup ON cgroup.id IS NULL")
	}
	if err := queryBuilder.Find(&prices).Error; err != nil {
		return nil, nil, r.HandleDBError(err)
	}
	return prices, &queryBuilder.RowsAffected, nil
}

func (r *MoneyAmountRepo) FindCurrencyMoneyAmounts(where []models.MoneyAmount) ([]models.MoneyAmount, *utils.ApplictaionError) {
	var results []models.MoneyAmount
	queryBuilder := r.Db().Model(&models.MoneyAmount{}).
		Joins("LEFT JOIN product_variant_money_amount pvma ON pvma.money_amount_id = ma.id").
		Select("pvma.variant_id, ma.id").
		Where("ma.region_id IS NULL").
		Where("ma.price_list_id IS NULL")
	for _, w := range where {
		queryBuilder = queryBuilder.Or("(pvma.variant_id = ? AND ma.currency_code = ?)", w.VariantId, w.CurrencyCode)
	}
	if err := queryBuilder.Find(&results).Error; err != nil {
		return nil, r.HandleDBError(err)
	}
	return results, nil
}

func (r *MoneyAmountRepo) FindRegionMoneyAmounts(where []models.MoneyAmount) ([]models.MoneyAmount, *utils.ApplictaionError) {
	var results []models.MoneyAmount
	queryBuilder := r.Db().Model(&models.MoneyAmount{}).
		Joins("LEFT JOIN product_variant_money_amount pvma ON pvma.money_amount_id = ma.id").
		Select("pvma.variant_id, ma.id").
		Where("ma.price_list_id IS NULL")
	for _, w := range where {
		queryBuilder = queryBuilder.Or("(pvma.variant_id = ? AND ma.region_id = ?)", w.VariantId, w.RegionId)
	}
	if err := queryBuilder.Find(&results).Error; err != nil {
		return nil, r.HandleDBError(err)
	}
	return results, nil
}

func (r *MoneyAmountRepo) FindManyForVariantsInRegion(variantIds uuid.UUIDs, regionId uuid.UUID, currencyCode string, customerId uuid.UUID, includeDiscountPrices bool, includeTaxInclusivePricing bool) (map[uuid.UUID][]models.MoneyAmount, *int64, *utils.ApplictaionError) {
	if len(variantIds) == 0 {
		return nil, nil, nil
	}
	var prices []models.MoneyAmount
	queryBuilder := r.Db().Model(&models.MoneyAmount{}).
		Joins("LEFT JOIN price_list ON price_list.id = ma.price_list_id").
		Joins("LEFT JOIN product_variant_money_amount pvma ON pvma.money_amount_id = ma.id").
		Where("pvma.variant_id IN ?", variantIds).
		Where("ma.price_list_id IS NULL OR price_list.status = 'active'").
		Where("price_list.ends_at IS NULL OR price_list.ends_at > ?", time.Now().UTC()).
		Where("price_list.starts_at IS NULL OR price_list.starts_at < ?", time.Now().UTC())
	if includeTaxInclusivePricing {
		queryBuilder = queryBuilder.
			Joins("LEFT JOIN currency ON currency.code = ma.currency_code").
			Joins("LEFT JOIN region ON region.id = ma.region_id").
			Select("currency.includes_tax, region.includes_tax")
	}
	if regionId != uuid.Nil || currencyCode != "" {
		queryBuilder = queryBuilder.Where(func(db *gorm.DB) *gorm.DB {
			if regionId != uuid.Nil && currencyCode == "" {
				return db.Where("ma.region_id = ?", regionId)
			}
			if regionId == uuid.Nil && currencyCode != "" {
				return db.Where("ma.currency_code = ?", currencyCode)
			}
			if regionId != uuid.Nil && currencyCode != "" {
				return db.Where("ma.region_id = ? OR ma.currency_code = ?", regionId, currencyCode)
			}
			return db.Where("price_list.id IS NULL")
		})
	} else if customerId == uuid.Nil && !includeDiscountPrices {
		queryBuilder = queryBuilder.Where("price_list.id IS NULL")
	}
	if customerId != uuid.Nil {
		queryBuilder = queryBuilder.
			Joins("LEFT JOIN customer_group_customers cgc ON cgc.customer_id = ?", customerId).
			Joins("LEFT JOIN customer_groups cgroup ON cgroup.id = cgc.customer_group_id").
			Where("cgroup.id IS NULL OR cgc.customer_id = ?", customerId)
	} else {
		queryBuilder = queryBuilder.
			Joins("LEFT JOIN customer_groups cgroup ON cgroup.id IS NULL")
	}
	if err := queryBuilder.Find(&prices).Error; err != nil {
		return nil, nil, r.HandleDBError(err)
	}
	groupedPrices := lo.GroupBy(prices, func(price models.MoneyAmount) uuid.UUID {
		return price.VariantId.UUID
	})
	return groupedPrices, &queryBuilder.RowsAffected, nil
}

func (r *MoneyAmountRepo) UpdatePriceListPrices(priceListId uuid.UUID, updates []models.MoneyAmount) ([]models.MoneyAmount, *utils.ApplictaionError) {
	var existingPrices []models.MoneyAmount
	var newPrices []models.MoneyAmount
	for _, update := range updates {
		if update.Id != uuid.Nil {
			existingPrices = append(existingPrices, update)
		} else {
			newPrices = append(newPrices, update)
		}
	}
	var newPriceEntities []models.MoneyAmount
	var joinTableValues []models.ProductVariantMoneyAmount
	for _, price := range newPrices {
		price.PriceListId = uuid.NullUUID{UUID: priceListId}
		joinTableValue := models.ProductVariantMoneyAmount{
			VariantId:     price.VariantId,
			MoneyAmountId: uuid.NullUUID{UUID: price.Id},
		}
		newPriceEntities = append(newPriceEntities, price)
		joinTableValues = append(joinTableValues, joinTableValue)
	}

	tx := r.Db().Begin()

	if err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"amount", "currency_code", "region_id", "price_list_id"}),
	}).Save(&existingPrices).Error; err != nil {
		tx.Rollback()
		return nil, r.HandleDBError(err)
	}

	if len(newPrices) > 0 {
		if err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"amount", "currency_code", "region_id", "price_list_id"}),
		}).Create(&newPriceEntities).Error; err != nil {
			tx.Rollback()
			return nil, r.HandleDBError(err)
		}
	}

	if err := tx.Where("money_amount_id IN ?", append(existingPrices, newPrices...)).Delete(&models.ProductVariantMoneyAmount{}).Error; err != nil {
		tx.Rollback()
		return nil, r.HandleDBError(err)
	}

	if err := tx.Create(&joinTableValues).Error; err != nil {
		tx.Rollback()
		return nil, r.HandleDBError(err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, r.HandleDBError(err)
	}

	return append(existingPrices, newPrices...), nil
}

func (r *MoneyAmountRepo) GetPricesForVariantInRegion(variantId uuid.UUID, regionId uuid.UUID) ([]models.MoneyAmount, *utils.ApplictaionError) {
	var prices []models.MoneyAmount
	err := r.Db().Model(&models.MoneyAmount{}).
		Joins("LEFT JOIN product_variant_money_amount pvma ON pvma.money_amount_id = ma.id").
		Where("pvma.variant_id = ?", variantId).
		Where("ma.region_id = ?", regionId).
		Find(&prices).Error
	if err != nil {
		return nil, r.HandleDBError(err)
	}
	return prices, nil
}

func (r *MoneyAmountRepo) CreateProductVariantMoneyAmounts(toCreate []models.MoneyAmount) *utils.ApplictaionError {
	queryBuilder := r.Db().Create(toCreate)
	if err := queryBuilder.Error; err != nil {
		return r.HandleDBError(err)
	}
	return nil
}

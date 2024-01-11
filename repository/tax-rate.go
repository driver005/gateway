package repository

import (
	"slices"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var resolveableFields = []string{
	"product_count",
	"product_type_count",
	"shipping_option_count",
}

type TaxRateRepo struct {
	sql.Repository[models.TaxRate]
}

func TaxRateRepository(db *gorm.DB) *TaxRateRepo {
	return &TaxRateRepo{*sql.NewRepository[models.TaxRate](db)}
}

func (r *TaxRateRepo) GetFindQueryBuilder(options sql.Query) *gorm.DB {
	qb := r.Db().Model(&models.TaxRate{})
	cleanOptions := options
	resolverFields := []string{}
	if options.Selects != nil {
		selectableCols := []string{}
		for _, k := range options.Selects {
			if !slices.Contains[[]string](resolveableFields, k) {
				selectableCols = append(selectableCols, k)
			} else {
				resolverFields = append(resolverFields, k)
			}
		}
		cleanOptions.Selects = selectableCols
	}
	qb.Set("gorm:query_option", cleanOptions)
	if len(resolverFields) > 0 {
		r.ApplyResolutionsToQueryBuilder(qb, resolverFields)
	}
	return qb
}

func (r *TaxRateRepo) FindWithResolution(options sql.Query) ([]models.TaxRate, *utils.ApplictaionError) {
	qb := r.GetFindQueryBuilder(options)
	var results []models.TaxRate
	err := qb.Find(&results)
	return results, r.HandleError(err)
}

func (r *TaxRateRepo) FindOneWithResolution(options sql.Query) (*models.TaxRate, *utils.ApplictaionError) {
	qb := r.GetFindQueryBuilder(options)
	var result models.TaxRate
	err := qb.First(&result)
	return &result, r.HandleError(err)
}

func (r *TaxRateRepo) FindAndCountWithResolution(options sql.Query) ([]models.TaxRate, *int64, *utils.ApplictaionError) {
	qb := r.GetFindQueryBuilder(options)
	var results []models.TaxRate
	var count *int64
	if err := qb.Find(&results); err.Error != nil {
		return nil, nil, r.HandleError(err)
	}
	if err := qb.Count(count); err.Error != nil {
		return nil, nil, r.HandleError(err)
	}
	return results, count, nil
}

func (r *TaxRateRepo) ApplyResolutionsToQueryBuilder(qb *gorm.DB, resolverFields []string) *gorm.DB {
	for _, k := range resolverFields {
		switch k {
		case "product_count":
			qb = qb.Joins("JOIN products ON products.tax_rate_id = tax_rates.id").
				Select("COUNT(DISTINCT products.id) AS product_count")
		case "product_type_count":
			qb = qb.Joins("JOIN product_types ON product_types.tax_rate_id = tax_rates.id").
				Select("COUNT(DISTINCT product_types.id) AS product_type_count")
		case "shipping_option_count":
			qb = qb.Joins("JOIN shipping_options ON shipping_options.tax_rate_id = tax_rates.id").
				Select("COUNT(DISTINCT shipping_options.id) AS shipping_option_count")
		}
	}
	return qb
}

func (r *TaxRateRepo) RemoveFromProduct(id uuid.UUID, productIds uuid.UUIDs) *utils.ApplictaionError {
	return r.HandleError(r.Db().Delete(&models.ProductTaxRate{}, "rate_id = ? AND product_id IN (?)", id, productIds))
}

func (r *TaxRateRepo) AddToProduct(id uuid.UUID, productIds uuid.UUIDs, overrideExisting bool) ([]models.ProductTaxRate, *utils.ApplictaionError) {
	toInsert := []models.ProductTaxRate{}
	for _, pId := range productIds {
		toInsert = append(toInsert, models.ProductTaxRate{RateId: uuid.NullUUID{UUID: id}, ProductId: uuid.NullUUID{UUID: pId}})
	}
	insertResult := r.Db().Clauses(clause.OnConflict{DoNothing: true}).Create(&toInsert)
	if overrideExisting {
		r.Db().Delete(&models.ProductTaxRate{}, "rate_id = ? AND product_id NOT IN (?)", id, productIds)
	}
	var results []models.ProductTaxRate
	err := r.Db().Model(&models.ProductTaxRate{}).Where("id IN (?)", insertResult.RowsAffected).Find(&results)
	return results, r.HandleError(err)
}

func (r *TaxRateRepo) RemoveFromProductType(id uuid.UUID, productTypeIds uuid.UUIDs) *utils.ApplictaionError {
	return r.HandleError(r.Db().Delete(&models.ProductTypeTaxRate{}, "rate_id = ? AND product_type_id IN (?)", id, productTypeIds))
}

func (r *TaxRateRepo) AddToProductType(id uuid.UUID, productTypeIds uuid.UUIDs, overrideExisting bool) ([]models.ProductTypeTaxRate, *utils.ApplictaionError) {
	toInsert := []models.ProductTypeTaxRate{}
	for _, pId := range productTypeIds {
		toInsert = append(toInsert, models.ProductTypeTaxRate{RateId: uuid.NullUUID{UUID: id}, ProductTypeId: uuid.NullUUID{UUID: pId}})
	}
	insertResult := r.Db().Clauses(clause.OnConflict{DoNothing: true}).Create(&toInsert)
	if overrideExisting {
		r.Db().Delete(&models.ProductTypeTaxRate{}, "rate_id = ? AND product_type_id NOT IN (?)", id, productTypeIds)
	}
	var results []models.ProductTypeTaxRate
	err := r.Db().Model(&models.ProductTypeTaxRate{}).Where("id IN (?)", insertResult.RowsAffected).Find(&results)
	return results, r.HandleError(err)
}

func (r *TaxRateRepo) RemoveFromShippingOption(id uuid.UUID, optionIds uuid.UUIDs) *utils.ApplictaionError {
	return r.HandleError(r.Db().Delete(&models.ShippingTaxRate{}, "rate_id = ? AND shipping_option_id IN (?)", id, optionIds))
}

func (r *TaxRateRepo) AddToShippingOption(id uuid.UUID, optionIds uuid.UUIDs, overrideExisting bool) ([]models.ShippingTaxRate, *utils.ApplictaionError) {
	toInsert := []models.ShippingTaxRate{}
	for _, pId := range optionIds {
		toInsert = append(toInsert, models.ShippingTaxRate{RateId: uuid.NullUUID{UUID: id}, ShippingOptionId: uuid.NullUUID{UUID: pId}})
	}
	insertResult := r.Db().Clauses(clause.OnConflict{DoNothing: true}).Create(&toInsert)
	if overrideExisting {
		r.Db().Delete(&models.ShippingTaxRate{}, "rate_id = ? AND shipping_option_id NOT IN (?)", id, optionIds)
	}
	var results []models.ShippingTaxRate
	err := r.Db().Model(&models.ShippingTaxRate{}).Where("id IN (?)", insertResult.RowsAffected).Find(&results)
	return results, r.HandleError(err)
}

func (r *TaxRateRepo) ListByProduct(productId uuid.UUID, config types.TaxRateListByConfig) ([]models.TaxRate, *utils.ApplictaionError) {
	productRates := r.Db().Model(&models.TaxRate{}).
		Joins("JOIN product_tax_rates ON product_tax_rates.rate_id = tax_rates.id").
		Joins("JOIN products ON products.id = product_tax_rates.product_id").
		Where("products.id = ?", productId)
	typeRates := r.Db().Model(&models.TaxRate{}).
		Joins("JOIN product_type_tax_rates ON product_type_tax_rates.rate_id = tax_rates.id").
		Joins("JOIN products ON products.type_id = product_type_tax_rates.product_type_id").
		Where("products.id = ?", productId)
	if config.RegionId != uuid.Nil {
		productRates = productRates.Where("tax_rates.region_id = ?", config.RegionId)
		typeRates = typeRates.Where("tax_rates.region_id = ?", config.RegionId)
	}
	var results []models.TaxRate
	if err := productRates.Find(&results); err.Error != nil {
		return nil, r.HandleError(err)
	}
	if err := typeRates.Find(&results); err.Error != nil {
		return nil, r.HandleError(err)
	}
	return results, nil
}

func (r *TaxRateRepo) ListByShippingOption(optionId uuid.UUID) ([]models.TaxRate, *utils.ApplictaionError) {
	rates := r.Db().Model(&models.TaxRate{}).
		Joins("JOIN shipping_tax_rates ON shipping_tax_rates.rate_id = tax_rates.id").
		Where("shipping_tax_rates.shipping_option_id = ?", optionId)
	var results []models.TaxRate
	err := rates.Find(&results)
	return results, r.HandleError(err)
}

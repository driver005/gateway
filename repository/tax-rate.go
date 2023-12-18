package repository

import (
	"slices"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
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
	Repository[models.TaxRate]
}

func TaxRateRepository(db *gorm.DB) TaxRateRepo {
	return TaxRateRepo{*NewRepository[models.TaxRate](db)}
}

func (r *TaxRateRepo) GetFindQueryBuilder(options Query) *gorm.DB {
	qb := r.db.Model(&models.TaxRate{})
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

func (r *TaxRateRepo) FindWithResolution(options Query) ([]models.TaxRate, error) {
	qb := r.GetFindQueryBuilder(options)
	var results []models.TaxRate
	err := qb.Find(&results).Error
	return results, err
}

func (r *TaxRateRepo) FindOneWithResolution(options Query) (*models.TaxRate, error) {
	qb := r.GetFindQueryBuilder(options)
	var result models.TaxRate
	err := qb.First(&result).Error
	return &result, err
}

func (r *TaxRateRepo) FindAndCountWithResolution(options Query) ([]models.TaxRate, *int64, error) {
	qb := r.GetFindQueryBuilder(options)
	var results []models.TaxRate
	var count *int64
	if err := qb.Find(&results).Error; err != nil {
		return nil, nil, err
	}
	if err := qb.Count(count).Error; err != nil {
		return nil, nil, err
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

func (r *TaxRateRepo) RemoveFromProduct(id uuid.UUID, productIds uuid.UUIDs) error {
	return r.db.Delete(&models.ProductTaxRate{}, "rate_id = ? AND product_id IN (?)", id, productIds).Error
}

func (r *TaxRateRepo) AddToProduct(id uuid.UUID, productIds uuid.UUIDs, overrideExisting bool) ([]models.ProductTaxRate, error) {
	toInsert := []models.ProductTaxRate{}
	for _, pId := range productIds {
		toInsert = append(toInsert, models.ProductTaxRate{RateId: uuid.NullUUID{UUID: id}, ProductId: uuid.NullUUID{UUID: pId}})
	}
	insertResult := r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&toInsert)
	if overrideExisting {
		r.db.Delete(&models.ProductTaxRate{}, "rate_id = ? AND product_id NOT IN (?)", id, productIds)
	}
	var results []models.ProductTaxRate
	err := r.db.Model(&models.ProductTaxRate{}).Where("id IN (?)", insertResult.RowsAffected).Find(&results).Error
	return results, err
}

func (r *TaxRateRepo) RemoveFromProductType(id uuid.UUID, productTypeIds uuid.UUIDs) error {
	return r.db.Delete(&models.ProductTypeTaxRate{}, "rate_id = ? AND product_type_id IN (?)", id, productTypeIds).Error
}

func (r *TaxRateRepo) AddToProductType(id uuid.UUID, productTypeIds uuid.UUIDs, overrideExisting bool) ([]models.ProductTypeTaxRate, error) {
	toInsert := []models.ProductTypeTaxRate{}
	for _, pId := range productTypeIds {
		toInsert = append(toInsert, models.ProductTypeTaxRate{RateId: uuid.NullUUID{UUID: id}, ProductTypeId: uuid.NullUUID{UUID: pId}})
	}
	insertResult := r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&toInsert)
	if overrideExisting {
		r.db.Delete(&models.ProductTypeTaxRate{}, "rate_id = ? AND product_type_id NOT IN (?)", id, productTypeIds)
	}
	var results []models.ProductTypeTaxRate
	err := r.db.Model(&models.ProductTypeTaxRate{}).Where("id IN (?)", insertResult.RowsAffected).Find(&results).Error
	return results, err
}

func (r *TaxRateRepo) RemoveFromShippingOption(id uuid.UUID, optionIds uuid.UUIDs) error {
	return r.db.Delete(&models.ShippingTaxRate{}, "rate_id = ? AND shipping_option_id IN (?)", id, optionIds).Error
}

func (r *TaxRateRepo) AddToShippingOption(id uuid.UUID, optionIds uuid.UUIDs, overrideExisting bool) ([]models.ShippingTaxRate, error) {
	toInsert := []models.ShippingTaxRate{}
	for _, pId := range optionIds {
		toInsert = append(toInsert, models.ShippingTaxRate{RateId: uuid.NullUUID{UUID: id}, ShippingOptionId: uuid.NullUUID{UUID: pId}})
	}
	insertResult := r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&toInsert)
	if overrideExisting {
		r.db.Delete(&models.ShippingTaxRate{}, "rate_id = ? AND shipping_option_id NOT IN (?)", id, optionIds)
	}
	var results []models.ShippingTaxRate
	err := r.db.Model(&models.ShippingTaxRate{}).Where("id IN (?)", insertResult.RowsAffected).Find(&results).Error
	return results, err
}

func (r *TaxRateRepo) ListByProduct(productId uuid.UUID, config types.TaxRateListByConfig) ([]models.TaxRate, error) {
	productRates := r.db.Model(&models.TaxRate{}).
		Joins("JOIN product_tax_rates ON product_tax_rates.rate_id = tax_rates.id").
		Joins("JOIN products ON products.id = product_tax_rates.product_id").
		Where("products.id = ?", productId)
	typeRates := r.db.Model(&models.TaxRate{}).
		Joins("JOIN product_type_tax_rates ON product_type_tax_rates.rate_id = tax_rates.id").
		Joins("JOIN products ON products.type_id = product_type_tax_rates.product_type_id").
		Where("products.id = ?", productId)
	if config.RegionId != uuid.Nil {
		productRates = productRates.Where("tax_rates.region_id = ?", config.RegionId)
		typeRates = typeRates.Where("tax_rates.region_id = ?", config.RegionId)
	}
	var results []models.TaxRate
	if err := productRates.Find(&results).Error; err != nil {
		return nil, err
	}
	if err := typeRates.Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (r *TaxRateRepo) ListByShippingOption(optionId uuid.UUID) ([]models.TaxRate, error) {
	rates := r.db.Model(&models.TaxRate{}).
		Joins("JOIN shipping_tax_rates ON shipping_tax_rates.rate_id = tax_rates.id").
		Where("shipping_tax_rates.shipping_option_id = ?", optionId)
	var results []models.TaxRate
	err := rates.Find(&results).Error
	return results, err
}

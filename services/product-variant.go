package services

import (
	"context"
	"fmt"
	"slices"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type ProductVariantService struct {
	ctx context.Context
	r   Registry
}

func NewProductVariantService(
	r Registry,
) *ProductVariantService {
	return &ProductVariantService{
		context.Background(),
		r,
	}
}

func (s *ProductVariantService) SetContext(context context.Context) *ProductVariantService {
	s.ctx = context
	return s
}

func (s *ProductVariantService) Retrieve(id uuid.UUID, config sql.Options) (*models.ProductVariant, *utils.ApplictaionError) {
	var res *models.ProductVariant

	query := sql.BuildQuery(models.ProductVariant{Model: core.Model{Id: id}}, config)

	if err := s.r.ProductVariantRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ProductVariantService) RetrieveBySKU(sku string, config sql.Options) (*models.ProductVariant, *utils.ApplictaionError) {
	priceIndex := -1
	if config.Relations != nil {
		for i, relation := range config.Relations {
			if relation == "prices" {
				priceIndex = i
				break
			}
		}
	}
	if priceIndex >= 0 && config.Relations != nil {
		config.Relations = append(config.Relations[:priceIndex], config.Relations[priceIndex+1:]...)
	}
	var res *models.ProductVariant

	query := sql.BuildQuery(models.ProductVariant{Sku: sku}, config)

	if err := s.r.ProductVariantRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ProductVariantService) Create(productId uuid.UUID, product *models.Product, variants []models.ProductVariant) ([]models.ProductVariant, *utils.ApplictaionError) {
	if productId != uuid.Nil {
		p, err := s.r.ProductService().SetContext(s.ctx).RetrieveById(productId, sql.Options{Relations: []string{"variants", "variants.options", "options"}})
		if err != nil {
			return nil, err
		}
		product = p
	}
	if product == nil || product.Id == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Product id missing",
			"500",
			nil,
		)
	}
	err := s.ValidateVariantsToCreate(product, variants)
	if err != nil {
		return nil, err
	}

	computedRank := len(product.Variants)
	var variantPricesToUpdate []models.MoneyAmount
	var results []models.ProductVariant
	for _, variant := range variants {
		variant.VariantRank = computedRank
		variant.ProductId = uuid.NullUUID{UUID: product.Id}

		if err := s.r.ProductVariantRepository().Save(s.ctx, &variant); err != nil {
			return nil, err
		}

		if len(variant.Prices) > 0 {
			variantPricesToUpdate = append(variantPricesToUpdate, variant.Prices...)
		}
		results = append(results, variant)
		computedRank++
	}
	if len(variantPricesToUpdate) > 0 {
		s.UpdateVariantPrices(variantPricesToUpdate)
	}
	// eventsToEmit := make([]struct {
	// 	EventName string
	// 	Data      struct {
	// 		ID        string
	// 		ProductID string
	// 	}
	// }, len(results))
	// for i, result := range results {
	// 	eventsToEmit[i] = struct {
	// 		EventName string
	// 		Data      struct {
	// 			ID        string
	// 			ProductID string
	// 		}
	// 	}{EventName: "product-variant.created", Data: struct {
	// 		ID        string
	// 		ProductID string
	// 	}{ID: result.ID, ProductID: result.ProductID}}
	// }
	// err = s.eventBus_.emit(eventsToEmit)
	// if err != nil {
	// 	return nil, err
	// }
	return results, nil
}

func (s *ProductVariantService) Update(id uuid.UUID, update *models.ProductVariant) (*models.ProductVariant, *utils.ApplictaionError) {
	variant, err := s.Retrieve(id, sql.Options{})
	if err != nil {
		return nil, err
	}

	if len(update.Prices) > 0 {
		s.UpdateVariantPrices(update.Prices)
	}

	for _, option := range update.Options {
		s.UpdateOptionValue(variant.Id, option.OptionId.UUID, option.Value)
	}

	update.Id = variant.Id

	if err := s.r.ProductVariantRepository().Update(s.ctx, update); err != nil {
		return nil, err
	}

	return update, nil
}

func (s *ProductVariantService) UpdateBatch(ids uuid.UUIDs, variants []models.ProductVariant) ([]models.ProductVariant, *utils.ApplictaionError) {
	if len(ids) != len(variants) {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			"length of `ids` and `variants` needs to be the same",
			"500",
			nil,
		)
	}

	var res []models.ProductVariant

	for i, variant := range variants {
		v, err := s.Update(ids[i], &variant)
		if err != nil {
			return nil, err
		}

		res = append(res, *v)
	}
	// events := make([]struct {
	// 	EventName string
	// 	Data      struct {
	// 		ID        string
	// 		ProductID string
	// 		Fields    []string
	// 	}
	// }, len(results))
	// for i, result := range results {
	// 	events[i] = struct {
	// 		EventName string
	// 		Data      struct {
	// 			ID        string
	// 			ProductID string
	// 			Fields    []string
	// 		}
	// 	}{EventName: "product-variant.updated", Data: struct {
	// 		ID        string
	// 		ProductID string
	// 		Fields    []string
	// 	}{ID: result.Variant.ID, ProductID: result.Variant.ProductID, Fields: Object.keys(result.UpdateData)}}
	// }
	// if len(events) > 0 {
	// 	s.eventBus_.emit(events)
	// }
	// return results.map(func(result struct {
	// 	Variant models.ProductVariant
	// 	UpdateData UpdateProductVariantInput
	// 	ShouldEmitUpdateEvent bool
	// }) models.ProductVariant {
	// 	return result.Variant
	// })

	return res, nil
}

func (s *ProductVariantService) UpdateVariantPrices(data []models.MoneyAmount) *utils.ApplictaionError {
	s.r.MoneyAmountRepository().DeleteVariantPricesNotIn(uuid.Nil, data)
	var regionIds uuid.UUIDs
	for _, d := range data {
		if d.RegionId.UUID != uuid.Nil {
			regionIds = append(regionIds, d.RegionId.UUID)
		}
	}

	regions, err := s.r.RegionService().SetContext(s.ctx).List(models.Region{}, sql.Options{
		Selects:       []string{"id", "currency_code"},
		Specification: []sql.Specification{sql.In("id", regionIds)},
	})
	if err != nil {
		return err
	}
	regionsMap := make(map[uuid.UUID]models.Region)
	for _, region := range regions {
		regionsMap[region.Id] = region
	}
	var dataRegionPrices []models.MoneyAmount
	var dataCurrencyPrices []models.MoneyAmount
	for _, d := range data {
		if d.RegionId.UUID != uuid.Nil {
			region := regionsMap[d.RegionId.UUID]
			dataRegionPrices = append(dataRegionPrices, models.MoneyAmount{
				VariantId:    d.VariantId,
				CurrencyCode: region.CurrencyCode,
				RegionId:     d.RegionId,
				Amount:       d.Amount,
			})
		} else {
			dataCurrencyPrices = append(dataCurrencyPrices, models.MoneyAmount{
				VariantId:    d.VariantId,
				CurrencyCode: d.CurrencyCode,
				Amount:       d.Amount,
			})
		}

	}
	if len(dataRegionPrices) > 0 {
		s.UpsertRegionPrices(dataRegionPrices)
	}
	if len(dataCurrencyPrices) > 0 {
		s.UpsertCurrencyPrices(dataCurrencyPrices)
	}
	return nil
}

func (s *ProductVariantService) UpsertRegionPrices(data []models.MoneyAmount) *utils.ApplictaionError {
	var where []models.MoneyAmount
	for _, d := range data {
		where = append(where, models.MoneyAmount{
			VariantId: d.VariantId,
			RegionId:  d.RegionId,
		})
	}
	moneyAmounts, err := s.r.MoneyAmountRepository().FindRegionMoneyAmounts(where)
	if err != nil {
		return err
	}
	moneyAmountsMapToVariantId := make(map[uuid.UUID][]models.MoneyAmount)
	for _, d := range moneyAmounts {
		mas := moneyAmountsMapToVariantId[d.VariantId.UUID]
		mas = append(mas, d)
		moneyAmountsMapToVariantId[d.VariantId.UUID] = mas
	}
	var dataToCreate []models.MoneyAmount
	var dataToUpdate []models.MoneyAmount
	for _, d := range data {
		variantMoneyAmounts := moneyAmountsMapToVariantId[d.VariantId.UUID]
		var moneyAmount *models.MoneyAmount
		for _, ma := range variantMoneyAmounts {
			if ma.RegionId == d.RegionId {
				moneyAmount = &ma
				break
			}
		}
		if moneyAmount != nil {
			if moneyAmount.Amount != d.Amount {
				dataToUpdate = append(dataToUpdate, models.MoneyAmount{
					Model: core.Model{
						Id: moneyAmount.Id,
					},
					Amount: d.Amount,
				})
			}
		} else {
			if err := s.r.MoneyAmountRepository().Create(s.ctx, &d); err != nil {
				return err
			}
			d.Variant = &models.ProductVariant{Model: core.Model{Id: d.VariantId.UUID}}
			dataToCreate = append(dataToCreate, d)
		}
	}
	if len(dataToCreate) > 0 {
		_, err := s.r.MoneyAmountRepository().InsertBulk(dataToCreate)
		if err != nil {
			return err
		}
	}
	if len(dataToUpdate) > 0 {
		for _, data := range dataToUpdate {
			if err := s.r.MoneyAmountRepository().Update(s.ctx, &data); err != nil {
				return err
			}
		}
	}
	if len(dataToCreate) > 0 || len(dataToUpdate) > 0 {
		var variantIds uuid.UUIDs
		for _, d := range data {
			variantIds = append(variantIds, d.VariantId.UUID)
		}
		s.r.PriceSelectionStrategy().OnVariantsPricesUpdate(variantIds)
	}
	return nil
}

func (s *ProductVariantService) UpsertCurrencyPrices(data []models.MoneyAmount) *utils.ApplictaionError {
	var where []models.MoneyAmount
	for _, d := range data {
		where = append(where, models.MoneyAmount{
			VariantId:    d.VariantId,
			CurrencyCode: d.CurrencyCode,
		})
	}
	moneyAmounts, err := s.r.MoneyAmountRepository().FindCurrencyMoneyAmounts(where)
	if err != nil {
		return err
	}
	moneyAmountsMapToVariantId := make(map[uuid.UUID][]models.MoneyAmount)
	for _, d := range moneyAmounts {
		mas := moneyAmountsMapToVariantId[d.VariantId.UUID]
		mas = append(mas, d)
		moneyAmountsMapToVariantId[d.VariantId.UUID] = mas
	}
	var dataToCreate []models.MoneyAmount
	var dataToUpdate []models.MoneyAmount
	for _, d := range data {
		variantMoneyAmounts := moneyAmountsMapToVariantId[d.VariantId.UUID]
		var moneyAmount *models.MoneyAmount
		for _, ma := range variantMoneyAmounts {
			if ma.CurrencyCode == d.CurrencyCode {
				moneyAmount = &ma
				break
			}
		}
		if moneyAmount != nil {
			if moneyAmount.Amount != d.Amount {
				dataToUpdate = append(dataToUpdate, models.MoneyAmount{
					Model: core.Model{
						Id: moneyAmount.Id,
					},
					Amount: d.Amount,
				})
			}
		} else {
			if err := s.r.MoneyAmountRepository().Create(s.ctx, &d); err != nil {
				return err
			}
			d.Variant = &models.ProductVariant{Model: core.Model{Id: d.VariantId.UUID}}
			dataToCreate = append(dataToCreate, d)
		}
	}
	if len(dataToCreate) > 0 {
		_, err := s.r.MoneyAmountRepository().InsertBulk(dataToCreate)
		if err != nil {
			return err
		}
	}
	if len(dataToUpdate) > 0 {
		for _, data := range dataToUpdate {
			if err := s.r.MoneyAmountRepository().Update(s.ctx, &data); err != nil {
				return err
			}
		}
	}
	if len(dataToCreate) > 0 || len(dataToUpdate) > 0 {
		var variantIds uuid.UUIDs
		for _, d := range data {
			variantIds = append(variantIds, d.VariantId.UUID)
		}
		s.r.PriceSelectionStrategy().OnVariantsPricesUpdate(variantIds)
	}
	return nil
}

// getRegionPrice gets the price specific to a region
func (s *ProductVariantService) GetRegionPrice(id uuid.UUID, data *interfaces.PricingContext) (*float64, *utils.ApplictaionError) {
	region, err := s.r.RegionService().SetContext(s.ctx).Retrieve(id, sql.Options{})
	if err != nil {
		return nil, err
	}

	prices, err := s.r.PriceSelectionStrategy().CalculateVariantPrice([]interfaces.Pricing{{VariantId: id, Quantity: data.Quantity}}, &interfaces.PricingContext{RegionId: data.RegionId, CurrencyCode: region.CurrencyCode, CustomerId: data.CustomerId, IncludeDiscountPrices: data.IncludeDiscountPrices})
	if err != nil {
		return nil, err
	}
	calculatedPrice := prices[id].CalculatedPrice
	return &calculatedPrice, nil
}

// setRegionPrice sets the default price of a specific region
func (s *ProductVariantService) SetRegionPrice(id uuid.UUID, price models.MoneyAmount) (*models.MoneyAmount, *utils.ApplictaionError) {
	var data *models.MoneyAmount
	moneyAmount, err := s.r.MoneyAmountRepository().GetPricesForVariantInRegion(id, price.RegionId.UUID)
	if err != nil {
		return nil, err
	}
	data = &moneyAmount[0]
	created := false
	if data == nil {
		data = &models.MoneyAmount{Amount: price.Amount, Variant: &models.ProductVariant{Model: core.Model{Id: id}}}
		created = true
	} else {
		data.Amount = price.Amount
	}
	if err := s.r.MoneyAmountRepository().Save(s.ctx, data); err != nil {
		return nil, err
	}
	if created {

		err = s.r.MoneyAmountRepository().CreateProductVariantMoneyAmounts([]models.MoneyAmount{{Model: core.Model{Id: data.Id}, VariantId: uuid.NullUUID{UUID: id}}})
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

// setCurrencyPrice sets the default price for the given currency
func (s *ProductVariantService) SetCurrencyPrice(id uuid.UUID, price models.MoneyAmount) (*models.MoneyAmount, *utils.ApplictaionError) {
	return s.r.MoneyAmountRepository().UpsertVariantCurrencyPrice(id, price)
}

// updateOptionValue updates variant's option value
func (s *ProductVariantService) UpdateOptionValue(id uuid.UUID, optionId uuid.UUID, optionValue string) (*models.ProductOptionValue, *utils.ApplictaionError) {
	var res *models.ProductOptionValue

	query := sql.BuildQuery(models.ProductOptionValue{VariantId: uuid.NullUUID{UUID: id}, OptionId: uuid.NullUUID{UUID: optionId}}, sql.Options{})

	if err := s.r.ProductOptionValueRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}

	res.Value = optionValue

	if err := s.r.ProductOptionValueRepository().Save(s.ctx, res); err != nil {
		return nil, err
	}

	return res, nil
}

// addOptionValue adds option value to a variant
func (s *ProductVariantService) AddOptionValue(id uuid.UUID, optionId uuid.UUID, optionValue string) (*models.ProductOptionValue, *utils.ApplictaionError) {
	var res *models.ProductOptionValue

	res.VariantId = uuid.NullUUID{UUID: id}
	res.OptionId = uuid.NullUUID{UUID: optionId}
	res.Value = optionValue

	if err := s.r.ProductOptionValueRepository().Save(s.ctx, res); err != nil {
		return nil, err
	}
	return res, nil
}

// deleteOptionValue deletes option value from given variant
func (s *ProductVariantService) DeleteOptionValue(id uuid.UUID, optionId uuid.UUID) *utils.ApplictaionError {
	var res *models.ProductOptionValue

	query := sql.BuildQuery(models.ProductOptionValue{VariantId: uuid.NullUUID{UUID: id}, OptionId: uuid.NullUUID{UUID: optionId}}, sql.Options{})

	if err := s.r.ProductOptionValueRepository().FindOne(s.ctx, res, query); err != nil {
		return err
	}
	if res == nil {
		return nil
	}
	if err := s.r.ProductOptionValueRepository().SoftRemove(s.ctx, res); err != nil {
		return err
	}
	return nil
}

// ListAndCount lists and counts the variants
func (s *ProductVariantService) ListAndCount(selector types.FilterableProductVariant, config sql.Options, q *string) ([]models.ProductVariant, *int64, *utils.ApplictaionError) {
	var res []models.ProductVariant

	if q != nil {
		v := sql.ILike(*q)
		selector.Title = []string{v}
		selector.SKU = []string{v}
		selector.Product.Title = v
	}

	query := sql.BuildQuery(selector, config)

	count, err := s.r.ProductVariantRepository().FindAndCount(s.ctx, res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

// list lists the variants
func (s *ProductVariantService) List(selector types.FilterableProductVariant, config sql.Options, q *string) ([]models.ProductVariant, *utils.ApplictaionError) {
	variant, _, err := s.ListAndCount(selector, config, q)
	if err != nil {
		return nil, err
	}
	return variant, nil
}

// Delete deletes the variant
func (s *ProductVariantService) Delete(variantIds uuid.UUIDs) *utils.ApplictaionError {
	// events := []map[string]interface{}{}
	for _, id := range variantIds {
		variant, err := s.Retrieve(id, sql.Options{})
		if err != nil {
			return nil
		}

		if variant == nil {
			return nil
		}

		// events = append(events, map[string]interface{}{
		// 	"eventName": "DELETED",
		// 	"data": map[string]interface{}{
		// 		"id":         variant.Id,
		// 		"product_id": variant.ProductId,
		// 		"metadata":   variant.Metadata,
		// 	},
		// })

		if err := s.r.ProductVariantRepository().SoftRemove(s.ctx, variant); err != nil {
			return nil
		}
	}
	// err = s.eventBus_.emit(events)
	// if err != nil {
	// 	return nil, err
	// }
	return nil
}

// IsVariantInSalesChannels checks if the variant is assigned to at least one of the provided sales channels
func (s *ProductVariantService) IsVariantInSalesChannels(id uuid.UUID, salesChannelIds uuid.UUIDs) (bool, *utils.ApplictaionError) {
	variant, err := s.Retrieve(id, sql.Options{
		Selects: []string{"product", "product.sales_channels"},
	})
	if err != nil {
		return false, err
	}

	var productsSalesChannels uuid.UUIDs

	for _, channel := range variant.Product.SalesChannels {
		productsSalesChannels = append(productsSalesChannels, channel.Id)
	}
	for _, id := range productsSalesChannels {
		if slices.Contains(salesChannelIds, id) {
			return true, nil
		}
	}
	return false, nil
}

// // getFreeTextQueryBuilder gets the free text query builder
// func (s *ProductVariantService) getFreeTextQueryBuilder(query FindWithRelationsOptions, q string) SelectQueryBuilder {
// 	where := query.where
// 	if where != nil {
// 		// if "title" in where {
// 		// 	Delete(where, "title")
// 		// }
// 	}
// 	qb := variantRepo.createQueryBuilder("pv").take(query.take).skip(max(query.skip, 0)).leftJoinAndSelect("pv.product", "product").Select([]string{"pv.id"}).where(where).andWhere(Brackets(func(qb SelectQueryBuilder) {
// 		qb.where(`product.title ILIKE :q`, map[string]interface{}{"q": `%${q}%`}).orWhere(`pv.title ILIKE :q`, map[string]interface{}{"q": `%${q}%`}).orWhere(`pv.sku ILIKE :q`, map[string]interface{}{"q": `%${q}%`})
// 	}))
// 	if query.withDeleted {
// 		qb = qb.withDeleted()
// 	}
// 	return qb
// }

// ValidateVariantsToCreate validates the variants to Create
func (s *ProductVariantService) ValidateVariantsToCreate(product *models.Product, variants []models.ProductVariant) *utils.ApplictaionError {
	for _, variant := range variants {
		if len(product.Options) != len(variant.Options) {
			return utils.NewApplictaionError(
				utils.INVALID_DATA,
				fmt.Sprintf("Product options length does not match variant options length. Product has %v and variant has %v.", len(product.Options), len(variant.Options)),
				"500",
				nil,
			)
		}
		for _, option := range product.Options {
			if !slices.ContainsFunc(variant.Options, func(v models.ProductOptionValue) bool {
				return option.Id == v.OptionId.UUID
			}) {
				return utils.NewApplictaionError(
					utils.INVALID_DATA,
					"Variant options do not contain value for "+option.Title,
					"500",
					nil,
				)
			}
		}

		var variantName string
		variantExists := slices.ContainsFunc(product.Variants, func(v models.ProductVariant) bool {
			for _, option := range v.Options {
				var variantOption models.ProductOptionValue
				for _, o := range variant.Options {
					if option.Id == o.Id {
						variantOption = o
						break
					}
				}
				if option.Value == variantOption.Value {
					variantName = v.Title
					return true
				}
			}

			return false
		})

		if variantExists {
			return utils.NewApplictaionError(
				utils.DUPLICATE_ERROR,
				"Variant with title "+variantName+" with provided options already exists",
				"500",
				nil,
			)
		}
	}
	return nil
}

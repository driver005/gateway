package services

import (
	"context"
	"fmt"
	"reflect"
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

func (s *ProductVariantService) Retrieve(id uuid.UUID, config *sql.Options) (*models.ProductVariant, *utils.ApplictaionError) {
	var res *models.ProductVariant

	query := sql.BuildQuery(models.ProductVariant{Model: core.Model{Id: id}}, config)

	if err := s.r.ProductVariantRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ProductVariantService) RetrieveBySKU(sku string, config *sql.Options) (*models.ProductVariant, *utils.ApplictaionError) {
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

func (s *ProductVariantService) Create(productId uuid.UUID, product *models.Product, variants []types.CreateProductVariantInput) ([]models.ProductVariant, *utils.ApplictaionError) {
	if productId != uuid.Nil {
		p, err := s.r.ProductService().SetContext(s.ctx).RetrieveById(productId, &sql.Options{Relations: []string{"variants", "variants.options", "options"}})
		if err != nil {
			return nil, err
		}
		product = p
	}
	if product == nil || product.Id == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Product id missing",
			nil,
		)
	}
	err := s.ValidateVariantsToCreate(product, variants)
	if err != nil {
		return nil, err
	}

	computedRank := len(product.Variants)
	var variantPricesToUpdate []types.UpdateVariantPricesData
	var results []models.ProductVariant
	for _, data := range variants {
		data.VariantRank = computedRank
		data.ProductId = product.Id

		model := &models.ProductVariant{
			Model: core.Model{
				Metadata: data.Metadata,
			},
			Title:             data.Title,
			ProductId:         uuid.NullUUID{UUID: data.ProductId},
			Sku:               data.SKU,
			Barcode:           data.Barcode,
			Ean:               data.EAN,
			Upc:               data.UPC,
			VariantRank:       data.VariantRank,
			InventoryQuantity: data.InventoryQuantity,
			AllowBackorder:    data.AllowBackorder,
			ManageInventory:   data.ManageInventory,
			HsCode:            data.HSCode,
			OriginCountry:     data.OriginCountry,
			MIdCode:           data.MIdCode,
			Material:          data.Material,
			Weight:            data.Weight,
			Length:            data.Length,
			Height:            data.Height,
			Width:             data.Width,
		}

		if err := s.r.ProductVariantRepository().Save(s.ctx, model); err != nil {
			return nil, err
		}

		if len(data.Prices) > 0 {
			variantPricesToUpdate = append(variantPricesToUpdate, types.UpdateVariantPricesData{
				VariantId: model.Id,
				Prices:    data.Prices,
			})
		}
		results = append(results, *model)
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

func (s *ProductVariantService) UpdateBatch(data []types.UpdateProductVariantData) ([]models.ProductVariant, *utils.ApplictaionError) {
	var res []models.ProductVariant
	var prices []types.UpdateVariantPricesData
	for _, d := range data {
		prices = append(prices, types.UpdateVariantPricesData{
			VariantId: d.Variant.Id,
			Prices:    d.UpdateData.Prices,
		})
	}

	if len(prices) > 0 {
		s.UpdateVariantPrices(prices)
	}

	for _, d := range data {
		for _, option := range d.UpdateData.Options {
			s.UpdateOptionValue(d.Variant.Id, option.OptionId, option.Value)
		}

		if !reflect.ValueOf(d.Variant.Metadata).IsZero() {
			d.Variant.Metadata = utils.MergeMaps(d.Variant.Metadata, d.UpdateData.Metadata)
		}
		if !reflect.ValueOf(d.Variant.Title).IsZero() {
			d.Variant.Title = d.UpdateData.Title
		}
		if !reflect.ValueOf(d.Variant.ProductId).IsZero() {
			d.Variant.ProductId = uuid.NullUUID{UUID: d.UpdateData.ProductId}
		}
		if !reflect.ValueOf(d.Variant.Sku).IsZero() {
			d.Variant.Sku = d.UpdateData.SKU
		}
		if !reflect.ValueOf(d.Variant.Barcode).IsZero() {
			d.Variant.Barcode = d.UpdateData.Barcode
		}
		if !reflect.ValueOf(d.Variant.Ean).IsZero() {
			d.Variant.Ean = d.UpdateData.EAN
		}
		if !reflect.ValueOf(d.Variant.Upc).IsZero() {
			d.Variant.Upc = d.UpdateData.UPC
		}
		if !reflect.ValueOf(d.Variant.VariantRank).IsZero() {
			d.Variant.VariantRank = d.UpdateData.VariantRank
		}
		if !reflect.ValueOf(d.Variant.InventoryQuantity).IsZero() {
			d.Variant.InventoryQuantity = d.UpdateData.InventoryQuantity
		}
		if !reflect.ValueOf(d.Variant.AllowBackorder).IsZero() {
			d.Variant.AllowBackorder = d.UpdateData.AllowBackorder
		}
		if !reflect.ValueOf(d.Variant.ManageInventory).IsZero() {
			d.Variant.ManageInventory = d.UpdateData.ManageInventory
		}
		if !reflect.ValueOf(d.Variant.HsCode).IsZero() {
			d.Variant.HsCode = d.UpdateData.HSCode
		}
		if !reflect.ValueOf(d.Variant.OriginCountry).IsZero() {
			d.Variant.OriginCountry = d.UpdateData.OriginCountry
		}
		if !reflect.ValueOf(d.Variant.MIdCode).IsZero() {
			d.Variant.MIdCode = d.UpdateData.MIdCode
		}
		if !reflect.ValueOf(d.Variant.Material).IsZero() {
			d.Variant.Material = d.UpdateData.Material
		}
		if !reflect.ValueOf(d.Variant.Weight).IsZero() {
			d.Variant.Weight = d.UpdateData.Weight
		}
		if !reflect.ValueOf(d.Variant.Length).IsZero() {
			d.Variant.Length = d.UpdateData.Length
		}
		if !reflect.ValueOf(d.Variant.Height).IsZero() {
			d.Variant.Height = d.UpdateData.Height
		}
		if !reflect.ValueOf(d.Variant.Width).IsZero() {
			d.Variant.Width = d.UpdateData.Width
		}

		if err := s.r.ProductVariantRepository().Update(s.ctx, d.Variant); err != nil {
			return nil, err
		}

		res = append(res, *d.Variant)
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

	return res, nil
}

func (s *ProductVariantService) Update(id uuid.UUID, variant *models.ProductVariant, data *types.UpdateProductVariantInput) (*models.ProductVariant, *utils.ApplictaionError) {
	model := variant
	if id == uuid.Nil {
		mod, err := s.Retrieve(id, &sql.Options{})
		if err != nil {
			return nil, err
		}
		model = mod
	}

	var update []types.UpdateProductVariantData

	if data != nil {
		if model.Id == uuid.Nil {
			return nil, utils.NewApplictaionError(
				utils.INVALID_DATA,
				"Variant id missing",
				nil,
			)
		}

		update = []types.UpdateProductVariantData{
			{
				Variant:    model,
				UpdateData: data,
			},
		}
	}

	res, err := s.UpdateBatch(update)
	if err != nil {
		return nil, err
	}
	return &res[0], nil
}

// func (s *ProductVariantService) UpdateVariantPrices(id uuid.UUID, variant types.UpdateVariantPricesData, data []types.ProductVariantPrice) *utils.ApplictaionError {
// 	return nil
// }

func (s *ProductVariantService) UpdateVariantPrices(data []types.UpdateVariantPricesData) *utils.ApplictaionError {
	var prices []models.MoneyAmount
	var regionIds uuid.UUIDs
	for _, d := range data {
		for _, p := range d.Prices {
			prices = append(prices, models.MoneyAmount{
				Model: core.Model{
					Id: p.Id,
				},
				CurrencyCode: p.CurrencyCode,
				RegionId:     uuid.NullUUID{UUID: p.RegionId},
				Amount:       p.Amount,
				MinQuantity:  p.MinQuantity,
				MaxQuantity:  p.MaxQuantity,
			})
			if p.RegionId != uuid.Nil {
				regionIds = append(regionIds, p.RegionId)
			}
		}
	}
	s.r.MoneyAmountRepository().DeleteVariantPricesNotIn(uuid.Nil, prices)

	regions, err := s.r.RegionService().SetContext(s.ctx).List(models.Region{}, &sql.Options{
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
	var dataRegionPrices []types.UpdateVariantRegionPriceData
	var dataCurrencyPrices []types.UpdateVariantCurrencyPriceData
	for _, d := range data {
		for _, p := range d.Prices {
			if p.RegionId != uuid.Nil {
				region := regionsMap[p.RegionId]
				dataRegionPrices = append(dataRegionPrices, types.UpdateVariantRegionPriceData{
					VariantId: d.VariantId,
					Price: &types.ProductVariantPrice{
						CurrencyCode: region.CurrencyCode,
						RegionId:     p.RegionId,
						Amount:       p.Amount,
					},
				})
			} else {
				dataCurrencyPrices = append(dataCurrencyPrices, types.UpdateVariantCurrencyPriceData{
					VariantId: d.VariantId,
					Price: &types.ProductVariantPrice{
						CurrencyCode: p.CurrencyCode,
						Amount:       p.Amount,
					},
				})
			}
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

func (s *ProductVariantService) UpsertRegionPrices(data []types.UpdateVariantRegionPriceData) *utils.ApplictaionError {
	var where []models.MoneyAmount
	for _, d := range data {
		where = append(where, models.MoneyAmount{
			VariantId: uuid.NullUUID{UUID: d.VariantId},
			RegionId:  uuid.NullUUID{UUID: d.Price.RegionId},
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
		variantMoneyAmounts := moneyAmountsMapToVariantId[d.VariantId]
		var moneyAmount *models.MoneyAmount
		for _, ma := range variantMoneyAmounts {
			if ma.RegionId.UUID == d.Price.RegionId {
				moneyAmount = &ma
				break
			}
		}
		if moneyAmount != nil {
			if moneyAmount.Amount != d.Price.Amount {
				dataToUpdate = append(dataToUpdate, models.MoneyAmount{
					Model: core.Model{
						Id: moneyAmount.Id,
					},
					Amount: d.Price.Amount,
				})
			}
		} else {
			model := &models.MoneyAmount{
				Model: core.Model{
					Id: d.Price.Id,
				},
				CurrencyCode: d.Price.CurrencyCode,
				RegionId:     uuid.NullUUID{UUID: d.Price.RegionId},
				Amount:       d.Price.Amount,
				MinQuantity:  d.Price.MinQuantity,
				MaxQuantity:  d.Price.MaxQuantity,
			}
			if err := s.r.MoneyAmountRepository().Create(s.ctx, model); err != nil {
				return err
			}
			model.Variant = &models.ProductVariant{Model: core.Model{Id: d.VariantId}}
			dataToCreate = append(dataToCreate, *model)
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
			variantIds = append(variantIds, d.VariantId)
		}
		s.r.PriceSelectionStrategy().OnVariantsPricesUpdate(variantIds)
	}
	return nil
}

func (s *ProductVariantService) UpsertCurrencyPrices(data []types.UpdateVariantCurrencyPriceData) *utils.ApplictaionError {
	var where []models.MoneyAmount
	for _, d := range data {
		where = append(where, models.MoneyAmount{
			VariantId:    uuid.NullUUID{UUID: d.VariantId},
			CurrencyCode: d.Price.CurrencyCode,
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
		variantMoneyAmounts := moneyAmountsMapToVariantId[d.VariantId]
		var moneyAmount *models.MoneyAmount
		for _, ma := range variantMoneyAmounts {
			if ma.CurrencyCode == d.Price.CurrencyCode {
				moneyAmount = &ma
				break
			}
		}
		if moneyAmount != nil {
			if moneyAmount.Amount != d.Price.Amount {
				dataToUpdate = append(dataToUpdate, models.MoneyAmount{
					Model: core.Model{
						Id: moneyAmount.Id,
					},
					Amount: d.Price.Amount,
				})
			}
		} else {
			model := &models.MoneyAmount{
				Model: core.Model{
					Id: d.Price.Id,
				},
				CurrencyCode: d.Price.CurrencyCode,
				RegionId:     uuid.NullUUID{UUID: d.Price.RegionId},
				Amount:       d.Price.Amount,
				MinQuantity:  d.Price.MinQuantity,
				MaxQuantity:  d.Price.MaxQuantity,
			}
			if err := s.r.MoneyAmountRepository().Create(s.ctx, model); err != nil {
				return err
			}
			model.Variant = &models.ProductVariant{Model: core.Model{Id: d.VariantId}}
			dataToCreate = append(dataToCreate, *model)
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
			variantIds = append(variantIds, d.VariantId)
		}
		s.r.PriceSelectionStrategy().OnVariantsPricesUpdate(variantIds)
	}
	return nil
}

// getRegionPrice gets the price specific to a region
func (s *ProductVariantService) GetRegionPrice(id uuid.UUID, data *types.GetRegionPriceContext) (*float64, *utils.ApplictaionError) {
	region, err := s.r.RegionService().SetContext(s.ctx).Retrieve(id, &sql.Options{})
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
func (s *ProductVariantService) SetRegionPrice(id uuid.UUID, price types.ProductVariantPrice) (*models.MoneyAmount, *utils.ApplictaionError) {
	var data *models.MoneyAmount
	moneyAmount, err := s.r.MoneyAmountRepository().GetPricesForVariantInRegion(id, price.RegionId)
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
func (s *ProductVariantService) SetCurrencyPrice(id uuid.UUID, data types.ProductVariantPrice) (*models.MoneyAmount, *utils.ApplictaionError) {
	return s.r.MoneyAmountRepository().UpsertVariantCurrencyPrice(id, models.MoneyAmount{
		Model: core.Model{
			Id: data.Id,
		},
		CurrencyCode: data.CurrencyCode,
		RegionId:     uuid.NullUUID{UUID: data.RegionId},
		Amount:       data.Amount,
		MinQuantity:  data.MinQuantity,
		MaxQuantity:  data.MaxQuantity,
	})
}

// updateOptionValue updates variant's option value
func (s *ProductVariantService) UpdateOptionValue(id uuid.UUID, optionId uuid.UUID, optionValue string) (*models.ProductOptionValue, *utils.ApplictaionError) {
	var res *models.ProductOptionValue

	query := sql.BuildQuery(models.ProductOptionValue{VariantId: uuid.NullUUID{UUID: id}, OptionId: uuid.NullUUID{UUID: optionId}}, &sql.Options{})

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

	query := sql.BuildQuery(models.ProductOptionValue{VariantId: uuid.NullUUID{UUID: id}, OptionId: uuid.NullUUID{UUID: optionId}}, &sql.Options{})

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
func (s *ProductVariantService) ListAndCount(selector types.FilterableProductVariant, config *sql.Options, q *string) ([]models.ProductVariant, *int64, *utils.ApplictaionError) {
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
func (s *ProductVariantService) List(selector types.FilterableProductVariant, config *sql.Options, q *string) ([]models.ProductVariant, *utils.ApplictaionError) {
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
		variant, err := s.Retrieve(id, &sql.Options{})
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
	variant, err := s.Retrieve(id, &sql.Options{
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
func (s *ProductVariantService) ValidateVariantsToCreate(product *models.Product, variants []types.CreateProductVariantInput) *utils.ApplictaionError {
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
			if !slices.ContainsFunc(variant.Options, func(v types.ProductVariantOption) bool {
				return option.Id == v.OptionId
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
					if option.Id == o.OptionId {
						variantOption = models.ProductOptionValue{
							OptionId: uuid.NullUUID{UUID: o.OptionId},
							Value:    o.Value,
						}
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

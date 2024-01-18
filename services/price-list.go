package services

import (
	"context"
	"reflect"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/icza/gox/gox"
)

type PriceListService struct {
	ctx context.Context
	r   Registry
}

func NewPriceListService(
	r Registry,
) *PriceListService {
	return &PriceListService{
		context.Background(),
		r,
	}
}

func (s *PriceListService) SetContext(context context.Context) *PriceListService {
	s.ctx = context
	return s
}

func (s *PriceListService) Retrieve(id uuid.UUID, config *sql.Options) (*models.PriceList, *utils.ApplictaionError) {
	if id == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_FOUND,
			`"id" must be defined`,
			nil,
		)
	}

	var res *models.PriceList
	query := sql.BuildQuery(models.OAuth{Model: core.Model{Id: id}}, config)

	if err := s.r.PriceListRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PriceListService) List(selector types.FilterablePriceList, config *sql.Options) ([]models.PriceList, *utils.ApplictaionError) {
	res, _, err := s.ListAndCount(selector, config)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PriceListService) ListAndCount(selector types.FilterablePriceList, config *sql.Options) ([]models.PriceList, *int64, *utils.ApplictaionError) {
	if reflect.DeepEqual(config, &sql.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(50)
		config.Order = gox.NewString("created_at DESC")
	}

	return s.r.PriceListRepository().ListAndCount(s.ctx, selector, config, config.Q)
}

func (s *PriceListService) ListPriceListsVariantIdsMap(priceListIds uuid.UUIDs) (map[string][]string, *utils.ApplictaionError) {
	if len(priceListIds) == 0 {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"priceListIds" must be defined`,
			nil,
		)
	}
	priceListsVariantIdsMap, err := s.r.PriceListRepository().ListPriceListsVariantIdsMap(priceListIds)
	if err != nil {
		return nil, err
	}
	return priceListsVariantIdsMap, nil
}

func (s *PriceListService) Create(data *types.CreatePriceListInput) (*models.PriceList, *utils.ApplictaionError) {
	feature := true

	model := &models.PriceList{
		Name:        data.Name,
		Description: data.Description,
		Type:        data.Type,
		Status:      data.Status,
		StartsAt:    data.StartsAt,
		EndsAt:      data.EndsAt,
		IncludesTax: data.IncludesTax,
	}

	if !feature {
		model.IncludesTax = false
	}

	if err := s.r.PriceListRepository().Save(s.ctx, model); err != nil {
		return nil, err
	}

	if data.Prices != nil {
		prices, err := s.AddCurrencyFromRegion(data.Prices)
		if err != nil {
			return nil, err
		}
		_, err = s.r.MoneyAmountRepository().AddPriceListPrices(model.Id, prices, false)
		if err != nil {
			return nil, err
		}
	}
	if data.CustomerGroups != nil {
		err := s.UpsertCustomerGroups(model.Id, data.CustomerGroups)
		if err != nil {
			return nil, err
		}
	}
	return s.Retrieve(model.Id, &sql.Options{Relations: []string{"prices", "customer_groups"}})
}

func (s *PriceListService) Update(id uuid.UUID, data *types.UpdatePriceListInput) (*models.PriceList, *utils.ApplictaionError) {
	priceList, err := s.Retrieve(id, &sql.Options{Selects: []string{"id"}})
	if err != nil {
		return nil, err
	}
	feature := true
	if !feature {
		priceList.IncludesTax = false
	}
	if data.Prices != nil {
		prices, err := s.AddCurrencyFromRegion(data.Prices)
		if err != nil {
			return nil, err
		}
		_, err = s.r.MoneyAmountRepository().AddPriceListPrices(priceList.Id, prices, false)
		if err != nil {
			return nil, err
		}
	}
	if data.CustomerGroups != nil {
		err := s.UpsertCustomerGroups(priceList.Id, data.CustomerGroups)
		if err != nil {
			return nil, err
		}
	}

	if !reflect.ValueOf(data.Name).IsZero() {
		priceList.Name = data.Name
	}
	if !reflect.ValueOf(data.Description).IsZero() {
		priceList.Description = data.Description
	}
	if !reflect.ValueOf(data.StartsAt).IsZero() {
		priceList.StartsAt = data.StartsAt
	}
	if !reflect.ValueOf(data.EndsAt).IsZero() {
		priceList.EndsAt = data.EndsAt
	}
	if !reflect.ValueOf(data.Status).IsZero() {
		priceList.Status = data.Status
	}
	if !reflect.ValueOf(data.Type).IsZero() {
		priceList.Type = data.Type
	}
	if !reflect.ValueOf(data.IncludesTax).IsZero() {
		priceList.IncludesTax = data.IncludesTax
	}

	if err := s.r.PriceListRepository().Save(s.ctx, priceList); err != nil {
		return nil, err
	}
	return s.Retrieve(priceList.Id, &sql.Options{Relations: []string{"prices", "customer_groups"}})
}

func (s *PriceListService) Delete(id uuid.UUID) *utils.ApplictaionError {
	priceList, err := s.Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}
	if priceList == nil {
		return nil
	}
	if err := s.r.PriceListRepository().Remove(s.ctx, priceList); err != nil {
		return err
	}
	return nil
}

func (s *PriceListService) AddPrices(id uuid.UUID, prices []types.PriceListPriceCreateInput, replace bool) (*models.PriceList, *utils.ApplictaionError) {
	priceList, err := s.Retrieve(id, &sql.Options{Selects: []string{"id"}})
	if err != nil {
		return nil, err
	}
	price, err := s.AddCurrencyFromRegion(prices)
	if err != nil {
		return nil, err
	}
	_, err = s.r.MoneyAmountRepository().AddPriceListPrices(priceList.Id, price, replace)
	if err != nil {
		return nil, err
	}
	return s.Retrieve(priceList.Id, &sql.Options{Relations: []string{"prices"}})
}

func (s *PriceListService) DeletePrices(id uuid.UUID, priceIds uuid.UUIDs) *utils.ApplictaionError {
	priceList, err := s.Retrieve(id, &sql.Options{Selects: []string{"id"}})
	if err != nil {
		return err
	}
	err = s.r.MoneyAmountRepository().DeletePriceListPrices(priceList.Id, priceIds)
	if err != nil {
		return err
	}
	return nil
}

func (s *PriceListService) ClearPrices(id uuid.UUID) *utils.ApplictaionError {
	priceList, err := s.Retrieve(id, &sql.Options{Selects: []string{"id"}})
	if err != nil {
		return err
	}
	if err = s.r.MoneyAmountRepository().Delete(s.ctx, &models.MoneyAmount{PriceListId: uuid.NullUUID{UUID: priceList.Id}}); err != nil {
		return err
	}
	return nil
}

func (s *PriceListService) UpsertCustomerGroups(id uuid.UUID, customerGroups []types.CustomerGroups) *utils.ApplictaionError {
	priceList, err := s.Retrieve(id, &sql.Options{Selects: []string{"id"}})
	if err != nil {
		return err
	}
	var groups []models.CustomerGroup
	for _, cg := range customerGroups {
		customerGroup, err := s.r.CustomerGroupService().Retrieve(cg.Id, &sql.Options{})
		if err != nil {
			return err
		}
		groups = append(groups, *customerGroup)
	}

	priceList.CustomerGroups = groups

	if err = s.r.PriceListRepository().Save(s.ctx, priceList); err != nil {
		return err
	}
	return nil
}

func (s *PriceListService) ListProducts(id uuid.UUID, selector types.FilterableProduct, config *sql.Options, requiresPriceList bool) ([]models.Product, *int64, *utils.ApplictaionError) {
	products, count, err := s.r.ProductService().SetContext(s.ctx).ListAndCount(selector, config)
	if err != nil {
		return nil, nil, err
	}

	productsWithPrices := []models.Product{}
	for _, p := range products {
		if len(p.Variants) > 0 {
			var variants []models.ProductVariant
			for _, v := range p.Variants {
				prices, _, err := s.r.MoneyAmountRepository().FindManyForVariantInPriceList(v.Id, id, requiresPriceList)
				if err != nil {
					return nil, nil, err
				}
				v.Prices = prices
				variants = append(variants, v)
			}
			p.Variants = variants
		}
		productsWithPrices = append(productsWithPrices, p)
	}
	return productsWithPrices, count, nil
}

func (s *PriceListService) ListVariants(id uuid.UUID, selector types.FilterableProductVariant, config *sql.Options, requiresPriceList bool) ([]models.ProductVariant, *int64, *utils.ApplictaionError) {
	variants, count, err := s.r.ProductVariantService().ListAndCount(selector, config)
	if err != nil {
		return nil, nil, err
	}

	var variantsWithPrices []models.ProductVariant
	for _, variant := range variants {
		prices, _, err := s.r.MoneyAmountRepository().FindManyForVariantInPriceList(variant.Id, id, requiresPriceList)
		if err != nil {
			return nil, nil, err
		}
		variant.Prices = prices
		variantsWithPrices = append(variantsWithPrices, variant)
	}
	return variantsWithPrices, count, nil
}

func (s *PriceListService) DeleteProductPrices(id uuid.UUID, productIds uuid.UUIDs) (uuid.UUIDs, *int, *utils.ApplictaionError) {
	products, count, err := s.ListProducts(id, types.FilterableProduct{FilterModel: core.FilterModel{Id: productIds}}, &sql.Options{Relations: []string{"variants"}}, true)
	if err != nil {
		return nil, nil, err
	}
	if count == nil {
		return nil, nil, nil
	}
	var priceIds uuid.UUIDs
	for _, product := range products {
		for _, variant := range product.Variants {
			for _, price := range variant.Prices {
				priceIds = append(priceIds, price.Id)
			}
		}
	}
	if len(priceIds) == 0 {
		return nil, nil, nil
	}
	if err = s.DeletePrices(id, priceIds); err != nil {
		return nil, nil, err
	}
	return priceIds, gox.NewInt(len(priceIds)), nil
}

func (s *PriceListService) DeleteVariantPrices(id uuid.UUID, variantIds uuid.UUIDs) (uuid.UUIDs, *int, *utils.ApplictaionError) {
	variants, count, err := s.ListVariants(id, types.FilterableProductVariant{FilterModel: core.FilterModel{Id: variantIds}}, &sql.Options{}, true)
	if err != nil {
		return nil, nil, err
	}
	if count == nil {
		return nil, nil, nil
	}
	var priceIds uuid.UUIDs
	for _, variant := range variants {
		for _, price := range variant.Prices {
			priceIds = append(priceIds, price.Id)
		}
	}
	if len(priceIds) == 0 {
		return nil, nil, nil
	}
	if err = s.DeletePrices(id, priceIds); err != nil {
		return nil, nil, err
	}
	return priceIds, gox.NewInt(len(priceIds)), nil
}

func (s *PriceListService) AddCurrencyFromRegion(prices []types.PriceListPriceCreateInput) ([]models.MoneyAmount, *utils.ApplictaionError) {
	var regionIds uuid.UUIDs
	var pricesData []models.MoneyAmount
	for _, p := range prices {
		regionIds = append(regionIds, p.RegionId)
	}
	regions, err := s.r.RegionService().SetContext(s.ctx).List(models.Region{}, &sql.Options{Specification: []sql.Specification{sql.In("id", regionIds)}})
	if err != nil {
		return nil, err
	}
	regionsMap := make(map[uuid.UUID]models.Region)
	for _, r := range regions {
		regionsMap[r.Id] = r
	}
	for _, price := range prices {
		if price.RegionId != uuid.Nil {
			region := regionsMap[price.RegionId]
			price.CurrencyCode = region.CurrencyCode
		}
		pricesData = append(pricesData, models.MoneyAmount{
			RegionId:     uuid.NullUUID{UUID: price.RegionId},
			CurrencyCode: price.CurrencyCode,
			VariantId:    uuid.NullUUID{UUID: price.VariantId},
			Amount:       price.Amount,
			MinQuantity:  price.MinQuantity,
			MaxQuantity:  price.MaxQuantity,
		})
	}
	return pricesData, nil
}

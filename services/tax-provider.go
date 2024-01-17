package services

import (
	"context"
	"fmt"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/sarulabs/di"
)

type TaxProviderService struct {
	ctx       context.Context
	container di.Container
	r         Registry
}

func NewTaxProviderService(
	container di.Container,
	r Registry,
) *TaxProviderService {
	return &TaxProviderService{
		context.Background(),
		container,
		r,
	}
}

func (s *TaxProviderService) SetContext(context context.Context) *TaxProviderService {
	s.ctx = context
	return s
}

func (s *TaxProviderService) RetrieveProvider(region models.Region) interfaces.ITaxService {
	var provider interfaces.ITaxService
	if region.TaxProviderId.UUID != uuid.Nil {
		objectInterface, _ := s.container.SafeGet("tp_" + region.TaxProviderId.UUID.String())
		provider, _ = objectInterface.(interfaces.ITaxService)
	} else {
		objectInterface, _ := s.container.SafeGet("systemTaxService")
		provider, _ = objectInterface.(interfaces.ITaxService)
	}

	return provider
}

func (s *TaxProviderService) ClearLineItemsTaxLines(itemIds uuid.UUIDs) *utils.ApplictaionError {
	if err := s.r.LineItemTaxLineRepository().Specification(sql.In[uuid.UUID]("item_id", itemIds)).Delete(s.ctx, &models.LineItemTaxLine{}); err != nil {
		return err
	}

	return nil
}

func (s *TaxProviderService) ClearTaxLines(cartId uuid.UUID) *utils.ApplictaionError {
	if err := s.r.LineItemTaxLineRepository().DeleteForCart(s.ctx, cartId); err != nil {
		return err
	}

	if err := s.r.ShippingMethodTaxLineRepository().DeleteForCart(s.ctx, cartId); err != nil {
		return err
	}

	return nil
}

func (s *TaxProviderService) CreateTaxLines(cart *models.Cart, lineItems []models.LineItem, calculationContext *interfaces.TaxCalculationContext) ([]models.ShippingMethodTaxLine, []models.LineItemTaxLine, *utils.ApplictaionError) {
	var taxline []models.LineItemTaxLine
	var smTaxLine []models.ShippingMethodTaxLine
	var err *utils.ApplictaionError

	if cart != nil {
		smTaxLine, taxline, err = s.GetTaxLines(cart.Items, calculationContext)
	} else {
		smTaxLine, taxline, err = s.GetTaxLines(lineItems, calculationContext)
	}
	if err != nil {
		return nil, nil, err
	}

	if err := s.r.LineItemTaxLineRepository().SaveSlice(s.ctx, taxline); err != nil {
		return nil, nil, err
	}

	if err := s.r.ShippingMethodTaxLineRepository().SaveSlice(s.ctx, smTaxLine); err != nil {
		return nil, nil, err
	}

	return smTaxLine, taxline, err
}

func (s *TaxProviderService) CreateShippingTaxLines(shippingMethod *models.ShippingMethod, calculationContext *interfaces.TaxCalculationContext) ([]models.ShippingMethodTaxLine, *utils.ApplictaionError) {
	taxline, err := s.GetShippingTaxLines(shippingMethod, calculationContext)
	if err != nil {
		return nil, err
	}

	if err := s.r.ShippingMethodTaxLineRepository().SaveSlice(s.ctx, taxline); err != nil {
		return nil, err
	}

	return taxline, nil
}

func (s *TaxProviderService) GetShippingTaxLines(shippingMethod *models.ShippingMethod, calculationContext *interfaces.TaxCalculationContext) ([]models.ShippingMethodTaxLine, *utils.ApplictaionError) {
	rates, err := s.GetRegionRatesForShipping(shippingMethod.Id, &calculationContext.Region)
	if err != nil {
		return nil, err
	}

	calculationLines := []interfaces.ShippingTaxCalculationLine{
		{
			ShippingMethod: models.ShippingMethod{
				Model: core.Model{
					Id: shippingMethod.Id,
				},
			},
			Rates: rates,
		},
	}
	taxProvider := s.RetrieveProvider(calculationContext.Region)
	providerLines, err := taxProvider.GetTaxLines([]interfaces.ItemTaxCalculationLine{}, calculationLines, calculationContext)
	if err != nil {
		return nil, err
	}

	var result []models.ShippingMethodTaxLine

	for _, pl := range providerLines {
		smTaxLine := models.ShippingMethodTaxLine{
			Model: core.Model{
				Metadata: pl.Metadata,
			},
			ShippingMethodId: uuid.NullUUID{
				UUID:  pl.ShippingMethodId,
				Valid: true,
			},
			Rate: pl.Rate,
			Name: pl.Name,
			Code: pl.Code,
		}
		result = append(result, smTaxLine)
	}
	return result, nil
}

func (s *TaxProviderService) GetTaxLines(lineItems []models.LineItem, calculationContext *interfaces.TaxCalculationContext) ([]models.ShippingMethodTaxLine, []models.LineItemTaxLine, *utils.ApplictaionError) {
	var productIds uuid.UUIDs

	for _, item := range lineItems {
		if item.Variant != nil && item.Variant.Id != uuid.Nil {
			productIds = append(productIds, item.Variant.Id)
		}
	}

	productRatesMap, err := s.GetRegionRatesForProduct(productIds, &calculationContext.Region)
	if err != nil {
		return nil, nil, err
	}

	var calculationLines []interfaces.ItemTaxCalculationLine

	for _, item := range lineItems {
		if item.IsReturn {
			continue
		}

		var rates []types.TaxServiceRate
		if item.Variant != nil && item.Variant.Id != uuid.Nil {
			rates = productRatesMap[item.Variant.Id]
		}

		calculationLines = append(calculationLines, interfaces.ItemTaxCalculationLine{
			Item:  item,
			Rates: rates,
		})
	}

	var shippingCalculationLines []interfaces.ShippingTaxCalculationLine

	for _, sm := range calculationContext.ShippingMethods {
		rates, err := s.GetRegionRatesForShipping(sm.ShippingOption.Id, &calculationContext.Region)
		if err != nil {
			return nil, nil, err
		}

		shippingCalculationLines = append(shippingCalculationLines, interfaces.ShippingTaxCalculationLine{
			ShippingMethod: sm,
			Rates:          rates,
		})
	}

	taxProvider := s.RetrieveProvider(calculationContext.Region)
	providerLines, err := taxProvider.GetTaxLines(calculationLines, shippingCalculationLines, calculationContext)
	if err != nil {
		return nil, nil, err
	}

	var smTaxLines []models.ShippingMethodTaxLine
	var liTaxLines []models.LineItemTaxLine

	for _, pl := range providerLines {
		if pl.ShippingMethodId != uuid.Nil {
			smTaxLines = append(smTaxLines, models.ShippingMethodTaxLine{
				Model: core.Model{
					Metadata: pl.Metadata,
				},
				ShippingMethodId: uuid.NullUUID{UUID: pl.ShippingMethodId},
				Rate:             pl.Rate,
				Name:             pl.Name,
				Code:             pl.Code,
			})
		} else if pl.ItemId == uuid.Nil {
			return nil, nil, utils.NewApplictaionError(
				utils.INVALID_DATA,
				"Tax Provider returned invalid tax lines",
				"500",
				nil,
			)
		} else {
			liTaxLines = append(liTaxLines, models.LineItemTaxLine{
				Model: core.Model{
					Metadata: pl.Metadata,
				},
				ItemId: uuid.NullUUID{UUID: pl.ItemId},
				Rate:   pl.Rate,
				Name:   pl.Name,
				Code:   pl.Code,
			})
		}
	}

	return smTaxLines, liTaxLines, nil
}

func (s *TaxProviderService) GetTaxLinesMap(items []models.LineItem, calculationContext *interfaces.TaxCalculationContext) (types.TaxLinesMaps, *utils.ApplictaionError) {
	lineItemsTaxLinesMap := make(map[uuid.UUID][]models.LineItemTaxLine)
	shippingMethodsTaxLinesMap := make(map[uuid.UUID][]models.ShippingMethodTaxLine)

	shippingMethodTaxLines, taxLines, err := s.GetTaxLines(items, calculationContext)
	if err != nil {
		return types.TaxLinesMaps{}, err
	}

	for _, taxLine := range taxLines {
		if taxLine.Id != uuid.Nil {
			itemTaxLines := lineItemsTaxLinesMap[taxLine.Id]
			itemTaxLines = append(itemTaxLines, taxLine)
			lineItemsTaxLinesMap[taxLine.Id] = itemTaxLines
		}
	}

	for _, shippingMethodTaxLine := range shippingMethodTaxLines {
		if shippingMethodTaxLine.Id != uuid.Nil {
			shippingMethodTaxLines := shippingMethodsTaxLinesMap[shippingMethodTaxLine.Id]
			shippingMethodTaxLines = append(shippingMethodTaxLines, shippingMethodTaxLine)
			shippingMethodsTaxLinesMap[shippingMethodTaxLine.Id] = shippingMethodTaxLines
		}
	}

	return types.TaxLinesMaps{
		LineItemsTaxLines:       lineItemsTaxLinesMap,
		ShippingMethodsTaxLines: shippingMethodsTaxLinesMap,
	}, nil
}

func (s *TaxProviderService) GetRegionRatesForShipping(optionId uuid.UUID, region *models.Region) ([]types.TaxServiceRate, *utils.ApplictaionError) {
	cachKey := s.GetCacheKey(optionId, region.Id)
	cacheHit, er := s.r.CacheService().Get(cachKey)
	if er != nil {
		return nil, utils.NewApplictaionError(
			utils.UNEXPECTED_STATE,
			er.Error(),
			nil,
		)
	}

	if cacheHit != nil {
		return cacheHit.([]types.TaxServiceRate), nil
	}

	var toReturn []types.TaxServiceRate
	optionRates, err := s.r.TaxRateService().SetContext(s.ctx).ListByShippingOption(optionId)
	if err != nil {
		return nil, err
	}

	if len(optionRates) == 0 {
		toReturn = []types.TaxServiceRate{
			{
				Rate: region.TaxRate,
				Name: "default",
				Code: "default",
			},
		}
	} else {
		for _, rate := range optionRates {
			toReturn = append(toReturn, types.TaxServiceRate{
				Rate: rate.Rate,
				Name: rate.Name,
				Code: rate.Code,
			})
		}
	}

	s.r.CacheService().Set(cachKey, toReturn, nil)
	return toReturn, nil
}

func (s *TaxProviderService) GetRegionRatesForProduct(productIds uuid.UUIDs, region *models.Region) (map[uuid.UUID][]types.TaxServiceRate, *utils.ApplictaionError) {
	var nonCachedProductIds uuid.UUIDs
	cacheKeysMap := make(map[uuid.UUID]string)
	productRatesMapResult := make(map[uuid.UUID][]types.TaxServiceRate)

	for _, p := range productIds {
		cacheKeysMap[p] = s.GetCacheKey(p, region.Id)
	}

	for key, value := range cacheKeysMap {
		cacheHit, err := s.r.CacheService().Get(value)
		if err != nil {
			return nil, utils.NewApplictaionError(
				utils.UNEXPECTED_STATE,
				err.Error(),
				"500",
				nil,
			)
		}

		if cacheHit == nil {
			nonCachedProductIds = append(nonCachedProductIds, key)
		} else {
			productRatesMapResult[key] = cacheHit.([]types.TaxServiceRate)
		}
	}

	if len(nonCachedProductIds) == 0 {
		return productRatesMapResult, nil
	}

	for _, value := range nonCachedProductIds {
		var toReturn []types.TaxServiceRate
		rates, err := s.r.TaxRateService().SetContext(s.ctx).ListByProduct(value, types.TaxRateListByConfig{RegionId: region.Id})
		if err != nil {
			return nil, err
		}

		if len(rates) == 0 {
			toReturn = []types.TaxServiceRate{
				{
					Rate: region.TaxRate,
					Name: "default",
					Code: "default",
				},
			}
		} else {
			for _, rate := range rates {
				toReturn = append(toReturn, types.TaxServiceRate{
					Rate: rate.Rate,
					Name: rate.Name,
					Code: rate.Code,
				})
			}
		}

		s.r.CacheService().Set(cacheKeysMap[value], toReturn, nil)
		productRatesMapResult[value] = toReturn
	}

	return productRatesMapResult, nil
}

func (s *TaxProviderService) GetCacheKey(id uuid.UUID, regionId uuid.UUID) string {
	return fmt.Sprintf("txrtcache:%s:%s", id.String(), regionId.String())
}

func (s *TaxProviderService) RegisterInstalledProviders(providers uuid.UUIDs) *utils.ApplictaionError {
	if err := s.r.TaxProviderRepository().Update(s.ctx, &models.TaxProvider{IsInstalled: false}); err != nil {
		return err
	}

	for _, p := range providers {
		var model *models.TaxProvider
		model.IsInstalled = true
		model.Id = p

		if err := s.r.TaxProviderRepository().Save(s.ctx, model); err != nil {
			return err
		}
	}

	return nil
}

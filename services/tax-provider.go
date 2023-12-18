package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/repository"
	"github.com/driver005/gateway/types"
	"github.com/google/uuid"
	"github.com/sarulabs/di"
)

type TaxProviderService struct {
	ctx            context.Context
	container      di.Container
	repo           *repository.TaxProviderRepo
	taxLineRepo    *repository.LineItemTaxLineRepo
	smTaxLineRepo  *repository.ShippingMethodTaxLineRepo
	cacheService   interfaces.ICacheService
	taxRateService *TaxRateService
}

func NewTaxProviderService(
	ctx context.Context,
	container di.Container,
	repo *repository.TaxProviderRepo,
	taxLineRepo *repository.LineItemTaxLineRepo,
	smTaxLineRepo *repository.ShippingMethodTaxLineRepo,
	cacheService interfaces.ICacheService,
	taxRateService *TaxRateService,
) *TaxProviderService {
	return &TaxProviderService{
		ctx,
		container,
		repo,
		taxLineRepo,
		smTaxLineRepo,
		cacheService,
		taxRateService,
	}
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

func (s *TaxProviderService) ClearLineItemsTaxLines(itemIds uuid.UUIDs) error {
	if err := s.taxLineRepo.Specification(repository.In[uuid.UUID]("item_id", itemIds)).Delete(s.ctx, &models.LineItemTaxLine{}); err != nil {
		return err
	}

	return nil
}

func (s *TaxProviderService) ClearTaxLines(cartId uuid.UUID) error {
	if err := s.taxLineRepo.DeleteForCart(s.ctx, cartId); err != nil {
		return err
	}

	if err := s.smTaxLineRepo.DeleteForCart(s.ctx, cartId); err != nil {
		return err
	}

	return nil
}

func (s *TaxProviderService) CreateTaxLines(cart *models.Cart, lineItems []models.LineItem, calculationContext *interfaces.TaxCalculationContext) ([]models.ShippingMethodTaxLine, []models.LineItemTaxLine, error) {
	var taxline []models.LineItemTaxLine
	var smTaxLine []models.ShippingMethodTaxLine
	var err error

	if cart != nil {
		smTaxLine, taxline, err = s.GetTaxLines(cart.Items, calculationContext)
	} else {
		smTaxLine, taxline, err = s.GetTaxLines(lineItems, calculationContext)
	}
	if err != nil {
		return nil, nil, err
	}

	if err := s.taxLineRepo.SaveSlice(s.ctx, taxline); err != nil {
		return nil, nil, err
	}

	if err := s.smTaxLineRepo.SaveSlice(s.ctx, smTaxLine); err != nil {
		return nil, nil, err
	}

	return smTaxLine, taxline, err
}

func (s *TaxProviderService) CreateShippingTaxLines(shippingMethod *models.ShippingMethod, calculationContext *interfaces.TaxCalculationContext) ([]models.ShippingMethodTaxLine, error) {
	taxline, err := s.GetShippingTaxLines(shippingMethod, calculationContext)
	if err != nil {
		return nil, err
	}

	if err := s.smTaxLineRepo.SaveSlice(s.ctx, taxline); err != nil {
		return nil, err
	}

	return taxline, nil
}

func (s *TaxProviderService) GetShippingTaxLines(shippingMethod *models.ShippingMethod, calculationContext *interfaces.TaxCalculationContext) ([]models.ShippingMethodTaxLine, error) {
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
				UUID:  pl.ShippingMethodID,
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

func (s *TaxProviderService) GetTaxLines(lineItems []models.LineItem, calculationContext *interfaces.TaxCalculationContext) ([]models.ShippingMethodTaxLine, []models.LineItemTaxLine, error) {
	var productIds []uuid.UUID
	var productIdsMap map[uuid.UUID]bool

	for _, item := range lineItems {
		if item.Variant != nil && item.Variant.Id != uuid.Nil {
			productIdsMap[item.Variant.Id] = true
		}
	}

	for id := range productIdsMap {
		productIds = append(productIds, id)
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
		if pl.ShippingMethodID != uuid.Nil {
			smTaxLines = append(smTaxLines, models.ShippingMethodTaxLine{
				Model: core.Model{
					Metadata: pl.Metadata,
				},
				ShippingMethodId: uuid.NullUUID{UUID: pl.ShippingMethodID},
				Rate:             pl.Rate,
				Name:             pl.Name,
				Code:             pl.Code,
			})
		} else if pl.ItemID == uuid.Nil {
			return nil, nil, errors.New("Tax Provider returned invalid tax lines")
		} else {
			liTaxLines = append(liTaxLines, models.LineItemTaxLine{
				Model: core.Model{
					Metadata: pl.Metadata,
				},
				ItemId: uuid.NullUUID{UUID: pl.ItemID},
				Rate:   pl.Rate,
				Name:   pl.Name,
				Code:   pl.Code,
			})
		}
	}

	return smTaxLines, liTaxLines, nil
}

func (s *TaxProviderService) GetTaxLinesMap(items []models.LineItem, calculationContext *interfaces.TaxCalculationContext) (types.TaxLinesMaps, error) {
	var lineItemsTaxLinesMap map[uuid.UUID][]models.LineItemTaxLine
	var shippingMethodsTaxLinesMap map[uuid.UUID][]models.ShippingMethodTaxLine

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

func (s *TaxProviderService) GetRegionRatesForShipping(optionId uuid.UUID, region *models.Region) ([]types.TaxServiceRate, error) {
	cachKey := s.GetCacheKey(optionId, region.Id)
	cacheHit, err := s.cacheService.Get(cachKey)
	if err != nil {
		return nil, err
	}

	if cacheHit != nil {
		return cacheHit.([]types.TaxServiceRate), nil
	}

	var toReturn []types.TaxServiceRate
	optionRates, err := s.taxRateService.ListByShippingOption(optionId)
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

	s.cacheService.Set(cachKey, toReturn, nil)
	return toReturn, nil
}

func (s *TaxProviderService) GetRegionRatesForProduct(productIds uuid.UUIDs, region *models.Region) (map[uuid.UUID][]types.TaxServiceRate, error) {
	var nonCachedProductIds uuid.UUIDs
	var cacheKeysMap map[uuid.UUID]string
	var productRatesMapResult map[uuid.UUID][]types.TaxServiceRate

	for _, p := range productIds {
		cacheKeysMap[p] = s.GetCacheKey(p, region.Id)
	}

	for key, value := range cacheKeysMap {
		cacheHit, err := s.cacheService.Get(value)
		if err != nil {
			return nil, err
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
		rates, err := s.taxRateService.ListByProduct(value, types.TaxRateListByConfig{RegionId: region.Id})
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

		s.cacheService.Set(cacheKeysMap[value], toReturn, nil)
		productRatesMapResult[value] = toReturn
	}

	return productRatesMapResult, nil
}

func (s *TaxProviderService) GetCacheKey(id uuid.UUID, regionId uuid.UUID) string {
	return fmt.Sprintf("txrtcache:%s:%s", id.String(), regionId.String())
}

func (s *TaxProviderService) RegisterInstalledProviders(providers []string) error {
	if err := s.repo.Update(s.ctx, &models.TaxProvider{IsInstalled: false}); err != nil {
		return err
	}

	for _, p := range providers {
		var model *models.TaxProvider
		model.IsInstalled = true
		if err := model.ParseUUID(p); err != nil {
			return err
		}

		if err := s.repo.Save(s.ctx, model); err != nil {
			return err
		}
	}

	return nil
}

package strategies

import (
	"context"
	"fmt"
	"reflect"

	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type PriceSelectionStrategy struct {
	ctx context.Context
	r   Registry
}

func NewPriceSelectionStrategy(
	r Registry,
) *PriceSelectionStrategy {
	return &PriceSelectionStrategy{
		context.Background(),
		r,
	}
}

func (s *PriceSelectionStrategy) SetContext(context context.Context) *PriceSelectionStrategy {
	s.ctx = context
	return s
}

func (s *PriceSelectionStrategy) CalculateVariantPrice(data []interfaces.Pricing, context *interfaces.PricingContext) (map[uuid.UUID]interfaces.PriceSelectionResult, *utils.ApplictaionError) {
	dataMap := make(map[uuid.UUID]interfaces.Pricing)
	for _, d := range data {
		dataMap[d.VariantId] = d
	}

	cacheKeysMap := make(map[uuid.UUID]string)
	for _, d := range data {
		cacheKeysMap[d.VariantId] = s.GetCacheKey(d.VariantId, context, d.Quantity)
	}

	nonCachedData := make([]interfaces.Pricing, 0)
	variantPricesMap := make(map[uuid.UUID]interfaces.PriceSelectionResult)

	if !context.IgnoreCache {
		var cacheHits []interfaces.PriceSelectionResult
		for _, cacheKey := range cacheKeysMap {
			cacheHit, err := s.r.CacheService().Get(cacheKey)
			if err != nil {
				return nil, utils.NewApplictaionError(
					utils.INVALID_DATA,
					err.Error(),
				)
			}
			cacheHits = append(cacheHits, cacheHit.(interfaces.PriceSelectionResult))
		}

		if len(cacheHits) == 0 {
			for _, d := range dataMap {
				nonCachedData = append(nonCachedData, d)
			}
		}

		for index, cacheHit := range cacheHits {
			variantId := data[index].VariantId
			if !reflect.DeepEqual(cacheHit, interfaces.PriceSelectionResult{}) {
				variantPricesMap[variantId] = cacheHit
				continue
			}
			nonCachedData = append(nonCachedData, dataMap[variantId])
		}
	} else {
		for _, d := range dataMap {
			nonCachedData = append(nonCachedData, d)
		}
	}

	// results := make(map[uuid.UUID]interfaces.PriceSelectionResult)
	results, err := s.CalculateVariantPriceNew(nonCachedData, context)
	if err != nil {
		return nil, err
	}

	for variantId, prices := range results {
		variantPricesMap[variantId] = prices
		if !context.IgnoreCache {
			err := s.r.CacheService().Set(cacheKeysMap[variantId], prices, nil)
			if err != nil {
				return nil, utils.NewApplictaionError(
					utils.INVALID_DATA,
					err.Error(),
				)
			}
		}
	}

	return variantPricesMap, nil
}

func (s *PriceSelectionStrategy) CalculateVariantPriceNew(data []interfaces.Pricing, context *interfaces.PricingContext) (map[uuid.UUID]interfaces.PriceSelectionResult, *utils.ApplictaionError) {
	var variantIDs uuid.UUIDs
	for _, d := range data {
		variantIDs = append(variantIDs, d.VariantId)
	}

	variantsPrices, _, err := s.r.MoneyAmountRepository().FindManyForVariantsInRegion(variantIDs, context.RegionId, context.CurrencyCode, context.CustomerId, context.IncludeDiscountPrices, true)
	if err != nil {
		return nil, err
	}

	variantPricesMap := make(map[uuid.UUID]interfaces.PriceSelectionResult)
	for variantId, prices := range variantsPrices {
		dataItem, _ := lo.Find(data, func(item interfaces.Pricing) bool {
			return item.VariantId == variantId
		})
		result := interfaces.PriceSelectionResult{
			Prices: prices,
		}

		if len(prices) == 0 || context != nil {
			variantPricesMap[variantId] = result
		}

		taxRate := 0.0
		for _, nextTaxRate := range context.TaxRates {
			taxRate += nextTaxRate.Rate / 100
		}

		for _, ma := range prices {
			isTaxInclusive := ma.Currency.IncludesTax
			if ma.PriceList != nil && ma.PriceList.IncludesTax {
				isTaxInclusive = ma.PriceList.IncludesTax
			} else if ma.Region != nil && ma.Region.IncludesTax {
				isTaxInclusive = ma.Region.IncludesTax
			}

			if ma.RegionId.UUID == context.RegionId && ma.PriceListId.UUID == uuid.Nil && reflect.ValueOf(ma.MinQuantity).IsZero() && reflect.ValueOf(ma.MaxQuantity).IsZero() {
				result.OriginalPriceIncludesTax = isTaxInclusive
				result.OriginalPrice = ma.Amount
			}

			if ma.CurrencyCode == context.CurrencyCode && ma.PriceListId.UUID == uuid.Nil && reflect.ValueOf(ma.MinQuantity).IsZero() && reflect.ValueOf(ma.MaxQuantity).IsZero() && reflect.ValueOf(result.OriginalPrice).IsZero() {
				result.OriginalPriceIncludesTax = isTaxInclusive
				result.OriginalPrice = ma.Amount
			}

			if IsValidQuantity(ma, dataItem.Quantity) && IsValidAmount(ma.Amount, result, isTaxInclusive, taxRate) && (ma.CurrencyCode == context.CurrencyCode || ma.RegionId.UUID == context.RegionId) {
				result.CalculatedPrice = ma.Amount
				result.CalculatedPriceType = string(ma.PriceList.Type)
				result.CalculatedPriceIncludesTax = isTaxInclusive
			}
		}

		variantPricesMap[variantId] = result
	}

	return variantPricesMap, nil
}

func (s *PriceSelectionStrategy) CalculateVariantPriceOld(data []interfaces.Pricing, context *interfaces.PricingContext) (map[uuid.UUID]interfaces.PriceSelectionResult, *utils.ApplictaionError) {
	var variantIDs uuid.UUIDs
	for _, d := range data {
		variantIDs = append(variantIDs, d.VariantId)
	}

	variantsPrices, _, err := s.r.MoneyAmountRepository().FindManyForVariantsInRegion(variantIDs, context.RegionId, context.CurrencyCode, context.CustomerId, context.IncludeDiscountPrices, true)
	if err != nil {
		return nil, err
	}

	variantPricesMap := make(map[uuid.UUID]interfaces.PriceSelectionResult)
	for variantId, prices := range variantsPrices {
		dataItem, _ := lo.Find(data, func(item interfaces.Pricing) bool {
			return item.VariantId == variantId
		})

		result := interfaces.PriceSelectionResult{
			Prices: prices,
		}

		if len(prices) == 0 || context != nil {
			variantPricesMap[variantId] = result
		}

		for _, ma := range prices {
			if ma.RegionId.UUID == context.RegionId && ma.PriceListId.UUID == uuid.Nil && reflect.ValueOf(ma.MinQuantity).IsZero() && reflect.ValueOf(ma.MaxQuantity).IsZero() {
				result.OriginalPrice = ma.Amount
			}

			if ma.CurrencyCode == context.CurrencyCode && ma.PriceListId.UUID == uuid.Nil && reflect.ValueOf(ma.MinQuantity).IsZero() && reflect.ValueOf(ma.MaxQuantity).IsZero() && reflect.ValueOf(result.OriginalPrice).IsZero() {
				result.OriginalPrice = ma.Amount
			}

			if IsValidQuantity(ma, dataItem.Quantity) && (reflect.ValueOf(result.CalculatedPrice).IsZero() || ma.Amount < result.CalculatedPrice) && (ma.CurrencyCode == context.CurrencyCode || ma.RegionId.UUID == context.RegionId) {
				result.CalculatedPrice = ma.Amount
				result.CalculatedPriceType = string(ma.PriceList.Type)
			}
		}

		variantPricesMap[variantId] = result
	}

	return variantPricesMap, nil
}

func (s *PriceSelectionStrategy) OnVariantsPricesUpdate(variantIds uuid.UUIDs) error {
	for _, id := range variantIds {
		err := s.r.CacheService().Invalidate(fmt.Sprintf("ps:%s:*", id))
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *PriceSelectionStrategy) GetCacheKey(variantID uuid.UUID, context *interfaces.PricingContext, quantity int) string {
	taxRate := 0.0
	for _, nextTaxRate := range context.TaxRates {
		taxRate += nextTaxRate.Rate / 100
	}

	return fmt.Sprintf("ps:%s:%s:%s:%s:%d:%t:%f", variantID, context.RegionId, context.CurrencyCode, context.CustomerId, quantity, context.IncludeDiscountPrices, taxRate)
}

func IsValidAmount(amount float64, result interfaces.PriceSelectionResult, isTaxInclusive bool, taxRate float64) bool {
	if reflect.ValueOf(result.CalculatedPrice).IsZero() {
		return true
	}

	if isTaxInclusive == result.CalculatedPriceIncludesTax {
		return amount < result.CalculatedPrice
	}

	if taxRate != 0.0 {
		if isTaxInclusive {
			return amount < (1+taxRate)*(result.CalculatedPrice)
		} else {
			return (1+taxRate)*amount < result.CalculatedPrice
		}
	}

	return false
}

func IsValidQuantity(price models.MoneyAmount, quantity int) bool {
	if quantity != 0 && IsValidPriceWithQuantity(price, quantity) {
		return true
	}

	if quantity == 0 && IsValidPriceWithoutQuantity(price) {
		return true
	}

	return false
}

func IsValidPriceWithoutQuantity(price models.MoneyAmount) bool {
	if reflect.ValueOf(price.MinQuantity).IsZero() && reflect.ValueOf(price.MaxQuantity).IsZero() {
		return true
	}

	if reflect.ValueOf(price.MinQuantity).IsZero() || price.MinQuantity == 0 {
		return !reflect.ValueOf(price.MaxQuantity).IsZero()
	}

	return false
}

func IsValidPriceWithQuantity(price models.MoneyAmount, quantity int) bool {
	if !reflect.ValueOf(price.MinQuantity).IsZero() || price.MinQuantity <= quantity {
		if !reflect.ValueOf(price.MaxQuantity).IsZero() || price.MaxQuantity >= quantity {
			return true
		}
	}

	return false
}

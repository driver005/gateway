package services

import (
	"context"
	"math"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type PricingService struct {
	ctx context.Context
	r   Registry
}

func NewPricingService(
	r Registry,
) *PricingService {
	return &PricingService{
		context.Background(),
		r,
	}
}

func (s *PricingService) SetContext(context context.Context) *PricingService {
	s.ctx = context
	return s
}

func (s *PricingService) collectPricingContext(context *interfaces.PricingContext) (*interfaces.PricingContext, *utils.ApplictaionError) {
	var region *models.Region = &models.Region{}
	var pricingContext *interfaces.PricingContext

	pricingContext.AutomaticTaxes = false
	pricingContext.CurrencyCode = context.CurrencyCode

	if context.RegionId != uuid.Nil {
		var err *utils.ApplictaionError
		region, err = s.r.RegionService().SetContext(s.ctx).Retrieve(context.RegionId, &sql.Options{
			Selects: []string{"id", "currency_code", "automatic_taxes", "tax_rate"},
		})
		if err != nil {
			return nil, err
		}

		pricingContext.AutomaticTaxes = region.AutomaticTaxes
		pricingContext.TaxRate = region.TaxRate
		pricingContext.CurrencyCode = region.CurrencyCode
	}
	return pricingContext, nil
}

func (s *PricingService) calculateTaxes(variantPricing types.ProductVariantPricing, productRates []types.TaxServiceRate) types.TaxedPricing {
	rate := 0.0
	for _, nextTaxRate := range productRates {
		rate += (nextTaxRate.Rate / 100)
	}
	taxedPricing := types.TaxedPricing{
		TaxRates: productRates,
	}
	if variantPricing.CalculatedPrice != 0.0 {
		includesTax := false
		feature := true
		if feature && variantPricing.CalculatedPriceIncludesTax {
			includesTax = true
		}
		taxedPricing.CalculatedTax = math.Round(utils.CalculatePriceTaxAmount(variantPricing.CalculatedPrice, rate, includesTax))
		if variantPricing.CalculatedPriceIncludesTax {
			taxedPricing.CalculatedPriceInclTax = variantPricing.CalculatedPrice
		} else {
			taxedPricing.CalculatedPriceInclTax = variantPricing.CalculatedPrice + taxedPricing.CalculatedTax
		}
	}
	if variantPricing.OriginalPrice != 0.0 {
		includesTax := false
		feature := true
		if feature && variantPricing.OriginalPriceIncludesTax {
			includesTax = true
		}
		taxedPricing.OriginalTax = math.Round(utils.CalculatePriceTaxAmount(variantPricing.OriginalPrice, rate, includesTax))
		if variantPricing.OriginalPriceIncludesTax {
			taxedPricing.OriginalPriceInclTax = variantPricing.OriginalPrice
		} else {
			taxedPricing.OriginalPriceInclTax = variantPricing.OriginalPrice + taxedPricing.OriginalTax
		}
	}
	return taxedPricing
}

// func (s *PricingService) getProductVariantPricingModulePricing(variantPriceData []interfaces.Pricing, context *interfaces.PricingContext) (map[uuid.UUID]types.ProductVariantPricing, *utils.ApplictaionError) {
// 	variables := map[string]interface{}{
// 		"variant_id": variantPriceData,
// 		"take":       nil,
// 	}
// 	query := map[string]interface{}{
// 		"product_variant_price_set": map[string]interface{}{
// 			"__args": variables,
// 			"fields": []string{"variant_id", "price_set_id"},
// 		},
// 	}
// 	variantPriceSets, err := s.r.ProductVariantService().SetContext(s.ctx).List(variantPriceData[0].VariantId, &sql.Options{Selects: []string{"variant_id", "price_set_id"}})
// 	if err != nil {
// 		return nil, err
// 	}
// 	var variantIdToPriceSetIdMap map[uuid.UUID]uuid.UUID
// 	for _, variantPriceSet := range variantPriceSets {
// 		variantIdToPriceSetIdMap[variantPriceSet.Id] = variantPriceSet.Prices
// 	}
// 	var priceSetIds uuid.UUIDs
// 	for _, variantPriceSet := range variantPriceSets {
// 		priceSetIds = append(priceSetIds, variantPriceSet.PriceSetID)
// 	}

// 	if context.CustomerId != uuid.Nil {
// 		customer, err := s.r.CustomerService().SetContext(s.ctx).RetrieveById(context.CustomerId, &sql.Options{
// 			Relations: []string{"groups"},
// 		})
// 		if err != nil {
// 			return nil, err
// 		}
// 		if len(customer.Groups) > 0 {
// 			var groupIDs uuid.UUIDs
// 			for _, group := range customer.Groups {
// 				groupIDs = append(groupIDs, group.Id)
// 			}
// 			context.CustomerGroupId = groupIDs
// 		}
// 	}
// 	var calculatedPrices []interfaces.CalculatedPriceSet
// 	if context.CurrencyCode != "" {
// 		prices, err := s.pricingModuleService.CalculatePrices(s.ctx, map[string]interface{}{}, priceSetIds)
// 		if err != nil {
// 			return nil, err
// 		}
// 		calculatedPrices = prices
// 	}

// 	calculatedPriceMap := make(map[uuid.UUID]interfaces.CalculatedPriceSet)
// 	for _, priceSet := range calculatedPrices {
// 		calculatedPriceMap[priceSet.Id] = priceSet
// 	}
// 	pricingResultMap := make(map[uuid.UUID]types.ProductVariantPricing)
// 	for _, pricedata := range variantPriceData {
// 		variantId := pricedata.VariantId
// 		priceSetId := variantIdToPriceSetIdMap[variantId]
// 		pricingResult := types.ProductVariantPricing{
// 			Prices: []models.MoneyAmount{},
// 		}
// 		if priceSetId != uuid.Nil {
// 			calculatedPrices, ok := calculatedPriceMap[priceSetId]
// 			if ok {
// 				pricingResult.Prices = append(pricingResult.Prices, models.MoneyAmount{
// 					Model: core.Model{
// 						Id: calculatedPrices.OriginalPrice.MoneyAmountId,
// 					},
// 					CurrencyCode: calculatedPrices.CurrencyCode,
// 					Amount:       calculatedPrices.OriginalAmount,
// 					MinQuantity:  calculatedPrices.OriginalPrice.MinQuantity,
// 					MaxQuantity:  calculatedPrices.OriginalPrice.MaxQuantity,
// 					PriceListId:  uuid.NullUUID{UUID: calculatedPrices.OriginalPrice.PriceListId},
// 				})
// 				if calculatedPrices.CalculatedPrice.MoneyAmountId != calculatedPrices.OriginalPrice.MoneyAmountId {
// 					pricingResult.Prices = append(pricingResult.Prices, models.MoneyAmount{
// 						Model: core.Model{
// 							Id: calculatedPrices.OriginalPrice.MoneyAmountId,
// 						},
// 						CurrencyCode: calculatedPrices.CurrencyCode,
// 						Amount:       calculatedPrices.CalculatedAmount,
// 						MinQuantity:  calculatedPrices.CalculatedPrice.MinQuantity,
// 						MaxQuantity:  calculatedPrices.CalculatedPrice.MaxQuantity,
// 						PriceListId:  uuid.NullUUID{UUID: calculatedPrices.OriginalPrice.PriceListId},
// 					})
// 				}
// 				pricingResult.OriginalPrice = calculatedPrices.OriginalAmount
// 				pricingResult.CalculatedPrice = calculatedPrices.CalculatedAmount
// 				pricingResult.CalculatedPriceType = calculatedPrices.CalculatedPrice.PriceListType
// 			}
// 		}
// 		pricingResultMap[variantId] = pricingResult
// 	}
// 	return pricingResultMap, nil
// }

func (s *PricingService) getProductVariantPricing(data []interfaces.Pricing, context *interfaces.PricingContext) (map[uuid.UUID]types.ProductVariantPricing, *utils.ApplictaionError) {
	// feature := true

	// if feature {
	// 	return s.getProductVariantPricingModulePricing(data, context)
	// }
	variantsPricing, err := s.r.PriceSelectionStrategy().CalculateVariantPrice(data, context)
	if err != nil {
		return nil, err
	}
	pricingResultMap := make(map[uuid.UUID]types.ProductVariantPricing)
	for variantId, pricing := range variantsPricing {
		pricingResult := types.ProductVariantPricing{
			Prices:                     pricing.Prices,
			OriginalPrice:              pricing.OriginalPrice,
			CalculatedPrice:            pricing.CalculatedPrice,
			CalculatedPriceType:        pricing.CalculatedPriceType,
			OriginalPriceIncludesTax:   pricing.OriginalPriceIncludesTax,
			CalculatedPriceIncludesTax: pricing.CalculatedPriceIncludesTax,
		}
		if context.AutomaticTaxes && context.RegionId != uuid.Nil {
			taxRates := context.TaxRates
			taxResults := s.calculateTaxes(pricingResult, taxRates)
			pricingResult.OriginalPriceInclTax = taxResults.OriginalPriceInclTax
			pricingResult.CalculatedPriceInclTax = taxResults.CalculatedPriceInclTax
			pricingResult.OriginalTax = taxResults.OriginalTax
			pricingResult.CalculatedTax = taxResults.CalculatedTax
			pricingResult.TaxRates = taxResults.TaxRates
		}
		pricingResultMap[variantId] = pricingResult
	}
	return pricingResultMap, nil
}

func (s *PricingService) GetProductVariantPricingById(variantId uuid.UUID, context *interfaces.PricingContext) (*types.ProductVariantPricing, *utils.ApplictaionError) {
	var pricingContext *interfaces.PricingContext
	if context.AutomaticTaxes {
		pricingContext = context
	} else {
		var err *utils.ApplictaionError
		pricingContext, err = s.collectPricingContext(context)
		if err != nil {
			return nil, err
		}
	}
	var productRates []types.TaxServiceRate
	if pricingContext.AutomaticTaxes && pricingContext.RegionId != uuid.Nil {
		product, err := s.r.ProductVariantService().SetContext(s.ctx).Retrieve(variantId, &sql.Options{
			Selects: []string{"id", "product_id"},
		})
		if err != nil {
			return nil, err
		}
		regionRatesForProduct, err := s.r.TaxProviderService().SetContext(s.ctx).GetRegionRatesForProduct(uuid.UUIDs{product.ProductId.UUID}, &models.Region{
			Model: core.Model{
				Id: pricingContext.RegionId,
			},
			TaxRate: pricingContext.TaxRate,
		})
		if err != nil {
			return nil, err
		}
		productRates = regionRatesForProduct[product.ProductId.UUID]
	}
	pricingContext.TaxRates = productRates
	productVariantPricing, err := s.getProductVariantPricing(
		[]interfaces.Pricing{
			{
				VariantId: variantId,
				Quantity:  pricingContext.Quantity,
			},
		},
		pricingContext,
	)
	if err != nil {
		return nil, err
	}

	res := productVariantPricing[variantId]
	return &res, nil
}

func (s *PricingService) GetProductVariantsPricing(data []interfaces.Pricing, context *interfaces.PricingContext) (map[uuid.UUID]types.ProductVariantPricing, *utils.ApplictaionError) {
	var pricingContext *interfaces.PricingContext
	if context.AutomaticTaxes {
		pricingContext = context
	} else {
		ctx, err := s.collectPricingContext(context)
		if err != nil {
			return nil, err
		}
		pricingContext = ctx
	}
	dataMap := make(map[uuid.UUID]interfaces.Pricing)
	var variantIds []uuid.UUID
	for _, d := range data {
		dataMap[d.VariantId] = d
		variantIds = append(variantIds, d.VariantId)
	}
	variants, err := s.r.ProductVariantService().SetContext(s.ctx).List(&types.FilterableProductVariant{
		FilterModel: core.FilterModel{Id: variantIds},
	}, &sql.Options{
		Selects: []string{"id", "product_id"},
	})
	if err != nil {
		return nil, err
	}

	if pricingContext.RegionId != uuid.Nil {
		productId := variants[0].ProductId.UUID
		productsRatesMap, err := s.r.TaxProviderService().SetContext(s.ctx).GetRegionRatesForProduct(uuid.UUIDs{productId}, &models.Region{
			Model: core.Model{
				Id: pricingContext.RegionId,
			},
			TaxRate: pricingContext.TaxRate,
		})
		if err != nil {
			return nil, err
		}
		pricingContext.TaxRates = productsRatesMap[productId]
	}
	var pricingData []interfaces.Pricing
	for _, v := range variants {
		pricingData = append(pricingData, interfaces.Pricing{
			VariantId: v.Id,
			Quantity:  dataMap[v.Id].Quantity,
		})
	}
	variantsPricingMap, err := s.getProductVariantPricing(pricingData, pricingContext)
	if err != nil {
		return nil, err
	}
	pricingResult := make(map[uuid.UUID]types.ProductVariantPricing)
	for _, d := range data {
		pricingResult[d.VariantId] = variantsPricingMap[d.VariantId]
	}
	return pricingResult, nil
}

func (s *PricingService) getProductPricing(data []interfaces.ProductPricing, context *interfaces.PricingContext) (map[uuid.UUID]map[uuid.UUID]types.ProductVariantPricing, *utils.ApplictaionError) {
	taxRatesMap := make(map[uuid.UUID][]types.TaxServiceRate)
	if context.AutomaticTaxes && context.RegionId != uuid.Nil {
		var err *utils.ApplictaionError
		var dataIds uuid.UUIDs
		for _, d := range data {
			dataIds = append(dataIds, d.ProductId)
		}
		taxRatesMap, err = s.r.TaxProviderService().SetContext(s.ctx).GetRegionRatesForProduct(dataIds, &models.Region{
			Model:   core.Model{Id: context.RegionId},
			TaxRate: context.TaxRate,
		})
		if err != nil {
			return nil, err
		}
	}
	productsPricingMap := make(map[uuid.UUID]map[uuid.UUID]types.ProductVariantPricing)
	for _, d := range data {
		var pricingData []interfaces.Pricing
		for _, v := range d.Variants {
			pricingData = append(pricingData, interfaces.Pricing{
				VariantId: v.Id,
				Quantity:  context.Quantity,
			})
		}

		if context.AutomaticTaxes && context.RegionId != uuid.Nil {
			context.TaxRates = taxRatesMap[d.ProductId]
		}
		variantsPricingMap, err := s.getProductVariantPricing(pricingData, context)
		if err != nil {
			return nil, err
		}
		productVariantsPricing := productsPricingMap[d.ProductId]
		for variantId, variantPricing := range variantsPricingMap {
			productVariantsPricing[variantId] = variantPricing
		}
		productsPricingMap[d.ProductId] = productVariantsPricing
	}
	return productsPricingMap, nil
}

func (s *PricingService) SetVariantPrices(variants []models.ProductVariant, context *interfaces.PricingContext) ([]models.ProductVariant, *utils.ApplictaionError) {
	pricingContext, err := s.collectPricingContext(context)
	if err != nil {
		return nil, err
	}

	var pricingData []interfaces.Pricing
	for _, v := range variants {
		pricingData = append(pricingData, interfaces.Pricing{
			VariantId: v.Id,
			Quantity:  context.Quantity,
		})
	}
	variantsPricingMap, err := s.GetProductVariantsPricing(pricingData, pricingContext)
	if err != nil {
		return nil, err
	}
	var result []models.ProductVariant
	for i, variant := range variants {
		variantPricing := variantsPricingMap[variant.Id]
		copier.CopyWithOption(&variant, variantPricing, copier.Option{IgnoreEmpty: true, DeepCopy: true})
		result[i] = variant
	}
	return result, nil
}

func (s *PricingService) SetProductPrices(products []models.Product, context *interfaces.PricingContext) ([]models.Product, *utils.ApplictaionError) {
	pricingContext, err := s.collectPricingContext(context)
	if err != nil {
		return nil, err
	}

	var pricingData []interfaces.ProductPricing
	for _, p := range products {
		pricingData = append(pricingData, interfaces.ProductPricing{
			ProductId: p.Id,
			Variants:  p.Variants,
		})
	}

	productsPricingMap, err := s.getProductPricing(pricingData, pricingContext)
	if err != nil {
		return nil, err
	}
	var pricedProducts []models.Product
	for i, product := range products {
		if len(product.Variants) == 0 {
			pricedProducts[i] = product
			continue
		}
		var pricedVariants []models.ProductVariant
		for j, variant := range product.Variants {
			variantPricing := productsPricingMap[product.Id][variant.Id]
			copier.CopyWithOption(&variant, variantPricing, copier.Option{IgnoreEmpty: true, DeepCopy: true})
			pricedVariants[j] = variant
		}
		product.Variants = pricedVariants
		pricedProducts[i] = product
	}
	return pricedProducts, nil
}

func (s *PricingService) GetProductPricing(product *models.Product, context *interfaces.PricingContext) (map[uuid.UUID]types.ProductVariantPricing, *utils.ApplictaionError) {
	pricingContext, err := s.collectPricingContext(context)
	if err != nil {
		return nil, err
	}
	pricingData := interfaces.ProductPricing{
		ProductId: product.Id,
		Variants:  product.Variants,
	}

	productsPricingMap, err := s.getProductPricing([]interfaces.ProductPricing{pricingData}, pricingContext)
	if err != nil {
		return nil, err
	}
	return productsPricingMap[product.Id], nil
}

func (s *PricingService) GetProductPricingById(productId uuid.UUID, context *interfaces.PricingContext) (map[uuid.UUID]types.ProductVariantPricing, *utils.ApplictaionError) {
	pricingContext, err := s.collectPricingContext(context)
	if err != nil {
		return nil, err
	}
	variants, err := s.r.ProductVariantService().SetContext(s.ctx).List(&types.FilterableProductVariant{
		ProductId: uuid.UUIDs{productId},
	}, &sql.Options{
		Selects: []string{"id"},
	})
	if err != nil {
		return nil, err
	}
	var pricingData []interfaces.Pricing
	for _, v := range variants {
		pricingData = append(pricingData, interfaces.Pricing{
			VariantId: v.Id,
			Quantity:  context.Quantity,
		})
	}

	variantsPricingMap, err := s.getProductVariantPricing(pricingData, pricingContext)
	if err != nil {
		return nil, err
	}
	return variantsPricingMap, nil
}

func (s *PricingService) getPricingModuleVariantMoneyAmounts(variantIds []uuid.UUID) (map[uuid.UUID][]models.MoneyAmount, *utils.ApplictaionError) {
	// variables := map[string]interface{}{
	// 	"variant_id": variantIds,
	// 	"take":       nil,
	// }
	// query := map[string]interface{}{
	// 	"product_variant_price_set": map[string]interface{}{
	// 		"__args": variables,
	// 		"fields": []string{"variant_id", "price_set_id"},
	// 	},
	// }
	// variantPriceSets, err := s.remoteQuery(query)
	// if err != nil {
	// 	return nil, err
	// }
	// priceSetIdToVariantIdMap := make(map[uuid.UUID]uuid.UUID)
	// for _, variantPriceSet := range variantPriceSets {
	// 	priceSetIdToVariantIdMap[variantPriceSet.price_set_id] = variantPriceSet.variant_id
	// }
	// var priceSetIds uuid.UUIDs
	// for i, variantPriceSet := range variantPriceSets {
	// 	priceSetIds[i] = variantPriceSet.price_set_id
	// }
	// priceSetMoneyAmounts, err := s.r.PricingModuleService().ListPriceSetMoneyAmounts(s.ctx, nil, interfaces.PriceSetMoneyAmount{
	// 	PriceSetId: priceSetIds,
	// }, &sql.Options{
	// 	Take:      nil,
	// 	Relations: []string{"money_amount", "price_list", "price_set", "price_rules", "price_rules.rule_type"},
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// variantIdMoneyAmountMap := make(map[uuid.UUID][]models.MoneyAmount)
	// for _, priceSetMoneyAmount := range priceSetMoneyAmounts {
	// 	variantId := priceSetIdToVariantIdMap[priceSetMoneyAmount.PriceSet.Id]
	// 	if variantId == uuid.Nil {
	// 		continue
	// 	}
	// 	var regionId uuid.UUID
	// 	for _, pr := range priceSetMoneyAmount.PriceRules {
	// 		if pr.RuleType.RuleAttribute == "region_id" {
	// 			regionId = pr.Value
	// 			break
	// 		}
	// 	}
	// 	priceSetMoneyAmount.PriceSet.MoneyAmounts = nil
	// 	moneyAmount := models.MoneyAmount{
	// 		Amount:       priceSetMoneyAmount.MoneyAmount.Amount,
	// 		CurrencyCode: priceSetMoneyAmount.MoneyAmount.CurrencyCode,
	// 		PriceListId:  uuid.NullUUID{UUID: priceSetMoneyAmount.PriceList.Id},
	// 		PriceList:    &priceSetMoneyAmount.PriceList,
	// 	}
	// 	if regionId != uuid.Nil {
	// 		moneyAmount.RegionId = uuid.NullUUID{UUID: regionId}
	// 	}
	// 	if _, ok := variantIdMoneyAmountMap[variantId]; ok {
	// 		variantIdMoneyAmountMap[variantId] = append(variantIdMoneyAmountMap[variantId], moneyAmount)
	// 	} else {
	// 		variantIdMoneyAmountMap[variantId] = []models.MoneyAmount{moneyAmount}
	// 	}
	// }
	// return variantIdMoneyAmountMap, nil

	return nil, nil
}

func (s *PricingService) SetAdminVariantPricing(variants []models.ProductVariant, context *interfaces.PricingContext) ([]models.ProductVariant, *utils.ApplictaionError) {
	feature := true
	if feature {
		return s.SetVariantPrices(variants, context)
	}
	var variantIds uuid.UUIDs
	for _, variant := range variants {
		variantIds = append(variantIds, variant.Id)
	}
	variantIdMoneyAmountMap, err := s.getPricingModuleVariantMoneyAmounts(variantIds)
	if err != nil {
		return nil, err
	}
	pricedVariants := make([]models.ProductVariant, len(variants))
	for i, variant := range variants {
		variant.Prices = variantIdMoneyAmountMap[variant.Id]
		pricedVariants[i] = variant
	}
	return pricedVariants, nil
}

func (s *PricingService) SetAdminProductPricing(products []models.Product) ([]models.Product, *utils.ApplictaionError) {
	feature := true
	if !feature {
		return s.SetProductPrices(products, &interfaces.PricingContext{})
	}
	var variantIds uuid.UUIDs
	for _, product := range products {
		for _, variant := range product.Variants {
			variantIds = append(variantIds, variant.Id)
		}
	}
	variantIdMoneyAmountMap, err := s.getPricingModuleVariantMoneyAmounts(variantIds)
	if err != nil {
		return nil, err
	}
	pricedProducts := make([]models.Product, len(products))
	for i, product := range products {
		if len(product.Variants) == 0 {
			pricedProducts[i] = product
			continue
		}
		pricedVariants := make([]models.ProductVariant, len(product.Variants))
		for j, productVariant := range product.Variants {
			productVariant.Prices = variantIdMoneyAmountMap[productVariant.Id]
			pricedVariants[j] = productVariant
		}
		product.Variants = pricedVariants
		pricedProducts[i] = product
	}
	return pricedProducts, nil
}

func (s *PricingService) GetShippingOptionPricing(shippingOption *models.ShippingOption, context *interfaces.PricingContext) (*types.PricedShippingOption, *utils.ApplictaionError) {
	if context != nil {
		c, err := s.collectPricingContext(context)
		if err != nil {
			return nil, err
		}
		context = c
	}
	var shippingOptionRates []types.TaxServiceRate
	if context.AutomaticTaxes && context.RegionId != uuid.Nil {
		tax, err := s.r.TaxProviderService().SetContext(s.ctx).GetRegionRatesForShipping(shippingOption.Id, &models.Region{
			Model:   core.Model{Id: context.RegionId},
			TaxRate: context.TaxRate,
		})
		if err != nil {
			return nil, err
		}
		shippingOptionRates = tax
	}
	price := shippingOption.Amount
	rate := 0.0
	for _, nextTaxRate := range shippingOptionRates {
		rate += nextTaxRate.Rate / 100
	}
	feature := true
	includesTax := feature && shippingOption.IncludesTax
	taxAmount := math.Round(utils.CalculatePriceTaxAmount(rate, price, includesTax))
	totalInclTax := price
	if includesTax {
		totalInclTax = price + taxAmount
	}

	return &types.PricedShippingOption{
		ShippingOption: shippingOption,
		PriceInclTax:   totalInclTax,
		TaxRates:       shippingOptionRates,
		TaxAmount:      taxAmount,
	}, nil
}

func (s *PricingService) SetShippingOptionPrices(shippingOptions []models.ShippingOption, context *interfaces.PricingContext) ([]types.PricedShippingOption, *utils.ApplictaionError) {
	var regions uuid.UUIDs
	for _, shippingOption := range shippingOptions {
		regions = append(regions, shippingOption.RegionId.UUID)
	}
	contexts := make([]struct {
		Context  *interfaces.PricingContext
		RegionId uuid.UUID
	}, 0)
	for _, regionId := range regions {
		c, err := s.collectPricingContext(&interfaces.PricingContext{
			RegionId: regionId,
		})
		if err != nil {
			return nil, err
		}
		contexts = append(contexts, struct {
			Context  *interfaces.PricingContext
			RegionId uuid.UUID
		}{
			Context:  c,
			RegionId: regionId,
		})
	}
	var pricedShippingOptions []types.PricedShippingOption
	for _, shippingOption := range shippingOptions {
		var pricingContext *interfaces.PricingContext
		for _, c := range contexts {
			if c.RegionId == shippingOption.RegionId.UUID {
				pricingContext = c.Context
				break
			}
		}
		if pricingContext.RegionId == uuid.Nil {
			return nil, utils.NewApplictaionError(
				utils.UNEXPECTED_STATE,
				"Could not find pricing context for shipping option",
				"500",
				nil,
			)
		}
		pricedShippingOption, err := s.GetShippingOptionPricing(&shippingOption, pricingContext)
		if err != nil {
			return nil, err
		}
		pricedShippingOptions = append(pricedShippingOptions, *pricedShippingOption)
	}
	return pricedShippingOptions, nil
}

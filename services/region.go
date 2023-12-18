package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type RegionService struct {
	ctx                           context.Context
	repo                          *repository.RegionRepo
	countryRepository             *repository.CountryRepo
	currencyRepository            *repository.CurrencyRepo
	paymentProviderRepository     *repository.PaymentProviderRepo
	fulfillmentProviderRepository *repository.FulfillmentProviderRepo
	taxProviderRepository         *repository.TaxProviderRepo
}

func NewRegionService(
	ctx context.Context,
	repo *repository.RegionRepo,
	countryRepository *repository.CountryRepo,
	currencyRepository *repository.CurrencyRepo,
	paymentProviderRepository *repository.PaymentProviderRepo,
	fulfillmentProviderRepository *repository.FulfillmentProviderRepo,
	taxProviderRepository *repository.TaxProviderRepo,
) *RegionService {
	return &RegionService{
		ctx,
		repo,
		countryRepository,
		currencyRepository,
		paymentProviderRepository,
		fulfillmentProviderRepository,
		taxProviderRepository,
	}
}

func (s *RegionService) create(data CreateRegionInput) (Region, error) {
    result, err := s.atomicPhase_(func(manager) (Region, error) {
        regionRepository := manager.withRepository(s.regionRepository_)
        currencyRepository := manager.withRepository(s.currencyRepository_)
        regionObject := data
        toValidate := data
        delete(toValidate, "metadata")
        delete(toValidate, "currency_code")
        validated, err := s.validateFields(toValidate)
        if err != nil {
            return Region{}, err
        }
        if s.featureFlagRouter_.isFeatureEnabled(TaxInclusivePricingFeatureFlag.key) {
            if includes_tax, ok := data["includes_tax"].(string); ok {
                regionObject["includes_tax"] = includes_tax
            }
        }
        if currency_code, ok := data["currency_code"].(string); ok {
            if err := s.validateCurrency(currency_code); err != nil {
                return Region{}, err
            }
            currency, err := currencyRepository.findOne(map[string]interface{}{"where": map[string]interface{}{"code": strings.ToLower(currency_code)}})
            if err != nil {
                return Region{}, MedusaError{Types: MedusaError.Types.INVALID_DATA, Message: fmt.Sprintf("Could not find currency with code %s", currency_code)}
            }
            regionObject["currency"] = currency
            regionObject["currency_code"] = strings.ToLower(currency_code)
        }
        if metadata, ok := data["metadata"]; ok {
            regionObject["metadata"] = setMetadata(regionObject["metadata"], metadata)
        }
        for key, value := range validated {
            regionObject[key] = value
        }
        created := regionRepository.create(regionObject)
        result, err := regionRepository.save(created)
        if err != nil {
            return Region{}, err
        }
        if err := s.eventBus_.withTransaction(manager).emit(RegionService.Events.CREATED, map[string]interface{}{"id": result.id}); err != nil {
            return Region{}, err
        }
        return result, nil
    })
    if err != nil {
        return Region{}, err
    }
    return result.(Region), nil
}

func (s *RegionService) update(regionId string, update UpdateRegionInput) (Region, error) {
    result, err := s.atomicPhase_(func(manager) (Region, error) {
        regionRepository := manager.withRepository(s.regionRepository_)
        currencyRepository := manager.withRepository(s.currencyRepository_)
        region, err := s.retrieve(regionId)
        if err != nil {
            return Region{}, err
        }
        toValidate := update
        delete(toValidate, "metadata")
        delete(toValidate, "currency_code")
        validated, err := s.validateFields(toValidate, region.id)
        if err != nil {
            return Region{}, err
        }
        if s.featureFlagRouter_.isFeatureEnabled(TaxInclusivePricingFeatureFlag.key) {
            if includes_tax, ok := update["includes_tax"].(string); ok {
                region.includes_tax = includes_tax
            }
        }
        if currency_code, ok := update["currency_code"].(string); ok {
            if err := s.validateCurrency(currency_code); err != nil {
                return Region{}, err
            }
            currency, err := currencyRepository.findOne(map[string]interface{}{"where": map[string]interface{}{"code": strings.ToLower(currency_code)}})
            if err != nil {
                return Region{}, MedusaError{Types: MedusaError.Types.INVALID_DATA, Message: fmt.Sprintf("Could not find currency with code %s", currency_code)}
            }
            region.currency_code = strings.ToLower(currency_code)
        }
        if metadata, ok := update["metadata"]; ok {
            region.metadata = setMetadata(region, metadata)
        }
        for key, value := range validated {
            region[key] = value
        }
        result, err := regionRepository.save(region)
        if err != nil {
            return Region{}, err
        }
        if err := s.eventBus_.withTransaction(manager).emit(RegionService.Events.UPDATED, map[string]interface{}{"id": result.id, "fields": Object.keys(update)}); err != nil {
            return Region{}, err
        }
        return result, nil
    })
    if err != nil {
        return Region{}, err
    }
    return result.(Region), nil
}

func (s *RegionService) validateFields(regionData interface{}, id string) (DeepPartial<Region>, error) {
    ppRepository := s.activeManager_.withRepository(s.paymentProviderRepository_)
    fpRepository := s.activeManager_.withRepository(s.fulfillmentProviderRepository_)
    tpRepository := s.activeManager_.withRepository(s.taxProviderRepository_)
    region := regionData
    if tax_rate, ok := regionData["tax_rate"].(float64); ok {
        if err := s.validateTaxRate(tax_rate); err != nil {
            return DeepPartial<Region>{}, err
        }
    }
    if countries, ok := regionData["countries"]; ok {
        region.countries = promiseAll(countries.map(func(countryCode string) {
            return s.validateCountry(countryCode, id)
        })).catch(func(err error) {
            return err
        })
    }
    if tax_provider_id, ok := regionData["tax_provider_id"].(string); ok {
        tp, err := tpRepository.findOne(map[string]interface{}{"where": map[string]interface{}{"id": tax_provider_id}})
        if err != nil {
            return DeepPartial<Region>{}, MedusaError{Types: MedusaError.Types.INVALID_DATA, Message: "Tax provider not found"}
        }
    }
    if payment_providers, ok := regionData["payment_providers"]; ok {
        region.payment_providers = promiseAll(payment_providers.map(func(pId string) {
            pp, err := ppRepository.findOne(map[string]interface{}{"where": map[string]interface{}{"id": pId}})
            if err != nil {
                return MedusaError{Types: MedusaError.Types.INVALID_DATA, Message: "Payment provider not found"}
            }
            return pp
        }))
    }
    if fulfillment_providers, ok := regionData["fulfillment_providers"]; ok {
        region.fulfillment_providers = promiseAll(fulfillment_providers.map(func(fId string) {
            fp, err := fpRepository.findOne(map[string]interface{}{"where": map[string]interface{}{"id": fId}})
            if err != nil {
                return MedusaError{Types: MedusaError.Types.INVALID_DATA, Message: "Fulfillment provider not found"}
            }
            return fp
        }))
    }
    return region, nil
}

func (s *RegionService) validateTaxRate(taxRate float64) error {
    if taxRate > 100 || taxRate < 0 {
        return MedusaError{Types: MedusaError.Types.INVALID_DATA, Message: "The tax_rate must be between 0 and 1"}
    }
    return nil
}

func (s *RegionService) validateCurrency(currencyCode string) error {
    store, err := s.storeService_.withTransaction(s.transactionManager_).retrieve(map[string]interface{}{"relations": []string{"currencies"}})
    if err != nil {
        return err
    }
    storeCurrencies := store.currencies.map(func(curr Currency) string {
        return curr.code
    })
    if !storeCurrencies.includes(strings.ToLower(currencyCode)) {
        return MedusaError{Types: MedusaError.Types.INVALID_DATA, Message: "Invalid currency code"}
    }
    return nil
}

func (s *RegionService) validateCountry(code string, regionId string) (Country, error) {
    countryRepository := s.activeManager_.withRepository(s.countryRepository_)
    countryCode := strings.ToUpper(code)
    isCountryExists := countries.some(func(country) bool {
        return country.alpha2 === countryCode
    })
    if !isCountryExists {
        return Country{}, MedusaError{Types: MedusaError.Types.INVALID_DATA, Message: "Invalid country code"}
    }
    country, err := countryRepository.findOne(map[string]interface{}{"where": map[string]interface{}{"iso_2": strings.ToLower(code)}})
    if err != nil {
        return Country{}, err
    }
    if country.region_id && country.region_id !== regionId {
        return Country{}, MedusaError{Types: MedusaError.Types.DUPLICATE_ERROR, Message: fmt.Sprintf("%s already exists in region %s", country.display_name, country.region_id)}
    }
    return country, nil
}

func (s *RegionService) retrieveByCountryCode(code string, config FindConfig<Region>) (Region, error) {
    countryRepository := s.activeManager_.withRepository(s.countryRepository_)
    query := buildQuery(map[string]interface{}{"iso_2": strings.ToLower(code)}, {})
    country, err := countryRepository.findOne(query)
    if err != nil {
        return Region{}, err
    }
    if !country.region_id {
        return Region{}, MedusaError{Types: MedusaError.Types.INVALID_DATA, Message: "Country does not belong to a region"}
    }
    return s.retrieve(country.region_id, config)
}

func (s *RegionService) retrieveByName(name string) (Region, error) {
    regions, err := s.list(map[string]interface{}{"name": name}, map[string]interface{}{"take": 1})
    if err != nil {
        return Region{}, err
    }
    if len(regions) == 0 {
        return Region{}, MedusaError{Types: MedusaError.Types.NOT_FOUND, Message: fmt.Sprintf("Region \"%s\" was not found", name)}
    }
    return regions[0], nil
}

func (s *RegionService) retrieve(regionId string, config FindConfig<Region>) (Region, error) {
    if regionId == "" {
        return Region{}, MedusaError{Types: MedusaError.Types.NOT_FOUND, Message: "\"regionId\" must be defined"}
    }
    regionRepository := s.activeManager_.withRepository(s.regionRepository_)
    query := buildQuery(map[string]interface{}{"id": regionId}, config)
    region, err := regionRepository.findOne(query)
    if err != nil {
        return Region{}, err
    }
    return region, nil
}

func (s *RegionService) list(selector Selector<Region>, config FindConfig<Region>) ([]Region, error) {
    regions, _, err := s.listAndCount(selector, config)
    if err != nil {
        return []Region{}, err
    }
    return regions, nil
}

func (s *RegionService) listAndCount(selector Selector<Region>, config FindConfig<Region>) ([]Region, int, error) {
    regionRepo := s.activeManager_.withRepository(s.regionRepository_)
    query := buildQuery(selector, config)
    regions, count, err := regionRepo.findAndCount(query)
    if err != nil {
        return []Region{}, 0, err
    }
    return regions, count, nil
}

func (s *RegionService) delete(regionId string) error {
    _, err := s.atomicPhase_(func(manager) (interface{}, error) {
        regionRepo := manager.withRepository(s.regionRepository_)
        countryRepo := manager.withRepository(s.countryRepository_)
        region, err := s.retrieve(regionId, map[string]interface{}{"relations": []string{"countries"}})
        if err != nil {
            return nil, err
        }
        if region == nil {
            return nil, nil
        }
        if err := regionRepo.softRemove(region); err != nil {
            return nil, err
        }
        if err := countryRepo.update(map[string]interface{}{"region_id": region.id}, map[string]interface{}{"region_id": nil}); err != nil {
            return nil, err
        }
        if err := s.eventBus_.withTransaction(manager).emit(RegionService.Events.DELETED, map[string]interface{}{"id": regionId}); err != nil {
            return nil, err
        }
        return nil, nil
    })
    if err != nil {
        return err
    }
    return nil
}

func (s *RegionService) addCountry(regionId string, code string) (Region, error) {
    result, err := s.atomicPhase_(func(manager) (Region, error) {
        regionRepo := manager.withRepository(s.regionRepository_)
        country, err := s.validateCountry(code, regionId)
        if err != nil {
            return Region{}, err
        }
        region, err := s.retrieve(regionId, map[string]interface{}{"relations": []string{"countries"}})
        if err != nil {
            return Region{}, err
        }
        // Check if region already has country
        if contains(region.countries, func(c Country) bool {
            return c.iso_2 == country.iso_2
        }) {
            return region, nil
        }
        region.countries = append(region.countries, country)
        updated, err := regionRepo.save(region)
        if err != nil {
            return Region{}, err
        }
        if err := s.eventBus_.withTransaction(manager).emit(RegionService.Events.UPDATED, map[string]interface{}{"id": updated.id, "fields": []string{"countries"}}); err != nil {
            return Region{}, err
        }
        return updated, nil
    })
    if err != nil {
        return Region{}, err
    }
    return result.(Region), nil
}

func (s *RegionService) removeCountry(regionId string, code string) (Region, error) {
    result, err := s.atomicPhase_(func(manager) (Region, error) {
        regionRepo := manager.withRepository(s.regionRepository_)
        region, err := s.retrieve(regionId, map[string]interface{}{"relations": []string{"countries"}})
        if err != nil {
            return Region{}, err
        }
        // Check if region contains country. If not, we simpy resolve
        if !contains(region.countries, func(c Country) bool {
            return c.iso_2 == code
        }) {
            return region, nil
        }
        region.countries = filter(region.countries, func(c Country) bool {
            return c.iso_2 != code
        })
        updated, err := regionRepo.save(region)
        if err != nil {
            return Region{}, err
        }
        if err := s.eventBus_.withTransaction(manager).emit(RegionService.Events.UPDATED, map[string]interface{}{"id": updated.id, "fields": []string{"countries"}}); err != nil {
            return Region{}, err
        }
        return updated, nil
    })
    if err != nil {
        return Region{}, err
    }
    return result.(Region), nil
}

func (s *RegionService) addPaymentProvider(regionId string, providerId string) (Region, error) {
    result, err := s.atomicPhase_(func(manager) (Region, error) {
        regionRepo := manager.withRepository(s.regionRepository_)
        ppRepo := manager.withRepository(s.paymentProviderRepository_)
        region, err := s.retrieve(regionId, map[string]interface{}{"relations": []string{"payment_providers"}})
        if err != nil {
            return Region{}, err
        }
        // Check if region already has payment provider
        if contains(region.payment_providers, func(pp PaymentProvider) bool {
            return pp.id == providerId
        }) {
            return region, nil
        }
        pp, err := ppRepo.findOne(map[string]interface{}{"where": map[string]interface{}{"id": providerId}})
        if err != nil {
            return Region{}, MedusaError{Types: MedusaError.Types.NOT_FOUND, Message: fmt.Sprintf("Payment provider %s was not found", providerId)}
        }
        region.payment_providers = append(region.payment_providers, pp)
        updated, err := regionRepo.save(region)
        if err != nil {
            return Region{}, err
        }
        if err := s.eventBus_.withTransaction(manager).emit(RegionService.Events.UPDATED, map[string]interface{}{"id": updated.id, "fields": []string{"payment_providers"}}); err != nil {
            return Region{}, err
        }
        return updated, nil
    })
    if err != nil {
        return Region{}, err
    }
    return result.(Region), nil
}

func (s *RegionService) addFulfillmentProvider(regionId string, providerId string) (Region, error) {
    result, err := s.atomicPhase_(func(manager) (Region, error) {
        regionRepo := manager.withRepository(s.regionRepository_)
        fpRepo := manager.withRepository(s.fulfillmentProviderRepository_)
        region, err := s.retrieve(regionId, map[string]interface{}{"relations": []string{"fulfillment_providers"}})
        if err != nil {
            return Region{}, err
        }
        // Check if region already has payment provider
        if contains(region.fulfillment_providers, func(fp FulfillmentProvider) bool {
            return fp.id == providerId
        }) {
            return region, nil
        }
        fp, err := fpRepo.findOne(map[string]interface{}{"where": map[string]interface{}{"id": providerId}})
        if err != nil {
            return Region{}, MedusaError{Types: MedusaError.Types.NOT_FOUND, Message: fmt.Sprintf("Fulfillment provider %s was not found", providerId)}
        }
        region.fulfillment_providers = append(region.fulfillment_providers, fp)
        updated, err := regionRepo.save(region)
        if err != nil {
            return Region{}, err
        }
        if err := s.eventBus_.withTransaction(manager).emit(RegionService.Events.UPDATED, map[string]interface{}{"id": updated.id, "fields": []string{"fulfillment_providers"}}); err != nil {
            return Region{}, err
        }
        return updated, nil
    })
    if err != nil {
        return Region{}, err
    }
    return result.(Region), nil
}

func (s *RegionService) removePaymentProvider(regionId string, providerId string) (Region, error) {
    result, err := s.atomicPhase_(func(manager) (Region, error) {
        regionRepo := manager.withRepository(s.regionRepository_)
        region, err := s.retrieve(regionId, map[string]interface{}{"relations": []string{"payment_providers"}})
        if err != nil {
            return Region{}, err
        }
        // Check if region already has payment provider
        if !contains(region.payment_providers, func(pp PaymentProvider) bool {
            return pp.id == providerId
        }) {
            return region, nil
        }
        region.payment_providers = filter(region.payment_providers, func(pp PaymentProvider) bool {
            return pp.id != providerId
        })
        updated, err := regionRepo.save(region)
        if err != nil {
            return Region{}, err
        }
        if err := s.eventBus_.withTransaction(manager).emit(RegionService.Events.UPDATED, map[string]interface{}{"id": updated.id, "fields": []string{"payment_providers"}}); err != nil {
            return Region{}, err
        }
        return updated, nil
    })
    if err != nil {
        return Region{}, err
    }
    return result.(Region), nil
}

func (s *RegionService) removeFulfillmentProvider(regionId string, providerId string) (Region, error) {
    result, err := s.atomicPhase_(func(manager) (Region, error) {
        regionRepo := manager.withRepository(s.regionRepository_)
        region, err := s.retrieve(regionId, map[string]interface{}{"relations": []string{"fulfillment_providers"}})
        if err != nil {
            return Region{}, err
        }
        // Check if region already has payment provider
        if !contains(region.fulfillment_providers, func(fp FulfillmentProvider) bool {
            return fp.id == providerId
        }) {
            return region, nil
        }
        region.fulfillment_providers = filter(region.fulfillment_providers, func(fp FulfillmentProvider) bool {
            return fp.id != providerId
        })
        updated, err := regionRepo.save(region)
        if err != nil {
            return Region{}, err
        }
        if err := s.eventBus_.withTransaction(manager).emit(RegionService.Events.UPDATED, map[string]interface{}{"id": updated.id, "fields": []string{"fulfillment_providers"}}); err != nil {
            return Region{}, err
        }
        return updated, nil
    })
    if err != nil {
        return Region{}, err
    }
    return result.(Region), nil
}




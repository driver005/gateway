package services

import (
	"context"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type RegionService struct {
	ctx context.Context
	r   Registry
}

func NewRegionService(
	r Registry,
) *RegionService {
	return &RegionService{
		context.Background(),
		r,
	}
}

func (s *RegionService) SetContext(context context.Context) *RegionService {
	s.ctx = context
	return s
}

func (s *RegionService) Create(data *types.CreateRegionInput) (*models.Region, *utils.ApplictaionError) {
	validated, err := s.validateFields(&types.UpdateRegionInput{
		Name:                 data.Name,
		CurrencyCode:         data.CurrencyCode,
		TaxCode:              data.TaxCode,
		TaxRate:              data.TaxRate,
		PaymentProviders:     data.PaymentProviders,
		FulfillmentProviders: data.FulfillmentProviders,
		Countries:            data.Countries,
		IncludesTax:          data.IncludesTax,
		Metadata:             data.Metadata,
	}, uuid.Nil)
	if err != nil {
		return nil, err
	}

	feature := true

	if feature {
		if !reflect.ValueOf(data.IncludesTax).IsZero() {
			validated.IncludesTax = data.IncludesTax
		}
	}
	if !reflect.ValueOf(data.CurrencyCode).IsZero() {
		if err := s.validateCurrency(data.CurrencyCode); err != nil {
			return nil, err
		}

		var currency *models.Currency = &models.Currency{}
		query := sql.BuildQuery(models.Currency{Code: strings.ToLower(data.CurrencyCode)}, &sql.Options{})
		if err := s.r.CurrencyRepository().FindOne(s.ctx, currency, query); err != nil {
			return nil, err
		}

		validated.Currency = currency
		validated.CurrencyCode = strings.ToLower(data.CurrencyCode)
	}

	if data.Metadata != nil {
		validated.Metadata = utils.MergeMaps(validated.Metadata, data.Metadata)
	}

	if err := s.r.RegionRepository().Save(s.ctx, validated); err != nil {
		return nil, err
	}
	// if err := s.eventBus_.withTransaction(manager).emit(RegionService.Events.CREATED, map[string]interface{}{"id": result.id}); err != nil {
	// 	return nil, err
	// }
	return validated, nil
}

func (s *RegionService) Update(regionId uuid.UUID, Update *types.UpdateRegionInput) (*models.Region, *utils.ApplictaionError) {
	region, err := s.Retrieve(regionId, &sql.Options{})
	if err != nil {
		return nil, err
	}

	validated, err := s.validateFields(Update, region.Id)
	if err != nil {
		return nil, err
	}
	feature := true

	if feature {
		if Update.IncludesTax {
			validated.IncludesTax = Update.IncludesTax
		}
	}
	if Update.CurrencyCode != "" {
		if err := s.validateCurrency(Update.CurrencyCode); err != nil {
			return nil, err
		}
		var currency *models.Currency = &models.Currency{}
		query := sql.BuildQuery(models.Currency{Code: strings.ToLower(Update.CurrencyCode)}, &sql.Options{})
		if err := s.r.CurrencyRepository().FindOne(s.ctx, currency, query); err != nil {
			return nil, err
		}

		validated.Currency = currency
		validated.CurrencyCode = strings.ToLower(Update.CurrencyCode)
	}

	validated.Id = region.Id

	if err := s.r.RegionRepository().Upsert(s.ctx, validated); err != nil {
		return nil, err
	}
	// if err := s.eventBus_.withTransaction(manager).emit(RegionService.Events.UPDATED, map[string]interface{}{"id": result.id, "fields": Object.keys(Update)}); err != nil {
	// 	return nil, err
	// }
	return validated, nil

}

func (s *RegionService) validateFields(data *types.UpdateRegionInput, id uuid.UUID) (*models.Region, *utils.ApplictaionError) {
	region := &models.Region{
		Model: core.Model{
			Metadata: data.Metadata,
		},
		Name:             data.Name,
		CurrencyCode:     data.CurrencyCode,
		TaxCode:          data.TaxCode,
		TaxRate:          data.TaxRate,
		GiftCardsTaxable: data.GiftCardsTaxable,
		AutomaticTaxes:   data.AutomaticTaxes,
		IncludesTax:      data.IncludesTax,
	}

	if err := s.validateTaxRate(data.TaxRate); err != nil {
		return nil, err
	}

	if data.Countries != nil {
		for _, countries := range data.Countries {
			validate, err := s.validateCountry(countries, id)
			if err != nil {
				return nil, err
			}

			region.Countries = append(region.Countries, *validate)
		}
	}

	if data.TaxProviderId != uuid.Nil {
		var provider *models.TaxProvider = &models.TaxProvider{}
		query := sql.BuildQuery(models.TaxProvider{Model: core.Model{Id: data.TaxProviderId}}, &sql.Options{})
		err := s.r.TaxProviderRepository().FindOne(s.ctx, provider, query)
		if err != nil {
			return nil, err
		}

		region.TaxProviderId = uuid.NullUUID{UUID: data.TaxProviderId}
		region.TaxProvider = provider
	}

	if data.PaymentProviders != nil {
		for _, paymentProvider := range data.PaymentProviders {
			var provider *models.PaymentProvider = &models.PaymentProvider{}
			query := sql.BuildQuery(models.PaymentProvider{Model: core.Model{Id: paymentProvider}}, &sql.Options{})
			err := s.r.PaymentProviderRepository().FindOne(s.ctx, provider, query)
			if err != nil {
				return nil, err
			}

			region.PaymentProviders = append(region.PaymentProviders, *provider)
		}
	}

	if data.FulfillmentProviders != nil {
		for _, fulfillmentProviders := range data.FulfillmentProviders {
			var provider *models.FulfillmentProvider = &models.FulfillmentProvider{}
			query := sql.BuildQuery(models.PaymentProvider{Model: core.Model{Id: fulfillmentProviders}}, &sql.Options{})
			err := s.r.FulfillmentProviderRepository().FindOne(s.ctx, provider, query)
			if err != nil {
				return nil, err
			}

			region.FulfillmentProviders = append(region.FulfillmentProviders, *provider)
		}
	}

	return region, nil
}

func (s *RegionService) validateTaxRate(taxRate float64) *utils.ApplictaionError {
	if taxRate > 100 || taxRate < 0 {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			"The tax_rate must be between 0 and 1",
			nil,
		)
	}
	return nil
}

func (s *RegionService) validateCurrency(currencyCode string) *utils.ApplictaionError {
	store, err := s.r.StoreService().SetContext(s.ctx).Retrieve(&sql.Options{Relations: []string{"currencies"}})
	if err != nil {
		return err
	}

	var storeCurrencies []string
	for _, currency := range store.Currencies {
		storeCurrencies = append(storeCurrencies, currency.Code)
	}

	if !slices.Contains(storeCurrencies, strings.ToLower(currencyCode)) {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Invalid currency code",
			nil,
		)
	}
	return nil
}

func (s *RegionService) validateCountry(code string, regionId uuid.UUID) (*models.Country, *utils.ApplictaionError) {
	countryCode := strings.ToUpper(code)

	isCountryExists := false
	for _, country := range utils.Countries {
		if country.Alpha2 == countryCode {
			isCountryExists = true
		}
	}

	if !isCountryExists {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Invalid country code",
			nil,
		)
	}

	var country *models.Country
	query := sql.BuildQuery(models.Country{Iso2: strings.ToLower(code)}, &sql.Options{})
	if err := s.r.CountryRepository().FindOne(s.ctx, country, query); err != nil {
		return nil, err
	}
	if country.RegionId.UUID != regionId {
		return nil, utils.NewApplictaionError(
			utils.DUPLICATE_ERROR,
			fmt.Sprintf("%s already exists in region %s", country.DisplayName, country.RegionId.UUID),
			nil,
		)
	}
	return country, nil
}

func (s *RegionService) RetrieveByCountryCode(code string, config *sql.Options) (*models.Region, *utils.ApplictaionError) {
	var country *models.Country
	query := sql.BuildQuery(models.Country{Iso2: strings.ToLower(code)}, &sql.Options{})
	if err := s.r.CountryRepository().FindOne(s.ctx, country, query); err != nil {
		return nil, err
	}

	if country.RegionId.UUID == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Country does not belong to a region",
			nil,
		)
	}
	return s.Retrieve(country.RegionId.UUID, config)
}

func (s *RegionService) RetrieveByName(name string) (*models.Region, *utils.ApplictaionError) {
	regions, err := s.List(&types.FilterableRegion{Name: name}, &sql.Options{Take: 1})
	if err != nil {
		return nil, err
	}
	if len(regions) == 0 {
		return nil, utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintf("Region \"%s\" was not found", name),
			nil,
		)
	}
	return &regions[0], nil
}

func (s *RegionService) Retrieve(regionId uuid.UUID, config *sql.Options) (*models.Region, *utils.ApplictaionError) {
	if regionId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"regionId" must be defined`,
			nil,
		)
	}

	var res *models.Region = &models.Region{}
	query := sql.BuildQuery(models.Region{Model: core.Model{Id: regionId}}, config)
	if err := s.r.RegionRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`Region with id `+regionId.String()+` was not found`,
			nil,
		)
	}
	return res, nil
}

func (s *RegionService) List(selector *types.FilterableRegion, config *sql.Options) ([]models.Region, *utils.ApplictaionError) {
	customerGroups, _, err := s.ListAndCount(selector, config)
	if err != nil {
		return nil, err
	}
	return customerGroups, nil
}

func (s *RegionService) ListAndCount(selector *types.FilterableRegion, config *sql.Options) ([]models.Region, *int64, *utils.ApplictaionError) {
	var res []models.Region

	query := sql.BuildQuery(selector, config)

	count, err := s.r.RegionRepository().FindAndCount(s.ctx, &res, query)
	if err != nil {
		return nil, nil, err
	}

	return res, count, nil
}

func (s *RegionService) Delete(regionId uuid.UUID) *utils.ApplictaionError {
	data, err := s.Retrieve(regionId, &sql.Options{
		Relations: []string{"countries"},
	})
	if err != nil {
		return err
	}

	if err := s.r.RegionRepository().SoftRemove(s.ctx, data); err != nil {
		return err
	}

	for _, country := range data.Countries {
		country.RegionId = uuid.NullUUID{}
		country.Region = &models.Region{}

		if err := s.r.CountryRepository().Update(s.ctx, &country); err != nil {
			return err
		}
	}

	// if err := s.eventBus_.withTransaction(manager).emit(RegionService.Events.DELETED, map[string]interface{}{"id": regionId}); err != nil {
	// 		return nil, err
	// 	}

	return nil
}

func (s *RegionService) AddCountry(regionId uuid.UUID, code string) (*models.Region, *utils.ApplictaionError) {
	country, err := s.validateCountry(code, regionId)
	if err != nil {
		return nil, err
	}
	region, err := s.Retrieve(regionId, &sql.Options{Relations: []string{"countries"}})
	if err != nil {
		return nil, err
	}
	// Check if region already has country
	if slices.ContainsFunc(region.Countries, func(c models.Country) bool {
		return c.Iso2 == country.Iso2
	}) {
		return region, nil
	}
	region.Countries = append(region.Countries, *country)
	if err := s.r.RegionRepository().Update(s.ctx, region); err != nil {
		return nil, err
	}
	// if err := s.eventBus_.withTransaction(manager).emit(RegionService.Events.UPDATED, map[string]interface{}{"id": updated.id, "fields": []string{"countries"}}); err != nil {
	// 	return nil, err
	// }
	return region, nil
}

func (s *RegionService) RemoveCountry(regionId uuid.UUID, code string) (*models.Region, *utils.ApplictaionError) {
	region, err := s.Retrieve(regionId, &sql.Options{Relations: []string{"countries"}})
	if err != nil {
		return nil, err
	}

	if !slices.ContainsFunc(region.Countries, func(c models.Country) bool {
		return c.Iso2 == code
	}) {
		return region, nil
	}

	var countries []models.Country
	for _, c := range region.Countries {
		if c.Iso2 != code {
			countries = append(countries, c)
		}
	}
	region.Countries = countries

	if err := s.r.RegionRepository().Update(s.ctx, region); err != nil {
		return nil, err
	}
	// if err := s.eventBus_.withTransaction(manager).emit(RegionService.Events.UPDATED, map[string]interface{}{"id": updated.id, "fields": []string{"countries"}}); err != nil {
	// 	return nil, err
	// }
	return region, nil
}

func (s *RegionService) AddPaymentProvider(regionId uuid.UUID, providerId uuid.UUID) (*models.Region, *utils.ApplictaionError) {
	region, err := s.Retrieve(regionId, &sql.Options{Relations: []string{"payment_providers"}})
	if err != nil {
		return nil, err
	}
	// Check if region already has payment provider
	if slices.ContainsFunc(region.PaymentProviders, func(pp models.PaymentProvider) bool {
		return pp.Id == providerId
	}) {
		return region, nil
	}

	var paymentProvider *models.PaymentProvider = &models.PaymentProvider{}
	query := sql.BuildQuery(models.PaymentProvider{Model: core.Model{Id: providerId}}, &sql.Options{})
	if err := s.r.PaymentProviderRepository().FindOne(s.ctx, paymentProvider, query); err != nil {
		return nil, err
	}

	region.PaymentProviders = append(region.PaymentProviders, *paymentProvider)

	if err := s.r.RegionRepository().Update(s.ctx, region); err != nil {
		return nil, err
	}
	// if err := s.eventBus_.withTransaction(manager).emit(RegionService.Events.UPDATED, map[string]interface{}{"id": updated.id, "fields": []string{"payment_providers"}}); err != nil {
	// 	return nil, err
	// }
	return region, nil
}

func (s *RegionService) RemovePaymentProvider(regionId uuid.UUID, providerId uuid.UUID) (*models.Region, *utils.ApplictaionError) {
	region, err := s.Retrieve(regionId, &sql.Options{Relations: []string{"payment_providers"}})
	if err != nil {
		return nil, err
	}
	// Check if region already has payment provider
	if !slices.ContainsFunc(region.PaymentProviders, func(pp models.PaymentProvider) bool {
		return pp.Id == providerId
	}) {
		return region, nil
	}

	var paymentProviders []models.PaymentProvider
	for _, pp := range region.PaymentProviders {
		if pp.Id != providerId {
			paymentProviders = append(paymentProviders, pp)
		}
	}

	region.PaymentProviders = paymentProviders

	if err := s.r.RegionRepository().Update(s.ctx, region); err != nil {
		return nil, err
	}
	// if err := s.eventBus_.withTransaction(manager).emit(RegionService.Events.UPDATED, map[string]interface{}{"id": updated.id, "fields": []string{"payment_providers"}}); err != nil {
	// 	return nil, err
	// }
	return region, nil
}

func (s *RegionService) AddFulfillmentProvider(regionId uuid.UUID, providerId uuid.UUID) (*models.Region, *utils.ApplictaionError) {
	region, err := s.Retrieve(regionId, &sql.Options{Relations: []string{"fulfillment_providers"}})
	if err != nil {
		return nil, err
	}
	// Check if region already has payment provider
	if slices.ContainsFunc(region.FulfillmentProviders, func(fp models.FulfillmentProvider) bool {
		return fp.Id == providerId
	}) {
		return region, nil
	}

	var fulfillmentProvider *models.FulfillmentProvider = &models.FulfillmentProvider{}
	query := sql.BuildQuery(models.FulfillmentProvider{Model: core.Model{Id: providerId}}, &sql.Options{})
	if err := s.r.FulfillmentProviderRepository().FindOne(s.ctx, fulfillmentProvider, query); err != nil {
		return nil, err
	}

	region.FulfillmentProviders = append(region.FulfillmentProviders, *fulfillmentProvider)

	if err := s.r.RegionRepository().Update(s.ctx, region); err != nil {
		return nil, err
	}
	// if err := s.eventBus_.withTransaction(manager).emit(RegionService.Events.UPDATED, map[string]interface{}{"id": updated.id, "fields": []string{"fulfillment_providers"}}); err != nil {
	// 	return nil, err
	// }
	return region, nil
}

func (s *RegionService) RemoveFulfillmentProvider(regionId uuid.UUID, providerId uuid.UUID) (*models.Region, *utils.ApplictaionError) {
	region, err := s.Retrieve(regionId, &sql.Options{Relations: []string{"fulfillment_providers"}})
	if err != nil {
		return nil, err
	}
	// Check if region already has payment provider
	if !slices.ContainsFunc(region.FulfillmentProviders, func(fp models.FulfillmentProvider) bool {
		return fp.Id == providerId
	}) {
		return region, nil
	}

	var fulfillmentProviders []models.FulfillmentProvider
	for _, fp := range region.FulfillmentProviders {
		if fp.Id != providerId {
			fulfillmentProviders = append(fulfillmentProviders, fp)
		}
	}

	region.FulfillmentProviders = fulfillmentProviders

	if err := s.r.RegionRepository().Update(s.ctx, region); err != nil {
		return nil, err
	}
	// if err := s.eventBus_.withTransaction(manager).emit(RegionService.Events.UPDATED, map[string]interface{}{"id": updated.id, "fields": []string{"fulfillment_providers"}}); err != nil {
	// 	return nil, err
	// }
	return region, nil
}

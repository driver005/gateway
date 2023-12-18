package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type FulfillmentProviderService struct {
	ctx  context.Context
	repo *repository.FulfillmentProviderRepo
}

func NewFulfillmentProviderService(
	ctx context.Context,
	repo *repository.FulfillmentProviderRepo,
) *FulfillmentProviderService {
	return &FulfillmentProviderService{
		ctx,
		repo,
	}
}

func (s *FulfillmentProviderService) RegisterInstalledProviders(providers []string) error {
	return s.atomicPhase(func(manager *typeorm.EntityManager) error {
		fulfillmentProviderRepo := manager.WithRepository(s.fulfillmentProviderRepository)
		err := fulfillmentProviderRepo.Update(map[string]interface{}{}, map[string]interface{}{"is_installed": false})
		if err != nil {
			return err
		}
		for _, p := range providers {
			n := fulfillmentProviderRepo.Create(map[string]interface{}{"id": p, "is_installed": true})
			err := fulfillmentProviderRepo.Save(n)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *FulfillmentProviderService) List() ([]FulfillmentProvider, error) {
	fpRepo := s.activeManager.WithRepository(s.fulfillmentProviderRepository)
	return fpRepo.Find(map[string]interface{}{})
}

func (s *FulfillmentProviderService) ListFulfillmentOptions(providerIDs []string) ([]FulfillmentOptions, error) {
	var fulfillmentOptions []FulfillmentOptions
	for _, p := range providerIDs {
		provider := s.RetrieveProvider(p)
		options, err := provider.GetFulfillmentOptions()
		if err != nil {
			return nil, err
		}
		fulfillmentOptions = append(fulfillmentOptions, FulfillmentOptions{
			ProviderID: p,
			Options:    options,
		})
	}
	return fulfillmentOptions, nil
}

func (s *FulfillmentProviderService) RetrieveProvider(providerID string) BaseFulfillmentService {
	provider, ok := s.container[providerID]
	if !ok {
		panic("Could not find a fulfillment provider with id: " + providerID)
	}
	return provider
}

func (s *FulfillmentProviderService) CreateFulfillment(method ShippingMethod, items []LineItem, order CreateFulfillmentOrder, fulfillment Fulfillment) (map[string]interface{}, error) {
	provider := s.RetrieveProvider(method.ShippingOption.ProviderID)
	return provider.CreateFulfillment(method.Data, items, order, fulfillment).(map[string]interface{}), nil
}

func (s *FulfillmentProviderService) CanCalculate(option CalculateOptionPriceInput) (bool, error) {
	provider := s.RetrieveProvider(option.ProviderID)
	return provider.CanCalculate(option.Data).(bool), nil
}

func (s *FulfillmentProviderService) ValidateFulfillmentData(option ShippingOption, data map[string]interface{}, cart Cart) (map[string]interface{}, error) {
	provider := s.RetrieveProvider(option.ProviderID)
	return provider.ValidateFulfillmentData(option.Data, data, cart).(map[string]interface{}), nil
}

func (s *FulfillmentProviderService) CancelFulfillment(fulfillment Fulfillment) (Fulfillment, error) {
	provider := s.RetrieveProvider(fulfillment.ProviderID)
	return provider.CancelFulfillment(fulfillment.Data).(Fulfillment), nil
}

func (s *FulfillmentProviderService) CalculatePrice(option ShippingOption, data map[string]interface{}, cart Order) (float64, error) {
	provider := s.RetrieveProvider(option.ProviderID)
	return provider.CalculatePrice(option.Data, data, cart).(float64), nil
}

func (s *FulfillmentProviderService) ValidateOption(option ShippingOption) (bool, error) {
	provider := s.RetrieveProvider(option.ProviderID)
	return provider.ValidateOption(option.Data).(bool), nil
}

func (s *FulfillmentProviderService) CreateReturn(returnOrder CreateReturnType) (map[string]interface{}, error) {
	option := returnOrder.ShippingMethod.ShippingOption
	provider := s.RetrieveProvider(option.ProviderID)
	return provider.CreateReturn(returnOrder).(map[string]interface{}), nil
}

func (s *FulfillmentProviderService) RetrieveDocuments(providerID string, fulfillmentData map[string]interface{}, documentType string) (interface{}, error) {
	provider := s.RetrieveProvider(providerID)
	return provider.RetrieveDocuments(fulfillmentData, documentType), nil
}

func main() {
	// Create a container
	container := &FulfillmentProviderContainer{
		fulfillmentProviderRepository: &FulfillmentProviderRepository{},
		manager:                       &typeorm.EntityManager{},
	}

	// Create a service
	service := NewFulfillmentProviderService(container)

	// Use the service
	providers := []string{"provider1", "provider2"}
	err := service.RegisterInstalledProviders(providers)
	if err != nil {
		panic(err)
	}

	fulfillmentOptions, err := service.ListFulfillmentOptions([]string{"provider1", "provider2"})
	if err != nil {
		panic(err)
	}

	provider := service.RetrieveProvider("provider1")
	fulfillment, err := provider.CreateFulfillment(ShippingMethod{}, []LineItem{}, CreateFulfillmentOrder{}, Fulfillment{})
	if err != nil {
		panic(err)
	}

	canCalculate, err := provider.CanCalculate(CalculateOptionPriceInput{})
	if err != nil {
		panic(err)
	}

	validatedData, err := provider.ValidateFulfillmentData(ShippingOption{}, map[string]interface{}{}, Cart{})
	if err != nil {
		panic(err)
	}

	cancelledFulfillment, err := provider.CancelFulfillment(Fulfillment{})
	if err != nil {
		panic(err)
	}

	price, err := provider.CalculatePrice(ShippingOption{}, map[string]interface{}{}, Order{})
	if err != nil {
		panic(err)
	}

	validOption, err := provider.ValidateOption(ShippingOption{})
	if err != nil {
		panic(err)
	}

	returnOrder, err := provider.CreateReturn(CreateReturnType{})
	if err != nil {
		panic(err)
	}

	documents, err := service.RetrieveDocuments("provider1", map[string]interface{}{}, "invoice")
	if err != nil {
		panic(err)
	}
}

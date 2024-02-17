package services

import (
	"context"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/sarulabs/di"
)

type FulfillmentProviderService struct {
	ctx       context.Context
	container di.Container
	r         Registry
}

func NewFulfillmentProviderService(
	container di.Container,
	r Registry,
) *FulfillmentProviderService {
	return &FulfillmentProviderService{
		context.Background(),
		container,
		r,
	}
}

func (s *FulfillmentProviderService) SetContext(context context.Context) *FulfillmentProviderService {
	s.ctx = context
	return s
}

func (s *FulfillmentProviderService) RetrieveProvider(providerID uuid.UUID) interfaces.IFulfillmentService {
	var provider interfaces.IFulfillmentService
	objectInterface, err := s.container.SafeGet(providerID.String())
	if err != nil {
		panic("Could not find a fulfillment provider with id: " + providerID.String())
	}
	provider, _ = objectInterface.(interfaces.IFulfillmentService)

	return provider
}

func (s *FulfillmentProviderService) RegisterInstalledProviders(providers uuid.UUIDs) *utils.ApplictaionError {
	if err := s.r.FulfillmentProviderRepository().Update(s.ctx, &models.FulfillmentProvider{IsInstalled: false}); err != nil {
		return err
	}

	for _, p := range providers {
		var model *models.FulfillmentProvider
		model.IsInstalled = true
		model.Id = p

		if err := s.r.FulfillmentProviderRepository().Save(s.ctx, model); err != nil {
			return err
		}
	}

	return nil

}

func (s *FulfillmentProviderService) List() ([]models.FulfillmentProvider, *utils.ApplictaionError) {
	var res []models.FulfillmentProvider
	if err := s.r.FulfillmentProviderRepository().Find(s.ctx, &res, sql.Query{}); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *FulfillmentProviderService) ListFulfillmentOptions(providerIDs uuid.UUIDs) ([]types.FulfillmentOptions, *utils.ApplictaionError) {
	var fulfillmentOptions []types.FulfillmentOptions
	for _, p := range providerIDs {
		provider := s.RetrieveProvider(p)
		options, err := provider.GetFulfillmentOptions()
		if err != nil {
			return nil, err
		}
		fulfillmentOptions = append(fulfillmentOptions, types.FulfillmentOptions{
			ProviderId: p,
			Options:    options,
		})
	}
	return fulfillmentOptions, nil
}

func (s *FulfillmentProviderService) CreateFulfillment(method *models.ShippingMethod, items []models.LineItem, order *types.CreateFulfillmentOrder, fulfillment *models.Fulfillment) (core.JSONB, *utils.ApplictaionError) {
	provider := s.RetrieveProvider(method.ShippingOption.ProviderId.UUID)
	return provider.CreateFulfillment(method, items, order, fulfillment), nil
}

func (s *FulfillmentProviderService) CanCalculate(option types.FulfillmentOptions) (bool, *utils.ApplictaionError) {
	provider := s.RetrieveProvider(option.ProviderId)
	return provider.CanCalculate(option.Options), nil
}

func (s *FulfillmentProviderService) ValidateFulfillmentData(option *models.ShippingOption, data map[string]interface{}, cart *models.Cart) (map[string]interface{}, *utils.ApplictaionError) {
	provider := s.RetrieveProvider(option.ProviderId.UUID)
	return provider.ValidateFulfillmentData(option, data, cart), nil
}

func (s *FulfillmentProviderService) CancelFulfillment(fulfillment *models.Fulfillment) (*models.Fulfillment, *utils.ApplictaionError) {
	provider := s.RetrieveProvider(fulfillment.ProviderId.UUID)
	return provider.CancelFulfillment(fulfillment), nil
}

func (s *FulfillmentProviderService) CalculatePrice(option *models.ShippingOption, data map[string]interface{}, cart *models.Cart) (float64, *utils.ApplictaionError) {
	provider := s.RetrieveProvider(option.ProviderId.UUID)
	return provider.CalculatePrice(option, data, cart), nil
}

func (s *FulfillmentProviderService) ValidateOption(option models.ShippingOption) (bool, *utils.ApplictaionError) {
	provider := s.RetrieveProvider(option.ProviderId.UUID)
	return provider.ValidateOption(option), nil
}

func (s *FulfillmentProviderService) CreateReturn(returnOrder *models.Return) (core.JSONB, *utils.ApplictaionError) {
	option := returnOrder.ShippingMethod.ShippingOption
	provider := s.RetrieveProvider(option.ProviderId.UUID)
	return provider.CreateReturn(returnOrder), nil
}

func (s *FulfillmentProviderService) RetrieveDocuments(providerId uuid.UUID, fulfillmentData map[string]interface{}, documentType string) (map[string]interface{}, *utils.ApplictaionError) {
	provider := s.RetrieveProvider(providerId)
	return provider.RetrieveDocuments(fulfillmentData, documentType), nil
}

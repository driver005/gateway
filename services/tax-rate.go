package services

import (
	"context"
	"errors"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/repository"
	"github.com/driver005/gateway/types"
	"github.com/google/uuid"
)

type TaxRateService struct {
	ctx                   context.Context
	repo                  *repository.TaxRateRepo
	productService        ProductService
	productTypeService    ProductTypeService
	shippingOptionService ShippingOptionService
}

func NewTaxRateService(
	ctx context.Context,
	repo *repository.TaxRateRepo,
	productService ProductService,
	productTypeService ProductTypeService,
	shippingOptionService ShippingOptionService,
) *TaxRateService {
	return &TaxRateService{
		ctx,
		repo,
		productService,
		productTypeService,
		shippingOptionService,
	}
}

func (s *TaxRateService) List(selector types.FilterableTaxRate, config repository.Options) ([]models.TaxRate, error) {
	query := repository.BuildQuery[types.FilterableTaxRate](selector, config)
	return s.repo.FindWithResolution(query)
}

func (s *TaxRateService) ListAndCount(selector types.FilterableTaxRate, config repository.Options) ([]models.TaxRate, int64, error) {
	query := repository.BuildQuery[types.FilterableTaxRate](selector, config)
	taxRates, count, err := s.repo.FindAndCountWithResolution(query)
	return taxRates, *count, err
}

func (s *TaxRateService) Retrieve(taxRateId uuid.UUID, config repository.Options) (*models.TaxRate, error) {
	if taxRateId == uuid.Nil {
		return nil, errors.New(`"taxRateId" must be defined`)
	}
	query := repository.BuildQuery[types.FilterableTaxRate](types.FilterableTaxRate{FilterModel: core.FilterModel{Id: uuid.UUIDs{taxRateId}}}, config)
	taxRate, err := s.repo.FindOneWithResolution(query)
	if err != nil {
		return nil, err
	}
	if taxRate == nil {
		return nil, errors.New(`TaxRate with ` + taxRateId.String() + ` was not found`)
	}
	return taxRate, nil
}

func (s *TaxRateService) Create(model *models.TaxRate) (*models.TaxRate, error) {
	if model.RegionId.UUID == uuid.Nil {
		return nil, errors.New("TaxRates must belong to a Region")
	}

	if err := s.repo.Save(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *TaxRateService) Update(id uuid.UUID, data types.UpdateTaxRateInput) (*models.TaxRate, error) {
	taxRate, err := s.Retrieve(id, repository.Options{})
	if err != nil {
		return nil, err
	}

	taxRate.Code = data.Code
	taxRate.Name = data.Name
	taxRate.Rate = data.Rate
	taxRate.RegionId = uuid.NullUUID{UUID: data.RegionId}

	if err := s.repo.Save(s.ctx, taxRate); err != nil {
		return nil, err
	}

	return taxRate, nil
}

func (s *TaxRateService) Delete(id uuid.UUID) error {
	return s.repo.Delete(s.ctx, &models.TaxRate{Model: core.Model{Id: id}})
}

func (s *TaxRateService) RemoveFromProduct(id uuid.UUID, productIds uuid.UUIDs) error {
	return s.repo.RemoveFromProduct(id, productIds)
}

func (s *TaxRateService) RemoveFromProductType(id uuid.UUID, typeIds uuid.UUIDs) error {
	return s.repo.RemoveFromProductType(id, typeIds)
}

func (s *TaxRateService) RemoveFromShippingOption(id uuid.UUID, optionIds uuid.UUIDs) error {
	return s.repo.RemoveFromShippingOption(id, optionIds)
}

func (s *TaxRateService) AddToProduct(id uuid.UUID, productIds uuid.UUIDs, replace bool) ([]models.ProductTaxRate, error) {
	return s.repo.AddToProduct(id, productIds, replace)
}

func (s *TaxRateService) AddToProductType(id uuid.UUID, productTypeIds uuid.UUIDs, replace bool) ([]models.ProductTypeTaxRate, error) {
	return s.repo.AddToProductType(id, productTypeIds, replace)
}

func (s *TaxRateService) AddToShippingOption(id uuid.UUID, optionIds uuid.UUIDs, replace bool) ([]models.ShippingTaxRate, error) {
	taxRate, err := s.Retrieve(id, repository.Options{Selects: []string{"id", "region_id"}})
	if err != nil {
		return nil, err
	}
	options, err := s.shippingOptionService.List(models.ShippingOption{Model: core.Model{Id: id}}, repository.Options{Selects: []string{"id", "region_id"}})
	if err != nil {
		return nil, err
	}
	for _, o := range options {
		if o.RegionId != taxRate.RegionId {
			return nil, errors.New(`Shipping Option and Tax Rate must belong to the same Region to be associated. Shipping Option with id: ` + o.Id.String() + ` belongs to Region with id: ` + o.RegionId.UUID.String() + ` and Tax Rate with id: ` + taxRate.Id.String() + ` belongs to Region with id: ` + taxRate.RegionId.UUID.String())
		}
	}
	return s.repo.AddToShippingOption(id, optionIds, replace)
}

func (s *TaxRateService) ListByProduct(productId uuid.UUID, config types.TaxRateListByConfig) ([]models.TaxRate, error) {
	return s.repo.ListByProduct(productId, config)
}

func (s *TaxRateService) ListByShippingOption(shippingOptionId uuid.UUID) ([]models.TaxRate, error) {
	return s.repo.ListByShippingOption(shippingOptionId)
}

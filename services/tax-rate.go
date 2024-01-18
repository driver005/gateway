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
)

type TaxRateService struct {
	ctx context.Context
	r   Registry
}

func NewTaxRateService(
	r Registry,
) *TaxRateService {
	return &TaxRateService{
		context.Background(),
		r,
	}
}

func (s *TaxRateService) SetContext(context context.Context) *TaxRateService {
	s.ctx = context
	return s
}

func (s *TaxRateService) List(selector types.FilterableTaxRate, config *sql.Options) ([]models.TaxRate, *utils.ApplictaionError) {
	query := sql.BuildQuery[types.FilterableTaxRate](selector, config)
	return s.r.TaxRateRepository().FindWithResolution(query)
}

func (s *TaxRateService) ListAndCount(selector types.FilterableTaxRate, config *sql.Options) ([]models.TaxRate, int64, *utils.ApplictaionError) {
	query := sql.BuildQuery[types.FilterableTaxRate](selector, config)
	taxRates, count, err := s.r.TaxRateRepository().FindAndCountWithResolution(query)
	return taxRates, *count, err
}

func (s *TaxRateService) Retrieve(taxRateId uuid.UUID, config *sql.Options) (*models.TaxRate, *utils.ApplictaionError) {
	if taxRateId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"taxRateId" must be defined`,
			nil,
		)
	}
	query := sql.BuildQuery[types.FilterableTaxRate](types.FilterableTaxRate{FilterModel: core.FilterModel{Id: uuid.UUIDs{taxRateId}}}, config)
	taxRate, err := s.r.TaxRateRepository().FindOneWithResolution(query)
	if err != nil {
		return nil, err
	}
	if taxRate == nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`TaxRate with `+taxRateId.String()+` was not found`,
			nil,
		)
	}
	return taxRate, nil
}

func (s *TaxRateService) Create(data *types.CreateTaxRateInput) (*models.TaxRate, *utils.ApplictaionError) {
	if data.RegionId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"TaxRates must belong to a Region",
			nil,
		)
	}

	model := &models.TaxRate{
		RegionId: uuid.NullUUID{UUID: data.RegionId},
		Code:     data.Code,
		Name:     data.Name,
		Rate:     data.Rate,
	}

	if err := s.r.TaxRateRepository().Save(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *TaxRateService) Update(id uuid.UUID, data types.UpdateTaxRateInput) (*models.TaxRate, *utils.ApplictaionError) {
	taxRate, err := s.Retrieve(id, &sql.Options{})
	if err != nil {
		return nil, err
	}

	if !reflect.ValueOf(data.Code).IsZero() {
		taxRate.Code = data.Code
	}
	if !reflect.ValueOf(data.Name).IsZero() {
		taxRate.Name = data.Name
	}
	if !reflect.ValueOf(data.Rate).IsZero() {
		taxRate.Rate = data.Rate
	}
	if !reflect.ValueOf(data.RegionId).IsZero() {
		taxRate.RegionId = uuid.NullUUID{UUID: data.RegionId}
	}

	if err := s.r.TaxRateRepository().Save(s.ctx, taxRate); err != nil {
		return nil, err
	}

	return taxRate, nil
}

func (s *TaxRateService) Delete(id uuid.UUID) *utils.ApplictaionError {
	return s.r.TaxRateRepository().Delete(s.ctx, &models.TaxRate{Model: core.Model{Id: id}})
}

func (s *TaxRateService) RemoveFromProduct(id uuid.UUID, productIds uuid.UUIDs) *utils.ApplictaionError {
	return s.r.TaxRateRepository().RemoveFromProduct(id, productIds)
}

func (s *TaxRateService) RemoveFromProductType(id uuid.UUID, typeIds uuid.UUIDs) *utils.ApplictaionError {
	return s.r.TaxRateRepository().RemoveFromProductType(id, typeIds)
}

func (s *TaxRateService) RemoveFromShippingOption(id uuid.UUID, optionIds uuid.UUIDs) *utils.ApplictaionError {
	return s.r.TaxRateRepository().RemoveFromShippingOption(id, optionIds)
}

func (s *TaxRateService) AddToProduct(id uuid.UUID, productIds uuid.UUIDs, replace bool) ([]models.ProductTaxRate, *utils.ApplictaionError) {
	return s.r.TaxRateRepository().AddToProduct(id, productIds, replace)
}

func (s *TaxRateService) AddToProductType(id uuid.UUID, productTypeIds uuid.UUIDs, replace bool) ([]models.ProductTypeTaxRate, *utils.ApplictaionError) {
	return s.r.TaxRateRepository().AddToProductType(id, productTypeIds, replace)
}

func (s *TaxRateService) AddToShippingOption(id uuid.UUID, optionIds uuid.UUIDs, replace bool) ([]models.ShippingTaxRate, *utils.ApplictaionError) {
	taxRate, err := s.Retrieve(id, &sql.Options{Selects: []string{"id", "region_id"}})
	if err != nil {
		return nil, err
	}
	options, err := s.r.ShippingOptionService().SetContext(s.ctx).List(models.ShippingOption{Model: core.Model{Id: id}}, &sql.Options{Selects: []string{"id", "region_id"}})
	if err != nil {
		return nil, err
	}
	for _, o := range options {
		if o.RegionId != taxRate.RegionId {
			return nil, utils.NewApplictaionError(
				utils.INVALID_DATA,
				`Shipping Option and Tax Rate must belong to the same Region to be associated. Shipping Option with id: `+o.Id.String()+` belongs to Region with id: `+o.RegionId.UUID.String()+` and Tax Rate with id: `+taxRate.Id.String()+` belongs to Region with id: `+taxRate.RegionId.UUID.String(),
				"500",
				nil,
			)
		}
	}
	return s.r.TaxRateRepository().AddToShippingOption(id, optionIds, replace)
}

func (s *TaxRateService) ListByProduct(productId uuid.UUID, config types.TaxRateListByConfig) ([]models.TaxRate, *utils.ApplictaionError) {
	return s.r.TaxRateRepository().ListByProduct(productId, config)
}

func (s *TaxRateService) ListByShippingOption(shippingOptionId uuid.UUID) ([]models.TaxRate, *utils.ApplictaionError) {
	return s.r.TaxRateRepository().ListByShippingOption(shippingOptionId)
}

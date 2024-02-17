package services

import (
	"context"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
)

type ProductTaxRateService struct {
	ctx context.Context
	r   Registry
}

func NewProductTaxRateService(
	r Registry,
) *ProductTaxRateService {
	return &ProductTaxRateService{
		context.Background(),
		r,
	}
}

func (s *ProductTaxRateService) SetContext(context context.Context) *ProductTaxRateService {
	s.ctx = context
	return s
}

func (s *ProductTaxRateService) List(selector types.FilterableProductTaxRate, config *sql.Options) ([]models.ProductTaxRate, *utils.ApplictaionError) {
	var discounts []models.ProductTaxRate
	query := sql.BuildQuery(selector, config)
	if err := s.r.ProductTaxRateRepository().Find(s.ctx, &discounts, query); err != nil {
		return nil, err
	}
	return discounts, nil
}

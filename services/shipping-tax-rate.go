package services

import (
	"context"
	"reflect"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/icza/gox/gox"
)

type ShippingTaxRateService struct {
	ctx context.Context
	r   Registry
}

func NewShippingTaxRateService(
	r Registry,
) *ShippingTaxRateService {
	return &ShippingTaxRateService{
		context.Background(),
		r,
	}
}

func (s *ShippingTaxRateService) SetContext(context context.Context) *ShippingTaxRateService {
	s.ctx = context
	return s
}

func (s *ShippingTaxRateService) List(selector types.FilterableShippingTaxRate, config *sql.Options) ([]models.ShippingTaxRate, *utils.ApplictaionError) {
	var res []models.ShippingTaxRate

	if reflect.DeepEqual(config, &sql.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(50)
		config.Order = gox.NewString("created_at DESC")
	}

	query := sql.BuildQuery(selector, config)

	if err := s.r.ShippingTaxRateRepository().Find(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

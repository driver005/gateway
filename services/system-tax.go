package services

import (
	"context"

	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
)

type SystemTaxService struct {
	ctx context.Context
	r   Registry
}

func NewSystemTaxService(r Registry) *SystemTaxService {
	return &SystemTaxService{
		context.Background(),
		r,
	}
}

func (s *SystemTaxService) SetContext(context context.Context) *SystemTaxService {
	s.ctx = context
	return s
}

func (s *SystemProviderService) GetTaxLines(itemLines []interfaces.ItemTaxCalculationLine, shippingLines []interfaces.ShippingTaxCalculationLine, context interfaces.TaxCalculationContext) ([]types.ProviderTaxLine, *utils.ApplictaionError) {
	var model []types.ProviderTaxLine
	for _, item := range itemLines {
		for _, r := range item.Rates {
			model = append(model, types.ProviderTaxLine{
				Rate: r.Rate,
				Name: r.Name,
				Code: r.Code,
			})
		}

	}
	for _, item := range shippingLines {
		for _, r := range item.Rates {
			model = append(model, types.ProviderTaxLine{
				Rate: r.Rate,
				Name: r.Name,
				Code: r.Code,
			})
		}
	}

	return model, nil
}

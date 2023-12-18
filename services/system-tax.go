package services

import (
	"context"

	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/types"
)

type SystemTaxService struct {
	ctx context.Context
	interfaces.ITaxService
}

func NewSystemTaxService(ctx context.Context, taxService interfaces.ITaxService) *SystemTaxService {
	return &SystemTaxService{
		ctx,
		taxService,
	}
}

func (s *SystemProviderService) GetTaxLines(itemLines []interfaces.ItemTaxCalculationLine, shippingLines []interfaces.ShippingTaxCalculationLine, context interfaces.TaxCalculationContext) ([]types.ProviderTaxLine, error) {
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

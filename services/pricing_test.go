package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

func TestNewPricingService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *PricingService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPricingService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPricingService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPricingService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *PricingService
		args args
		want *PricingService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PricingService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPricingService_collectPricingContext(t *testing.T) {
	type args struct {
		context *interfaces.PricingContext
	}
	tests := []struct {
		name  string
		s     *PricingService
		args  args
		want  *interfaces.PricingContext
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.collectPricingContext(tt.args.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PricingService.collectPricingContext() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PricingService.collectPricingContext() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPricingService_calculateTaxes(t *testing.T) {
	type args struct {
		variantPricing types.ProductVariantPricing
		productRates   []types.TaxServiceRate
	}
	tests := []struct {
		name string
		s    *PricingService
		args args
		want types.TaxedPricing
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.calculateTaxes(tt.args.variantPricing, tt.args.productRates); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PricingService.calculateTaxes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPricingService_getProductVariantPricing(t *testing.T) {
	type args struct {
		data    []interfaces.Pricing
		context *interfaces.PricingContext
	}
	tests := []struct {
		name  string
		s     *PricingService
		args  args
		want  map[uuid.UUID]types.ProductVariantPricing
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.getProductVariantPricing(tt.args.data, tt.args.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PricingService.getProductVariantPricing() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PricingService.getProductVariantPricing() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPricingService_GetProductVariantPricingById(t *testing.T) {
	type args struct {
		variantId uuid.UUID
		context   *interfaces.PricingContext
	}
	tests := []struct {
		name  string
		s     *PricingService
		args  args
		want  *types.ProductVariantPricing
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetProductVariantPricingById(tt.args.variantId, tt.args.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PricingService.GetProductVariantPricingById() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PricingService.GetProductVariantPricingById() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPricingService_GetProductVariantsPricing(t *testing.T) {
	type args struct {
		data    []interfaces.Pricing
		context *interfaces.PricingContext
	}
	tests := []struct {
		name  string
		s     *PricingService
		args  args
		want  map[uuid.UUID]types.ProductVariantPricing
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetProductVariantsPricing(tt.args.data, tt.args.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PricingService.GetProductVariantsPricing() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PricingService.GetProductVariantsPricing() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPricingService_getProductPricing(t *testing.T) {
	type args struct {
		data    []interfaces.ProductPricing
		context *interfaces.PricingContext
	}
	tests := []struct {
		name  string
		s     *PricingService
		args  args
		want  map[uuid.UUID]map[uuid.UUID]types.ProductVariantPricing
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.getProductPricing(tt.args.data, tt.args.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PricingService.getProductPricing() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PricingService.getProductPricing() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPricingService_SetVariantPrices(t *testing.T) {
	type args struct {
		variants []models.ProductVariant
		context  *interfaces.PricingContext
	}
	tests := []struct {
		name  string
		s     *PricingService
		args  args
		want  []models.ProductVariant
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.SetVariantPrices(tt.args.variants, tt.args.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PricingService.SetVariantPrices() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PricingService.SetVariantPrices() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPricingService_SetProductPrices(t *testing.T) {
	type args struct {
		products []models.Product
		context  *interfaces.PricingContext
	}
	tests := []struct {
		name  string
		s     *PricingService
		args  args
		want  []models.Product
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.SetProductPrices(tt.args.products, tt.args.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PricingService.SetProductPrices() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PricingService.SetProductPrices() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPricingService_GetProductPricing(t *testing.T) {
	type args struct {
		product *models.Product
		context *interfaces.PricingContext
	}
	tests := []struct {
		name  string
		s     *PricingService
		args  args
		want  map[uuid.UUID]types.ProductVariantPricing
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetProductPricing(tt.args.product, tt.args.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PricingService.GetProductPricing() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PricingService.GetProductPricing() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPricingService_GetProductPricingById(t *testing.T) {
	type args struct {
		productId uuid.UUID
		context   *interfaces.PricingContext
	}
	tests := []struct {
		name  string
		s     *PricingService
		args  args
		want  map[uuid.UUID]types.ProductVariantPricing
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetProductPricingById(tt.args.productId, tt.args.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PricingService.GetProductPricingById() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PricingService.GetProductPricingById() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPricingService_getPricingModuleVariantMoneyAmounts(t *testing.T) {
	type args struct {
		variantIds []uuid.UUID
	}
	tests := []struct {
		name  string
		s     *PricingService
		args  args
		want  map[uuid.UUID][]models.MoneyAmount
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.getPricingModuleVariantMoneyAmounts(tt.args.variantIds)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PricingService.getPricingModuleVariantMoneyAmounts() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PricingService.getPricingModuleVariantMoneyAmounts() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPricingService_SetAdminVariantPricing(t *testing.T) {
	type args struct {
		variants []models.ProductVariant
		context  *interfaces.PricingContext
	}
	tests := []struct {
		name  string
		s     *PricingService
		args  args
		want  []models.ProductVariant
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.SetAdminVariantPricing(tt.args.variants, tt.args.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PricingService.SetAdminVariantPricing() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PricingService.SetAdminVariantPricing() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPricingService_SetAdminProductPricing(t *testing.T) {
	type args struct {
		products []models.Product
	}
	tests := []struct {
		name  string
		s     *PricingService
		args  args
		want  []models.Product
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.SetAdminProductPricing(tt.args.products)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PricingService.SetAdminProductPricing() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PricingService.SetAdminProductPricing() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPricingService_GetShippingOptionPricing(t *testing.T) {
	type args struct {
		shippingOption *models.ShippingOption
		context        *interfaces.PricingContext
	}
	tests := []struct {
		name  string
		s     *PricingService
		args  args
		want  *types.PricedShippingOption
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetShippingOptionPricing(tt.args.shippingOption, tt.args.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PricingService.GetShippingOptionPricing() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PricingService.GetShippingOptionPricing() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPricingService_SetShippingOptionPrices(t *testing.T) {
	type args struct {
		shippingOptions []models.ShippingOption
		context         *interfaces.PricingContext
	}
	tests := []struct {
		name  string
		s     *PricingService
		args  args
		want  []types.PricedShippingOption
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.SetShippingOptionPrices(tt.args.shippingOptions, tt.args.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PricingService.SetShippingOptionPrices() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PricingService.SetShippingOptionPrices() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

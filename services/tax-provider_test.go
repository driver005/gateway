package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/sarulabs/di"
)

func TestNewTaxProviderService(t *testing.T) {
	type args struct {
		container di.Container
		r         Registry
	}
	tests := []struct {
		name string
		args args
		want *TaxProviderService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTaxProviderService(tt.args.container, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaxProviderService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaxProviderService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *TaxProviderService
		args args
		want *TaxProviderService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxProviderService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaxProviderService_RetrieveProvider(t *testing.T) {
	type args struct {
		region models.Region
	}
	tests := []struct {
		name string
		s    *TaxProviderService
		args args
		want interfaces.ITaxService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.RetrieveProvider(tt.args.region); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxProviderService.RetrieveProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaxProviderService_List(t *testing.T) {
	type args struct {
		selector *models.TaxProvider
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *TaxProviderService
		args  args
		want  []models.TaxProvider
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxProviderService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TaxProviderService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTaxProviderService_ListAndCount(t *testing.T) {
	type args struct {
		selector *models.TaxProvider
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *TaxProviderService
		args  args
		want  []models.TaxProvider
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxProviderService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TaxProviderService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("TaxProviderService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestTaxProviderService_ClearLineItemsTaxLines(t *testing.T) {
	type args struct {
		itemIds uuid.UUIDs
	}
	tests := []struct {
		name string
		s    *TaxProviderService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ClearLineItemsTaxLines(tt.args.itemIds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxProviderService.ClearLineItemsTaxLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaxProviderService_ClearTaxLines(t *testing.T) {
	type args struct {
		cartId uuid.UUID
	}
	tests := []struct {
		name string
		s    *TaxProviderService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ClearTaxLines(tt.args.cartId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxProviderService.ClearTaxLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaxProviderService_CreateTaxLines(t *testing.T) {
	type args struct {
		cart               *models.Cart
		lineItems          []models.LineItem
		calculationContext *interfaces.TaxCalculationContext
	}
	tests := []struct {
		name  string
		s     *TaxProviderService
		args  args
		want  []models.ShippingMethodTaxLine
		want1 []models.LineItemTaxLine
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.CreateTaxLines(tt.args.cart, tt.args.lineItems, tt.args.calculationContext)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxProviderService.CreateTaxLines() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TaxProviderService.CreateTaxLines() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("TaxProviderService.CreateTaxLines() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestTaxProviderService_CreateShippingTaxLines(t *testing.T) {
	type args struct {
		shippingMethod     *models.ShippingMethod
		calculationContext *interfaces.TaxCalculationContext
	}
	tests := []struct {
		name  string
		s     *TaxProviderService
		args  args
		want  []models.ShippingMethodTaxLine
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateShippingTaxLines(tt.args.shippingMethod, tt.args.calculationContext)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxProviderService.CreateShippingTaxLines() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TaxProviderService.CreateShippingTaxLines() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTaxProviderService_GetShippingTaxLines(t *testing.T) {
	type args struct {
		shippingMethod     *models.ShippingMethod
		calculationContext *interfaces.TaxCalculationContext
	}
	tests := []struct {
		name  string
		s     *TaxProviderService
		args  args
		want  []models.ShippingMethodTaxLine
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetShippingTaxLines(tt.args.shippingMethod, tt.args.calculationContext)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxProviderService.GetShippingTaxLines() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TaxProviderService.GetShippingTaxLines() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTaxProviderService_GetTaxLines(t *testing.T) {
	type args struct {
		lineItems          []models.LineItem
		calculationContext *interfaces.TaxCalculationContext
	}
	tests := []struct {
		name  string
		s     *TaxProviderService
		args  args
		want  []models.ShippingMethodTaxLine
		want1 []models.LineItemTaxLine
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.GetTaxLines(tt.args.lineItems, tt.args.calculationContext)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxProviderService.GetTaxLines() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TaxProviderService.GetTaxLines() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("TaxProviderService.GetTaxLines() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestTaxProviderService_GetTaxLinesMap(t *testing.T) {
	type args struct {
		items              []models.LineItem
		calculationContext *interfaces.TaxCalculationContext
	}
	tests := []struct {
		name  string
		s     *TaxProviderService
		args  args
		want  types.TaxLinesMaps
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetTaxLinesMap(tt.args.items, tt.args.calculationContext)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxProviderService.GetTaxLinesMap() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TaxProviderService.GetTaxLinesMap() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTaxProviderService_GetRegionRatesForShipping(t *testing.T) {
	type args struct {
		optionId uuid.UUID
		region   *models.Region
	}
	tests := []struct {
		name  string
		s     *TaxProviderService
		args  args
		want  []types.TaxServiceRate
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetRegionRatesForShipping(tt.args.optionId, tt.args.region)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxProviderService.GetRegionRatesForShipping() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TaxProviderService.GetRegionRatesForShipping() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTaxProviderService_GetRegionRatesForProduct(t *testing.T) {
	type args struct {
		productIds uuid.UUIDs
		region     *models.Region
	}
	tests := []struct {
		name  string
		s     *TaxProviderService
		args  args
		want  map[uuid.UUID][]types.TaxServiceRate
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetRegionRatesForProduct(tt.args.productIds, tt.args.region)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxProviderService.GetRegionRatesForProduct() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TaxProviderService.GetRegionRatesForProduct() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTaxProviderService_GetCacheKey(t *testing.T) {
	type args struct {
		id       uuid.UUID
		regionId uuid.UUID
	}
	tests := []struct {
		name string
		s    *TaxProviderService
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetCacheKey(tt.args.id, tt.args.regionId); got != tt.want {
				t.Errorf("TaxProviderService.GetCacheKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaxProviderService_RegisterInstalledProviders(t *testing.T) {
	type args struct {
		providers uuid.UUIDs
	}
	tests := []struct {
		name string
		s    *TaxProviderService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.RegisterInstalledProviders(tt.args.providers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxProviderService.RegisterInstalledProviders() = %v, want %v", got, tt.want)
			}
		})
	}
}

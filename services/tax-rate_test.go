package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

func TestNewTaxRateService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *TaxRateService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTaxRateService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaxRateService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaxRateService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *TaxRateService
		args args
		want *TaxRateService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxRateService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaxRateService_List(t *testing.T) {
	type args struct {
		selector *types.FilterableTaxRate
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *TaxRateService
		args  args
		want  []models.TaxRate
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxRateService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TaxRateService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTaxRateService_ListAndCount(t *testing.T) {
	type args struct {
		selector *types.FilterableTaxRate
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *TaxRateService
		args  args
		want  []models.TaxRate
		want1 int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxRateService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("TaxRateService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("TaxRateService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestTaxRateService_Retrieve(t *testing.T) {
	type args struct {
		taxRateId uuid.UUID
		config    *sql.Options
	}
	tests := []struct {
		name  string
		s     *TaxRateService
		args  args
		want  *models.TaxRate
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.taxRateId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxRateService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TaxRateService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTaxRateService_Create(t *testing.T) {
	type args struct {
		data *types.CreateTaxRateInput
	}
	tests := []struct {
		name  string
		s     *TaxRateService
		args  args
		want  *models.TaxRate
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxRateService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TaxRateService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTaxRateService_Update(t *testing.T) {
	type args struct {
		id   uuid.UUID
		data *types.UpdateTaxRateInput
	}
	tests := []struct {
		name  string
		s     *TaxRateService
		args  args
		want  *models.TaxRate
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.id, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxRateService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TaxRateService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTaxRateService_Delete(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name string
		s    *TaxRateService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxRateService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaxRateService_RemoveFromProduct(t *testing.T) {
	type args struct {
		id         uuid.UUID
		productIds uuid.UUIDs
	}
	tests := []struct {
		name string
		s    *TaxRateService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.RemoveFromProduct(tt.args.id, tt.args.productIds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxRateService.RemoveFromProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaxRateService_RemoveFromProductType(t *testing.T) {
	type args struct {
		id      uuid.UUID
		typeIds uuid.UUIDs
	}
	tests := []struct {
		name string
		s    *TaxRateService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.RemoveFromProductType(tt.args.id, tt.args.typeIds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxRateService.RemoveFromProductType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaxRateService_RemoveFromShippingOption(t *testing.T) {
	type args struct {
		id        uuid.UUID
		optionIds uuid.UUIDs
	}
	tests := []struct {
		name string
		s    *TaxRateService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.RemoveFromShippingOption(tt.args.id, tt.args.optionIds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxRateService.RemoveFromShippingOption() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaxRateService_AddToProduct(t *testing.T) {
	type args struct {
		id         uuid.UUID
		productIds uuid.UUIDs
		replace    bool
	}
	tests := []struct {
		name  string
		s     *TaxRateService
		args  args
		want  []models.ProductTaxRate
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AddToProduct(tt.args.id, tt.args.productIds, tt.args.replace)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxRateService.AddToProduct() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TaxRateService.AddToProduct() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTaxRateService_AddToProductType(t *testing.T) {
	type args struct {
		id             uuid.UUID
		productTypeIds uuid.UUIDs
		replace        bool
	}
	tests := []struct {
		name  string
		s     *TaxRateService
		args  args
		want  []models.ProductTypeTaxRate
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AddToProductType(tt.args.id, tt.args.productTypeIds, tt.args.replace)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxRateService.AddToProductType() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TaxRateService.AddToProductType() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTaxRateService_AddToShippingOption(t *testing.T) {
	type args struct {
		id        uuid.UUID
		optionIds uuid.UUIDs
		replace   bool
	}
	tests := []struct {
		name  string
		s     *TaxRateService
		args  args
		want  []models.ShippingTaxRate
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AddToShippingOption(tt.args.id, tt.args.optionIds, tt.args.replace)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxRateService.AddToShippingOption() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TaxRateService.AddToShippingOption() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTaxRateService_ListByProduct(t *testing.T) {
	type args struct {
		productId uuid.UUID
		config    types.TaxRateListByConfig
	}
	tests := []struct {
		name  string
		s     *TaxRateService
		args  args
		want  []models.TaxRate
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ListByProduct(tt.args.productId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxRateService.ListByProduct() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TaxRateService.ListByProduct() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTaxRateService_ListByShippingOption(t *testing.T) {
	type args struct {
		shippingOptionId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *TaxRateService
		args  args
		want  []models.TaxRate
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ListByShippingOption(tt.args.shippingOptionId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxRateService.ListByShippingOption() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TaxRateService.ListByShippingOption() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

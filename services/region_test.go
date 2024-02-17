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

func TestNewRegionService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *RegionService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRegionService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRegionService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegionService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *RegionService
		args args
		want *RegionService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegionService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegionService_Create(t *testing.T) {
	type args struct {
		data *types.CreateRegionInput
	}
	tests := []struct {
		name  string
		s     *RegionService
		args  args
		want  *models.Region
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegionService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RegionService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRegionService_Update(t *testing.T) {
	type args struct {
		regionId uuid.UUID
		Update   *types.UpdateRegionInput
	}
	tests := []struct {
		name  string
		s     *RegionService
		args  args
		want  *models.Region
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.regionId, tt.args.Update)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegionService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RegionService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRegionService_validateFields(t *testing.T) {
	type args struct {
		data *types.UpdateRegionInput
		id   uuid.UUID
	}
	tests := []struct {
		name  string
		s     *RegionService
		args  args
		want  *models.Region
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.validateFields(tt.args.data, tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegionService.validateFields() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RegionService.validateFields() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRegionService_validateTaxRate(t *testing.T) {
	type args struct {
		taxRate float64
	}
	tests := []struct {
		name string
		s    *RegionService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.validateTaxRate(tt.args.taxRate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegionService.validateTaxRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegionService_validateCurrency(t *testing.T) {
	type args struct {
		currencyCode string
	}
	tests := []struct {
		name string
		s    *RegionService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.validateCurrency(tt.args.currencyCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegionService.validateCurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegionService_validateCountry(t *testing.T) {
	type args struct {
		code     string
		regionId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *RegionService
		args  args
		want  *models.Country
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.validateCountry(tt.args.code, tt.args.regionId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegionService.validateCountry() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RegionService.validateCountry() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRegionService_RetrieveByCountryCode(t *testing.T) {
	type args struct {
		code   string
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *RegionService
		args  args
		want  *models.Region
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveByCountryCode(tt.args.code, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegionService.RetrieveByCountryCode() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RegionService.RetrieveByCountryCode() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRegionService_RetrieveByName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name  string
		s     *RegionService
		args  args
		want  *models.Region
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveByName(tt.args.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegionService.RetrieveByName() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RegionService.RetrieveByName() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRegionService_Retrieve(t *testing.T) {
	type args struct {
		regionId uuid.UUID
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *RegionService
		args  args
		want  *models.Region
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.regionId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegionService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RegionService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRegionService_List(t *testing.T) {
	type args struct {
		selector *types.FilterableRegion
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *RegionService
		args  args
		want  []models.Region
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegionService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RegionService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRegionService_ListAndCount(t *testing.T) {
	type args struct {
		selector *types.FilterableRegion
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *RegionService
		args  args
		want  []models.Region
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegionService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RegionService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("RegionService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestRegionService_Delete(t *testing.T) {
	type args struct {
		regionId uuid.UUID
	}
	tests := []struct {
		name string
		s    *RegionService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.regionId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegionService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegionService_AddCountry(t *testing.T) {
	type args struct {
		regionId uuid.UUID
		code     string
	}
	tests := []struct {
		name  string
		s     *RegionService
		args  args
		want  *models.Region
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AddCountry(tt.args.regionId, tt.args.code)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegionService.AddCountry() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RegionService.AddCountry() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRegionService_RemoveCountry(t *testing.T) {
	type args struct {
		regionId uuid.UUID
		code     string
	}
	tests := []struct {
		name  string
		s     *RegionService
		args  args
		want  *models.Region
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RemoveCountry(tt.args.regionId, tt.args.code)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegionService.RemoveCountry() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RegionService.RemoveCountry() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRegionService_AddPaymentProvider(t *testing.T) {
	type args struct {
		regionId   uuid.UUID
		providerId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *RegionService
		args  args
		want  *models.Region
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AddPaymentProvider(tt.args.regionId, tt.args.providerId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegionService.AddPaymentProvider() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RegionService.AddPaymentProvider() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRegionService_RemovePaymentProvider(t *testing.T) {
	type args struct {
		regionId   uuid.UUID
		providerId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *RegionService
		args  args
		want  *models.Region
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RemovePaymentProvider(tt.args.regionId, tt.args.providerId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegionService.RemovePaymentProvider() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RegionService.RemovePaymentProvider() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRegionService_AddFulfillmentProvider(t *testing.T) {
	type args struct {
		regionId   uuid.UUID
		providerId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *RegionService
		args  args
		want  *models.Region
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AddFulfillmentProvider(tt.args.regionId, tt.args.providerId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegionService.AddFulfillmentProvider() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RegionService.AddFulfillmentProvider() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRegionService_RemoveFulfillmentProvider(t *testing.T) {
	type args struct {
		regionId   uuid.UUID
		providerId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *RegionService
		args  args
		want  *models.Region
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RemoveFulfillmentProvider(tt.args.regionId, tt.args.providerId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegionService.RemoveFulfillmentProvider() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RegionService.RemoveFulfillmentProvider() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

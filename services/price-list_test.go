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

func TestNewPriceListService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *PriceListService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPriceListService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPriceListService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPriceListService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *PriceListService
		args args
		want *PriceListService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PriceListService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPriceListService_Retrieve(t *testing.T) {
	type args struct {
		id     uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *PriceListService
		args  args
		want  *models.PriceList
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.id, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PriceListService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PriceListService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPriceListService_List(t *testing.T) {
	type args struct {
		selector *types.FilterablePriceList
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *PriceListService
		args  args
		want  []models.PriceList
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PriceListService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PriceListService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPriceListService_ListAndCount(t *testing.T) {
	type args struct {
		selector *types.FilterablePriceList
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *PriceListService
		args  args
		want  []models.PriceList
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PriceListService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PriceListService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("PriceListService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestPriceListService_ListPriceListsVariantIdsMap(t *testing.T) {
	type args struct {
		priceListIds uuid.UUIDs
	}
	tests := []struct {
		name  string
		s     *PriceListService
		args  args
		want  map[string][]string
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ListPriceListsVariantIdsMap(tt.args.priceListIds)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PriceListService.ListPriceListsVariantIdsMap() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PriceListService.ListPriceListsVariantIdsMap() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPriceListService_Create(t *testing.T) {
	type args struct {
		data *types.CreatePriceListInput
	}
	tests := []struct {
		name  string
		s     *PriceListService
		args  args
		want  *models.PriceList
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PriceListService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PriceListService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPriceListService_Update(t *testing.T) {
	type args struct {
		id   uuid.UUID
		data *types.UpdatePriceListInput
	}
	tests := []struct {
		name  string
		s     *PriceListService
		args  args
		want  *models.PriceList
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.id, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PriceListService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PriceListService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPriceListService_Delete(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name string
		s    *PriceListService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PriceListService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPriceListService_AddPrices(t *testing.T) {
	type args struct {
		id      uuid.UUID
		prices  []types.PriceListPriceCreateInput
		replace bool
	}
	tests := []struct {
		name  string
		s     *PriceListService
		args  args
		want  *models.PriceList
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AddPrices(tt.args.id, tt.args.prices, tt.args.replace)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PriceListService.AddPrices() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PriceListService.AddPrices() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPriceListService_DeletePrices(t *testing.T) {
	type args struct {
		id       uuid.UUID
		priceIds uuid.UUIDs
	}
	tests := []struct {
		name string
		s    *PriceListService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.DeletePrices(tt.args.id, tt.args.priceIds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PriceListService.DeletePrices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPriceListService_ClearPrices(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name string
		s    *PriceListService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ClearPrices(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PriceListService.ClearPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPriceListService_UpsertCustomerGroups(t *testing.T) {
	type args struct {
		id             uuid.UUID
		customerGroups []types.CustomerGroups
	}
	tests := []struct {
		name string
		s    *PriceListService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UpsertCustomerGroups(tt.args.id, tt.args.customerGroups); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PriceListService.UpsertCustomerGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPriceListService_ListProducts(t *testing.T) {
	type args struct {
		id                uuid.UUID
		selector          *types.FilterableProduct
		config            *sql.Options
		requiresPriceList bool
	}
	tests := []struct {
		name  string
		s     *PriceListService
		args  args
		want  []models.Product
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListProducts(tt.args.id, tt.args.selector, tt.args.config, tt.args.requiresPriceList)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PriceListService.ListProducts() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PriceListService.ListProducts() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("PriceListService.ListProducts() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestPriceListService_ListVariants(t *testing.T) {
	type args struct {
		id                uuid.UUID
		selector          *types.FilterableProductVariant
		config            *sql.Options
		requiresPriceList bool
	}
	tests := []struct {
		name  string
		s     *PriceListService
		args  args
		want  []models.ProductVariant
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListVariants(tt.args.id, tt.args.selector, tt.args.config, tt.args.requiresPriceList)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PriceListService.ListVariants() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PriceListService.ListVariants() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("PriceListService.ListVariants() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestPriceListService_DeleteProductPrices(t *testing.T) {
	type args struct {
		id         uuid.UUID
		productIds uuid.UUIDs
	}
	tests := []struct {
		name  string
		s     *PriceListService
		args  args
		want  uuid.UUIDs
		want1 *int
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.DeleteProductPrices(tt.args.id, tt.args.productIds)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PriceListService.DeleteProductPrices() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PriceListService.DeleteProductPrices() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("PriceListService.DeleteProductPrices() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestPriceListService_DeleteVariantPrices(t *testing.T) {
	type args struct {
		id         uuid.UUID
		variantIds uuid.UUIDs
	}
	tests := []struct {
		name  string
		s     *PriceListService
		args  args
		want  uuid.UUIDs
		want1 *int
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.DeleteVariantPrices(tt.args.id, tt.args.variantIds)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PriceListService.DeleteVariantPrices() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PriceListService.DeleteVariantPrices() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("PriceListService.DeleteVariantPrices() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestPriceListService_AddCurrencyFromRegion(t *testing.T) {
	type args struct {
		prices []types.PriceListPriceCreateInput
	}
	tests := []struct {
		name  string
		s     *PriceListService
		args  args
		want  []models.MoneyAmount
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AddCurrencyFromRegion(tt.args.prices)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PriceListService.AddCurrencyFromRegion() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PriceListService.AddCurrencyFromRegion() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

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

func TestNewProductVariantService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *ProductVariantService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProductVariantService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductVariantService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductVariantService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *ProductVariantService
		args args
		want *ProductVariantService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductVariantService_Retrieve(t *testing.T) {
	type args struct {
		id     uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *ProductVariantService
		args  args
		want  *models.ProductVariant
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.id, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductVariantService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductVariantService_RetrieveBySKU(t *testing.T) {
	type args struct {
		sku    string
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *ProductVariantService
		args  args
		want  *models.ProductVariant
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveBySKU(tt.args.sku, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantService.RetrieveBySKU() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductVariantService.RetrieveBySKU() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductVariantService_Create(t *testing.T) {
	type args struct {
		productId uuid.UUID
		product   *models.Product
		variants  []types.CreateProductVariantInput
	}
	tests := []struct {
		name  string
		s     *ProductVariantService
		args  args
		want  []models.ProductVariant
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.productId, tt.args.product, tt.args.variants)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductVariantService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductVariantService_UpdateBatch(t *testing.T) {
	type args struct {
		data []types.UpdateProductVariantData
	}
	tests := []struct {
		name  string
		s     *ProductVariantService
		args  args
		want  []models.ProductVariant
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.UpdateBatch(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantService.UpdateBatch() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductVariantService.UpdateBatch() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductVariantService_Update(t *testing.T) {
	type args struct {
		id      uuid.UUID
		variant *models.ProductVariant
		data    *types.UpdateProductVariantInput
	}
	tests := []struct {
		name  string
		s     *ProductVariantService
		args  args
		want  *models.ProductVariant
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.id, tt.args.variant, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductVariantService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductVariantService_UpdateVariantPrices(t *testing.T) {
	type args struct {
		data []types.UpdateVariantPricesData
	}
	tests := []struct {
		name string
		s    *ProductVariantService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UpdateVariantPrices(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantService.UpdateVariantPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductVariantService_UpsertRegionPrices(t *testing.T) {
	type args struct {
		data []types.UpdateVariantRegionPriceData
	}
	tests := []struct {
		name string
		s    *ProductVariantService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UpsertRegionPrices(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantService.UpsertRegionPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductVariantService_UpsertCurrencyPrices(t *testing.T) {
	type args struct {
		data []types.UpdateVariantCurrencyPriceData
	}
	tests := []struct {
		name string
		s    *ProductVariantService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UpsertCurrencyPrices(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantService.UpsertCurrencyPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductVariantService_GetRegionPrice(t *testing.T) {
	type args struct {
		id   uuid.UUID
		data *types.GetRegionPriceContext
	}
	tests := []struct {
		name  string
		s     *ProductVariantService
		args  args
		want  *float64
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetRegionPrice(tt.args.id, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantService.GetRegionPrice() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductVariantService.GetRegionPrice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductVariantService_SetRegionPrice(t *testing.T) {
	type args struct {
		id    uuid.UUID
		price types.ProductVariantPrice
	}
	tests := []struct {
		name  string
		s     *ProductVariantService
		args  args
		want  *models.MoneyAmount
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.SetRegionPrice(tt.args.id, tt.args.price)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantService.SetRegionPrice() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductVariantService.SetRegionPrice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductVariantService_SetCurrencyPrice(t *testing.T) {
	type args struct {
		id   uuid.UUID
		data types.ProductVariantPrice
	}
	tests := []struct {
		name  string
		s     *ProductVariantService
		args  args
		want  *models.MoneyAmount
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.SetCurrencyPrice(tt.args.id, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantService.SetCurrencyPrice() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductVariantService.SetCurrencyPrice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductVariantService_UpdateOptionValue(t *testing.T) {
	type args struct {
		id          uuid.UUID
		optionId    uuid.UUID
		optionValue string
	}
	tests := []struct {
		name  string
		s     *ProductVariantService
		args  args
		want  *models.ProductOptionValue
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.UpdateOptionValue(tt.args.id, tt.args.optionId, tt.args.optionValue)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantService.UpdateOptionValue() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductVariantService.UpdateOptionValue() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductVariantService_AddOptionValue(t *testing.T) {
	type args struct {
		id          uuid.UUID
		optionId    uuid.UUID
		optionValue string
	}
	tests := []struct {
		name  string
		s     *ProductVariantService
		args  args
		want  *models.ProductOptionValue
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AddOptionValue(tt.args.id, tt.args.optionId, tt.args.optionValue)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantService.AddOptionValue() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductVariantService.AddOptionValue() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductVariantService_DeleteOptionValue(t *testing.T) {
	type args struct {
		id       uuid.UUID
		optionId uuid.UUID
	}
	tests := []struct {
		name string
		s    *ProductVariantService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.DeleteOptionValue(tt.args.id, tt.args.optionId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantService.DeleteOptionValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductVariantService_ListAndCount(t *testing.T) {
	type args struct {
		selector *types.FilterableProductVariant
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *ProductVariantService
		args  args
		want  []models.ProductVariant
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductVariantService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("ProductVariantService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestProductVariantService_List(t *testing.T) {
	type args struct {
		selector *types.FilterableProductVariant
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *ProductVariantService
		args  args
		want  []models.ProductVariant
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductVariantService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductVariantService_Delete(t *testing.T) {
	type args struct {
		variantIds uuid.UUIDs
	}
	tests := []struct {
		name string
		s    *ProductVariantService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.variantIds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductVariantService_IsVariantInSalesChannels(t *testing.T) {
	type args struct {
		id              uuid.UUID
		salesChannelIds uuid.UUIDs
	}
	tests := []struct {
		name  string
		s     *ProductVariantService
		args  args
		want  bool
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.IsVariantInSalesChannels(tt.args.id, tt.args.salesChannelIds)
			if got != tt.want {
				t.Errorf("ProductVariantService.IsVariantInSalesChannels() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductVariantService.IsVariantInSalesChannels() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductVariantService_ValidateVariantsToCreate(t *testing.T) {
	type args struct {
		product  *models.Product
		variants []types.CreateProductVariantInput
	}
	tests := []struct {
		name string
		s    *ProductVariantService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ValidateVariantsToCreate(tt.args.product, tt.args.variants); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantService.ValidateVariantsToCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}

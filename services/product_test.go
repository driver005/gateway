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

func TestNewProductService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *ProductService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProductService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *ProductService
		args args
		want *ProductService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductService_List(t *testing.T) {
	type args struct {
		selector *types.FilterableProduct
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *ProductService
		args  args
		want  []models.Product
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductService_ListAndCount(t *testing.T) {
	type args struct {
		selector *types.FilterableProduct
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *ProductService
		args  args
		want  []models.Product
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("ProductService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestProductService_Count(t *testing.T) {
	type args struct {
		selector models.Product
	}
	tests := []struct {
		name  string
		s     *ProductService
		args  args
		want  *int64
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Count(tt.args.selector)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.Count() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductService.Count() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductService_RetrieveById(t *testing.T) {
	type args struct {
		id     uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *ProductService
		args  args
		want  *models.Product
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveById(tt.args.id, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.RetrieveById() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductService.RetrieveById() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductService_RetrieveByHandle(t *testing.T) {
	type args struct {
		productHandle string
		config        *sql.Options
	}
	tests := []struct {
		name  string
		s     *ProductService
		args  args
		want  *models.Product
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveByHandle(tt.args.productHandle, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.RetrieveByHandle() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductService.RetrieveByHandle() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductService_RetrieveByExternalId(t *testing.T) {
	type args struct {
		externalId string
		config     *sql.Options
	}
	tests := []struct {
		name  string
		s     *ProductService
		args  args
		want  *models.Product
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveByExternalId(tt.args.externalId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.RetrieveByExternalId() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductService.RetrieveByExternalId() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductService_Retrieve(t *testing.T) {
	type args struct {
		selector models.Product
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *ProductService
		args  args
		want  *models.Product
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductService_RetrieveVariants(t *testing.T) {
	type args struct {
		id     uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *ProductService
		args  args
		want  []models.ProductVariant
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveVariants(tt.args.id, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.RetrieveVariants() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductService.RetrieveVariants() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductService_FilterProductsBySalesChannel(t *testing.T) {
	type args struct {
		productIds     uuid.UUIDs
		salesChannelId uuid.UUID
		config         *sql.Options
	}
	tests := []struct {
		name  string
		s     *ProductService
		args  args
		want  []models.Product
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.FilterProductsBySalesChannel(tt.args.productIds, tt.args.salesChannelId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.FilterProductsBySalesChannel() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductService.FilterProductsBySalesChannel() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductService_ListTypes(t *testing.T) {
	tests := []struct {
		name  string
		s     *ProductService
		want  []models.ProductType
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ListTypes()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.ListTypes() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductService.ListTypes() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductService_ListTagsByUsage(t *testing.T) {
	type args struct {
		take int
	}
	tests := []struct {
		name  string
		s     *ProductService
		args  args
		want  []models.ProductTag
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ListTagsByUsage(tt.args.take)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.ListTagsByUsage() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductService.ListTagsByUsage() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductService_IsProductInSalesChannels(t *testing.T) {
	type args struct {
		id              uuid.UUID
		salesChannelIds uuid.UUIDs
	}
	tests := []struct {
		name  string
		s     *ProductService
		args  args
		want  bool
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.IsProductInSalesChannels(tt.args.id, tt.args.salesChannelIds)
			if got != tt.want {
				t.Errorf("ProductService.IsProductInSalesChannels() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductService.IsProductInSalesChannels() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductService_Create(t *testing.T) {
	type args struct {
		data *types.CreateProductInput
	}
	tests := []struct {
		name  string
		s     *ProductService
		args  args
		want  *models.Product
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductService_Update(t *testing.T) {
	type args struct {
		id   uuid.UUID
		data *types.UpdateProductInput
	}
	tests := []struct {
		name  string
		s     *ProductService
		args  args
		want  *models.Product
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.id, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductService_Delete(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name string
		s    *ProductService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductService_AddOption(t *testing.T) {
	type args struct {
		id          uuid.UUID
		optionTitle string
	}
	tests := []struct {
		name  string
		s     *ProductService
		args  args
		want  *models.Product
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AddOption(tt.args.id, tt.args.optionTitle)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.AddOption() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductService.AddOption() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductService_ReorderVariants(t *testing.T) {
	type args struct {
		id           uuid.UUID
		variantOrder uuid.UUIDs
	}
	tests := []struct {
		name  string
		s     *ProductService
		args  args
		want  *models.Product
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ReorderVariants(tt.args.id, tt.args.variantOrder)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.ReorderVariants() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductService.ReorderVariants() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductService_UpdateOption(t *testing.T) {
	type args struct {
		id       uuid.UUID
		optionId uuid.UUID
		data     *types.ProductOptionInput
	}
	tests := []struct {
		name  string
		s     *ProductService
		args  args
		want  *models.Product
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.UpdateOption(tt.args.id, tt.args.optionId, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.UpdateOption() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductService.UpdateOption() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductService_RetrieveOptionByTitle(t *testing.T) {
	type args struct {
		title string
		id    uuid.UUID
	}
	tests := []struct {
		name  string
		s     *ProductService
		args  args
		want  *models.ProductOption
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveOptionByTitle(tt.args.title, tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.RetrieveOptionByTitle() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductService.RetrieveOptionByTitle() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductService_DeleteOption(t *testing.T) {
	type args struct {
		id       uuid.UUID
		optionId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *ProductService
		args  args
		want  *models.Product
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.DeleteOption(tt.args.id, tt.args.optionId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.DeleteOption() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductService.DeleteOption() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductService_UpdateShippingProfile(t *testing.T) {
	type args struct {
		productIds uuid.UUIDs
		profileId  uuid.UUID
	}
	tests := []struct {
		name  string
		s     *ProductService
		args  args
		want  []models.Product
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.UpdateShippingProfile(tt.args.productIds, tt.args.profileId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.UpdateShippingProfile() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductService.UpdateShippingProfile() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductService_DecorateProductsWithSalesChannels(t *testing.T) {
	type args struct {
		products []models.Product
	}
	tests := []struct {
		name  string
		s     *ProductService
		args  args
		want  []models.Product
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.DecorateProductsWithSalesChannels(tt.args.products)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.DecorateProductsWithSalesChannels() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductService.DecorateProductsWithSalesChannels() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductService_getSalesChannelModuleChannels(t *testing.T) {
	type args struct {
		products []models.Product
	}
	tests := []struct {
		name  string
		s     *ProductService
		args  args
		want  map[uuid.UUID][]models.SalesChannel
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.getSalesChannelModuleChannels(tt.args.products)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.getSalesChannelModuleChannels() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductService.getSalesChannelModuleChannels() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

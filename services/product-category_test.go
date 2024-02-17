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

func TestNewProductCategoryService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *ProductCategoryService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProductCategoryService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductCategoryService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductCategoryService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *ProductCategoryService
		args args
		want *ProductCategoryService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductCategoryService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductCategoryService_List(t *testing.T) {
	type args struct {
		selector *types.FilterableProductCategory
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *ProductCategoryService
		args  args
		want  []models.ProductCategory
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductCategoryService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductCategoryService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductCategoryService_ListAndCount(t *testing.T) {
	type args struct {
		selector *types.FilterableProductCategory
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *ProductCategoryService
		args  args
		want  []models.ProductCategory
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductCategoryService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductCategoryService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("ProductCategoryService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestProductCategoryService_Retrieve(t *testing.T) {
	type args struct {
		selector models.ProductCategory
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *ProductCategoryService
		args  args
		want  *models.ProductCategory
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductCategoryService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductCategoryService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductCategoryService_RetrieveById(t *testing.T) {
	type args struct {
		id     uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *ProductCategoryService
		args  args
		want  *models.ProductCategory
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveById(tt.args.id, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductCategoryService.RetrieveById() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductCategoryService.RetrieveById() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductCategoryService_RetrieveByHandle(t *testing.T) {
	type args struct {
		handle string
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *ProductCategoryService
		args  args
		want  *models.ProductCategory
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveByHandle(tt.args.handle, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductCategoryService.RetrieveByHandle() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductCategoryService.RetrieveByHandle() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductCategoryService_Create(t *testing.T) {
	type args struct {
		data *types.CreateProductCategoryInput
	}
	tests := []struct {
		name  string
		s     *ProductCategoryService
		args  args
		want  *models.ProductCategory
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductCategoryService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductCategoryService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductCategoryService_Update(t *testing.T) {
	type args struct {
		productCategoryId uuid.UUID
		data              *types.UpdateProductCategoryInput
	}
	tests := []struct {
		name  string
		s     *ProductCategoryService
		args  args
		want  *models.ProductCategory
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.productCategoryId, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductCategoryService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductCategoryService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductCategoryService_Delete(t *testing.T) {
	type args struct {
		productCategoryId uuid.UUID
	}
	tests := []struct {
		name string
		s    *ProductCategoryService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.productCategoryId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductCategoryService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductCategoryService_AddProducts(t *testing.T) {
	type args struct {
		productCategoryId uuid.UUID
		productIDs        uuid.UUIDs
	}
	tests := []struct {
		name string
		s    *ProductCategoryService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.AddProducts(tt.args.productCategoryId, tt.args.productIDs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductCategoryService.AddProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductCategoryService_RemoveProducts(t *testing.T) {
	type args struct {
		productCategoryId uuid.UUID
		productIDs        uuid.UUIDs
	}
	tests := []struct {
		name string
		s    *ProductCategoryService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.RemoveProducts(tt.args.productCategoryId, tt.args.productIDs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductCategoryService.RemoveProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductCategoryService_fetchReorderConditions(t *testing.T) {
	type args struct {
		productCategory     *models.ProductCategory
		input               *types.UpdateProductCategoryInput
		shouldDeleteElement bool
	}
	tests := []struct {
		name string
		s    *ProductCategoryService
		args args
		want types.ReorderConditions
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.fetchReorderConditions(tt.args.productCategory, tt.args.input, tt.args.shouldDeleteElement); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductCategoryService.fetchReorderConditions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductCategoryService_performReordering(t *testing.T) {
	type args struct {
		conditions types.ReorderConditions
	}
	tests := []struct {
		name string
		s    *ProductCategoryService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.performReordering(tt.args.conditions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductCategoryService.performReordering() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductCategoryService_shiftSiblings(t *testing.T) {
	type args struct {
		conditions types.ReorderConditions
	}
	tests := []struct {
		name string
		s    *ProductCategoryService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.shiftSiblings(tt.args.conditions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductCategoryService.shiftSiblings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductCategoryService_transformParentIdToEntity(t *testing.T) {
	type args struct {
		update *types.UpdateProductCategoryInput
	}
	tests := []struct {
		name string
		s    *ProductCategoryService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.transformParentIdToEntity(tt.args.update); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductCategoryService.transformParentIdToEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}

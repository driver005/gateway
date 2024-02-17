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

func TestNewProductCollectionService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *ProductCollectionService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProductCollectionService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductCollectionService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductCollectionService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *ProductCollectionService
		args args
		want *ProductCollectionService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductCollectionService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductCollectionService_Retrieve(t *testing.T) {
	type args struct {
		collectionId uuid.UUID
		config       *sql.Options
	}
	tests := []struct {
		name  string
		s     *ProductCollectionService
		args  args
		want  *models.ProductCollection
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.collectionId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductCollectionService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductCollectionService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductCollectionService_RetrieveByHandle(t *testing.T) {
	type args struct {
		collectionHandle string
		config           *sql.Options
	}
	tests := []struct {
		name  string
		s     *ProductCollectionService
		args  args
		want  *models.ProductCollection
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveByHandle(tt.args.collectionHandle, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductCollectionService.RetrieveByHandle() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductCollectionService.RetrieveByHandle() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductCollectionService_List(t *testing.T) {
	type args struct {
		selector *types.FilterableCollection
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *ProductCollectionService
		args  args
		want  []models.ProductCollection
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductCollectionService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductCollectionService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductCollectionService_ListAndCount(t *testing.T) {
	type args struct {
		selector *types.FilterableCollection
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *ProductCollectionService
		args  args
		want  []models.ProductCollection
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductCollectionService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductCollectionService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("ProductCollectionService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestProductCollectionService_Create(t *testing.T) {
	type args struct {
		data *types.CreateProductCollection
	}
	tests := []struct {
		name  string
		s     *ProductCollectionService
		args  args
		want  *models.ProductCollection
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductCollectionService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductCollectionService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductCollectionService_Update(t *testing.T) {
	type args struct {
		collectionId uuid.UUID
		data         *types.UpdateProductCollection
	}
	tests := []struct {
		name  string
		s     *ProductCollectionService
		args  args
		want  *models.ProductCollection
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.collectionId, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductCollectionService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductCollectionService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductCollectionService_Delete(t *testing.T) {
	type args struct {
		collectionId uuid.UUID
	}
	tests := []struct {
		name string
		s    *ProductCollectionService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.collectionId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductCollectionService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductCollectionService_AddProducts(t *testing.T) {
	type args struct {
		collectionId uuid.UUID
		productIds   uuid.UUIDs
	}
	tests := []struct {
		name  string
		s     *ProductCollectionService
		args  args
		want  *models.ProductCollection
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AddProducts(tt.args.collectionId, tt.args.productIds)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductCollectionService.AddProducts() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductCollectionService.AddProducts() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductCollectionService_RemoveProducts(t *testing.T) {
	type args struct {
		collectionId uuid.UUID
		productIds   uuid.UUIDs
	}
	tests := []struct {
		name string
		s    *ProductCollectionService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.RemoveProducts(tt.args.collectionId, tt.args.productIds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductCollectionService.RemoveProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func TestNewProductTagService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *ProductTagService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProductTagService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductTagService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductTagService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *ProductTagService
		args args
		want *ProductTagService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductTagService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductTagService_Retrieve(t *testing.T) {
	type args struct {
		id     uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *ProductTagService
		args  args
		want  *models.ProductTag
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.id, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductTagService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductTagService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductTagService_Create(t *testing.T) {
	type args struct {
		tag *models.ProductTag
	}
	tests := []struct {
		name  string
		s     *ProductTagService
		args  args
		want  *models.ProductTag
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.tag)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductTagService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductTagService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductTagService_List(t *testing.T) {
	type args struct {
		selector *types.FilterableProductTag
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *ProductTagService
		args  args
		want  []models.ProductTag
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductTagService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductTagService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductTagService_ListAndCount(t *testing.T) {
	type args struct {
		selector *types.FilterableProductTag
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *ProductTagService
		args  args
		want  []models.ProductTag
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductTagService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductTagService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("ProductTagService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

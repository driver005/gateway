package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
)

func TestNewProductTaxRateService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *ProductTaxRateService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProductTaxRateService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductTaxRateService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductTaxRateService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *ProductTaxRateService
		args args
		want *ProductTaxRateService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductTaxRateService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductTaxRateService_List(t *testing.T) {
	type args struct {
		selector types.FilterableProductTaxRate
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *ProductTaxRateService
		args  args
		want  []models.ProductTaxRate
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductTaxRateService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductTaxRateService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

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

func TestNewCustomShippingOptionService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *CustomShippingOptionService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomShippingOptionService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomShippingOptionService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomShippingOptionService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *CustomShippingOptionService
		args args
		want *CustomShippingOptionService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomShippingOptionService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomShippingOptionService_Retrieve(t *testing.T) {
	type args struct {
		id     uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *CustomShippingOptionService
		args  args
		want  *models.CustomShippingOption
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.id, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomShippingOptionService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomShippingOptionService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCustomShippingOptionService_List(t *testing.T) {
	type args struct {
		selector models.CustomShippingOption
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *CustomShippingOptionService
		args  args
		want  []models.CustomShippingOption
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomShippingOptionService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomShippingOptionService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCustomShippingOptionService_Create(t *testing.T) {
	type args struct {
		data []types.CreateCustomShippingOptionInput
	}
	tests := []struct {
		name  string
		s     *CustomShippingOptionService
		args  args
		want  []models.CustomShippingOption
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomShippingOptionService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomShippingOptionService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

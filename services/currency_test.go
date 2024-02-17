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

func TestNewCurrencyService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *CurrencyService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCurrencyService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCurrencyService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCurrencyService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *CurrencyService
		args args
		want *CurrencyService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CurrencyService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCurrencyService_RetrieveByCode(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name  string
		s     *CurrencyService
		args  args
		want  *models.Currency
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveByCode(tt.args.code)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CurrencyService.RetrieveByCode() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CurrencyService.RetrieveByCode() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCurrencyService_ListAndCount(t *testing.T) {
	type args struct {
		selector *types.FilterableCurrencyProps
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *CurrencyService
		args  args
		want  []models.Currency
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CurrencyService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CurrencyService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("CurrencyService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestCurrencyService_Update(t *testing.T) {
	type args struct {
		code string
		data *types.UpdateCurrencyInput
	}
	tests := []struct {
		name  string
		s     *CurrencyService
		args  args
		want  *models.Currency
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.code, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CurrencyService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CurrencyService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

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

func TestNewStoreService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *StoreService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStoreService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStoreService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStoreService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *StoreService
		args args
		want *StoreService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StoreService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStoreService_Create(t *testing.T) {
	tests := []struct {
		name  string
		s     *StoreService
		want  *models.Store
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StoreService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("StoreService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStoreService_Retrieve(t *testing.T) {
	type args struct {
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *StoreService
		args  args
		want  *models.Store
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StoreService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("StoreService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStoreService_GetDefaultCurrency(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		s    *StoreService
		args args
		want *models.Currency
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetDefaultCurrency(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StoreService.GetDefaultCurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStoreService_Update(t *testing.T) {
	type args struct {
		data *types.UpdateStoreInput
	}
	tests := []struct {
		name  string
		s     *StoreService
		args  args
		want  *models.Store
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StoreService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("StoreService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStoreService_AddCurrency(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name  string
		s     *StoreService
		args  args
		want  *models.Store
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AddCurrency(tt.args.code)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StoreService.AddCurrency() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("StoreService.AddCurrency() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStoreService_RemoveCurrency(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name  string
		s     *StoreService
		args  args
		want  *models.Store
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RemoveCurrency(tt.args.code)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StoreService.RemoveCurrency() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("StoreService.RemoveCurrency() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

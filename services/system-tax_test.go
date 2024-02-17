package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
)

func TestNewSystemTaxService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *SystemTaxService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSystemTaxService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSystemTaxService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemTaxService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *SystemTaxService
		args args
		want *SystemTaxService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SystemTaxService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemProviderService_GetTaxLines(t *testing.T) {
	type args struct {
		itemLines     []interfaces.ItemTaxCalculationLine
		shippingLines []interfaces.ShippingTaxCalculationLine
		context       interfaces.TaxCalculationContext
	}
	tests := []struct {
		name  string
		s     *SystemProviderService
		args  args
		want  []types.ProviderTaxLine
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetTaxLines(tt.args.itemLines, tt.args.shippingLines, tt.args.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SystemProviderService.GetTaxLines() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SystemProviderService.GetTaxLines() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

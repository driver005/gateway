package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

func TestNewNewTotalsServices(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *NewTotalsService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNewTotalsServices(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNewTotalsServices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTotalsService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *NewTotalsService
		args args
		want *NewTotalsService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTotalsService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTotalsService_GetLineItemTotals(t *testing.T) {
	type args struct {
		items              []models.LineItem
		includeTax         bool
		calculationContext *interfaces.TaxCalculationContext
		taxRate            *float64
	}
	tests := []struct {
		name  string
		s     *NewTotalsService
		args  args
		want  map[uuid.UUID]models.LineItem
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetLineItemTotals(tt.args.items, tt.args.includeTax, tt.args.calculationContext, tt.args.taxRate)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTotalsService.GetLineItemTotals() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewTotalsService.GetLineItemTotals() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewTotalsService_getLineItemTotals(t *testing.T) {
	type args struct {
		item               models.LineItem
		taxRate            *float64
		includeTax         bool
		lineItemAllocation types.LineAllocations
		taxLines           []models.LineItemTaxLine
		calculationContext *interfaces.TaxCalculationContext
	}
	tests := []struct {
		name  string
		s     *NewTotalsService
		args  args
		want  *models.LineItem
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.getLineItemTotals(tt.args.item, tt.args.taxRate, tt.args.includeTax, tt.args.lineItemAllocation, tt.args.taxLines, tt.args.calculationContext)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTotalsService.getLineItemTotals() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewTotalsService.getLineItemTotals() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewTotalsService_getLineItemTotalsLegacy(t *testing.T) {
	type args struct {
		item               models.LineItem
		taxRate            *float64
		includeTax         bool
		lineItemAllocation types.LineAllocations
		taxLines           []models.LineItemTaxLine
		calculationContext *interfaces.TaxCalculationContext
	}
	tests := []struct {
		name  string
		s     *NewTotalsService
		args  args
		want  *models.LineItem
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.getLineItemTotalsLegacy(tt.args.item, tt.args.taxRate, tt.args.includeTax, tt.args.lineItemAllocation, tt.args.taxLines, tt.args.calculationContext)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTotalsService.getLineItemTotalsLegacy() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewTotalsService.getLineItemTotalsLegacy() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewTotalsService_GetLineItemRefund(t *testing.T) {
	type args struct {
		lineItem           models.LineItem
		calculationContext *interfaces.TaxCalculationContext
		taxRate            *float64
	}
	tests := []struct {
		name  string
		s     *NewTotalsService
		args  args
		want  float64
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetLineItemRefund(tt.args.lineItem, tt.args.calculationContext, tt.args.taxRate)
			if got != tt.want {
				t.Errorf("NewTotalsService.GetLineItemRefund() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewTotalsService.GetLineItemRefund() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewTotalsService_getLineItemRefundLegacy(t *testing.T) {
	type args struct {
		lineItem           models.LineItem
		calculationContext *interfaces.TaxCalculationContext
		taxRate            *float64
	}
	tests := []struct {
		name  string
		s     *NewTotalsService
		args  args
		want  float64
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.getLineItemRefundLegacy(tt.args.lineItem, tt.args.calculationContext, tt.args.taxRate)
			if got != tt.want {
				t.Errorf("NewTotalsService.getLineItemRefundLegacy() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewTotalsService.getLineItemRefundLegacy() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewTotalsService_GetGiftCardTotals(t *testing.T) {
	type args struct {
		giftCardableAmount   float64
		giftCardTransactions []models.GiftCardTransaction
		region               *models.Region
		giftCards            []models.GiftCard
	}
	tests := []struct {
		name  string
		s     *NewTotalsService
		args  args
		want  *Total
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetGiftCardTotals(tt.args.giftCardableAmount, tt.args.giftCardTransactions, tt.args.region, tt.args.giftCards)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTotalsService.GetGiftCardTotals() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewTotalsService.GetGiftCardTotals() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewTotalsService_GetGiftCardTransactionsTotals(t *testing.T) {
	type args struct {
		giftCardTransactions []models.GiftCardTransaction
		region               *models.Region
	}
	tests := []struct {
		name string
		s    *NewTotalsService
		args args
		want *Total
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetGiftCardTransactionsTotals(tt.args.giftCardTransactions, tt.args.region); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTotalsService.GetGiftCardTransactionsTotals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTotalsService_GetShippingMethodTotals(t *testing.T) {
	type args struct {
		shippingMethods    []models.ShippingMethod
		includeTax         bool
		discounts          []models.Discount
		taxRate            *float64
		calculationContext *interfaces.TaxCalculationContext
	}
	tests := []struct {
		name  string
		s     *NewTotalsService
		args  args
		want  map[uuid.UUID]models.ShippingMethod
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetShippingMethodTotals(tt.args.shippingMethods, tt.args.includeTax, tt.args.discounts, tt.args.taxRate, tt.args.calculationContext)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTotalsService.GetShippingMethodTotals() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewTotalsService.GetShippingMethodTotals() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewTotalsService_GetGiftCardableAmount(t *testing.T) {
	type args struct {
		giftCardsTaxable bool
		subtotal         float64
		shippingTotal    float64
		discountTotal    float64
		taxTotal         float64
	}
	tests := []struct {
		name string
		s    *NewTotalsService
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetGiftCardableAmount(tt.args.giftCardsTaxable, tt.args.subtotal, tt.args.shippingTotal, tt.args.discountTotal, tt.args.taxTotal); got != tt.want {
				t.Errorf("NewTotalsService.GetGiftCardableAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTotalsService_getShippingMethodTotals(t *testing.T) {
	type args struct {
		shippingMethod     models.ShippingMethod
		includeTax         bool
		calculationContext *interfaces.TaxCalculationContext
		taxLines           []models.ShippingMethodTaxLine
		discounts          []models.Discount
		taxRate            float64
	}
	tests := []struct {
		name  string
		s     *NewTotalsService
		args  args
		want  *models.ShippingMethod
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.getShippingMethodTotals(tt.args.shippingMethod, tt.args.includeTax, tt.args.calculationContext, tt.args.taxLines, tt.args.discounts, tt.args.taxRate)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTotalsService.getShippingMethodTotals() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewTotalsService.getShippingMethodTotals() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewTotalsService_getShippingMethodTotalsLegacy(t *testing.T) {
	type args struct {
		shippingMethod     models.ShippingMethod
		includeTax         bool
		calculationContext *interfaces.TaxCalculationContext
		taxLines           []models.ShippingMethodTaxLine
		discounts          []models.Discount
		taxRate            float64
	}
	tests := []struct {
		name  string
		s     *NewTotalsService
		args  args
		want  *models.ShippingMethod
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.getShippingMethodTotalsLegacy(tt.args.shippingMethod, tt.args.includeTax, tt.args.calculationContext, tt.args.taxLines, tt.args.discounts, tt.args.taxRate)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTotalsService.getShippingMethodTotalsLegacy() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewTotalsService.getShippingMethodTotalsLegacy() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

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

func TestNewTotalService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *TotalsService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTotalService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTotalService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTotalsService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *TotalsService
		args args
		want *TotalsService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TotalsService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTotalsService_GetTotal(t *testing.T) {
	type args struct {
		cart    *models.Cart
		order   *models.Order
		options GetTotalsOptions
	}
	tests := []struct {
		name  string
		s     *TotalsService
		args  args
		want  float64
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetTotal(tt.args.cart, tt.args.order, tt.args.options)
			if got != tt.want {
				t.Errorf("TotalsService.GetTotal() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TotalsService.GetTotal() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTotalsService_GetPaidTotal(t *testing.T) {
	type args struct {
		order *models.Order
	}
	tests := []struct {
		name string
		s    *TotalsService
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetPaidTotal(tt.args.order); got != tt.want {
				t.Errorf("TotalsService.GetPaidTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTotalsService_GetSwapTotal(t *testing.T) {
	type args struct {
		order *models.Order
	}
	tests := []struct {
		name string
		s    *TotalsService
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetSwapTotal(tt.args.order); got != tt.want {
				t.Errorf("TotalsService.GetSwapTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTotalsService_GetShippingMethodTotals(t *testing.T) {
	type args struct {
		shippingMethod models.ShippingMethod
		cart           *models.Cart
		order          *models.Order
		opts           GetShippingMethodTotalsOptions
	}
	tests := []struct {
		name  string
		s     *TotalsService
		args  args
		want  *ShippingMethodTotals
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetShippingMethodTotals(tt.args.shippingMethod, tt.args.cart, tt.args.order, tt.args.opts)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TotalsService.GetShippingMethodTotals() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TotalsService.GetShippingMethodTotals() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTotalsService_GetSubtotal(t *testing.T) {
	type args struct {
		cart  *models.Cart
		order *models.Order
		opts  types.SubtotalOptions
	}
	tests := []struct {
		name  string
		s     *TotalsService
		args  args
		want  float64
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetSubtotal(tt.args.cart, tt.args.order, tt.args.opts)
			if got != tt.want {
				t.Errorf("TotalsService.GetSubtotal() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TotalsService.GetSubtotal() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTotalsService_GetShippingTotal(t *testing.T) {
	type args struct {
		cart  *models.Cart
		order *models.Order
	}
	tests := []struct {
		name  string
		s     *TotalsService
		args  args
		want  float64
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetShippingTotal(tt.args.cart, tt.args.order)
			if got != tt.want {
				t.Errorf("TotalsService.GetShippingTotal() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TotalsService.GetShippingTotal() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTotalsService_GetTaxTotal(t *testing.T) {
	type args struct {
		cart       *models.Cart
		order      *models.Order
		forceTaxes bool
	}
	tests := []struct {
		name  string
		s     *TotalsService
		args  args
		want  float64
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetTaxTotal(tt.args.cart, tt.args.order, tt.args.forceTaxes)
			if got != tt.want {
				t.Errorf("TotalsService.GetTaxTotal() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TotalsService.GetTaxTotal() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTotalsService_GetAllocationMap(t *testing.T) {
	type args struct {
		calculationContextData types.CalculationContextData
		options                AllocationMapOptions
	}
	tests := []struct {
		name  string
		s     *TotalsService
		args  args
		want  types.LineAllocationsMap
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetAllocationMap(tt.args.calculationContextData, tt.args.options)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TotalsService.GetAllocationMap() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TotalsService.GetAllocationMap() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTotalsService_GetRefundedTotal(t *testing.T) {
	type args struct {
		order *models.Order
	}
	tests := []struct {
		name string
		s    *TotalsService
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetRefundedTotal(tt.args.order); got != tt.want {
				t.Errorf("TotalsService.GetRefundedTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTotalsService_GetLineItemRefund(t *testing.T) {
	type args struct {
		order    *models.Order
		lineItem models.LineItem
	}
	tests := []struct {
		name  string
		s     *TotalsService
		args  args
		want  float64
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetLineItemRefund(tt.args.order, tt.args.lineItem)
			if got != tt.want {
				t.Errorf("TotalsService.GetLineItemRefund() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TotalsService.GetLineItemRefund() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTotalsService_GetRefundTotal(t *testing.T) {
	type args struct {
		order     *models.Order
		lineItems []models.LineItem
	}
	tests := []struct {
		name  string
		s     *TotalsService
		args  args
		want  float64
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetRefundTotal(tt.args.order, tt.args.lineItems)
			if got != tt.want {
				t.Errorf("TotalsService.GetRefundTotal() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TotalsService.GetRefundTotal() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTotalsService_CalculateDiscount(t *testing.T) {
	type args struct {
		lineItem     models.LineItem
		variant      uuid.UUID
		variantPrice float64
		value        float64
		discountType models.DiscountRuleType
	}
	tests := []struct {
		name string
		s    *TotalsService
		args args
		want types.LineDiscount
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.CalculateDiscount(tt.args.lineItem, tt.args.variant, tt.args.variantPrice, tt.args.value, tt.args.discountType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TotalsService.CalculateDiscount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTotalsService_GetAllocationItemDiscounts(t *testing.T) {
	type args struct {
		discount models.Discount
		cart     models.Cart
	}
	tests := []struct {
		name string
		s    *TotalsService
		args args
		want []types.LineDiscount
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetAllocationItemDiscounts(tt.args.discount, tt.args.cart); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TotalsService.GetAllocationItemDiscounts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTotalsService_GetLineItemDiscountAdjustment(t *testing.T) {
	type args struct {
		lineItem models.LineItem
		discount models.Discount
	}
	tests := []struct {
		name string
		s    *TotalsService
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetLineItemDiscountAdjustment(tt.args.lineItem, tt.args.discount); got != tt.want {
				t.Errorf("TotalsService.GetLineItemDiscountAdjustment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTotalsService_GetLineItemAdjustmentsTotal(t *testing.T) {
	type args struct {
		cart  *models.Cart
		order *models.Order
	}
	tests := []struct {
		name string
		s    *TotalsService
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetLineItemAdjustmentsTotal(tt.args.cart, tt.args.order); got != tt.want {
				t.Errorf("TotalsService.GetLineItemAdjustmentsTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTotalsService_GetLineDiscounts(t *testing.T) {
	type args struct {
		calculationContextData types.CalculationContextData
		discount               models.Discount
	}
	tests := []struct {
		name string
		s    *TotalsService
		args args
		want []types.LineDiscountAmount
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetLineDiscounts(tt.args.calculationContextData, tt.args.discount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TotalsService.GetLineDiscounts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTotalsService_GetLineItemTotals(t *testing.T) {
	type args struct {
		lineItem models.LineItem
		cart     *models.Cart
		order    *models.Order
		options  *LineItemTotalsOptions
	}
	tests := []struct {
		name  string
		s     *TotalsService
		args  args
		want  *models.LineItem
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetLineItemTotals(tt.args.lineItem, tt.args.cart, tt.args.order, tt.args.options)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TotalsService.GetLineItemTotals() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TotalsService.GetLineItemTotals() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTotalsService_GetLineItemTotal(t *testing.T) {
	type args struct {
		lineItem models.LineItem
		cart     *models.Cart
		order    *models.Order
		options  GetLineItemTotalOptions
	}
	tests := []struct {
		name  string
		s     *TotalsService
		args  args
		want  float64
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetLineItemTotal(tt.args.lineItem, tt.args.cart, tt.args.order, tt.args.options)
			if got != tt.want {
				t.Errorf("TotalsService.GetLineItemTotal() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TotalsService.GetLineItemTotal() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTotalsService_GetGiftCardableAmount(t *testing.T) {
	type args struct {
		cart  *models.Cart
		order *models.Order
	}
	tests := []struct {
		name  string
		s     *TotalsService
		args  args
		want  float64
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetGiftCardableAmount(tt.args.cart, tt.args.order)
			if got != tt.want {
				t.Errorf("TotalsService.GetGiftCardableAmount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TotalsService.GetGiftCardableAmount() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTotalsService_GetGiftCardTotal(t *testing.T) {
	type args struct {
		cart  *models.Cart
		order *models.Order
		opts  map[string]interface{}
	}
	tests := []struct {
		name  string
		s     *TotalsService
		args  args
		want  *Total
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetGiftCardTotal(tt.args.cart, tt.args.order, tt.args.opts)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TotalsService.GetGiftCardTotal() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TotalsService.GetGiftCardTotal() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTotalsService_GetDiscountTotal(t *testing.T) {
	type args struct {
		cart  *models.Cart
		order *models.Order
	}
	tests := []struct {
		name  string
		s     *TotalsService
		args  args
		want  float64
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetDiscountTotal(tt.args.cart, tt.args.order)
			if got != tt.want {
				t.Errorf("TotalsService.GetDiscountTotal() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TotalsService.GetDiscountTotal() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTotalsService_GetCalculationContext(t *testing.T) {
	type args struct {
		calculationContextData types.CalculationContextData
		options                CalculationContextOptions
	}
	tests := []struct {
		name  string
		s     *TotalsService
		args  args
		want  *interfaces.TaxCalculationContext
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetCalculationContext(tt.args.calculationContextData, tt.args.options)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TotalsService.GetCalculationContext() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TotalsService.GetCalculationContext() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTotalsService_Rounded(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		s    *TotalsService
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Rounded(tt.args.value); got != tt.want {
				t.Errorf("TotalsService.Rounded() = %v, want %v", got, tt.want)
			}
		})
	}
}

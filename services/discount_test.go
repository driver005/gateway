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

func TestNewDiscountService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *DiscountService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDiscountService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDiscountService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiscountService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *DiscountService
		args args
		want *DiscountService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscountService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiscountService_ValidateDiscountRule(t *testing.T) {
	type args struct {
		discountRule *types.CreateDiscountRuleInput
	}
	tests := []struct {
		name  string
		s     *DiscountService
		args  args
		want  *types.CreateDiscountRuleInput
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ValidateDiscountRule(tt.args.discountRule)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscountService.ValidateDiscountRule() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DiscountService.ValidateDiscountRule() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDiscountService_List(t *testing.T) {
	type args struct {
		selector types.FilterableDiscount
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *DiscountService
		args  args
		want  []models.Discount
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscountService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DiscountService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDiscountService_ListAndCount(t *testing.T) {
	type args struct {
		selector *types.FilterableDiscount
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *DiscountService
		args  args
		want  []models.Discount
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscountService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DiscountService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("DiscountService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestDiscountService_Create(t *testing.T) {
	type args struct {
		data *types.CreateDiscountInput
	}
	tests := []struct {
		name  string
		s     *DiscountService
		args  args
		want  *models.Discount
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscountService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DiscountService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDiscountService_Retrieve(t *testing.T) {
	type args struct {
		discountId uuid.UUID
		config     *sql.Options
	}
	tests := []struct {
		name  string
		s     *DiscountService
		args  args
		want  *models.Discount
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.discountId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscountService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DiscountService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDiscountService_RetrieveByCode(t *testing.T) {
	type args struct {
		discountCode string
		config       *sql.Options
	}
	tests := []struct {
		name  string
		s     *DiscountService
		args  args
		want  *models.Discount
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveByCode(tt.args.discountCode, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscountService.RetrieveByCode() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DiscountService.RetrieveByCode() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDiscountService_ListByCodes(t *testing.T) {
	type args struct {
		discountCodes []string
		config        sql.Query
	}
	tests := []struct {
		name  string
		s     *DiscountService
		args  args
		want  []models.Discount
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ListByCodes(tt.args.discountCodes, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscountService.ListByCodes() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DiscountService.ListByCodes() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDiscountService_Update(t *testing.T) {
	type args struct {
		discountId uuid.UUID
		data       *types.UpdateDiscountInput
	}
	tests := []struct {
		name  string
		s     *DiscountService
		args  args
		want  *models.Discount
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.discountId, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscountService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DiscountService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDiscountService_CreateDynamicCode(t *testing.T) {
	type args struct {
		discountId uuid.UUID
		data       *types.CreateDynamicDiscountInput
	}
	tests := []struct {
		name  string
		s     *DiscountService
		args  args
		want  *models.Discount
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateDynamicCode(tt.args.discountId, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscountService.CreateDynamicCode() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DiscountService.CreateDynamicCode() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDiscountService_DeleteDynamicCode(t *testing.T) {
	type args struct {
		discountId uuid.UUID
		code       string
	}
	tests := []struct {
		name string
		s    *DiscountService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.DeleteDynamicCode(tt.args.discountId, tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscountService.DeleteDynamicCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiscountService_AddRegion(t *testing.T) {
	type args struct {
		discountId uuid.UUID
		regionId   uuid.UUID
	}
	tests := []struct {
		name  string
		s     *DiscountService
		args  args
		want  *models.Discount
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AddRegion(tt.args.discountId, tt.args.regionId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscountService.AddRegion() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DiscountService.AddRegion() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDiscountService_RemoveRegion(t *testing.T) {
	type args struct {
		discountId uuid.UUID
		regionId   uuid.UUID
	}
	tests := []struct {
		name  string
		s     *DiscountService
		args  args
		want  *models.Discount
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RemoveRegion(tt.args.discountId, tt.args.regionId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscountService.RemoveRegion() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DiscountService.RemoveRegion() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDiscountService_Delete(t *testing.T) {
	type args struct {
		discountId uuid.UUID
	}
	tests := []struct {
		name string
		s    *DiscountService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.discountId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscountService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiscountService_ValidateDiscountForProduct(t *testing.T) {
	type args struct {
		discountRuleId uuid.UUID
		productId      uuid.UUID
	}
	tests := []struct {
		name  string
		s     *DiscountService
		args  args
		want  bool
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ValidateDiscountForProduct(tt.args.discountRuleId, tt.args.productId)
			if got != tt.want {
				t.Errorf("DiscountService.ValidateDiscountForProduct() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DiscountService.ValidateDiscountForProduct() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDiscountService_CalculateDiscountForLineItem(t *testing.T) {
	type args struct {
		discountId             uuid.UUID
		lineItem               *models.LineItem
		calculationContextData types.CalculationContextData
	}
	tests := []struct {
		name  string
		s     *DiscountService
		args  args
		want  float64
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CalculateDiscountForLineItem(tt.args.discountId, tt.args.lineItem, tt.args.calculationContextData)
			if got != tt.want {
				t.Errorf("DiscountService.CalculateDiscountForLineItem() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DiscountService.CalculateDiscountForLineItem() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDiscountService_validateDiscountForCartOrThrow(t *testing.T) {
	type args struct {
		cart      *models.Cart
		discounts []models.Discount
	}
	tests := []struct {
		name string
		s    *DiscountService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.validateDiscountForCartOrThrow(tt.args.cart, tt.args.discounts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscountService.validateDiscountForCartOrThrow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiscountService_hasCustomersGroupCondition(t *testing.T) {
	type args struct {
		discount models.Discount
	}
	tests := []struct {
		name string
		s    *DiscountService
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.hasCustomersGroupCondition(tt.args.discount); got != tt.want {
				t.Errorf("DiscountService.hasCustomersGroupCondition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiscountService_hasReachedLimit(t *testing.T) {
	type args struct {
		discount models.Discount
	}
	tests := []struct {
		name string
		s    *DiscountService
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.hasReachedLimit(tt.args.discount); got != tt.want {
				t.Errorf("DiscountService.hasReachedLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiscountService_hasNotStarted(t *testing.T) {
	type args struct {
		discount models.Discount
	}
	tests := []struct {
		name string
		s    *DiscountService
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.hasNotStarted(tt.args.discount); got != tt.want {
				t.Errorf("DiscountService.hasNotStarted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiscountService_hasExpired(t *testing.T) {
	type args struct {
		discount models.Discount
	}
	tests := []struct {
		name string
		s    *DiscountService
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.hasExpired(tt.args.discount); got != tt.want {
				t.Errorf("DiscountService.hasExpired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiscountService_isDisabled(t *testing.T) {
	type args struct {
		discount models.Discount
	}
	tests := []struct {
		name string
		s    *DiscountService
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.isDisabled(tt.args.discount); got != tt.want {
				t.Errorf("DiscountService.isDisabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiscountService_isValidForRegion(t *testing.T) {
	type args struct {
		discount  models.Discount
		region_id uuid.UUID
	}
	tests := []struct {
		name  string
		s     *DiscountService
		args  args
		want  bool
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.isValidForRegion(tt.args.discount, tt.args.region_id)
			if got != tt.want {
				t.Errorf("DiscountService.isValidForRegion() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DiscountService.isValidForRegion() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDiscountService_canApplyForCustomer(t *testing.T) {
	type args struct {
		discountRuleId uuid.UUID
		customerId     uuid.UUID
	}
	tests := []struct {
		name  string
		s     *DiscountService
		args  args
		want  bool
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.canApplyForCustomer(tt.args.discountRuleId, tt.args.customerId)
			if got != tt.want {
				t.Errorf("DiscountService.canApplyForCustomer() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DiscountService.canApplyForCustomer() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

func TestNewShippingOptionService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *ShippingOptionService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewShippingOptionService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewShippingOptionService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShippingOptionService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *ShippingOptionService
		args args
		want *ShippingOptionService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingOptionService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShippingOptionService_ValidateRequirement(t *testing.T) {
	type args struct {
		data     *types.ValidateRequirementTypeInput
		optionId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *ShippingOptionService
		args  args
		want  *models.ShippingOptionRequirement
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ValidateRequirement(tt.args.data, tt.args.optionId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingOptionService.ValidateRequirement() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingOptionService.ValidateRequirement() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingOptionService_List(t *testing.T) {
	type args struct {
		selector *types.FilterableShippingOption
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *ShippingOptionService
		args  args
		want  []models.ShippingOption
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingOptionService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingOptionService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingOptionService_ListAndCount(t *testing.T) {
	type args struct {
		selector *types.FilterableShippingOption
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *ShippingOptionService
		args  args
		want  []models.ShippingOption
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingOptionService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingOptionService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("ShippingOptionService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestShippingOptionService_Retrieve(t *testing.T) {
	type args struct {
		optionId uuid.UUID
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *ShippingOptionService
		args  args
		want  *models.ShippingOption
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.optionId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingOptionService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingOptionService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingOptionService_UpdateShippingMethod(t *testing.T) {
	type args struct {
		id   uuid.UUID
		data *types.ShippingMethodUpdate
	}
	tests := []struct {
		name  string
		s     *ShippingOptionService
		args  args
		want  *models.ShippingMethod
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.UpdateShippingMethod(tt.args.id, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingOptionService.UpdateShippingMethod() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingOptionService.UpdateShippingMethod() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingOptionService_DeleteShippingMethods(t *testing.T) {
	type args struct {
		shippingMethods []models.ShippingMethod
	}
	tests := []struct {
		name string
		s    *ShippingOptionService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.DeleteShippingMethods(tt.args.shippingMethods); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingOptionService.DeleteShippingMethods() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShippingOptionService_CreateShippingMethod(t *testing.T) {
	type args struct {
		optionId uuid.UUID
		data     map[string]interface{}
		config   *types.CreateShippingMethodDto
	}
	tests := []struct {
		name  string
		s     *ShippingOptionService
		args  args
		want  *models.ShippingMethod
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateShippingMethod(tt.args.optionId, tt.args.data, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingOptionService.CreateShippingMethod() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingOptionService.CreateShippingMethod() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingOptionService_ValidateCartOption(t *testing.T) {
	type args struct {
		option *models.ShippingOption
		cart   *models.Cart
	}
	tests := []struct {
		name  string
		s     *ShippingOptionService
		args  args
		want  *models.ShippingOption
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ValidateCartOption(tt.args.option, tt.args.cart)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingOptionService.ValidateCartOption() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingOptionService.ValidateCartOption() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingOptionService_ValidateAndMutatePrice(t *testing.T) {
	type args struct {
		option     *models.ShippingOption
		option2    *types.CreateShippingOptionInput
		priceInput types.ValidatePriceTypeAndAmountInput
	}
	tests := []struct {
		name  string
		s     *ShippingOptionService
		args  args
		want  *models.ShippingOption
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ValidateAndMutatePrice(tt.args.option, tt.args.option2, tt.args.priceInput)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingOptionService.ValidateAndMutatePrice() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingOptionService.ValidateAndMutatePrice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingOptionService_validatePriceType(t *testing.T) {
	type args struct {
		priceType models.ShippingOptionPriceType
		option    *models.ShippingOption
	}
	tests := []struct {
		name  string
		s     *ShippingOptionService
		args  args
		want  models.ShippingOptionPriceType
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.validatePriceType(tt.args.priceType, tt.args.option)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingOptionService.validatePriceType() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingOptionService.validatePriceType() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingOptionService_Create(t *testing.T) {
	type args struct {
		data *types.CreateShippingOptionInput
	}
	tests := []struct {
		name  string
		s     *ShippingOptionService
		args  args
		want  *models.ShippingOption
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingOptionService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingOptionService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingOptionService_Update(t *testing.T) {
	type args struct {
		optionId uuid.UUID
		data     *types.UpdateShippingOptionInput
	}
	tests := []struct {
		name  string
		s     *ShippingOptionService
		args  args
		want  *models.ShippingOption
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.optionId, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingOptionService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingOptionService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingOptionService_Delete(t *testing.T) {
	type args struct {
		optionId uuid.UUID
	}
	tests := []struct {
		name string
		s    *ShippingOptionService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.optionId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingOptionService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShippingOptionService_AddRequirement(t *testing.T) {
	type args struct {
		optionId    uuid.UUID
		requirement *models.ShippingOptionRequirement
	}
	tests := []struct {
		name  string
		s     *ShippingOptionService
		args  args
		want  *models.ShippingOption
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AddRequirement(tt.args.optionId, tt.args.requirement)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingOptionService.AddRequirement() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingOptionService.AddRequirement() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingOptionService_RemoveRequirement(t *testing.T) {
	type args struct {
		requirementId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *ShippingOptionService
		args  args
		want  *models.ShippingOptionRequirement
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RemoveRequirement(tt.args.requirementId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingOptionService.RemoveRequirement() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingOptionService.RemoveRequirement() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingOptionService_UpdateShippingProfile(t *testing.T) {
	type args struct {
		optionIds uuid.UUIDs
		profileId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *ShippingOptionService
		args  args
		want  []models.ShippingOption
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.UpdateShippingProfile(tt.args.optionIds, tt.args.profileId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingOptionService.UpdateShippingProfile() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingOptionService.UpdateShippingProfile() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingOptionService_GetPrice(t *testing.T) {
	type args struct {
		option *models.ShippingOption
		data   core.JSONB
		cart   *models.Cart
	}
	tests := []struct {
		name  string
		s     *ShippingOptionService
		args  args
		want  float64
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetPrice(tt.args.option, tt.args.data, tt.args.cart)
			if got != tt.want {
				t.Errorf("ShippingOptionService.GetPrice() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingOptionService.GetPrice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

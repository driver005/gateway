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

func TestNewShippingProfileService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *ShippingProfileService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewShippingProfileService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewShippingProfileService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShippingProfileService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *ShippingProfileService
		args args
		want *ShippingProfileService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingProfileService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShippingProfileService_List(t *testing.T) {
	type args struct {
		selector *types.FilterableShippingProfile
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *ShippingProfileService
		args  args
		want  []models.ShippingProfile
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingProfileService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingProfileService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingProfileService_GetMapProfileIdsByProductIds(t *testing.T) {
	type args struct {
		productIds uuid.UUIDs
	}
	tests := []struct {
		name  string
		s     *ShippingProfileService
		args  args
		want  *map[uuid.UUID]uuid.UUID
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetMapProfileIdsByProductIds(tt.args.productIds)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingProfileService.GetMapProfileIdsByProductIds() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingProfileService.GetMapProfileIdsByProductIds() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingProfileService_Retrieve(t *testing.T) {
	type args struct {
		profileId uuid.UUID
		config    *sql.Options
	}
	tests := []struct {
		name  string
		s     *ShippingProfileService
		args  args
		want  *models.ShippingProfile
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.profileId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingProfileService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingProfileService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingProfileService_RetrieveForProducts(t *testing.T) {
	type args struct {
		productIds uuid.UUIDs
	}
	tests := []struct {
		name  string
		s     *ShippingProfileService
		args  args
		want  map[string]models.ShippingProfile
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveForProducts(tt.args.productIds)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingProfileService.RetrieveForProducts() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingProfileService.RetrieveForProducts() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingProfileService_RetrieveDefault(t *testing.T) {
	tests := []struct {
		name  string
		s     *ShippingProfileService
		want  *models.ShippingProfile
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveDefault()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingProfileService.RetrieveDefault() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingProfileService.RetrieveDefault() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingProfileService_CreateDefault(t *testing.T) {
	tests := []struct {
		name  string
		s     *ShippingProfileService
		want  *models.ShippingProfile
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateDefault()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingProfileService.CreateDefault() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingProfileService.CreateDefault() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingProfileService_RetrieveGiftCardDefault(t *testing.T) {
	tests := []struct {
		name  string
		s     *ShippingProfileService
		want  *models.ShippingProfile
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveGiftCardDefault()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingProfileService.RetrieveGiftCardDefault() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingProfileService.RetrieveGiftCardDefault() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingProfileService_CreateGiftCardDefault(t *testing.T) {
	tests := []struct {
		name  string
		s     *ShippingProfileService
		want  *models.ShippingProfile
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateGiftCardDefault()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingProfileService.CreateGiftCardDefault() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingProfileService.CreateGiftCardDefault() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingProfileService_Create(t *testing.T) {
	type args struct {
		data *types.CreateShippingProfile
	}
	tests := []struct {
		name  string
		s     *ShippingProfileService
		args  args
		want  *models.ShippingProfile
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingProfileService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingProfileService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingProfileService_Update(t *testing.T) {
	type args struct {
		profileId uuid.UUID
		data      *types.UpdateShippingProfile
	}
	tests := []struct {
		name  string
		s     *ShippingProfileService
		args  args
		want  *models.ShippingProfile
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.profileId, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingProfileService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingProfileService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingProfileService_Delete(t *testing.T) {
	type args struct {
		profileId uuid.UUID
	}
	tests := []struct {
		name string
		s    *ShippingProfileService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.profileId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingProfileService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShippingProfileService_AddProduct(t *testing.T) {
	type args struct {
		profileId uuid.UUID
		productId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *ShippingProfileService
		args  args
		want  *models.ShippingProfile
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AddProduct(tt.args.profileId, tt.args.productId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingProfileService.AddProduct() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingProfileService.AddProduct() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingProfileService_AddProducts(t *testing.T) {
	type args struct {
		profileId  uuid.UUID
		productIds uuid.UUIDs
	}
	tests := []struct {
		name  string
		s     *ShippingProfileService
		args  args
		want  *models.ShippingProfile
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AddProducts(tt.args.profileId, tt.args.productIds)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingProfileService.AddProducts() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingProfileService.AddProducts() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingProfileService_RemoveProducts(t *testing.T) {
	type args struct {
		profileId  uuid.UUID
		productIds uuid.UUIDs
	}
	tests := []struct {
		name  string
		s     *ShippingProfileService
		args  args
		want  *models.ShippingProfile
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RemoveProducts(tt.args.profileId, tt.args.productIds)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingProfileService.RemoveProducts() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingProfileService.RemoveProducts() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingProfileService_AddShippingOption(t *testing.T) {
	type args struct {
		profileId uuid.UUID
		optionId  uuid.UUID
	}
	tests := []struct {
		name  string
		s     *ShippingProfileService
		args  args
		want  *models.ShippingProfile
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AddShippingOption(tt.args.profileId, tt.args.optionId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingProfileService.AddShippingOption() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingProfileService.AddShippingOption() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingProfileService_AddShippingOptions(t *testing.T) {
	type args struct {
		profileId uuid.UUID
		optionIds uuid.UUIDs
	}
	tests := []struct {
		name  string
		s     *ShippingProfileService
		args  args
		want  *models.ShippingProfile
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AddShippingOptions(tt.args.profileId, tt.args.optionIds)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingProfileService.AddShippingOptions() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingProfileService.AddShippingOptions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingProfileService_FetchCartOptions(t *testing.T) {
	type args struct {
		cart *models.Cart
	}
	tests := []struct {
		name  string
		s     *ShippingProfileService
		args  args
		want  []models.ShippingOption
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.FetchCartOptions(tt.args.cart)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingProfileService.FetchCartOptions() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingProfileService.FetchCartOptions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestShippingProfileService_GetProfilesInCart(t *testing.T) {
	type args struct {
		cart *models.Cart
	}
	tests := []struct {
		name  string
		s     *ShippingProfileService
		args  args
		want  uuid.UUIDs
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetProfilesInCart(tt.args.cart)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingProfileService.GetProfilesInCart() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShippingProfileService.GetProfilesInCart() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/sarulabs/di"
)

func TestNewFulfillmentProviderService(t *testing.T) {
	type args struct {
		container di.Container
		r         Registry
	}
	tests := []struct {
		name string
		args args
		want *FulfillmentProviderService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFulfillmentProviderService(tt.args.container, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFulfillmentProviderService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFulfillmentProviderService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *FulfillmentProviderService
		args args
		want *FulfillmentProviderService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FulfillmentProviderService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFulfillmentProviderService_RetrieveProvider(t *testing.T) {
	type args struct {
		providerID uuid.UUID
	}
	tests := []struct {
		name string
		s    *FulfillmentProviderService
		args args
		want interfaces.IFulfillmentService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.RetrieveProvider(tt.args.providerID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FulfillmentProviderService.RetrieveProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFulfillmentProviderService_RegisterInstalledProviders(t *testing.T) {
	type args struct {
		providers uuid.UUIDs
	}
	tests := []struct {
		name string
		s    *FulfillmentProviderService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.RegisterInstalledProviders(tt.args.providers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FulfillmentProviderService.RegisterInstalledProviders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFulfillmentProviderService_List(t *testing.T) {
	tests := []struct {
		name  string
		s     *FulfillmentProviderService
		want  []models.FulfillmentProvider
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FulfillmentProviderService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FulfillmentProviderService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFulfillmentProviderService_ListFulfillmentOptions(t *testing.T) {
	type args struct {
		providerIDs uuid.UUIDs
	}
	tests := []struct {
		name  string
		s     *FulfillmentProviderService
		args  args
		want  []types.FulfillmentOptions
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ListFulfillmentOptions(tt.args.providerIDs)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FulfillmentProviderService.ListFulfillmentOptions() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FulfillmentProviderService.ListFulfillmentOptions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFulfillmentProviderService_CreateFulfillment(t *testing.T) {
	type args struct {
		method      *models.ShippingMethod
		items       []models.LineItem
		order       *types.CreateFulfillmentOrder
		fulfillment *models.Fulfillment
	}
	tests := []struct {
		name  string
		s     *FulfillmentProviderService
		args  args
		want  core.JSONB
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateFulfillment(tt.args.method, tt.args.items, tt.args.order, tt.args.fulfillment)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FulfillmentProviderService.CreateFulfillment() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FulfillmentProviderService.CreateFulfillment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFulfillmentProviderService_CanCalculate(t *testing.T) {
	type args struct {
		option types.FulfillmentOptions
	}
	tests := []struct {
		name  string
		s     *FulfillmentProviderService
		args  args
		want  bool
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CanCalculate(tt.args.option)
			if got != tt.want {
				t.Errorf("FulfillmentProviderService.CanCalculate() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FulfillmentProviderService.CanCalculate() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFulfillmentProviderService_ValidateFulfillmentData(t *testing.T) {
	type args struct {
		option *models.ShippingOption
		data   map[string]interface{}
		cart   *models.Cart
	}
	tests := []struct {
		name  string
		s     *FulfillmentProviderService
		args  args
		want  map[string]interface{}
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ValidateFulfillmentData(tt.args.option, tt.args.data, tt.args.cart)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FulfillmentProviderService.ValidateFulfillmentData() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FulfillmentProviderService.ValidateFulfillmentData() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFulfillmentProviderService_CancelFulfillment(t *testing.T) {
	type args struct {
		fulfillment *models.Fulfillment
	}
	tests := []struct {
		name  string
		s     *FulfillmentProviderService
		args  args
		want  *models.Fulfillment
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CancelFulfillment(tt.args.fulfillment)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FulfillmentProviderService.CancelFulfillment() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FulfillmentProviderService.CancelFulfillment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFulfillmentProviderService_CalculatePrice(t *testing.T) {
	type args struct {
		option *models.ShippingOption
		data   map[string]interface{}
		cart   *models.Cart
	}
	tests := []struct {
		name  string
		s     *FulfillmentProviderService
		args  args
		want  float64
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CalculatePrice(tt.args.option, tt.args.data, tt.args.cart)
			if got != tt.want {
				t.Errorf("FulfillmentProviderService.CalculatePrice() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FulfillmentProviderService.CalculatePrice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFulfillmentProviderService_ValidateOption(t *testing.T) {
	type args struct {
		option models.ShippingOption
	}
	tests := []struct {
		name  string
		s     *FulfillmentProviderService
		args  args
		want  bool
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ValidateOption(tt.args.option)
			if got != tt.want {
				t.Errorf("FulfillmentProviderService.ValidateOption() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FulfillmentProviderService.ValidateOption() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFulfillmentProviderService_CreateReturn(t *testing.T) {
	type args struct {
		returnOrder *models.Return
	}
	tests := []struct {
		name  string
		s     *FulfillmentProviderService
		args  args
		want  core.JSONB
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateReturn(tt.args.returnOrder)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FulfillmentProviderService.CreateReturn() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FulfillmentProviderService.CreateReturn() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFulfillmentProviderService_RetrieveDocuments(t *testing.T) {
	type args struct {
		providerId      uuid.UUID
		fulfillmentData map[string]interface{}
		documentType    string
	}
	tests := []struct {
		name  string
		s     *FulfillmentProviderService
		args  args
		want  map[string]interface{}
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveDocuments(tt.args.providerId, tt.args.fulfillmentData, tt.args.documentType)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FulfillmentProviderService.RetrieveDocuments() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FulfillmentProviderService.RetrieveDocuments() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

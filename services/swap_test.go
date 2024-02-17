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

func TestNewSwapService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *SwapService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSwapService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSwapService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwapService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *SwapService
		args args
		want *SwapService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwapService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwapService_transformQueryForCart(t *testing.T) {
	type args struct {
		config *sql.Options
	}
	tests := []struct {
		name string
		s    *SwapService
		args args
		want map[string]*sql.Options
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.transformQueryForCart(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwapService.transformQueryForCart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwapService_Retrieve(t *testing.T) {
	type args struct {
		swapId uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *SwapService
		args  args
		want  *models.Swap
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.swapId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwapService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SwapService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSwapService_RetrieveByCartId(t *testing.T) {
	type args struct {
		cartId    uuid.UUID
		relations []string
	}
	tests := []struct {
		name  string
		s     *SwapService
		args  args
		want  *models.Swap
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveByCartId(tt.args.cartId, tt.args.relations)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwapService.RetrieveByCartId() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SwapService.RetrieveByCartId() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSwapService_List(t *testing.T) {
	type args struct {
		selector *types.FilterableSwap
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *SwapService
		args  args
		want  []models.Swap
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwapService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SwapService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSwapService_ListAndCount(t *testing.T) {
	type args struct {
		selector *types.FilterableSwap
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *SwapService
		args  args
		want  []models.Swap
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwapService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SwapService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("SwapService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestSwapService_Create(t *testing.T) {
	type args struct {
		order           *models.Order
		returnItems     []types.OrderReturnItem
		additionalItems []types.CreateClaimItemAdditionalItemInput
		returnShipping  *types.CreateClaimReturnShippingInput
		custom          map[string]interface{}
	}
	tests := []struct {
		name  string
		s     *SwapService
		args  args
		want  *models.Swap
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.order, tt.args.returnItems, tt.args.additionalItems, tt.args.returnShipping, tt.args.custom)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwapService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SwapService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSwapService_ProcessDifference(t *testing.T) {
	type args struct {
		swapId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *SwapService
		args  args
		want  *models.Swap
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ProcessDifference(tt.args.swapId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwapService.ProcessDifference() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SwapService.ProcessDifference() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSwapService_Update(t *testing.T) {
	type args struct {
		swapId uuid.UUID
		Update *models.Swap
	}
	tests := []struct {
		name  string
		s     *SwapService
		args  args
		want  *models.Swap
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.swapId, tt.args.Update)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwapService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SwapService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSwapService_CreateCart(t *testing.T) {
	type args struct {
		swapId                uuid.UUID
		customShippingOptions []types.CreateCustomShippingOptionInput
		context               map[string]interface{}
	}
	tests := []struct {
		name  string
		s     *SwapService
		args  args
		want  *models.Swap
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateCart(tt.args.swapId, tt.args.customShippingOptions, tt.args.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwapService.CreateCart() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SwapService.CreateCart() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSwapService_RegisterCartCompletion(t *testing.T) {
	type args struct {
		swapId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *SwapService
		args  args
		want  *models.Swap
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RegisterCartCompletion(tt.args.swapId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwapService.RegisterCartCompletion() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SwapService.RegisterCartCompletion() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSwapService_Cancel(t *testing.T) {
	type args struct {
		swapId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *SwapService
		args  args
		want  *models.Swap
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Cancel(tt.args.swapId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwapService.Cancel() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SwapService.Cancel() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSwapService_CreateFulfillment(t *testing.T) {
	type args struct {
		swapId uuid.UUID
		config *types.CreateShipmentConfig
	}
	tests := []struct {
		name  string
		s     *SwapService
		args  args
		want  *models.Swap
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateFulfillment(tt.args.swapId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwapService.CreateFulfillment() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SwapService.CreateFulfillment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSwapService_CancelFulfillment(t *testing.T) {
	type args struct {
		fulfillmentId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *SwapService
		args  args
		want  *models.Swap
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CancelFulfillment(tt.args.fulfillmentId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwapService.CancelFulfillment() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SwapService.CancelFulfillment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSwapService_CreateShipment(t *testing.T) {
	type args struct {
		swapId        uuid.UUID
		fulfillmentId uuid.UUID
		trackingLinks []models.TrackingLink
		config        *types.CreateShipmentConfig
	}
	tests := []struct {
		name  string
		s     *SwapService
		args  args
		want  *models.Swap
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateShipment(tt.args.swapId, tt.args.fulfillmentId, tt.args.trackingLinks, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwapService.CreateShipment() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SwapService.CreateShipment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSwapService_DeleteMetadata(t *testing.T) {
	type args struct {
		swapId string
		key    string
	}
	tests := []struct {
		name  string
		s     *SwapService
		args  args
		want  *models.Swap
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.DeleteMetadata(tt.args.swapId, tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwapService.DeleteMetadata() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SwapService.DeleteMetadata() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSwapService_RegisterReceived(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name  string
		s     *SwapService
		args  args
		want  *models.Swap
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RegisterReceived(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwapService.RegisterReceived() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SwapService.RegisterReceived() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSwapService_AreReturnItemsValid(t *testing.T) {
	type args struct {
		returnItems []types.OrderReturnItem
	}
	tests := []struct {
		name  string
		s     *SwapService
		args  args
		want  bool
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AreReturnItemsValid(tt.args.returnItems)
			if got != tt.want {
				t.Errorf("SwapService.AreReturnItemsValid() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SwapService.AreReturnItemsValid() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

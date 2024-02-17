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

func TestNewFulfillmentService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *FulfillmentService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFulfillmentService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFulfillmentService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFulfillmentService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *FulfillmentService
		args args
		want *FulfillmentService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FulfillmentService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFulfillmentService_PartitionItems(t *testing.T) {
	type args struct {
		shippingMethods []models.ShippingMethod
		items           []models.LineItem
	}
	tests := []struct {
		name string
		s    *FulfillmentService
		args args
		want []types.FulfillmentItemPartition
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.PartitionItems(tt.args.shippingMethods, tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FulfillmentService.PartitionItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFulfillmentService_GetFulfillmentItems(t *testing.T) {
	type args struct {
		order *types.CreateFulfillmentOrder
		items []types.FulFillmentItemType
	}
	tests := []struct {
		name  string
		s     *FulfillmentService
		args  args
		want  []models.LineItem
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetFulfillmentItems(tt.args.order, tt.args.items)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FulfillmentService.GetFulfillmentItems() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FulfillmentService.GetFulfillmentItems() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFulfillmentService_ValidateFulfillmentLineItem(t *testing.T) {
	type args struct {
		item     *models.LineItem
		quantity int
	}
	tests := []struct {
		name  string
		s     *FulfillmentService
		args  args
		want  *models.LineItem
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ValidateFulfillmentLineItem(tt.args.item, tt.args.quantity)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FulfillmentService.ValidateFulfillmentLineItem() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FulfillmentService.ValidateFulfillmentLineItem() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFulfillmentService_Retrieve(t *testing.T) {
	type args struct {
		fulfillmentId uuid.UUID
		config        *sql.Options
	}
	tests := []struct {
		name  string
		s     *FulfillmentService
		args  args
		want  *models.Fulfillment
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.fulfillmentId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FulfillmentService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FulfillmentService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFulfillmentService_CreateFulfillment(t *testing.T) {
	type args struct {
		order          *types.CreateFulfillmentOrder
		itemsToFulfill []types.FulFillmentItemType
		custom         models.Fulfillment
	}
	tests := []struct {
		name  string
		s     *FulfillmentService
		args  args
		want  []models.Fulfillment
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateFulfillment(tt.args.order, tt.args.itemsToFulfill, tt.args.custom)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FulfillmentService.CreateFulfillment() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FulfillmentService.CreateFulfillment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFulfillmentService_CancelFulfillment(t *testing.T) {
	type args struct {
		fulfillmentId uuid.UUID
		fulfillment   *models.Fulfillment
	}
	tests := []struct {
		name  string
		s     *FulfillmentService
		args  args
		want  *models.Fulfillment
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CancelFulfillment(tt.args.fulfillmentId, tt.args.fulfillment)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FulfillmentService.CancelFulfillment() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FulfillmentService.CancelFulfillment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFulfillmentService_CreateShipment(t *testing.T) {
	type args struct {
		fulfillmentId uuid.UUID
		trackingLinks []models.TrackingLink
		config        *types.CreateShipmentConfig
	}
	tests := []struct {
		name  string
		s     *FulfillmentService
		args  args
		want  *models.Fulfillment
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateShipment(tt.args.fulfillmentId, tt.args.trackingLinks, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FulfillmentService.CreateShipment() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FulfillmentService.CreateShipment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

func TestNewSalesChannelInventoryService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *SalesChannelInventoryService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSalesChannelInventoryService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSalesChannelInventoryService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSalesChannelInventoryService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *SalesChannelInventoryService
		args args
		want *SalesChannelInventoryService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesChannelInventoryService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSalesChannelInventoryService_RetrieveAvailableItemQuantity(t *testing.T) {
	type args struct {
		salesChannelId  uuid.UUID
		inventoryItemId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *SalesChannelInventoryService
		args  args
		want  int
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveAvailableItemQuantity(tt.args.salesChannelId, tt.args.inventoryItemId)
			if got != tt.want {
				t.Errorf("SalesChannelInventoryService.RetrieveAvailableItemQuantity() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SalesChannelInventoryService.RetrieveAvailableItemQuantity() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

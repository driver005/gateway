package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

func TestNewSalesChannelLocationService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *SalesChannelLocationService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSalesChannelLocationService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSalesChannelLocationService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSalesChannelLocationService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *SalesChannelLocationService
		args args
		want *SalesChannelLocationService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesChannelLocationService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSalesChannelLocationService_RemoveLocation(t *testing.T) {
	type args struct {
		salesChannelId uuid.UUID
		locationId     uuid.UUID
	}
	tests := []struct {
		name string
		s    *SalesChannelLocationService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.RemoveLocation(tt.args.salesChannelId, tt.args.locationId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesChannelLocationService.RemoveLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSalesChannelLocationService_AssociateLocation(t *testing.T) {
	type args struct {
		salesChannelId uuid.UUID
		locationId     uuid.UUID
	}
	tests := []struct {
		name string
		s    *SalesChannelLocationService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.AssociateLocation(tt.args.salesChannelId, tt.args.locationId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesChannelLocationService.AssociateLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSalesChannelLocationService_ListLocationIds(t *testing.T) {
	type args struct {
		salesChannelId uuid.UUIDs
	}
	tests := []struct {
		name  string
		s     *SalesChannelLocationService
		args  args
		want  uuid.UUIDs
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ListLocationIds(tt.args.salesChannelId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesChannelLocationService.ListLocationIds() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SalesChannelLocationService.ListLocationIds() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSalesChannelLocationService_ListSalesChannelIds(t *testing.T) {
	type args struct {
		locationId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *SalesChannelLocationService
		args  args
		want  uuid.UUIDs
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ListSalesChannelIds(tt.args.locationId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesChannelLocationService.ListSalesChannelIds() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SalesChannelLocationService.ListSalesChannelIds() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

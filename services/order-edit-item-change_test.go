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

func TestNewOrderItemChangeService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *OrderItemChangeService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrderItemChangeService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrderItemChangeService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderItemChangeService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *OrderItemChangeService
		args args
		want *OrderItemChangeService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderItemChangeService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderItemChangeService_Retrieve(t *testing.T) {
	type args struct {
		id     uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *OrderItemChangeService
		args  args
		want  *models.OrderItemChange
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.id, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderItemChangeService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderItemChangeService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderItemChangeService_List(t *testing.T) {
	type args struct {
		selector models.OrderItemChange
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *OrderItemChangeService
		args  args
		want  []models.OrderItemChange
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderItemChangeService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderItemChangeService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderItemChangeService_Create(t *testing.T) {
	type args struct {
		data *types.CreateOrderEditItemChangeInput
	}
	tests := []struct {
		name  string
		s     *OrderItemChangeService
		args  args
		want  *models.OrderItemChange
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderItemChangeService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderItemChangeService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderItemChangeService_Delete(t *testing.T) {
	type args struct {
		itemChangeIds uuid.UUIDs
	}
	tests := []struct {
		name string
		s    *OrderItemChangeService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.itemChangeIds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderItemChangeService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

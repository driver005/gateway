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

func TestNewReturnService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *ReturnService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReturnService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReturnService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReturnService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *ReturnService
		args args
		want *ReturnService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReturnService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReturnService_GetFulfillmentItems(t *testing.T) {
	type args struct {
		order       *models.Order
		items       []types.OrderReturnItem
		transformer Transformer
	}
	tests := []struct {
		name  string
		s     *ReturnService
		args  args
		want  []models.ReturnItem
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetFulfillmentItems(tt.args.order, tt.args.items, tt.args.transformer)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReturnService.GetFulfillmentItems() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ReturnService.GetFulfillmentItems() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReturnService_List(t *testing.T) {
	type args struct {
		selector *types.FilterableReturn
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *ReturnService
		args  args
		want  []models.Return
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReturnService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ReturnService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReturnService_ListAndCount(t *testing.T) {
	type args struct {
		selector *types.FilterableReturn
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *ReturnService
		args  args
		want  []models.Return
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReturnService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ReturnService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("ReturnService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestReturnService_Cancel(t *testing.T) {
	type args struct {
		returnId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *ReturnService
		args  args
		want  *models.Return
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Cancel(tt.args.returnId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReturnService.Cancel() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ReturnService.Cancel() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReturnService_ValidateReturnStatuses(t *testing.T) {
	type args struct {
		order *models.Order
	}
	tests := []struct {
		name string
		s    *ReturnService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ValidateReturnStatuses(tt.args.order); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReturnService.ValidateReturnStatuses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReturnService_ValidateReturnLineItem(t *testing.T) {
	type args struct {
		item       *models.LineItem
		quantity   int
		additional *types.OrderReturnItem
	}
	tests := []struct {
		name  string
		s     *ReturnService
		args  args
		want  *models.ReturnItem
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ValidateReturnLineItem(tt.args.item, tt.args.quantity, tt.args.additional)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReturnService.ValidateReturnLineItem() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ReturnService.ValidateReturnLineItem() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReturnService_Retrieve(t *testing.T) {
	type args struct {
		returnId uuid.UUID
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *ReturnService
		args  args
		want  *models.Return
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.returnId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReturnService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ReturnService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReturnService_RetrieveBySwap(t *testing.T) {
	type args struct {
		swapId    uuid.UUID
		relations []string
	}
	tests := []struct {
		name  string
		s     *ReturnService
		args  args
		want  *models.Return
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveBySwap(tt.args.swapId, tt.args.relations)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReturnService.RetrieveBySwap() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ReturnService.RetrieveBySwap() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReturnService_Update(t *testing.T) {
	type args struct {
		returnId uuid.UUID
		data     *types.UpdateReturnInput
	}
	tests := []struct {
		name  string
		s     *ReturnService
		args  args
		want  *models.Return
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.returnId, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReturnService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ReturnService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReturnService_Create(t *testing.T) {
	type args struct {
		data *types.CreateReturnInput
	}
	tests := []struct {
		name  string
		s     *ReturnService
		args  args
		want  *models.Return
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReturnService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ReturnService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReturnService_Fulfill(t *testing.T) {
	type args struct {
		returnId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *ReturnService
		args  args
		want  *models.Return
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Fulfill(tt.args.returnId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReturnService.Fulfill() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ReturnService.Fulfill() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReturnService_Receive(t *testing.T) {
	type args struct {
		returnId      uuid.UUID
		receivedItems []types.OrderReturnItem
		refundAmount  *float64
		allowMismatch bool
		context       map[string]interface{}
	}
	tests := []struct {
		name  string
		s     *ReturnService
		args  args
		want  *models.Return
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Receive(tt.args.returnId, tt.args.receivedItems, tt.args.refundAmount, tt.args.allowMismatch, tt.args.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReturnService.Receive() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ReturnService.Receive() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

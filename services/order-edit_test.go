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

func TestNewOrderEditService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *OrderEditService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrderEditService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrderEditService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderEditService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *OrderEditService
		args args
		want *OrderEditService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderEditService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderEditService_Retrieve(t *testing.T) {
	type args struct {
		id     uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *OrderEditService
		args  args
		want  *models.OrderEdit
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.id, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderEditService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderEditService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderEditService_ListAndCount(t *testing.T) {
	type args struct {
		selector *types.FilterableOrderEdit
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *OrderEditService
		args  args
		want  []models.OrderEdit
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderEditService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderEditService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("OrderEditService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestOrderEditService_List(t *testing.T) {
	type args struct {
		selector *types.FilterableOrderEdit
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *OrderEditService
		args  args
		want  []models.OrderEdit
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderEditService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderEditService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderEditService_Create(t *testing.T) {
	type args struct {
		data      *types.CreateOrderEditInput
		createdBy uuid.UUID
	}
	tests := []struct {
		name  string
		s     *OrderEditService
		args  args
		want  *models.OrderEdit
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data, tt.args.createdBy)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderEditService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderEditService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderEditService_Update(t *testing.T) {
	type args struct {
		id   uuid.UUID
		data *models.OrderEdit
	}
	tests := []struct {
		name  string
		s     *OrderEditService
		args  args
		want  *models.OrderEdit
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.id, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderEditService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderEditService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderEditService_Delete(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name string
		s    *OrderEditService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderEditService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderEditService_Decline(t *testing.T) {
	type args struct {
		id             uuid.UUID
		declinedBy     uuid.UUID
		declinedReason string
	}
	tests := []struct {
		name  string
		s     *OrderEditService
		args  args
		want  *models.OrderEdit
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Decline(tt.args.id, tt.args.declinedBy, tt.args.declinedReason)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderEditService.Decline() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderEditService.Decline() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderEditService_UpdateLineItem(t *testing.T) {
	type args struct {
		id       uuid.UUID
		itemId   uuid.UUID
		quantity int
	}
	tests := []struct {
		name string
		s    *OrderEditService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UpdateLineItem(tt.args.id, tt.args.itemId, tt.args.quantity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderEditService.UpdateLineItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderEditService_RemoveLineItem(t *testing.T) {
	type args struct {
		id         uuid.UUID
		lineItemId uuid.UUID
	}
	tests := []struct {
		name string
		s    *OrderEditService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.RemoveLineItem(tt.args.id, tt.args.lineItemId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderEditService.RemoveLineItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderEditService_RefreshAdjustments(t *testing.T) {
	type args struct {
		id                        uuid.UUID
		preserveCustomAdjustments bool
	}
	tests := []struct {
		name string
		s    *OrderEditService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.RefreshAdjustments(tt.args.id, tt.args.preserveCustomAdjustments); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderEditService.RefreshAdjustments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderEditService_DecorateTotals(t *testing.T) {
	type args struct {
		orderEdit *models.OrderEdit
	}
	tests := []struct {
		name  string
		s     *OrderEditService
		args  args
		want  *models.OrderEdit
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.DecorateTotals(tt.args.orderEdit)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderEditService.DecorateTotals() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderEditService.DecorateTotals() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderEditService_AddLineItem(t *testing.T) {
	type args struct {
		id   uuid.UUID
		data *types.AddOrderEditLineItemInput
	}
	tests := []struct {
		name string
		s    *OrderEditService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.AddLineItem(tt.args.id, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderEditService.AddLineItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderEditService_DeleteItemChange(t *testing.T) {
	type args struct {
		id           uuid.UUID
		itemChangeId uuid.UUID
	}
	tests := []struct {
		name string
		s    *OrderEditService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.DeleteItemChange(tt.args.id, tt.args.itemChangeId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderEditService.DeleteItemChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderEditService_RequestConfirmation(t *testing.T) {
	type args struct {
		id          uuid.UUID
		requestedBy uuid.UUID
	}
	tests := []struct {
		name  string
		s     *OrderEditService
		args  args
		want  *models.OrderEdit
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RequestConfirmation(tt.args.id, tt.args.requestedBy)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderEditService.RequestConfirmation() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderEditService.RequestConfirmation() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderEditService_Cancel(t *testing.T) {
	type args struct {
		id         uuid.UUID
		canceledBy uuid.UUID
	}
	tests := []struct {
		name  string
		s     *OrderEditService
		args  args
		want  *models.OrderEdit
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Cancel(tt.args.id, tt.args.canceledBy)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderEditService.Cancel() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderEditService.Cancel() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderEditService_Confirm(t *testing.T) {
	type args struct {
		id          uuid.UUID
		confirmedBy uuid.UUID
	}
	tests := []struct {
		name  string
		s     *OrderEditService
		args  args
		want  *models.OrderEdit
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Confirm(tt.args.id, tt.args.confirmedBy)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderEditService.Confirm() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderEditService.Confirm() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderEditService_RetrieveActive(t *testing.T) {
	type args struct {
		orderId uuid.UUID
		config  *sql.Options
	}
	tests := []struct {
		name  string
		s     *OrderEditService
		args  args
		want  *models.OrderEdit
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveActive(tt.args.orderId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderEditService.RetrieveActive() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderEditService.RetrieveActive() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderEditService_DeleteClonedItems(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name string
		s    *OrderEditService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.DeleteClonedItems(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderEditService.DeleteClonedItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderEditService_isOrderEditActive(t *testing.T) {
	type args struct {
		orderEdit *models.OrderEdit
	}
	tests := []struct {
		name string
		s    *OrderEditService
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.isOrderEditActive(tt.args.orderEdit); got != tt.want {
				t.Errorf("OrderEditService.isOrderEditActive() = %v, want %v", got, tt.want)
			}
		})
	}
}

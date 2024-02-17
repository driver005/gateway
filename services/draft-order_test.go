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

func TestNewDraftOrderService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *DraftOrderService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDraftOrderService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDraftOrderService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDraftOrderService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *DraftOrderService
		args args
		want *DraftOrderService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DraftOrderService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDraftOrderService_Retrieve(t *testing.T) {
	type args struct {
		draftOrderId uuid.UUID
		config       *sql.Options
	}
	tests := []struct {
		name  string
		s     *DraftOrderService
		args  args
		want  *models.DraftOrder
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.draftOrderId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DraftOrderService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DraftOrderService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDraftOrderService_RetrieveByCartId(t *testing.T) {
	type args struct {
		cartId uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *DraftOrderService
		args  args
		want  *models.DraftOrder
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveByCartId(tt.args.cartId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DraftOrderService.RetrieveByCartId() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DraftOrderService.RetrieveByCartId() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDraftOrderService_Delete(t *testing.T) {
	type args struct {
		draftOrderId uuid.UUID
	}
	tests := []struct {
		name string
		s    *DraftOrderService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.draftOrderId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DraftOrderService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDraftOrderService_ListAndCount(t *testing.T) {
	type args struct {
		selector *types.FilterableDraftOrder
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *DraftOrderService
		args  args
		want  []models.DraftOrder
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DraftOrderService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DraftOrderService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("DraftOrderService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestDraftOrderService_List(t *testing.T) {
	type args struct {
		selector *models.DraftOrder
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *DraftOrderService
		args  args
		want  []models.DraftOrder
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DraftOrderService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DraftOrderService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDraftOrderService_Create(t *testing.T) {
	type args struct {
		data *types.DraftOrderCreate
	}
	tests := []struct {
		name  string
		s     *DraftOrderService
		args  args
		want  *models.DraftOrder
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DraftOrderService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DraftOrderService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDraftOrderService_RegisterCartCompletion(t *testing.T) {
	type args struct {
		id      uuid.UUID
		orderId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *DraftOrderService
		args  args
		want  *models.DraftOrder
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RegisterCartCompletion(tt.args.id, tt.args.orderId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DraftOrderService.RegisterCartCompletion() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DraftOrderService.RegisterCartCompletion() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDraftOrderService_Update(t *testing.T) {
	type args struct {
		id   uuid.UUID
		data *models.DraftOrder
	}
	tests := []struct {
		name  string
		s     *DraftOrderService
		args  args
		want  *models.DraftOrder
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.id, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DraftOrderService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DraftOrderService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

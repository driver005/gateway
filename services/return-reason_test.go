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

func TestNewReturnReasonService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *ReturnReasonService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReturnReasonService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReturnReasonService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReturnReasonService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *ReturnReasonService
		args args
		want *ReturnReasonService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReturnReasonService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReturnReasonService_Create(t *testing.T) {
	type args struct {
		data *types.CreateReturnReason
	}
	tests := []struct {
		name  string
		s     *ReturnReasonService
		args  args
		want  *models.ReturnReason
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReturnReasonService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ReturnReasonService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReturnReasonService_Update(t *testing.T) {
	type args struct {
		id   uuid.UUID
		data *types.UpdateReturnReason
	}
	tests := []struct {
		name  string
		s     *ReturnReasonService
		args  args
		want  *models.ReturnReason
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.id, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReturnReasonService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ReturnReasonService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReturnReasonService_List(t *testing.T) {
	type args struct {
		selector *types.FilterableReturnReason
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *ReturnReasonService
		args  args
		want  []models.ReturnReason
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReturnReasonService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ReturnReasonService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReturnReasonService_Retrieve(t *testing.T) {
	type args struct {
		returnReasonId uuid.UUID
		config         *sql.Options
	}
	tests := []struct {
		name  string
		s     *ReturnReasonService
		args  args
		want  *models.ReturnReason
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.returnReasonId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReturnReasonService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ReturnReasonService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReturnReasonService_Delete(t *testing.T) {
	type args struct {
		returnReasonId uuid.UUID
	}
	tests := []struct {
		name string
		s    *ReturnReasonService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.returnReasonId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReturnReasonService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

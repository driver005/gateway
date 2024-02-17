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

func TestNewDiscountConditionService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *DiscountConditionService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDiscountConditionService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDiscountConditionService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiscountConditionService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *DiscountConditionService
		args args
		want *DiscountConditionService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscountConditionService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiscountConditionService_Retrieve(t *testing.T) {
	type args struct {
		conditionId uuid.UUID
		config      *sql.Options
	}
	tests := []struct {
		name  string
		s     *DiscountConditionService
		args  args
		want  *models.DiscountCondition
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.conditionId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscountConditionService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DiscountConditionService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDiscountConditionService_ResolveConditionType(t *testing.T) {
	type args struct {
		data *types.DiscountConditionInput
	}
	tests := []struct {
		name  string
		s     *DiscountConditionService
		args  args
		want  *models.DiscountCondition
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ResolveConditionType(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscountConditionService.ResolveConditionType() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DiscountConditionService.ResolveConditionType() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDiscountConditionService_UpsertCondition(t *testing.T) {
	type args struct {
		data             *types.DiscountConditionInput
		overrideExisting bool
	}
	tests := []struct {
		name  string
		s     *DiscountConditionService
		args  args
		want  []models.DiscountCondition
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.UpsertCondition(tt.args.data, tt.args.overrideExisting)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscountConditionService.UpsertCondition() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DiscountConditionService.UpsertCondition() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDiscountConditionService_RemoveResources(t *testing.T) {
	type args struct {
		data *types.DiscountConditionInput
	}
	tests := []struct {
		name string
		s    *DiscountConditionService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.RemoveResources(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscountConditionService.RemoveResources() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiscountConditionService_Delete(t *testing.T) {
	type args struct {
		conditionId uuid.UUID
	}
	tests := []struct {
		name string
		s    *DiscountConditionService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.conditionId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscountConditionService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

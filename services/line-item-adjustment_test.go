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

func TestNewLineItemAdjustmentService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *LineItemAdjustmentService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLineItemAdjustmentService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLineItemAdjustmentService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineItemAdjustmentService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *LineItemAdjustmentService
		args args
		want *LineItemAdjustmentService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineItemAdjustmentService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineItemAdjustmentService_Retrieve(t *testing.T) {
	type args struct {
		id     uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *LineItemAdjustmentService
		args  args
		want  *models.LineItemAdjustment
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.id, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineItemAdjustmentService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LineItemAdjustmentService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLineItemAdjustmentService_Create(t *testing.T) {
	type args struct {
		data *models.LineItemAdjustment
	}
	tests := []struct {
		name  string
		s     *LineItemAdjustmentService
		args  args
		want  *models.LineItemAdjustment
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineItemAdjustmentService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LineItemAdjustmentService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLineItemAdjustmentService_Update(t *testing.T) {
	type args struct {
		id     uuid.UUID
		Update *models.LineItemAdjustment
	}
	tests := []struct {
		name  string
		s     *LineItemAdjustmentService
		args  args
		want  *models.LineItemAdjustment
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.id, tt.args.Update)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineItemAdjustmentService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LineItemAdjustmentService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLineItemAdjustmentService_List(t *testing.T) {
	type args struct {
		selector types.FilterableLineItemAdjustmentProps
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *LineItemAdjustmentService
		args  args
		want  []models.LineItemAdjustment
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineItemAdjustmentService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LineItemAdjustmentService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLineItemAdjustmentService_Delete(t *testing.T) {
	type args struct {
		id       uuid.UUID
		selector *models.LineItemAdjustment
		config   *sql.Options
	}
	tests := []struct {
		name string
		s    *LineItemAdjustmentService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.id, tt.args.selector, tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineItemAdjustmentService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineItemAdjustmentService_DeleteSlice(t *testing.T) {
	type args struct {
		ids      uuid.UUIDs
		selector []models.LineItemAdjustment
	}
	tests := []struct {
		name string
		s    *LineItemAdjustmentService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.DeleteSlice(tt.args.ids, tt.args.selector); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineItemAdjustmentService.DeleteSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineItemAdjustmentService_GenerateAdjustments(t *testing.T) {
	type args struct {
		calculationContextData types.CalculationContextData
		generatedLineItem      *models.LineItem
		context                *models.ProductVariant
	}
	tests := []struct {
		name  string
		s     *LineItemAdjustmentService
		args  args
		want  []models.LineItemAdjustment
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GenerateAdjustments(tt.args.calculationContextData, tt.args.generatedLineItem, tt.args.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineItemAdjustmentService.GenerateAdjustments() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LineItemAdjustmentService.GenerateAdjustments() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLineItemAdjustmentService_CreateAdjustmentForLineItem(t *testing.T) {
	type args struct {
		cart     *models.Cart
		lineItem *models.LineItem
	}
	tests := []struct {
		name  string
		s     *LineItemAdjustmentService
		args  args
		want  []models.LineItemAdjustment
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateAdjustmentForLineItem(tt.args.cart, tt.args.lineItem)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineItemAdjustmentService.CreateAdjustmentForLineItem() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LineItemAdjustmentService.CreateAdjustmentForLineItem() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLineItemAdjustmentService_CreateAdjustments(t *testing.T) {
	type args struct {
		cart     *models.Cart
		lineItem *models.LineItem
	}
	tests := []struct {
		name  string
		s     *LineItemAdjustmentService
		args  args
		want  []models.LineItemAdjustment
		want1 [][]models.LineItemAdjustment
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.CreateAdjustments(tt.args.cart, tt.args.lineItem)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineItemAdjustmentService.CreateAdjustments() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LineItemAdjustmentService.CreateAdjustments() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("LineItemAdjustmentService.CreateAdjustments() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

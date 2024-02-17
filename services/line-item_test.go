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

func TestNewLineItemService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *LineItemService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLineItemService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLineItemService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineItemService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *LineItemService
		args args
		want *LineItemService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineItemService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineItemService_Retrieve(t *testing.T) {
	type args struct {
		lineItemId uuid.UUID
		config     *sql.Options
	}
	tests := []struct {
		name  string
		s     *LineItemService
		args  args
		want  *models.LineItem
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.lineItemId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineItemService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LineItemService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLineItemService_List(t *testing.T) {
	type args struct {
		selector models.LineItem
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *LineItemService
		args  args
		want  []models.LineItem
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineItemService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LineItemService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLineItemService_CreateReturnLines(t *testing.T) {
	type args struct {
		returnId uuid.UUID
		cartId   uuid.UUID
	}
	tests := []struct {
		name  string
		s     *LineItemService
		args  args
		want  *models.LineItem
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateReturnLines(tt.args.returnId, tt.args.cartId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineItemService.CreateReturnLines() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LineItemService.CreateReturnLines() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLineItemService_Generate(t *testing.T) {
	type args struct {
		variantId uuid.UUID
		variant   []types.GenerateInputData
		regionId  uuid.UUID
		quantity  int
		context   types.GenerateLineItemContext
	}
	tests := []struct {
		name  string
		s     *LineItemService
		args  args
		want  []models.LineItem
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Generate(tt.args.variantId, tt.args.variant, tt.args.regionId, tt.args.quantity, tt.args.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineItemService.Generate() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LineItemService.Generate() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLineItemService_generateLineItem(t *testing.T) {
	type args struct {
		variant  models.ProductVariant
		quantity int
		context  GenerateLineItemContext
	}
	tests := []struct {
		name  string
		s     *LineItemService
		args  args
		want  *models.LineItem
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.generateLineItem(tt.args.variant, tt.args.quantity, tt.args.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineItemService.generateLineItem() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LineItemService.generateLineItem() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLineItemService_Create(t *testing.T) {
	type args struct {
		data []models.LineItem
	}
	tests := []struct {
		name  string
		s     *LineItemService
		args  args
		want  []models.LineItem
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineItemService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LineItemService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLineItemService_Update(t *testing.T) {
	type args struct {
		id       uuid.UUID
		selector *models.LineItem
		Update   *models.LineItem
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *LineItemService
		args  args
		want  *models.LineItem
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.id, tt.args.selector, tt.args.Update, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineItemService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LineItemService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLineItemService_Delete(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name string
		s    *LineItemService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineItemService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineItemService_DeleteWithTaxLines(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name string
		s    *LineItemService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.DeleteWithTaxLines(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineItemService.DeleteWithTaxLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineItemService_CreateTaxLine(t *testing.T) {
	type args struct {
		data *models.LineItemTaxLine
	}
	tests := []struct {
		name  string
		s     *LineItemService
		args  args
		want  *models.LineItemTaxLine
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateTaxLine(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineItemService.CreateTaxLine() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LineItemService.CreateTaxLine() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLineItemService_CloneTo(t *testing.T) {
	type args struct {
		ids     uuid.UUIDs
		data    *models.LineItem
		options map[string]interface{}
	}
	tests := []struct {
		name  string
		s     *LineItemService
		args  args
		want  []models.LineItem
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CloneTo(tt.args.ids, tt.args.data, tt.args.options)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineItemService.CloneTo() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LineItemService.CloneTo() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLineItemService_validateGenerateArguments(t *testing.T) {
	type args struct {
		variantId uuid.UUID
		variant   []types.GenerateInputData
		regionId  uuid.UUID
		quantity  int
		context   types.GenerateLineItemContext
	}
	tests := []struct {
		name    string
		s       *LineItemService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.validateGenerateArguments(tt.args.variantId, tt.args.variant, tt.args.regionId, tt.args.quantity, tt.args.context); (err != nil) != tt.wantErr {
				t.Errorf("LineItemService.validateGenerateArguments() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

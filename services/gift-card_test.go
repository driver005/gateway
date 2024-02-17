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

func TestNewGiftCardService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *GiftCardService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGiftCardService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGiftCardService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGiftCardService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *GiftCardService
		args args
		want *GiftCardService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GiftCardService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGiftCardService_GenerateCode(t *testing.T) {
	tests := []struct {
		name string
		s    *GiftCardService
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GenerateCode(); got != tt.want {
				t.Errorf("GiftCardService.GenerateCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGiftCardService_ListAndCount(t *testing.T) {
	type args struct {
		selector *types.FilterableGiftCard
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *GiftCardService
		args  args
		want  []models.GiftCard
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GiftCardService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GiftCardService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("GiftCardService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestGiftCardService_List(t *testing.T) {
	type args struct {
		selector *types.FilterableGiftCard
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *GiftCardService
		args  args
		want  []models.GiftCard
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GiftCardService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GiftCardService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGiftCardService_CreateTransaction(t *testing.T) {
	type args struct {
		data *types.CreateGiftCardTransactionInput
	}
	tests := []struct {
		name  string
		s     *GiftCardService
		args  args
		want  uuid.UUID
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateTransaction(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GiftCardService.CreateTransaction() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GiftCardService.CreateTransaction() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGiftCardService_Create(t *testing.T) {
	type args struct {
		data *types.CreateGiftCardInput
	}
	tests := []struct {
		name  string
		s     *GiftCardService
		args  args
		want  *models.GiftCard
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GiftCardService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GiftCardService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGiftCardService_ResolveTaxRate(t *testing.T) {
	type args struct {
		giftCardTaxRate float64
		region          *models.Region
	}
	tests := []struct {
		name string
		s    *GiftCardService
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ResolveTaxRate(tt.args.giftCardTaxRate, tt.args.region); got != tt.want {
				t.Errorf("GiftCardService.ResolveTaxRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGiftCardService_Retrieve(t *testing.T) {
	type args struct {
		selector *models.GiftCard
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *GiftCardService
		args  args
		want  *models.GiftCard
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GiftCardService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GiftCardService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGiftCardService_RetrieveById(t *testing.T) {
	type args struct {
		id     uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *GiftCardService
		args  args
		want  *models.GiftCard
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveById(tt.args.id, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GiftCardService.RetrieveById() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GiftCardService.RetrieveById() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGiftCardService_RetrieveByCode(t *testing.T) {
	type args struct {
		code   string
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *GiftCardService
		args  args
		want  *models.GiftCard
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveByCode(tt.args.code, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GiftCardService.RetrieveByCode() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GiftCardService.RetrieveByCode() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGiftCardService_Update(t *testing.T) {
	type args struct {
		id   uuid.UUID
		data *types.UpdateGiftCardInput
	}
	tests := []struct {
		name  string
		s     *GiftCardService
		args  args
		want  *models.GiftCard
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.id, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GiftCardService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GiftCardService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGiftCardService_Delete(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name string
		s    *GiftCardService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GiftCardService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

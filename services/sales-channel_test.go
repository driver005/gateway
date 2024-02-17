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

func TestNewSalesChannelService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *SalesChannelService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSalesChannelService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSalesChannelService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSalesChannelService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *SalesChannelService
		args args
		want *SalesChannelService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesChannelService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSalesChannelService_Retrieve(t *testing.T) {
	type args struct {
		selector *models.SalesChannel
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *SalesChannelService
		args  args
		want  *models.SalesChannel
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesChannelService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SalesChannelService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSalesChannelService_RetrieveById(t *testing.T) {
	type args struct {
		salesChannelId uuid.UUID
		config         *sql.Options
	}
	tests := []struct {
		name  string
		s     *SalesChannelService
		args  args
		want  *models.SalesChannel
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveById(tt.args.salesChannelId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesChannelService.RetrieveById() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SalesChannelService.RetrieveById() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSalesChannelService_RetrieveByName(t *testing.T) {
	type args struct {
		name   string
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *SalesChannelService
		args  args
		want  *models.SalesChannel
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveByName(tt.args.name, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesChannelService.RetrieveByName() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SalesChannelService.RetrieveByName() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSalesChannelService_List(t *testing.T) {
	type args struct {
		selector *types.FilterableSalesChannel
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *SalesChannelService
		args  args
		want  []models.SalesChannel
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesChannelService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SalesChannelService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSalesChannelService_ListAndCount(t *testing.T) {
	type args struct {
		selector *types.FilterableSalesChannel
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *SalesChannelService
		args  args
		want  []models.SalesChannel
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesChannelService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SalesChannelService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("SalesChannelService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestSalesChannelService_Create(t *testing.T) {
	type args struct {
		data *types.CreateSalesChannelInput
	}
	tests := []struct {
		name  string
		s     *SalesChannelService
		args  args
		want  *models.SalesChannel
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesChannelService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SalesChannelService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSalesChannelService_Update(t *testing.T) {
	type args struct {
		salesChannelId uuid.UUID
		data           *types.UpdateSalesChannelInput
	}
	tests := []struct {
		name  string
		s     *SalesChannelService
		args  args
		want  *models.SalesChannel
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.salesChannelId, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesChannelService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SalesChannelService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSalesChannelService_Delete(t *testing.T) {
	type args struct {
		salesChannelId uuid.UUID
	}
	tests := []struct {
		name string
		s    *SalesChannelService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.salesChannelId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesChannelService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSalesChannelService_CreateDefault(t *testing.T) {
	tests := []struct {
		name  string
		s     *SalesChannelService
		want  *models.SalesChannel
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateDefault()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesChannelService.CreateDefault() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SalesChannelService.CreateDefault() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSalesChannelService_RetrieveDefault(t *testing.T) {
	tests := []struct {
		name  string
		s     *SalesChannelService
		want  *models.SalesChannel
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveDefault()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesChannelService.RetrieveDefault() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SalesChannelService.RetrieveDefault() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSalesChannelService_ListProductIdsBySalesChannelIds(t *testing.T) {
	type args struct {
		salesChannelIds uuid.UUIDs
	}
	tests := []struct {
		name  string
		s     *SalesChannelService
		args  args
		want  map[string][]string
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ListProductIdsBySalesChannelIds(tt.args.salesChannelIds)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesChannelService.ListProductIdsBySalesChannelIds() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SalesChannelService.ListProductIdsBySalesChannelIds() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSalesChannelService_AddProducts(t *testing.T) {
	type args struct {
		salesChannelId uuid.UUID
		productIds     uuid.UUIDs
	}
	tests := []struct {
		name  string
		s     *SalesChannelService
		args  args
		want  *models.SalesChannel
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AddProducts(tt.args.salesChannelId, tt.args.productIds)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesChannelService.AddProducts() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SalesChannelService.AddProducts() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSalesChannelService_RemoveProducts(t *testing.T) {
	type args struct {
		salesChannelId uuid.UUID
		productIds     uuid.UUIDs
	}
	tests := []struct {
		name  string
		s     *SalesChannelService
		args  args
		want  *models.SalesChannel
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RemoveProducts(tt.args.salesChannelId, tt.args.productIds)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesChannelService.RemoveProducts() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SalesChannelService.RemoveProducts() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

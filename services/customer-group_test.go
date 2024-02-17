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

func TestNewCustomerGroupService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *CustomerGroupService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomerGroupService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomerGroupService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerGroupService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *CustomerGroupService
		args args
		want *CustomerGroupService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerGroupService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerGroupService_Retrieve(t *testing.T) {
	type args struct {
		customerGroupId uuid.UUID
		config          *sql.Options
	}
	tests := []struct {
		name  string
		s     *CustomerGroupService
		args  args
		want  *models.CustomerGroup
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.customerGroupId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerGroupService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerGroupService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCustomerGroupService_Create(t *testing.T) {
	type args struct {
		data *types.CreateCustomerGroup
	}
	tests := []struct {
		name  string
		s     *CustomerGroupService
		args  args
		want  *models.CustomerGroup
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerGroupService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerGroupService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCustomerGroupService_AddCustomers(t *testing.T) {
	type args struct {
		id          uuid.UUID
		customerIds uuid.UUIDs
	}
	tests := []struct {
		name  string
		s     *CustomerGroupService
		args  args
		want  *models.CustomerGroup
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AddCustomers(tt.args.id, tt.args.customerIds)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerGroupService.AddCustomers() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerGroupService.AddCustomers() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCustomerGroupService_Update(t *testing.T) {
	type args struct {
		customerGroupId uuid.UUID
		data            *types.UpdateCustomerGroup
	}
	tests := []struct {
		name  string
		s     *CustomerGroupService
		args  args
		want  *models.CustomerGroup
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.customerGroupId, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerGroupService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerGroupService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCustomerGroupService_Delete(t *testing.T) {
	type args struct {
		groupId uuid.UUID
	}
	tests := []struct {
		name string
		s    *CustomerGroupService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.groupId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerGroupService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerGroupService_List(t *testing.T) {
	type args struct {
		selector *types.FilterableCustomerGroup
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *CustomerGroupService
		args  args
		want  []models.CustomerGroup
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerGroupService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerGroupService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCustomerGroupService_ListAndCount(t *testing.T) {
	type args struct {
		selector *types.FilterableCustomerGroup
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *CustomerGroupService
		args  args
		want  []models.CustomerGroup
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerGroupService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerGroupService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("CustomerGroupService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestCustomerGroupService_RemoveCustomer(t *testing.T) {
	type args struct {
		id          uuid.UUID
		customerIds uuid.UUIDs
	}
	tests := []struct {
		name  string
		s     *CustomerGroupService
		args  args
		want  *models.CustomerGroup
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RemoveCustomer(tt.args.id, tt.args.customerIds)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerGroupService.RemoveCustomer() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerGroupService.RemoveCustomer() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCustomerGroupService_handleCreationFail(t *testing.T) {
	type args struct {
		id  uuid.UUID
		ids uuid.UUIDs
		err *utils.ApplictaionError
	}
	tests := []struct {
		name string
		s    *CustomerGroupService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.handleCreationFail(tt.args.id, tt.args.ids, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerGroupService.handleCreationFail() = %v, want %v", got, tt.want)
			}
		})
	}
}

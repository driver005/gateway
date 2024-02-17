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

func TestNewCustomerService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *CustomerService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomerService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomerService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *CustomerService
		args args
		want *CustomerService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerService_HashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		s       *CustomerService
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.HashPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerService.HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CustomerService.HashPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerService_GenerateResetPasswordToken(t *testing.T) {
	type args struct {
		customerId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *CustomerService
		args  args
		want  *string
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GenerateResetPasswordToken(tt.args.customerId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerService.GenerateResetPasswordToken() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerService.GenerateResetPasswordToken() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCustomerService_List(t *testing.T) {
	type args struct {
		selector *types.FilterableCustomer
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *CustomerService
		args  args
		want  []models.Customer
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCustomerService_ListAndCount(t *testing.T) {
	type args struct {
		selector *types.FilterableCustomer
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *CustomerService
		args  args
		want  []models.Customer
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("CustomerService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestCustomerService_Count(t *testing.T) {
	tests := []struct {
		name  string
		s     *CustomerService
		want  *int64
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Count()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerService.Count() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerService.Count() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCustomerService_Retrieve(t *testing.T) {
	type args struct {
		selector models.Customer
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *CustomerService
		args  args
		want  *models.Customer
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCustomerService_RetrieveByEmail(t *testing.T) {
	type args struct {
		email  string
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *CustomerService
		args  args
		want  *models.Customer
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveByEmail(tt.args.email, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerService.RetrieveByEmail() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerService.RetrieveByEmail() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCustomerService_RetrieveUnregisteredByEmail(t *testing.T) {
	type args struct {
		email  string
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *CustomerService
		args  args
		want  *models.Customer
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveUnregisteredByEmail(tt.args.email, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerService.RetrieveUnregisteredByEmail() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerService.RetrieveUnregisteredByEmail() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCustomerService_RetrieveRegisteredByEmail(t *testing.T) {
	type args struct {
		email  string
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *CustomerService
		args  args
		want  *models.Customer
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveRegisteredByEmail(tt.args.email, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerService.RetrieveRegisteredByEmail() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerService.RetrieveRegisteredByEmail() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCustomerService_ListByEmail(t *testing.T) {
	type args struct {
		email  string
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *CustomerService
		args  args
		want  []models.Customer
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ListByEmail(tt.args.email, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerService.ListByEmail() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerService.ListByEmail() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCustomerService_RetrieveByPhone(t *testing.T) {
	type args struct {
		phone  string
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *CustomerService
		args  args
		want  *models.Customer
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveByPhone(tt.args.phone, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerService.RetrieveByPhone() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerService.RetrieveByPhone() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCustomerService_RetrieveById(t *testing.T) {
	type args struct {
		id     uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *CustomerService
		args  args
		want  *models.Customer
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveById(tt.args.id, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerService.RetrieveById() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerService.RetrieveById() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCustomerService_Create(t *testing.T) {
	type args struct {
		data *types.CreateCustomerInput
	}
	tests := []struct {
		name  string
		s     *CustomerService
		args  args
		want  *models.Customer
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCustomerService_Update(t *testing.T) {
	type args struct {
		userId uuid.UUID
		data   *types.UpdateCustomerInput
	}
	tests := []struct {
		name  string
		s     *CustomerService
		args  args
		want  *models.Customer
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.userId, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCustomerService_UpdateBillingAddress(t *testing.T) {
	type args struct {
		model   *models.Customer
		id      uuid.UUID
		address *models.Address
	}
	tests := []struct {
		name string
		s    *CustomerService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UpdateBillingAddress(tt.args.model, tt.args.id, tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerService.UpdateBillingAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerService_UpdateAddress(t *testing.T) {
	type args struct {
		customerId uuid.UUID
		addressId  uuid.UUID
		model      *models.Address
	}
	tests := []struct {
		name  string
		s     *CustomerService
		args  args
		want  *models.Address
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.UpdateAddress(tt.args.customerId, tt.args.addressId, tt.args.model)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerService.UpdateAddress() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerService.UpdateAddress() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCustomerService_RemoveAddress(t *testing.T) {
	type args struct {
		customerId uuid.UUID
		addressId  uuid.UUID
	}
	tests := []struct {
		name    string
		s       *CustomerService
		args    args
		wantErr *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotErr := tt.s.RemoveAddress(tt.args.customerId, tt.args.addressId); !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("CustomerService.RemoveAddress() = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestCustomerService_AddAddress(t *testing.T) {
	type args struct {
		customerId uuid.UUID
		address    *models.Address
	}
	tests := []struct {
		name  string
		s     *CustomerService
		args  args
		want  *models.Customer
		want1 *models.Address
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.AddAddress(tt.args.customerId, tt.args.address)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerService.AddAddress() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerService.AddAddress() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("CustomerService.AddAddress() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestCustomerService_Delete(t *testing.T) {
	type args struct {
		customerId uuid.UUID
	}
	tests := []struct {
		name string
		s    *CustomerService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.customerId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

package services

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

func TestNewUserService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *UserService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *UserService
		args args
		want *UserService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_HashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		s       *UserService
		args    args
		wantErr bool
	}{
		{
			name:    "ValidPassword",
			s:       &UserService{}, // Assuming you have an instance of UserService
			args:    args{password: "password123"},
			wantErr: false,
		},
		{
			name:    "EmptyPassword",
			s:       &UserService{},
			args:    args{password: ""},
			wantErr: true, // Assuming empty password should result in an error
		},
		{
			name:    "WeakPassword",
			s:       &UserService{},
			args:    args{password: "weak"},
			wantErr: false, // Assuming weak password should result in an error
		},
		{
			name:    "LongPassword",
			s:       &UserService{},
			args:    args{password: "averylongpasswordthatexceedsthelimitofcharacters"},
			wantErr: false,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.s.HashPassword(tt.args.password)
			fmt.Println(err)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && err == nil {
				t.Logf("UserService.HashPassword() did not return an error, but it was expected")
			}
			// Optionally, you can add additional checks on the hashed password
			// For example, if you're using a specific hashing algorithm, you can validate the format of the hashed password
		})
	}
}

func TestUserService_List(t *testing.T) {
	type args struct {
		selector *types.FilterableUser
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *UserService
		args  args
		want  []models.User
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("UserService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUserService_ListAndCount(t *testing.T) {
	type args struct {
		selector *types.FilterableUser
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *UserService
		args  args
		want  []models.User
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("UserService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("UserService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestUserService_Retrieve(t *testing.T) {
	type args struct {
		userId uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *UserService
		args  args
		want  *models.User
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.userId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("UserService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUserService_RetrieveByApiToken(t *testing.T) {
	type args struct {
		apiToken  string
		relations []string
	}
	tests := []struct {
		name  string
		s     *UserService
		args  args
		want  *models.User
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveByApiToken(tt.args.apiToken, tt.args.relations)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.RetrieveByApiToken() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("UserService.RetrieveByApiToken() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUserService_RetrieveByEmail(t *testing.T) {
	type args struct {
		email  string
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *UserService
		args  args
		want  *models.User
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveByEmail(tt.args.email, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.RetrieveByEmail() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("UserService.RetrieveByEmail() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUserService_Create(t *testing.T) {
	type args struct {
		data *types.CreateUserInput
	}
	tests := []struct {
		name  string
		s     *UserService
		args  args
		want  *models.User
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("UserService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUserService_Update(t *testing.T) {
	type args struct {
		userId uuid.UUID
		data   *types.UpdateUserInput
	}
	tests := []struct {
		name  string
		s     *UserService
		args  args
		want  *models.User
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.userId, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("UserService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUserService_Delete(t *testing.T) {
	type args struct {
		userId uuid.UUID
	}
	tests := []struct {
		name string
		s    *UserService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.userId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_SetPassword(t *testing.T) {
	type args struct {
		userId   uuid.UUID
		password string
	}
	tests := []struct {
		name  string
		s     *UserService
		args  args
		want  *models.User
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.SetPassword(tt.args.userId, tt.args.password)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.SetPassword() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("UserService.SetPassword() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUserService_GenerateResetPasswordToken(t *testing.T) {
	type args struct {
		userId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *UserService
		args  args
		want  *string
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GenerateResetPasswordToken(tt.args.userId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.GenerateResetPasswordToken() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("UserService.GenerateResetPasswordToken() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

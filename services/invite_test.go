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

func TestNewInviteService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *InviteService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInviteService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInviteService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInviteService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *InviteService
		args args
		want *InviteService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InviteService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInviteService_List(t *testing.T) {
	type args struct {
		selector *types.FilterableInvite
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *InviteService
		args  args
		want  []models.Invite
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InviteService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("InviteService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestInviteService_Create(t *testing.T) {
	type args struct {
		data          *types.CreateInviteInput
		validDuration int
	}
	tests := []struct {
		name string
		s    *InviteService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Create(tt.args.data, tt.args.validDuration); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InviteService.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInviteService_Delete(t *testing.T) {
	type args struct {
		inviteId uuid.UUID
	}
	tests := []struct {
		name string
		s    *InviteService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.inviteId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InviteService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInviteService_Accept(t *testing.T) {
	type args struct {
		token string
		user  *models.User
	}
	tests := []struct {
		name  string
		s     *InviteService
		args  args
		want  *models.User
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Accept(tt.args.token, tt.args.user)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InviteService.Accept() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("InviteService.Accept() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestInviteService_Resend(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name string
		s    *InviteService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Resend(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InviteService.Resend() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
	"github.com/sarulabs/di"
)

func TestNewOAuthService(t *testing.T) {
	type args struct {
		container di.Container
		r         Registry
	}
	tests := []struct {
		name string
		args args
		want *OAuthService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOAuthService(tt.args.container, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOAuthService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOAuthService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *OAuthService
		args args
		want *OAuthService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OAuthService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOAuthService_RetrieveByName(t *testing.T) {
	type args struct {
		appName string
		config  *sql.Options
	}
	tests := []struct {
		name  string
		s     *OAuthService
		args  args
		want  *models.OAuth
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveByName(tt.args.appName, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OAuthService.RetrieveByName() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OAuthService.RetrieveByName() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOAuthService_Retrieve(t *testing.T) {
	type args struct {
		id     uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *OAuthService
		args  args
		want  *models.OAuth
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.id, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OAuthService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OAuthService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOAuthService_List(t *testing.T) {
	type args struct {
		selector *models.OAuth
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *OAuthService
		args  args
		want  []models.OAuth
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OAuthService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OAuthService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOAuthService_Create(t *testing.T) {
	type args struct {
		data *types.CreateOauthInput
	}
	tests := []struct {
		name  string
		s     *OAuthService
		args  args
		want  *models.OAuth
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OAuthService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OAuthService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOAuthService_Update(t *testing.T) {
	type args struct {
		id   uuid.UUID
		data *types.UpdateOauthInput
	}
	tests := []struct {
		name  string
		s     *OAuthService
		args  args
		want  *models.OAuth
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.id, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OAuthService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OAuthService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOAuthService_RegisterOauthApp(t *testing.T) {
	type args struct {
		appDetails *types.CreateOauthInput
	}
	tests := []struct {
		name  string
		s     *OAuthService
		args  args
		want  *models.OAuth
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RegisterOauthApp(tt.args.appDetails)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OAuthService.RegisterOauthApp() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OAuthService.RegisterOauthApp() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOAuthService_GenerateToken(t *testing.T) {
	type args struct {
		appName string
		code    string
		state   string
	}
	tests := []struct {
		name  string
		s     *OAuthService
		args  args
		want  *models.OAuth
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GenerateToken(tt.args.appName, tt.args.code, tt.args.state)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OAuthService.GenerateToken() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OAuthService.GenerateToken() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOAuthService_RefreshToken(t *testing.T) {
	type args struct {
		appName string
	}
	tests := []struct {
		name  string
		s     *OAuthService
		args  args
		want  *models.OAuth
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RefreshToken(tt.args.appName)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OAuthService.RefreshToken() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OAuthService.RefreshToken() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/driver005/gateway/types"
)

func TestNewAuthService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *AuthService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *AuthService
		args args
		want *AuthService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthService_ComparePassword(t *testing.T) {
	type args struct {
		password string
		hash     string
	}
	tests := []struct {
		name string
		s    *AuthService
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ComparePassword(tt.args.password, tt.args.hash); got != tt.want {
				t.Errorf("AuthService.ComparePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthService_AuthenticateAPIToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name string
		s    *AuthService
		args args
		want types.AuthenticateResult
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.AuthenticateAPIToken(tt.args.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthService.AuthenticateAPIToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthService_Authenticate(t *testing.T) {
	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name string
		s    *AuthService
		args args
		want types.AuthenticateResult
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Authenticate(tt.args.email, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthService.Authenticate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthService_AuthenticateCustomer(t *testing.T) {
	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name string
		s    *AuthService
		args args
		want types.AuthenticateResult
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.AuthenticateCustomer(tt.args.email, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthService.AuthenticateCustomer() = %v, want %v", got, tt.want)
			}
		})
	}
}

package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang-jwt/jwt"
)

func TestNewTockenService(t *testing.T) {
	type args struct {
		secretKey []byte
	}
	tests := []struct {
		name string
		args args
		want *TockenService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTockenService(tt.args.secretKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTockenService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTockenService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *TockenService
		args args
		want *TockenService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TockenService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTockenService_VerifyToken(t *testing.T) {
	type args struct {
		tocken string
	}
	tests := []struct {
		name    string
		s       *TockenService
		args    args
		want    *jwt.Token
		want1   jwt.MapClaims
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.s.VerifyToken(tt.args.tocken)
			if (err != nil) != tt.wantErr {
				t.Errorf("TockenService.VerifyToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TockenService.VerifyToken() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TockenService.VerifyToken() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTockenService_VerifyTokenWithSecret(t *testing.T) {
	type args struct {
		tocken string
		secret []byte
	}
	tests := []struct {
		name    string
		s       *TockenService
		args    args
		want    *jwt.Token
		want1   jwt.MapClaims
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.s.VerifyTokenWithSecret(tt.args.tocken, tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("TockenService.VerifyTokenWithSecret() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TockenService.VerifyTokenWithSecret() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TockenService.VerifyTokenWithSecret() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTockenService_SignToken(t *testing.T) {
	type args struct {
		data map[string]interface{}
	}
	tests := []struct {
		name    string
		s       *TockenService
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.SignToken(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("TockenService.SignToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TockenService.SignToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTockenService_SignTokenWithSecret(t *testing.T) {
	type args struct {
		data   map[string]interface{}
		secret []byte
	}
	tests := []struct {
		name    string
		s       *TockenService
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.SignTokenWithSecret(tt.args.data, tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("TockenService.SignTokenWithSecret() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TockenService.SignTokenWithSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
)

func TestNewIdempotencyKeyService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *IdempotencyKeyService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIdempotencyKeyService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIdempotencyKeyService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIdempotencyKeyService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *IdempotencyKeyService
		args args
		want *IdempotencyKeyService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IdempotencyKeyService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIdempotencyKeyService_InitializeRequest(t *testing.T) {
	type args struct {
		headerKey string
		reqMethod string
		reqParams core.JSONB
		reqPath   string
	}
	tests := []struct {
		name  string
		s     *IdempotencyKeyService
		args  args
		want  *models.IdempotencyKey
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.InitializeRequest(tt.args.headerKey, tt.args.reqMethod, tt.args.reqParams, tt.args.reqPath)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IdempotencyKeyService.InitializeRequest() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("IdempotencyKeyService.InitializeRequest() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestIdempotencyKeyService_Create(t *testing.T) {
	type args struct {
		payload *types.CreateIdempotencyKeyInput
	}
	tests := []struct {
		name  string
		s     *IdempotencyKeyService
		args  args
		want  *models.IdempotencyKey
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.payload)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IdempotencyKeyService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("IdempotencyKeyService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestIdempotencyKeyService_Retrieve(t *testing.T) {
	type args struct {
		idempotencyKey string
	}
	tests := []struct {
		name  string
		s     *IdempotencyKeyService
		args  args
		want  *models.IdempotencyKey
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.idempotencyKey)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IdempotencyKeyService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("IdempotencyKeyService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestIdempotencyKeyService_Lock(t *testing.T) {
	type args struct {
		idempotencyKey string
	}
	tests := []struct {
		name  string
		s     *IdempotencyKeyService
		args  args
		want  *models.IdempotencyKey
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Lock(tt.args.idempotencyKey)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IdempotencyKeyService.Lock() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("IdempotencyKeyService.Lock() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestIdempotencyKeyService_Update(t *testing.T) {
	type args struct {
		idempotencyKey string
		data           *models.IdempotencyKey
	}
	tests := []struct {
		name  string
		s     *IdempotencyKeyService
		args  args
		want  *models.IdempotencyKey
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.idempotencyKey, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IdempotencyKeyService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("IdempotencyKeyService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestIdempotencyKeyService_WorkStage(t *testing.T) {
	type args struct {
		idempotencyKey string
		callback       func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError)
	}
	tests := []struct {
		name  string
		s     *IdempotencyKeyService
		args  args
		want  *models.IdempotencyKey
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.WorkStage(tt.args.idempotencyKey, tt.args.callback)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IdempotencyKeyService.WorkStage() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("IdempotencyKeyService.WorkStage() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

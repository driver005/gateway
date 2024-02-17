package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

func TestNewAnalyticsConfigService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *AnalyticsConfigService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAnalyticsConfigService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAnalyticsConfigService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyticsConfigService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *AnalyticsConfigService
		args args
		want *AnalyticsConfigService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AnalyticsConfigService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyticsConfigService_Retrive(t *testing.T) {
	type args struct {
		userId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *AnalyticsConfigService
		args  args
		want  *models.AnalyticsConfig
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrive(tt.args.userId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AnalyticsConfigService.Retrive() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("AnalyticsConfigService.Retrive() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAnalyticsConfigService_Create(t *testing.T) {
	type args struct {
		userId uuid.UUID
		data   *types.CreateAnalyticsConfig
	}
	tests := []struct {
		name  string
		s     *AnalyticsConfigService
		args  args
		want  *models.AnalyticsConfig
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.userId, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AnalyticsConfigService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("AnalyticsConfigService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAnalyticsConfigService_Update(t *testing.T) {
	type args struct {
		userId uuid.UUID
		data   *types.UpdateAnalyticsConfig
	}
	tests := []struct {
		name  string
		s     *AnalyticsConfigService
		args  args
		want  *models.AnalyticsConfig
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.userId, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AnalyticsConfigService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("AnalyticsConfigService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAnalyticsConfigService_Delete(t *testing.T) {
	type args struct {
		userId uuid.UUID
	}
	tests := []struct {
		name string
		s    *AnalyticsConfigService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.userId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AnalyticsConfigService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

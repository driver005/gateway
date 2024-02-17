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

func TestNewStagedJobService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *StagedJobService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStagedJobService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStagedJobService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStagedJobService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *StagedJobService
		args args
		want *StagedJobService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StagedJobService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStagedJobService_List(t *testing.T) {
	type args struct {
		selector models.StagedJob
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *StagedJobService
		args  args
		want  []models.StagedJob
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StagedJobService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("StagedJobService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStagedJobService_Delete(t *testing.T) {
	type args struct {
		stagedJobIds uuid.UUIDs
	}
	tests := []struct {
		name string
		s    *StagedJobService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.stagedJobIds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StagedJobService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStagedJobService_Create(t *testing.T) {
	type args struct {
		data []types.EmitData
	}
	tests := []struct {
		name string
		s    *StagedJobService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Create(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StagedJobService.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

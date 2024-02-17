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

func TestNewClaimItemService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *ClaimItemService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClaimItemService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClaimItemService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClaimItemService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *ClaimItemService
		args args
		want *ClaimItemService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClaimItemService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClaimItemService_Retrieve(t *testing.T) {
	type args struct {
		claimItemId uuid.UUID
		config      *sql.Options
	}
	tests := []struct {
		name  string
		s     *ClaimItemService
		args  args
		want  *models.ClaimItem
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.claimItemId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClaimItemService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ClaimItemService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestClaimItemService_List(t *testing.T) {
	type args struct {
		selector models.ClaimItem
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *ClaimItemService
		args  args
		want  []models.ClaimItem
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClaimItemService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ClaimItemService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestClaimItemService_Update(t *testing.T) {
	type args struct {
		id   uuid.UUID
		data *types.UpdateClaimItemInput
	}
	tests := []struct {
		name  string
		s     *ClaimItemService
		args  args
		want  *models.ClaimItem
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.id, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClaimItemService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ClaimItemService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestClaimItemService_Create(t *testing.T) {
	type args struct {
		data *types.CreateClaimItemInput
	}
	tests := []struct {
		name  string
		s     *ClaimItemService
		args  args
		want  *models.ClaimItem
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClaimItemService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ClaimItemService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

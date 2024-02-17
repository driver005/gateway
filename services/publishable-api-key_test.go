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

func TestNewPublishableApiKeyService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *PublishableApiKeyService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPublishableApiKeyService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPublishableApiKeyService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPublishableApiKeyService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *PublishableApiKeyService
		args args
		want *PublishableApiKeyService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PublishableApiKeyService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPublishableApiKeyService_Create(t *testing.T) {
	type args struct {
		data           *types.CreatePublishableApiKeyInput
		loggedInUserId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *PublishableApiKeyService
		args  args
		want  *models.PublishableApiKey
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data, tt.args.loggedInUserId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PublishableApiKeyService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PublishableApiKeyService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPublishableApiKeyService_RetrieveById(t *testing.T) {
	type args struct {
		id     uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *PublishableApiKeyService
		args  args
		want  *models.PublishableApiKey
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveById(tt.args.id, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PublishableApiKeyService.RetrieveById() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PublishableApiKeyService.RetrieveById() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPublishableApiKeyService_Retrieve(t *testing.T) {
	type args struct {
		selector *models.PublishableApiKey
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *PublishableApiKeyService
		args  args
		want  *models.PublishableApiKey
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PublishableApiKeyService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PublishableApiKeyService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPublishableApiKeyService_ListAndCount(t *testing.T) {
	type args struct {
		selector *models.PublishableApiKey
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *PublishableApiKeyService
		args  args
		want  []models.PublishableApiKey
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PublishableApiKeyService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PublishableApiKeyService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("PublishableApiKeyService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestPublishableApiKeyService_Update(t *testing.T) {
	type args struct {
		id   uuid.UUID
		data *types.UpdatePublishableApiKeyInput
	}
	tests := []struct {
		name  string
		s     *PublishableApiKeyService
		args  args
		want  *models.PublishableApiKey
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.id, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PublishableApiKeyService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PublishableApiKeyService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPublishableApiKeyService_Delete(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name string
		s    *PublishableApiKeyService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PublishableApiKeyService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPublishableApiKeyService_Revoke(t *testing.T) {
	type args struct {
		id             uuid.UUID
		loggedInUserId uuid.UUID
	}
	tests := []struct {
		name string
		s    *PublishableApiKeyService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Revoke(tt.args.id, tt.args.loggedInUserId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PublishableApiKeyService.Revoke() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPublishableApiKeyService_IsValid(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name  string
		s     *PublishableApiKeyService
		args  args
		want  bool
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.IsValid(tt.args.id)
			if got != tt.want {
				t.Errorf("PublishableApiKeyService.IsValid() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PublishableApiKeyService.IsValid() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPublishableApiKeyService_AddSalesChannels(t *testing.T) {
	type args struct {
		id              uuid.UUID
		salesChannelIds uuid.UUIDs
	}
	tests := []struct {
		name string
		s    *PublishableApiKeyService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.AddSalesChannels(tt.args.id, tt.args.salesChannelIds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PublishableApiKeyService.AddSalesChannels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPublishableApiKeyService_RemoveSalesChannels(t *testing.T) {
	type args struct {
		id              uuid.UUID
		salesChannelIds uuid.UUIDs
	}
	tests := []struct {
		name string
		s    *PublishableApiKeyService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.RemoveSalesChannels(tt.args.id, tt.args.salesChannelIds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PublishableApiKeyService.RemoveSalesChannels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPublishableApiKeyService_ListSalesChannels(t *testing.T) {
	type args struct {
		id uuid.UUID
		q  *string
	}
	tests := []struct {
		name  string
		s     *PublishableApiKeyService
		args  args
		want  []models.SalesChannel
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ListSalesChannels(tt.args.id, tt.args.q)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PublishableApiKeyService.ListSalesChannels() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PublishableApiKeyService.ListSalesChannels() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPublishableApiKeyService_GetResourceScopes(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name  string
		s     *PublishableApiKeyService
		args  args
		want  *types.PublishableApiKeyScopes
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetResourceScopes(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PublishableApiKeyService.GetResourceScopes() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PublishableApiKeyService.GetResourceScopes() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

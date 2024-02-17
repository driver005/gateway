package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/sarulabs/di"
)

func TestNewNotificationService(t *testing.T) {
	type args struct {
		container           di.Container
		subscribers         map[string]uuid.UUIDs
		attachmentGenerator interface{}
		r                   Registry
	}
	tests := []struct {
		name string
		args args
		want *NotificationService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNotificationService(tt.args.container, tt.args.subscribers, tt.args.attachmentGenerator, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNotificationService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotificationService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *NotificationService
		args args
		want *NotificationService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotificationService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotificationService_RegisterAttachmentGenerator(t *testing.T) {
	type args struct {
		service interface{}
	}
	tests := []struct {
		name string
		s    *NotificationService
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.RegisterAttachmentGenerator(tt.args.service)
		})
	}
}

func TestNotificationService_RegisterInstalledProviders(t *testing.T) {
	type args struct {
		providers uuid.UUIDs
	}
	tests := []struct {
		name string
		s    *NotificationService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.RegisterInstalledProviders(tt.args.providers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotificationService.RegisterInstalledProviders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotificationService_List(t *testing.T) {
	type args struct {
		selector *types.FilterableNotification
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *NotificationService
		args  args
		want  []models.Notification
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotificationService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NotificationService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNotificationService_ListAndCount(t *testing.T) {
	type args struct {
		selector *types.FilterableNotification
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *NotificationService
		args  args
		want  []models.Notification
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotificationService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NotificationService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("NotificationService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestNotificationService_Retrieve(t *testing.T) {
	type args struct {
		id     uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *NotificationService
		args  args
		want  *models.Notification
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.id, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotificationService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NotificationService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNotificationService_Subscribe(t *testing.T) {
	type args struct {
		eventName  string
		providerId uuid.UUID
	}
	tests := []struct {
		name string
		s    *NotificationService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Subscribe(tt.args.eventName, tt.args.providerId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotificationService.Subscribe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotificationService_RetrieveProvider(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name  string
		s     *NotificationService
		args  args
		want  interfaces.INotificationService
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveProvider(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotificationService.RetrieveProvider() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NotificationService.RetrieveProvider() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNotificationService_HandleEvent(t *testing.T) {
	type args struct {
		eventName string
		data      map[string]interface{}
	}
	tests := []struct {
		name string
		s    *NotificationService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.HandleEvent(tt.args.eventName, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotificationService.HandleEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotificationService_Send(t *testing.T) {
	type args struct {
		event      string
		eventData  map[string]interface{}
		providerId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *NotificationService
		args  args
		want  *models.Notification
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Send(tt.args.event, tt.args.eventData, tt.args.providerId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotificationService.Send() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NotificationService.Send() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNotificationService_Resend(t *testing.T) {
	type args struct {
		id     uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *NotificationService
		args  args
		want  *models.Notification
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Resend(tt.args.id, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotificationService.Resend() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NotificationService.Resend() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

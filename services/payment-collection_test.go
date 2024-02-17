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

func TestNewPaymentCollectionService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *PaymentCollectionService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPaymentCollectionService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPaymentCollectionService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentCollectionService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *PaymentCollectionService
		args args
		want *PaymentCollectionService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentCollectionService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentCollectionService_Retrieve(t *testing.T) {
	type args struct {
		id     uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *PaymentCollectionService
		args  args
		want  *models.PaymentCollection
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.id, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentCollectionService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentCollectionService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentCollectionService_Create(t *testing.T) {
	type args struct {
		data *types.CreatePaymentCollectionInput
	}
	tests := []struct {
		name  string
		s     *PaymentCollectionService
		args  args
		want  *models.PaymentCollection
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentCollectionService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentCollectionService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentCollectionService_Update(t *testing.T) {
	type args struct {
		id     uuid.UUID
		Update *models.PaymentCollection
	}
	tests := []struct {
		name  string
		s     *PaymentCollectionService
		args  args
		want  *models.PaymentCollection
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.id, tt.args.Update)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentCollectionService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentCollectionService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentCollectionService_Delete(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name string
		s    *PaymentCollectionService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentCollectionService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentCollectionService_IsValidTotalAmount(t *testing.T) {
	type args struct {
		total         float64
		sessionsInput []types.PaymentCollectionsSessionsBatchInput
	}
	tests := []struct {
		name string
		s    *PaymentCollectionService
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsValidTotalAmount(tt.args.total, tt.args.sessionsInput); got != tt.want {
				t.Errorf("PaymentCollectionService.IsValidTotalAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentCollectionService_SetPaymentSessionsBatch(t *testing.T) {
	type args struct {
		id                uuid.UUID
		paymentCollection *models.PaymentCollection
		sessionsInput     []types.PaymentCollectionsSessionsBatchInput
		customerId        uuid.UUID
	}
	tests := []struct {
		name  string
		s     *PaymentCollectionService
		args  args
		want  *models.PaymentCollection
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.SetPaymentSessionsBatch(tt.args.id, tt.args.paymentCollection, tt.args.sessionsInput, tt.args.customerId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentCollectionService.SetPaymentSessionsBatch() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentCollectionService.SetPaymentSessionsBatch() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentCollectionService_SetPaymentSession(t *testing.T) {
	type args struct {
		id           uuid.UUID
		sessionInput *types.SessionsInput
		customerId   uuid.UUID
	}
	tests := []struct {
		name  string
		s     *PaymentCollectionService
		args  args
		want  *models.PaymentCollection
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.SetPaymentSession(tt.args.id, tt.args.sessionInput, tt.args.customerId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentCollectionService.SetPaymentSession() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentCollectionService.SetPaymentSession() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentCollectionService_RefreshPaymentSession(t *testing.T) {
	type args struct {
		id         uuid.UUID
		sessionId  uuid.UUID
		customerId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *PaymentCollectionService
		args  args
		want  *models.PaymentSession
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RefreshPaymentSession(tt.args.id, tt.args.sessionId, tt.args.customerId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentCollectionService.RefreshPaymentSession() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentCollectionService.RefreshPaymentSession() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentCollectionService_MarkAsAuthorized(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name  string
		s     *PaymentCollectionService
		args  args
		want  *models.PaymentCollection
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.MarkAsAuthorized(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentCollectionService.MarkAsAuthorized() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentCollectionService.MarkAsAuthorized() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentCollectionService_AuthorizePaymentSessions(t *testing.T) {
	type args struct {
		id         uuid.UUID
		sessionIds uuid.UUIDs
		context    map[string]interface{}
	}
	tests := []struct {
		name  string
		s     *PaymentCollectionService
		args  args
		want  *models.PaymentCollection
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AuthorizePaymentSessions(tt.args.id, tt.args.sessionIds, tt.args.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentCollectionService.AuthorizePaymentSessions() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentCollectionService.AuthorizePaymentSessions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

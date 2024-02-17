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

func TestNewPaymentService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *PaymentService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPaymentService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPaymentService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *PaymentService
		args args
		want *PaymentService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentService_Retrieve(t *testing.T) {
	type args struct {
		id     uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *PaymentService
		args  args
		want  *models.Payment
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.id, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentService_List(t *testing.T) {
	type args struct {
		selector models.Payment
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *PaymentService
		args  args
		want  []models.Payment
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentService_Create(t *testing.T) {
	type args struct {
		data *types.CreatePaymentInput
	}
	tests := []struct {
		name  string
		s     *PaymentService
		args  args
		want  *models.Payment
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentService_Update(t *testing.T) {
	type args struct {
		id   uuid.UUID
		data *types.UpdatePaymentInput
	}
	tests := []struct {
		name  string
		s     *PaymentService
		args  args
		want  *models.Payment
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.id, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentService_Capture(t *testing.T) {
	type args struct {
		id      uuid.UUID
		payment *models.Payment
	}
	tests := []struct {
		name  string
		s     *PaymentService
		args  args
		want  *models.Payment
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Capture(tt.args.id, tt.args.payment)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentService.Capture() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentService.Capture() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentService_Refund(t *testing.T) {
	type args struct {
		id      uuid.UUID
		payment *models.Payment
		amount  float64
		reason  models.RefundReason
		note    *string
	}
	tests := []struct {
		name  string
		s     *PaymentService
		args  args
		want  *models.Refund
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Refund(tt.args.id, tt.args.payment, tt.args.amount, tt.args.reason, tt.args.note)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentService.Refund() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentService.Refund() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

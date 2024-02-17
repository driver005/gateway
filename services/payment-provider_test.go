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

func TestNewPaymentProviderService(t *testing.T) {
	type args struct {
		container di.Container
		r         Registry
	}
	tests := []struct {
		name string
		args args
		want *PaymentProviderService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPaymentProviderService(tt.args.container, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPaymentProviderService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentProviderService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *PaymentProviderService
		args args
		want *PaymentProviderService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentProviderService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentProviderService_RegisterInstalledProviders(t *testing.T) {
	type args struct {
		providers uuid.UUIDs
	}
	tests := []struct {
		name string
		s    *PaymentProviderService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.RegisterInstalledProviders(tt.args.providers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentProviderService.RegisterInstalledProviders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentProviderService_RetrieveProvider(t *testing.T) {
	type args struct {
		providerId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *PaymentProviderService
		args  args
		want  interfaces.IPaymentProcessor
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveProvider(tt.args.providerId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentProviderService.RetrieveProvider() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentProviderService.RetrieveProvider() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentProviderService_List(t *testing.T) {
	tests := []struct {
		name  string
		s     *PaymentProviderService
		want  []models.PaymentProvider
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentProviderService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentProviderService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentProviderService_RetrievePayment(t *testing.T) {
	type args struct {
		id        uuid.UUID
		relations []string
	}
	tests := []struct {
		name  string
		s     *PaymentProviderService
		args  args
		want  *models.Payment
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrievePayment(tt.args.id, tt.args.relations)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentProviderService.RetrievePayment() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentProviderService.RetrievePayment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentProviderService_ListPayments(t *testing.T) {
	type args struct {
		selector models.Payment
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *PaymentProviderService
		args  args
		want  []models.Payment
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ListPayments(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentProviderService.ListPayments() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentProviderService.ListPayments() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentProviderService_RetrieveSession(t *testing.T) {
	type args struct {
		id        uuid.UUID
		relations []string
	}
	tests := []struct {
		name  string
		s     *PaymentProviderService
		args  args
		want  *models.PaymentSession
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveSession(tt.args.id, tt.args.relations)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentProviderService.RetrieveSession() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentProviderService.RetrieveSession() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentProviderService_CreateSession(t *testing.T) {
	type args struct {
		providerId uuid.UUID
		session    *types.PaymentSessionInput
	}
	tests := []struct {
		name  string
		s     *PaymentProviderService
		args  args
		want  *models.PaymentSession
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateSession(tt.args.providerId, tt.args.session)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentProviderService.CreateSession() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentProviderService.CreateSession() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentProviderService_RefreshSession(t *testing.T) {
	type args struct {
		paymentSession *models.PaymentSession
		sessionInput   *types.PaymentSessionInput
	}
	tests := []struct {
		name  string
		s     *PaymentProviderService
		args  args
		want  *models.PaymentSession
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RefreshSession(tt.args.paymentSession, tt.args.sessionInput)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentProviderService.RefreshSession() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentProviderService.RefreshSession() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentProviderService_UpdateSession(t *testing.T) {
	type args struct {
		paymentSession *models.PaymentSession
		sessionInput   *types.PaymentSessionInput
	}
	tests := []struct {
		name  string
		s     *PaymentProviderService
		args  args
		want  *models.PaymentSession
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.UpdateSession(tt.args.paymentSession, tt.args.sessionInput)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentProviderService.UpdateSession() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentProviderService.UpdateSession() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentProviderService_DeleteSession(t *testing.T) {
	type args struct {
		paymentSession *models.PaymentSession
	}
	tests := []struct {
		name string
		s    *PaymentProviderService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.DeleteSession(tt.args.paymentSession); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentProviderService.DeleteSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentProviderService_CreatePayment(t *testing.T) {
	type args struct {
		data *types.CreatePaymentInput
	}
	tests := []struct {
		name  string
		s     *PaymentProviderService
		args  args
		want  *models.Payment
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreatePayment(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentProviderService.CreatePayment() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentProviderService.CreatePayment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentProviderService_UpdatePayment(t *testing.T) {
	type args struct {
		id   uuid.UUID
		data *types.UpdatePaymentInput
	}
	tests := []struct {
		name  string
		s     *PaymentProviderService
		args  args
		want  *models.Payment
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.UpdatePayment(tt.args.id, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentProviderService.UpdatePayment() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentProviderService.UpdatePayment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentProviderService_AuthorizePayment(t *testing.T) {
	type args struct {
		paymentSession *models.PaymentSession
		context        map[string]interface{}
	}
	tests := []struct {
		name  string
		s     *PaymentProviderService
		args  args
		want  *models.PaymentSession
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AuthorizePayment(tt.args.paymentSession, tt.args.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentProviderService.AuthorizePayment() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentProviderService.AuthorizePayment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentProviderService_UpdateSessionData(t *testing.T) {
	type args struct {
		paymentSession *models.PaymentSession
		data           map[string]interface{}
	}
	tests := []struct {
		name  string
		s     *PaymentProviderService
		args  args
		want  *models.PaymentSession
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.UpdateSessionData(tt.args.paymentSession, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentProviderService.UpdateSessionData() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentProviderService.UpdateSessionData() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentProviderService_CancelPayment(t *testing.T) {
	type args struct {
		data *models.Payment
	}
	tests := []struct {
		name  string
		s     *PaymentProviderService
		args  args
		want  *models.Payment
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CancelPayment(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentProviderService.CancelPayment() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentProviderService.CancelPayment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentProviderService_GetStatus(t *testing.T) {
	type args struct {
		payment *models.Payment
	}
	tests := []struct {
		name  string
		s     *PaymentProviderService
		args  args
		want  *models.PaymentSessionStatus
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetStatus(tt.args.payment)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentProviderService.GetStatus() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentProviderService.GetStatus() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentProviderService_CapturePayment(t *testing.T) {
	type args struct {
		data *models.Payment
	}
	tests := []struct {
		name  string
		s     *PaymentProviderService
		args  args
		want  *models.Payment
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CapturePayment(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentProviderService.CapturePayment() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentProviderService.CapturePayment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentProviderService_RefundPayments(t *testing.T) {
	type args struct {
		data   []models.Payment
		amount float64
		reason models.RefundReason
		note   *string
	}
	tests := []struct {
		name  string
		s     *PaymentProviderService
		args  args
		want  *models.Refund
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RefundPayments(tt.args.data, tt.args.amount, tt.args.reason, tt.args.note)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentProviderService.RefundPayments() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentProviderService.RefundPayments() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentProviderService_RefundFromPayment(t *testing.T) {
	type args struct {
		payment *models.Payment
		amount  float64
		reason  models.RefundReason
		note    *string
	}
	tests := []struct {
		name  string
		s     *PaymentProviderService
		args  args
		want  *models.Refund
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RefundFromPayment(tt.args.payment, tt.args.amount, tt.args.reason, tt.args.note)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentProviderService.RefundFromPayment() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentProviderService.RefundFromPayment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentProviderService_RetrieveRefund(t *testing.T) {
	type args struct {
		id     uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *PaymentProviderService
		args  args
		want  *models.Refund
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveRefund(tt.args.id, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentProviderService.RetrieveRefund() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentProviderService.RetrieveRefund() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentProviderService_buildPaymentProcessorContext(t *testing.T) {
	type args struct {
		data *types.PaymentSessionInput
	}
	tests := []struct {
		name string
		s    *PaymentProviderService
		args args
		want *interfaces.PaymentProcessorContext
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.buildPaymentProcessorContext(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentProviderService.buildPaymentProcessorContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentProviderService_SaveSession(t *testing.T) {
	type args struct {
		id         uuid.UUID
		providerId uuid.UUID
		data       *models.PaymentSession
	}
	tests := []struct {
		name  string
		s     *PaymentProviderService
		args  args
		want  *models.PaymentSession
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.SaveSession(tt.args.id, tt.args.providerId, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentProviderService.SaveSession() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentProviderService.SaveSession() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentProviderService_processUpdateRequestsData(t *testing.T) {
	type args struct {
		data            *models.Customer
		paymentResponse *interfaces.PaymentProcessorSessionResponse
	}
	tests := []struct {
		name string
		s    *PaymentProviderService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.processUpdateRequestsData(tt.args.data, tt.args.paymentResponse); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentProviderService.processUpdateRequestsData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentProviderService_throwFromPaymentProcessorError(t *testing.T) {
	type args struct {
		errObj *interfaces.PaymentProcessorError
	}
	tests := []struct {
		name string
		s    *PaymentProviderService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.throwFromPaymentProcessorError(tt.args.errObj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentProviderService.throwFromPaymentProcessorError() = %v, want %v", got, tt.want)
			}
		})
	}
}

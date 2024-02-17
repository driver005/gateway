package services

import (
	"context"
	"reflect"
	"testing"
)

func TestNewSystemProviderService(t *testing.T) {
	tests := []struct {
		name string
		want *SystemProviderService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSystemProviderService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSystemProviderService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemProviderService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *SystemProviderService
		args args
		want *SystemProviderService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SystemProviderService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemProviderService_CreatePayment(t *testing.T) {
	tests := []struct {
		name string
		s    *SystemProviderService
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.CreatePayment(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SystemProviderService.CreatePayment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemProviderService_GetStatus(t *testing.T) {
	tests := []struct {
		name string
		s    *SystemProviderService
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetStatus(); got != tt.want {
				t.Errorf("SystemProviderService.GetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemProviderService_GetPaymentData(t *testing.T) {
	tests := []struct {
		name string
		s    *SystemProviderService
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetPaymentData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SystemProviderService.GetPaymentData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemProviderService_AuthorizePayment(t *testing.T) {
	tests := []struct {
		name string
		s    *SystemProviderService
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.AuthorizePayment(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SystemProviderService.AuthorizePayment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemProviderService_UpdatePaymentData(t *testing.T) {
	tests := []struct {
		name string
		s    *SystemProviderService
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UpdatePaymentData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SystemProviderService.UpdatePaymentData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemProviderService_UpdatePayment(t *testing.T) {
	tests := []struct {
		name string
		s    *SystemProviderService
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UpdatePayment(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SystemProviderService.UpdatePayment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemProviderService_DeletePayment(t *testing.T) {
	tests := []struct {
		name string
		s    *SystemProviderService
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.DeletePayment(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SystemProviderService.DeletePayment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemProviderService_CapturePayment(t *testing.T) {
	tests := []struct {
		name string
		s    *SystemProviderService
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.CapturePayment(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SystemProviderService.CapturePayment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemProviderService_RefundPayment(t *testing.T) {
	tests := []struct {
		name string
		s    *SystemProviderService
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.RefundPayment(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SystemProviderService.RefundPayment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemProviderService_CancelPayment(t *testing.T) {
	tests := []struct {
		name string
		s    *SystemProviderService
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.CancelPayment(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SystemProviderService.CancelPayment() = %v, want %v", got, tt.want)
			}
		})
	}
}

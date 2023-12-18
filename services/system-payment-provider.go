package services

import (
	"context"
)

type SystemProviderService struct {
	ctx context.Context
}

func NewSystemProviderService(ctx context.Context) *SystemProviderService {
	return &SystemProviderService{
		ctx,
	}
}

func (s *SystemProviderService) CreatePayment() map[string]interface{} {
	return make(map[string]interface{})
}

func (s *SystemProviderService) GetStatus() string {
	return "authorized"
}

func (s *SystemProviderService) GetPaymentData() map[string]interface{} {
	return make(map[string]interface{})
}

func (s *SystemProviderService) AuthorizePayment() map[string]interface{} {
	return map[string]interface{}{
		"data":   make(map[string]interface{}),
		"status": "authorized",
	}
}

func (s *SystemProviderService) UpdatePaymentData() map[string]interface{} {
	return make(map[string]interface{})
}

func (s *SystemProviderService) UpdatePayment() map[string]interface{} {
	return make(map[string]interface{})
}

func (s *SystemProviderService) DeletePayment() map[string]interface{} {
	return make(map[string]interface{})
}

func (s *SystemProviderService) CapturePayment() map[string]interface{} {
	return make(map[string]interface{})
}

func (s *SystemProviderService) RefundPayment() map[string]interface{} {
	return make(map[string]interface{})
}

func (s *SystemProviderService) CancelPayment() map[string]interface{} {
	return make(map[string]interface{})
}

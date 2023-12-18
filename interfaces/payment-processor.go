package interfaces

import (
	"github.com/driver005/gateway/models"
)

type PaymentProcessorContext struct {
	BillingAddress     *models.Address
	Email              string
	CurrencyCode       string
	Amount             float64
	ResourceID         string
	Customer           *models.Customer
	Context            map[string]interface{}
	PaymentSessionData map[string]interface{}
}

type PaymentProcessorSessionResponse struct {
	UpdateRequests map[string]interface{}
	SessionData    map[string]interface{}
}

type PaymentProcessorError struct {
	Error  string
	Code   string
	Detail interface{}
}

type PaymentProcessor interface {
	GetIdentifier() string
	InitiatePayment(context PaymentProcessorContext) (*PaymentProcessorError, *PaymentProcessorSessionResponse)
	UpdatePayment(context PaymentProcessorContext) (*PaymentProcessorError, *PaymentProcessorSessionResponse, error)
	RefundPayment(paymentSessionData map[string]interface{}, refundAmount float64) (*PaymentProcessorError, map[string]interface{})
	AuthorizePayment(paymentSessionData map[string]interface{}, context map[string]interface{}) (*PaymentProcessorError, *models.PaymentSessionStatus, map[string]interface{})
	CapturePayment(paymentSessionData map[string]interface{}) (*PaymentProcessorError, map[string]interface{})
	DeletePayment(paymentSessionData map[string]interface{}) (*PaymentProcessorError, map[string]interface{})
	RetrievePayment(paymentSessionData map[string]interface{}) (*PaymentProcessorError, map[string]interface{})
	CancelPayment(paymentSessionData map[string]interface{}) (*PaymentProcessorError, map[string]interface{})
	GetPaymentStatus(paymentSessionData map[string]interface{}) (*models.PaymentSessionStatus, error)
	UpdatePaymentData(sessionID string, data map[string]interface{}) (*PaymentProcessorError, map[string]interface{})
}

func IsPaymentProcessor(obj interface{}) bool {
	_, ok := obj.(*PaymentProcessor)
	return ok
}

func IsPaymentProcessorError(obj interface{}) bool {
	err, ok := obj.(*PaymentProcessorError)
	return ok && (err.Error != "" || err.Code != "" || err.Detail != nil)
}

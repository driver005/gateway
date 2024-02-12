package interfaces

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/google/uuid"
)

type PaymentProcessorContext struct {
	Id                 uuid.UUID
	BillingAddress     *models.Address
	Email              string
	CurrencyCode       string
	Amount             float64
	ResourceId         uuid.UUID
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

type IPaymentProcessor interface {
	GetIdentifier() string
	InitiatePayment(context *PaymentProcessorContext) (*PaymentProcessorSessionResponse, *PaymentProcessorError)
	UpdatePayment(context *PaymentProcessorContext) (*PaymentProcessorSessionResponse, *PaymentProcessorError)
	RefundPayment(paymentSessionData map[string]interface{}, refundAmount float64) (map[string]interface{}, *PaymentProcessorError)
	AuthorizePayment(paymentSessionData map[string]interface{}, context map[string]interface{}) (*models.PaymentSessionStatus, map[string]interface{}, *PaymentProcessorError)
	CapturePayment(paymentSessionData map[string]interface{}) (map[string]interface{}, *PaymentProcessorError)
	DeletePayment(paymentSessionData map[string]interface{}) (map[string]interface{}, *PaymentProcessorError)
	RetrievePayment(paymentSessionData map[string]interface{}) (map[string]interface{}, *PaymentProcessorError)
	CancelPayment(paymentSessionData map[string]interface{}) (map[string]interface{}, *PaymentProcessorError)
	GetPaymentStatus(paymentSessionData map[string]interface{}) (*models.PaymentSessionStatus, *PaymentProcessorError)
	UpdatePaymentData(sessionId uuid.UUID, data map[string]interface{}) (map[string]interface{}, *PaymentProcessorError)

	RetrieveSavedMethods(customer *models.Customer) []types.PaymentMethod
}

func IsPaymentProcessor(obj interface{}) bool {
	_, ok := obj.(*IPaymentProcessor)
	return ok
}

func IsPaymentProcessorError(obj interface{}) bool {
	err, ok := obj.(*PaymentProcessorError)
	return ok && (err.Error != "" || err.Code != "" || err.Detail != nil)
}

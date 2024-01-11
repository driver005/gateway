package services

import (
	"context"
	"fmt"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/sarulabs/di"
)

type PaymentProviderService struct {
	ctx       context.Context
	container di.Container
	r         Registry
}

func NewPaymentProviderService(
	container di.Container,
	r Registry,
) *PaymentProviderService {
	return &PaymentProviderService{
		context.Background(),
		container,
		r,
	}
}

func (s *PaymentProviderService) SetContext(context context.Context) *PaymentProviderService {
	s.ctx = context
	return s
}

func (s *PaymentProviderService) RegisterInstalledProviders(providers uuid.UUIDs) *utils.ApplictaionError {
	if err := s.r.PaymentProviderRepository().Update(s.ctx, &models.PaymentProvider{IsInstalled: false}); err != nil {
		return err
	}

	for _, p := range providers {
		var model *models.PaymentProvider
		model.IsInstalled = true
		model.Id = p

		if err := s.r.PaymentProviderRepository().Save(s.ctx, model); err != nil {
			return err
		}
	}

	return nil
}

func (s *PaymentProviderService) RetrieveProvider(providerId uuid.UUID) (interfaces.IPaymentProcessor, *utils.ApplictaionError) {
	var provider interfaces.IPaymentProcessor
	name := "system"
	if providerId != uuid.Nil {
		name = "pp_" + providerId.String()
	}
	objectInterface, err := s.container.SafeGet(name)
	if err != nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			fmt.Sprintf("Could not find a notification provider with id: %s.", providerId),
			"500",
			nil,
		)
	}
	provider, _ = objectInterface.(interfaces.IPaymentProcessor)

	return provider, nil
}

func (s *PaymentProviderService) List() ([]models.PaymentProvider, *utils.ApplictaionError) {
	var paymentProviders []models.PaymentProvider
	if err := s.r.PaymentProviderRepository().Find(s.ctx, paymentProviders, sql.Query{}); err != nil {
		return nil, err
	}
	return paymentProviders, nil
}

func (s *PaymentProviderService) RetrievePayment(id uuid.UUID, relations []string) (*models.Payment, *utils.ApplictaionError) {
	return s.r.PaymentService().SetContext(s.ctx).Retrieve(id, sql.Options{Relations: relations})
}

func (s *PaymentProviderService) ListPayments(selector models.Payment, config sql.Options) ([]models.Payment, *utils.ApplictaionError) {
	return s.r.PaymentService().SetContext(s.ctx).List(selector, config)
}

func (s *PaymentProviderService) RetrieveSession(id uuid.UUID, relations []string) (*models.PaymentSession, *utils.ApplictaionError) {
	if id == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"id" must be defined`,
			"500",
			nil,
		)
	}
	var res *models.PaymentSession
	query := sql.BuildQuery(models.PaymentSession{Model: core.Model{Id: id}}, sql.Options{Relations: relations})
	if err := s.r.PaymentSessionRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PaymentProviderService) CreateSession(providerId uuid.UUID, session *types.PaymentSessionInput) (*models.PaymentSession, *utils.ApplictaionError) {
	if session != nil {
		providerId = session.ProviderId
	}

	provider, err := s.RetrieveProvider(providerId)
	if err != nil {
		return nil, err
	}
	context := s.buildPaymentProcessorContext(session)
	if context.CurrencyCode == "" || context.Amount == 0 {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"`currency_code` and `amount` are required to Create payment session.",
			"500",
			nil,
		)
	}

	paymnet, fail := provider.InitiatePayment(context)
	if fail != nil {
		return nil, s.throwFromPaymentProcessorError(fail)
	}

	err = s.processUpdateRequestsData(&models.Customer{Model: core.Model{Id: context.Customer.Id}}, paymnet)
	if err != nil {
		return nil, err
	}
	return s.SaveSession(session.PaymentSessionId, providerId, &models.PaymentSession{
		CartId:      uuid.NullUUID{UUID: context.Id},
		Data:        paymnet.SessionData,
		Status:      models.PaymentSessionStatusPending,
		IsInitiated: true,
		Amount:      context.Amount,
	})
}

func (s *PaymentProviderService) RefreshSession(paymentSession *models.PaymentSession, sessionInput *types.PaymentSessionInput) (*models.PaymentSession, *utils.ApplictaionError) {
	session, err := s.RetrieveSession(paymentSession.Id, []string{})
	if err != nil {
		return nil, err
	}
	provider, err := s.RetrieveProvider(paymentSession.ProviderId.UUID)
	if err != nil {
		return nil, err
	}

	_, fail := provider.DeletePayment(session.Data)
	if fail != nil {
		return nil, s.throwFromPaymentProcessorError(fail)
	}

	if err := s.r.PaymentSessionRepository().Remove(s.ctx, session); err != nil {
		return nil, err
	}
	return s.CreateSession(uuid.Nil, sessionInput)
}

func (s *PaymentProviderService) UpdateSession(paymentSession *models.PaymentSession, sessionInput *types.PaymentSessionInput) (*models.PaymentSession, *utils.ApplictaionError) {
	provider, err := s.RetrieveProvider(paymentSession.ProviderId.UUID)
	if err != nil {
		return nil, err
	}
	context := s.buildPaymentProcessorContext(sessionInput)

	context.PaymentSessionData = paymentSession.Data

	paymnet, fail := provider.UpdatePayment(context)
	if fail != nil {
		return nil, s.throwFromPaymentProcessorError(fail)
	}

	if paymnet.SessionData != nil {
		return s.RetrieveSession(paymentSession.Id, []string{})
	}

	err = s.processUpdateRequestsData(&models.Customer{Model: core.Model{Id: context.Customer.Id}}, paymnet)
	if err != nil {
		return nil, err
	}
	return s.SaveSession(paymentSession.Id, paymentSession.ProviderId.UUID, &models.PaymentSession{
		Data:        paymnet.SessionData,
		IsInitiated: true,
		Amount:      context.Amount,
	})
}

func (s *PaymentProviderService) DeleteSession(paymentSession *models.PaymentSession) *utils.ApplictaionError {
	session, err := s.RetrieveSession(paymentSession.Id, []string{})
	if err != nil {
		return err
	}
	provider, err := s.RetrieveProvider(paymentSession.ProviderId.UUID)
	if err != nil {
		return err
	}
	_, fail := provider.DeletePayment(paymentSession.Data)
	if fail != nil {
		return s.throwFromPaymentProcessorError(fail)
	}
	if err := s.r.PaymentSessionRepository().Remove(s.ctx, session); err != nil {
		return err
	}
	return nil
}

func (s *PaymentProviderService) CreatePayment(data *models.Payment) (*models.Payment, *utils.ApplictaionError) {
	provider, err := s.RetrieveProvider(data.ProviderId.UUID)
	if err != nil {
		return nil, err
	}
	payment, fail := provider.RetrievePayment(data.Data)
	if fail != nil {
		return nil, s.throwFromPaymentProcessorError(fail)
	}

	data.Data = payment

	if err := s.r.PaymentRepository().Save(s.ctx, data); err != nil {
		return nil, err
	}
	return data, nil
}

func (s *PaymentProviderService) UpdatePayment(id uuid.UUID, Update *models.Payment) (*models.Payment, *utils.ApplictaionError) {
	return s.r.PaymentService().SetContext(s.ctx).Update(id, Update)
}

func (s *PaymentProviderService) AuthorizePayment(paymentSession *models.PaymentSession, context map[string]interface{}) (*models.PaymentSession, *utils.ApplictaionError) {
	session, err := s.RetrieveSession(paymentSession.Id, []string{})
	if err != nil {
		return nil, err
	}
	provider, err := s.RetrieveProvider(paymentSession.ProviderId.UUID)
	if err != nil {
		return nil, err
	}

	status, payment, fail := provider.AuthorizePayment(paymentSession.Data, context)
	if err != nil {
		return nil, s.throwFromPaymentProcessorError(fail)
	}
	session.Data = payment
	session.Status = *status

	if session.Status == models.PaymentSessionStatusAuthorized {
		now := time.Now()
		session.PaymentAuthorizedAt = &now
	}
	if err := s.r.PaymentSessionRepository().Save(s.ctx, session); err != nil {
		return nil, err
	}
	return paymentSession, nil
}

func (s *PaymentProviderService) UpdateSessionData(paymentSession *models.PaymentSession, data map[string]interface{}) (*models.PaymentSession, *utils.ApplictaionError) {
	session, err := s.RetrieveSession(paymentSession.Id, []string{})
	if err != nil {
		return nil, err
	}
	provider, err := s.RetrieveProvider(paymentSession.ProviderId.UUID)
	if err != nil {
		return nil, err
	}
	paymnet, fail := provider.UpdatePaymentData(paymentSession.Id, data)
	if fail != nil {
		return nil, s.throwFromPaymentProcessorError(fail)
	}

	paymentSession.Id = session.Id
	paymentSession.Data = paymnet
	paymentSession.Status = session.Status

	if err := s.r.PaymentSessionRepository().Update(s.ctx, paymentSession); err != nil {
		return nil, err
	}
	return paymentSession, nil
}

func (s *PaymentProviderService) CancelPayment(data *models.Payment) (*models.Payment, *utils.ApplictaionError) {
	payment, err := s.RetrievePayment(data.Id, []string{})
	if err != nil {
		return payment, err
	}
	provider, err := s.RetrieveProvider(payment.ProviderId.UUID)
	if err != nil {
		return nil, err
	}
	pay, fail := provider.CancelPayment(payment.Data)
	if fail != nil {
		return nil, s.throwFromPaymentProcessorError(fail)
	}
	now := time.Now()
	payment.CanceledAt = &now
	payment.Data = pay

	if err := s.r.PaymentRepository().Update(s.ctx, payment); err != nil {
		return nil, err
	}
	return payment, nil
}

func (s *PaymentProviderService) GetStatus(payment *models.Payment) (*models.PaymentSessionStatus, *utils.ApplictaionError) {
	provider, err := s.RetrieveProvider(payment.ProviderId.UUID)
	if err != nil {
		return nil, err
	}

	status, fail := provider.GetPaymentStatus(payment.Data)
	if fail != nil {
		return nil, s.throwFromPaymentProcessorError(fail)
	}
	return status, err
}

func (s *PaymentProviderService) CapturePayment(data *models.Payment) (*models.Payment, *utils.ApplictaionError) {
	payment, err := s.RetrievePayment(data.Id, []string{})
	if err != nil {
		return payment, err
	}
	provider, err := s.RetrieveProvider(payment.ProviderId.UUID)
	if err != nil {
		return nil, err
	}
	res, fail := provider.CapturePayment(payment.Data)
	if fail != nil {
		return nil, s.throwFromPaymentProcessorError(fail)
	}

	payment.Data = res
	now := time.Now()
	payment.CapturedAt = &now

	if err := s.r.PaymentRepository().Update(s.ctx, payment); err != nil {
		return nil, err
	}
	return nil, err
}

func (s *PaymentProviderService) RefundPayments(data []models.Payment, amount float64, reason string, note *string) (*models.Refund, *utils.ApplictaionError) {
	var ids uuid.UUIDs
	for _, d := range data {
		ids = append(ids, d.Id)
	}
	payments, err := s.ListPayments(models.Payment{}, sql.Options{Specification: []sql.Specification{sql.In("id", ids)}})
	if err != nil {
		return nil, err
	}

	var orderId uuid.NullUUID
	var refundable float64
	var refunds []models.Payment
	for _, payment := range payments {
		orderId = payment.OrderId
		if payment.CapturedAt != nil {
			refundable += payment.Amount - payment.AmountRefunded
		}

		if payment.Amount-payment.AmountRefunded > 0 {
			refunds = append(refunds, payment)
		}
	}
	if refundable < amount {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Refund amount is greater than the refundable amount",
			"500",
			nil,
		)
	}

	var used uuid.UUIDs
	balance := amount

	for _, refund := range refunds {
		currentRefundable := refund.Amount - refund.AmountRefunded
		refundAmount := min(currentRefundable, balance)
		provider, err := s.RetrieveProvider(refund.ProviderId.UUID)
		if err != nil {
			return nil, err
		}
		res, fail := provider.RefundPayment(refund.Data, refundAmount)
		if fail != nil {
			return nil, s.throwFromPaymentProcessorError(fail)
		}
		refund.Data = res
		refund.AmountRefunded += refundAmount
		if err := s.r.PaymentRepository().Update(s.ctx, &refund); err != nil {
			return nil, err
		}
		balance -= refundAmount
		used = append(used, refund.Id)
		if balance > 0 {
			refunds = append(refunds, refund)
		} else {
			continue
		}
	}

	res := &models.Refund{
		OrderId: orderId,
		Amount:  amount,
		Reason:  reason,
		Note:    *note,
	}

	if err := s.r.RefundRepository().Update(s.ctx, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PaymentProviderService) RefundFromPayment(payment *models.Payment, amount float64, reason string, note *string) (*models.Refund, *utils.ApplictaionError) {
	return s.RefundPayments([]models.Payment{*payment}, amount, reason, note)
}

func (s *PaymentProviderService) RetrieveRefund(id uuid.UUID, config sql.Options) (*models.Refund, *utils.ApplictaionError) {
	if id == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"id" must be defined`,
			"500",
			nil,
		)
	}
	var res *models.Refund
	query := sql.BuildQuery(models.Refund{Model: core.Model{Id: id}}, config)
	if err := s.r.RefundRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PaymentProviderService) buildPaymentProcessorContext(data *types.PaymentSessionInput) *interfaces.PaymentProcessorContext {
	var processor *interfaces.PaymentProcessorContext

	processor.Id = data.Cart.Id
	processor.Amount = data.Amount
	processor.BillingAddress = data.Customer.BillingAddress
	processor.Context = data.Context
	processor.Customer = data.Customer
	processor.Email = data.Customer.Email
	processor.PaymentSessionData = data.PaymentSessionData
	processor.ResourceId = data.ResourceId

	return processor
}

func (s *PaymentProviderService) SaveSession(id uuid.UUID, providerId uuid.UUID, data *models.PaymentSession) (*models.PaymentSession, *utils.ApplictaionError) {
	if id != uuid.Nil {
		session, err := s.RetrieveSession(id, []string{})
		if err != nil {
			return nil, err
		}

		session.Data = data.Data
		session.Status = data.Status
		session.Amount = data.Amount
		session.IsInitiated = data.IsInitiated
		session.IsSelected = data.IsSelected

		if err := s.r.PaymentSessionRepository().Save(s.ctx, session); err != nil {
			return nil, err
		}
		return session, nil
	}

	data.ProviderId = uuid.NullUUID{UUID: providerId}

	if err := s.r.PaymentSessionRepository().Save(s.ctx, data); err != nil {
		return nil, err
	}
	return data, nil
}

func (s *PaymentProviderService) processUpdateRequestsData(data *models.Customer, paymentResponse *interfaces.PaymentProcessorSessionResponse) *utils.ApplictaionError {
	if paymentResponse.UpdateRequests == nil {
		return nil
	}

	if paymentResponse.UpdateRequests["customer_metadata"] != nil && data.Id != uuid.Nil {
		_, err := s.r.CustomerService().SetContext(s.ctx).Update(data.Id, &models.Customer{Model: core.Model{Metadata: paymentResponse.UpdateRequests["customer_metadata"].(map[string]interface{})}})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *PaymentProviderService) throwFromPaymentProcessorError(errObj *interfaces.PaymentProcessorError) *utils.ApplictaionError {
	return utils.NewApplictaionError(
		utils.INVALID_DATA,
		fmt.Sprintf("%s %s", errObj.Error, errObj.Detail),
		errObj.Code,
		nil,
	)
}

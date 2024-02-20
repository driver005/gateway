package services

import (
	"context"
	"fmt"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type PaymentService struct {
	ctx context.Context
	r   Registry
}

func NewPaymentService(
	r Registry,
) *PaymentService {
	return &PaymentService{
		context.Background(),
		r,
	}
}

func (s *PaymentService) SetContext(context context.Context) *PaymentService {
	s.ctx = context
	return s
}

func (s *PaymentService) Retrieve(id uuid.UUID, config *sql.Options) (*models.Payment, *utils.ApplictaionError) {
	if id == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"id" must be defined`,
			nil,
		)
	}
	var res *models.Payment = &models.Payment{}
	query := sql.BuildQuery(models.Payment{Model: core.Model{Id: id}}, config)
	if err := s.r.PaymentRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PaymentService) List(selector models.Payment, config *sql.Options) ([]models.Payment, *utils.ApplictaionError) {
	var res []models.Payment
	query := sql.BuildQuery(selector, config)
	if err := s.r.PaymentRepository().Find(s.ctx, &res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PaymentService) Create(data *types.CreatePaymentInput) (*models.Payment, *utils.ApplictaionError) {
	model := &models.Payment{
		CartId:       uuid.NullUUID{UUID: data.CartId},
		Amount:       data.Amount,
		CurrencyCode: data.CurrencyCode,
		ProviderId:   uuid.NullUUID{UUID: data.ProviderId},
	}
	if err := s.r.PaymentRepository().Save(s.ctx, model); err != nil {
		return nil, err
	}
	// err = s.EventBusService.Emit(PaymentService.Events.CREATED, saved)
	return model, nil
}

func (s *PaymentService) Update(id uuid.UUID, data *types.UpdatePaymentInput) (*models.Payment, *utils.ApplictaionError) {
	payment, err := s.Retrieve(id, &sql.Options{})
	if err != nil {
		return nil, err
	}

	if data.OrderId != uuid.Nil {
		payment.OrderId = uuid.NullUUID{UUID: data.OrderId}
	}

	if data.SwapId != uuid.Nil {
		payment.SwapId = uuid.NullUUID{UUID: data.SwapId}
	}

	if err := s.r.PaymentRepository().Save(s.ctx, payment); err != nil {
		return nil, err
	}
	// err = s.EventBusService.Emit(PaymentService.Events.UPDATED, updated)
	return payment, nil
}

func (s *PaymentService) Capture(id uuid.UUID, payment *models.Payment) (*models.Payment, *utils.ApplictaionError) {
	if id != uuid.Nil {
		p, err := s.Retrieve(id, &sql.Options{})
		if err != nil {
			return nil, err
		}
		payment = p
	}

	if payment.CapturedAt != nil {
		return payment, nil
	}

	capturedPayment, err := s.r.PaymentProviderService().SetContext(s.ctx).CapturePayment(payment)
	if err != nil {
		// err = s.EventBusService.Emit(PaymentService.Events.PAYMENT_CAPTURE_FAILED, map[string]interface{}{
		// 	"payment":                 payment,
		// 	"*utils.ApplictaionError": captureError,
		// })
		// if err != nil {
		// 	return nil, err
		// }
		return nil, err
	}
	// err = s.EventBusService.Emit(PaymentService.Events.PAYMENT_CAPTURED, capturedPayment)
	return capturedPayment, nil
}

func (s *PaymentService) Refund(id uuid.UUID, payment *models.Payment, amount float64, reason models.RefundReason, note *string) (*models.Refund, *utils.ApplictaionError) {
	if id != uuid.Nil {
		p, err := s.Retrieve(id, &sql.Options{})
		if err != nil {
			return nil, err
		}
		payment = p
	}

	if payment.CapturedAt == nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`Payment `+payment.Id.String()+` is not captured`,
			nil,
		)
	}

	refundable := payment.Amount - payment.AmountRefunded
	if amount > refundable {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			fmt.Sprintf("Only %f can be refunded from models.Payment %v", refundable, payment.Id),
			nil,
		)
	}

	refund, err := s.r.PaymentProviderService().SetContext(s.ctx).RefundFromPayment(payment, amount, reason, note)
	if err != nil {
		// err = s.EventBusService.Emit(PaymentService.Events.REFUND_FAILED, map[string]interface{}{
		// 	"payment":                 payment,
		// 	"*utils.ApplictaionError": refundError,
		// })
		// if err != nil {
		// 	return nil, err
		// }
		return nil, err
	}
	// err = s.EventBusService.Emit(PaymentService.Events.REFUND_CREATED, refund)
	return refund, nil
}

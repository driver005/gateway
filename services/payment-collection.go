package services

import (
	"context"
	"fmt"
	"reflect"
	"slices"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type PaymentCollectionService struct {
	ctx context.Context
	r   Registry
}

func NewPaymentCollectionService(
	r Registry,
) *PaymentCollectionService {
	return &PaymentCollectionService{
		context.Background(),
		r,
	}
}

func (s *PaymentCollectionService) SetContext(context context.Context) *PaymentCollectionService {
	s.ctx = context
	return s
}

func (s *PaymentCollectionService) Retrieve(id uuid.UUID, config *sql.Options) (*models.PaymentCollection, *utils.ApplictaionError) {
	if id == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_FOUND,
			`"id" must be defined`,
			nil,
		)
	}

	var res *models.PaymentCollection
	query := sql.BuildQuery(models.OAuth{Model: core.Model{Id: id}}, config)

	if err := s.r.PaymentCollectionRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PaymentCollectionService) Create(data *types.CreatePaymentCollectionInput) (*models.PaymentCollection, *utils.ApplictaionError) {
	model := &models.PaymentCollection{
		Model: core.Model{
			Metadata: data.Metadata,
		},
		Status:       models.PaymentCollectionStatusNotPaid,
		RegionId:     uuid.NullUUID{UUID: data.RegionId},
		Type:         data.Type,
		CurrencyCode: data.CurrencyCode,
		Amount:       data.Amount,
		CreatedBy:    data.CreatedBy,
		Description:  data.Description,
	}

	if err := s.r.PaymentCollectionRepository().Save(s.ctx, model); err != nil {
		return nil, err
	}
	// eventBusService.emit(PaymentCollectionService.Events.CREATED, paymentCollection)
	return model, nil
}

func (s *PaymentCollectionService) Update(id uuid.UUID, Update *models.PaymentCollection) (*models.PaymentCollection, *utils.ApplictaionError) {
	paymentCollection, err := s.Retrieve(id, &sql.Options{})
	if err != nil {
		return nil, err
	}

	Update.Id = paymentCollection.Id

	if err := s.r.PaymentCollectionRepository().Save(s.ctx, Update); err != nil {
		return nil, err
	}
	// eventBusService.emit(PaymentCollectionService.Events.UPDATED, result)
	return Update, nil
}

func (s *PaymentCollectionService) Delete(id uuid.UUID) *utils.ApplictaionError {
	paymentCollection, err := s.Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}
	if paymentCollection == nil {
		return nil
	}
	if err := s.r.PaymentCollectionRepository().Remove(s.ctx, paymentCollection); err != nil {
		return err
	}
	// eventBusService.emit(PaymentCollectionService.Events.DELETED, paymentCollection)
	return nil
}

func (s *PaymentCollectionService) IsValidTotalAmount(total float64, sessionsInput []types.PaymentCollectionsSessionsBatchInput) bool {
	sum := 0.0
	for _, sess := range sessionsInput {
		sum += sess.Amount
	}
	return total == sum
}

func (s *PaymentCollectionService) SetPaymentSessionsBatch(id uuid.UUID, paymentCollection *models.PaymentCollection, sessionsInput []types.PaymentCollectionsSessionsBatchInput, customerId uuid.UUID) (*models.PaymentCollection, *utils.ApplictaionError) {
	if id != uuid.Nil {
		p, err := s.Retrieve(id, &sql.Options{
			Relations: []string{"region", "region.payment_providers", "payment_sessions"},
		})
		if err != nil {
			return nil, err
		}
		paymentCollection = p
	}

	if paymentCollection.Status != models.PaymentCollectionStatusNotPaid {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			fmt.Sprintf(`Cannot set payment sessions for a payment collection with status %v`, paymentCollection.Status),
			nil,
		)
	}

	payColRegionProviderMap := make(map[uuid.UUID]interface{})
	for _, provider := range paymentCollection.Region.PaymentProviders {
		payColRegionProviderMap[provider.Id] = provider
	}

	var sessions []types.PaymentCollectionsSessionsBatchInput

	for _, input := range sessionsInput {
		if payColRegionProviderMap[input.ProviderId] != nil {
			sessions = append(sessions, input)
		}
	}

	if !s.IsValidTotalAmount(paymentCollection.Amount, sessions) {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			fmt.Sprintf(`The sum of sessions is not equal to %f on Payment Collection`, paymentCollection.Amount),
			nil,
		)
	}

	var customer *models.Customer
	if customerId != uuid.Nil {
		c, err := s.r.CustomerService().SetContext(s.ctx).RetrieveById(customerId, &sql.Options{
			Selects: []string{"id", "email", "metadata"},
		})
		if err != nil {
			return nil, err
		}
		customer = c
	}

	payColSessionMap := make(map[uuid.UUID]models.PaymentSession)
	for _, session := range paymentCollection.PaymentSessions {
		payColSessionMap[session.Id] = session
	}

	var selectedSessionIds uuid.UUIDs
	var paymentSessions []models.PaymentSession
	for _, session := range sessions {
		pay, ok := payColSessionMap[session.SessionId]

		var existingSession *models.PaymentSession
		if ok {
			existingSession = &pay
		}

		inputData := &types.PaymentSessionInput{
			Cart: &models.Cart{
				Email:           customer.Email,
				Context:         core.JSONB{},
				ShippingMethods: []models.ShippingMethod{},
				ShippingAddress: &models.Address{},
				RegionId:        paymentCollection.RegionId,
				Total:           session.Amount,
			},
			ResourceId:   paymentCollection.Id,
			CurrencyCode: paymentCollection.CurrencyCode,
			Amount:       session.Amount,
			ProviderId:   session.ProviderId,
			Customer:     customer,
		}

		var paymentSession *models.PaymentSession
		if existingSession != nil {
			p, err := s.r.PaymentProviderService().SetContext(s.ctx).UpdateSession(existingSession, inputData)
			if err != nil {
				return nil, err
			}
			paymentSession = p
		} else {
			p, err := s.r.PaymentProviderService().SetContext(s.ctx).CreateSession(uuid.Nil, inputData)
			if err != nil {
				return nil, err
			}
			paymentSession = p
		}
		selectedSessionIds = append(selectedSessionIds, paymentSession.Id)
		paymentSessions = append(paymentSessions, *paymentSession)
	}

	if len(paymentCollection.PaymentSessions) > 0 {
		var removeSessions []models.PaymentSession
		for _, sess := range paymentCollection.PaymentSessions {
			if !slices.Contains(selectedSessionIds, sess.Id) {
				removeSessions = append(removeSessions, sess)
			}
		}
		if len(removeSessions) > 0 {
			// paymentCollectionRepository.Delete(removeSessions.map(func(sess PaymentSession) string {
			// 	return sess.id
			// }))
			for _, sess := range removeSessions {
				err := s.r.PaymentProviderService().SetContext(s.ctx).DeleteSession(&sess)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	paymentCollection.PaymentSessions = paymentSessions

	if err := s.r.PaymentCollectionRepository().Save(s.ctx, paymentCollection); err != nil {
		return nil, err
	}

	return paymentCollection, nil
}

func (s *PaymentCollectionService) SetPaymentSession(id uuid.UUID, sessionInput *types.PaymentCollectionsSessionsInput, customerId uuid.UUID) (*models.PaymentCollection, *utils.ApplictaionError) {
	paymentCollection, err := s.Retrieve(id, &sql.Options{
		Relations: []string{"region", "region.payment_providers", "payment_sessions"},
	})
	if err != nil {
		return nil, err
	}

	hasProvider := false
	for _, p := range paymentCollection.Region.PaymentProviders {
		if p.Id == sessionInput.ProviderId {
			hasProvider = true
			break
		}
	}

	if !hasProvider {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Payment provider not found",
			nil,
		)
	}

	var existingSession models.PaymentSession
	for _, sess := range paymentCollection.PaymentSessions {
		if sessionInput.ProviderId == sess.ProviderId.UUID {
			existingSession = sess
			break
		}
	}

	return s.SetPaymentSessionsBatch(uuid.Nil, paymentCollection, []types.PaymentCollectionsSessionsBatchInput{
		{
			ProviderId: sessionInput.ProviderId,
			Amount:     paymentCollection.Amount,
			SessionId:  existingSession.Id,
		},
	}, customerId)
}

func (s *PaymentCollectionService) RefreshPaymentSession(id uuid.UUID, sessionId uuid.UUID, customerId uuid.UUID) (*models.PaymentSession, *utils.ApplictaionError) {
	paymentCollection, err := s.r.PaymentCollectionRepository().GetPaymentCollectionIdBySessionId(sessionId, &sql.Options{
		Relations: []string{"region", "payment_sessions"},
	})
	if err != nil {
		return nil, err
	}

	if id != paymentCollection.Id {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`Payment Session `+sessionId.String()+` does not belong to Payment Collection `+id.String(),
			nil,
		)
	}

	var session models.PaymentSession
	for _, sess := range paymentCollection.PaymentSessions {
		if sessionId == sess.Id {
			session = sess
			break
		}
	}

	if reflect.DeepEqual(session, models.PaymentSession{}) {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`Session with id `+sessionId.String()+` was not found`,
			nil,
		)
	}

	var customer *models.Customer
	if customerId != uuid.Nil {
		c, err := s.r.CustomerService().SetContext(s.ctx).RetrieveById(customerId, &sql.Options{
			Selects: []string{"id", "email", "metadata"},
		})
		if err != nil {
			return nil, err
		}

		customer = c
	}

	inputData := &types.PaymentSessionInput{
		ResourceId:   paymentCollection.Id,
		CurrencyCode: paymentCollection.CurrencyCode,
		Amount:       session.Amount,
		ProviderId:   session.ProviderId.UUID,
		Customer:     customer,
	}

	sessionRefreshed, err := s.r.PaymentProviderService().SetContext(s.ctx).RefreshSession(&session, inputData)
	if err != nil {
		return nil, err
	}

	var paymentSessions []models.PaymentSession
	for _, sess := range paymentCollection.PaymentSessions {
		if sess.Id == sessionId {
			paymentSessions = append(paymentSessions, *sessionRefreshed)
		}
	}
	paymentCollection.PaymentSessions = paymentSessions

	if session.PaymentAuthorizedAt != nil && paymentCollection.AuthorizedAmount != 0.0 {
		paymentCollection.AuthorizedAmount -= session.Amount
	}

	if err := s.r.PaymentCollectionRepository().Save(s.ctx, paymentCollection); err != nil {
		return nil, err
	}

	return sessionRefreshed, nil
}

func (s *PaymentCollectionService) MarkAsAuthorized(id uuid.UUID) (*models.PaymentCollection, *utils.ApplictaionError) {
	paymentCollection, err := s.Retrieve(id, &sql.Options{})
	if err != nil {
		return nil, err
	}
	paymentCollection.Status = models.PaymentCollectionStatusAuthorized
	paymentCollection.AuthorizedAmount = paymentCollection.Amount
	if err := s.r.PaymentCollectionRepository().Save(s.ctx, paymentCollection); err != nil {
		return nil, err
	}
	// eventBusService.emit(PaymentCollectionService.Events.PAYMENT_AUTHORIZED, result)
	return paymentCollection, nil
}

func (s *PaymentCollectionService) AuthorizePaymentSessions(id uuid.UUID, sessionIds uuid.UUIDs, context map[string]interface{}) (*models.PaymentCollection, *utils.ApplictaionError) {
	paymentCollection, err := s.Retrieve(id, &sql.Options{
		Relations: []string{"payment_sessions", "payments"},
	})
	if err != nil {
		return nil, err
	}

	if paymentCollection.AuthorizedAmount == paymentCollection.Amount {
		return paymentCollection, nil
	}

	if paymentCollection.Amount <= 0 {
		paymentCollection.AuthorizedAmount = 0
		paymentCollection.Status = models.PaymentCollectionStatusAuthorized
		if err := s.r.PaymentCollectionRepository().Save(s.ctx, paymentCollection); err != nil {
			return nil, err
		}

		return paymentCollection, nil
	}

	if len(paymentCollection.PaymentSessions) == 0 {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"You cannot complete a Payment Collection without a payment session.",
			nil,
		)
	}

	authorizedAmount := 0.0
	for i := 0; i < len(paymentCollection.PaymentSessions); i++ {
		session := paymentCollection.PaymentSessions[i]
		if session.PaymentAuthorizedAt != nil {
			authorizedAmount += session.Amount
			continue
		}
		if !slices.Contains(sessionIds, session.Id) {
			continue
		}
		paymentSession, err := s.r.PaymentProviderService().SetContext(s.ctx).AuthorizePayment(&session, context)
		if err != nil {
			return nil, err
		}
		if paymentSession != nil {
			paymentCollection.PaymentSessions[i] = *paymentSession
		}
		if paymentSession.Status == models.PaymentSessionStatusAuthorized {
			authorizedAmount += session.Amount
			data, err := s.r.PaymentProviderService().SetContext(s.ctx).CreatePayment(&types.CreatePaymentInput{
				Amount:         session.Amount,
				CurrencyCode:   paymentCollection.CurrencyCode,
				ProviderId:     session.ProviderId.UUID,
				ResourceId:     paymentCollection.Id,
				PaymentSession: paymentSession,
			})
			if err != nil {
				return nil, err
			}
			paymentCollection.Payments = append(paymentCollection.Payments, *data)
		}
	}

	if authorizedAmount == 0 {
		paymentCollection.Status = models.PaymentCollectionStatusAwaiting
	} else if authorizedAmount < paymentCollection.Amount {
		paymentCollection.Status = models.PaymentCollectionStatusPartiallyAuthorized
	} else if authorizedAmount == paymentCollection.Amount {
		paymentCollection.Status = models.PaymentCollectionStatusAuthorized
	}

	paymentCollection.AuthorizedAmount = authorizedAmount
	if err := s.r.PaymentCollectionRepository().Save(s.ctx, paymentCollection); err != nil {
		return nil, err
	}
	// eventBusService.emit(PaymentCollectionService.Events.PAYMENT_AUTHORIZED, payColCopy)
	return paymentCollection, nil
}

package services

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/sarulabs/di"
)

type NotificationService struct {
	ctx                 context.Context
	container           di.Container
	subscribers         map[string]uuid.UUIDs
	attachmentGenerator interface{}
	r                   Registry
}

func NewNotificationService(
	container di.Container,
	subscribers map[string]uuid.UUIDs,
	attachmentGenerator interface{},
	r Registry,
) *NotificationService {
	return &NotificationService{
		context.Background(),
		container,
		subscribers,
		attachmentGenerator,
		r,
	}
}

func (s *NotificationService) SetContext(context context.Context) *NotificationService {
	s.ctx = context
	return s
}

func (s *NotificationService) RegisterAttachmentGenerator(service interface{}) {
	s.attachmentGenerator = service
}

func (s *NotificationService) RegisterInstalledProviders(providers uuid.UUIDs) *utils.ApplictaionError {
	if err := s.r.NotificationProviderRepository().Update(s.ctx, &models.NotificationProvider{IsInstalled: false}); err != nil {
		return err
	}

	for _, p := range providers {
		var model *models.NotificationProvider = &models.NotificationProvider{}
		model.IsInstalled = true
		model.Id = uuid.NullUUID{UUID: p}

		if err := s.r.NotificationProviderRepository().Save(s.ctx, model); err != nil {
			return err
		}
	}

	return nil
}

func (s *NotificationService) List(selector *types.FilterableNotification, config *sql.Options) ([]models.Notification, *utils.ApplictaionError) {
	notifications, _, err := s.ListAndCount(selector, config)
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (s *NotificationService) ListAndCount(selector *types.FilterableNotification, config *sql.Options) ([]models.Notification, *int64, *utils.ApplictaionError) {
	if reflect.DeepEqual(config, &sql.Options{}) {
		config.Skip = 0
		config.Take = 50
		config.Order = "created_at DESC"
	}

	var res []models.Notification

	query := sql.BuildQuery(selector, config)

	count, err := s.r.NotificationRepository().FindAndCount(s.ctx, &res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

func (s *NotificationService) Retrieve(id uuid.UUID, config *sql.Options) (*models.Notification, *utils.ApplictaionError) {
	var res *models.Notification = &models.Notification{}
	query := sql.BuildQuery(&models.Notification{BaseModel: core.BaseModel{Id: id}}, config)
	if err := s.r.NotificationRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *NotificationService) Subscribe(eventName string, providerId uuid.UUID) *utils.ApplictaionError {
	if providerId == uuid.Nil {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			"providerId must be a string",
			nil,
		)
	}
	if s.subscribers[eventName] != nil {
		s.subscribers[eventName] = append(s.subscribers[eventName], providerId)
	} else {
		s.subscribers[eventName] = uuid.UUIDs{providerId}
	}
	return nil
}

func (s *NotificationService) RetrieveProvider(id uuid.UUID) (interfaces.INotificationService, *utils.ApplictaionError) {
	var provider interfaces.INotificationService
	objectInterface, err := s.container.SafeGet("noti_" + id.String())
	if err != nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			fmt.Sprintf("Could not find a notification provider with id: %s.", id),
			nil,
		)
	}
	provider, _ = objectInterface.(interfaces.INotificationService)

	return provider, nil
}

func (s *NotificationService) HandleEvent(eventName string, data map[string]interface{}) *utils.ApplictaionError {
	subs := s.subscribers[eventName]
	if subs == nil {
		return nil
	}
	if data["no_notification"] == true {
		return nil
	}
	for _, providerId := range subs {
		_, err := s.Send(eventName, data, providerId)
		if err != nil {
			return err
			// s.logger_.Log(err)
			// s.logger_.Warn(fmt.Sprintf("An *utils.ApplictaionError occured while %s was processing a notification for %s: %s", providerId, eventName, err.Error()))
		}
	}
	return nil
}

func (s *NotificationService) Send(event string, eventData map[string]interface{}, providerId uuid.UUID) (*models.Notification, *utils.ApplictaionError) {
	provider, err := s.RetrieveProvider(providerId)
	if err != nil {
		return nil, err
	}
	result, er := provider.SendNotification(event, eventData, s.attachmentGenerator)
	if er != nil {
		return nil, err
	}
	if result == nil {
		return nil, err
	}
	resourceType := event[:strings.Index(event, ".")]
	resourceId := eventData["id"].(uuid.UUID)
	customerId := eventData["customer_id"].(uuid.UUID)
	res := &models.Notification{
		ResourceType: resourceType,
		ResourceId:   uuid.NullUUID{UUID: resourceId},
		CustomerId:   uuid.NullUUID{UUID: customerId},
		To:           result.To,
		Data:         result.Data,
		EventName:    event,
		ProviderId:   uuid.NullUUID{UUID: providerId},
	}

	if err := s.r.NotificationRepository().Save(s.ctx, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *NotificationService) Resend(id uuid.UUID, config *sql.Options) (*models.Notification, *utils.ApplictaionError) {
	notification, err := s.Retrieve(id, config)
	if err != nil {
		return nil, err
	}
	provider, err := s.RetrieveProvider(notification.ProviderId.UUID)
	if err != nil {
		return nil, err
	}
	result, er := provider.ResendNotification(notification, config, s.attachmentGenerator)
	if er != nil {
		return nil, err
	}

	notification.To = result.To
	notification.Data = result.Data
	notification.ParentId = uuid.NullUUID{UUID: id}

	if err := s.r.NotificationRepository().Save(s.ctx, notification); err != nil {
		return nil, err
	}
	return notification, err
}

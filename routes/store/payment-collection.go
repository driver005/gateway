package store

import (
	"fmt"

	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type PaymentCollection struct {
	r Registry
}

func NewPaymentCollection(r Registry) *PaymentCollection {
	m := PaymentCollection{r: r}
	return &m
}

func (m *PaymentCollection) SetRoutes(router fiber.Router) {
	route := router.Group("/payment-collections")
	route.Get("/:id", m.Get)

	route.Post("/:id/sessions/batch", m.PaymentSessionManageBatch)
	route.Post("/:id/sessions/batch/authorize", m.PaymentSessionAuthorizeBatch)
	route.Post("/:id/sessions", m.PaymentSessionManage)
	route.Post("/:id/sessions/:session_id", m.PaymentSessionRefresh)
	route.Post("/:id/sessions/:session_id/authorize", m.PaymentSessionAuthorize)
}

func (m *PaymentCollection) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	result, err := m.r.PaymentCollectionService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *PaymentCollection) PaymentSessionAuthorizeBatch(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.PaymentCollectionsAuthorizeBatch](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.PaymentCollectionService().SetContext(context.Context()).AuthorizePaymentSessions(id, model.SessionIds, context.Locals("request_context").(map[string]interface{}))
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *PaymentCollection) PaymentSessionAuthorize(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	sessionId, err := api.BindDelete(context, "session_id")
	if err != nil {
		return err
	}

	paymentCollection, err := m.r.PaymentCollectionService().SetContext(context.Context()).AuthorizePaymentSessions(id, uuid.UUIDs{sessionId}, context.Locals("request_context").(map[string]interface{}))
	if err != nil {
		return err
	}

	result, ok := lo.Find(paymentCollection.PaymentSessions, func(item models.PaymentSession) bool {
		return item.Id == sessionId
	})

	if !ok {
		return utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintln("Could not find Payment Session with id %s", id),
		)
	}

	if result.Status != models.PaymentSessionStatusAuthorized {
		return utils.NewApplictaionError(
			utils.PAYMENT_AUTHORIZATION_ERROR,
			fmt.Sprintln("Failed to authorize Payment Session id %s", id),
		)
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *PaymentCollection) PaymentSessionManageBatch(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.PaymentCollectionsSessionsBatch](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	customerId := context.Locals("customer_id").(uuid.UUID)

	result, err := m.r.PaymentCollectionService().SetContext(context.Context()).SetPaymentSessionsBatch(id, nil, model.Sessions, customerId)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *PaymentCollection) PaymentSessionManage(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.PaymentCollectionsSessionsInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	customerId := context.Locals("customer_id").(uuid.UUID)

	result, err := m.r.PaymentCollectionService().SetContext(context.Context()).SetPaymentSession(id, model, customerId)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *PaymentCollection) PaymentSessionRefresh(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	sessionId, err := api.BindDelete(context, "session_id")
	if err != nil {
		return err
	}

	customerId := context.Locals("customer_id").(uuid.UUID)

	result, err := m.r.PaymentCollectionService().SetContext(context.Context()).RefreshPaymentSession(id, sessionId, customerId)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

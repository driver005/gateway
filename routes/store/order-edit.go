package store

import (
	"fmt"

	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/samber/lo"
)

type OrderEdit struct {
	r Registry
}

func NewOrderEdit(r Registry) *OrderEdit {
	m := OrderEdit{r: r}
	return &m
}

func (m *OrderEdit) SetRoutes(router fiber.Router) {
	route := router.Group("/order-edits")
	route.Get("/:id", m.Get)

	route.Post("/:id/decline", m.Decline)
	route.Post("/:id/complete", m.Complete)
}

func (m *OrderEdit) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	result, err := m.r.OrderEditService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	result, err = m.r.OrderEditService().SetContext(context.Context()).DecorateTotals(result)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *OrderEdit) Complete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	userId := api.GetUserStore(context)

	orderEdit, err := m.r.OrderEditService().SetContext(context.Context()).Retrieve(id, &sql.Options{Relations: []string{"payment_collection", "payment_collection.payments"}})
	if err != nil {
		return err
	}

	allowedStatus := []models.OrderEditStatus{models.OrderEditStatusConfirmed, models.OrderEditStatusRequested}
	if orderEdit.PaymentCollection != nil && lo.Contains(allowedStatus, orderEdit.Status) {
		if orderEdit.PaymentCollection.Status != models.PaymentCollectionStatusAuthorized {
			return utils.NewApplictaionError(
				utils.NOT_ALLOWED,
				"Unable to complete an order edit if the payment is not authorized",
			)
		}

		if orderEdit.PaymentCollection != nil {
			for _, payment := range orderEdit.PaymentCollection.Payments {
				if payment.OrderId != orderEdit.OrderId {
					if _, err := m.r.PaymentProviderService().SetContext(context.Context()).UpdatePayment(payment.Id, &types.UpdatePaymentInput{
						OrderId: orderEdit.OrderId.UUID,
					}); err != nil {
						return err
					}

				}
			}
		}

		if orderEdit.Status != models.OrderEditStatusConfirmed {
			if orderEdit.Status != models.OrderEditStatusRequested {
				return utils.NewApplictaionError(
					utils.NOT_ALLOWED,
					fmt.Sprintf("Cannot complete an order edit with status %s", orderEdit.Status),
				)
			}

			if _, err := m.r.OrderEditService().SetContext(context.Context()).Confirm(id, userId); err != nil {
				return err
			}
		}
	}

	result, err := m.r.OrderEditService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err = m.r.OrderEditService().SetContext(context.Context()).DecorateTotals(result)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *OrderEdit) Decline(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.OrderEditsDecline](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	userId := api.GetUserStore(context)

	if _, err := m.r.OrderEditService().SetContext(context.Context()).Decline(id, userId, model.DeclinedReason); err != nil {
		return err
	}

	result, err := m.r.OrderEditService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err = m.r.OrderEditService().SetContext(context.Context()).DecorateTotals(result)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

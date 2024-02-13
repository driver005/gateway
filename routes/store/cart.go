package store

import (
	"reflect"
	"strings"

	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/services"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Cart struct {
	r Registry
}

func NewCart(r Registry) *Cart {
	m := Cart{r: r}
	return &m
}

func (m *Cart) SetRoutes(router fiber.Router) {
	route := router.Group("/gift-cards")
	route.Get("/:id", m.Get)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)

	route.Post("/:id/complete", m.Complete)
	route.Post("/:id/line-items", m.CreateLineItem)
	route.Post("/:id/line-items/:line_id", m.UpdateLineItem)
	route.Delete("/:id/line-items/:line_id", m.DeleteLineItem)
	route.Post("/:id/payment-session", m.SetPaymentSession)
	route.Post("/:id/payment-sessions", m.CreatePaymentSession)
	route.Post("/:id/payment-sessions/:provider_id", m.UpdatePaymentSession)
	route.Delete("/:id/payment-sessions/:provider_id", m.DeletePaymentSession)
	route.Post("/:id/payment-sessions/:provider_id/refresh", m.RefreshPaymentSession)
	route.Post("/:id/shipping-methods", m.AddShippingMethod)
	route.Post("/:id/taxes", m.CalculateTaxes)
	route.Delete("/:id/discounts/:code", m.DeleteDiscount)
}

func (m *Cart) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	cart, err := m.r.CartService().SetContext(context.Context()).Retrieve(id, &sql.Options{Selects: []string{"id", "customer_id"}}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	customerId := context.Locals("customer_id").(uuid.UUID)

	if context.Locals("user") != nil && customerId != uuid.Nil {
		if cart.CustomerId.UUID == uuid.Nil || cart.Email == "" || cart.CustomerId.UUID != customerId {
			if _, err := m.r.CartService().SetContext(context.Context()).Update(id, nil, &types.CartUpdateProps{CustomerId: customerId}); err != nil {
				return err
			}
		}
	}

	config.Selects = append(config.Selects, "sales_channel_id")

	result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, config, services.TotalsConfig{})
	if err != nil {
		return err
	}

	if lo.Contains(config.Relations, "variant") {
		var variants []models.ProductVariant
		for _, item := range result.Items {
			variants = append(variants, *item.Variant)
		}
		if _, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(variants, uuid.UUIDs{result.SalesChannelId.UUID}, &services.AvailabilityContext{}); err != nil {
			return err
		}
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Cart) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateCart](context, m.r.Validator())
	if err != nil {
		return err
	}

	var regionId uuid.UUID
	if reflect.ValueOf(model.RegionId).IsZero() {
		regionId = model.RegionId
	} else {
		regions, err := m.r.CartService().SetContext(context.Context()).List(types.FilterableCartProps{}, &sql.Options{})
		if err != nil {
			return err
		}

		if len(regions) == 0 {
			return utils.NewApplictaionError(
				utils.INVALID_DATA,
				"A region is required to create a cart",
			)
		}

		regionId = regions[0].RegionId.UUID
	}

	model.Context = utils.MergeMaps(model.Context, map[string]interface{}{
		"ip":         context.IP(),
		"user_agent": context.Get("user-agent"),
	})

	data := &types.CartCreateProps{
		RegionId:       regionId,
		SalesChannelId: model.SalesChannelId,
		Context:        model.Context,
	}

	customerId := context.Locals("customer_id").(uuid.UUID)

	if customerId != uuid.Nil {
		customer, err := m.r.CustomerService().SetContext(context.Context()).RetrieveById(customerId, &sql.Options{})
		if err != nil {
			return err
		}

		data.CustomerId = customer.Id
		data.Email = customer.Email
	}

	if model.CountryCode != "" {
		data.ShippingAddress = &types.AddressPayload{
			CountryCode: strings.ToLower(model.CountryCode),
		}
	}

	publishableApiKeyScopes, ok := context.Locals("publishableApiKeyScopes").(types.PublishableApiKeyScopes)
	if data.SalesChannelId != uuid.Nil && ok {
		if len(publishableApiKeyScopes.SalesChannelIds) > 1 {
			return utils.NewApplictaionError(
				utils.UNEXPECTED_STATE,
				"The PublishableApiKey provided in the request header has multiple associated sales channels.",
			)
		}

		data.SalesChannelId = publishableApiKeyScopes.SalesChannelIds[0]
	}

	cart, err := m.r.CartService().SetContext(context.Context()).Create(data)
	if err != nil {
		return err
	}

	if len(model.Items) != 0 {
		var generateInputData []types.GenerateInputData
		for _, item := range model.Items {
			generateInputData = append(generateInputData, types.GenerateInputData{
				VariantId: item.VariantId,
				Quantity:  item.Quantity,
			})
		}

		//TODO: Check Quantity
		generatedLineItems, err := m.r.LineItemService().SetContext(context.Context()).Generate(uuid.Nil, generateInputData, uuid.Nil, 0, types.GenerateLineItemContext{
			RegionId:   regionId,
			CustomerId: customerId,
		})
		if err != nil {
			return err
		}

		if err := m.r.CartService().SetContext(context.Context()).AddOrUpdateLineItems(cart.Id, generatedLineItems, true); err != nil {
			return err
		}
	}

	result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(cart.Id, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	var variants []models.ProductVariant
	for _, item := range result.Items {
		variants = append(variants, *item.Variant)
	}
	if _, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(variants, uuid.UUIDs{result.SalesChannelId.UUID}, &services.AvailabilityContext{}); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Cart) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.CartUpdateProps](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	customerId := context.Locals("customer_id").(uuid.UUID)

	if customerId != uuid.Nil {
		model.CustomerId = customerId
	}

	if _, err := m.r.CartService().SetContext(context.Context()).Update(id, nil, model); err != nil {
		return err
	}

	updated, err := m.r.CartService().SetContext(context.Context()).Retrieve(id, &sql.Options{Relations: []string{"payment_sessions", "shipping_methods"}}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	if len(updated.PaymentSessions) != 0 && model.RegionId != uuid.Nil {
		if err := m.r.CartService().SetContext(context.Context()).SetPaymentSessions(id, nil); err != nil {
			return err
		}
	}

	result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	var variants []models.ProductVariant
	for _, item := range result.Items {
		variants = append(variants, *item.Variant)
	}
	if _, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(variants, uuid.UUIDs{result.SalesChannelId.UUID}, &services.AvailabilityContext{}); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Cart) Complete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	headerKey := context.Get("Idempotency-Key")

	idempotencyKey := &models.IdempotencyKey{}
	idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).InitializeRequest(headerKey, context.Method(), nil, context.Path())
	if err != nil {
		return err
	}

	context.Set("Access-Control-Expose-Headers", "Idempotency-Key")
	context.Set("Idempotency-Key", idempotencyKey.IdempotencyKey)

	//TODO: add req.request_context
	result, err := m.r.CartCompletionStrategy().Complete(id, idempotencyKey, types.RequestContext{})
	if err != nil {
		return err
	}

	return context.Status(result.ResponseCode).JSON(result.ResponseBody)
}

func (m *Cart) CreateLineItem(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.AddOrderEditLineItemInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	customerId := context.Locals("customer_id").(uuid.UUID)

	headerKey := context.Get("Idempotency-Key")

	idempotencyKey := &models.IdempotencyKey{}
	idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).InitializeRequest(headerKey, context.Method(), nil, context.Path())
	if err != nil {
		return err
	}

	context.Set("Access-Control-Expose-Headers", "Idempotency-Key")
	context.Set("Idempotency-Key", idempotencyKey.IdempotencyKey)

	if idempotencyKey.RecoveryPoint == "started" {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).WorkStage(idempotencyKey.IdempotencyKey, func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
			cart, err := m.r.CartService().SetContext(context.Context()).Retrieve(id, &sql.Options{Selects: []string{"id", "region_id", "customer_id"}}, services.TotalsConfig{})
			if err != nil {
				return nil, err
			}

			cusId := customerId
			if cusId == uuid.Nil {
				cusId = cart.CustomerId.UUID
			}

			line, err := m.r.LineItemService().SetContext(context.Context()).Generate(id, nil, cart.RegionId.UUID, model.Quantity, types.GenerateLineItemContext{
				CustomerId: cusId,
				Metadata:   model.Metadata,
			})
			if err != nil {
				return nil, err
			}

			if err := m.r.CartService().SetContext(context.Context()).AddOrUpdateLineItems(cart.Id, line, true); err != nil {
				return nil, err
			}

			return &types.IdempotencyCallbackResult{RecoveryPoint: "set-payment-sessions"}, nil
		})
		if err != nil {
			return err
		}
	} else if idempotencyKey.RecoveryPoint == "set-payment-sessions" {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).WorkStage(idempotencyKey.IdempotencyKey, func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
			//TODO: add defaultStoreCartRelations
			cart, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{Relations: []string{"region.tax_rates", "customer"}}, services.TotalsConfig{})
			if err != nil {
				return nil, err
			}

			if len(cart.PaymentSessions) > 0 {
				if err := m.r.CartService().SetContext(context.Context()).SetPaymentSessions(uuid.Nil, cart); err != nil {
					return nil, err
				}
			}

			//TODO: add defaultStoreCartRelations
			result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{Relations: []string{"region.tax_rates", "customer"}}, services.TotalsConfig{})
			if err != nil {
				return nil, err
			}

			var variants []models.ProductVariant
			for _, item := range result.Items {
				variants = append(variants, *item.Variant)
			}

			if len(variants) > 0 {
				if _, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(variants, uuid.UUIDs{result.SalesChannelId.UUID}, &services.AvailabilityContext{}); err != nil {
					return nil, err
				}
			}

			return &types.IdempotencyCallbackResult{ResponseCode: 200, ResponseBody: core.JSONB{"cart": result}}, nil
		})
		if err != nil {
			return err
		}
	} else {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).Update(idempotencyKey.IdempotencyKey, &models.IdempotencyKey{
			RecoveryPoint: "finished",
			ResponseCode:  500,
			ResponseBody:  core.JSONB{"message": "Unknown recovery point"},
		})
		if err != nil {
			return err
		}
	}

	return context.Status(idempotencyKey.ResponseCode).JSON(idempotencyKey.ResponseBody)
}

func (m *Cart) UpdateLineItem(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateLineItem](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	lineId, err := api.BindDelete(context, "line_id")
	if err != nil {
		return err
	}

	if model.Quantity == 0 {
		if err = m.r.CartService().SetContext(context.Context()).RemoveLineItem(id, uuid.UUIDs{lineId}); err != nil {
			return err
		}
	} else {
		cart, err := m.r.CartService().SetContext(context.Context()).Retrieve(id, &sql.Options{Relations: []string{"items", "items.variant", "shipping_methods"}}, services.TotalsConfig{})
		if err != nil {
			return err
		}

		existing, ok := lo.Find(cart.Items, func(item models.LineItem) bool {
			return item.Id == lineId
		})

		if !ok {
			return utils.NewApplictaionError(
				utils.INVALID_DATA,
				"Could not find the line item",
			)
		}

		data := &types.LineItemUpdate{
			VariantId:             existing.Variant.Id,
			RegionId:              cart.RegionId.UUID,
			Quantity:              model.Quantity,
			Metadata:              model.Metadata,
			ShouldCalculatePrices: true,
		}

		if _, err = m.r.CartService().SetContext(context.Context()).UpdateLineItem(id, lineId, data); err != nil {
			return err
		}
	}

	updated, err := m.r.CartService().SetContext(context.Context()).Retrieve(id, &sql.Options{Relations: []string{"payment_sessions"}}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	if len(updated.PaymentSessions) > 0 {
		if err := m.r.CartService().SetContext(context.Context()).SetPaymentSessions(id, nil); err != nil {
			return err
		}
	}

	result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	var variants []models.ProductVariant
	for _, item := range result.Items {
		variants = append(variants, *item.Variant)
	}
	if _, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(variants, uuid.UUIDs{result.SalesChannelId.UUID}, &services.AvailabilityContext{}); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Cart) DeleteLineItem(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	lineId, err := api.BindDelete(context, "line_id")
	if err != nil {
		return err
	}

	if err := m.r.CartService().SetContext(context.Context()).RemoveLineItem(id, uuid.UUIDs{lineId}); err != nil {
		return err
	}

	updated, err := m.r.CartService().SetContext(context.Context()).Retrieve(id, &sql.Options{Relations: []string{"payment_sessions"}}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	if len(updated.PaymentSessions) > 0 {
		if err := m.r.CartService().SetContext(context.Context()).SetPaymentSessions(id, nil); err != nil {
			return err
		}
	}

	result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	var variants []models.ProductVariant
	for _, item := range result.Items {
		variants = append(variants, *item.Variant)
	}
	if _, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(variants, uuid.UUIDs{result.SalesChannelId.UUID}, &services.AvailabilityContext{}); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Cart) CreatePaymentSession(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	headerKey := context.Get("Idempotency-Key")

	idempotencyKey := &models.IdempotencyKey{}
	idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).InitializeRequest(headerKey, context.Method(), nil, context.Path())
	if err != nil {
		return err
	}

	context.Set("Access-Control-Expose-Headers", "Idempotency-Key")
	context.Set("Idempotency-Key", idempotencyKey.IdempotencyKey)

	if idempotencyKey.RecoveryPoint == "started" {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).WorkStage(idempotencyKey.IdempotencyKey, func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
			//TODO: add defaultStoreCartRelations
			cart, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{Relations: []string{"region.tax_rates", "customer"}}, services.TotalsConfig{})
			if err != nil {
				return nil, err
			}

			if len(cart.PaymentSessions) > 0 {
				if err := m.r.CartService().SetContext(context.Context()).SetPaymentSessions(uuid.Nil, cart); err != nil {
					return nil, err
				}
			}

			//TODO: add defaultStoreCartRelations
			result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{Relations: []string{"region.tax_rates", "customer"}}, services.TotalsConfig{})
			if err != nil {
				return nil, err
			}

			var variants []models.ProductVariant
			for _, item := range result.Items {
				variants = append(variants, *item.Variant)
			}

			if len(variants) > 0 {
				if _, err := m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(variants, uuid.UUIDs{result.SalesChannelId.UUID}, &services.AvailabilityContext{}); err != nil {
					return nil, err
				}
			}

			return &types.IdempotencyCallbackResult{ResponseCode: 200, ResponseBody: core.JSONB{"cart": result}}, nil
		})
		if err != nil {
			return err
		}
	} else {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).Update(idempotencyKey.IdempotencyKey, &models.IdempotencyKey{
			RecoveryPoint: "finished",
			ResponseCode:  500,
			ResponseBody:  core.JSONB{"message": "Unknown recovery point"},
		})
		if err != nil {
			return err
		}
	}

	return context.Status(idempotencyKey.ResponseCode).JSON(idempotencyKey.ResponseBody)
}

func (m *Cart) UpdatePaymentSession(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdatePaymentSession](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	providerId, err := api.BindDelete(context, "provider_id")
	if err != nil {
		return err
	}

	if err := m.r.CartService().SetContext(context.Context()).SetPaymentSession(id, providerId); err != nil {
		return err
	}

	if _, err := m.r.CartService().SetContext(context.Context()).UpdatePaymentSession(id, model.Data); err != nil {
		return err
	}

	result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	var variants []models.ProductVariant
	for _, item := range result.Items {
		variants = append(variants, *item.Variant)
	}
	if _, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(variants, uuid.UUIDs{result.SalesChannelId.UUID}, &services.AvailabilityContext{}); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Cart) DeletePaymentSession(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	providerId, err := api.BindDelete(context, "provider_id")
	if err != nil {
		return err
	}

	if err := m.r.CartService().SetContext(context.Context()).DeletePaymentSession(id, providerId); err != nil {
		return err
	}

	result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	var variants []models.ProductVariant
	for _, item := range result.Items {
		variants = append(variants, *item.Variant)
	}
	if _, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(variants, uuid.UUIDs{result.SalesChannelId.UUID}, &services.AvailabilityContext{}); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Cart) RefreshPaymentSession(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	providerId, err := api.BindDelete(context, "provider_id")
	if err != nil {
		return err
	}

	if err := m.r.CartService().SetContext(context.Context()).RefreshPaymentSession(id, providerId); err != nil {
		return err
	}

	result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{
		Relations: []string{
			"region",
			"region.countries",
			"region.payment_providers",
			"shipping_methods",
			"payment_sessions",
			"shipping_methods.shipping_option",
		},
	}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Cart) SetPaymentSession(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.SessionsInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if err := m.r.CartService().SetContext(context.Context()).SetPaymentSession(id, model.ProviderId); err != nil {
		return err
	}

	result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	var variants []models.ProductVariant
	for _, item := range result.Items {
		variants = append(variants, *item.Variant)
	}
	if _, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(variants, uuid.UUIDs{result.SalesChannelId.UUID}, &services.AvailabilityContext{}); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Cart) AddShippingMethod(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.AddShippingMethod](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if _, err := m.r.CartService().SetContext(context.Context()).AddShippingMethod(id, nil, model.OptionId, model.Data); err != nil {
		return err
	}

	updated, err := m.r.CartService().SetContext(context.Context()).Retrieve(id, &sql.Options{Relations: []string{"payment_sessions"}}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	if len(updated.PaymentSessions) > 0 {
		if err := m.r.CartService().SetContext(context.Context()).SetPaymentSessions(id, nil); err != nil {
			return err
		}
	}

	result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	var variants []models.ProductVariant
	for _, item := range result.Items {
		variants = append(variants, *item.Variant)
	}
	if _, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(variants, uuid.UUIDs{result.SalesChannelId.UUID}, &services.AvailabilityContext{}); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Cart) CalculateTaxes(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	headerKey := context.Get("Idempotency-Key")

	idempotencyKey := &models.IdempotencyKey{}
	idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).InitializeRequest(headerKey, context.Method(), nil, context.Path())
	if err != nil {
		return err
	}

	context.Set("Access-Control-Expose-Headers", "Idempotency-Key")
	context.Set("Idempotency-Key", idempotencyKey.IdempotencyKey)

	if idempotencyKey.RecoveryPoint == "started" {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).WorkStage(idempotencyKey.IdempotencyKey, func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
			result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{}, services.TotalsConfig{ForceTaxes: true})
			if err != nil {
				return nil, err
			}

			return &types.IdempotencyCallbackResult{ResponseCode: 200, ResponseBody: core.JSONB{"cart": result}}, nil
		})
		if err != nil {
			return err
		}
	} else {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).Update(idempotencyKey.IdempotencyKey, &models.IdempotencyKey{
			RecoveryPoint: "finished",
			ResponseCode:  500,
			ResponseBody:  core.JSONB{"message": "Unknown recovery point"},
		})
		if err != nil {
			return err
		}
	}

	return context.Status(idempotencyKey.ResponseCode).JSON(idempotencyKey.ResponseBody)
}

func (m *Cart) DeleteDiscount(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	code := context.Params("code")

	if _, err := m.r.CartService().SetContext(context.Context()).RemoveDiscount(id, code); err != nil {
		return err
	}

	updated, err := m.r.CartService().SetContext(context.Context()).Retrieve(id, &sql.Options{Relations: []string{"payment_sessions"}}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	if len(updated.PaymentSessions) > 0 {
		if err := m.r.CartService().SetContext(context.Context()).SetPaymentSessions(id, nil); err != nil {
			return err
		}
	}

	result, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	var variants []models.ProductVariant
	for _, item := range result.Items {
		variants = append(variants, *item.Variant)
	}
	if _, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(variants, uuid.UUIDs{result.SalesChannelId.UUID}, &services.AvailabilityContext{}); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)

}

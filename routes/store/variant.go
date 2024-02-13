package store

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/services"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type Variant struct {
	r Registry
}

func NewVariant(r Registry) *Variant {
	m := Variant{r: r}
	return &m
}

func (m *Variant) SetRoutes(router fiber.Router) {
	route := router.Group("/variants")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
}

func (m *Variant) Get(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.ProductVariantVariant](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	customerId := context.Locals("customer_id").(uuid.UUID)

	variant, err := m.r.ProductVariantService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	salesChannelId := model.SalesChannelId
	publishableApiKeyScopes, ok := context.Locals("publishableApiKeyScopes").(types.PublishableApiKeyScopes)
	if ok {
		salesChannelId = publishableApiKeyScopes.SalesChannelIds[0]
	}

	regionId := model.RegionId
	currencyCode := model.CurrencyCode
	if model.CartId != uuid.Nil {
		cart, err := m.r.CartService().SetContext(context.Context()).Retrieve(model.CartId, &sql.Options{Selects: []string{"id", "region_id"}}, services.TotalsConfig{})
		if err != nil {
			return err
		}
		region, err := m.r.RegionService().SetContext(context.Context()).Retrieve(cart.RegionId.UUID, &sql.Options{Selects: []string{"id", "currency_code"}})
		if err != nil {
			return err
		}
		regionId = region.Id
		currencyCode = region.CurrencyCode
	}

	prices, err := m.r.PricingService().SetContext(context.Context()).SetVariantPrices([]models.ProductVariant{*variant}, &interfaces.PricingContext{
		CartId:                model.CartId,
		CustomerId:            customerId,
		RegionId:              regionId,
		CurrencyCode:          currencyCode,
		IncludeDiscountPrices: true,
	})
	if err != nil {
		return err
	}

	result, err := m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(prices, uuid.UUIDs{salesChannelId}, &services.AvailabilityContext{})
	if err != nil {
		return err
	}

	//TODO: Result only variant not list
	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Variant) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.ProductVariantParams](context)
	if err != nil {
		return err
	}

	customerId := context.Locals("customer_id").(uuid.UUID)
	publishableApiKeyScopes, ok := context.Locals("publishableApiKeyScopes").(types.PublishableApiKeyScopes)
	if ok {
		model.SalesChannelId = publishableApiKeyScopes.SalesChannelIds[0]
	}

	variants, err := m.r.ProductVariantService().SetContext(context.Context()).List(&types.FilterableProductVariant{
		FilterModel: core.FilterModel{
			Id: []uuid.UUID{model.Id},
		},
		Title:             model.Title,
		InventoryQuantity: model.InventoryQuantity,
	}, config)
	if err != nil {
		return err
	}

	regionId := model.RegionId
	currencyCode := model.CurrencyCode
	if model.CartId != uuid.Nil {
		cart, err := m.r.CartService().SetContext(context.Context()).Retrieve(model.CartId, &sql.Options{Selects: []string{"id", "region_id"}}, services.TotalsConfig{})
		if err != nil {
			return err
		}
		region, err := m.r.RegionService().SetContext(context.Context()).Retrieve(cart.RegionId.UUID, &sql.Options{Selects: []string{"id", "currency_code"}})
		if err != nil {
			return err
		}
		regionId = region.Id
		currencyCode = region.CurrencyCode
	}

	prices, err := m.r.PricingService().SetContext(context.Context()).SetVariantPrices(variants, &interfaces.PricingContext{
		CartId:                model.CartId,
		CustomerId:            customerId,
		RegionId:              regionId,
		CurrencyCode:          currencyCode,
		IncludeDiscountPrices: true,
	})
	if err != nil {
		return err
	}

	result, err := m.r.ProductVariantInventoryService().SetContext(context.Context()).SetVariantAvailability(prices, uuid.UUIDs{model.SalesChannelId}, &services.AvailabilityContext{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

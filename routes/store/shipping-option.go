package store

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/services"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"golang.org/x/exp/maps"
)

type ShippingOption struct {
	r Registry
}

func NewShippingOption(r Registry) *ShippingOption {
	m := ShippingOption{r: r}
	return &m
}

func (m *ShippingOption) SetRoutes(router fiber.Router) {
	route := router.Group("/shipping-options")
	route.Get("", m.ListOptions)
	route.Get("/:cart_id", m.List)
}

func (m *ShippingOption) List(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "cart_id")
	if err != nil {
		return err
	}
	cart, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(id, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	options, err := m.r.ShippingProfileService().SetContext(context.Context()).FetchCartOptions(cart)
	if err != nil {
		return err
	}

	result, err := m.r.PricingService().SetContext(context.Context()).SetShippingOptionPrices(options, &interfaces.PricingContext{
		CartId: id,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"shipping_options": result,
	})
}

func (m *ShippingOption) ListOptions(context fiber.Ctx) error {
	model, config, err := api.BindList[types.ShippingOptionParams](context)
	if err != nil {
		return err
	}

	query := &types.FilterableShippingOption{}

	if context.Query("is_return") != "" {
		query.IsReturn = model.IsReturn == "true"
	}

	if model.RegionId != uuid.Nil {
		query.RegionId = model.RegionId
	}

	query.AdminOnly = false

	if len(model.ProductIds) > 0 {
		productShippinProfileMap, err := m.r.ShippingProfileService().SetContext(context.Context()).GetMapProfileIdsByProductIds(model.ProductIds)

		query.ProfileId = maps.Values(*productShippinProfileMap)
		if err != nil {
			return err
		}
	}

	result, err := m.r.ShippingOptionService().SetContext(context.Context()).List(query, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"shipping_options": result,
	})
}

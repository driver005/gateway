package store

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/services"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Product struct {
	r Registry
}

func NewProduct(r Registry) *Product {
	m := Product{r: r}
	return &m
}

func (m *Product) SetRoutes(router fiber.Router) {
	route := router.Group("/regions")
	route.Get("/:id", m.Get)
	route.Get("", m.List)

	route.Post("", m.Search)
}

func (m *Product) Get(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.ProductVariantVariant](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	customerId := context.Locals("customer_id").(uuid.UUID)

	product, err := m.r.ProductService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	salesChannelId := model.SalesChannelId
	if context.Locals("publishableApiKeyScopes").SalesChannelIds != nil {
		salesChannelId = context.Locals("publishableApiKeyScopes").SalesChannelIds[0]
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

	result := []models.Product{*product}

	if lo.ContainsBy(config.Relations, func(item string) bool {
		return item == "variants" || item == "variants.prices"
	}) {
		result, err = m.r.PricingService().SetContext(context.Context()).SetProductPrices(result, &interfaces.PricingContext{
			CartId:                model.CartId,
			CustomerId:            customerId,
			RegionId:              regionId,
			CurrencyCode:          currencyCode,
			IncludeDiscountPrices: true,
		})
		if err != nil {
			return err
		}
	}

	if lo.ContainsBy(config.Relations, func(item string) bool {
		return item == "variants"
	}) {
		result, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetProductAvailability(result, uuid.UUIDs{salesChannelId})
		if err != nil {
			return err
		}
	}

	//TODO: Result only variant not list
	return context.Status(fiber.StatusOK).JSON(result[0])
}

func (m *Product) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableProduct](context)
	if err != nil {
		return err
	}

	customerId := context.Locals("customer_id").(uuid.UUID)
	regionId := model.RegionId
	currencyCode := model.CurrencyCode

	model.Status = []models.ProductStatus{models.ProductStatusPublished}

	if context.Locals("publishableApiKeyScopes").SalesChannelIds != nil {
		if model.SalesChannelId == nil {
			model.SalesChannelId = context.Locals("publishableApiKeyScopes").SalesChannelIds[0]
		}

		if !lo.Contains(config.Relations, "listConfig.relations") {
			config.Relations = append(config.Relations, "sales_channels")
		}
	}

	result, count, err := m.r.ProductService().SetContext(context.Context()).ListAndCount(model, config)
	if err != nil {
		return err
	}

	cart := &models.Cart{}

	if model.CartId != uuid.Nil {
		cart, err = m.r.CartService().SetContext(context.Context()).Retrieve(model.CartId, &sql.Options{Selects: []string{"id", "region_id"}, Relations: []string{"region"}}, services.TotalsConfig{})
		if err != nil {
			return err
		}

		regionId = cart.RegionId.UUID
		currencyCode = cart.Region.CurrencyCode
	}

	if lo.ContainsBy(config.Relations, func(item string) bool {
		return item == "variants" || item == "variants.prices"
	}) {
		result, err = m.r.PricingService().SetContext(context.Context()).SetProductPrices(result, &interfaces.PricingContext{
			CartId:                model.CartId,
			CustomerId:            customerId,
			RegionId:              regionId,
			CurrencyCode:          currencyCode,
			IncludeDiscountPrices: true,
		})
		if err != nil {
			return err
		}
	}

	if lo.ContainsBy(config.Relations, func(item string) bool {
		return item == "variants"
	}) {
		result, err = m.r.ProductVariantInventoryService().SetContext(context.Context()).SetProductAvailability(result, model.SalesChannelId)
		if err != nil {
			return err
		}
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   result,
		"count":  count,
		"offset": config.Skip,
		"limit":  config.Take,
	})
}

func (m *Product) Search(context fiber.Ctx) error {
	model, err := api.Bind[types.ProductSearch](context, m.r.Validator())
	if err != nil {
		return err
	}

	result := m.r.DefaultSearchService().SetContext(context.Context()).Search("product", model.Q, map[string]interface{}{
		"limit":  model.Limit,
		"offset": model.Offset,
		"filter": model.Filter,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

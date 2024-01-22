package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type PriceList struct {
	r Registry
}

func NewPriceList(r Registry) *PriceList {
	m := PriceList{r: r}
	return &m
}

func (m *PriceList) SetRoutes(router fiber.Router) {
	route := router.Group("/price-lists")
	route.Get("/:id", m.Get)
	route.Get("/", m.List)
	route.Post("/", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)

	route.Get("/:id/products", m.ListPriceListProducts)
	route.Delete("/:id/products/:product_id/prices", m.DeleteProductPrices)
	route.Delete("/:id/products/prices/batch", m.DeleteProductPricesBatch)
	route.Delete("/:id/variants/:variant_id/prices", m.DeleteVariantPrices)
	route.Delete("/:id/prices/batch", m.DeletePricesBatch)
	route.Post("/:id/prices/batch", m.AddPricesBatch)
}

func (m *PriceList) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.PriceListService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *PriceList) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterablePriceList](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.PriceListService().SetContext(context.Context()).ListAndCount(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   result,
		"count":  count,
		"offset": config.Skip,
		"limit":  config.Take,
	})
}

func (m *PriceList) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreatePriceListInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.PriceListService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *PriceList) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdatePriceListInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.PriceListService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *PriceList) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.PriceListService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "price-list",
		"deleted": true,
	})
}

func (m *PriceList) AddPricesBatch(context fiber.Ctx) error {
	return nil
}

func (m *PriceList) DeletePricesBatch(context fiber.Ctx) error {
	return nil
}

func (m *PriceList) DeleteProductPrices(context fiber.Ctx) error {
	return nil
}

func (m *PriceList) DeleteProductPricesBatch(context fiber.Ctx) error {
	return nil
}

func (m *PriceList) DeleteVariantPrices(context fiber.Ctx) error {
	return nil
}

func (m *PriceList) ListPriceListProducts(context fiber.Ctx) error {
	return nil
}

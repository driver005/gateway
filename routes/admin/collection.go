package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type Collection struct {
	r Registry
}

func NewCollection(r Registry) *Collection {
	m := Collection{r: r}
	return &m
}

func (m *Collection) SetRoutes(router fiber.Router) {
	route := router.Group("/collections")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)

	route.Post("/:id/products", m.AddProducts)
	route.Delete("/:id/products", m.RemoveProducts)
}

func (m *Collection) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.ProductCollectionService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Collection) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableCollection](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.ProductCollectionService().SetContext(context.Context()).ListAndCount(model, config)
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

func (m *Collection) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateProductCollection](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ProductCollectionService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Collection) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateProductCollection](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ProductCollectionService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Collection) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.ProductCollectionService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "product-collection",
		"deleted": true,
	})
}

func (m *Collection) AddProducts(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.AddProductsToCollectionInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ProductCollectionService().AddProducts(id, model.ProductIds)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Collection) RemoveProducts(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.AddProductsToCollectionInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if err := m.r.ProductCollectionService().RemoveProducts(id, model.ProductIds); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":               id,
		"object":           "product-collection",
		"removed_products": model.ProductIds,
	})
}

package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type Product struct {
	r Registry
}

func NewProduct(r Registry) *Product {
	m := Product{r: r}
	return &m
}

func (m *Product) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.ProductService().SetContext(context.Context()).RetrieveById(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Product) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableProduct](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.ProductService().SetContext(context.Context()).ListAndCount(model, config)
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

func (m *Product) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateProductInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ProductService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Product) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateProductInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ProductService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Product) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.ProductService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "product",
		"deleted": true,
	})
}

func (m *Product) AddOption(context fiber.Ctx) error {
	return nil
}

func (m *Product) DeletOption(context fiber.Ctx) error {
	return nil
}

func (m *Product) CreateVariant(context fiber.Ctx) error {
	return nil
}

func (m *Product) DeletVariant(context fiber.Ctx) error {
	return nil
}

func (m *Product) ListTagUsageCount(context fiber.Ctx) error {
	return nil
}

func (m *Product) ListTypes(context fiber.Ctx) error {
	return nil
}

func (m *Product) ListVariants(context fiber.Ctx) error {
	return nil
}

func (m *Product) SetMetadata(context fiber.Ctx) error {
	return nil
}

func (m *Product) UpdateOption(context fiber.Ctx) error {
	return nil
}

func (m *Product) UpdateVariant(context fiber.Ctx) error {
	return nil
}

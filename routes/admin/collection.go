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
	return nil
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
	return nil
}

func (m *Collection) Delete(context fiber.Ctx) error {
	return nil
}

func (m *Collection) AddProducts(context fiber.Ctx) error {
	return nil
}

func (m *Collection) RemoveProducts(context fiber.Ctx) error {
	return nil
}

package admin

import "github.com/gofiber/fiber/v3"

type Product struct {
	r Registry
}

func NewProduct(r Registry) *Product {
	m := Product{r: r}
	return &m
}

func (m *Product) Get(context fiber.Ctx) error {
	return nil
}

func (m *Product) List(context fiber.Ctx) error {
	return nil
}

func (m *Product) Create(context fiber.Ctx) error {
	return nil
}

func (m *Product) Update(context fiber.Ctx) error {
	return nil
}

func (m *Product) Delete(context fiber.Ctx) error {
	return nil
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

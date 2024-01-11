package admin

import "github.com/gofiber/fiber/v3"

type Region struct {
	r Registry
}

func NewRegion(r Registry) *Region {
	m := Region{r: r}
	return &m
}

func (m *Region) Get(context fiber.Ctx) error {
	return nil
}

func (m *Region) List(context fiber.Ctx) error {
	return nil
}

func (m *Region) Create(context fiber.Ctx) error {
	return nil
}

func (m *Region) Update(context fiber.Ctx) error {
	return nil
}

func (m *Region) Delete(context fiber.Ctx) error {
	return nil
}

func (m *Region) AddCountry(context fiber.Ctx) error {
	return nil
}

func (m *Region) AddFullfilmentProvider(context fiber.Ctx) error {
	return nil
}

func (m *Region) AddPaymentProvider(context fiber.Ctx) error {
	return nil
}

func (m *Region) GetFulfillmentOptions(context fiber.Ctx) error {
	return nil
}

func (m *Region) RemoveCountry(context fiber.Ctx) error {
	return nil
}

func (m *Region) RemoveFullfilmentProvider(context fiber.Ctx) error {
	return nil
}

func (m *Region) RemovePaymentProvider(context fiber.Ctx) error {
	return nil
}

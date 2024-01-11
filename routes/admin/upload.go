package admin

import "github.com/gofiber/fiber/v3"

type Upload struct {
	r Registry
}

func NewUpload(r Registry) *Upload {
	m := Upload{r: r}
	return &m
}

func (m *Upload) Get(context fiber.Ctx) error {
	return nil
}

func (m *Upload) Create(context fiber.Ctx) error {
	return nil
}

func (m *Upload) Delete(context fiber.Ctx) error {
	return nil
}

func (m *Upload) CreateProtectedUpload(context fiber.Ctx) error {
	return nil
}

package admin

import "github.com/gofiber/fiber/v3"

type ProductTag struct {
	r Registry
}

func NewProductTag(r Registry) *ProductTag {
	m := ProductTag{r: r}
	return &m
}

func (m *ProductTag) List(context fiber.Ctx) error {
	return nil
}

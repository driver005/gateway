package admin

import "github.com/gofiber/fiber/v3"

type ProductType struct {
	r Registry
}

func NewProductType(r Registry) *ProductType {
	m := ProductType{r: r}
	return &m
}

func (m *ProductType) List(context fiber.Ctx) error {
	return nil
}

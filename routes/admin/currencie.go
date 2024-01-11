package admin

import "github.com/gofiber/fiber/v3"

type Currencie struct {
	r Registry
}

func NewCurrencie(r Registry) *Currencie {
	m := Currencie{r: r}
	return &m
}

func (m *Currencie) List(context fiber.Ctx) error {
	return nil
}

func (m *Currencie) Update(context fiber.Ctx) error {
	return nil
}

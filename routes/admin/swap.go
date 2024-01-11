package admin

import "github.com/gofiber/fiber/v3"

type Swap struct {
	r Registry
}

func NewSwap(r Registry) *Swap {
	m := Swap{r: r}
	return &m
}

func (m *Swap) Get(context fiber.Ctx) error {
	return nil
}

func (m *Swap) List(context fiber.Ctx) error {
	return nil
}

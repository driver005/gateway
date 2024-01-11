package admin

import "github.com/gofiber/fiber/v3"

type Note struct {
	r Registry
}

func NewNote(r Registry) *Note {
	m := Note{r: r}
	return &m
}

func (m *Note) Get(context fiber.Ctx) error {
	return nil
}

func (m *Note) List(context fiber.Ctx) error {
	return nil
}

func (m *Note) Create(context fiber.Ctx) error {
	return nil
}

func (m *Note) Update(context fiber.Ctx) error {
	return nil
}

func (m *Note) Delete(context fiber.Ctx) error {
	return nil
}

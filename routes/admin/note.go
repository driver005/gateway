package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

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
	model, err := api.BindCreate[types.CreateNoteInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.NoteService().SetContext(context.Context()).Create(model, map[string]interface{}{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Note) Update(context fiber.Ctx) error {
	return nil
}

func (m *Note) Delete(context fiber.Ctx) error {
	return nil
}

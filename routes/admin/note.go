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
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.NoteService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Note) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableNote](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.NoteService().SetContext(context.Context()).ListAndCount(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   result,
		"count":  count,
		"offset": config.Skip,
		"limit":  config.Take,
	})
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
	model, id, err := api.BindUpdate[types.UpdateNoteInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.NoteService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Note) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.NoteService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "note",
		"deleted": true,
	})
}

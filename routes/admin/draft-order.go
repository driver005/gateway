package admin

import "github.com/gofiber/fiber/v3"

type DraftOrder struct {
	r Registry
}

func NewDraftOrder(r Registry) *DraftOrder {
	m := DraftOrder{r: r}
	return &m
}

func (m *DraftOrder) Get(context fiber.Ctx) error {
	return nil
}

func (m *DraftOrder) List(context fiber.Ctx) error {
	return nil
}

func (m *DraftOrder) Create(context fiber.Ctx) error {
	return nil
}

func (m *DraftOrder) Update(context fiber.Ctx) error {
	return nil
}

func (m *DraftOrder) Delete(context fiber.Ctx) error {
	return nil
}

func (m *DraftOrder) CreateLineItem(context fiber.Ctx) error {
	return nil
}

func (m *DraftOrder) RegisterPayment(context fiber.Ctx) error {
	return nil
}

func (m *DraftOrder) UpdateLineItem(context fiber.Ctx) error {
	return nil
}

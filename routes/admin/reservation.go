package admin

import "github.com/gofiber/fiber/v3"

type Reservation struct {
	r Registry
}

func NewReservation(r Registry) *Reservation {
	m := Reservation{r: r}
	return &m
}

func (m *Reservation) Get(context fiber.Ctx) error {
	return nil
}

func (m *Reservation) List(context fiber.Ctx) error {
	return nil
}

func (m *Reservation) Create(context fiber.Ctx) error {
	return nil
}

func (m *Reservation) Update(context fiber.Ctx) error {
	return nil
}

func (m *Reservation) Delete(context fiber.Ctx) error {
	return nil
}

package admin

import "github.com/gofiber/fiber/v3"

type User struct {
	r Registry
}

func NewUser(r Registry) *User {
	m := User{r: r}
	return &m
}

func (m *User) Get(context fiber.Ctx) error {
	return nil
}

func (m *User) List(context fiber.Ctx) error {
	return nil
}

func (m *User) Create(context fiber.Ctx) error {
	return nil
}

func (m *User) Update(context fiber.Ctx) error {
	return nil
}

func (m *User) Delete(context fiber.Ctx) error {
	return nil
}

func (m *User) ResetPassword(context fiber.Ctx) error {
	return nil
}

func (m *User) ResetPasswordTocken(context fiber.Ctx) error {
	return nil
}

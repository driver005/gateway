package utils

import "github.com/gofiber/fiber/v3"

func ConvertMiddleware(m []func(fiber.Ctx) error) []any {
	var result []any
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

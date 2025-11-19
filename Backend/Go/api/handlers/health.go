package handlers

import (
	"github.com/gofiber/fiber/v3"
)

// HealthCheck simple health endpoint
func HealthCheck(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"ok": true,
	})
}

package controller

import (
	services "city/services"

	"github.com/gofiber/fiber/v2"
)

func SingIn(c *fiber.Ctx) error {
	return services.SingIn(c)
}

func SingUp(c *fiber.Ctx) error {
	return services.SingUp(c)
}

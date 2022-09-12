package services

import "github.com/gofiber/fiber/v2"

func SingIn(c *fiber.Ctx) error {
	return c.SendString("Sing")
}

func SingUp(c *fiber.Ctx) error {
	return c.SendString("SingUp")
}

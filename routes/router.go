package router

import "github.com/gofiber/fiber/v2"

func Router(app *fiber.App) {
	auth(app)
}

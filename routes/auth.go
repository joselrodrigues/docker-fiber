package routes

import (
	c "city/controllers"

	"github.com/gofiber/fiber/v2"
)

func auth(app *fiber.App) {
	app.Get("/singin", c.SingIn)
	app.Get("/singup", c.SingUp)
}

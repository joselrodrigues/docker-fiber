package main

import (
	r "city/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	r.Router(app)

	app.Listen(":4001")
}

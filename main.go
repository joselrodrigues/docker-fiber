package main

import (
	db "city/db"
	logger "city/log"
	routes "city/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	logger.Setup()
	routes.Setup(app)
	db.Configuration()
	app.Listen(":4001")
}

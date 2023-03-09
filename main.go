package main

import (
	"github/co-codin/service-layer-fiber/src/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, da 👋!")
	})

	app.Listen(":3000")
}

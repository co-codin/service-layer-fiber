package main

import (
	"github/co-codin/service-layer-fiber/src/database"
	"github/co-codin/service-layer-fiber/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	database.AutoMigrate()

	app := fiber.New()

	app.Use(cors.New(
		cors.Config{
			AllowCredentials: true,
		},
	))

	routes.SetUp(app)

	app.Listen(":3000")
}

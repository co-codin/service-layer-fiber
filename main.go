package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	_, err := gorm.Open(
		postgres.Open("postgres://postgres:root@localhost:5432/service_layer_fiber"),
		&gorm.Config{},
	)

	if err != nil {
		panic("can not connect to postgres")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")
}

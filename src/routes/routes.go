package routes

import (
	"github/co-codin/service-layer-fiber/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	api := app.Group("api")
	admin := api.Group("admin")

	admin.Post("/register", controllers.Register)
	admin.Post("/login", controllers.Login)
	admin.Get("/user", controllers.User)
}

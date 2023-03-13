package controllers

import (
	"github/co-codin/service-layer-fiber/src/database"
	"github/co-codin/service-layer-fiber/src/models"

	"github.com/gofiber/fiber/v2"
)

func Products(c *fiber.Ctx) error {
	var products []models.Product

	database.DB.Find(&products)

	return c.JSON(products)
}

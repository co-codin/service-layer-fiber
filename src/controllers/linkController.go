package controllers

import (
	"github/co-codin/service-layer-fiber/src/database"
	"github/co-codin/service-layer-fiber/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Link(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var links []models.Link

	database.DB.Where("user_id = ?", id).Find(&links)

	return c.JSON(links)
}

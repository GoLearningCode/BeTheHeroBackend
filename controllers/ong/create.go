package ong

import (
	"github.com/GoLearningCode/BeTheHeroBackend/database"
	"github.com/GoLearningCode/BeTheHeroBackend/models"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	ong := models.Ong{}
	c.BodyParser(&ong)

	db := database.ConnectToDb()
	db.Table("ong")

	db.Create(&ong)

	return c.JSON(map[string]models.Ong{"ong": ong})
}

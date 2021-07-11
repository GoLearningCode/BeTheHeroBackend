package ong

import (
	"github.com/GoLearningCode/BeTheHeroBackend/database"
	"github.com/GoLearningCode/BeTheHeroBackend/models"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	db := database.ConnectToDb()
	db.Table("ong")
	ongs := []models.Ong{}
	db.Find(&ongs)
	return c.JSON(map[string][]models.Ong{"ongs": ongs})
}

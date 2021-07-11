package session

import (
	"github.com/GoLearningCode/BeTheHeroBackend/database"
	"github.com/GoLearningCode/BeTheHeroBackend/models"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	id := struct {
		ID int
	}{}
	c.BodyParser(&id)

	db := database.ConnectToDb()

	db.Table("ong")

	ong := models.Ong{}

	db.Where("id = ?", id.ID).First(&ong)

	return c.JSON(map[string]string{"ong": ong.Name})
}

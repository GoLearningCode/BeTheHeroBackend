package incidents

import (
	"github.com/GoLearningCode/BeTheHeroBackend/database"
	"github.com/GoLearningCode/BeTheHeroBackend/models"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	incident := models.Incident{}

	c.BodyParser(&incident)

	db := database.ConnectToDb()

	db.Table("incident")

	db.Create(&incident)

	return c.JSON(map[string]models.Incident{"incident": incident})
}

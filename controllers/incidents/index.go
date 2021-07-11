package incidents

import (
	"github.com/GoLearningCode/BeTheHeroBackend/database"
	"github.com/GoLearningCode/BeTheHeroBackend/models"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	db := database.ConnectToDb()

	db.Table("incident")

	incidents := []models.Incident{}

	db.Find(&incidents)

	return c.JSON(map[string][]models.Incident{"incidents": incidents})
}

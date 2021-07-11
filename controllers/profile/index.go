package profile

import (
	"github.com/GoLearningCode/BeTheHeroBackend/database"
	"github.com/GoLearningCode/BeTheHeroBackend/models"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	db := database.ConnectToDb()
	db.Table("incidents")

	id := c.Params("id")

	incident := []models.Incident{}
	db.Where("ong_id = ?", id).Find(&incident)
	return c.JSON(map[string][]models.Incident{"incidents": incident})
}

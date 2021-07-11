package incidents

import (
	"github.com/GoLearningCode/BeTheHeroBackend/database"
	"github.com/GoLearningCode/BeTheHeroBackend/models"
	"github.com/gofiber/fiber/v2"
)

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.ConnectToDb()

	db.Table("incident")

	incident := models.Incident{}

	db.First(&incident, id)

	db.Delete(&incident, id)
	return c.JSON(map[string]models.Incident{"incident": incident})
}

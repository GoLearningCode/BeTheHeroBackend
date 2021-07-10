package main

import (
	"github.com/GoLearningCode/BeTheHeroBackend/database"
	"github.com/GoLearningCode/BeTheHeroBackend/models"
	"github.com/GoLearningCode/BeTheHeroBackend/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := database.ConnectToDb()
	db.AutoMigrate(&models.Incident{}, &models.Ong{})
	app := fiber.New()
	routes.Route(app)
	app.Listen(":3000")
}

package main

import (
	"github.com/GoLearningCode/BeTheHeroBackend/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.Route(app)
	app.Listen(":3000")
}

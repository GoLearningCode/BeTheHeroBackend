package routes

import (
	"github.com/GoLearningCode/BeTheHeroBackend/controllers/ong"
	"github.com/GoLearningCode/BeTheHeroBackend/controllers/profile"
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	//app.Post("/sessions", session.Create)
	app.Get("/ongs", ong.Index)
	app.Post("/ongs", ong.Create)
	app.Get("/profile", profile.Index)
	//app.Get("/incidents", incidents.Index)
	//app.Post("/incidents", incidents.Create)
	//app.Delete("/incidents/:id", incidents.Delete)

}

package main

import "github.com/gofiber/fiber/v2"

func Route(app *fiber.App) {
	/*app.Post("/sessions", controllers.session.create)
	app.Get("/ongs", controllers.ong.index)
	app.Post("/ongs" controllers.ong.create)
	app.Get("/profile", controllers.profile.create)
	*/
	app.Get("/incidents", controllers.incidents.index)
	//app.Post("/incidents", controllers.incidents.create)
	//app.Delete("/incidents/:id", controllers.incidents.delete)

}

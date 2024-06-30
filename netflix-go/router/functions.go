package router

import (
	"server/modules/movies"
	"server/modules/users"

	"github.com/gofiber/fiber/v2"
)

func ConfigureRoute(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	users.Route(v1)
	movies.Route(v1)

}

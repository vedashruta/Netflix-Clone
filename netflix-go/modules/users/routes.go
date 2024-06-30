package users

import (
	"server/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

func Route(router fiber.Router) {
	users := router.Group("/users")

	users.Post("/create", timeout.NewWithContext(create, config.Config.Server.Timeout))
	users.Post("/getWatchlist", timeout.NewWithContext(getWatchlist, config.Config.Server.Timeout))
}

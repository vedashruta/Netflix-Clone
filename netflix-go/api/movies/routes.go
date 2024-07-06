package movies

import (
	"server/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

func Route(router fiber.Router) {
	movies := router.Group("/movies")

	movies.Post("/insertMovie", timeout.NewWithContext(insertMovie, config.Config.Server.Timeout))
	movies.Post("/getAllMovies", timeout.NewWithContext(getAllMovies, config.Config.Server.Timeout))
	movies.Post("/getMoviesByLanguage", timeout.NewWithContext(getMoviesByLanguage, config.Config.Server.Timeout))
	movies.Post("/getMovieById", timeout.NewWithContext(getMovieById, config.Config.Server.Timeout))
	movies.Post("/getLatestMovies", timeout.NewWithContext(getLatestMovies, config.Config.Server.Timeout))
	movies.Get("/search", timeout.NewWithContext(search, config.Config.Server.Timeout))
}

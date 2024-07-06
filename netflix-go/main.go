package main

import (
	"log"
	"server/api/movies"
	"server/api/users"
	"server/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatalln(err)
	}
	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")
	users.Route(v1)
	movies.Route(v1)
	err = app.Listen(config.Config.Server.Port)
	if err != nil {
		log.Fatalln(err)
	}
}

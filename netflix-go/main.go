package main

import (
	"context"
	"log"
	"server/config"
	"server/router"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatalln(err)
	}
	app := fiber.New(fiber.Config{
		StrictRouting: true,
		CaseSensitive: true,
	})
	router.ConfigureRoute(app)
	app.Get("/test", func(c *fiber.Ctx) error {
		id, _ := primitive.ObjectIDFromHex("6680cc69ed36e27f3e9d3cca")
		filter := bson.D{
			{
				Key:   "_id",
				Value: id,
			},
		}
		var res any
		err = config.Movies.FindOne(context.TODO(), filter).Decode(&res)
		if err != nil {
			c.SendString("error")
		}
		return c.JSON(res)
	})
	err = app.Listen(config.Config.Server.Port)
	if err != nil {
		log.Fatalln(err)
	}
}

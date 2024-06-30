package users

import (
	"context"
	"server/config"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func create(c *fiber.Ctx) (err error) {
	var payload map[string]interface{}
	err = c.BodyParser(&payload)
	if err != nil {
		return c.SendString("err : no payload")
	}
	email, ok := payload["email"]
	if !ok {
		return c.SendString("err : email key not found")
	}
	response := map[string]interface{}{
		"email": email,
	}
	return c.JSON(response)
}

func getWatchlist(c *fiber.Ctx) (err error) {
	var payload map[string]interface{}
	err = c.BodyParser(&payload)
	if err != nil {
		return c.SendString("err : no payload")
	}
	userID, ok := payload["user_id"]
	if !ok {
		return c.SendString("key : user_id not found")
	}
	profileID, ok := payload["profile_id"]
	if !ok {
		return c.SendString("key : user_id not found")
	}
	skip, ok := payload["skip"]
	if !ok {
		return c.SendString("key : genre not found")
	}
	filter := bson.D{
		{
			Key:   "_id",
			Value: userID.(primitive.ObjectID),
		},
		{
			Key:   "profile.id",
			Value: profileID.(primitive.ObjectID),
		},
	}
	limit := 20
	opts := options.Find()
	skip = int64((skip.(int64) - 1) * int64(limit))
	opts.SetSkip(skip.(int64))
	ctx, cancel := context.WithTimeout(context.Background(), config.Config.Database.Timeout)
	defer cancel()
	cursor, err := config.Movies.Find(ctx, filter, opts)
	if err != nil {
		return c.JSON(err)
	}
	defer cursor.Close(ctx)
	var movies []Model
	for cursor.Next(ctx) {
		var movie Model
		if err = cursor.Decode(&movie); err != nil {
			return c.JSON(err)
		}
		movies = append(movies, movie)
	}
	if err != nil {
		return c.JSON(err)
	}
	return c.JSON(movies)
}

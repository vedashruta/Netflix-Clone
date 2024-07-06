package movies

import (
	"context"
	"encoding/json"
	"server/config"
	"server/packages"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func insertMovie(c *fiber.Ctx) (err error) {
	var payload []interface{}
	bytes := c.Body()
	err = json.Unmarshal(bytes, &payload)
	if err != nil {
		return c.SendString("err : no payload")
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.Config.Database.Timeout)
	defer cancel()
	_, err = config.Movies.InsertMany(ctx, payload)
	if err != nil {
		return c.JSON(err)
	}
	return
}

func getAllMovies(c *fiber.Ctx) (err error) {
	var payload map[string]interface{}
	err = c.BodyParser(&payload)
	if err != nil {
		return c.SendString("err : no payload")
	}
	movieType, ok := payload["type"]
	if !ok {
		return c.SendString("key : type not found")
	}
	genre, ok := payload["genre"]
	if !ok {
		return c.SendString("key : genre not found")
	}
	skip, ok := payload["skip"]
	if !ok {
		return c.SendString("key : genre not found")
	}
	skipValue := packages.GetInt64(skip)
	res, err := getMovies(movieType.(string), genre.(string), skipValue)
	if err != nil {
		return c.JSON(err)
	}
	return c.JSON(res)
}

func getMoviesByLanguage(c *fiber.Ctx) (err error) {
	var payload map[string]interface{}
	err = c.BodyParser(&payload)
	if err != nil {
		return c.SendString("err : no payload")
	}
	language, ok := payload["type"]
	if !ok {
		return c.SendString("key : type not found")
	}
	skip, ok := payload["skip"]
	if !ok {
		return c.SendString("key : genre not found")
	}
	skipValue := packages.GetInt64(skip)
	res, err := getByLanguauge(language.(string), skipValue)
	if err != nil {
		return c.JSON(err)
	}
	return c.JSON(res)
}

func getMovieById(c *fiber.Ctx) (err error) {
	var payload map[string]interface{}
	err = c.BodyParser(&payload)
	if err != nil {
		return c.SendString("err : no payload")
	}
	id, ok := payload["_id"]
	if !ok {
		return c.SendString("key : _id not found")
	}
	res := &Model{}
	filter := bson.D{
		{
			Key:   "_id",
			Value: id.(primitive.ObjectID),
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.Config.Database.Timeout)
	defer cancel()
	err = config.Movies.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return c.JSON(err)
	}
	return c.JSON(res)
}

func getLatestMovies(c *fiber.Ctx) (err error) {
	var payload map[string]interface{}
	err = c.BodyParser(&payload)
	if err != nil {
		return c.SendString("err : no payload")
	}
	skip, ok := payload["skip"]
	if !ok {
		return c.SendString("key : skip not found")
	}
	skipValue := packages.GetInt64(skip)
	res, err := getLatest(skipValue)
	if err != nil {
		return c.JSON(err)
	}
	return c.JSON(res)
}

func search(c *fiber.Ctx) (err error) {
	query := c.Query("query")
	filter := bson.D{
		{
			Key: "title",
			Value: bson.D{
				{
					Key:   "$regex",
					Value: query,
				},
				{
					Key:   "$options",
					Value: "i",
				},
			},
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.Config.Database.Timeout)
	defer cancel()
	cursor, err := config.Movies.Find(ctx, filter)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	movies := []Model{}
	for cursor.Next(ctx) {
		var movie Model
		if err = cursor.Decode(&movie); err != nil {
			return
		}
		movies = append(movies, movie)
	}
	return c.JSON(movies)
}

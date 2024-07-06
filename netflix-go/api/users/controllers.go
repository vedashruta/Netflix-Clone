package users

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"server/config"
	"server/packages"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func create(c *fiber.Ctx) (err error) {
	var payload map[string]interface{}
	err = c.BodyParser(&payload)
	if err != nil {
		return c.SendString("err : no payload")
	}
	tempEmail, ok := payload["email"]
	if !ok {
		return c.SendString("err : email key not found")
	}
	tempPass, ok := payload["password"]
	if !ok {
		return c.SendString("err : password key not found")
	}
	email := packages.GetString(tempEmail)
	password := packages.GetString(tempPass)
	filter := bson.D{
		{
			Key:   "email",
			Value: email,
		},
	}
	opts := options.Find()
	limit := int64(1)
	opts.SetLimit(limit)
	ctx, cancel := context.WithTimeout(context.Background(), config.Config.Database.Timeout)
	defer cancel()
	var res Model
	err = config.Users.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return c.SendString("err : email already exists")

		}
	}
	enc, err := config.Encrypt([]byte(password), "salt")
	if err != nil {
		c.JSON(err)
	}
	password = hex.EncodeToString(enc)
	user := Model{
		Email:    email,
		Password: password,
	}
	_, err = config.Users.InsertOne(ctx, user)
	if err != nil {
		c.JSON(err)
	}
	return
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

func update(c *fiber.Ctx) (err error) {
	var payload map[string]interface{}
	err = c.BodyParser(&payload)
	if err != nil {
		return c.SendString("err : no payload")
	}
	tempEmail, ok := payload["email"]
	if !ok {
		return c.SendString("key : email not found")
	}
	updateEmail, ok := payload["updateEmail"]
	if !ok {
		return c.SendString("key : email not found")
	}
	tempName, ok := payload["name"]
	if !ok {
		return c.SendString("object : name not found")
	}
	bytes, err := json.Marshal(tempName)
	if err != nil {
		c.JSON(err)
	}
	name := Name{}
	err = json.Unmarshal(bytes, name)
	if err != nil {
		c.JSON(err)
	}
	tempPhone, ok := payload["phone"]
	if !ok {
		return c.SendString("key : phone not found")
	}
	filter := bson.D{
		{
			Key:   "email",
			Value: packages.GetString(tempEmail),
		},
	}
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{
					Key:   "email",
					Value: packages.GetString(updateEmail),
				},
				{
					Key:   "name",
					Value: name,
				},
				{
					Key:   "phone",
					Value: packages.GetInt64(tempPhone),
				},
			},
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.Config.Database.Timeout)
	defer cancel()
	res := config.Users.FindOneAndUpdate(ctx, filter, update)
	return c.JSON(res)
}

func delete(c *fiber.Ctx) (err error) {
	var payload map[string]interface{}
	err = c.BodyParser(&payload)
	if err != nil {
		return c.SendString("err : no payload")
	}
	tempEmail, ok := payload["email"]
	if !ok {
		return c.SendString("key : email not found")
	}
	filter := bson.D{
		{
			Key:   "email",
			Value: packages.GetString(tempEmail),
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.Config.Database.Timeout)
	defer cancel()
	res := config.Users.FindOneAndDelete(ctx, filter)
	return c.JSON(res)
}

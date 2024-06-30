package config

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v2"
)

func Init() (err error) {
	fileName := "env/.yaml"
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	err = yaml.NewDecoder(file).Decode(&Config)
	if err != nil {
		return
	}
	Db, err = mongoConnect()
	if err != nil {
		return
	}
	Users = Db.Collection("users")
	Movies = Db.Collection("movies")
	return
}

func mongoConnect() (db *mongo.Database, err error) {
	clientOptions := options.Client().ApplyURI(Config.Database.URI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return
	}
	db = client.Database(Config.Database.Db)
	return
}

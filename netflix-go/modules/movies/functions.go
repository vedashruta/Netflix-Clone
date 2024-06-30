package movies

import (
	"context"
	"server/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getMovies(movieType string, genre string, skip int64) (movies []Model, err error) {
	filter := bson.D{
		{
			Key:   "type",
			Value: movieType,
		},
		{
			Key:   "genre",
			Value: genre,
		},
	}
	limit := 20
	opts := options.Find()
	skip = int64((skip - 1)) * int64(limit)
	opts.SetSkip(skip)
	ctx, cancel := context.WithTimeout(context.Background(), config.Config.Database.Timeout)
	defer cancel()
	cursor, err := config.Movies.Find(ctx, filter, opts)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	movies = []Model{}
	for cursor.Next(ctx) {
		var movie Model
		if err = cursor.Decode(&movie); err != nil {
			return
		}
		movies = append(movies, movie)
	}
	return
}

func getByLanguauge(language string, skip int64) (movies []Model, err error) {
	filter := bson.D{
		{
			Key: "language",
			Value: bson.D{
				{
					Key: "$in",
					Value: bson.A{
						language,
					},
				},
			},
		},
	}
	limit := 20
	opts := options.Find()
	skip = int64((skip - 1) * int64(limit))
	opts.SetSkip(skip)
	ctx, cancel := context.WithTimeout(context.Background(), config.Config.Database.Timeout)
	defer cancel()
	cursor, err := config.Movies.Find(ctx, filter, opts)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	movies = []Model{}
	for cursor.Next(ctx) {
		var movie Model
		if err = cursor.Decode(&movie); err != nil {
			return
		}
		movies = append(movies, movie)
	}
	return
}

func getLatest(skip int64) (movies []Model, err error) {
	filter := bson.D{}
	opts := options.Find()
	limit := 20
	skip = int64((skip - 1) * int64(limit))
	sort := bson.D{
		{
			Key:   "_id",
			Value: -1,
		},
	}
	opts.SetSort(sort)
	opts.SetSkip(skip)
	ctx, cancel := context.WithTimeout(context.Background(), config.Config.Database.Timeout)
	defer cancel()
	cursor, err := config.Movies.Find(ctx, filter, opts)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	movies = []Model{}
	for cursor.Next(ctx) {
		var movie Model
		if err = cursor.Decode(&movie); err != nil {
			return
		}
		movies = append(movies, movie)
	}
	return
}

package movies

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Type        string             `json:"type" bson:"type"`
	Cast        []string           `json:"cast" bson:"cast"`
	DirectorID  primitive.ObjectID `json:"director_id" bson:"director_id"`
	Year        int                `json:"year" bson:"year"`
	Genre       []string           `json:"genre" bson:"genre"`
	Language    []string           `json:"language" bson:"language"`
	Description string             `json:"description" bson:"description"`
	Runtime     string             `json:"runtime" bson:"runtime"`
	Rating      float64            `json:"rating" bson:"rating"`
	Thumbnail   string             `json:"thumbnail" bson:"thumbnail"`
	Image       string             `json:"image" bson:"image"`
	TrailerUrl  string             `json:"trailerUrl" bson:"trailerUrl"`
}

type Director struct {
	ID     primitive.ObjectID   `json:"_id" bson:"_id"`
	Name   string               `json:"name" bson:"name"`
	Movies []primitive.ObjectID `json:"movies" bson:"movies"`
}

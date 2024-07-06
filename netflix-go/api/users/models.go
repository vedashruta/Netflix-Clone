package users

import "go.mongodb.org/mongo-driver/bson/primitive"

type Model struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     Name               `json:"name" bson:"name"`
	Phone    int                `json:"phone" bson:"phone"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
	Profiles []Profile          `json:"profile" bson:"profile"`
	Billing  Billing            `json:"billing" bson:"billing"`
}

type Name struct {
	First string `json:"first" bson:"first"`
	Last  string `json:"last" bson:"last"`
}

type Profile struct {
	ID           primitive.ObjectID   `json:"id" bson:"id"`
	ProfileName  string               `json:"profileName" bson:"profileName"`
	Avatar       string               `json:"avatar" bson:"avatar"`
	Watchlist    []primitive.ObjectID `json:"watchlist" bson:"watchlist"`
	Watchhistory []primitive.ObjectID `json:"watchhistory" bson:"watchhistory"`
	Liked        []primitive.ObjectID `json:"liked" bson:"liked"`
	Disliked     []primitive.ObjectID `json:"disliked" bson:"disliked"`
}

type Billing struct {
	ID          primitive.ObjectID `json:"id" bson:"id"`
	Plan        Plan               `json:"plan" bson:"plan"`
	PurchasedOn string             `json:"purchasedOn" bson:"purchasedOn"`
	CardNumber  string             `json:"cardNumber" bson:"cardNumber"`
}

type Plan struct {
	Name  string  `json:"name" bson:"name"`
	Price float64 `json:"price" bson:"price"`
}

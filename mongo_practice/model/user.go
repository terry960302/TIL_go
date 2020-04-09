package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id    primitive.ObjectID `bson:"id"`
	Name  string             `bson:"name"`
	Email string             `bson:"email"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name" validate:"required,min=3,max=50"`
	Email    string             `json:"email" validate:"email,required"`
	Phone    string             `json:"phone" validate:"required"`
	Password string             `json:"Password" validate:"required,min=6"`
}

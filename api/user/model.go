package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id primitive.ObjectID `json:"-" bson:"_id"`
	Nickname string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password" bson:"-"`
	PasswordHash []byte `json:"-" bson:"passHash"`
	PasswordSalt []byte `json:"-" bson:"passSalt"`
}
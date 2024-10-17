package model

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UID      primitive.ObjectID	`json:"uid,omitempty" bson:"uid,omitempty"`
	IDRole   uint	`json:"idrole,omitempty" bson:"idrole,omitempty"`
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
	Phone    string `json:"phone,omitempty" bson:"phone,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Address  string `json:"address,omitempty" bson:"address,omitempty"`
}

type Role struct {
	IDRole uint	`json:"idrole,omitempty" bson:"idrole,omitempty"`
	Role   string `json:"role,omitempty" bson:"role,omitempty"`
}

type JWTClaims struct {
	jwt.StandardClaims
	UID      uint	`json:"uid,omitempty" bson:"uid,omitempty"`
	IDRole   uint	`json:"idrole,omitempty" bson:"idrole,omitempty"`
}
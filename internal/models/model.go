package models

import "github.com/dgrijalva/jwt-go"

type User struct {
	UserID   string `json:"userID" bson:"userID"`
	Name     string `json:"name" bson:"name"`
	Surname  string `json:"surname" bson:"surname"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Type     string `json:"type" bson:"type"`
}

type Login struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Auth struct {
	Token string `json:"token"`
}

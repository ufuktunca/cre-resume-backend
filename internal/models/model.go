package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	UserID       string `json:"userID" bson:"userID"`
	Name         string `json:"name" bson:"name"`
	Surname      string `json:"surname" bson:"surname"`
	Email        string `json:"email" bson:"email"`
	Password     string `json:"password" bson:"password"`
	Type         string `json:"type" bson:"type"`
	UserActivate bool   `json:"userActivate" bson:"userActivate"`
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

type JobPost struct {
	ID        string    `json:"id" bson:"id"`
	Title     string    `json:"title" bson:"title"`
	Company   string    `json:"company" bson:"company"`
	Content   string    `json:"content" bson:"content"`
	Salary    int       `json:"salary" bson:"salary"`
	Category  string    `json:"category" bson:"category"`
	Location  string    `json:"location" bson:"location"`
	Image     string    `json:"image" bson:"image"`
	Type      string    `json:"type" bson:"type"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt" bson:"deletedAt"`
}

var RegistirationMailContent = "You have successfully registered to Cre-Resume please click this link for activation.\n "

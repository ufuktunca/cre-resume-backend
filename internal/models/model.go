package models

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	UserID     string `json:"userID" bson:"userID"`
	Name       string `json:"name" bson:"name"`
	Surname    string `json:"surname" bson:"surname"`
	Email      string `json:"email" bson:"email"`
	Password   string `json:"password" bson:"password"`
	Type       string `json:"type" bson:"type"`
	Activation bool   `json:"activation" bson:"activation"`
}

type Login struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type ReSend struct {
	Email string `json:"email"`
}

type Claims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}

type Auth struct {
	Token string `json:"token"`
}

type JobPost struct {
	ID        string    `json:"id" bson:"id"`
	OwnerID   string    `json:"ownerId" bson:"ownerId"`
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

type CV struct {
	ID          string       `json:"id" bson:"id"`
	OwnerID     string       `json:"ownerId" bson:"ownerId"`
	CVName      string       `json:"cvName" bson:"cvName"`
	NameSurname string       `json:"nameSurname" bson:"nameSurname"`
	AboutMe     string       `json:"aboutMe" bson:"aboutMe"`
	JobTitle    string       `json:"jobTitle" bson:"jobTitle"`
	PhoneNumber string       `json:"phoneNumber" bson:"phoneNumber"`
	Email       string       `json:"email" bson:"email"`
	Hobbies     []string     `json:"hobbies" bson:"hobbies"`
	Photo       string       `json:"photo" bson:"photo"`
	Education   []Education  `json:"education" bson:"education"`
	Experience  []Experience `json:"experience" bson:"experience"`
	Github      string       `json:"github" bson:"github"`
	Linkedin    string       `json:"linkedin" bson:"linkedin"`
	OtherSM     []string     `json:"otherSM" bson:"otherSM"`
	PDFCV       string       `json:"pdfCV" bson:"pdfCV"`
	Languages   []Language   `json:"languages" bson:"languages"`
	Skills      []string     `json:"skills" bson:"skilss"`
}

type Experience struct {
	StartDate   string `json:"startDate" bson:"startDate"`
	EndDate     string `json:"endDate" bson:"endDate"`
	Company     string `json:"company" bson:"company"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
}

type Education struct {
	School     string `json:"school" bson:"school"`
	Department string `json:"department" bson:"department"`
	StartDate  string `json:"startDate" bson:"startDate"`
	EndDate    string `json:"endDate" bson:"endDate"`
}

type Language struct {
	Language string `json:"language" bson:"language"`
	Level    string `json:"level" bson:"level"`
}

type ApplyJobPost struct {
	ID          string `json:"id" bson:"id"`
	JobPostID   string `json:"jobPostId" bson:"jobPostId"`
	CVID        string `json:"cvId" bson:"cvId"`
	ApplierID   string `json:"applierId" bson:"applierId"`
	PostOwnerID string `json:"postOwnerId" bson:"postOwnerId"`
}

type ApplyJobPostDTO struct {
	CVID string `json:"cvId" bson:"cvId"`
}

var RegistirationMailContent = "You have successfully registered to Cre-Resume please click this link for activation.\n "
var ActivationError = errors.New("User is not activated")
var EmailError = errors.New("This email address already used")

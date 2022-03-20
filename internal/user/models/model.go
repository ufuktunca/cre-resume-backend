package models

type User struct {
	UserID   string `json:"userID" bson:"userID"`
	Name     string `json:"name" bson:"name"`
	Surname  string `json:"surname" bson:"surname"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Type     string `json:"type" bson:"type"`
}

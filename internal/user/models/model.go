package models

type User struct {
	UserID   string `json:"userID"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

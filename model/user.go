package model

type User struct {
	ID       int    `json:"id" bson:"id"`
	Email    string `json:"-" bson:"email"`
	Username string `json:"username" bson:"username"`
}

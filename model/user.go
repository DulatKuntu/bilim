package model

type User struct {
	ID       int    `json:"id" bson:"id"`
	Email    string `json:"email" bson:"email"`
	Username string `json:"username" bson:"username"`
	Name     string `json:"name" bson:"name"`
	Surname  string `json:"surname" bson:"surname"`
	Bio      string `json:"bio" bson:"bio"`
}

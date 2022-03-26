package model

type BuddyPost struct {
	ID          int    `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Surname     string `json:"surname" bson:"surname"`
	Description string `json:"description" bson:"description"`
}

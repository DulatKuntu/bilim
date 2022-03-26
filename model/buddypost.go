package model

type RequestBuddyPost struct {
	UserID      int    `json:"userID" bson:"userID"`
	Name        string `json:"name" bson:"name"`
	Surname     string `json:"surname" bson:"surname"`
	Description string `json:"description" bson:"description"`
}

type BuddyPost struct {
	ID          int    `json:"id" bson:"id"`
	UserID      int    `json:"userID" bson:"userID"`
	Name        string `json:"name" bson:"name"`
	Surname     string `json:"surname" bson:"surname"`
	Description string `json:"description" bson:"description"`
}

package model

type User struct {
	ID        int    `json:"id" bson:"id"`
	Email     string `json:"email" bson:"email"`
	Username  string `json:"username" bson:"username"`
	Name      string `json:"name" bson:"name"`
	Surname   string `json:"surname" bson:"surname"`
	Password  string `json:"password" bson"password"`
	Bio       string `json:"bio" bson:"bio"`             //mentor or student
	Interests string `json;"interests" bson:"interests"` //for example, Biology, CS, Engineering
}

/*type Mentor struct {
	ID       int    `json:"id" bson:"id"`
	Email    string `json:"email" bson:"email"`
	Username string `json:"username" bson:"username"`
	Name     string `json:"name" bson:"name"`
	Surname  string `json:"surname" bson:"surname"`
	Bio      string `json:"bio" bson:"bio"`
}*/

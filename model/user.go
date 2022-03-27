package model

type User struct {
	ID        int      `json:"id" bson:"id"`
	Email     string   `json:"email" bson:"email"`
	Token     string   `json:"token" bson:"token"`
	Username  string   `json:"username" bson:"username"`
	Name      string   `json:"name" bson:"name"`
	Surname   string   `json:"surname" bson:"surname"`
	Password  string   `json:"password" bson:"password"`
	Image     string   `json:"image" bson:"image"`
	Bio       string   `json:"bio" bson:"bio"`             //mentor or student
	Interests []int    `json:"interests" bson:"interests"` //for example, Biology, CS, Engineering
	Groups    []string `json:"groups" bson:"groups"`
}

type RequestUser struct {
	Email     string   `json:"email" bson:"email"`
	Username  string   `json:"username" bson:"username"`
	Password  string   `json:"password" bson:"password"`
	Name      string   `json:"name" bson:"name"`
	Surname   string   `json:"surname" bson:"surname"`
	Bio       string   `json:"bio" bson:"bio"`
	Interests []int    `json:"interests" bson:"interests"` //for example, Biology, CS, Engineering
	Groups    []string `json:"groups" bson:"groups"`
}

type Mentor struct {
	ID        int    `json:"id" bson:"id"`
	Email     string `json:"email" bson:"email"`
	Token     string `json:"token" bson:"token"`
	Username  string `json:"username" bson:"username"`
	Password  string `json:"password" bson:"password"`
	Name      string `json:"name" bson:"name"`
	Surname   string `json:"surname" bson:"surname"`
	Bio       string `json:"bio" bson:"bio"`
	Image     string `json:"image" bson:"image"`
	Interests []int  `json:"interests" bson:"interests"`
	Students  []int  `json:"students" bson:"students"`
	Pending   []int  `json:"pending" bson:"pending"`
}

type RequestMentor struct {
	Email     string `json:"email" bson:"email"`
	Username  string `json:"username" bson:"username"`
	Password  string `json:"password" bson:"password"`
	Name      string `json:"name" bson:"name"`
	Surname   string `json:"surname" bson:"surname"`
	Bio       string `json:"bio" bson:"bio"`
	Interests []int  `json:"interests" bson:"interests"`
	Students  []int  `json:"students" bson:"students"`
	Pending   []int  `json:"pending" bson:"pending"`
}

type LoginInfo struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type LoginMentor struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type Interests struct {
	InterestID int    `json:"id" bson:"id"`
	Name       string `json:"name" bson:"name"`
}

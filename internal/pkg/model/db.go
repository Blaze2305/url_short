package model

// Shorten struct is the basic data structure for the url document
type Shorten struct {
	Token   string `json:"token" bson:"_id"`
	Forward string `json:"forward" bson:"forward"`
	Created string `json:"created" bson:"created"`
	User    string `json:"user" bson:"user"`
}

// User struct defines the user model structure
type User struct {
	ID       string `json:"id" bson:"_id"`
	UserName string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Salt     string `json:"salt" bson:"salt"`
}

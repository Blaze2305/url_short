package model

// Shorten struct is the basic data structure for the url document
type Shorten struct {
	Token   string `json:"token" bson:"_id"`
	Forward string `json:"forward" bson:"forward"`
	Created string `json:"created" bson:"created"`
}

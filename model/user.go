package model

type User struct {
	UserName string `bson:"username"`
	ID       string `bson:"_id,omitempty"`
	Email    string `bson:"email"`
}

type UserRequest struct {
	UserName string `json:"username"`
	ID       string `json:"id,omitempty"`
	Email    string `json:"email"`
}

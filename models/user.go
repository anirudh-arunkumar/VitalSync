package models

type User struct {
	Username string `json:"username" bson:"username`
	Password string `json:"password" bson:"password`
	Email string `json:"email" bson:"email`
	ID string `json:"id,omitempty" bson:"id,omitempty`
}
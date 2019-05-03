package models

type User struct {
	ID       int32  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

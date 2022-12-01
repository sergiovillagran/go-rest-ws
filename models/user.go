package models

type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"Email"`
	Password string `json:"password"`
}

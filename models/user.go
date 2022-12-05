package models

type User struct {
	Id       string `json:"id"`
	Email    string `json:"Email"`
	Password string `json:"password"`
}

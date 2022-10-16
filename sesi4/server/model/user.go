package model

import "time"

type BaseModel struct {
	Id        int `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
type User struct {
	BaseModel
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Age      string
	Contact  string
	Address  string
	Photo    string
	Password string `json:"password"`
}

type Product struct {
	BaseModel
	Name  string
	Stock int
}

var Users = []User{}

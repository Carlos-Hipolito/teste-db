package models

import "gorm.io/gorm"

type Users []User

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

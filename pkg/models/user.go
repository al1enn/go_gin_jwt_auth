package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}
type LoginStruct struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

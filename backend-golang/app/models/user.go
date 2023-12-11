package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int `gorm:"primarykey"`
	FirstName string
	LastName  string
	Username  string
	Password  string
	Phone     string
	Email     string
}

package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int `gorm:"primarykey"`
	FirstName string
	LastName  string
	Username  string `gorm:"unique_index;not null"`
	Password  string
	Phone     string
	Email     string
}

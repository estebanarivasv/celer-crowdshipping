package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int `gorm:"primarykey"`
	Name     string
	Surname  string
	Username string
	// Password string
	Phone    string
	Email    string
	BankName string
	BankCBU  string
}

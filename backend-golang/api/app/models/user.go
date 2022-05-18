package models

import "gorm.io/gorm"

// User TODO: Field-Level Permission gorm
type User struct {
	gorm.Model
	ID       int    `gorm:"primarykey" json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	BankName string `json:"bank_name"`
	BankCBU  string `json:"bank_CBU"`
}

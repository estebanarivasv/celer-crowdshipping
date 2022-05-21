package entities

import (
	"time"
)

type UserInDTO struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	BankName string `json:"bank_name"`
	BankCBU  string `json:"bank_CBU"`
}

type UserOutDTO struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Username  string    `json:"username"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	BankName  string    `json:"bank_name"`
	BankCBU   string    `json:"bank_CBU"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

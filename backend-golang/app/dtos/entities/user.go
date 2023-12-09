package entities

type UserInDTO struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	BankName string `json:"bankName"`
	BankCBU  string `json:"bankCbu"`
}

type UserOutDTO struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	BankName string `json:"bankName"`
	BankCBU  string `json:"bankCbu"`
}

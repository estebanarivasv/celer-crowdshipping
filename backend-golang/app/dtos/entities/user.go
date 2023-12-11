package entities

type UserInDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

type UserOutDTO struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

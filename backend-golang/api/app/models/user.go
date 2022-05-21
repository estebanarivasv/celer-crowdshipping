package models

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"gorm.io/gorm"
)

// User TODO: Field-Level Permission gorm
type User struct {
	gorm.Model
	ID       int    `gorm:"primarykey" json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	// Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	BankName string `json:"bank_name"`
	BankCBU  string `json:"bank_CBU"`
}

func (model User) FromDTO(dto entities.UserInDTO) interface{} {

	return User{
		Name:     dto.Name,
		Surname:  dto.Surname,
		Username: dto.Username,
		Phone:    dto.Phone,
		Email:    dto.Email,
		BankName: dto.BankName,
		BankCBU:  dto.BankCBU,
	}

}

func (model User) ToDTO() entities.UserOutDTO {

	return entities.UserOutDTO{
		ID:        model.ID,
		Name:      model.Name,
		Surname:   model.Surname,
		Username:  model.Username,
		Phone:     model.Phone,
		Email:     model.Email,
		BankName:  model.BankName,
		BankCBU:   model.BankCBU,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		DeletedAt: model.DeletedAt.Time,
	}

}

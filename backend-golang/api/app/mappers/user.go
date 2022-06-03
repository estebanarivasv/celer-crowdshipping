package mappers

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/models"
)

type UserMapper struct {
}

// FromDTO Return a model with the dto that comes from the service
func (m UserMapper) FromDTO(dto *entities.UserInDTO) models.User {

	return models.User{
		Name:     dto.Name,
		Surname:  dto.Surname,
		Username: dto.Username,
		Phone:    dto.Phone,
		Email:    dto.Email,
		BankName: dto.BankName,
		BankCBU:  dto.BankCBU,
	}

}

// ToDTO Return a dto with the model that comes from the database
func (m UserMapper) ToDTO(model *models.User) interface{} {

	// If model is empty, it returns an empty interface
	if model.ID == 0 {
		return nil
	}

	return entities.UserOutDTO{
		ID:       model.ID,
		Name:     model.Name,
		Surname:  model.Surname,
		Username: model.Username,
		Phone:    model.Phone,
		Email:    model.Email,
		BankName: model.BankName,
		BankCBU:  model.BankCBU,
		//CreatedAt:   model.CreatedAt,
		//UpdatedAt:   model.UpdatedAt,
		//DeletedAt:   model.DeletedAt.Time,
	}

}

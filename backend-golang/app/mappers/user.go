package mappers

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/models"
)

type UserMapper struct {
}

// NewUserMapper Returns a new instance
func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

// FromDTO Return a model with the dto that comes from the service
func (m UserMapper) FromDTO(dto *entities.UserInDTO) models.User {

	return models.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Username:  dto.Username,
		Password:  dto.Password,
		Phone:     dto.Phone,
		Email:     dto.Email,
	}

}

// ToDTO Return a dto with the model that comes from the database
func (m UserMapper) ToDTO(model *models.User) interface{} {

	// If model is empty, it returns an empty interface
	if model.ID == 0 {
		return nil
	}

	return entities.UserOutDTO{
		ID:        model.ID,
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Username:  model.Username,
		Phone:     model.Phone,
		Email:     model.Email,
		//CreatedAt:   model.CreatedAt,
		//UpdatedAt:   model.UpdatedAt,
		//DeletedAt:   model.DeletedAt.Time,
	}

}

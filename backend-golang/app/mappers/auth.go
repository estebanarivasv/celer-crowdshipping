package mappers

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/models"
)

type AuthMapper struct {
}

// NewAuthMapper Returns a new instance
func NewAuthMapper() *AuthMapper {
	return &AuthMapper{}
}

// FromLoginDTO Return a model with the dto that comes from the service
func (m AuthMapper) FromLoginDTO(dto *entities.UserLoginInDTO) models.User {

	return models.User{
		Username: dto.Username,
		Password: dto.Password,
	}

}

// ToDTO Return a dto with the model that comes from the generated JWT
func (m AuthMapper) ToDTO(token string) entities.JWTOutput {

	return entities.JWTOutput{
		Token: token,
	}

}

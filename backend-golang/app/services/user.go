package services

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/mappers"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
	userMapper     *mappers.UserMapper
}

// NewUserService Returns a new instance
func NewUserService() *UserService {

	var repo = repositories.NewUserRepository()
	var userMapper = mappers.NewUserMapper()

	return &UserService{
		userRepository: repo,
		userMapper:     userMapper,
	}
}

func (s *UserService) CreateUser(userInDTO *entities.UserInDTO) dtos.Response {

	// Convert the dto to an entity
	model := s.userMapper.FromDTO(userInDTO)

	query, err := s.userRepository.Create(model)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Mapping into a dto
	dto := s.userMapper.ToDTO(&query)

	return dtos.Response{Success: true, Data: dto}
}

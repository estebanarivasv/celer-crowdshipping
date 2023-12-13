package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/mappers"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/models"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/repositories"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type AuthService struct {
	userRepository *repositories.UserRepository
	authMapper     *mappers.AuthMapper
}

// NewUserService Returns a new instance
func NewAuthService() *AuthService {

	var repo = repositories.NewUserRepository()
	var authMapper = mappers.NewAuthMapper()

	return &AuthService{
		userRepository: repo,
		authMapper:     authMapper,
	}
}

func (s *AuthService) LoginUser(userInDTO *entities.UserLoginInDTO) dtos.Response {

	user, err := s.userRepository.FindOneByUsername(userInDTO.Username)

	pwdHashFromDB := []byte(user.Password)

	// Compares the hash pwd from DB to the pwd inserted via a DTO
	if err := bcrypt.CompareHashAndPassword(pwdHashFromDB, []byte(userInDTO.Password)); err != nil {
		return dtos.Response{Success: false, Error: "Wrong credentials"}
	}

	// Generate JWT token
	jwtDto, err := s.GenerateJWT(user.Username)
	if err != nil {
		return dtos.Response{Success: false, Error: "Error generating JWT token"}
	}

	// Mapping into a dto
	dto := s.authMapper.ToDTO(jwtDto)

	return dtos.Response{Success: true, Data: dto}
}

func (s *AuthService) GenerateJWT(username string) (string, error) {

	jwtKey := os.Getenv("SECRET_JWT_KEY")
	var secretKey = []byte(jwtKey)

	claims := entities.JWTClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *AuthService) IsJWTValid(tokenFromHeader string) (bool, *jwt.Token, error) {

	jwtKey := os.Getenv("SECRET_JWT_KEY")
	var secretKey = []byte(jwtKey)

	// Validate the token
	token, err := jwt.Parse(tokenFromHeader, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		if err != nil {
			return false, token, err
		} else {
			return false, token, errors.New("invalid token")
		}
	}
	return true, token, nil
}

func (s *AuthService) GetUserFromClaims(token *jwt.Token) (*models.User, error) {

	// Access token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("error obtaining token claims")
	}

	// Access username from claims
	username, ok := claims["username"].(string)
	if !ok {
		return nil, errors.New("error obtaining token claims")
	}

	// Get user from DB
	user, err := s.userRepository.FindOneByUsername(username)
	if err != nil {
		return nil, errors.New("error obtaining the user data")
	}

	return &user, nil
}

package handlers

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/services"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/utils/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	authService *services.AuthService
	userService *services.UserService
}

// NewAuthController Returns a new instance
func NewAuthController() *AuthController {

	var authService = services.NewAuthService()
	var userService = services.NewUserService()

	return &AuthController{
		authService: authService,
		userService: userService,
	}
}

// Login
// @Summary Authenticates a user
// @Description Authenticates the user and returns a JWT if the credentials are correct
// @Consume application/json
// @Accept json
// @Produce json
// @Param Package body entities.UserLoginInDTO true "Fill the body to create a new user"
// @Success 200 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /auth/login [post]
func (c *AuthController) Login(context *gin.Context) {

	dto := controllers.ShouldBindDTO(context, entities.UserLoginInDTO{})

	var responseDto = c.authService.LoginUser(&dto)
	if !responseDto.Success {
		context.JSON(http.StatusBadRequest, responseDto)
		return
	}
	context.JSON(http.StatusOK, responseDto)
	return

}

// Register
// @Summary Creates a new user
// @Description Creates a new user in the database with a hashed password
// @Consume application/json
// @Accept json
// @Produce json
// @Param Package body entities.UserInDTO true "Fill the body to create a new user"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /auth/register [post]
func (c *AuthController) Register(context *gin.Context) {

	dto := controllers.ShouldBindDTO(context, entities.UserInDTO{})

	var responseDto = c.userService.CreateUser(&dto)
	if !responseDto.Success {
		context.JSON(http.StatusBadRequest, responseDto)
		return
	}
	context.JSON(http.StatusCreated, responseDto)
	return

}

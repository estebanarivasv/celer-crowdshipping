package middlewares

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthMiddleware struct {
	authService *services.AuthService
}

// NewAuthMiddleware Returns a new instance
func NewAuthMiddleware() *AuthMiddleware {

	var service = services.NewAuthService()

	return &AuthMiddleware{
		authService: service,
	}
}

func (m AuthMiddleware) UseMiddleware(c *gin.Context) {

	// Get token from Authentication header
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token not provided"})
		c.Abort()
		return
	}

	valid, parsedToken, err := m.authService.IsJWTValid(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token is not valid"})
		c.Abort()
		return
	}

	user, err := m.authService.GetUserFromClaims(parsedToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las reclamaciones del token"})
		c.Abort()
		return
	}

	// Store user in gin context for controllers to have this item accessible
	c.Set("user_id", user.ID)

}

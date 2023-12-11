package entities

import (
	"github.com/dgrijalva/jwt-go"
)

type UserLoginInDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type JWTOutput struct {
	Token string `json:"token"`
}

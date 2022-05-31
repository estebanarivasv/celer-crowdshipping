package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetCorsConfig() gin.HandlerFunc {
	// TODO: Edit cors configuration

	/*
		config := cors.Config{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
			AllowCredentials: false,
			MaxAge:           12 * time.Hour,
		}
		return cors.New(config)
	*/

	// This will allow all HTTP methods and origins,
	return cors.Default()
}

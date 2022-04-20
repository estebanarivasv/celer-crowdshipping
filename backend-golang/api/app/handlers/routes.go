package handlers

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	// Create a get method and associate it with the welcome funct
	router.GET("/", welcome)
}

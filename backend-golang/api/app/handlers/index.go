package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Hello world",
	})
	return
}

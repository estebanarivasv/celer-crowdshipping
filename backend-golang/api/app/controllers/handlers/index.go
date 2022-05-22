package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome to celer-crowdshipping API - Docs: http://localhost:5001/docs/swagger",
	})
	return
}

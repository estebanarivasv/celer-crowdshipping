package handlers

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func IndexHandler(c *gin.Context) {

	config.LoadEnv()
	address := os.Getenv("ADDR")

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome to celer-crowdshipping API - Docs: http://localhost" + address + "/docs/swagger/index.html",
	})
	return
}

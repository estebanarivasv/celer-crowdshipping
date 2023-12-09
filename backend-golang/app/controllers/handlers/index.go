package handlers

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type IndexController struct {
}

// NewIndexController Returns a new instance
func NewIndexController() *IndexController {
	return &IndexController{}
}

func (c *IndexController) IndexHandler(context *gin.Context) {

	config.LoadEnv()
	address := os.Getenv("ADDR")

	context.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome to celer-crowdshipping API - Docs: http://localhost" + address + "/docs/swagger/index.html",
	})
	return
}

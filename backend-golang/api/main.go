package main

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Route Handlers / Endpoints
	handlers.Routes(router)

	router.Run(":5000")
}

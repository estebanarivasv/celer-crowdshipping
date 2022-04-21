package main

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/config"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Connect to database
	config.ConnectToDb()

	// Load REST controllers
	handlers.Routes(router)

	router.Run(":5000")
}

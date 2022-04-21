package main

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/config"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/handlers"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Connect to database
	db := config.ConnectToDb()

	// AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes
	db.AutoMigrate(&models.Shipping{})

	// Load REST controllers
	handlers.Routes(router)

	router.Run(":5000")
}

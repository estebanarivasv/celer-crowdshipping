package main

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/config"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/controllers"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/models"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {

	config.LoadEnv()
	config.LoadSwaggerConfig()

	server := gin.Default()

	// Connect to database
	db := config.ConnectToDb()

	// AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes
	err := db.AutoMigrate(
		&models.Offer{},
		&models.Package{},
		&models.RouteDetail{},
		&models.Shipping{},
		&models.User{})
	if err != nil {
		panic(err)
		return
	}

	// Load REST handlers
	controllers.GenerateRouting(server)

	err = server.Run(os.Getenv("ADDR"))
	if err != nil {
		panic(err)
		return
	}
}

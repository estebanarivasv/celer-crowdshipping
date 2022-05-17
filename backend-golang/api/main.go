package main

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/config"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/handlers"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

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

	// Load REST controllers
	handlers.Routes(router)

	err = router.Run(":5001")
	if err != http.ErrServerClosed {
		panic("Server is closed")
		return
	}
}

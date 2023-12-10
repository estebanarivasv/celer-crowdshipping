package main

import (
	"fmt"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/config"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/controllers"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/models"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/services"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {

	config.LoadEnv()
	config.LoadSwaggerConfig()

	server := gin.Default()

	// Add cors configuration
	server.Use(config.GetCorsConfig())

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

	router := controllers.NewRouter()
	// Load REST handlers
	router.GenerateRouting(server)

	// Push Camunda workflow engine
	var camundaService = services.NewCamundaService()

	engineIsAlive, err := camundaService.CheckIfEngineIsAlive()
	if err != nil {
		fmt.Println("Error checking if engine is alive:", err)
		return
	}

	if engineIsAlive {
		_, err = camundaService.PushBPMNDiagram()
		if err != nil {
			fmt.Println("Error pushing BPMN diagram:", err)
			return
		}
	}

	err = server.Run(os.Getenv("ADDR"))
	if err != nil {
		panic(err)
		return
	}
}

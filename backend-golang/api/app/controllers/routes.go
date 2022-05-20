package controllers

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/controllers/handlers"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/controllers/handlers/distributor"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/controllers/handlers/sender"
	"github.com/estebanarivasv/Celer/backend-golang/api/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func GenerateRouting(server *gin.Engine) {
	// Create a get method and associate it with the welcome function
	// Todo delete these comments
	/*
		server.GET("/", indexHandler)
		server.POST("/sender/shippings", newShippingHandler)
		server.GET("/sender/shippings", getAllShippingsHandler)
		server.GET("/sender/shippings/:id", getShippingByIDHandler)
		server.PUT("/sender/shippings/:id", updateShippingByIDHandler)
		server.DELETE("/sender/shippings/:id", deleteShippingByIDHandler)
		server.POST("/camunda-proc/create", createShippingProc)
		server.POST("/camunda-proc/message", sendMessageToProc)
	*/

	// ex: localhost/api/v1
	mainPath := server.Group(docs.SwaggerInfo.BasePath)
	{
		senderPath := mainPath.Group("/sender")
		{
			senderShipPath := senderPath.Group("/shippings")
			{
				senderShipPath.POST("", sender.NewShipping)
				senderShipPath.GET("", sender.GetAllShippings)
				senderShipPath.GET("/:id", handlers.GetShippingByID)
				senderShipPath.PUT("/:id", sender.UpdateShippingByID)
				senderShipPath.PATCH("/:id", handlers.UpdateShippingStateByID)
				senderShipPath.DELETE("/:id", sender.DeleteShippingByID)
				senderShipPath.GET("/:id/offers", sender.GetShippingOffersByID)
				senderShipPath.GET("/:id/package", sender.GetShippingPackageByID)
				senderShipPath.GET("/:id/route", sender.GetShippingRouteByID)
			}
		}
		distributorPath := mainPath.Group("/distributor")
		{
			distributorReqPath := distributorPath.Group("/requests")
			{
				distributorReqPath.GET("", distributor.SearchRequests)
				distributorReqPath.POST("/:id/offers", distributor.NewRequestOffer)
				distributorReqPath.GET("/:id/offers/:offerId", distributor.GetOfferByID)
				distributorReqPath.DELETE("/:id/offers/:offerId", distributor.DeleteOfferByID)
			}
			distributorPath.GET("/offers", distributor.GetAllOffersByDistributor)
			distributorShipPath := distributorPath.Group("/shippings")
			{
				distributorShipPath.GET("/:id", handlers.GetShippingByID)
				distributorShipPath.PATCH("/:id", handlers.UpdateShippingStateByID)
				distributorShipPath.POST("/:id/route/coordinates", distributor.NewShippingCoordinate)
			}
		}
		mainPath.GET("/shippings/:id/state", handlers.GetShippingStateByID)
		mainPath.GET("/users/:id", handlers.GetUserByID)
	}

	server.NoRoute(handlers.IndexHandler)
	// Append swagger middleware
	server.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

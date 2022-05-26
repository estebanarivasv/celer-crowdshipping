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

	// ex: localhost/api/v1
	mainPath := server.Group(docs.SwaggerInfo.BasePath)
	{
		senderPath := mainPath.Group("/sender")
		{
			senderShipPath := senderPath.Group("/shippings")
			{
				senderShipPath.POST("", sender.NewShipping)
				senderShipPath.GET("", sender.GetAllShippings)
				senderShipPath.GET("/:id", sender.GetShippingByID)
				senderShipPath.PATCH("/:id", sender.UpdateShippingStateByID)
				senderShipPath.DELETE("/:id", sender.DeleteShippingByID)
				senderShipPath.GET("/:id/offers", sender.GetShippingOffersByID)
				senderShipPath.PATCH("/:id/offers/selected", sender.UpdateSelectedOfferByID)
				senderShipPath.GET("/:id/offers/selected", sender.GetSelectedOfferByID)
				senderShipPath.GET("/:id/route", sender.GetShippingRouteByID)
			}
		}
		distributorPath := mainPath.Group("/distributor")
		{
			distributorReqPath := distributorPath.Group("/requests")
			{
				distributorReqPath.GET("", distributor.SearchRequests)
				distributorReqPath.POST("/:id/offers", distributor.NewRequestOffer)
				distributorReqPath.GET("/:id/offers", distributor.GetRequestOffers)
			}
			distributorPath.GET("/offers", distributor.GetAllOffersByDistributor)
			distributorPath.DELETE("/offers/:id", distributor.DeleteOfferByID)
			distributorShipPath := distributorPath.Group("/shippings")
			{
				distributorShipPath.GET("", distributor.GetShippingsInProcess)
				distributorShipPath.GET("/:id", distributor.GetShippingByID)
				distributorShipPath.PATCH("/:id", distributor.UpdateShippingStateByID)
				distributorShipPath.POST("/:id/route", distributor.NewShippingCoordinates)
			}
		}
		mainPath.GET("/shippings/:id/state", handlers.GetShippingStateByID)
		mainPath.GET("/users/:id", handlers.GetUserByID)
		mainPath.POST("/packages", handlers.NewPackage)
		mainPath.DELETE("/packages/:id", handlers.DeletePackageByID)

	}

	server.NoRoute(handlers.IndexHandler)
	// Append swagger middleware
	server.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

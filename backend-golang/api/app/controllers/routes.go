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
				senderShipPath.POST("", sender.NewShipping)                  // TODO DONE
				senderShipPath.GET("", sender.GetAllShippings)               // TODO DONE
				senderShipPath.GET("/:id", sender.GetShippingByID)           // TODO DONE
				senderShipPath.PATCH("/:id", sender.UpdateShippingStateByID) // TODO DONE
				senderShipPath.DELETE("/:id", sender.DeleteShippingByID)     // TODO DONE
				senderShipPath.PATCH("/:id/offers/selected", sender.UpdateSelectedOfferByID)
				senderShipPath.GET("/:id/offers", sender.GetShippingOffersByID)
				senderShipPath.GET("/:id/distributor", sender.GetShippingDistributorByID)
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
				distributorReqPath.DELETE("/:id/offers/:offerId", distributor.DeleteOfferByID)
			}
			distributorPath.GET("/offers", distributor.GetAllOffersByDistributor)
			distributorShipPath := distributorPath.Group("/shippings")
			{
				distributorShipPath.GET("/:id", distributor.GetShippingByID)           // TODO DONE
				distributorShipPath.PATCH("/:id", distributor.UpdateShippingStateByID) // TODO DONE
				distributorShipPath.POST("/:id/route/coordinates", distributor.NewShippingCoordinate)
			}
		}
		mainPath.GET("/shippings/:id/state", handlers.GetShippingStateByID) // TODO DONE
		mainPath.GET("/users/:id", handlers.GetUserByID)
		mainPath.POST("/packages", handlers.NewPackage)
		mainPath.DELETE("/packages/:id", handlers.DeletePackageByID)

	}

	server.NoRoute(handlers.IndexHandler)
	// Append swagger middleware
	server.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

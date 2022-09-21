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

type Router struct {
}

// NewShippingController Returns a new instance
func NewRouter() *Router {

	return &Router{}
}

func (r *Router) GenerateRouting(server *gin.Engine) {

	dOfferController := distributor.NewOfferController()
	dRequestController := distributor.NewRequestController()
	dShippingController := distributor.NewDistShippingController()
	sShippingController := sender.NewSendShippingController()
	indexController := handlers.NewIndexController()
	packageController := handlers.NewPackageController()
	shippingController := handlers.NewShippingController()

	// Create a get method and associate it with the welcome function
	// ex: localhost/api/v1
	mainPath := server.Group(docs.SwaggerInfo.BasePath)
	{
		senderPath := mainPath.Group("/sender")
		{
			senderShipPath := senderPath.Group("/shippings")
			{
				senderShipPath.POST("", sShippingController.NewShipping)
				senderShipPath.GET("", sShippingController.GetAllShippings)
				senderShipPath.GET("/:id", sShippingController.GetShippingByID)
				senderShipPath.PATCH("/:id", sShippingController.UpdateShippingStateByID)
				senderShipPath.DELETE("/:id", sShippingController.DeleteShippingByID)
				senderShipPath.GET("/:id/offers", sShippingController.GetShippingOffersByID)
				senderShipPath.PATCH("/:id/offers/selected", sShippingController.UpdateSelectedOfferByID)
				senderShipPath.GET("/:id/offers/selected", sShippingController.GetSelectedOfferByID)
				senderShipPath.GET("/:id/route", sShippingController.GetShippingRouteByID)
			}
		}
		distributorPath := mainPath.Group("/distributor")
		{
			distributorReqPath := distributorPath.Group("/requests")
			{
				distributorReqPath.GET("", dRequestController.SearchRequests)
				distributorReqPath.POST("/:id/offers", dRequestController.NewRequestOffer)
				distributorReqPath.GET("/:id/offers", dOfferController.GetRequestOffers)
			}
			distributorPath.GET("/offers", dOfferController.GetAllOffersByDistributor)
			distributorPath.DELETE("/offers/:id", dOfferController.DeleteOfferByID)
			distributorShipPath := distributorPath.Group("/shippings")
			{
				distributorShipPath.GET("", dShippingController.GetShippingsInProcess)
				distributorShipPath.GET("/:id", dShippingController.GetShippingByID)
				distributorShipPath.PATCH("/:id", dShippingController.UpdateShippingStateByID)
				distributorShipPath.POST("/:id/route", dShippingController.NewShippingCoordinates)
			}
		}
		mainPath.GET("/shippings/:id/state", shippingController.GetShippingStateByID)

		//mainPath.GET("/users/:id", handlers.GetUserByID)

		mainPath.POST("/packages", packageController.NewPackage)
		mainPath.DELETE("/packages/:id", packageController.DeletePackageByID)

	}

	server.NoRoute(indexController.IndexHandler)
	// Append swagger middleware
	server.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

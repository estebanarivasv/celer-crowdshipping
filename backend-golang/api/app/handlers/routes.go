package handlers

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/handlers/sender"
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
		// TODO Finish methods
		senderPath := mainPath.Group("/sender")
		{
			senderShipPath := senderPath.Group("/shippings")
			{
				senderShipPath.POST("", sender.NewShippingHandler)
				senderShipPath.GET("", sender.GetAllShippingsHandler)
				senderShipPath.GET("/:id", sender.GetShippingByIDHandler)
				senderShipPath.PUT("/:id", sender.UpdateShippingByIDHandler)
				senderShipPath.PATCH("/:id")
				senderShipPath.DELETE("/:id", sender.DeleteShippingByIDHandler)
				senderShipPath.GET("/:id/offers")
				senderShipPath.GET("/:id/package")
				senderShipPath.GET("/:id/route")
			}
		}
		distributorPath := mainPath.Group("/distributor")
		{
			distributorReqPath := distributorPath.Group("/requests")
			{
				distributorReqPath.GET("")
				distributorReqPath.POST("/:id/offers")
				distributorReqPath.GET("/:id/offers/:offerId")
				distributorReqPath.DELETE("/:id/offers/:offerId")
			}
			distributorPath.GET("/offers")
			distributorShipPath := distributorPath.Group("/shippings")
			{
				distributorShipPath.GET("/:id")
				distributorShipPath.PATCH("/:id")
				distributorShipPath.POST("/:id/route/coordinates")
			}
		}
		mainPath.GET("/users/:id")
	}

	// Append middleware
	server.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

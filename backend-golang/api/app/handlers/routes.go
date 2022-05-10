package handlers

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	// Create a get method and associate it with the welcome function
	router.GET("/", indexHandler)
	router.POST("/sender/shippings", newShippingHandler)
	router.GET("/sender/shippings", getAllShippingsHandler)
	router.GET("/sender/shippings/:id", getShippingByIDHandler)
	router.PUT("/sender/shippings/:id", updateShippingByIDHandler)
	router.DELETE("/sender/shippings/:id", deleteShippingByIDHandler)
	// Todo delete these routes
	router.POST("/shipping-proc/create", createShippingProc)
	router.POST("/shipping-proc/message", sendMessageToProc)
}

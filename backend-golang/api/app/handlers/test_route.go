package handlers

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/camunda"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func sendMessageToProc(c *gin.Context) {
	var jsonMessage shipping.MessageToProc
	if err := c.ShouldBindJSON(&jsonMessage); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}
	err := services.SendMessage(jsonMessage.MessageName, jsonMessage.ProcessInstanceId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, http.NoBody)
	return
}

func createShippingProc(c *gin.Context) {
	var data, err = services.CreateShippingProcInstance()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, data)
	return
}

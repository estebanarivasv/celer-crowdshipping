package handlers

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/services"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/utils/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ShippingController struct {
	shippingService *services.ShippingService
}

// NewShippingController Returns a new instance
func NewShippingController() *ShippingController {

	var shippingService = services.NewShippingService()

	return &ShippingController{
		shippingService: shippingService,
	}
}

// GetShippingStateByID
// @Summary Get Shipping state
// @Description Get Shipping state making reference to Camunda instance
// @Consume application/json
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Shipping ID"
// @Success 200 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Failure 500 {object} dtos.Response
// @Router /shippings/{id}/state [get]
func (c *ShippingController) GetShippingStateByID(context *gin.Context) {
	id, err := controllers.ConvertParamToInt(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	dto := c.shippingService.FindShippingStateById(id)
	if !dto.Success {
		if dto.Error == "process ID not found" {
			context.JSON(http.StatusNotFound, dto)
			return
		}
		context.JSON(http.StatusInternalServerError, dto)
		return
	}

	context.JSON(http.StatusOK, dto)
	return
}

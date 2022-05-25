package handlers

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/services"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/utils/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetShippingStateByID
// @Summary Get Shipping state
// @Description Get Shipping state making reference to Camunda instance
// @Consume application/json
// @Accept json
// @Produce json
// @Param id path int true "Shipping ID"
// @Success 200 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Failure 500 {object} dtos.Response
// @Router /shippings/{id}/state [get]
func GetShippingStateByID(c *gin.Context) {
	id, err := controllers.ConvertParamToInt(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	dto := services.FindShippingStateById(id)
	if !dto.Success {
		if dto.Error == "process ID not found" {
			c.JSON(http.StatusNotFound, dto)
			return
		}
		c.JSON(http.StatusInternalServerError, dto)
		return
	}

	c.JSON(http.StatusOK, dto)
	return
}

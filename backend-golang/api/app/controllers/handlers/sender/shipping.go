package sender

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/services"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/utils/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// NewShipping
// @Summary Creates a new shipping
// @Description Creates a new shipping instance in camunda and in the database
// @Consume application/json
// @Accept json
// @Produce json
// @Param Shipping body entities.ShippingInDTO true "Fill the body to create a new shipping"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /sender/shippings [post]
func NewShipping(c *gin.Context) {
	// TODO: Replace with generic should bind - TEST IF IT WORKS
	/*
		if err := c.ShouldBindJSON(&shipping); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": err.Error()})
			return
		}
	*/

	// TODO CREATE CAMUNDA INSTANCE
	dto := controllers.ShouldBindDTO(c, entities.ShippingInDTO{})

	var responseDto = services.CreateShipping(&dto)
	if !responseDto.Success {
		c.JSON(http.StatusInternalServerError, responseDto)
		return
	}
	c.JSON(http.StatusCreated, responseDto)
	return
}

// GetAllShippings
// @Summary Get all shippings
// @Description Get all shippings stored in the database
// @Produce json
// @Success 200 {object} dtos.Response
// @Failure 500 {object} dtos.Response
// @Router /sender/shippings [get]
func GetAllShippings(c *gin.Context) {
	var responseDto = services.FindAllShippings()
	if !responseDto.Success {
		c.JSON(http.StatusInternalServerError, responseDto)
		return
	}
	c.JSON(http.StatusOK, responseDto)
	return
}

// UpdateShippingByID
// @Summary Update Shipping
// @Description Update a Shipping by ID which is stored in the database
// @Consume application/json
// @Accept json
// @Produce json
// @Param id path int true "Shipping ID"
// @Param Shipping body entities.ShippingInPutDTO true "Fill the body to update a new shipping"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /sender/shippings/{id} [put]
func UpdateShippingByID(c *gin.Context) {
	id, err := controllers.ConvertParamToInt(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	dto := controllers.ShouldBindDTO(c, entities.ShippingInPutDTO{})
	var responseDto = services.UpdateShippingById(id, dto)
	if !responseDto.Success {
		c.JSON(http.StatusInternalServerError, responseDto)
	}
	c.JSON(http.StatusOK, responseDto)

	return
}

// DeleteShippingByID
// @Summary Delete Shipping
// @Description Delete a Shipping by ID which is stored in the database
// @Consume application/json
// @Accept json
// @Produce json
// @Param id path int true "Shipping ID"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /sender/shippings/{id} [delete]
func DeleteShippingByID(c *gin.Context) {

	var id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var responseDto = services.DeleteShippingById(id)
	if !responseDto.Success {
		c.JSON(http.StatusInternalServerError, responseDto)
	}
	c.JSON(http.StatusAccepted, responseDto)

	return
}

func GetShippingOffersByID(c *gin.Context) {

}

func GetShippingDistributorByID(c *gin.Context) {

}

func GetShippingRouteByID(c *gin.Context) {

}

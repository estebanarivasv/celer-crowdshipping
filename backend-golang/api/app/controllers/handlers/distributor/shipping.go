package distributor

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/services"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/utils/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

// NewShippingCoordinates
// @Summary Create new coordinates
// @Description Add new coordinates to a shipping and stores it in the database
// @Consume application/json
// @Accept json
// @Produce json
// @Param Offer body entities.NewRouteDetailInDTO true "Fill the body to update coordinates"
// @Param id path int true "Shipping ID"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Failure 500 {object} dtos.Response
// @Router /distributor/shippings/{id}/route [post]
func NewShippingCoordinates(c *gin.Context) {

	shippingId, err := controllers.ConvertParamToInt(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	dto := controllers.ShouldBindDTO(c, entities.RouteDetailInDTO{})
	responseDto := services.AddNewShippingCoordinates(shippingId, dto)
	if !responseDto.Success {
		c.JSON(http.StatusBadRequest, dto)
		return
	}

	c.JSON(http.StatusOK, dto)
	return

}

// GetShippingByID
// @Summary Get Shipping
// @Description Get Shipping stored in the database by passing an ID
// @Consume application/json
// @Accept json
// @Produce json
// @Param id path int true "Shipping ID"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Failure 404 {object} dtos.Response
// @Router /distributor/shippings/{id} [get]
func GetShippingByID(c *gin.Context) {
	// TODO: Verify user has access to this information - AUTH
	id, err := controllers.ConvertParamToInt(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	dto := services.FindShippingById(id)
	if !dto.Success {
		if dto.Error == "record not found" {
			c.JSON(http.StatusNotFound, dto)
			return
		}
		c.JSON(http.StatusInternalServerError, dto)
		return
	}
	c.JSON(http.StatusOK, dto)
	return
}

// UpdateShippingStateByID
// @Summary Update Shipping State
// @Description Change shipping state by sending a message to a camunda process
// @Consume application/json
// @Accept json
// @Produce json
// @Param id path int true "Shipping ID"
// @Param Shipping body dtos.MessageToProcessInDTO true "Fill the body to change shipping state"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /distributor/shippings/{id} [patch]
func UpdateShippingStateByID(c *gin.Context) {

	id, err := controllers.ConvertParamToInt(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	dto := controllers.ShouldBindDTO(c, dtos.MessageToProcessInDTO{})

	// Verify message is valid for sender
	valid := controllers.ContainsValidDistributorMsg(dto.MessageName)
	if !valid {
		c.JSON(
			http.StatusBadRequest,
			dtos.Response{Success: false, Error: "not a valid message"})
		return
	}

	responseDto := services.UpdateShippingState(id, dto.MessageName)
	if !responseDto.Success {
		c.JSON(http.StatusInternalServerError, responseDto)
		return
	}

	c.JSON(http.StatusCreated, responseDto)
	return
}

// GetShippingsInProcess
// @Summary Get shippings in progress
// @Description Get shippings that started their track into the recipient
// @Consume application/json
// @Accept json
// @Produce json
// @Success 200 {object} dtos.Response
// @Failure 500 {object} dtos.Response
// @Router /distributor/shippings [get]
func GetShippingsInProcess(c *gin.Context) {

	responseDto := services.FindActiveShippings()
	if !responseDto.Success {
		c.JSON(http.StatusInternalServerError, responseDto)
		return
	}

	c.JSON(http.StatusOK, responseDto)
	return

}

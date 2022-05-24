package sender

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/services"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/utils/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
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
// @Router /sender/shippings/{id} [get]
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

	id, err := controllers.ConvertParamToInt(c.Param("id"))
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
// @Router /sender/shippings/{id} [patch]
func UpdateShippingStateByID(c *gin.Context) {

	id, err := controllers.ConvertParamToInt(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	dto := controllers.ShouldBindDTO(c, dtos.MessageToProcessInDTO{})

	// Verify message is valid for sender
	valid := controllers.ContainsValidSenderMsg(dto.MessageName)
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

// UpdateShippingStateByID
// @Summary Update Shipping's selected offer
// @Description Add selected offer and send a message to a camunda process
// @Consume application/json
// @Accept json
// @Produce json
// @Param id path int true "Shipping ID"
// @Param Shipping body entities.ShippingInPatchDTO true "Add selected offer ID"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /sender/shippings/{id}/offers/selected [patch]
func UpdateSelectedOfferByID(c *gin.Context) {

	id, err := controllers.ConvertParamToInt(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	dto := controllers.ShouldBindDTO(c, entities.ShippingInPatchDTO{})

	responseDto := services.UpdateSelectedOffer(id, dto)

	if !responseDto.Success {
		c.JSON(http.StatusInternalServerError, responseDto)
		return
	}
	c.JSON(http.StatusCreated, responseDto)
	return
}

func GetShippingOffersByID(c *gin.Context) {

}

func GetShippingDistributorByID(c *gin.Context) {

}

func GetShippingRouteByID(c *gin.Context) {

}

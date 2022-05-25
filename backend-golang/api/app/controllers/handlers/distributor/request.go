package distributor

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/services"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/utils/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SearchRequests
// @Summary Get all pending shipping requests
// @Description Get all requests which are stored in the database and expect offers
// @Produce json
// @Success 200 {object} dtos.Response
// @Failure 500 {object} dtos.Response
// @Router /distributor/requests [get]
func SearchRequests(c *gin.Context) {

	responseDto := services.FindShippingRequests()
	if !responseDto.Success {
		c.JSON(http.StatusInternalServerError, responseDto)
		return
	}
	c.JSON(http.StatusOK, responseDto)
	return
}

// NewRequestOffer
// @Summary Creates a new offer for a shipping request
// @Description Creates a new shipping request offer instance and stores it in the database
// @Consume application/json
// @Accept json
// @Produce json
// @Param Offer body entities.NewOfferInDTO true "Fill the body to create a new shipping request offer"
// @Param id path int true "Shipping request ID"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Failure 500 {object} dtos.Response
// @Router /distributor/requests/{id}/offers [post]
func NewRequestOffer(c *gin.Context) {

	shippingID, err := controllers.ConvertParamToInt(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	dto := controllers.ShouldBindDTO(c, entities.OfferInDTO{})
	dto.ShippingID = shippingID

	responseDto := services.CreateOffer(&dto)
	if !responseDto.Success {
		if responseDto.Error == "this shipping id does not refer to a shipping request" {
			c.JSON(http.StatusBadRequest, responseDto)
			return
		}
		c.JSON(http.StatusInternalServerError, responseDto)
		return
	}
	c.JSON(http.StatusCreated, responseDto)
	return

}

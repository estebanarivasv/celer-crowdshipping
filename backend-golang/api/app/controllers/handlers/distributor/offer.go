package distributor

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/services"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/utils/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllOffersByDistributor(c *gin.Context) {

}

// GetRequestOffers
// @Summary Get request offers
// @Description Get offers associated to a request from the database by passing an ID
// @Consume application/json
// @Accept json
// @Produce json
// @Param id path int true "Shipping request ID"
// @Success 200 {object} dtos.Response
// @Failure 500 {object} dtos.Response
// @Router /distributor/requests/{id}/offers [get]
func GetRequestOffers(c *gin.Context) {
	shippingId, err := controllers.ConvertParamToInt(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	dto := services.FindOffersByRequestID(shippingId)
	if !dto.Success {
		c.JSON(http.StatusBadRequest, dto)
		return
	}

	c.JSON(http.StatusOK, dto)
	return

}

// DeleteOfferByID
// @Summary Delete offer
// @Description Delete offer from the database by passing an ID
// @Consume application/json
// @Accept json
// @Produce json
// @Param id path int true "Offer ID"
// @Success 204 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Failure 500 {object} dtos.Response
// @Router /distributor/requests/offers/{id} [delete]
func DeleteOfferByID(c *gin.Context) {

	offerId, err := controllers.ConvertParamToInt(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	dto := services.DeleteOfferById(offerId)
	if !dto.Success {
		c.JSON(http.StatusBadRequest, dto)
		return
	}

	c.JSON(http.StatusNoContent, dto)
	return

}

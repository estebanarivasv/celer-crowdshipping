package distributor

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/services"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/utils/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OfferController struct {
	offerService *services.OfferService
}

// NewOfferController Returns a new instance
func NewOfferController() *OfferController {

	var offerService = services.NewOfferService()

	return &OfferController{
		offerService: offerService,
	}
}

// GetAllOffersByDistributor
// @Summary Get offers made by a distributor
// @Description Get offers associated to a distributor from the database by passing an ID
// @Consume application/json
// @Produce json
// @Security ApiKeyAuth
// @Param distributorId query int true "Distributor ID"
// @Success 200 {object} dtos.Response
// @Failure 500 {object} dtos.Response
// @Router /distributor/offers [get]
func (c *OfferController) GetAllOffersByDistributor(context *gin.Context) {
	distributorId, err := controllers.ConvertParamToInt(context.Query("distributorId"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	dto := c.offerService.FindOffersByDistributorID(distributorId)
	if !dto.Success {
		context.JSON(http.StatusBadRequest, dto)
		return
	}

	context.JSON(http.StatusOK, dto)
	return

}

// DeleteOfferByID
// @Summary Delete offer
// @Description Delete offer from the database by passing an ID
// @Consume application/json
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Offer ID"
// @Success 204 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Failure 500 {object} dtos.Response
// @Router /distributor/offers/{id} [delete]
func (c *OfferController) DeleteOfferByID(context *gin.Context) {

	offerId, err := controllers.ConvertParamToInt(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	dto := c.offerService.DeleteOfferById(offerId)
	if !dto.Success {
		context.JSON(http.StatusBadRequest, dto)
		return
	}

	context.JSON(http.StatusNoContent, dto)
	return

}

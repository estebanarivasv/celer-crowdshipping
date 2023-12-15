package distributor

import (
	"fmt"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/services"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/utils/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RequestController struct {
	shippingService *services.ShippingService
	offerService    *services.OfferService
}

// NewRequestController Returns a new instance
func NewRequestController() *RequestController {

	var shippingService = services.NewShippingService()
	var offerService = services.NewOfferService()

	return &RequestController{
		shippingService: shippingService,
		offerService:    offerService,
	}
}

// SearchRequests
// @Summary Get all pending shipping requests
// @Description Get all requests which are stored in the database and expect offers
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} dtos.Response
// @Failure 500 {object} dtos.Response
// @Router /distributor/requests [get]
func (c *RequestController) SearchRequests(context *gin.Context) {

	responseDto := c.shippingService.FindShippingRequests()
	if !responseDto.Success {
		context.JSON(http.StatusInternalServerError, responseDto)
		return
	}
	context.JSON(http.StatusOK, responseDto)
	return
}

// NewRequestOffer
// @Summary Creates a new offer for a shipping request
// @Description Creates a new shipping request offer instance and stores it in the database
// @Consume application/json
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Offer body entities.NewOfferInDTO true "Fill the body to create a new shipping request offer"
// @Param id path int true "Shipping request ID"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Failure 500 {object} dtos.Response
// @Router /distributor/requests/{id}/offers [post]
func (c *RequestController) NewRequestOffer(context *gin.Context) {

	// Get user_id from context
	userID, exists := context.Get("user_id")
	if !exists {
		// Handle in case user does not exist in the context
		context.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't retrieve user information"})
		return
	}

	stringifiedUserID := fmt.Sprintf("%v", userID)
	parsedUserID, err := strconv.Atoi(stringifiedUserID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't retrieve user information"})
		return
	}

	shippingID, err := controllers.ConvertParamToInt(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	dto := controllers.ShouldBindDTO(context, entities.OfferInDTO{})
	dto.ShippingID = shippingID

	responseDto := c.offerService.CreateOffer(&dto, parsedUserID)
	if !responseDto.Success {
		if responseDto.Error == "this shipping id does not refer to a shipping request" {
			context.JSON(http.StatusBadRequest, responseDto)
			return
		}
		context.JSON(http.StatusInternalServerError, responseDto)
		return
	}
	context.JSON(http.StatusCreated, responseDto)
	return

}

// GetRequestOffers
// @Summary Get request offers
// @Description Get offers associated to a request from the database by passing an ID
// @Consume application/json
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Shipping request ID"
// @Success 200 {object} dtos.Response
// @Failure 500 {object} dtos.Response
// @Router /distributor/requests/{id}/offers [get]
func (c *OfferController) GetRequestOffers(context *gin.Context) {
	shippingId, err := controllers.ConvertParamToInt(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	dto := c.offerService.FindOffersByRequestID(shippingId)
	if !dto.Success {
		context.JSON(http.StatusBadRequest, dto)
		return
	}

	context.JSON(http.StatusOK, dto)
	return

}

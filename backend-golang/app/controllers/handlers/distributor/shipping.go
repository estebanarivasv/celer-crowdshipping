package distributor

import (
	"fmt"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/services"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/utils/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DistShippingController struct {
	shippingService *services.ShippingService
	routeService    *services.RouteService
}

// NewDistShippingController Returns a new instance
func NewDistShippingController() *DistShippingController {

	var shippingService = services.NewShippingService()
	var routeService = services.NewRouteService()

	return &DistShippingController{
		shippingService: shippingService,
		routeService:    routeService,
	}
}

// NewShippingCoordinates
// @Summary Create new coordinates
// @Description Add new coordinates to a shipping and stores it in the database
// @Consume application/json
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Offer body entities.NewRouteDetailInDTO true "Fill the body to update coordinates"
// @Param id path int true "Shipping ID"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Failure 500 {object} dtos.Response
// @Router /distributor/shippings/{id}/route [post]
func (c *DistShippingController) NewShippingCoordinates(context *gin.Context) {

	shippingId, err := controllers.ConvertParamToInt(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	dto := controllers.ShouldBindDTO(context, entities.RouteDetailInDTO{})
	responseDto := c.routeService.AddNewShippingCoordinates(shippingId, dto)
	if !responseDto.Success {
		context.JSON(http.StatusBadRequest, dto)
		return
	}

	context.JSON(http.StatusOK, dto)
	return

}

// GetShippingByID
// @Summary Get Shipping
// @Description Get Shipping stored in the database by passing an ID
// @Consume application/json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Shipping ID"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Failure 404 {object} dtos.Response
// @Router /distributor/shippings/{id} [get]
func (c *DistShippingController) GetShippingByID(context *gin.Context) {

	id, err := controllers.ConvertParamToInt(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	dto := c.shippingService.FindShippingById(id)
	if !dto.Success {
		if dto.Error == "record not found" {
			context.JSON(http.StatusNotFound, dto)
			return
		}
		context.JSON(http.StatusInternalServerError, dto)
		return
	}
	context.JSON(http.StatusOK, dto)
	return
}

// UpdateShippingStateByID
// @Summary Update Shipping State
// @Description Change shipping state by sending a message to a camunda process
// @Consume application/json
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Shipping ID"
// @Param Shipping body dtos.MessageToProcessInDTO true "Fill the body to change shipping state"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /distributor/shippings/{id} [patch]
func (c *DistShippingController) UpdateShippingStateByID(context *gin.Context) {

	id, err := controllers.ConvertParamToInt(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	dto := controllers.ShouldBindDTO(context, dtos.MessageToProcessInDTO{})

	// Verify message is valid for sender
	valid := controllers.ContainsValidDistributorMsg(dto.MessageName)
	if !valid {
		context.JSON(
			http.StatusBadRequest,
			dtos.Response{Success: false, Error: "not a valid message"})
		return
	}

	responseDto := c.shippingService.UpdateShippingState(id, dto.MessageName)
	if !responseDto.Success {
		context.JSON(http.StatusInternalServerError, responseDto)
		return
	}

	context.JSON(http.StatusCreated, responseDto)
	return
}

// GetShippingsInProcessByDistributor
// @Summary Get shippings in progress
// @Description Get shippings that started their track into the recipient
// @Consume application/json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} dtos.Response
// @Failure 500 {object} dtos.Response
// @Router /distributor/shippings [get]
func (c *DistShippingController) GetShippingsInProcessByDistributor(context *gin.Context) {

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

	responseDto := c.shippingService.FindActiveShippingsByDistributorID(parsedUserID)
	if !responseDto.Success {
		context.JSON(http.StatusInternalServerError, responseDto)
		return
	}

	context.JSON(http.StatusOK, responseDto)
	return

}

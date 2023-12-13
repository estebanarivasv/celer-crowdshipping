package sender

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

type SendShippingController struct {
	shippingService *services.ShippingService
	routeService    *services.RouteService
	offerService    *services.OfferService
}

// NewSendShippingController Returns a new instance
func NewSendShippingController() *SendShippingController {

	var shippingService = services.NewShippingService()
	var routeService = services.NewRouteService()

	return &SendShippingController{
		shippingService: shippingService,
		routeService:    routeService,
	}
}

// NewShipping
// @Summary Creates a new shipping
// @Description Creates a new shipping instance in camunda and in the database
// @Consume application/json
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Shipping body entities.ShippingInDTO true "Fill the body to create a new shipping"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /sender/shippings [post]
func (c *SendShippingController) NewShipping(context *gin.Context) {
	dto := controllers.ShouldBindDTO(context, entities.ShippingInDTO{})

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

	var responseDto = c.shippingService.CreateShipping(&dto, parsedUserID)
	if !responseDto.Success {
		context.JSON(http.StatusInternalServerError, responseDto)
		return
	}
	context.JSON(http.StatusCreated, responseDto)
	return
}

// GetAllShippings
// @Summary Get all shippings
// @Description Get all shippings stored in the database
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} dtos.Response
// @Failure 500 {object} dtos.Response
// @Router /sender/shippings [get]
func (c *SendShippingController) GetAllShippings(context *gin.Context) {
	var responseDto = c.shippingService.FindAllShippings()
	if !responseDto.Success {
		context.JSON(http.StatusInternalServerError, responseDto)
		return
	}
	context.JSON(http.StatusOK, responseDto)
	return
}

// GetShippingByID
// @Summary Get Shipping
// @Description Get Shipping stored in the database by passing an ID
// @Consume application/json
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Shipping ID"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Failure 404 {object} dtos.Response
// @Router /sender/shippings/{id} [get]
func (c *SendShippingController) GetShippingByID(context *gin.Context) {
	// TODO: Verify user has access to this information - AUTH
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

// DeleteShippingByID
// @Summary Delete Shipping
// @Description Delete a Shipping by ID which is stored in the database
// @Consume application/json
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Shipping ID"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /sender/shippings/{id} [delete]
func (c *SendShippingController) DeleteShippingByID(context *gin.Context) {

	id, err := controllers.ConvertParamToInt(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	var responseDto = c.shippingService.DeleteShippingById(id)
	if !responseDto.Success {
		context.JSON(http.StatusInternalServerError, responseDto)
	}
	context.JSON(http.StatusAccepted, responseDto)

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
// @Router /sender/shippings/{id} [patch]
func (c *SendShippingController) UpdateShippingStateByID(context *gin.Context) {

	id, err := controllers.ConvertParamToInt(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	dto := controllers.ShouldBindDTO(context, dtos.MessageToProcessInDTO{})

	// Verify message is valid for sender
	valid := controllers.ContainsValidSenderMsg(dto.MessageName)
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

// UpdateSelectedOfferByID
// @Summary Update Shipping's selected offer
// @Description Add selected offer and send a message to a camunda process
// @Consume application/json
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Shipping ID"
// @Param Shipping body entities.ShippingInPatchDTO true "Add selected offer ID"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /sender/shippings/{id}/offers/selected [patch]
func (c *SendShippingController) UpdateSelectedOfferByID(context *gin.Context) {

	id, err := controllers.ConvertParamToInt(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	dto := controllers.ShouldBindDTO(context, entities.ShippingInPatchDTO{})

	responseDto := c.shippingService.UpdateSelectedOffer(id, dto)

	if !responseDto.Success {
		context.JSON(http.StatusBadRequest, responseDto)
		return
	}
	context.JSON(http.StatusCreated, responseDto)
	return
}

// GetShippingOffersByID
// @Summary Get shipping offers
// @Description Get offers associated to a shipping from the database by passing an ID
// @Consume application/json
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Shipping ID"
// @Success 200 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Failure 500 {object} dtos.Response
// @Router /sender/shippings/{id}/offers [get]
func (c *SendShippingController) GetShippingOffersByID(context *gin.Context) {

	shippingId, err := controllers.ConvertParamToInt(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	dto := c.offerService.FindOffersByShippingID(shippingId)
	if !dto.Success {
		context.JSON(http.StatusBadRequest, dto)
		return
	}

	context.JSON(http.StatusOK, dto)
	return
}

// GetSelectedOfferByID
// @Summary Get Shipping selected offer
// @Description Get a shipping selected route from the database by passing an ID
// @Consume application/json
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Shipping ID"
// @Success 202 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Failure 500 {object} dtos.Response
// @Router /sender/shippings/{id}/offers/selected [get]
func (c *SendShippingController) GetSelectedOfferByID(context *gin.Context) {

	shippingId, err := controllers.ConvertParamToInt(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	responseDto := c.offerService.FindSelectedOfferByShippingId(shippingId)
	if !responseDto.Success {
		context.JSON(http.StatusBadRequest, responseDto)
		return
	}

	context.JSON(http.StatusAccepted, responseDto)
	return

}

// GetShippingRouteByID
// @Summary Get Shipping route
// @Description Get a shipping route from the database by passing an ID
// @Consume application/json
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Shipping ID"
// @Success 202 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Failure 500 {object} dtos.Response
// @Router /sender/shippings/{id}/route [get]
func (c *SendShippingController) GetShippingRouteByID(context *gin.Context) {

	id, err := controllers.ConvertParamToInt(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	var responseDto = c.routeService.FindRouteByShippingID(id)
	if !responseDto.Success {
		context.JSON(http.StatusInternalServerError, responseDto)
	}
	context.JSON(http.StatusAccepted, responseDto)

	return
}

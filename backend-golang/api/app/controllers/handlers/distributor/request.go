package distributor

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/services"
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

func NewRequestOffer(c *gin.Context) {

}

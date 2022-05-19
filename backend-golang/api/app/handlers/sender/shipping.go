package sender

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/models"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// NewShippingHandler
// @Summary Creates a new shipping
// @Description Creates a new shipping instance in camunda and in the database
// @Consume application/json
// @Accept json
// @Produce json
// -		object			model				description			TODO create dto
// @Param shippings body dtos.Response true "shipping info"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /sender/shippings [post]
func NewShippingHandler(c *gin.Context) {
	var shipping models.Shipping
	if err := c.ShouldBindJSON(&shipping); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}

	var dao = services.CreateShipping(&shipping)
	if !dao.Success {
		c.JSON(http.StatusInternalServerError, dao.Error)
		return
	}
	c.JSON(http.StatusCreated, dao.Data)
	return
}

func GetAllShippingsHandler(c *gin.Context) {
	var dao = services.FindAllShippings()
	if !dao.Success {
		c.JSON(http.StatusInternalServerError, dao.Error)
		return
	}
	c.JSON(http.StatusOK, dao.Data)
	return
}

func GetShippingByIDHandler(c *gin.Context) {
	var id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var dao = services.FindShippingById(id)
	if !dao.Success {
		if dao.Error == "record not found" {
			c.JSON(http.StatusNotFound, dao.Error)
			return
		}
		c.JSON(http.StatusInternalServerError, dao.Error)
		return
	}
	c.JSON(http.StatusOK, dao.Data)
	return
}

func UpdateShippingByIDHandler(c *gin.Context) {
	var id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var newShipping models.Shipping
	if err := c.ShouldBindJSON(&newShipping); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}

	var dao = services.UpdateShippingById(id, &newShipping)
	if !dao.Success {
		c.JSON(http.StatusInternalServerError, dao.Error)
	}
	c.JSON(http.StatusOK, dao.Data)

	return
}

func DeleteShippingByIDHandler(c *gin.Context) {

	var id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var dao = services.DeleteShippingById(id)
	if !dao.Success {
		c.JSON(http.StatusInternalServerError, dao.Error)
	}
	c.JSON(http.StatusAccepted, dao.Data)

	return
}

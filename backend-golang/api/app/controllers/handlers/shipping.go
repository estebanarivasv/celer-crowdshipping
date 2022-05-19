package handlers

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetShippingByID(c *gin.Context) {
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

func UpdateShippingStateByID(c *gin.Context) {

}

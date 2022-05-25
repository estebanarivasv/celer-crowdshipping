package handlers

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/services"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/utils/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

// NewPackage
// @Summary Creates a new package
// @Description Creates a new package in the database
// @Consume application/json
// @Accept json
// @Produce json
// @Param Package body entities.PackageInDTO true "Fill the body to create a new package"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /packages [post]
func NewPackage(c *gin.Context) {
	dto := controllers.ShouldBindDTO(c, entities.PackageInDTO{})

	var responseDto = services.CreatePackage(&dto)
	if !responseDto.Success {
		c.JSON(http.StatusInternalServerError, responseDto)
		return
	}
	c.JSON(http.StatusCreated, responseDto)
	return
}

// DeletePackageByID
// @Summary Delete Package
// @Description Delete a Package by ID which is stored in the database
// @Consume application/json
// @Accept json
// @Produce json
// @Param id path int true "Package ID"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /packages/{id} [delete]
func DeletePackageByID(c *gin.Context) {

	id, err := controllers.ConvertParamToInt(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var responseDto = services.DeletePackageById(id)
	if !responseDto.Success {
		c.JSON(http.StatusInternalServerError, responseDto)
	}
	c.JSON(http.StatusAccepted, responseDto)

	return
}

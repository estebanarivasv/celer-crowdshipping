package handlers

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/services"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/utils/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PackageController struct {
	packageService *services.PackageService
}

// NewPackageController Returns a new instance
func NewPackageController() *PackageController {

	var service = services.NewPackageService()

	return &PackageController{
		packageService: service,
	}
}

// NewPackage
// @Summary Creates a new package
// @Description Creates a new package in the database
// @Consume application/json
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Package body entities.PackageInDTO true "Fill the body to create a new package"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /packages [post]
func (c *PackageController) NewPackage(context *gin.Context) {
	dto := controllers.ShouldBindDTO(context, entities.PackageInDTO{})

	var responseDto = c.packageService.CreatePackage(&dto)
	if !responseDto.Success {
		context.JSON(http.StatusInternalServerError, responseDto)
		return
	}
	context.JSON(http.StatusCreated, responseDto)
	return
}

// DeletePackageByID
// @Summary Delete Package
// @Description Delete a Package by ID which is stored in the database
// @Consume application/json
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Package ID"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /packages/{id} [delete]
func (c *PackageController) DeletePackageByID(context *gin.Context) {

	id, err := controllers.ConvertParamToInt(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	var responseDto = c.packageService.DeletePackageById(id)
	if !responseDto.Success {
		context.JSON(http.StatusInternalServerError, responseDto)
	}
	context.JSON(http.StatusAccepted, responseDto)

	return
}

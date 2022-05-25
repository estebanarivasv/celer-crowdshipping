package services

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/mappers"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/repositories"
)

var packageRepository = repositories.NewPackageRepository()
var packageMapper = mappers.PackageMapper{}

func CreatePackage(packageInDTO *entities.PackageInDTO) dtos.Response {

	// Convert the dto to an entity
	model := packageMapper.FromDTO(packageInDTO)

	query, err := packageRepository.Create(model)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Mapping into a dto
	dto := packageMapper.ToDTO(&query)

	return dtos.Response{Success: true, Data: dto}
}

func DeletePackageById(id int) dtos.Response {

	err := packageRepository.DeleteById(id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	return dtos.Response{Success: true}

}

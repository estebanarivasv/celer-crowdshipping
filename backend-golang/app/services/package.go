package services

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/mappers"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/repositories"
)

type PackageService struct {
	packageRepo *repositories.PackageRepository
	mapper      *mappers.PackageMapper
}

// NewPackageService Returns a new instance
func NewPackageService() *PackageService {

	var repo = repositories.NewPackageRepository()
	var mapper = mappers.NewPackageMapper()

	return &PackageService{
		packageRepo: repo,
		mapper:      mapper,
	}
}

func (s *PackageService) CreatePackage(packageInDTO *entities.PackageInDTO) dtos.Response {

	// Convert the dto to an entity
	model := s.mapper.FromDTO(packageInDTO)

	query, err := s.packageRepo.Create(model)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Mapping into a dto
	dto := s.mapper.ToDTO(&query)

	return dtos.Response{Success: true, Data: dto}
}

func (s *PackageService) DeletePackageById(id int) dtos.Response {

	err := s.packageRepo.DeleteById(id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	return dtos.Response{Success: true}

}
